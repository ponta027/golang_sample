package md2pdf

import (
	"log"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func htmlToPdf2(html string, pdffile string) error {
	log.Printf("htmlToPdf")
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))
	err = pdfg.Create()
	if err != nil {
		return err
	}
	log.Printf("Output Pdf:%s", pdffile)
	err = pdfg.WriteFile(pdffile)
	return err
}
