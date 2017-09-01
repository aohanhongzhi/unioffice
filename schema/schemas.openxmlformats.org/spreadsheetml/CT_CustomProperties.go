// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_CustomProperties struct {
	// Custom Property
	CustomPr []*CT_CustomProperty
}

func NewCT_CustomProperties() *CT_CustomProperties {
	ret := &CT_CustomProperties{}
	return ret
}

func (m *CT_CustomProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	secustomPr := xml.StartElement{Name: xml.Name{Local: "x:customPr"}}
	e.EncodeElement(m.CustomPr, secustomPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustomProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CustomProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "customPr":
				tmp := NewCT_CustomProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CustomPr = append(m.CustomPr, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CustomProperties and its children
func (m *CT_CustomProperties) Validate() error {
	return m.ValidateWithPath("CT_CustomProperties")
}

// ValidateWithPath validates the CT_CustomProperties and its children, prefixing error messages with path
func (m *CT_CustomProperties) ValidateWithPath(path string) error {
	for i, v := range m.CustomPr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CustomPr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
