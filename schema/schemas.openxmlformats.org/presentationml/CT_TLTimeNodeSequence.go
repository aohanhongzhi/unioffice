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
	"strconv"
)

type CT_TLTimeNodeSequence struct {
	// Concurrent
	ConcurrentAttr *bool
	// Previous Action
	PrevAcAttr ST_TLPreviousActionType
	// Next Action
	NextAcAttr ST_TLNextActionType
	// Common TimeNode Properties
	CTn *CT_TLCommonTimeNodeData
	// Previous Conditions List
	PrevCondLst *CT_TLTimeConditionList
	// Next Conditions List
	NextCondLst *CT_TLTimeConditionList
}

func NewCT_TLTimeNodeSequence() *CT_TLTimeNodeSequence {
	ret := &CT_TLTimeNodeSequence{}
	ret.CTn = NewCT_TLCommonTimeNodeData()
	return ret
}

func (m *CT_TLTimeNodeSequence) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ConcurrentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "concurrent"},
			Value: fmt.Sprintf("%v", *m.ConcurrentAttr)})
	}
	if m.PrevAcAttr != ST_TLPreviousActionTypeUnset {
		attr, err := m.PrevAcAttr.MarshalXMLAttr(xml.Name{Local: "prevAc"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.NextAcAttr != ST_TLNextActionTypeUnset {
		attr, err := m.NextAcAttr.MarshalXMLAttr(xml.Name{Local: "nextAc"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	secTn := xml.StartElement{Name: xml.Name{Local: "p:cTn"}}
	e.EncodeElement(m.CTn, secTn)
	if m.PrevCondLst != nil {
		seprevCondLst := xml.StartElement{Name: xml.Name{Local: "p:prevCondLst"}}
		e.EncodeElement(m.PrevCondLst, seprevCondLst)
	}
	if m.NextCondLst != nil {
		senextCondLst := xml.StartElement{Name: xml.Name{Local: "p:nextCondLst"}}
		e.EncodeElement(m.NextCondLst, senextCondLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLTimeNodeSequence) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CTn = NewCT_TLCommonTimeNodeData()
	for _, attr := range start.Attr {
		if attr.Name.Local == "concurrent" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ConcurrentAttr = &parsed
		}
		if attr.Name.Local == "prevAc" {
			m.PrevAcAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "nextAc" {
			m.NextAcAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_TLTimeNodeSequence:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cTn":
				if err := d.DecodeElement(m.CTn, &el); err != nil {
					return err
				}
			case "prevCondLst":
				m.PrevCondLst = NewCT_TLTimeConditionList()
				if err := d.DecodeElement(m.PrevCondLst, &el); err != nil {
					return err
				}
			case "nextCondLst":
				m.NextCondLst = NewCT_TLTimeConditionList()
				if err := d.DecodeElement(m.NextCondLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLTimeNodeSequence
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLTimeNodeSequence and its children
func (m *CT_TLTimeNodeSequence) Validate() error {
	return m.ValidateWithPath("CT_TLTimeNodeSequence")
}

// ValidateWithPath validates the CT_TLTimeNodeSequence and its children, prefixing error messages with path
func (m *CT_TLTimeNodeSequence) ValidateWithPath(path string) error {
	if err := m.PrevAcAttr.ValidateWithPath(path + "/PrevAcAttr"); err != nil {
		return err
	}
	if err := m.NextAcAttr.ValidateWithPath(path + "/NextAcAttr"); err != nil {
		return err
	}
	if err := m.CTn.ValidateWithPath(path + "/CTn"); err != nil {
		return err
	}
	if m.PrevCondLst != nil {
		if err := m.PrevCondLst.ValidateWithPath(path + "/PrevCondLst"); err != nil {
			return err
		}
	}
	if m.NextCondLst != nil {
		if err := m.NextCondLst.ValidateWithPath(path + "/NextCondLst"); err != nil {
			return err
		}
	}
	return nil
}
