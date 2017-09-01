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

type CT_Consolidation struct {
	// Auto Page
	AutoPageAttr *bool
	// Page Item Values
	Pages *CT_Pages
	// Range Sets
	RangeSets *CT_RangeSets
}

func NewCT_Consolidation() *CT_Consolidation {
	ret := &CT_Consolidation{}
	ret.RangeSets = NewCT_RangeSets()
	return ret
}

func (m *CT_Consolidation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.AutoPageAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoPage"},
			Value: fmt.Sprintf("%v", *m.AutoPageAttr)})
	}
	e.EncodeToken(start)
	if m.Pages != nil {
		sepages := xml.StartElement{Name: xml.Name{Local: "x:pages"}}
		e.EncodeElement(m.Pages, sepages)
	}
	serangeSets := xml.StartElement{Name: xml.Name{Local: "x:rangeSets"}}
	e.EncodeElement(m.RangeSets, serangeSets)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Consolidation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.RangeSets = NewCT_RangeSets()
	for _, attr := range start.Attr {
		if attr.Name.Local == "autoPage" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoPageAttr = &parsed
		}
	}
lCT_Consolidation:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pages":
				m.Pages = NewCT_Pages()
				if err := d.DecodeElement(m.Pages, &el); err != nil {
					return err
				}
			case "rangeSets":
				if err := d.DecodeElement(m.RangeSets, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Consolidation
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Consolidation and its children
func (m *CT_Consolidation) Validate() error {
	return m.ValidateWithPath("CT_Consolidation")
}

// ValidateWithPath validates the CT_Consolidation and its children, prefixing error messages with path
func (m *CT_Consolidation) ValidateWithPath(path string) error {
	if m.Pages != nil {
		if err := m.Pages.ValidateWithPath(path + "/Pages"); err != nil {
			return err
		}
	}
	if err := m.RangeSets.ValidateWithPath(path + "/RangeSets"); err != nil {
		return err
	}
	return nil
}
