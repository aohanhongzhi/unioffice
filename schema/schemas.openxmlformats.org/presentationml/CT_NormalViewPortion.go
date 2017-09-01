// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_NormalViewPortion struct {
	// Normal View Dimension Size
	SzAttr drawingml.ST_PositiveFixedPercentage
	// Auto Adjust Normal View
	AutoAdjustAttr *bool
}

func NewCT_NormalViewPortion() *CT_NormalViewPortion {
	ret := &CT_NormalViewPortion{}
	return ret
}

func (m *CT_NormalViewPortion) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sz"},
		Value: fmt.Sprintf("%v", m.SzAttr)})
	if m.AutoAdjustAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoAdjust"},
			Value: fmt.Sprintf("%v", *m.AutoAdjustAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NormalViewPortion) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "sz" {
			parsed, err := ParseUnionST_PositiveFixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.SzAttr = parsed
		}
		if attr.Name.Local == "autoAdjust" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoAdjustAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_NormalViewPortion: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_NormalViewPortion and its children
func (m *CT_NormalViewPortion) Validate() error {
	return m.ValidateWithPath("CT_NormalViewPortion")
}

// ValidateWithPath validates the CT_NormalViewPortion and its children, prefixing error messages with path
func (m *CT_NormalViewPortion) ValidateWithPath(path string) error {
	if err := m.SzAttr.ValidateWithPath(path + "/SzAttr"); err != nil {
		return err
	}
	return nil
}
