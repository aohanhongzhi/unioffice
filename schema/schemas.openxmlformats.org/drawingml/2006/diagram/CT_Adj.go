// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_Adj struct {
	IdxAttr uint32
	ValAttr float64
}

func NewCT_Adj() *CT_Adj {
	ret := &CT_Adj{}
	ret.IdxAttr = 1
	return ret
}

func (m *CT_Adj) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "idx"},
		Value: fmt.Sprintf("%v", m.IdxAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Adj) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.IdxAttr = 1
	for _, attr := range start.Attr {
		if attr.Name.Local == "idx" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdxAttr = uint32(parsed)
		}
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
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
			return fmt.Errorf("parsing CT_Adj: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Adj and its children
func (m *CT_Adj) Validate() error {
	return m.ValidateWithPath("CT_Adj")
}

// ValidateWithPath validates the CT_Adj and its children, prefixing error messages with path
func (m *CT_Adj) ValidateWithPath(path string) error {
	if m.IdxAttr < 1 {
		return fmt.Errorf("%s/m.IdxAttr must be >= 1 (have %v)", path, m.IdxAttr)
	}
	return nil
}
