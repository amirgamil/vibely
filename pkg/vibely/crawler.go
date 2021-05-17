package vibely

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rhnvrm/lyric-api-go/goquery_helpers"
)

func crawlGetSong(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	defer resp.Body.Close()
	var songFormatted string
	if resp.StatusCode != 200 {
		log.Println("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("ERROR CREATING DOCUMENT! ", err)
		songFormatted = "Error getting the song lyrics"
	}
	result := document.Find(".Lyrics__Container-sc-1ynbvzw-6")
	// for i := range result.Nodes {
	// 	className, _ := result.Eq(i).Attr("class")
	// 	fmt.Println(className)
	// }
	songFormatted = strings.TrimSpace(goquery_helpers.RenderSelection(result, "\n"))
	r, _ := regexp.Compile("[\\(\\[].*?[\\)\\]]")
	songFormatted = r.ReplaceAllString(songFormatted, "")
	return songFormatted
}
