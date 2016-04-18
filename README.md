# Html2Text

# htmlcheck
converts html to text

Install
```
go get github.com/BlackEspresso/web2text
```

Example 1
``` Go
package main

import (
	"fmt"
	"github.com/BlackEspresso/web2text"
)

func main() {
	htmlText := `<div>Link:</div> <a href="https://www.google.com">google</a>`

	conf := web2text.NewSettings()
	text, err := web2text.Html2Text(htmlText, conf)
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}

```

prints

```
Link:  google (https://www.google.com)
```

Example 2
``` Go
package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"github.com/BlackEspresso/web2text"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	checkerr(err)
	
	content, err := ioutil.ReadAll(resp.Body)
	checkerr(err)
	
	conf := web2text.NewSettings()
	conf.IncludeLinkUrls = false
	text, err := web2text.Html2Text(string(content), conf)
	checkerr(err)
	
	fmt.Println(text)
}

func checkerr(err error){
	if err != nil{
		panic(err)
	}
}

```
prints

```

```
