package rest

import (
	"testing"
	"encoding/json"
)

/*** */
func TestConvertTemplateNotFound(t *testing.T) {
	var root Root
	path := "../testdata/test_notfound.tmpl"
	err := convertTemplate(&root, path)
	if err == nil {
		t.Fatalf("failed test ")
	}

}
/*** */
func TestConvertTemplate(t *testing.T) {
    result := "{}"
    b := []byte(result)
	var root Root
	if err := json.Unmarshal(b, &root); err != nil { } 
	path := "../testdata/test.tmpl"
	err := convertTemplate(&root, path)
	if err == nil {
		t.Fatalf("failed test ")
	}

}


func TestRequest(t *testing.T) {
	url := "http://google.co.jp"
	result, err := request(url)
	if err != "" {
		t.Fatalf("failed test %#v", err)
	}
	if len(result) == 0 {
		t.Fatalf("failed test %#v", err)
	}
}

/**
 */
func TestParseConfig(t *testing.T) {
	result, err := parseConfig("../testdata/test.json")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result.TemplatePath != "./internal/redmine.tmpl" {
		t.Fatalf("failed test %#v", err)
	}
	if result.BaseUrl != "https://www.redmine.org/projects/redmine/issues.json?limit=2" {
		t.Fatalf("failed test %#v", err)
	}
}

/**
 */
func TestParseConfigNotFound(t *testing.T) {
	result, err := parseConfig("../testdata/test_notfound.json")
	if err == nil {
		t.Fatalf("failed test %#v", result)
	}
}
