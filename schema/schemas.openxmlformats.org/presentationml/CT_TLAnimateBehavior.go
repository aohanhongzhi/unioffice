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

type CT_TLAnimateBehavior struct {
	// By
	ByAttr *string
	// From
	FromAttr *string
	// To
	ToAttr *string
	// Calculation Mode
	CalcmodeAttr ST_TLAnimateBehaviorCalcMode
	// Value Type
	ValueTypeAttr ST_TLAnimateBehaviorValueType
	CBhvr         *CT_TLCommonBehaviorData
	// Time Animated Value List
	TavLst *CT_TLTimeAnimateValueList
}

func NewCT_TLAnimateBehavior() *CT_TLAnimateBehavior {
	ret := &CT_TLAnimateBehavior{}
	ret.CBhvr = NewCT_TLCommonBehaviorData()
	return ret
}

func (m *CT_TLAnimateBehavior) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ByAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "by"},
			Value: fmt.Sprintf("%v", *m.ByAttr)})
	}
	if m.FromAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "from"},
			Value: fmt.Sprintf("%v", *m.FromAttr)})
	}
	if m.ToAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "to"},
			Value: fmt.Sprintf("%v", *m.ToAttr)})
	}
	if m.CalcmodeAttr != ST_TLAnimateBehaviorCalcModeUnset {
		attr, err := m.CalcmodeAttr.MarshalXMLAttr(xml.Name{Local: "calcmode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ValueTypeAttr != ST_TLAnimateBehaviorValueTypeUnset {
		attr, err := m.ValueTypeAttr.MarshalXMLAttr(xml.Name{Local: "valueType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	secBhvr := xml.StartElement{Name: xml.Name{Local: "p:cBhvr"}}
	e.EncodeElement(m.CBhvr, secBhvr)
	if m.TavLst != nil {
		setavLst := xml.StartElement{Name: xml.Name{Local: "p:tavLst"}}
		e.EncodeElement(m.TavLst, setavLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLAnimateBehavior) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CBhvr = NewCT_TLCommonBehaviorData()
	for _, attr := range start.Attr {
		if attr.Name.Local == "by" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ByAttr = &parsed
		}
		if attr.Name.Local == "from" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FromAttr = &parsed
		}
		if attr.Name.Local == "to" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ToAttr = &parsed
		}
		if attr.Name.Local == "calcmode" {
			m.CalcmodeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "valueType" {
			m.ValueTypeAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_TLAnimateBehavior:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cBhvr":
				if err := d.DecodeElement(m.CBhvr, &el); err != nil {
					return err
				}
			case "tavLst":
				m.TavLst = NewCT_TLTimeAnimateValueList()
				if err := d.DecodeElement(m.TavLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLAnimateBehavior
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLAnimateBehavior and its children
func (m *CT_TLAnimateBehavior) Validate() error {
	return m.ValidateWithPath("CT_TLAnimateBehavior")
}

// ValidateWithPath validates the CT_TLAnimateBehavior and its children, prefixing error messages with path
func (m *CT_TLAnimateBehavior) ValidateWithPath(path string) error {
	if err := m.CalcmodeAttr.ValidateWithPath(path + "/CalcmodeAttr"); err != nil {
		return err
	}
	if err := m.ValueTypeAttr.ValidateWithPath(path + "/ValueTypeAttr"); err != nil {
		return err
	}
	if err := m.CBhvr.ValidateWithPath(path + "/CBhvr"); err != nil {
		return err
	}
	if m.TavLst != nil {
		if err := m.TavLst.ValidateWithPath(path + "/TavLst"); err != nil {
			return err
		}
	}
	return nil
}
