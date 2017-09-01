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

type CT_BaseStylesOverride struct {
	ClrScheme  *CT_ColorScheme
	FontScheme *CT_FontScheme
	FmtScheme  *CT_StyleMatrix
}

func NewCT_BaseStylesOverride() *CT_BaseStylesOverride {
	ret := &CT_BaseStylesOverride{}
	return ret
}

func (m *CT_BaseStylesOverride) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.ClrScheme != nil {
		seclrScheme := xml.StartElement{Name: xml.Name{Local: "a:clrScheme"}}
		e.EncodeElement(m.ClrScheme, seclrScheme)
	}
	if m.FontScheme != nil {
		sefontScheme := xml.StartElement{Name: xml.Name{Local: "a:fontScheme"}}
		e.EncodeElement(m.FontScheme, sefontScheme)
	}
	if m.FmtScheme != nil {
		sefmtScheme := xml.StartElement{Name: xml.Name{Local: "a:fmtScheme"}}
		e.EncodeElement(m.FmtScheme, sefmtScheme)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BaseStylesOverride) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_BaseStylesOverride:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "clrScheme":
				m.ClrScheme = NewCT_ColorScheme()
				if err := d.DecodeElement(m.ClrScheme, &el); err != nil {
					return err
				}
			case "fontScheme":
				m.FontScheme = NewCT_FontScheme()
				if err := d.DecodeElement(m.FontScheme, &el); err != nil {
					return err
				}
			case "fmtScheme":
				m.FmtScheme = NewCT_StyleMatrix()
				if err := d.DecodeElement(m.FmtScheme, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BaseStylesOverride
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BaseStylesOverride and its children
func (m *CT_BaseStylesOverride) Validate() error {
	return m.ValidateWithPath("CT_BaseStylesOverride")
}

// ValidateWithPath validates the CT_BaseStylesOverride and its children, prefixing error messages with path
func (m *CT_BaseStylesOverride) ValidateWithPath(path string) error {
	if m.ClrScheme != nil {
		if err := m.ClrScheme.ValidateWithPath(path + "/ClrScheme"); err != nil {
			return err
		}
	}
	if m.FontScheme != nil {
		if err := m.FontScheme.ValidateWithPath(path + "/FontScheme"); err != nil {
			return err
		}
	}
	if m.FmtScheme != nil {
		if err := m.FmtScheme.ValidateWithPath(path + "/FmtScheme"); err != nil {
			return err
		}
	}
	return nil
}
