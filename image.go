package docx

import (
	"encoding/xml"
	"fmt"
)

// ref: http://officeopenxml.com/drwPic.php

type Image struct {
	XMLName xml.Name `xml:"w:p"`
	Data    []interface{}

	file *File
}

type Picture struct {
	XMLName xml.Name `xml:"pic"`

	Name string
	Data []byte
}

func (para *Paragraph) AddImage(name string) error {
	if name == "" {
		return fmt.Errorf("'name' MUST not be null or empty")
	}

	var pic *Picture = nil
	for _, entry := range para.file.Document.Pictures {
		if name == entry.Name {
			pic = entry
			break
		}
	}
	if pic == nil {
		return fmt.Errorf("image named %s not found in para.file.Document.Pictures", name)
	}

	// TODO: add to the paragraph XML...

	return nil
}
