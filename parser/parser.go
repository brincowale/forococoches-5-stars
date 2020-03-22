package parser

import (
	"forococoches-5-stars/models"
	"github.com/gocolly/colly"
	"regexp"
	"strconv"
)

var threadSelector = "//td[starts-with(@id, 'td_threadtitle_')]"
var threadTitle = "//a[starts-with(@id, 'thread_title_')]"
var ratingSelector = "//img[@src='//st.forocoches.com/foro/images/rating/rating_5.gif']"
var regexId, _ = regexp.Compile(`(\d+$)`)

func Parse(subForumURL string) []models.Thread {
	c := colly.NewCollector()
	var threads []models.Thread
	c.OnXML(threadSelector, func(e *colly.XMLElement) {
		if e.ChildAttr(ratingSelector, "src") != "" {
			threadId, _ := strconv.Atoi(regexId.FindString(e.ChildAttr(threadTitle, "href")))
			thread := models.Thread{
				Title: e.ChildText(threadTitle),
				URL:   "https://www.forocoches.com/foro/" + e.ChildAttr(threadTitle, "href"),
				Id:    threadId,
			}
			threads = append(threads, thread)
		}
	})
	c.Visit(subForumURL)
	return threads
}
