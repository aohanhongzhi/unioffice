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
)

type CT_SignedTwipsMeasure struct {
	// Positive or Negative Value in Twentieths of a Point
	ValAttr ST_SignedTwipsMeasure
}

func NewCT_SignedTwipsMeasure() *CT_SignedTwipsMeasure {
	ret := &CT_SignedTwipsMeasure{}
	return ret
}

func (m *CT_SignedTwipsMeasure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SignedTwipsMeasure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_SignedTwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SignedTwipsMeasure: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SignedTwipsMeasure and its children
func (m *CT_SignedTwipsMeasure) Validate() error {
	return m.ValidateWithPath("CT_SignedTwipsMeasure")
}

// ValidateWithPath validates the CT_SignedTwipsMeasure and its children, prefixing error messages with path
func (m *CT_SignedTwipsMeasure) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
