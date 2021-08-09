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
}

type TableGrid struct {
	XMLName xml.Name `xml:"w:tblGrid"`
	Data    []interface{}
}

type tableGridCol struct {
	XMLName xml.Name `xml:"w:gridCol"`
	Width   int      `xml:"w:w,attr"`
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
	XMLName xml.Name `xml:"w:tc"`
	Data    []interface{}
	row     *Row
}

type CellProperties struct {
	XMLName xml.Name `xml:"w:rPr"`
	Data    []interface{}
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

// AddCell is adding cell to row
func (r *Row) AddCell() *Cell {
	c := &Cell{
		row: r,
	}
	r.Data = append(r.Data, c)
	cell := &tableGridCol{
		Width: 4098,
	}
	r.table.grid.Data = append(r.table.grid.Data, cell)
	r.table.properties.width.Width += 4098

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
