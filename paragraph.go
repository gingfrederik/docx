package docx

import "encoding/xml"

const (
	JUSTIFY_START      = "start"
	JUSTIFY_END        = "end"
	JUSTIFY_BOTH       = "both"
	JUSTIFY_CENTER     = "center"
	JUSTIFY_DISTRIBUTE = "distribute"
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

func (prp *ParagraphProperties) Justification(justificaiton string) {
	prp.Data = append(prp.Data, &Justification{Val: justificaiton})
}
