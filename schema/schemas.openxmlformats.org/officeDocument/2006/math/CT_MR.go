// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_MR struct {
	E []*CT_OMathArg
}

func NewCT_MR() *CT_MR {
	ret := &CT_MR{}
	return ret
}

func (m *CT_MR) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MR) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_MR:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "e":
				tmp := NewCT_OMathArg()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.E = append(m.E, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MR
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_MR and its children
func (m *CT_MR) Validate() error {
	return m.ValidateWithPath("CT_MR")
}

// ValidateWithPath validates the CT_MR and its children, prefixing error messages with path
func (m *CT_MR) ValidateWithPath(path string) error {
	for i, v := range m.E {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/E[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
