// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetDrawing

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_ConnectorNonVisual struct {
	CNvPr      *drawingml.CT_NonVisualDrawingProps
	CNvCxnSpPr *drawingml.CT_NonVisualConnectorProperties
}

func NewCT_ConnectorNonVisual() *CT_ConnectorNonVisual {
	ret := &CT_ConnectorNonVisual{}
	ret.CNvPr = drawingml.NewCT_NonVisualDrawingProps()
	ret.CNvCxnSpPr = drawingml.NewCT_NonVisualConnectorProperties()
	return ret
}

func (m *CT_ConnectorNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvCxnSpPr := xml.StartElement{Name: xml.Name{Local: "cNvCxnSpPr"}}
	e.EncodeElement(m.CNvCxnSpPr, secNvCxnSpPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ConnectorNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = drawingml.NewCT_NonVisualDrawingProps()
	m.CNvCxnSpPr = drawingml.NewCT_NonVisualConnectorProperties()
lCT_ConnectorNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cNvPr":
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case "cNvCxnSpPr":
				if err := d.DecodeElement(m.CNvCxnSpPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ConnectorNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ConnectorNonVisual and its children
func (m *CT_ConnectorNonVisual) Validate() error {
	return m.ValidateWithPath("CT_ConnectorNonVisual")
}

// ValidateWithPath validates the CT_ConnectorNonVisual and its children, prefixing error messages with path
func (m *CT_ConnectorNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvCxnSpPr.ValidateWithPath(path + "/CNvCxnSpPr"); err != nil {
		return err
	}
	return nil
}
