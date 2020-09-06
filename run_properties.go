package docx

import "encoding/xml"

const (
	HYPERLINK_STYLE = "a1"
)

type RunProperties struct {
	XMLName  xml.Name  `xml:"w:rPr"`
	Color    *Color    `xml:"w:color,omitempty"`
	Size     *Size     `xml:"w:sz,omitempty"`
	RunStyle *RunStyle `xml:"w:rStyle,omitempty"`
}

type RunStyle struct {
	XMLName xml.Name `xml:"w:rStyle"`
	Val     string   `xml:"w:val,attr"`
}

type Color struct {
	XMLName xml.Name `xml:"w:color"`
	Val     string   `xml:"w:val,attr"`
}

type Size struct {
	XMLName xml.Name `xml:"w:sz"`
	Val     int      `xml:"w:val,attr"`
}
