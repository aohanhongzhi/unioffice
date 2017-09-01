// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_MdxMemeberProp struct {
	// Member Unique Name Index
	NAttr uint32
	// Property Name Index
	NpAttr uint32
}

func NewCT_MdxMemeberProp() *CT_MdxMemeberProp {
	ret := &CT_MdxMemeberProp{}
	return ret
}

func (m *CT_MdxMemeberProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "n"},
		Value: fmt.Sprintf("%v", m.NAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "np"},
		Value: fmt.Sprintf("%v", m.NpAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MdxMemeberProp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "n" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.NAttr = uint32(parsed)
		}
		if attr.Name.Local == "np" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.NpAttr = uint32(parsed)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_MdxMemeberProp: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_MdxMemeberProp and its children
func (m *CT_MdxMemeberProp) Validate() error {
	return m.ValidateWithPath("CT_MdxMemeberProp")
}

// ValidateWithPath validates the CT_MdxMemeberProp and its children, prefixing error messages with path
func (m *CT_MdxMemeberProp) ValidateWithPath(path string) error {
	return nil
}
