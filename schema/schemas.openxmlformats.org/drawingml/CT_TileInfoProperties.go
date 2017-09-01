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

type CT_TileInfoProperties struct {
	TxAttr   *ST_Coordinate
	TyAttr   *ST_Coordinate
	SxAttr   *ST_Percentage
	SyAttr   *ST_Percentage
	FlipAttr ST_TileFlipMode
	AlgnAttr ST_RectAlignment
}

func NewCT_TileInfoProperties() *CT_TileInfoProperties {
	ret := &CT_TileInfoProperties{}
	return ret
}

func (m *CT_TileInfoProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.TxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tx"},
			Value: fmt.Sprintf("%v", *m.TxAttr)})
	}
	if m.TyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ty"},
			Value: fmt.Sprintf("%v", *m.TyAttr)})
	}
	if m.SxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sx"},
			Value: fmt.Sprintf("%v", *m.SxAttr)})
	}
	if m.SyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sy"},
			Value: fmt.Sprintf("%v", *m.SyAttr)})
	}
	if m.FlipAttr != ST_TileFlipModeUnset {
		attr, err := m.FlipAttr.MarshalXMLAttr(xml.Name{Local: "flip"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AlgnAttr != ST_RectAlignmentUnset {
		attr, err := m.AlgnAttr.MarshalXMLAttr(xml.Name{Local: "algn"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TileInfoProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "tx" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.TxAttr = &parsed
		}
		if attr.Name.Local == "ty" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.TyAttr = &parsed
		}
		if attr.Name.Local == "sx" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.SxAttr = &parsed
		}
		if attr.Name.Local == "sy" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.SyAttr = &parsed
		}
		if attr.Name.Local == "flip" {
			m.FlipAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "algn" {
			m.AlgnAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TileInfoProperties: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TileInfoProperties and its children
func (m *CT_TileInfoProperties) Validate() error {
	return m.ValidateWithPath("CT_TileInfoProperties")
}

// ValidateWithPath validates the CT_TileInfoProperties and its children, prefixing error messages with path
func (m *CT_TileInfoProperties) ValidateWithPath(path string) error {
	if m.TxAttr != nil {
		if err := m.TxAttr.ValidateWithPath(path + "/TxAttr"); err != nil {
			return err
		}
	}
	if m.TyAttr != nil {
		if err := m.TyAttr.ValidateWithPath(path + "/TyAttr"); err != nil {
			return err
		}
	}
	if m.SxAttr != nil {
		if err := m.SxAttr.ValidateWithPath(path + "/SxAttr"); err != nil {
			return err
		}
	}
	if m.SyAttr != nil {
		if err := m.SyAttr.ValidateWithPath(path + "/SyAttr"); err != nil {
			return err
		}
	}
	if err := m.FlipAttr.ValidateWithPath(path + "/FlipAttr"); err != nil {
		return err
	}
	if err := m.AlgnAttr.ValidateWithPath(path + "/AlgnAttr"); err != nil {
		return err
	}
	return nil
}
