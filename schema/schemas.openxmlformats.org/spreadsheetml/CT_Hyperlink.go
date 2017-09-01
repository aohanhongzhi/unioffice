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
)

type CT_Hyperlink struct {
	// Reference
	RefAttr string
	IdAttr  *string
	// Location
	LocationAttr *string
	// Tool Tip
	TooltipAttr *string
	// Display String
	DisplayAttr *string
}

func NewCT_Hyperlink() *CT_Hyperlink {
	ret := &CT_Hyperlink{}
	return ret
}

func (m *CT_Hyperlink) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ref"},
		Value: fmt.Sprintf("%v", m.RefAttr)})
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.LocationAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "location"},
			Value: fmt.Sprintf("%v", *m.LocationAttr)})
	}
	if m.TooltipAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tooltip"},
			Value: fmt.Sprintf("%v", *m.TooltipAttr)})
	}
	if m.DisplayAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "display"},
			Value: fmt.Sprintf("%v", *m.DisplayAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Hyperlink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "ref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefAttr = parsed
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
		}
		if attr.Name.Local == "location" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LocationAttr = &parsed
		}
		if attr.Name.Local == "tooltip" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TooltipAttr = &parsed
		}
		if attr.Name.Local == "display" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DisplayAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Hyperlink: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Hyperlink and its children
func (m *CT_Hyperlink) Validate() error {
	return m.ValidateWithPath("CT_Hyperlink")
}

// ValidateWithPath validates the CT_Hyperlink and its children, prefixing error messages with path
func (m *CT_Hyperlink) ValidateWithPath(path string) error {
	return nil
}
