package main

import (
	"os"

	"github.com/gingfrederik/docx"
)

func main() {
	f := docx.NewFile()
	// add new paragraph
	para := f.AddParagraph()
	// add text
	para.AddText("test")

	para.AddText("test font size").Size(22)
	para.AddText("test color").Color("808080")
	para.AddText("test font size and color").Size(22).Color("121212")

	nextPara := f.AddParagraph()
	nextPara.AddLink("google", `http://google.com`)

	imageData, _ := os.ReadFile("./TEACHING_GOPHER.png")

	err := f.AddImage("TEACHING_GOPHER.png", imageData)
	if err != nil {
		panic(err)
	}

	nextPara = f.AddParagraph()
	err = nextPara.AddImage("TEACHING_GOPHER.png" /* Add options for the image */)
	if err != nil {
		panic(err)
	}

	f.Save("./test.docx")
}
