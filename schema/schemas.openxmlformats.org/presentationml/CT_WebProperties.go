// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_WebProperties struct {
	// Show animation in HTML output
	ShowAnimationAttr *bool
	// Resize graphics in HTML output
	ResizeGraphicsAttr *bool
	// Allow PNG in HTML output
	AllowPngAttr *bool
	// Rely on VML for HTML output
	RelyOnVmlAttr *bool
	// Organize HTML output in folders
	OrganizeInFoldersAttr *bool
	// Use long file names in HTML output
	UseLongFilenamesAttr *bool
	// Image size for HTML output
	ImgSzAttr ST_WebScreenSize
	// Encoding for HTML output
	EncodingAttr *string
	// Slide Navigation Colors for HTML output
	ClrAttr ST_WebColorType
	ExtLst  *CT_ExtensionList
}

func NewCT_WebProperties() *CT_WebProperties {
	ret := &CT_WebProperties{}
	return ret
}

func (m *CT_WebProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ShowAnimationAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showAnimation"},
			Value: fmt.Sprintf("%v", *m.ShowAnimationAttr)})
	}
	if m.ResizeGraphicsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "resizeGraphics"},
			Value: fmt.Sprintf("%v", *m.ResizeGraphicsAttr)})
	}
	if m.AllowPngAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "allowPng"},
			Value: fmt.Sprintf("%v", *m.AllowPngAttr)})
	}
	if m.RelyOnVmlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "relyOnVml"},
			Value: fmt.Sprintf("%v", *m.RelyOnVmlAttr)})
	}
	if m.OrganizeInFoldersAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "organizeInFolders"},
			Value: fmt.Sprintf("%v", *m.OrganizeInFoldersAttr)})
	}
	if m.UseLongFilenamesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "useLongFilenames"},
			Value: fmt.Sprintf("%v", *m.UseLongFilenamesAttr)})
	}
	if m.ImgSzAttr != ST_WebScreenSizeUnset {
		attr, err := m.ImgSzAttr.MarshalXMLAttr(xml.Name{Local: "imgSz"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.EncodingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "encoding"},
			Value: fmt.Sprintf("%v", *m.EncodingAttr)})
	}
	if m.ClrAttr != ST_WebColorTypeUnset {
		attr, err := m.ClrAttr.MarshalXMLAttr(xml.Name{Local: "clr"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_WebProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "showAnimation" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowAnimationAttr = &parsed
		}
		if attr.Name.Local == "resizeGraphics" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ResizeGraphicsAttr = &parsed
		}
		if attr.Name.Local == "allowPng" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AllowPngAttr = &parsed
		}
		if attr.Name.Local == "relyOnVml" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RelyOnVmlAttr = &parsed
		}
		if attr.Name.Local == "organizeInFolders" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OrganizeInFoldersAttr = &parsed
		}
		if attr.Name.Local == "useLongFilenames" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UseLongFilenamesAttr = &parsed
		}
		if attr.Name.Local == "imgSz" {
			m.ImgSzAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "encoding" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.EncodingAttr = &parsed
		}
		if attr.Name.Local == "clr" {
			m.ClrAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_WebProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_WebProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_WebProperties and its children
func (m *CT_WebProperties) Validate() error {
	return m.ValidateWithPath("CT_WebProperties")
}

// ValidateWithPath validates the CT_WebProperties and its children, prefixing error messages with path
func (m *CT_WebProperties) ValidateWithPath(path string) error {
	if err := m.ImgSzAttr.ValidateWithPath(path + "/ImgSzAttr"); err != nil {
		return err
	}
	if err := m.ClrAttr.ValidateWithPath(path + "/ClrAttr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
