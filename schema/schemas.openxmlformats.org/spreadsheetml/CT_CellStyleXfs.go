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
	"log"
	"strconv"
)

type CT_CellStyleXfs struct {
	// Style Count
	CountAttr *uint32
	// Formatting Elements
	Xf []*CT_Xf
}

func NewCT_CellStyleXfs() *CT_CellStyleXfs {
	ret := &CT_CellStyleXfs{}
	return ret
}

func (m *CT_CellStyleXfs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	sexf := xml.StartElement{Name: xml.Name{Local: "x:xf"}}
	e.EncodeElement(m.Xf, sexf)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CellStyleXfs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
	}
lCT_CellStyleXfs:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "xf":
				tmp := NewCT_Xf()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Xf = append(m.Xf, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CellStyleXfs
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CellStyleXfs and its children
func (m *CT_CellStyleXfs) Validate() error {
	return m.ValidateWithPath("CT_CellStyleXfs")
}

// ValidateWithPath validates the CT_CellStyleXfs and its children, prefixing error messages with path
func (m *CT_CellStyleXfs) ValidateWithPath(path string) error {
	for i, v := range m.Xf {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Xf[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
