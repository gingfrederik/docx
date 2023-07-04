package docx

import (
	"encoding/xml"
	"fmt"
)

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

func (p *Paragraph) AddImage(name string) error {
	if name == "" {
		return fmt.Errorf("'name' MUST not be null or empty")
	}

	var pic *Picture = nil
	for _, entry := range p.file.Document.Pictures {
		if name == entry.Name {
			pic = entry
			break
		}
	}
	if pic == nil {
		return fmt.Errorf("image named %s not found in p.file.Document.Pictures", name)
	}

	pic.RelID = p.file.addImageRelation(pic)

	p.Data = append(p.Data, &Run{Drawing: pic.Drawing()})

	return nil
}
