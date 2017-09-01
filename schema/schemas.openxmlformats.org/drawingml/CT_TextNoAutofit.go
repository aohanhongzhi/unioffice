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
)

type CT_TextNoAutofit struct {
}

func NewCT_TextNoAutofit() *CT_TextNoAutofit {
	ret := &CT_TextNoAutofit{}
	return ret
}

func (m *CT_TextNoAutofit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextNoAutofit) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextNoAutofit: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TextNoAutofit and its children
func (m *CT_TextNoAutofit) Validate() error {
	return m.ValidateWithPath("CT_TextNoAutofit")
}

// ValidateWithPath validates the CT_TextNoAutofit and its children, prefixing error messages with path
func (m *CT_TextNoAutofit) ValidateWithPath(path string) error {
	return nil
}
