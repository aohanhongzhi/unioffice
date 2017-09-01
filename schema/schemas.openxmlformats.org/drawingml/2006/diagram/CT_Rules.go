// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_Rules struct {
	Rule []*CT_NumericRule
}

func NewCT_Rules() *CT_Rules {
	ret := &CT_Rules{}
	return ret
}

func (m *CT_Rules) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Rule != nil {
		serule := xml.StartElement{Name: xml.Name{Local: "rule"}}
		e.EncodeElement(m.Rule, serule)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Rules) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Rules:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "rule":
				tmp := NewCT_NumericRule()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rule = append(m.Rule, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Rules
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Rules and its children
func (m *CT_Rules) Validate() error {
	return m.ValidateWithPath("CT_Rules")
}

// ValidateWithPath validates the CT_Rules and its children, prefixing error messages with path
func (m *CT_Rules) ValidateWithPath(path string) error {
	for i, v := range m.Rule {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rule[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
