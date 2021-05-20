package md2pdf

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/russross/blackfriday"
)

/**
 */
func ConvertPdf(markdown string, pdffile string) error {
	log.Printf("ConvertPdf %s->%s", markdown, pdffile)
	md, err := ioutil.ReadFile(markdown)
	if err != nil {
		return err
	}
	output := blackfriday.MarkdownCommon([]byte(md))
	header := `<html><meta charset="utf-8"<head><title></title></head>%s</body></html>`
	html := fmt.Sprintf(header, string(output))
	return htmlToPdf2(html, pdffile)
}
