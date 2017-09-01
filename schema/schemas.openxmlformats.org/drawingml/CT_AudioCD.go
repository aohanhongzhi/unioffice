// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type CT_AudioCD struct {
	St     *CT_AudioCDTime
	End    *CT_AudioCDTime
	ExtLst *CT_OfficeArtExtensionList
}

func NewCT_AudioCD() *CT_AudioCD {
	ret := &CT_AudioCD{}
	ret.St = NewCT_AudioCDTime()
	ret.End = NewCT_AudioCDTime()
	return ret
}

func (m *CT_AudioCD) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	sest := xml.StartElement{Name: xml.Name{Local: "a:st"}}
	e.EncodeElement(m.St, sest)
	seend := xml.StartElement{Name: xml.Name{Local: "a:end"}}
	e.EncodeElement(m.End, seend)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AudioCD) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.St = NewCT_AudioCDTime()
	m.End = NewCT_AudioCDTime()
lCT_AudioCD:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "st":
				if err := d.DecodeElement(m.St, &el); err != nil {
					return err
				}
			case "end":
				if err := d.DecodeElement(m.End, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AudioCD
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AudioCD and its children
func (m *CT_AudioCD) Validate() error {
	return m.ValidateWithPath("CT_AudioCD")
}

// ValidateWithPath validates the CT_AudioCD and its children, prefixing error messages with path
func (m *CT_AudioCD) ValidateWithPath(path string) error {
	if err := m.St.ValidateWithPath(path + "/St"); err != nil {
		return err
	}
	if err := m.End.ValidateWithPath(path + "/End"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
