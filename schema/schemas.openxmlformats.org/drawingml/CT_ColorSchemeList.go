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

type CT_ColorSchemeList struct {
	ExtraClrScheme []*CT_ColorSchemeAndMapping
}

func NewCT_ColorSchemeList() *CT_ColorSchemeList {
	ret := &CT_ColorSchemeList{}
	return ret
}

func (m *CT_ColorSchemeList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.ExtraClrScheme != nil {
		seextraClrScheme := xml.StartElement{Name: xml.Name{Local: "a:extraClrScheme"}}
		e.EncodeElement(m.ExtraClrScheme, seextraClrScheme)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ColorSchemeList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ColorSchemeList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "extraClrScheme":
				tmp := NewCT_ColorSchemeAndMapping()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ExtraClrScheme = append(m.ExtraClrScheme, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ColorSchemeList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ColorSchemeList and its children
func (m *CT_ColorSchemeList) Validate() error {
	return m.ValidateWithPath("CT_ColorSchemeList")
}

// ValidateWithPath validates the CT_ColorSchemeList and its children, prefixing error messages with path
func (m *CT_ColorSchemeList) ValidateWithPath(path string) error {
	for i, v := range m.ExtraClrScheme {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ExtraClrScheme[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
