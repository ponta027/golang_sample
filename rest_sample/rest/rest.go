package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Config struct {
	TemplatePath string `json:"templatePath"`
	BaseUrl      string `json:"base_url"`
}

type Statuses struct {
	Name string `json:"name"`
}

type Trackers struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Issue struct {
	Tracker   Trackers `json:"tracker"`
	Status    Statuses `json:"status"`
	Desc      string   `json:"description"`
	Subject   string   `json:"subject"`
	DoneRatio int      `json:"done_ratio"`
	Id        int      `json:"id"`
}
type Root struct {
	Issues []Issue `json:"issues"`
}

func Request(configPath string) {
	config, er := parseConfig(configPath)
	if er != nil {
		fmt.Println(er)
        os.Exit(1)
	}

	result, err := request(config.BaseUrl)
	if err != "" {
	}
	var root Root
	if err := json.Unmarshal(result, &root); err != nil {
	}
	if err := convertTemplate(&root, config.TemplatePath) ; err != nil {
    }
}

/** */
func request(url string) ([]byte, string) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, ""
}

/** */
func existFile(tpath string) bool {
	_, err := os.Stat(tpath)
	return err == nil
}

/** */
func convertTemplate(root *Root, tpath string) error{
    if !existFile(tpath) {
        return fmt.Errorf("file not exist %s",tpath)
    }
	tmpl := template.Must(template.ParseFiles(tpath))
	if err := tmpl.Execute(os.Stdout, root); err != nil {
		log.Fatal(err)
        return err
	}
    return nil
}

/** */
func parseConfig(configPath string) (Config, error) {
	raw, err := ioutil.ReadFile(configPath)
	var root Config
	if err != nil {
		return root, err
	}
	json.Unmarshal(raw, &root)
	return root, nil

}
