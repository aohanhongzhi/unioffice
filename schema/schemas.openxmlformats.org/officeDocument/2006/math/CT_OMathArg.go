// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_OMathArg struct {
	ArgPr                *CT_OMathArgPr
	EG_OMathMathElements []*EG_OMathMathElements
	CtrlPr               *CT_CtrlPr
}

func NewCT_OMathArg() *CT_OMathArg {
	ret := &CT_OMathArg{}
	return ret
}

func (m *CT_OMathArg) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.ArgPr != nil {
		seargPr := xml.StartElement{Name: xml.Name{Local: "m:argPr"}}
		e.EncodeElement(m.ArgPr, seargPr)
	}
	if m.EG_OMathMathElements != nil {
		for _, c := range m.EG_OMathMathElements {
			c.MarshalXML(e, start)
		}
	}
	if m.CtrlPr != nil {
		sectrlPr := xml.StartElement{Name: xml.Name{Local: "m:ctrlPr"}}
		e.EncodeElement(m.CtrlPr, sectrlPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OMathArg) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_OMathArg:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "argPr":
				m.ArgPr = NewCT_OMathArgPr()
				if err := d.DecodeElement(m.ArgPr, &el); err != nil {
					return err
				}
			case "acc":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Acc = NewCT_Acc()
				if err := d.DecodeElement(tmpomathmathelements.Acc, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "bar":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Bar = NewCT_Bar()
				if err := d.DecodeElement(tmpomathmathelements.Bar, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "box":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Box = NewCT_Box()
				if err := d.DecodeElement(tmpomathmathelements.Box, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "borderBox":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.BorderBox = NewCT_BorderBox()
				if err := d.DecodeElement(tmpomathmathelements.BorderBox, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "d":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.D = NewCT_D()
				if err := d.DecodeElement(tmpomathmathelements.D, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "eqArr":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.EqArr = NewCT_EqArr()
				if err := d.DecodeElement(tmpomathmathelements.EqArr, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "f":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.F = NewCT_F()
				if err := d.DecodeElement(tmpomathmathelements.F, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "func":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Func = NewCT_Func()
				if err := d.DecodeElement(tmpomathmathelements.Func, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "groupChr":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.GroupChr = NewCT_GroupChr()
				if err := d.DecodeElement(tmpomathmathelements.GroupChr, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "limLow":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.LimLow = NewCT_LimLow()
				if err := d.DecodeElement(tmpomathmathelements.LimLow, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "limUpp":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.LimUpp = NewCT_LimUpp()
				if err := d.DecodeElement(tmpomathmathelements.LimUpp, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "m":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.M = NewCT_M()
				if err := d.DecodeElement(tmpomathmathelements.M, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "nary":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Nary = NewCT_Nary()
				if err := d.DecodeElement(tmpomathmathelements.Nary, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "phant":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Phant = NewCT_Phant()
				if err := d.DecodeElement(tmpomathmathelements.Phant, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "rad":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Rad = NewCT_Rad()
				if err := d.DecodeElement(tmpomathmathelements.Rad, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "sPre":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.SPre = NewCT_SPre()
				if err := d.DecodeElement(tmpomathmathelements.SPre, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "sSub":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.SSub = NewCT_SSub()
				if err := d.DecodeElement(tmpomathmathelements.SSub, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "sSubSup":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.SSubSup = NewCT_SSubSup()
				if err := d.DecodeElement(tmpomathmathelements.SSubSup, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "sSup":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.SSup = NewCT_SSup()
				if err := d.DecodeElement(tmpomathmathelements.SSup, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "r":
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.R = NewCT_R()
				if err := d.DecodeElement(tmpomathmathelements.R, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case "ctrlPr":
				m.CtrlPr = NewCT_CtrlPr()
				if err := d.DecodeElement(m.CtrlPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OMathArg
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OMathArg and its children
func (m *CT_OMathArg) Validate() error {
	return m.ValidateWithPath("CT_OMathArg")
}

// ValidateWithPath validates the CT_OMathArg and its children, prefixing error messages with path
func (m *CT_OMathArg) ValidateWithPath(path string) error {
	if m.ArgPr != nil {
		if err := m.ArgPr.ValidateWithPath(path + "/ArgPr"); err != nil {
			return err
		}
	}
	for i, v := range m.EG_OMathMathElements {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_OMathMathElements[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.CtrlPr != nil {
		if err := m.CtrlPr.ValidateWithPath(path + "/CtrlPr"); err != nil {
			return err
		}
	}
	return nil
}
