package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var spaces = regexp.MustCompile("\\s+")

func homepath(p string) string {
	if strings.HasPrefix(p, "~/") {
		// relative to home
		p := p[2:]
		u, _ := user.Current()
		return filepath.Join(u.HomeDir, p)
	}

	return p
}

func main() {
	dbpath := flag.String("db", "~/Library/Messages/chat.db", "path to user Messages database")
	service := flag.String("service", "SMS", "message service to query")
	pattern := flag.String("pattern", "Your verification code is %.", "message text pattern")
	filter := flag.Bool("filter", true, "filter out pattern")

	flag.Parse()

	db, err := sql.Open("sqlite3", homepath(*dbpath))
	if err != nil {
		log.Fatal(*dbpath, ": ", err)
	}

	defer db.Close()

	var text string

	err = db.QueryRow("SELECT text FROM message WHERE service=? AND text LIKE ? ORDER BY date DESC",
		*service, *pattern).Scan(&text)
	if err != nil {
		log.Fatal(err)
	}

	if *filter {
		re := strings.Replace(*pattern, "%", "(.+)", -1)
		re = spaces.ReplaceAllLiteralString(re, `\s+`)
		match := regexp.MustCompile(re).FindStringSubmatch(text)
		if len(match) > 1 {
			text = strings.Join(match[1:], " ")
		}
	}

	fmt.Println(text)
}
