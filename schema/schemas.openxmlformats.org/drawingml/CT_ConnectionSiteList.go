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
	"log"
)

type CT_ConnectionSiteList struct {
	Cxn []*CT_ConnectionSite
}

func NewCT_ConnectionSiteList() *CT_ConnectionSiteList {
	ret := &CT_ConnectionSiteList{}
	return ret
}

func (m *CT_ConnectionSiteList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Cxn != nil {
		secxn := xml.StartElement{Name: xml.Name{Local: "a:cxn"}}
		e.EncodeElement(m.Cxn, secxn)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ConnectionSiteList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ConnectionSiteList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cxn":
				tmp := NewCT_ConnectionSite()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cxn = append(m.Cxn, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ConnectionSiteList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ConnectionSiteList and its children
func (m *CT_ConnectionSiteList) Validate() error {
	return m.ValidateWithPath("CT_ConnectionSiteList")
}

// ValidateWithPath validates the CT_ConnectionSiteList and its children, prefixing error messages with path
func (m *CT_ConnectionSiteList) ValidateWithPath(path string) error {
	for i, v := range m.Cxn {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cxn[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
