// Dapatkan jumlah Star, Watch dari sebuah repository di github.com
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
)

type repo struct {
	url        string
	start      int
	watch      int
	lastCommit string
}

func writeHelp() {
	fmt.Println("repovalue.exe golang/go golang/proposal")
}

func (r repo) getPage() string {
	b, e := soup.Get(r.url)
	if e != nil {
		fmt.Println("Failed to get ", r.url, ":", e)
		return ""
	}
	return b
}

func (r *repo) setRepoValueBasedOn(page string) (bool, error) {
	doc := soup.HTMLParse(page)
	elem := doc.Find("ul", "class", "pagehead-actions")
	aSocialCount := elem.FindAll("a", "class", "social-count")
	var e error
	r.watch, e = strconv.Atoi(strings.TrimSpace(aSocialCount[0].Text()))
	if e != nil {
		return false, e
	}
	r.start, e = strconv.Atoi(strings.TrimSpace(aSocialCount[1].Text()))
	if e != nil {
		return false, e
	}
	return true, nil
}

func (r repo) toString() string {
	return "URL: " + r.url + " watch: " + string(r.watch) + " start: " + string(r.start)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		writeHelp()
		os.Exit(1)
	}
	repos := make([]repo, len(args)-1)
	for i := 1; i < len(args); i++ {
		repos[i-1].url = fmt.Sprintf("https://github.com/%s", args[i])
		fmt.Println("Menganalisa ", repos[i-1].url)

	}

}
