package html2text

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
)

func DefaultSkipTags() map[string]bool {
	k := make(map[string]bool)
	k["script"] = true
	k["link"] = true
	k["meta"] = true
	k["style"] = true
	k["noscript"] = true
	return k
}

func NewSettings() TexterSettings {
	conf := TexterSettings{}
	conf.SkipTags = DefaultSkipTags()
	conf.IncludeLinkUrls = true
	return conf
}

type TexterSettings struct {
	SkipTags        map[string]bool
	IncludeLinkUrls bool
}

func Html2Text(html string, conf *TexterSettings) (string, error) {
	buffer := bytes.NewBuffer([]byte(html))
	return Html2TextFromReader(buffer, conf)
}

func Html2TextFromReader(r io.Reader, conf *TexterSettings) (string, error) {
	d := html.NewTokenizer(r)

	lastToken := html.Token{}
	buffer := bytes.Buffer{}
	textAdded := false

	for {
		// token type
		tokenType := d.Next()
		if tokenType == html.ErrorToken {
			if d.Err() == io.EOF {
				break
			}

			return "", d.Err()
		}

		token := d.Token()

		if isTextTag(token) {
			if isStartTag(lastToken) || lastToken.Data == "" {
				_, ok := conf.SkipTags[lastToken.Data]
				if !ok {
					if conf.IncludeLinkUrls && isTagName(lastToken, "a") {
						buffer.WriteString(token.Data + " (" + getHrefUrl(lastToken) + ")")
					} else {
						buffer.WriteString(token.Data)
					}
					textAdded = true
				}
			}
			if isEndTag(lastToken) {
				buffer.WriteString(token.Data)
				textAdded = true
			}
		}

		if isEndTag(token) && textAdded {
			textAdded = false
			if isTagName(token, "p") {
				buffer.WriteString("\n")
			} else {
				buffer.WriteString(" ")
			}
		}

		if isTagName(token, "br") {
			buffer.WriteString("\n")
		}

		if isStartTag(token) && isTagName(token, "p") {
			buffer.WriteString("\n")
		}

		lastToken = token
	}

	return buffer.String(), nil
}

func getHrefUrl(token html.Token) string {
	if isStartTag(token) && isTagName(token, "a") {
		for _, i := range token.Attr {
			if i.Key == "href" {
				return i.Val
			}
		}
	}
	return ""
}

func isTextTag(token html.Token) bool {
	return token.Type == html.TextToken
}

func isTagName(token html.Token, tagName string) bool {
	return token.Data == tagName
}

func isEndTag(token html.Token) bool {
	return token.Type == html.EndTagToken
}

func isStartTag(token html.Token) bool {
	return token.Type == html.StartTagToken
}
