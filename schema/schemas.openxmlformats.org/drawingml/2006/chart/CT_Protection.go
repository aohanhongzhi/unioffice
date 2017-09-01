// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"
)

type CT_Protection struct {
	ChartObject   *CT_Boolean
	Data          *CT_Boolean
	Formatting    *CT_Boolean
	Selection     *CT_Boolean
	UserInterface *CT_Boolean
}

func NewCT_Protection() *CT_Protection {
	ret := &CT_Protection{}
	return ret
}

func (m *CT_Protection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.ChartObject != nil {
		sechartObject := xml.StartElement{Name: xml.Name{Local: "chartObject"}}
		e.EncodeElement(m.ChartObject, sechartObject)
	}
	if m.Data != nil {
		sedata := xml.StartElement{Name: xml.Name{Local: "data"}}
		e.EncodeElement(m.Data, sedata)
	}
	if m.Formatting != nil {
		seformatting := xml.StartElement{Name: xml.Name{Local: "formatting"}}
		e.EncodeElement(m.Formatting, seformatting)
	}
	if m.Selection != nil {
		seselection := xml.StartElement{Name: xml.Name{Local: "selection"}}
		e.EncodeElement(m.Selection, seselection)
	}
	if m.UserInterface != nil {
		seuserInterface := xml.StartElement{Name: xml.Name{Local: "userInterface"}}
		e.EncodeElement(m.UserInterface, seuserInterface)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Protection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Protection:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "chartObject":
				m.ChartObject = NewCT_Boolean()
				if err := d.DecodeElement(m.ChartObject, &el); err != nil {
					return err
				}
			case "data":
				m.Data = NewCT_Boolean()
				if err := d.DecodeElement(m.Data, &el); err != nil {
					return err
				}
			case "formatting":
				m.Formatting = NewCT_Boolean()
				if err := d.DecodeElement(m.Formatting, &el); err != nil {
					return err
				}
			case "selection":
				m.Selection = NewCT_Boolean()
				if err := d.DecodeElement(m.Selection, &el); err != nil {
					return err
				}
			case "userInterface":
				m.UserInterface = NewCT_Boolean()
				if err := d.DecodeElement(m.UserInterface, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Protection
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Protection and its children
func (m *CT_Protection) Validate() error {
	return m.ValidateWithPath("CT_Protection")
}

// ValidateWithPath validates the CT_Protection and its children, prefixing error messages with path
func (m *CT_Protection) ValidateWithPath(path string) error {
	if m.ChartObject != nil {
		if err := m.ChartObject.ValidateWithPath(path + "/ChartObject"); err != nil {
			return err
		}
	}
	if m.Data != nil {
		if err := m.Data.ValidateWithPath(path + "/Data"); err != nil {
			return err
		}
	}
	if m.Formatting != nil {
		if err := m.Formatting.ValidateWithPath(path + "/Formatting"); err != nil {
			return err
		}
	}
	if m.Selection != nil {
		if err := m.Selection.ValidateWithPath(path + "/Selection"); err != nil {
			return err
		}
	}
	if m.UserInterface != nil {
		if err := m.UserInterface.ValidateWithPath(path + "/UserInterface"); err != nil {
			return err
		}
	}
	return nil
}
