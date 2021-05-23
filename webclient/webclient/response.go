package webclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Ping struct {
	Status int
	Result string
}

func parseJson(data *[]byte) Ping {
	var root Ping
	log.Printf(string(*data))
	if err := json.Unmarshal(*data, &root); err != nil {
		log.Fatal(err)
	}
	return root
}
func parseResponse(resp *http.Response) error {
	headers := resp.Header
	verboseHeader(&headers)
	tmp := getHeader("Content-Type", &headers)
	if strings.Index(tmp, "application/json") < 0 {
		return errors.New("Response Content-Type Error")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	root := parseJson(&data)
	log.Printf("{Status:%d,Result=%s}", root.Status, root.Result)
	return nil
}

/**
 */
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
