package page

import (
	"math"
	"net/http"
	"reflect"

	"tiers/model"
	"tiers/profile"
	"tiers/session"
)

type progressAP struct {
	Current  uint
	Required uint
	Missing  uint
	Expected uint
}

type progressBadge struct {
	RankName string
	Badge    string
	Current  uint
	Required uint
	Expected uint
}

type ProgressPage struct {
	User int

	Warning string
	AP      progressAP
	Badges  map[string]progressBadge
}

type badgeProgress struct {
	Badge    string
	Rank     int64
	Current  uint
	Required uint
}

func possibleBadges(p profile.Profile, rank int64) []badgeProgress {
	var b []badgeProgress

	s := reflect.ValueOf(&p.Badges).Elem()
	typeOf := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		name := typeOf.Field(i).Name

		if f.FieldByName("Rank").Int() < rank {
			b = append(b, badgeProgress{
				name,
				rank,
				uint(f.FieldByName("Current").Uint()),
				profile.BadgeRanks[name][rank-1],
			})
		}
	}

	return b
}

func lineaRegression(x, y []uint) (float64, float64, float64) {
	var (
		sum_x  uint = 0
		sum_y  uint = 0
		sum_xy uint = 0
		sum_xx uint = 0
		sum_yy uint = 0
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

	var viewData ProgressPage
	viewData.User = userid.(int)

	profiles := model.GetAllProfiles(viewData.User)

	if len(profiles) == 1 {
		viewData.Warning = "Not enough data to generate page. :("
		templates.ExecuteTemplate(w, "progress", viewData)
		return
	}

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

	if newestProfile.AP < newestProfile.NextLevel.AP {
		slope, intercept, _ := lineaRegression(x, y["AP"])
		ts := uint(((float64(newestProfile.NextLevel.AP) - intercept) / slope)) + 1262304000

		viewData.AP = progressAP{
			newestProfile.AP,
			newestProfile.NextLevel.AP,
			newestProfile.NextLevel.AP - newestProfile.AP,
			ts,
		}
	}

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
		viewData.Badges = make(map[string]progressBadge)
	}

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

	templates.ExecuteTemplate(w, "progress", &viewData)
}
