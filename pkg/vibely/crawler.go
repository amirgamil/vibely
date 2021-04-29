package vibely

import (
	"log"

	"github.com/anaskhan96/soup"
)

func crawlGetSong(url string) string {
	resp, err := soup.Get(url)
	if err != nil {
		log.Printf("Error trying to crawl the song")
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "class", "lyrics").FindAll("a", "p")

	return ""
}
