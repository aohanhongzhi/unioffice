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

type TblStyleLst struct {
	CT_TableStyleList
}

func NewTblStyleLst() *TblStyleLst {
	ret := &TblStyleLst{}
	ret.CT_TableStyleList = *NewCT_TableStyleList()
	return ret
}

func (m *TblStyleLst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "a:tblStyleLst"
	return m.CT_TableStyleList.MarshalXML(e, start)
}

func (m *TblStyleLst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_TableStyleList = *NewCT_TableStyleList()
	for _, attr := range start.Attr {
		if attr.Name.Local == "def" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DefAttr = parsed
		}
	}
lTblStyleLst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tblStyle":
				tmp := NewCT_TableStyle()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TblStyle = append(m.TblStyle, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lTblStyleLst
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the TblStyleLst and its children
func (m *TblStyleLst) Validate() error {
	return m.ValidateWithPath("TblStyleLst")
}

// ValidateWithPath validates the TblStyleLst and its children, prefixing error messages with path
func (m *TblStyleLst) ValidateWithPath(path string) error {
	if err := m.CT_TableStyleList.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
