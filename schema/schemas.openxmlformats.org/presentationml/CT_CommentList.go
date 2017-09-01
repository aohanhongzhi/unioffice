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
	"log"
)

type CT_CommentList struct {
	// Comment
	Cm []*CT_Comment
}

func NewCT_CommentList() *CT_CommentList {
	ret := &CT_CommentList{}
	return ret
}

func (m *CT_CommentList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Cm != nil {
		secm := xml.StartElement{Name: xml.Name{Local: "p:cm"}}
		e.EncodeElement(m.Cm, secm)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CommentList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CommentList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cm":
				tmp := NewCT_Comment()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cm = append(m.Cm, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CommentList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CommentList and its children
func (m *CT_CommentList) Validate() error {
	return m.ValidateWithPath("CT_CommentList")
}

// ValidateWithPath validates the CT_CommentList and its children, prefixing error messages with path
func (m *CT_CommentList) ValidateWithPath(path string) error {
	for i, v := range m.Cm {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cm[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
