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

type EG_FillModeProperties struct {
	Tile    *CT_TileInfoProperties
	Stretch *CT_StretchInfoProperties
}

func NewEG_FillModeProperties() *EG_FillModeProperties {
	ret := &EG_FillModeProperties{}
	return ret
}

func (m *EG_FillModeProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Tile != nil {
		setile := xml.StartElement{Name: xml.Name{Local: "a:tile"}}
		e.EncodeElement(m.Tile, setile)
	}
	if m.Stretch != nil {
		sestretch := xml.StartElement{Name: xml.Name{Local: "a:stretch"}}
		e.EncodeElement(m.Stretch, sestretch)
	}
	return nil
}

func (m *EG_FillModeProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_FillModeProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tile":
				m.Tile = NewCT_TileInfoProperties()
				if err := d.DecodeElement(m.Tile, &el); err != nil {
					return err
				}
			case "stretch":
				m.Stretch = NewCT_StretchInfoProperties()
				if err := d.DecodeElement(m.Stretch, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_FillModeProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_FillModeProperties and its children
func (m *EG_FillModeProperties) Validate() error {
	return m.ValidateWithPath("EG_FillModeProperties")
}

// ValidateWithPath validates the EG_FillModeProperties and its children, prefixing error messages with path
func (m *EG_FillModeProperties) ValidateWithPath(path string) error {
	if m.Tile != nil {
		if err := m.Tile.ValidateWithPath(path + "/Tile"); err != nil {
			return err
		}
	}
	if m.Stretch != nil {
		if err := m.Stretch.ValidateWithPath(path + "/Stretch"); err != nil {
			return err
		}
	}
	return nil
}
