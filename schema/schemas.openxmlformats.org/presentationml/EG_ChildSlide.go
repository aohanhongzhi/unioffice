// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type EG_ChildSlide struct {
	// Color Scheme Map Override
	ClrMapOvr *drawingml.CT_ColorMappingOverride
}

func NewEG_ChildSlide() *EG_ChildSlide {
	ret := &EG_ChildSlide{}
	return ret
}

func (m *EG_ChildSlide) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ClrMapOvr != nil {
		seclrMapOvr := xml.StartElement{Name: xml.Name{Local: "p:clrMapOvr"}}
		e.EncodeElement(m.ClrMapOvr, seclrMapOvr)
	}
	return nil
}

func (m *EG_ChildSlide) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ChildSlide:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "clrMapOvr":
				m.ClrMapOvr = drawingml.NewCT_ColorMappingOverride()
				if err := d.DecodeElement(m.ClrMapOvr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ChildSlide
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ChildSlide and its children
func (m *EG_ChildSlide) Validate() error {
	return m.ValidateWithPath("EG_ChildSlide")
}

// ValidateWithPath validates the EG_ChildSlide and its children, prefixing error messages with path
func (m *EG_ChildSlide) ValidateWithPath(path string) error {
	if m.ClrMapOvr != nil {
		if err := m.ClrMapOvr.ValidateWithPath(path + "/ClrMapOvr"); err != nil {
			return err
		}
	}
	return nil
}
