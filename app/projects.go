package app

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type project struct {
	Name      string    `db:"name"`
	Lang      string    `db:"lang"`
	Token     string    `db:"token"`
	Coverage  float64   `db:"coverage"`
	Timestamp time.Time `db:"timestamp"`
	Raw       string    `db:"raw"`
}

func Push(conn *sqlx.DB, name, lang, token, raw string) error {
	p := project{
		Name:      name,
		Lang:      lang,
		Token:     token,
		Coverage:  0,
		Timestamp: time.Now(),
		Raw:       raw,
	}
	switch p.Lang {
	case "golang":
		p.Coverage = parseGolangData(p.Raw)
	}
	r, err := conn.NamedExec("UPDATE project SET `lang`=:lang, `token`=:token, `coverage`=:coverage, `timestamp`=:timestamp, `raw`=:raw WHERE `name`=:name AND (`token`=:token OR `token`='') LIMIT 1", p)
	if err != nil {
		return err
	}
	a, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if a == 0 {
		_, err = conn.NamedExec("INSERT INTO project (`name`, `lang`, `token`, `coverage`, `timestamp`, `raw`) VALUES (:name, :lang, :token, :coverage, :timestamp, :raw)", p)
	}
	return err
}

func Get(conn *sqlx.DB, name string) (float64, error) {
	var coverage float64
	return coverage, conn.Get(&coverage, "SELECT `coverage` FROM project WHERE `name`=? LIMIT 1", name)
}
