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

type CT_ExternalSheetNames struct {
	// Sheet Name
	SheetName []*CT_ExternalSheetName
}

func NewCT_ExternalSheetNames() *CT_ExternalSheetNames {
	ret := &CT_ExternalSheetNames{}
	return ret
}

func (m *CT_ExternalSheetNames) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	sesheetName := xml.StartElement{Name: xml.Name{Local: "x:sheetName"}}
	e.EncodeElement(m.SheetName, sesheetName)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ExternalSheetNames) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ExternalSheetNames:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sheetName":
				tmp := NewCT_ExternalSheetName()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SheetName = append(m.SheetName, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ExternalSheetNames
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ExternalSheetNames and its children
func (m *CT_ExternalSheetNames) Validate() error {
	return m.ValidateWithPath("CT_ExternalSheetNames")
}

// ValidateWithPath validates the CT_ExternalSheetNames and its children, prefixing error messages with path
func (m *CT_ExternalSheetNames) ValidateWithPath(path string) error {
	for i, v := range m.SheetName {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SheetName[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
