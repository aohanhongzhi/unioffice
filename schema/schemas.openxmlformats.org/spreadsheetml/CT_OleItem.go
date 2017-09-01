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

type CT_OleItem struct {
	// Object Name
	NameAttr string
	// Icon
	IconAttr *bool
	// Advise
	AdviseAttr *bool
	// Object is an Image
	PreferPicAttr *bool
}

func NewCT_OleItem() *CT_OleItem {
	ret := &CT_OleItem{}
	return ret
}

func (m *CT_OleItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	if m.IconAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "icon"},
			Value: fmt.Sprintf("%v", *m.IconAttr)})
	}
	if m.AdviseAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "advise"},
			Value: fmt.Sprintf("%v", *m.AdviseAttr)})
	}
	if m.PreferPicAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "preferPic"},
			Value: fmt.Sprintf("%v", *m.PreferPicAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OleItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
		if attr.Name.Local == "icon" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IconAttr = &parsed
		}
		if attr.Name.Local == "advise" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AdviseAttr = &parsed
		}
		if attr.Name.Local == "preferPic" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PreferPicAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_OleItem: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_OleItem and its children
func (m *CT_OleItem) Validate() error {
	return m.ValidateWithPath("CT_OleItem")
}

// ValidateWithPath validates the CT_OleItem and its children, prefixing error messages with path
func (m *CT_OleItem) ValidateWithPath(path string) error {
	return nil
}
