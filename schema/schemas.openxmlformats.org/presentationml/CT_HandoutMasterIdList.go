// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"
)

type CT_HandoutMasterIdList struct {
	// Handout Master ID
	HandoutMasterId *CT_HandoutMasterIdListEntry
}

func NewCT_HandoutMasterIdList() *CT_HandoutMasterIdList {
	ret := &CT_HandoutMasterIdList{}
	return ret
}

func (m *CT_HandoutMasterIdList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.HandoutMasterId != nil {
		sehandoutMasterId := xml.StartElement{Name: xml.Name{Local: "p:handoutMasterId"}}
		e.EncodeElement(m.HandoutMasterId, sehandoutMasterId)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_HandoutMasterIdList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_HandoutMasterIdList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "handoutMasterId":
				m.HandoutMasterId = NewCT_HandoutMasterIdListEntry()
				if err := d.DecodeElement(m.HandoutMasterId, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_HandoutMasterIdList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_HandoutMasterIdList and its children
func (m *CT_HandoutMasterIdList) Validate() error {
	return m.ValidateWithPath("CT_HandoutMasterIdList")
}

// ValidateWithPath validates the CT_HandoutMasterIdList and its children, prefixing error messages with path
func (m *CT_HandoutMasterIdList) ValidateWithPath(path string) error {
	if m.HandoutMasterId != nil {
		if err := m.HandoutMasterId.ValidateWithPath(path + "/HandoutMasterId"); err != nil {
			return err
		}
	}
	return nil
}
