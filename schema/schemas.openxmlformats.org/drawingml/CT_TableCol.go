// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_TableCol struct {
	WAttr  ST_Coordinate
	ExtLst *CT_OfficeArtExtensionList
}

func NewCT_TableCol() *CT_TableCol {
	ret := &CT_TableCol{}
	return ret
}

func (m *CT_TableCol) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w"},
		Value: fmt.Sprintf("%v", m.WAttr)})
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TableCol) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "w" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.WAttr = parsed
		}
	}
lCT_TableCol:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
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
			break lCT_TableCol
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TableCol and its children
func (m *CT_TableCol) Validate() error {
	return m.ValidateWithPath("CT_TableCol")
}

// ValidateWithPath validates the CT_TableCol and its children, prefixing error messages with path
func (m *CT_TableCol) ValidateWithPath(path string) error {
	if err := m.WAttr.ValidateWithPath(path + "/WAttr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
