package docx

import "encoding/xml"

const (
	XMLNS_W = `http://schemas.openxmlformats.org/wordprocessingml/2006/main`
	XMLNS_R = `http://schemas.openxmlformats.org/officeDocument/2006/relationships`
	// Drawing namespace
	XMLNS_DRAWING                 = `http://schemas.openxmlformats.org/drawingml/2006/main`
	XMLNS_DRAWING_WORD_PROCESSING = `http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing`
	XMLNS_PICTURE                 = `http://schemas.openxmlformats.org/drawingml/2006/picture`
	XMLNS_BLIP                    = `http://schemas.openxmlformats.org/drawingml/2006/main`
)

type Document struct {
	XMLName  xml.Name `xml:"w:document"`
	XMLW     string   `xml:"xmlns:w,attr"`
	XMLR     string   `xml:"xmlns:r,attr"`
	XMLWP    string   `xml:"xmlns:wp,attr"`
	Body     *Body
	Pictures []*Picture `xml:"-"`
}

type Body struct {
	XMLName   xml.Name `xml:"w:body"`
	Paragraph []*Paragraph
}
