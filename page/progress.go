package page

import (
	"fmt"
	"math"
	"net/http"
	"reflect"
	"sort"
	"strings"

	"tiers/model"
	"tiers/profile"
	"tiers/session"
	"time"
)

type tier struct {
	Name       string
	Percentage float64
}

type progress struct {
	Name string
	Rank int
	Icon string

	Tiers []tier

	Current  int64
	Required int64

	Expected int64
}

type progressValues []progress

func (p progressValues) Len() int {
	return len(p)
}

func (p progressValues) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type progressByExpected struct {
	progressValues
}

func (p progressByExpected) Less(i, j int) bool {
	var (
		a = p.progressValues[i].Expected
		b = p.progressValues[j].Expected
	)

	if a < 0 {
		return false
	} else if b < 0 {
		return true
	}

	return a < b
}

type progressByName struct {
	progressValues
}

func (p progressByName) Less(i, j int) bool {
	var (
		a = p.progressValues[i].Name
		b = p.progressValues[j].Name
	)

	return a < b
}

type view struct {
	User int

	Queue int

	NextLevel uint

	Next []map[string]string

	AP        progress
	Progress  []progress
	Completed []progress

	Warning string
}

var rankNames = []string{
	"bronze", "silver", "gold", "platinum", "onyx",
}

func lineaRegression(x, y []int64) (float64, float64, float64) {
	var (
		sum_x  int64 = 0
		sum_y  int64 = 0
		sum_xy int64 = 0
		sum_xx int64 = 0
		sum_yy int64 = 0
	)

	for i := 0; i < len(x); i++ {
		x := x[i] - 1262304000
		y := y[i]
		sum_x += x
		sum_y += y
		sum_xy += (x * y)
		sum_xx += (x * x)
		sum_yy += (y * y)
	}

	var (
		n        = float64(len(x))
		float_x  = float64(sum_x)
		float_y  = float64(sum_y)
		float_xy = float64(sum_xy)
		float_xx = float64(sum_xx)
		float_yy = float64(sum_yy)
	)

	slope := (n*float_xy - float_x*float_y) / (n*float_xx - float_x*float_x)
	intercept := (float_y - slope*float_x) / n
	r2 := math.Pow((n*float_xy-float_x*float_y)/math.Sqrt((n*float_xx-float_x*float_x)*(n*float_yy-float_y*float_y)), 2)

	return slope, intercept, r2
}

func shouldIncludeBadge(name string, timestamp, newY int64) bool {
	// The most random timestamp ever! Also known as the timestamp of the first
	// profile with mod deployed data.
	if name == "Engineer" && newY == 0 && timestamp < 1418672414 {
		return false
		// Badge was implemented around 22:30 CET, timestamp is form the last profile
		// submitted before it was public.
	} else if name == "Translator" && newY == 0 && timestamp < 1422564883 {
		return false
	}

	return true
}

func ProgressHandler(w http.ResponseWriter, r *http.Request) {
	templates := loadTemplates(
		"header.html",
		"footer.html",
		"nav.html",
		"progress.html",
	)

	session, _ := session.Get(r, "tiers")
	userid, ok := session.Values["user"]

	if !ok {
		http.Redirect(w, r, "/", 302)
		return
	}

	var view view
	view.User = userid.(int)

	limit := time.Now().Unix() - int64((time.Hour * 24 * 30).Seconds())
	profiles := model.GetAllProfiles(view.User, limit)

	if len(profiles) == 0 {
		view.Warning = "Not enough data to generate page. :("
		templates.ExecuteTemplate(w, "progress", &view)
		return
	}

	view.Queue = model.GetNumQueuedProfiles(userid.(int))

	var x = map[string][]int64{}
	var y = map[string][]int64{}

	for k := range profiles {
		p := &profiles[k]
		profile.HandleProfile(p)
		x["AP"] = append(x["AP"], p.Timestamp)
		y["AP"] = append(y["AP"], p.AP)
	}

	newestProfile := profiles[len(profiles)-1]
	if newestProfile.AP < newestProfile.NextLevel.AP {
		var ts int64

		if y["AP"][0] == y["AP"][len(y["AP"])-1] {
			ts = -1
		} else {
			slope, intercept, _ := lineaRegression(x["AP"], y["AP"])
			ts = int64(((float64(newestProfile.NextLevel.AP) - intercept) / slope)) + 1262304000
		}

		perc := (float64(int(float64(newestProfile.AP)/float64(newestProfile.NextLevel.AP)*1000)) / 10)
		view.AP = progress{
			Name: "AP",

			Tiers: []tier{
				{
					Name:       "info",
					Percentage: perc,
				},
			},

			Current:  newestProfile.AP,
			Required: newestProfile.NextLevel.AP,

			Expected: ts,
		}
	} else {
		view.AP = progress{
			Name: "AP",

			Tiers: []tier{
				{
					Name:       "info",
					Percentage: 100,
				},
			},

			Current:  newestProfile.AP,
			Required: newestProfile.NextLevel.AP,

			Expected: -2,
		}
	}

	var badges []progress

	for badgeName := range profile.BadgeRanks {
		if badgeName == "Innovator" {
			continue
		}

		badges = append(badges, progress{
			Name: badgeName,
		})
	}

	for _, badge := range badges {
		for _, p := range profiles {
			s := reflect.ValueOf(&p.Badges).Elem()
			typeOf := s.Type()

			for i := 0; i < s.NumField(); i++ {
				f := s.Field(i)
				name := typeOf.Field(i).Name
				current := f.FieldByName("Current").Int()

				if badge.Name == name && shouldIncludeBadge(badge.Name, p.Timestamp, current) {
					x[name] = append(x[name], p.Timestamp)
					y[name] = append(y[name], current)
				}
			}
		}

		if badge.Name == "Guardian" && len(x["Guardian"]) > 1 {
			var i int
			var n = len(x["Guardian"]) - 1

			xValue := x["Guardian"][n]
			yValue := y["Guardian"][n]

			for i = n - 1; i > 0; i-- {
				var curX = x["Guardian"][i]
				var curY = y["Guardian"][i]

				diff := xValue - curX
				if curY != yValue {
					if diff <= 86400 {
						xValue = curX
						yValue = curY
					}
				} else if diff > 86400 {
					i++
					break
				}
			}

			x["Guardian"] = x["Guardian"][i : n+1]
			y["Guardian"] = y["Guardian"][i : n+1]
		}

		s := reflect.ValueOf(&newestProfile.Badges).Elem()
		typeOf := s.Type()

		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			name := typeOf.Field(i).Name

			if badge.Name == name {
				badge.Rank = int(f.FieldByName("Rank").Int())
				badge.Current = f.FieldByName("Current").Int()
				badge.Required = f.FieldByName("Next").Int()

				if len(y[name]) > 0 && y[name][0] == y[name][len(y[name])-1] {
					badge.Expected = -2
				} else if badge.Current < badge.Required {
					slope, intercept, _ := lineaRegression(x[badge.Name], y[badge.Name])
					ts := int64(((float64(badge.Required) - intercept) / slope)) + 1262304000

					badge.Expected = ts
				} else {
					badge.Expected = -1
				}

				var total float64
				max := float64(badge.Required)
				for rank, value := range profile.BadgeRanks[name] {
					if rank <= badge.Rank {
						perc := (float64(int(float64(value)/max*1000)) / 10) - total

						badge.Tiers = append(badge.Tiers, tier{
							Name:       rankNames[rank],
							Percentage: perc,
						})

						total += perc
					}
				}

				if badge.Rank != 4 {
					badge.Tiers = append(badge.Tiers, tier{
						Name:       rankNames[badge.Rank+1],
						Percentage: (float64(int(float64(badge.Current)/max*1000)) / 10) - total,
					})
				}

				badge.Icon = fmt.Sprintf("images/badges/%s/%d.png", strings.ToLower(badge.Name), badge.Rank+1)
			}
		}

		view.NextLevel = newestProfile.NextLevel.Level
		view.Progress = append(view.Progress, badge)
	}

	var inprogress []progress
	var completed []progress
	for _, v := range view.Progress {
		switch v.Expected {
		case -1:
			completed = append(completed, v)
		default:
			inprogress = append(inprogress, v)
		}
	}

	sort.Sort(progressByExpected{inprogress})
	sort.Sort(progressByName{completed})

	view.Progress = inprogress
	view.Completed = completed

	templates.ExecuteTemplate(w, "progress", &view)
}
