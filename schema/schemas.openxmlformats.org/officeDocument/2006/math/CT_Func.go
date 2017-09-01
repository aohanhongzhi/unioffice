// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"log"
)

type CT_Func struct {
	FuncPr *CT_FuncPr
	FName  *CT_OMathArg
	E      *CT_OMathArg
}

func NewCT_Func() *CT_Func {
	ret := &CT_Func{}
	ret.FName = NewCT_OMathArg()
	ret.E = NewCT_OMathArg()
	return ret
}

func (m *CT_Func) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.FuncPr != nil {
		sefuncPr := xml.StartElement{Name: xml.Name{Local: "m:funcPr"}}
		e.EncodeElement(m.FuncPr, sefuncPr)
	}
	sefName := xml.StartElement{Name: xml.Name{Local: "m:fName"}}
	e.EncodeElement(m.FName, sefName)
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Func) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.FName = NewCT_OMathArg()
	m.E = NewCT_OMathArg()
lCT_Func:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "funcPr":
				m.FuncPr = NewCT_FuncPr()
				if err := d.DecodeElement(m.FuncPr, &el); err != nil {
					return err
				}
			case "fName":
				if err := d.DecodeElement(m.FName, &el); err != nil {
					return err
				}
			case "e":
				if err := d.DecodeElement(m.E, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Func
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Func and its children
func (m *CT_Func) Validate() error {
	return m.ValidateWithPath("CT_Func")
}

// ValidateWithPath validates the CT_Func and its children, prefixing error messages with path
func (m *CT_Func) ValidateWithPath(path string) error {
	if m.FuncPr != nil {
		if err := m.FuncPr.ValidateWithPath(path + "/FuncPr"); err != nil {
			return err
		}
	}
	if err := m.FName.ValidateWithPath(path + "/FName"); err != nil {
		return err
	}
	if err := m.E.ValidateWithPath(path + "/E"); err != nil {
		return err
	}
	return nil
}
