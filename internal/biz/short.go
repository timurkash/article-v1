package biz

import (
	"bytes"
	"errors"
	"fmt"
	"html"
	"strings"

	"github.com/grokify/html-strip-tags-go"
)

func GetTitleHtml(htmlString string) (string, string, error) {
	badArticleFormat := errors.New("bad article format")
	if !strings.HasPrefix(htmlString, "<h1>") {
		return "", "", badArticleFormat
	}
	pos := strings.Index(htmlString, "</h1>")
	if pos == -1 {
		return "", "", badArticleFormat
	}
	body := htmlString[pos+5:]
	title := htmlString[4:pos]
	title = strings.ReplaceAll(title, "&quot;", "\"")
	return title, body, nil
}

func GetTitleHtml0(title, html string) string {
	return fmt.Sprintf(fmt.Sprintf("<h1>%s</h1>%s", title, html))
}

func GetTextAndShort(htmlString string, limit int) string {
	htmlString = html.UnescapeString(htmlString)
	htmlString = strings.ReplaceAll(htmlString, "</p><p>", "</p> <p>") //Чтобы был разрыв между абзацами
	htmlString = strip.StripTags(htmlString)
	htmlString = strings.TrimSpace(htmlString)
	if limit == 0 {
		return htmlString
	}
	return Shorten(htmlString, limit)
}

func Shorten(str string, limit int) string {
	if len(str) < limit {
		return str
	}
	str = FirstN(str, limit-3) + "..."
	return str
}

func FirstN(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}

var (
	TAGS       = []string{"<h1>", "</h1>", "<p>", "</p>"}
	EMPTY_TAGS = make([]string, 0)
)

func init() {
	tags := ""
	for i, tag := range TAGS {
		tags = tags + tag
		if i%2 == 1 {
			EMPTY_TAGS = append(EMPTY_TAGS, tags)
			tags = ""
		}
	}
}

func shortenHtml(html string, limit int32) string {
	str := []rune(html)
	i, j, over := 0, 0, false
	total := make([]rune, len(str)+3)
	for {
		r := str[i:]
		offset := tag(r)
		if offset > 1 {
			total = append(total, r[:offset]...)
		} else if !over {
			total = append(total, r[0])
			j++
			if j == int(limit) {
				over = true
				total = append(total, []rune("…")...)
			}
		}
		i = i + offset
		if i >= len(str) {
			break
		}
	}
	totalBytes := bytes.Trim([]byte(string(total)), "\x00")
	totalStr := string(totalBytes)
	totalStr = DeleteImg(totalStr)
	totalStr = DeleteEmpty(totalStr)
	return totalStr
}

func DeleteImg(html string) string {
	for {
		if strings.Contains(html, "<img ") {
			p := strings.Index(html, "<img ")
			tail := html[p:]
			//fmt.Println(tail)
			q := strings.Index(tail, ">")

			html = html[:p] + tail[q+1:]
		} else {
			break
		}
	}
	return html
}

func DeleteEmpty(str string) string {
	str = strings.ReplaceAll(str, "\n", " ")
	str = strings.ReplaceAll(str, "&nbsp;", " ")
	for {
		var c bool
		str, c = deleteEmpty(str)
		if !c {
			break
		}
	}
	return str
}

func tag(r []rune) int {
	str := string(r)
	for _, tag := range TAGS {
		if strings.HasPrefix(str, tag) {
			return len(tag)
		}
	}
	return 1
}

func deleteEmpty(str string) (string, bool) {
	c := false
	if strings.Contains(str, "  ") {
		c = true
		str = strings.ReplaceAll(str, "  ", " ")
	}
	for _, emptyTags := range EMPTY_TAGS {
		if strings.Contains(str, emptyTags) {
			c = true
			str = strings.ReplaceAll(str, emptyTags, "")
		}
	}
	return str, c
}
