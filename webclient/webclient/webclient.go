package webclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Ping struct {
	Status int
	Result string
}

/**
* https://golang.org/pkg/net/http/
**/
func Get(url string, query string) error {
	log.Printf("Start Request:%s", url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	headers := resp.Header
	tmp := getHeader("Content-Type", &headers)
	if strings.Index(tmp, "application/json") > -1 {
		fmt.Printf("JSON")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	root := parseJson(&data)
	log.Printf("{Status:%d,Result=%s}", root.Status, root.Result)

	//	verboseHeader(&headers)
	return nil
}

func parseJson(data *[]byte) Ping {
	var root Ping
	if err := json.Unmarshal(*data, &root); err != nil {
		log.Fatal(err)
	}
	return root
}

func verboseHeader(headers *http.Header) {
	for i, v := range *headers {
		log.Printf("key:%s", i)
		for _, k := range v {
			log.Printf("value:%s", k)
		}
	}
}
func getHeader(name string, header *http.Header) string {
	return header.Get(name)
}
