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
	"strconv"
)

type CT_SoftEdgesEffect struct {
	RadAttr int64
}

func NewCT_SoftEdgesEffect() *CT_SoftEdgesEffect {
	ret := &CT_SoftEdgesEffect{}
	ret.RadAttr = 0
	return ret
}

func (m *CT_SoftEdgesEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rad"},
		Value: fmt.Sprintf("%v", m.RadAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SoftEdgesEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.RadAttr = 0
	for _, attr := range start.Attr {
		if attr.Name.Local == "rad" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.RadAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SoftEdgesEffect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SoftEdgesEffect and its children
func (m *CT_SoftEdgesEffect) Validate() error {
	return m.ValidateWithPath("CT_SoftEdgesEffect")
}

// ValidateWithPath validates the CT_SoftEdgesEffect and its children, prefixing error messages with path
func (m *CT_SoftEdgesEffect) ValidateWithPath(path string) error {
	if m.RadAttr < 0 {
		return fmt.Errorf("%s/m.RadAttr must be >= 0 (have %v)", path, m.RadAttr)
	}
	if m.RadAttr > 27273042316900 {
		return fmt.Errorf("%s/m.RadAttr must be <= 27273042316900 (have %v)", path, m.RadAttr)
	}
	return nil
}
