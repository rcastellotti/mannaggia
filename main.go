package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/k3a/html2text"
)

func main() {
	_, m, d := time.Now().Date()
	res, err := http.Get(fmt.Sprintf("https://santiebeati.it/%02d/%d", m, d))
	if err != nil {
		panic(err)
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	left := "<FONT SIZE=\"-2\">"
	right := "</FONT>&nbsp;"
	rx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `(.*?)` + regexp.QuoteMeta(right))
	matches := rx.FindAllStringSubmatch(string(bodyBytes), -1)

	for _, k := range matches {
		strippedSaint := html2text.HTML2Text(k[0])
		fmt.Printf("%s%s\n", "\033[1mmannaggia \033[0m", strippedSaint)
		time.Sleep(500 * time.Millisecond)
	}
}
