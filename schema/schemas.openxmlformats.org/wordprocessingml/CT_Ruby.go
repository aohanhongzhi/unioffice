// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type CT_Ruby struct {
	// Phonetic Guide Properties
	RubyPr *CT_RubyPr
	// Phonetic Guide Text
	Rt *CT_RubyContent
	// Phonetic Guide Base Text
	RubyBase *CT_RubyContent
}

func NewCT_Ruby() *CT_Ruby {
	ret := &CT_Ruby{}
	ret.RubyPr = NewCT_RubyPr()
	ret.Rt = NewCT_RubyContent()
	ret.RubyBase = NewCT_RubyContent()
	return ret
}

func (m *CT_Ruby) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	serubyPr := xml.StartElement{Name: xml.Name{Local: "w:rubyPr"}}
	e.EncodeElement(m.RubyPr, serubyPr)
	sert := xml.StartElement{Name: xml.Name{Local: "w:rt"}}
	e.EncodeElement(m.Rt, sert)
	serubyBase := xml.StartElement{Name: xml.Name{Local: "w:rubyBase"}}
	e.EncodeElement(m.RubyBase, serubyBase)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Ruby) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.RubyPr = NewCT_RubyPr()
	m.Rt = NewCT_RubyContent()
	m.RubyBase = NewCT_RubyContent()
lCT_Ruby:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "rubyPr":
				if err := d.DecodeElement(m.RubyPr, &el); err != nil {
					return err
				}
			case "rt":
				if err := d.DecodeElement(m.Rt, &el); err != nil {
					return err
				}
			case "rubyBase":
				if err := d.DecodeElement(m.RubyBase, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Ruby
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Ruby and its children
func (m *CT_Ruby) Validate() error {
	return m.ValidateWithPath("CT_Ruby")
}

// ValidateWithPath validates the CT_Ruby and its children, prefixing error messages with path
func (m *CT_Ruby) ValidateWithPath(path string) error {
	if err := m.RubyPr.ValidateWithPath(path + "/RubyPr"); err != nil {
		return err
	}
	if err := m.Rt.ValidateWithPath(path + "/Rt"); err != nil {
		return err
	}
	if err := m.RubyBase.ValidateWithPath(path + "/RubyBase"); err != nil {
		return err
	}
	return nil
}
