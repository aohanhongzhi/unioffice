// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package terms

import (
	"encoding/xml"
	"fmt"
)

type ISO639_2 struct {
}

func NewISO639_2() *ISO639_2 {
	ret := &ISO639_2{}
	return ret
}

func (m *ISO639_2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Name.Local = "ISO639-2"
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *ISO639_2) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing ISO639_2: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the ISO639_2 and its children
func (m *ISO639_2) Validate() error {
	return m.ValidateWithPath("ISO639_2")
}

// ValidateWithPath validates the ISO639_2 and its children, prefixing error messages with path
func (m *ISO639_2) ValidateWithPath(path string) error {
	return nil
}
