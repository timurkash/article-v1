package others

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/grokify/html-strip-tags-go"
	"html"
	"strings"
)

func GetTitleHtml(html string) (string, string, error) {
	badArticleFormat := errors.New("bad article format")
	if !strings.HasPrefix(html, "<h1>") {
		return "", "", badArticleFormat
	}
	pos := strings.Index(html, "</h1>")
	if pos == -1 {
		return "", "", badArticleFormat
	}
	body := html[pos+5:]
	title := html[4:pos]
	title = strings.ReplaceAll(title, "&quot;", "\"")
	return title, body, nil
}

func GetTitleHtml0(title, html string) string {
	return fmt.Sprintf(fmt.Sprintf("<h1>%s</h1>%s", title, html))
}

func GetText(htmlString string, limit int) string {
	htmlString = html.UnescapeString(htmlString)
	htmlString = strings.ReplaceAll(htmlString, "</p><p>", "</p> <p>") //Чтобы был разрыв между абзацами
	htmlString = strip.StripTags(htmlString)
	htmlString = strings.TrimSpace(htmlString)
	if limit == 0 {
		return htmlString
	}
	return Shorten(htmlString, limit)
}

var (
	TAGS      = []string{"<h1>", "</h1>", "<p>", "</p>"}
	EmptyTags = make([]string, 0)
)

func init() {
	tags := ""
	for i, tag := range TAGS {
		tags = tags + tag
		if i%2 == 1 {
			EmptyTags = append(EmptyTags, tags)
			tags = ""
		}
	}
}

func ShortenHtml(html string, limit int) string {
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
			if j == limit {
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

func tag(r []rune) int {
	str := string(r)
	for _, tag := range TAGS {
		if strings.HasPrefix(str, tag) {
			return len(tag)
		}
	}
	return 1
}
