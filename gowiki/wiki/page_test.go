package page

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
 "github.com/google/go-cmp/cmp"
)

func cmpFile(a string, b string) error {
	bodya, erra := ioutil.ReadFile(a)
	if erra != nil {
		return erra
	}
	bodyb, errb := ioutil.ReadFile(b)
	if errb != nil {
		return errb
	}
	if !reflect.DeepEqual(bodya, bodyb) {
		return fmt.Errorf("not match")
	}
	return nil
}

func TestSave(t *testing.T) {

	p := &Page{Title: "test", Body: []byte("ABC")}
	err := p.Save()
	if err != nil {
		t.Fatal(err)
	}
	err = cmpFile("log/test.txt", "../testdata/page_save.log")
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadPage(t *testing.T) {

	expect := &Page{Title: "test", Body: []byte("ABC")}
	page, err := LoadPage("test")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(expect, page); diff != "" {
		t.Errorf("%s", diff)
	}

}
