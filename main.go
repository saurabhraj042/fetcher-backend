package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gocolly/colly"
)

type Notice struct {
	Link 	string	`json:"link"`
	Title	string	`json:"title"`
}

func EncodeToJson(opportunities *[]Notice, w io.Writer){
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")
	encoder.Encode(opportunities)
}

func ScrapeAmityPortal(URL string) []Notice {	
	collector := colly.NewCollector()

	opportunities := make([]Notice, 0)
	
	collector.OnHTML("ul[class]", func(e *colly.HTMLElement) {
		if e.Attr("class") == "notices" {
			e.ForEach("a", func(i int, h *colly.HTMLElement) {
				cur_link := h.Request.AbsoluteURL(h.Attr("href"))
				cur_title := h.Text
				
				curNotice := Notice{
					Link: cur_link,
					Title: cur_title,
				}

				opportunities = append(opportunities, curNotice)
			})
		}
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Request.URL)
	})

	collector.Visit(URL)

	return opportunities
}

func HomePage(w http.ResponseWriter, r *http.Request, notices []Notice){
	URL1 := "https://www.amity.edu/placement/upcoming-recruitment.asp"
	URL2 := "https://www.amity.edu/placement/recruitment-result.asp"
	
	notices = nil
	notices = append(notices, ScrapeAmityPortal(URL1)...)
	notices = append(notices, ScrapeAmityPortal(URL2)...)
	EncodeToJson(&notices, w)
}

func main(){
	port := os.Getenv("PORT")
	server := &http.Server{Addr: ":" + port}
	notices := make([]Notice, 0)

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		HomePage(rw, r, notices)
	})

	http.HandleFunc("/shutdown", func(rw http.ResponseWriter, r *http.Request) {
		server.Shutdown(context.Background())
	})

	log.Fatal(server.ListenAndServe())
}