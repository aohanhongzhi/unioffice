// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_TblGridCol struct {
	// Grid Column Width
	WAttr *sharedTypes.ST_TwipsMeasure
}

func NewCT_TblGridCol() *CT_TblGridCol {
	ret := &CT_TblGridCol{}
	return ret
}

func (m *CT_TblGridCol) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.WAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:w"},
			Value: fmt.Sprintf("%v", *m.WAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TblGridCol) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "w" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.WAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TblGridCol: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TblGridCol and its children
func (m *CT_TblGridCol) Validate() error {
	return m.ValidateWithPath("CT_TblGridCol")
}

// ValidateWithPath validates the CT_TblGridCol and its children, prefixing error messages with path
func (m *CT_TblGridCol) ValidateWithPath(path string) error {
	if m.WAttr != nil {
		if err := m.WAttr.ValidateWithPath(path + "/WAttr"); err != nil {
			return err
		}
	}
	return nil
}
