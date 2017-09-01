// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"
)

type CT_AxDataSource struct {
	Choice *CT_AxDataSourceChoice
}

func NewCT_AxDataSource() *CT_AxDataSource {
	ret := &CT_AxDataSource{}
	ret.Choice = NewCT_AxDataSourceChoice()
	return ret
}

func (m *CT_AxDataSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	m.Choice.MarshalXML(e, start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AxDataSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Choice = NewCT_AxDataSourceChoice()
lCT_AxDataSource:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "multiLvlStrRef":
				m.Choice = NewCT_AxDataSourceChoice()
				if err := d.DecodeElement(&m.Choice.MultiLvlStrRef, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "numRef":
				m.Choice = NewCT_AxDataSourceChoice()
				if err := d.DecodeElement(&m.Choice.NumRef, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "numLit":
				m.Choice = NewCT_AxDataSourceChoice()
				if err := d.DecodeElement(&m.Choice.NumLit, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "strRef":
				m.Choice = NewCT_AxDataSourceChoice()
				if err := d.DecodeElement(&m.Choice.StrRef, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "strLit":
				m.Choice = NewCT_AxDataSourceChoice()
				if err := d.DecodeElement(&m.Choice.StrLit, &el); err != nil {
					return err
				}
				_ = m.Choice
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AxDataSource
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AxDataSource and its children
func (m *CT_AxDataSource) Validate() error {
	return m.ValidateWithPath("CT_AxDataSource")
}

// ValidateWithPath validates the CT_AxDataSource and its children, prefixing error messages with path
func (m *CT_AxDataSource) ValidateWithPath(path string) error {
	if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
		return err
	}
	return nil
}
