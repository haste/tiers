package page

import (
	"html/template"
	"log"
	"strconv"

	"github.com/GeertJohan/go.rice"
)

func comma(n uint) string {
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

func loadTemplates(temps ...string) *template.Template {
	box, err := rice.FindBox("templates")
	if err != nil {
		log.Fatal(err)
	}

	templates := template.New("").Funcs(template.FuncMap{
		"comma": comma,
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
