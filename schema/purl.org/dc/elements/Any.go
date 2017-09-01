// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package elements

import (
	"encoding/xml"
	"fmt"
)

type Any struct {
	SimpleLiteral
}

func NewAny() *Any {
	ret := &Any{}
	ret.SimpleLiteral = *NewSimpleLiteral()
	return ret
}

func (m *Any) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	return m.SimpleLiteral.MarshalXML(e, start)
}

func (m *Any) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SimpleLiteral = *NewSimpleLiteral()
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Any: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Any and its children
func (m *Any) Validate() error {
	return m.ValidateWithPath("Any")
}

// ValidateWithPath validates the Any and its children, prefixing error messages with path
func (m *Any) ValidateWithPath(path string) error {
	if err := m.SimpleLiteral.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
