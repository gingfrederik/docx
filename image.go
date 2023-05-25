package docx

import (
	"encoding/xml"
	"fmt"
	"strings"

	u "github.com/bcicen/go-units"
)

// REFERENCES
// 1. http://officeopenxml.com/drwPic.php
// 2. https://startbigthinksmall.wordpress.com/2010/01/04/points-inches-and-emus-measuring-units-in-office-open-xml/
// 3. https://developers.google.com/slides/api/reference/rest/v1/Unit

// compressionState for embeded images
type compressionState = string

var (
	CompressionStateEmail   = "email"   // email compression
	CompressionStateHQprint = "hqprint" // - high quality printing compression
	CompressionStateNone    = "none"    // none
	CompressionStatePrint   = "print"   // - printing compression
	CompressionStateScreen  = "screen"  // screen viewing compression
)

var unitEMU = u.NewUnit("emu", "emu")
var unitPixel = u.NewUnit("px", "px")

var A4WidthCentimeters = 21
var A4HeightCentimeters = 29.7

func init() {
	// An English Metric Unit (EMU) is defined as 1/360,000 of a centimeter and
	// thus there are 914,400 EMUs per inch, and 12,700 EMUs per point.
	u.NewRatioConversion(u.CentiMeter, unitEMU, 360_000)
	u.NewRatioConversion(u.CentiMeter, unitPixel, 37.7952755906)
}

type Drawing struct {
	XMLName xml.Name `xml:"w:drawing"`
	Inline  inlineDrawing
}

type Picture struct {
	XMLName xml.Name `xml:"pic:pic"`
	NS      string   `xml:"xmlns:pic,attr"`

	NonVisualProperties *nvPicPr
	Fill                *blipFill

	ID    string `xml:"-"` // unique ID for the image (auto-incrementing?)
	RelID string `xml:"-"` // to use as embed id in relationship
	Name  string `xml:"-"` // name of the file
	Data  []byte `xml:"-"` // data of the file to be written to the media/ directory
}

func (p *Picture) MediaName() string {
	return fmt.Sprintf("media/%s", p.Name)
}

func (p *Picture) ContentType() string {
	if strings.HasSuffix(p.Name, ".png") {
		return "image/png"
	}
	if strings.HasSuffix(p.Name, ".jpeg") {
		return "image/jpeg"
	}

	return "image"
}

func (p *Picture) Drawing() *Drawing {
	p.NS = XMLNS_PICTURE
	nvp := &nvPicPr{}
	if p.NonVisualProperties == nil {
		nvp.Properties = make([]cNvPicPr, 1)

		nameProp := cNvPicPr{
			ID:   "0",
			Name: p.Name,
		}

		nvp.Properties = append(nvp.Properties, nameProp)

		p.NonVisualProperties = nvp
	}

	if p.Fill == nil {
		p.Fill = defaultBlipFill(p)
	}

	// TODO: fix cx, cy extents based on the size of the image proper not just the dimensions of A4
	// TODO: also use the actual page's width and not hardcoded A4 dimensions..
	widthPadding := 2 // pad by 5 cm
	cx := u.MustConvertFloat(float64(A4WidthCentimeters-widthPadding), u.CentiMeter, unitEMU)
	cy := u.MustConvertFloat(float64(A4HeightCentimeters/2), u.CentiMeter, unitEMU)

	return &Drawing{
		Inline: inlineDrawing{
			TopDistance:    "0",
			BottomDistance: "0",
			LeftDistance:   "0",
			RightDistance:  "0",
			SimplePos: simplePos{
				X: "0",
				Y: "0",
			},
			Extent: extent{
				CX: strings.TrimSuffix(cx.String(), " emus"),
				CY: strings.TrimSuffix(cy.String(), " emus"),
			},
			Graphic: graphicElement{
				NS: "http://schemas.openxmlformats.org/drawingml/2006/main",
				GraphicData: graphicData{
					URI:     "http://schemas.openxmlformats.org/drawingml/2006/picture",
					Picture: p,
				},
			},
		},
	}
}

type blipFill struct {
	XMLName xml.Name `xml:"pic:blipFill"`

	Blip    blipProperties    `xml:"a:blip,omitempty"`
	Stretch stretchProperties `xml:"a:stretch,omitempty"`
}

type blipProperties struct {
	Embed            string           `xml:"r:embed,attr,omitempty"`
	CompressionState compressionState `xml:"cstate,attr,omitempty"`
}

type stretchProperties struct {
	FillRect interface{} `xml:"a:fillRect"`
}

func defaultBlipFill(p *Picture) *blipFill {
	return &blipFill{
		Blip: blipProperties{
			Embed: p.RelID,
			// CompressionState: CompressionStateScreen,
		},
		Stretch: stretchProperties{},
	}
}

type nvPicPr struct {
	XMLName    xml.Name `xml:"pic:nvPicPr"`
	Properties []cNvPicPr
}

type cNvPicPr struct {
	XMLName xml.Name `xml:"pic:cNvPr"`
	ID      string   `xml:"id,attr,omitempty"`
	Name    string   `xml:"name,attr,omitempty"`
}

type inlineDrawing struct {
	XMLName        xml.Name `xml:"wp:inline,omitempty"`
	TopDistance    string   `xml:"distT,attr,default:0"`
	BottomDistance string   `xml:"distB,attr,default:0"`
	LeftDistance   string   `xml:"distL,attr,default:0"`
	RightDistance  string   `xml:"distR,attr,default:0"`

	SimplePos simplePos
	Extent    extent
	Graphic   graphicElement
}

type simplePos struct {
	XMLName xml.Name `xml:"wp:simplePos"`
	X       string   `xml:"x,attr"`
	Y       string   `xml:"y,attr"`
}

type extent struct {
	XMLName xml.Name `xml:"wp:extent"`
	CX      string   `xml:"cx,attr"` // CX is measured EMUs
	CY      string   `xml:"cy,attr"` // CY is measured in EMUs
}

type graphicElement struct {
	XMLName     xml.Name `xml:"a:graphic"`
	NS          string   `xml:"xmlns:a,attr"`
	GraphicData graphicData
}

type graphicData struct {
	XMLName xml.Name `xml:"a:graphicData"`
	URI     string   `xml:"uri,attr"`
	Picture *Picture
}
