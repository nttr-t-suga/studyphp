package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {
	values := url.Values{}
	values.Add("alphabet", "Z")
	values2 := url.Values{}
	values2.Add("alphabet", "a")
	resp, _ := http.PostForm("http://localhost:8080/alpha", values)
	resp2, _ := http.PostForm("http://localhost:8080/alpha", values2) //<--- new
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	htmlData2, err := ioutil.ReadAll(resp2.Body) //<--- new
	if err != nil {                              //<--- new
		fmt.Println(err) //<--- new
		os.Exit(1)       //<--- new
	} //<--- new

	fmt.Println(string(htmlData))                   //<--- new
	fmt.Println(string(htmlData2))                  //<--- new
	split := strings.Split(string(htmlData), ":")   //<--- new
	split2 := strings.Split(split[1], "}")          //<--- new
	split3 := strings.Split(string(htmlData2), ":") //<--- new
	split4 := strings.Split(split3[1], "}")         //<--- new
	fmt.Println(split2[0])                          //<--- new
	fmt.Println(split4[0])                          //<--- new
	num, _ := strconv.Atoi(split2[0])               //<--- new
	num2, _ := strconv.Atoi(split4[0])              //<--- new

	if resp.StatusCode == 200 { //<--- new
		fmt.Println("response 200") //<--- new
	} else if resp.StatusCode == 406 { //<--- new
		num = -num                  //<--- new
		fmt.Println("response 406") //<--- new
	} //<--- new

	if resp2.StatusCode == 200 { //<--- new
		fmt.Println("response 200") //<--- new
	} else if resp2.StatusCode == 406 { //<--- new
		num2 = -num2                //<--- new
		fmt.Println("response 406") //<--- new
	} //<--- new

	var sum int
	sum = 0
	sum += num       //<--- new
	sum += num2      //<--- new
	fmt.Println(sum) //<--- new
}
