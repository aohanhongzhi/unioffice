// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"log"
)

type Revisions struct {
	CT_Revisions
}

func NewRevisions() *Revisions {
	ret := &Revisions{}
	ret.CT_Revisions = *NewCT_Revisions()
	return ret
}

func (m *Revisions) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "x:revisions"
	return m.CT_Revisions.MarshalXML(e, start)
}

func (m *Revisions) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Revisions = *NewCT_Revisions()
lRevisions:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "rrc":
				tmp := NewCT_RevisionRowColumn()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rrc = append(m.Rrc, tmp)
			case "rm":
				tmp := NewCT_RevisionMove()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rm = append(m.Rm, tmp)
			case "rcv":
				tmp := NewCT_RevisionCustomView()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcv = append(m.Rcv, tmp)
			case "rsnm":
				tmp := NewCT_RevisionSheetRename()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rsnm = append(m.Rsnm, tmp)
			case "ris":
				tmp := NewCT_RevisionInsertSheet()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ris = append(m.Ris, tmp)
			case "rcc":
				tmp := NewCT_RevisionCellChange()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcc = append(m.Rcc, tmp)
			case "rfmt":
				tmp := NewCT_RevisionFormatting()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rfmt = append(m.Rfmt, tmp)
			case "raf":
				tmp := NewCT_RevisionAutoFormatting()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Raf = append(m.Raf, tmp)
			case "rdn":
				tmp := NewCT_RevisionDefinedName()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rdn = append(m.Rdn, tmp)
			case "rcmt":
				tmp := NewCT_RevisionComment()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcmt = append(m.Rcmt, tmp)
			case "rqt":
				tmp := NewCT_RevisionQueryTableField()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rqt = append(m.Rqt, tmp)
			case "rcft":
				tmp := NewCT_RevisionConflict()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcft = append(m.Rcft, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lRevisions
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Revisions and its children
func (m *Revisions) Validate() error {
	return m.ValidateWithPath("Revisions")
}

// ValidateWithPath validates the Revisions and its children, prefixing error messages with path
func (m *Revisions) ValidateWithPath(path string) error {
	if err := m.CT_Revisions.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
