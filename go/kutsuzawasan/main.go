package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	// 対象の文字列
	str := "KUT"
	split := strings.Split(str, "")
	values := url.Values{}
	for _, s := range split {
		fmt.Println(s)
		// values := url.Values{}
		values.Add("alphabet", s)
		go resp, _ := http.PostForm("http://localhost:8080/alpha", values)
		htmlData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(htmlData)) //<--- new
	}

}
