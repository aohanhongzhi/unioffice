// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
)

type CT_RelativeRect struct {
	LAttr *ST_Percentage
	TAttr *ST_Percentage
	RAttr *ST_Percentage
	BAttr *ST_Percentage
}

func NewCT_RelativeRect() *CT_RelativeRect {
	ret := &CT_RelativeRect{}
	return ret
}

func (m *CT_RelativeRect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.LAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "l"},
			Value: fmt.Sprintf("%v", *m.LAttr)})
	}
	if m.TAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "t"},
			Value: fmt.Sprintf("%v", *m.TAttr)})
	}
	if m.RAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
			Value: fmt.Sprintf("%v", *m.RAttr)})
	}
	if m.BAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "b"},
			Value: fmt.Sprintf("%v", *m.BAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RelativeRect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "l" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.LAttr = &parsed
		}
		if attr.Name.Local == "t" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.TAttr = &parsed
		}
		if attr.Name.Local == "r" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.RAttr = &parsed
		}
		if attr.Name.Local == "b" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.BAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_RelativeRect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_RelativeRect and its children
func (m *CT_RelativeRect) Validate() error {
	return m.ValidateWithPath("CT_RelativeRect")
}

// ValidateWithPath validates the CT_RelativeRect and its children, prefixing error messages with path
func (m *CT_RelativeRect) ValidateWithPath(path string) error {
	if m.LAttr != nil {
		if err := m.LAttr.ValidateWithPath(path + "/LAttr"); err != nil {
			return err
		}
	}
	if m.TAttr != nil {
		if err := m.TAttr.ValidateWithPath(path + "/TAttr"); err != nil {
			return err
		}
	}
	if m.RAttr != nil {
		if err := m.RAttr.ValidateWithPath(path + "/RAttr"); err != nil {
			return err
		}
	}
	if m.BAttr != nil {
		if err := m.BAttr.ValidateWithPath(path + "/BAttr"); err != nil {
			return err
		}
	}
	return nil
}
