// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_DrawingHF struct {
	IdAttr string
	// Left Header for Odd Pages
	LhoAttr *uint32
	// Left Header for Even Pages
	LheAttr *uint32
	// Left Header for First Page
	LhfAttr *uint32
	// Center Header for Odd Pages
	ChoAttr *uint32
	// Center Header for Even Pages
	CheAttr *uint32
	// Center Header for First Page
	ChfAttr *uint32
	// Right Header for Odd Pages
	RhoAttr *uint32
	// Right Header for Even Pages
	RheAttr *uint32
	// Right Header for First Page
	RhfAttr *uint32
	// Left Footer for Odd Pages
	LfoAttr *uint32
	// Left Footer for Even Pages
	LfeAttr *uint32
	// Left Footer for First Page
	LffAttr *uint32
	// Center Footer for Odd Pages
	CfoAttr *uint32
	// Center Footer for Even Pages
	CfeAttr *uint32
	// Center Footer for First Page
	CffAttr *uint32
	// Right Footer for Odd Pages
	RfoAttr *uint32
	// Right Footer for Even Pages
	RfeAttr *uint32
	// Right Footer for First Page
	RffAttr *uint32
}

func NewCT_DrawingHF() *CT_DrawingHF {
	ret := &CT_DrawingHF{}
	return ret
}

func (m *CT_DrawingHF) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	if m.LhoAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lho"},
			Value: fmt.Sprintf("%v", *m.LhoAttr)})
	}
	if m.LheAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lhe"},
			Value: fmt.Sprintf("%v", *m.LheAttr)})
	}
	if m.LhfAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lhf"},
			Value: fmt.Sprintf("%v", *m.LhfAttr)})
	}
	if m.ChoAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cho"},
			Value: fmt.Sprintf("%v", *m.ChoAttr)})
	}
	if m.CheAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "che"},
			Value: fmt.Sprintf("%v", *m.CheAttr)})
	}
	if m.ChfAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "chf"},
			Value: fmt.Sprintf("%v", *m.ChfAttr)})
	}
	if m.RhoAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rho"},
			Value: fmt.Sprintf("%v", *m.RhoAttr)})
	}
	if m.RheAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rhe"},
			Value: fmt.Sprintf("%v", *m.RheAttr)})
	}
	if m.RhfAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rhf"},
			Value: fmt.Sprintf("%v", *m.RhfAttr)})
	}
	if m.LfoAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lfo"},
			Value: fmt.Sprintf("%v", *m.LfoAttr)})
	}
	if m.LfeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lfe"},
			Value: fmt.Sprintf("%v", *m.LfeAttr)})
	}
	if m.LffAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lff"},
			Value: fmt.Sprintf("%v", *m.LffAttr)})
	}
	if m.CfoAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cfo"},
			Value: fmt.Sprintf("%v", *m.CfoAttr)})
	}
	if m.CfeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cfe"},
			Value: fmt.Sprintf("%v", *m.CfeAttr)})
	}
	if m.CffAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cff"},
			Value: fmt.Sprintf("%v", *m.CffAttr)})
	}
	if m.RfoAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rfo"},
			Value: fmt.Sprintf("%v", *m.RfoAttr)})
	}
	if m.RfeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rfe"},
			Value: fmt.Sprintf("%v", *m.RfeAttr)})
	}
	if m.RffAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rff"},
			Value: fmt.Sprintf("%v", *m.RffAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DrawingHF) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
		}
		if attr.Name.Local == "lho" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.LhoAttr = &pt
		}
		if attr.Name.Local == "lhe" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.LheAttr = &pt
		}
		if attr.Name.Local == "lhf" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.LhfAttr = &pt
		}
		if attr.Name.Local == "cho" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ChoAttr = &pt
		}
		if attr.Name.Local == "che" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CheAttr = &pt
		}
		if attr.Name.Local == "chf" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ChfAttr = &pt
		}
		if attr.Name.Local == "rho" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RhoAttr = &pt
		}
		if attr.Name.Local == "rhe" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RheAttr = &pt
		}
		if attr.Name.Local == "rhf" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RhfAttr = &pt
		}
		if attr.Name.Local == "lfo" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.LfoAttr = &pt
		}
		if attr.Name.Local == "lfe" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.LfeAttr = &pt
		}
		if attr.Name.Local == "lff" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.LffAttr = &pt
		}
		if attr.Name.Local == "cfo" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CfoAttr = &pt
		}
		if attr.Name.Local == "cfe" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CfeAttr = &pt
		}
		if attr.Name.Local == "cff" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CffAttr = &pt
		}
		if attr.Name.Local == "rfo" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RfoAttr = &pt
		}
		if attr.Name.Local == "rfe" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RfeAttr = &pt
		}
		if attr.Name.Local == "rff" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RffAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_DrawingHF: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_DrawingHF and its children
func (m *CT_DrawingHF) Validate() error {
	return m.ValidateWithPath("CT_DrawingHF")
}

// ValidateWithPath validates the CT_DrawingHF and its children, prefixing error messages with path
func (m *CT_DrawingHF) ValidateWithPath(path string) error {
	return nil
}
