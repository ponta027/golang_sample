package md2pdf

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/russross/blackfriday"
)

/**
 */
func ConvertPdf(markdown string, pdffile string, password string) error {
	log.Printf("ConvertPdf %s->%s", markdown, pdffile)

	md, err := ioutil.ReadFile(markdown)
	if err != nil {
		return err
	}
	output := blackfriday.MarkdownCommon([]byte(md))
	header := `<html><meta charset="utf-8"<head><title></title></head>%s</body></html>`
	html := fmt.Sprintf(header, string(output))
	if len(password) > 0 {
		return htmlToPdfWithPassword(html, pdffile, password)
	} else {
		return htmlToPdf(html, pdffile)
	}
}
