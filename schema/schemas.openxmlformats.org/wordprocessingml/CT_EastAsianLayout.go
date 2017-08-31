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
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_EastAsianLayout struct {
	// East Asian Typography Run ID
	IdAttr *int64
	// Two Lines in One
	CombineAttr *sharedTypes.ST_OnOff
	// Display Brackets Around Two Lines in One
	CombineBracketsAttr ST_CombineBrackets
	// Horizontal in Vertical (Rotate Text)
	VertAttr *sharedTypes.ST_OnOff
	// Compress Rotated Text to Line Height
	VertCompressAttr *sharedTypes.ST_OnOff
}

func NewCT_EastAsianLayout() *CT_EastAsianLayout {
	ret := &CT_EastAsianLayout{}
	return ret
}
func (m *CT_EastAsianLayout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.CombineAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:combine"},
			Value: fmt.Sprintf("%v", *m.CombineAttr)})
	}
	if m.CombineBracketsAttr != ST_CombineBracketsUnset {
		attr, err := m.CombineBracketsAttr.MarshalXMLAttr(xml.Name{Local: "w:combineBrackets"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.VertAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:vert"},
			Value: fmt.Sprintf("%v", *m.VertAttr)})
	}
	if m.VertCompressAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:vertCompress"},
			Value: fmt.Sprintf("%v", *m.VertCompressAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_EastAsianLayout) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
		}
		if attr.Name.Local == "combine" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.CombineAttr = &parsed
		}
		if attr.Name.Local == "combineBrackets" {
			m.CombineBracketsAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "vert" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.VertAttr = &parsed
		}
		if attr.Name.Local == "vertCompress" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.VertCompressAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_EastAsianLayout: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_EastAsianLayout) Validate() error {
	return m.ValidateWithPath("CT_EastAsianLayout")
}
func (m *CT_EastAsianLayout) ValidateWithPath(path string) error {
	if m.CombineAttr != nil {
		if err := m.CombineAttr.ValidateWithPath(path + "/CombineAttr"); err != nil {
			return err
		}
	}
	if err := m.CombineBracketsAttr.ValidateWithPath(path + "/CombineBracketsAttr"); err != nil {
		return err
	}
	if m.VertAttr != nil {
		if err := m.VertAttr.ValidateWithPath(path + "/VertAttr"); err != nil {
			return err
		}
	}
	if m.VertCompressAttr != nil {
		if err := m.VertCompressAttr.ValidateWithPath(path + "/VertCompressAttr"); err != nil {
			return err
		}
	}
	return nil
}
