package main
import (
	"log"
	"github.com/signintech/gopdf"
)

func badge(textB, NameFile) {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{ PageSize: *gopdf.PageSizeA4 })
	pdf.AddPage()
	err := pdf.AddTTFFont("wts11", "../ttf/wts11.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, textB)
	pdf.WritePdf(NameFile)

}

func imageBadge(textB, imadeB, NameFile) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4 })
	pdf.AddPage()
	var err error
	err = pdf.AddTTFFont("loma", "../ttf/Loma.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.Image("../imgs/" + imadeB + ".jpg", 200, 50, nil) //print image
	err = pdf.SetFont("loma", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetXY(250, 200) //move current location
	pdf.Cell(nil, textB) //print text
	pdf.WritePdf(NameFile)
}
