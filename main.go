package main

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/NomNes/go-errors-sentry"
	"github.com/getsentry/sentry-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"coverage-badge/app"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Debug:            true,
		AttachStacktrace: true,
		BeforeSend:       errors.SentryBeforeSend,
	})
	if err != nil {
		panic(err)
	}
	defer sentry.Flush(2 * time.Second)
	defer func() {
		defer sentry.Recover()
		err := recover()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}()

	conn, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@(%s:3306)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASS"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DB")))
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = conn.Close()
	}()
	conn.SetMaxOpenConns(200)

	panic(http.ListenAndServe(":80", handler(func(w http.ResponseWriter, r *http.Request) error {
		if r.RequestURI == "/" {
			if r.Method != http.MethodPost {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return nil
			}
			name, data, token, lang, err := app.ParseBody(r)
			if err != nil {
				return err
			}
			err = app.Push(conn, name, lang, token, data)
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("OK"))
			return err
		}
		name := strings.Trim(r.RequestURI, "/")
		coverage, err := app.Get(conn, name)
		if err != nil {
			return err
		}
		svg := app.GetSvg(int(math.Round(coverage)))
		w.Header().Set("content-type", "image/svg+xml")
		_, err = w.Write([]byte(svg))
		if err != nil {
			return err
		}
		return nil
	})))
}

func handler(h func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
