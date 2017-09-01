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

type CT_QueryCache struct {
	// Cached Query Count
	CountAttr *uint32
	// Query
	Query []*CT_Query
}

func NewCT_QueryCache() *CT_QueryCache {
	ret := &CT_QueryCache{}
	return ret
}

func (m *CT_QueryCache) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	sequery := xml.StartElement{Name: xml.Name{Local: "x:query"}}
	e.EncodeElement(m.Query, sequery)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_QueryCache) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_QueryCache:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "query":
				tmp := NewCT_Query()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Query = append(m.Query, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_QueryCache
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_QueryCache and its children
func (m *CT_QueryCache) Validate() error {
	return m.ValidateWithPath("CT_QueryCache")
}

// ValidateWithPath validates the CT_QueryCache and its children, prefixing error messages with path
func (m *CT_QueryCache) ValidateWithPath(path string) error {
	for i, v := range m.Query {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Query[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
