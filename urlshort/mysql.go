package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // New import
)

// SnippetModel wraps a sql.DB connection pool
type gomysql struct {
	DB *sql.DB
}

// ErrorNoRecord will be the error we will return
// if there is no match
var ErrorNoRecord = errors.New("models: no matching record found")

// Snippet is an object
type dataModel struct {
	path string
	url  string
}

type config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

func newMysql() (*gomysql, error) {
	conf, err := getConfig()
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("%s:%s@/%s?parseTime=true",
		conf.Username, conf.Password, "")
	log.Println("dsn: ", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err = db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+conf.Dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return nil, err
	}
	dsn = fmt.Sprintf("%s:%s@/%s?parseTime=true",
		conf.Username, conf.Password, conf.Dbname)
	log.Println("dsn: ", dsn)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &gomysql{DB: db}, nil
}

func getConfig() (*config, error) {
	f, err := ioutil.ReadFile("./.config.json")
	if err != nil {
		return nil, err
	}
	c := &config{}
	err = json.Unmarshal(f, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Insert inserts anew snippet into database
func (m *gomysql) Insert() (int, error) {
	// stmt := `INSERT INTO snippets (title, content, created, expires)
	// VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// result, err := m.DB.Exec(stmt, title, content, expires)
	// if err != nil {
	// 	return 0, nil
	// }
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	return 0, err
	// }
	// return int(id), nil
	return 0, nil
}

// Get returns the spesific snippet we want based on its id
func (m *gomysql) getData(path string) (string, bool) {
	stmt := `SELECT path, url FROM dataurl WHERE path = ?`

	row := m.DB.QueryRow(stmt, path)
	s := &dataModel{}
	err := row.Scan(&s.path, &s.url)
	log.Println("gomysql s: ", s)
	if err == sql.ErrNoRows {
		return "", false
	} else if err != nil {
		return "", false
	}

	return s.url, true
}
