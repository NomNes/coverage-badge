package app

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func ParseBody(r *http.Request) (name, data, token, lang string, err error) {
	if err = r.ParseMultipartForm(32 << 20); err != nil {
		return
	}

	f, _, err := r.FormFile("data")
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	data = string(b)

	name = r.FormValue("name")
	token = r.FormValue("token")
	lang = r.FormValue("lang")
	return
}

func parseGolangData(data string) float64 {
	lines := strings.Split(data, "\n")
	var total, covered int
	for _, l := range lines {
		parts := strings.Split(l, ":")
		if len(parts) == 2 {
			items := strings.Split(parts[1], " ")
			if len(items) == 3 {
				t, err := strconv.Atoi(items[1])
				if err == nil {
					total += t
					if items[2] != "0" {
						covered += t
					}
				}
			}
		}
	}
	return (float64(covered) / float64(total)) * 100
}
