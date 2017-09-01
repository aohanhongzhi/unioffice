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
)

type CT_PixelsMeasure struct {
	// Measurement in Pixels
	ValAttr uint64
}

func NewCT_PixelsMeasure() *CT_PixelsMeasure {
	ret := &CT_PixelsMeasure{}
	return ret
}

func (m *CT_PixelsMeasure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PixelsMeasure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.ValAttr = uint64(parsed)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PixelsMeasure: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PixelsMeasure and its children
func (m *CT_PixelsMeasure) Validate() error {
	return m.ValidateWithPath("CT_PixelsMeasure")
}

// ValidateWithPath validates the CT_PixelsMeasure and its children, prefixing error messages with path
func (m *CT_PixelsMeasure) ValidateWithPath(path string) error {
	return nil
}
