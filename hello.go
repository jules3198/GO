package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Tp1 (w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
	currentTime := time.Now()
	var resp string = strconv.Itoa(currentTime.Hour())+"h"+ strconv.Itoa(currentTime.Minute())
	fmt.Fprintf(w, resp)
	}
}

func writeText (w http.ResponseWriter,data string) {
	oldText := readText()
	newText := oldText + "\n" + data
	f, err := os.Create("db.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    _, err2 := f.WriteString(newText)
    if err2 != nil {
        log.Fatal(err2)
    }
	resp := strings.Replace(data, "=>", ":", 1)
    fmt.Fprintf(w,resp)
}

func readText () string {
	text := ""
	f, err := os.Open("db.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
		text = text + "\n" +scanner.Text()
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	return text
}

func Tp2 (w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
	if err := req.ParseForm(); err != nil {
	fmt.Println("Something went bad")
	fmt.Fprintln(w, "Something went bad")
	return
	}
	var data string = ""
	for _, value := range req.PostForm {
	data = data + "=>" +value[0]
	}
	data = data + "\n"
	data = strings.Replace(data, "=>","",1)
	writeText(w,data)
	}
}

func Tp3 (w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		data := strings.Split(readText(),"\n")
		text := ""
		for val := range data {
			if data[val] != "" {
				dat := strings.Split(data[val], "=>")
				text = text + dat[1] + "\n"
			}
		}
	fmt.Fprintf(w, text)
	}
}

func main() {
	http.HandleFunc("/", Tp1)
	http.HandleFunc("/add", Tp2)
	http.HandleFunc("/entries", Tp3)
 	http.ListenAndServe(":4567", nil)

}