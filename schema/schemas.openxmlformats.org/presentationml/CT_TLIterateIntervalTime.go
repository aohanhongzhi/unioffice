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
)

type CT_TLIterateIntervalTime struct {
	// Time
	ValAttr ST_TLTime
}

func NewCT_TLIterateIntervalTime() *CT_TLIterateIntervalTime {
	ret := &CT_TLIterateIntervalTime{}
	return ret
}

func (m *CT_TLIterateIntervalTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLIterateIntervalTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_TLTime(attr.Value)
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
			return fmt.Errorf("parsing CT_TLIterateIntervalTime: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TLIterateIntervalTime and its children
func (m *CT_TLIterateIntervalTime) Validate() error {
	return m.ValidateWithPath("CT_TLIterateIntervalTime")
}

// ValidateWithPath validates the CT_TLIterateIntervalTime and its children, prefixing error messages with path
func (m *CT_TLIterateIntervalTime) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
