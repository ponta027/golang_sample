package webserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type HelloJson struct {
	Name string `json:"name"`
}
type Ping struct {
	Status int
	Result string
}

func Start(port int) {
	log.Printf("port:%d\n", port)
	http.HandleFunc("/", handler)
	http.HandleFunc("/json", handlerJson)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func parseJson(query *url.Values) HelloJson {
	hello := HelloJson{query.Get("name")}
	return hello
}

func parseBody(r *http.Request) HelloJson {
	body := r.Body
	buf := new(bytes.Buffer)
	io.Copy(buf, body)
	log.Printf("parseBody:%s", string(buf.Bytes()))
	var hello HelloJson
	json.Unmarshal(buf.Bytes(), &hello)
	return hello
}

func handleJsonBody(w *http.ResponseWriter, query *url.Values) {
	var ping Ping
	if query != nil {
		log.Printf((*query).Get("name"))
		ping = Ping{http.StatusOK, query.Get("name")}
	}
	res, err := json.Marshal(ping)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Write(res)
}
func handleJsonBody2(w *http.ResponseWriter, hello *HelloJson) {
	ping := Ping{http.StatusOK, hello.Name}
	res, err := json.Marshal(ping)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Write(res)
}

/** */
func handlerJson(w http.ResponseWriter, r *http.Request) {
	log.Printf("Method:%s", r.Method)
	switch r.Method {
	case http.MethodGet:
		log.Printf("path:%s", r.URL.Path)
		query := r.URL.Query()
		hello := parseJson(&query)
		//handleJsonBody(&w, &query)
		handleJsonBody2(&w, &hello)
	case http.MethodPost:
		hello := parseBody(r)
		log.Printf("PostData:%s", hello)
		handleJsonBody2(&w, &hello)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
