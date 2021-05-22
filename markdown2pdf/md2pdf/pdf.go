package md2pdf

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func htmlToPdfWithPassword(html string, pdffile string, password string) error {
	pdffile_org := pdffile + "org"
	err := htmlToPdf(html, pdffile_org)
	exec_command := exec.Command("pdftk", pdffile_org, "output", pdffile, "owner_pw", password+"owner", "user_pw", password)
	std_out, std_err := exec_command.Output()
	status := exec_command.ProcessState.ExitCode()
	log.Printf(string(status))
	if std_err != nil {
		log.Fatal(std_err.Error())
	} else {
		log.Printf(string(std_out))
	}
	os.Remove(pdffile_org)
	return err
}

func htmlToPdf(html string, pdffile string) error {
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

func passwordSet() {
}
