package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
    "ponta027.dip.jp/helloworld"
)

type Statuses struct{
    Name string `json:"name"`
}

type Trackers struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Issue struct {
	Tracker   Trackers `json:"tracker"`
    Status    Statuses  `json:"status"`
	Desc      string   `json:"description"`
	Subject   string   `json:"subject"`
	DoneRatio int      `json:"done_ratio"`
	Id        int      `json:"id"`
}
type Root struct {
	Issues []Issue `json:"issues"`
}

func request(req string) ([]byte, string) {
	url := "https://www.redmine.org/projects/redmine/issues.json?limit=2"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	//
	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, ""
}

func main() {
	var (
		configPath = flag.String("path", "", "specify config path")
	)

	flag.Parse()
	fmt.Println("START")
	fmt.Println(*configPath)

	result, err := request(*configPath)
	if err != "" {
	}
	var root Root
	if err := json.Unmarshal(result, &root); err != nil {
	}
	for _, p := range root.Issues {
		fmt.Printf("[%s]#%d[%s],DoneRatiod=%d,subject=%s", p.Tracker.Name, p.Id, p.Status.Name,p.DoneRatio, p.Subject)
		fmt.Println("")
	}
    helloworld.Hello()
}
