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

type CT_PivotTableStyle struct {
	// Table Style Name
	NameAttr *string
	// Show Row Header Formatting
	ShowRowHeadersAttr *bool
	// Show Table Style Column Header Formatting
	ShowColHeadersAttr *bool
	// Show Row Stripes
	ShowRowStripesAttr *bool
	// Show Column Stripes
	ShowColStripesAttr *bool
	// Show Last Column
	ShowLastColumnAttr *bool
}

func NewCT_PivotTableStyle() *CT_PivotTableStyle {
	ret := &CT_PivotTableStyle{}
	return ret
}

func (m *CT_PivotTableStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.ShowRowHeadersAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showRowHeaders"},
			Value: fmt.Sprintf("%v", *m.ShowRowHeadersAttr)})
	}
	if m.ShowColHeadersAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showColHeaders"},
			Value: fmt.Sprintf("%v", *m.ShowColHeadersAttr)})
	}
	if m.ShowRowStripesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showRowStripes"},
			Value: fmt.Sprintf("%v", *m.ShowRowStripesAttr)})
	}
	if m.ShowColStripesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showColStripes"},
			Value: fmt.Sprintf("%v", *m.ShowColStripesAttr)})
	}
	if m.ShowLastColumnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showLastColumn"},
			Value: fmt.Sprintf("%v", *m.ShowLastColumnAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PivotTableStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
		if attr.Name.Local == "showRowHeaders" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowRowHeadersAttr = &parsed
		}
		if attr.Name.Local == "showColHeaders" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowColHeadersAttr = &parsed
		}
		if attr.Name.Local == "showRowStripes" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowRowStripesAttr = &parsed
		}
		if attr.Name.Local == "showColStripes" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowColStripesAttr = &parsed
		}
		if attr.Name.Local == "showLastColumn" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowLastColumnAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PivotTableStyle: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PivotTableStyle and its children
func (m *CT_PivotTableStyle) Validate() error {
	return m.ValidateWithPath("CT_PivotTableStyle")
}

// ValidateWithPath validates the CT_PivotTableStyle and its children, prefixing error messages with path
func (m *CT_PivotTableStyle) ValidateWithPath(path string) error {
	return nil
}
