// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type EG_EffectProperties struct {
	EffectLst *CT_EffectList
	EffectDag *CT_EffectContainer
}

func NewEG_EffectProperties() *EG_EffectProperties {
	ret := &EG_EffectProperties{}
	return ret
}

func (m *EG_EffectProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.EffectLst != nil {
		seeffectLst := xml.StartElement{Name: xml.Name{Local: "a:effectLst"}}
		e.EncodeElement(m.EffectLst, seeffectLst)
	}
	if m.EffectDag != nil {
		seeffectDag := xml.StartElement{Name: xml.Name{Local: "a:effectDag"}}
		e.EncodeElement(m.EffectDag, seeffectDag)
	}
	return nil
}

func (m *EG_EffectProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_EffectProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "effectLst":
				m.EffectLst = NewCT_EffectList()
				if err := d.DecodeElement(m.EffectLst, &el); err != nil {
					return err
				}
			case "effectDag":
				m.EffectDag = NewCT_EffectContainer()
				if err := d.DecodeElement(m.EffectDag, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_EffectProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_EffectProperties and its children
func (m *EG_EffectProperties) Validate() error {
	return m.ValidateWithPath("EG_EffectProperties")
}

// ValidateWithPath validates the EG_EffectProperties and its children, prefixing error messages with path
func (m *EG_EffectProperties) ValidateWithPath(path string) error {
	if m.EffectLst != nil {
		if err := m.EffectLst.ValidateWithPath(path + "/EffectLst"); err != nil {
			return err
		}
	}
	if m.EffectDag != nil {
		if err := m.EffectDag.ValidateWithPath(path + "/EffectDag"); err != nil {
			return err
		}
	}
	return nil
}
