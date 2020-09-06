package docx

import "encoding/xml"

type Paragraph struct {
	XMLName xml.Name `xml:"w:p"`
	Data    []interface{}

	file *File
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
