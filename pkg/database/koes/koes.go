package koes

import (
	"log"
	"errors"
	"database/sql"
	"github.com/otofuto/powloan/pkg/database"
	//"../../database"
)

type Koes struct {
	Who string `json:"who"`
	Comment string `json:"comment"`
	CreatedAt string `json:"created_at"`
}

func (k * Koes) Insert() bool {
	db := database.Connect()
	defer db.Close()

	sql := "insert into koes (who, comment) values ($1, $2)"

	_, err := db.Exec(sql, &k.Who, &k.Comment)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func All() ([]Koes, error) {
	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("select who, comment, created_at from koes order by created_at desc")
	if err != nil {
		log.Println(err)
		return make([]Koes, 0), errors.New("select failed at koes.All()")
	}

	defer rows.Close()
	var ret []Koes
	for rows.Next() {
		var k Koes
		err = rows.Scan(&k.Who, &k.Comment, &k.CreatedAt)
		if err != nil {
			log.Println(err)
			return make([]Koes, 0), errors.New("row scan failed at koes.All()")
		}
		ret = append(ret, k)
	}
	return ret, nil
}