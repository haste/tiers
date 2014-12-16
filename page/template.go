package page

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/GeertJohan/go.rice"
)

func comma(n int64) string {
	var h []byte
	var s = strconv.Itoa(int(n))

	for i := len(s) - 1; i >= 0; i-- {
		o := len(s) - 1 - i
		if o%3 == 0 && o != 0 {
			h = append(h, ',')
		}

		h = append(h, s[i])
	}

	for i, j := 0, len(h)-1; i < j; i, j = i+1, j-1 {
		h[i], h[j] = h[j], h[i]
	}

	return string(h)
}

func round(dividend int64, divisor float64) int64 {
	return int64(math.Floor((float64(dividend) / divisor) + .5))
}

func relativeTime(ts int64) string {
	if ts == 0 {
		return "Never"
	}

	now := time.Now().Unix()

	var seconds int64
	if ts > now {
		seconds = ts - now
	} else {
		seconds = now - ts
	}

	minutes := round(seconds, 60)
	hours := round(minutes, 60)
	days := round(hours, 24)
	months := round(days, 30)
	years := round(months, 12)

	switch {
	case seconds < 45:
		return "A few seconds"
	case minutes == 1:
		return "A minute"
	case minutes < 45:
		return fmt.Sprintf("%d minutes", minutes)
	case hours == 1:
		return "An hour"
	case hours < 22:
		return fmt.Sprintf("%d hours", hours)
	case days == 1:
		return "A day"
	case days <= 26:
		return fmt.Sprintf("%d days", days)
	case months == 1:
		return "A month"
	case months <= 11:
		return fmt.Sprintf("%d months", months)
	case years == 1:
		return "A year"
	case days < 3650:
		return fmt.Sprintf("%d years", years)
	default:
		return "A good while"
	}
}

func relativeTimeAgo(ts int64) string {
	now := time.Now().Unix()
	s := relativeTime(ts)

	if ts > now {
		return s
	} else {
		return s + " ago"
	}
}

func loadTemplates(temps ...string) *template.Template {
	box, err := rice.FindBox("templates")
	if err != nil {
		log.Fatal(err)
	}

	templates := template.New("").Funcs(template.FuncMap{
		"comma":       comma,
		"relative":    relativeTime,
		"relativeAgo": relativeTimeAgo,
	})

	for _, temp := range temps {
		templateString, err := box.String(temp)
		if err != nil {
			log.Fatal(err)
		}

		_, err = templates.New(temp).Parse(templateString)
		if err != nil {
			log.Fatal(err)
		}

	}

	return templates
}
