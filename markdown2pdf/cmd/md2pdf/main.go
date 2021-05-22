package main

import (
	"flag"
	"log"
	"markdown2pdf/md2pdf"
)

func main() {
	var (
		mdfile   = flag.String("i", "test.md", "input markdown file")
		pdffile  = flag.String("o", "test.pdf", "output pdf file")
		password = flag.String("p", "", "password.if password is empty, not protect")
	)
	flag.Parse()

	log.Printf("START")
	err := md2pdf.ConvertPdf(*mdfile, *pdffile, *password)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("END")

}
