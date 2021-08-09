package docx

import "encoding/xml"

const (
	JUSTIFY_START      = "start"
	JUSTIFY_END        = "end"
	JUSTIFY_BOTH       = "both"
	JUSTIFY_CENTER     = "center"
	JUSTIFY_DISTRIBUTE = "distribute"
	CM                 = 567
	INCH               = 1440
)

type Paragraph struct {
	XMLName    xml.Name `xml:"w:p"`
	Data       []interface{}
	Properties *ParagraphProperties
	file       *File
}

type ParagraphProperties struct {
	XMLName xml.Name `xml:"w:pPr"`
	Data    []interface{}
}

type Justification struct {
	XMLName xml.Name `xml:"w:jc"`
	Val     string   `xml:"w:val,attr"`
}

type Indentation struct {
	XMLName xml.Name `xml:"w:ind"`
	Left    int      `xml:"w:left,attr"`
	Right   int      `xml:"w:right,attr"`
}

type LineBreak struct {
	XMLName xml.Name `xml:"w:br"`
	Type    string   `xml:"w:type,attr"`
}

// AddParagraph add new paragraph
func (f *File) AddParagraph() *Paragraph {
	props := &ParagraphProperties{}
	p := &Paragraph{
		Data:       make([]interface{}, 0),
		Properties: props,
		file:       f,
	}
	p.Data = append(p.Data, props)

	f.Document.Body.Content = append(f.Document.Body.Content, p)
	return p
}

// AddText add text to paragraph
func (p *Paragraph) AddText(text string) *Run {
	t := &Text{
		Text: text,
	}

	run := &Run{
		Text:          t,
		RunProperties: &RunProperties{},
	}

	p.Data = append(p.Data, run)

	return run
}

// AddLink add hyperlink to paragraph
func (p *Paragraph) AddLink(text string, link string) *Hyperlink {
	rId := p.file.addLinkRelation(link)
	hyperlink := &Hyperlink{
		ID: rId,
		Run: Run{
			RunProperties: &RunProperties{
				RunStyle: &RunStyle{
					Val: HYPERLINK_STYLE,
				},
			},
			InstrText: text,
		},
	}

	p.Data = append(p.Data, hyperlink)

	return hyperlink
}

// AddNewLine adds break with type textWrapping, newline after previous
func (p *Paragraph) AddNewLine() {
	p.Data = append(p.Data, &LineBreak{Type: "textWrapping"})
}

// AddNewPage adds break with type page, new page after previous
func (p *Paragraph) AddNewPage() {
	p.Data = append(p.Data, &LineBreak{Type: "page"})
}

// Justification takes justificaiton type
// Possible values are in constants:
// JUSTIFY_START align to the left
// JUSTIFY_END align to the right
// JUSTIFY_BOTH align to both sides with regular spaces
// JUSTIFY_DISTRIBUTE align to both sides with wider spaces
// JUSTIFY_CENTER align to the center of line
func (prp *ParagraphProperties) Justification(justificaiton string) {
	prp.Data = append(prp.Data, &Justification{Val: justificaiton})
}

// Indentation takes left and right indentation for paragraph in twips
// There is constant CM which is equal to 567 twips per centimeter
// There is constant INCH which is equal to 1440 twips per inch
func (prp *ParagraphProperties) Indentation(left, right int) {
	prp.Data = append(prp.Data, &Indentation{Left: left, Right: right})
}
