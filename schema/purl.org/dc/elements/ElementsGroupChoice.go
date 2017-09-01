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
	"log"
)

type ElementsGroupChoice struct {
	Any []*Any
}

func NewElementsGroupChoice() *ElementsGroupChoice {
	ret := &ElementsGroupChoice{}
	return ret
}

func (m *ElementsGroupChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Any != nil {
		seany := xml.StartElement{Name: xml.Name{Local: "dc:any"}}
		e.EncodeElement(m.Any, seany)
	}
	return nil
}

func (m *ElementsGroupChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lElementsGroupChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "any":
				tmp := NewAny()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Any = append(m.Any, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lElementsGroupChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ElementsGroupChoice and its children
func (m *ElementsGroupChoice) Validate() error {
	return m.ValidateWithPath("ElementsGroupChoice")
}

// ValidateWithPath validates the ElementsGroupChoice and its children, prefixing error messages with path
func (m *ElementsGroupChoice) ValidateWithPath(path string) error {
	for i, v := range m.Any {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Any[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
