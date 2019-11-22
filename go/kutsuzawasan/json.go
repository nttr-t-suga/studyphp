package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
)

func main() {
	// type Json struct {
	// 	Num int `json:"num"`
	// }
	values := url.Values{}
	values.Add("alphabet", "Z")
	// resp, _ := http.PostForm("http://localhost:8080/alpha", values)
	resp, _ := http.PostForm("http://alphabet.ms10.cf001.goo.ne.jp/alpha", values)
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	fmt.Println(reflect.TypeOf(resp))
	fmt.Println(reflect.TypeOf(htmlData))
	fmt.Println(htmlData)
	fmt.Println(string(htmlData))

	// var x = string(htmlData)
	// fmt.Println(x.name)
	// var buf bytes.Buffer
	// enc := json.NewEncoder(resp.Body)
	// if err := enc.Encode(buf); err != nil {
	// 	fmt.Println(1111)
	// 	log.Fatal(err)
	// }
	// fmt.Println(htmlData)
	// if resp.StatusCode == 200 { //<--- new
	// 	fmt.Println("response 200") //<--- new
	// } else if resp.StatusCode == 406 { //<--- new
	// 	// num = -num                  //<--- new
	// 	fmt.Println("response 406") //<--- new
	// } //<--- newtype Person struct {

	// fmt.Println(reflect.TypeOf(p.Name))
	// fmt.Println(err)

}
