package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type result struct {
	Num int `json:"num"`
}

func main() {
	http.HandleFunc("/alpha", GetNum)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func getNum(x string) (int, error) {
	x = strings.ToLower(x)
	numRune := []rune(x)
	for _, numChar := range numRune {
		if numChar == ' ' {
			return 0, nil
		}
		return int(numChar) - 96, nil
	}
	return 0, fmt.Errorf("something")
}

func GetNum(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "This sever allow POST only")
		return
	}

	alphabet := r.FormValue("alphabet")
	log.Println(alphabet)
	if len(alphabet) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Do not post more than a char")
		return
	}
	num, err := getNum(alphabet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "It occurs some trouble: %s", err.Error())
		return
	}

	time.Sleep(2 * time.Second)

	result := result{Num: num}
	res, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "It occurs some trouble: %s", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if alphabet == "Z" {
		w.WriteHeader(http.StatusNotAcceptable)

	}
	w.Write(res)
}
