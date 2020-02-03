package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/prologic/go-gopher"
	log "github.com/sirupsen/logrus"
	"jaytaylor.com/html2text"
)

type proxy struct{}

func (p *proxy) ServeGopher(w gopher.ResponseWriter, r *gopher.Request) {
	log.Infof("Selector: %s", r.Selector)
	url := strings.TrimPrefix(r.Selector, "/")
	url = fmt.Sprintf("https://%s", url)

	res, err := http.Get(url)
	if err != nil {
		msg := fmt.Sprintf("error fetching web resource %s: %s", url, err)
		log.WithError(err).WithField("url", url).Error(msg)
		w.WriteError(msg)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		msg := fmt.Sprintf("error reading web resource body: %s", err)
		log.WithError(err).WithField("url", url).Error(msg)
		w.WriteError(msg)
		return
	}

	html := string(body)
	text, err := html2text.FromString(html, html2text.Options{PrettyTables: true})
	if err != nil {
		msg := fmt.Sprintf("error converting html to text: %s", err)
		log.WithError(err).WithField("url", url).Error(msg)
		w.WriteError(msg)
		return
	}

	// TODO: Handle links
	// TODO: Write Info items
	w.Write([]byte(text))
}

func main() {
	log.Fatal(gopher.ListenAndServe(":7000", &proxy{}))
}
