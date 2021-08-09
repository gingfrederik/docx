package main

import (
	"github.com/srdolor/docx"
)

func main() {
	f := docx.NewFile()
	// add new table
	t := f.AddTable()

	for i := 0; i < 3; i++ {
		row := t.AddRow()
		for i := 0; i < 3; i++ {
			c := row.AddCell()
			c.AddText("Hello").Size(20).Color("fbfcda")
		}
	}
	row := t.AddRow()
	c := row.AddCell()
	c.AddText("One column")
	row = t.AddRow()
	c = row.AddCell()
	c.AddText("Two columns")
	c = row.AddCell()
	c.AddText("Two columns")
	// add new paragraph
	para := f.AddParagraph()
	// add text
	para.AddText("test")

	para.AddText("test font size").Size(22)
	para.AddText("test color").Color("808080")
	para.AddText("test font size and color").Size(22).Color("121212")

	nextPara := f.AddParagraph()
	nextPara.AddLink("google", `http://google.com`)

	f.Save("./test.docx")
}
