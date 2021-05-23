package webclient

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

/**
* https://golang.org/pkg/net/http/
**/
func Get(url string, query string) error {
	log.Printf("Start Request:%s", url)

	resp, err := request("GET", url, query)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	parseResponse(resp)
	return nil
}

func Post(urlpath string, query string) error {
	log.Printf("Start Post Request:%s", urlpath)
	resp, err := request("POST", urlpath, query)
	if err != nil {
		return err
	}

	if err != nil {
		log.Printf(err.Error())
		return err
	}
	defer resp.Body.Close()

	parseResponse(resp)
	return nil

}

func request(method string, urlPath string, query string) (*http.Response, error) {
	var req *http.Request = nil
	var err error
	client := &http.Client{}
	switch method {
	case "GET":
		req, err = http.NewRequest("GET", urlPath, nil)
	case "POST":
		req, err = http.NewRequest("POST", urlPath, strings.NewReader(query))
	default:
		return nil, errors.New("Not Allowed Method")
	}

	if err != nil {
		return nil, err
	}

	return client.Do(req)
}
