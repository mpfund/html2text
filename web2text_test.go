package web2text

import (
	"fmt"
	"strings"
	"testing"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func TestHtml2TextSkipTags(t *testing.T) {
	conf := NewSettings()
	conf.IncludeLinkUrls = false

	text, err := Html2Text("q<script>ddgft</script>mm", &conf)
	checkerr(err)
	if text != "q mm" {
		t.Fail()
	}
}

func TestHtml2FullHTML1(t *testing.T) {
	conf := NewSettings()
	conf.IncludeLinkUrls = true

	html := `<!DOCTYPE html><html><body>hallo</body></html>`
	text, err := Html2Text(html, &conf)
	checkerr(err)

	if strings.TrimSpace(text) != "hallo" {
		t.Fail()
	}

	fmt.Println(text)
}

func TestHtml2TextLinks(t *testing.T) {
	conf := NewSettings()
	conf.IncludeLinkUrls = true

	text, err := Html2Text("mm <a href=\"https://test/\">Test</a>", &conf)
	checkerr(err)

	if text != `mm Test (https://test/) ` {
		t.Fail()
	}
}
