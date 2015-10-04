package web2text

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"testing"
)

func checkerr (err error){
	if err != nil {
		panic(err)
	}
}

func TestHtml2TextSkipTags(t *testing.T){
	conf := NewSettings()
	str, err := Html2Text("q<script>ddgft</script>mm",conf)
	checkerr(err)
	fmt.Println(str)
}

func TestHtml2TextLinks(t *testing.T){
	conf := NewSettings()
	conf.IncludeLinkUrls = true
	str, err := Html2Text("mm <a href=\"https://test/\">Test</a>",conf)
	checkerr(err)
	fmt.Println(str)
}

func TestHtml2TextGoogle(t *testing.T) {
	resp, err := http.Get("https://www.reddit.com/r/golang/")
	checkerr(err)
	c, err := ioutil.ReadAll(resp.Body)
	checkerr(err)
	conf := NewSettings()
	str, err := Html2Text(string(c),conf)
	checkerr(err)
	fmt.Println(str)
}
