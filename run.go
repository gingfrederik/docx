package docx

import "encoding/xml"

type Run struct {
	XMLName       xml.Name       `xml:"w:r"`
	RunProperties *RunProperties `xml:"w:rPr,omitempty"`
	InstrText     string         `xml:"w:instrText,omitempty"`
	Text          *Text
}

type Text struct {
	XMLName  xml.Name `xml:"w:t"`
	XMLSpace string   `xml:"xml:space,attr,omitempty"`
	Text     string   `xml:",chardata"`
}

// Color set run color
func (r *Run) Color(color string) *Run {
	r.RunProperties.Color = &Color{
		Val: color,
	}

	return r
}

// Size set run size
func (r *Run) Size(size int) *Run {
	r.RunProperties.Size = &Size{
		Val: size * 2,
	}
	return r
}

type Hyperlink struct {
	XMLName xml.Name `xml:"w:hyperlink"`
	ID      string   `xml:"r:id,attr"`
	Run     Run
}
