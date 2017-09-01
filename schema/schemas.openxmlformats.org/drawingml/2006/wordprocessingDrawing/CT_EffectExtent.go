// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingDrawing

import (
	"encoding/xml"
	"fmt"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_EffectExtent struct {
	LAttr drawingml.ST_Coordinate
	TAttr drawingml.ST_Coordinate
	RAttr drawingml.ST_Coordinate
	BAttr drawingml.ST_Coordinate
}

func NewCT_EffectExtent() *CT_EffectExtent {
	ret := &CT_EffectExtent{}
	return ret
}

func (m *CT_EffectExtent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "l"},
		Value: fmt.Sprintf("%v", m.LAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "t"},
		Value: fmt.Sprintf("%v", m.TAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
		Value: fmt.Sprintf("%v", m.RAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "b"},
		Value: fmt.Sprintf("%v", m.BAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_EffectExtent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "l" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.LAttr = parsed
		}
		if attr.Name.Local == "t" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.TAttr = parsed
		}
		if attr.Name.Local == "r" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.RAttr = parsed
		}
		if attr.Name.Local == "b" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.BAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_EffectExtent: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_EffectExtent and its children
func (m *CT_EffectExtent) Validate() error {
	return m.ValidateWithPath("CT_EffectExtent")
}

// ValidateWithPath validates the CT_EffectExtent and its children, prefixing error messages with path
func (m *CT_EffectExtent) ValidateWithPath(path string) error {
	if err := m.LAttr.ValidateWithPath(path + "/LAttr"); err != nil {
		return err
	}
	if err := m.TAttr.ValidateWithPath(path + "/TAttr"); err != nil {
		return err
	}
	if err := m.RAttr.ValidateWithPath(path + "/RAttr"); err != nil {
		return err
	}
	if err := m.BAttr.ValidateWithPath(path + "/BAttr"); err != nil {
		return err
	}
	return nil
}
