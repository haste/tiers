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
	return p.progressValues[i].Expected < p.progressValues[j].Expected
}

type view struct {
	User int

	NextLevel uint

	Next []map[string]string

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

	profiles := model.GetAllProfiles(view.User)

	if len(profiles) == 0 {
		view.Warning = "Not enough data to generate page. :("
		templates.ExecuteTemplate(w, "progress", &view)
		return
	}

	var x []int64
	var y = map[string][]int64{}

	for k := range profiles {
		p := &profiles[k]
		profile.HandleProfile(p)
		x = append(x, p.Timestamp)
		y["AP"] = append(y["AP"], p.AP)
	}

	newestProfile := profiles[len(profiles)-1]
	if newestProfile.AP < newestProfile.NextLevel.AP {
		slope, intercept, _ := lineaRegression(x, y["AP"])
		ts := int64(((float64(newestProfile.NextLevel.AP) - intercept) / slope)) + 1262304000

		perc := (float64(int(float64(newestProfile.AP)/float64(newestProfile.NextLevel.AP)*1000)) / 10)
		view.Progress = append(view.Progress, progress{
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
		})
	} else {
		view.Progress = append(view.Progress, progress{
			Name: "AP",

			Tiers: []tier{
				{
					Name:       "info",
					Percentage: 100,
				},
			},

			Current: newestProfile.AP,
		})
	}

	var badges []progress

	for badgeName := range profile.BadgeRanks {
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

				if badge.Name == name {
					y[name] = append(y[name], f.FieldByName("Current").Int())
				}
			}
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

				if y[name][0] == y[name][len(y[name])-1] {
					badge.Expected = time.Now().AddDate(20, 0, 0).Unix()
				} else if badge.Current < badge.Required {
					slope, intercept, _ := lineaRegression(x, y[badge.Name])
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

	var priority []progress
	var inprogress []progress
	var completed []progress
	for _, v := range view.Progress {
		switch v.Expected {
		case -1:
			completed = append(completed, v)
		case 0:
			priority = append(priority, v)
		default:
			inprogress = append(inprogress, v)
		}
	}

	sort.Sort(progressByExpected{priority})
	sort.Sort(progressByExpected{inprogress})
	sort.Sort(progressByExpected{completed})

	view.Progress = append(priority, inprogress...)
	view.Completed = completed

	/*
		badges := profile.BuildBadgeProgress(newestProfile)
		for _, badge := range badges {
			fmt.Println(badge)
			var total float64
			var tiers []tier

			max := float64(badge.Ranges[len(badge.Ranges)-1])
			for rank, value := range badge.Ranges {
				perc := (float64(int(float64(value)/max*1000)) / 10) - total

				tiers = append(tiers, tier{
					Name:       rankNames[rank],
					Percentage: perc,
				})

				total += perc
			}
		}
	*/

	/*
		var badges []progress
		if newestProfile.Bronze < newestProfile.NextLevel.Bronze {
			badges = append(badges, possibleBadges(newestProfile, 1)...)
		}

		if newestProfile.Silver < newestProfile.NextLevel.Silver {
			badges = append(badges, possibleBadges(newestProfile, 2)...)
		}

		if newestProfile.Gold < newestProfile.NextLevel.Gold {
			badges = append(badges, possibleBadges(newestProfile, 3)...)
		}

		if newestProfile.Platinum < newestProfile.NextLevel.Platinum {
			badges = append(badges, possibleBadges(newestProfile, 4)...)
		}

		if newestProfile.Onyx < newestProfile.NextLevel.Onyx {
			badges = append(badges, possibleBadges(newestProfile, 5)...)
		}

		if len(badges) > 0 {
			view.Progress = badges
		}

		/*
			var x []uint
			var y = map[string][]uint{}
			var badges []badgeProgress
			var ranks = []string{
				"", "Bronze", "Silver", "Gold", "Platinum", "Onyx",
			}

			for k := range profiles {
				p := &profiles[k]
				profile.HandleProfile(p)
				x = append(x, p.Timestamp)
				y["AP"] = append(y["AP"], p.AP)
			}

			newestProfile := profiles[len(profiles)-1]

			for _, badge := range badges {
				for _, p := range profiles {
					s := reflect.ValueOf(&p.Badges).Elem()
					typeOf := s.Type()

					for i := 0; i < s.NumField(); i++ {
						f := s.Field(i)
						name := typeOf.Field(i).Name

						if badge.Badge == name {
							y[name] = append(y[name], uint(f.FieldByName("Current").Uint()))
						}
					}
				}
			}

			for _, badge := range badges {
				slope, intercept, _ := lineaRegression(x, y[badge.Badge])
				ts := uint(((float64(badge.Required) - intercept) / slope)) + 1262304000

				viewData.Badges[badge.Badge] = progressBadge{
					ranks[badge.Rank],
					badge.Badge,
					badge.Current,
					badge.Required,
					ts,
				}
			}
	*/

	templates.ExecuteTemplate(w, "progress", &view)
}
