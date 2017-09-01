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
)

type CT_ControlList struct {
	// Embedded Control
	Control []*CT_Control
}

func NewCT_ControlList() *CT_ControlList {
	ret := &CT_ControlList{}
	return ret
}

func (m *CT_ControlList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Control != nil {
		secontrol := xml.StartElement{Name: xml.Name{Local: "p:control"}}
		e.EncodeElement(m.Control, secontrol)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ControlList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ControlList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "control":
				tmp := NewCT_Control()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Control = append(m.Control, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ControlList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ControlList and its children
func (m *CT_ControlList) Validate() error {
	return m.ValidateWithPath("CT_ControlList")
}

// ValidateWithPath validates the CT_ControlList and its children, prefixing error messages with path
func (m *CT_ControlList) ValidateWithPath(path string) error {
	for i, v := range m.Control {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Control[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
