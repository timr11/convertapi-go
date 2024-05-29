package main

import (
	"fmt"
	"github.com/timr11/convertapi-go"
	"github.com/timr11/convertapi-go/config"
	"github.com/timr11/convertapi-go/param"
	"os"
)

func main() {
	config.Default.Secret = os.Getenv("CONVERTAPI_SECRET") // Get your secret at https://www.convertapi.com/a

	fmt.Println("Converting DOCX to PDF and JPG in parallel using same source file")
	fileParam := param.NewPath("file", "assets/test.docx", nil)
	pdfRes := convertapi.ConvDef("docx", "pdf", fileParam)
	jpgRes := convertapi.ConvDef("docx", "jpg", fileParam)

	// Downloading and saving files also in parallel
	c1 := save(pdfRes)
	c2 := save(jpgRes)
	<-c1
	<-c2
}

func save(res *convertapi.Result) (finish chan struct{}) {
	finish = make(chan struct{})
	go func() {
		defer close(finish)
		if files, errs := res.ToPath("/tmp"); errs == nil {
			fmt.Println("File saved to: ", files[0].Name())
		} else {
			fmt.Println(errs)
		}
	}()
	return
}
