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

type CT_MeasureDimensionMap struct {
	// Measure Group Id
	MeasureGroupAttr *uint32
	// Dimension Id
	DimensionAttr *uint32
}

func NewCT_MeasureDimensionMap() *CT_MeasureDimensionMap {
	ret := &CT_MeasureDimensionMap{}
	return ret
}

func (m *CT_MeasureDimensionMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.MeasureGroupAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "measureGroup"},
			Value: fmt.Sprintf("%v", *m.MeasureGroupAttr)})
	}
	if m.DimensionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dimension"},
			Value: fmt.Sprintf("%v", *m.DimensionAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MeasureDimensionMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "measureGroup" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.MeasureGroupAttr = &pt
		}
		if attr.Name.Local == "dimension" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DimensionAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_MeasureDimensionMap: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_MeasureDimensionMap and its children
func (m *CT_MeasureDimensionMap) Validate() error {
	return m.ValidateWithPath("CT_MeasureDimensionMap")
}

// ValidateWithPath validates the CT_MeasureDimensionMap and its children, prefixing error messages with path
func (m *CT_MeasureDimensionMap) ValidateWithPath(path string) error {
	return nil
}
