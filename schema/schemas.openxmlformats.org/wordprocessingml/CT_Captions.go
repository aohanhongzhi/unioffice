// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_Captions struct {
	// Single Caption Type Definition
	Caption []*CT_Caption
	// Automatic Captioning Settings
	AutoCaptions *CT_AutoCaptions
}

func NewCT_Captions() *CT_Captions {
	ret := &CT_Captions{}
	return ret
}

func (m *CT_Captions) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	secaption := xml.StartElement{Name: xml.Name{Local: "w:caption"}}
	e.EncodeElement(m.Caption, secaption)
	if m.AutoCaptions != nil {
		seautoCaptions := xml.StartElement{Name: xml.Name{Local: "w:autoCaptions"}}
		e.EncodeElement(m.AutoCaptions, seautoCaptions)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Captions) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Captions:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "caption":
				tmp := NewCT_Caption()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Caption = append(m.Caption, tmp)
			case "autoCaptions":
				m.AutoCaptions = NewCT_AutoCaptions()
				if err := d.DecodeElement(m.AutoCaptions, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Captions
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Captions and its children
func (m *CT_Captions) Validate() error {
	return m.ValidateWithPath("CT_Captions")
}

// ValidateWithPath validates the CT_Captions and its children, prefixing error messages with path
func (m *CT_Captions) ValidateWithPath(path string) error {
	for i, v := range m.Caption {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Caption[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.AutoCaptions != nil {
		if err := m.AutoCaptions.ValidateWithPath(path + "/AutoCaptions"); err != nil {
			return err
		}
	}
	return nil
}
