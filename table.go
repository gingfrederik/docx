package docx

import "encoding/xml"

type Table struct {
	XMLName    xml.Name `xml:"w:tbl"`
	Data       []interface{}
	grid       *TableGrid
	properties *TableProperties
	file       *File
}

type TableProperties struct {
	XMLName xml.Name `xml:"w:tblPr"`
	Data    []interface{}
	width   *TableWidth
}

type TableWidth struct {
	XMLName xml.Name `xml:"w:tblW"`
	Width   int      `xml:"w:w,attr"`
	Type    string   `xml:"w:type,attr"`
}

type TableGrid struct {
	XMLName xml.Name `xml:"w:tblGrid"`
	Data    []interface{}
}

type tableGridCol struct {
	XMLName xml.Name `xml:"w:gridCol"`
	Width   int      `xml:"w:w,attr"`
	Type    string   `xml:"w:type,attr"`
}

type Row struct {
	XMLName xml.Name `xml:"w:tr"`
	Data    []interface{}
	table   *Table
}

type RowProperties struct {
	XMLName xml.Name `xml:"w:trPr"`
	Data    []interface{}
	width   *TableWidth
}

type Cell struct {
	XMLName    xml.Name `xml:"w:tc"`
	Data       []interface{}
	Properties *CellProperties
	row        *Row
}

type CellProperties struct {
	XMLName xml.Name `xml:"w:tcPr"`
	Data    []interface{}
}

type cellWidth struct {
	XMLName xml.Name `xml:"w:tcW"`
	Width   int      `xml:"w:w,attr"`
	Type    string   `xml:"w:type,attr"`
}

// AddTable is adding table to document
func (f *File) AddTable() *Table {
	grid := &TableGrid{
		Data: make([]interface{}, 0),
	}
	tblW := &TableWidth{}
	props := &TableProperties{
		Data:  make([]interface{}, 0),
		width: tblW,
	}
	props.Data = append(props.Data, tblW)
	t := &Table{
		Data:       make([]interface{}, 0),
		grid:       grid,
		properties: props,
		file:       f,
	}
	t.Data = append(t.Data, props)
	t.Data = append(t.Data, grid)

	f.Document.Body.Content = append(f.Document.Body.Content, t)
	return t
}

// AddRow is adding row to table
func (t *Table) AddRow() *Row {
	r := &Row{
		table: t,
	}
	t.Data = append(t.Data, r)

	return r
}

// AddCell is adding cell to row, takes width in twips
func (r *Row) AddCell(width int) *Cell {
	tcProps := &CellProperties{}
	c := &Cell{
		row:        r,
		Properties: tcProps,
	}
	r.Data = append(r.Data, c)
	gridCol := &tableGridCol{
		Width: width,
		Type:  "dxa",
	}
	tcW := &cellWidth{
		Width: width,
		Type:  "dxa",
	}
	tcProps.Data = append(tcProps.Data, tcW)
	// Do not add more width than we have columns
	if len(r.Data) > len(r.table.grid.Data) {
		r.table.grid.Data = append(r.table.grid.Data, gridCol)
		r.table.properties.width.Width += width
	}

	return c
}

// AddText is adding text to cell
func (c *Cell) AddText(text string) *Run {
	p := &Paragraph{
		Data: make([]interface{}, 0),
	}
	txt := &Text{
		Text: text,
	}
	run := &Run{
		Text:          txt,
		RunProperties: &RunProperties{},
	}

	p.Data = append(p.Data, run)
	c.Data = append(c.Data, p)

	return run
}
