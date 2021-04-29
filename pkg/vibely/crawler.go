package vibely

import (
	"fmt"
	"log"

	"github.com/anaskhan96/soup"
)

func crawlGetSong(url string) string {
	resp, err := soup.Get(url)
	if err != nil {
		log.Printf("Error trying to crawl the song")
	}
	doc := soup.HTMLParse(resp)
	links := doc.FindStrict("div", "class", "lyrics").Find("p")
	fmt.Println(links)
	return ""
}
