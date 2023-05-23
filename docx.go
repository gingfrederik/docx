package docx

import "encoding/xml"

const (
	XMLNS_W = `http://schemas.openxmlformats.org/wordprocessingml/2006/main`
	XMLNS_R = `http://schemas.openxmlformats.org/officeDocument/2006/relationships`
	// Drawing namespace
	XMLNS_D = `xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main"`
)

type Document struct {
	XMLName  xml.Name `xml:"w:document"`
	XMLW     string   `xml:"xmlns:w,attr"`
	XMLR     string   `xml:"xmlns:r,attr"`
	Body     *Body
	Pictures []*Picture
}

type Body struct {
	XMLName   xml.Name `xml:"w:body"`
	Paragraph []*Paragraph
}
