package postgres

import (
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
	// load pq library with side effects
	_ "github.com/lib/pq"
)

func NewDB(url *url.URL, schema string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", url.String())
	if err != nil {
		return nil, err
	} else if schema != "" {
		_, err := db.Exec(fmt.Sprintf(`set search_path='%s'`, schema))
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
