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
	"log"
	"strconv"
)

type CT_CacheField struct {
	// PivotCache Field Name
	NameAttr string
	// PivotCache Field Caption
	CaptionAttr *string
	// Property Name
	PropertyNameAttr *string
	// Server-based Field
	ServerFieldAttr *bool
	// Unique List Retrieved
	UniqueListAttr *bool
	// Number Format Id
	NumFmtIdAttr *uint32
	// Calculated Field Formula
	FormulaAttr *string
	// SQL Data Type
	SqlTypeAttr *int32
	// Hierarchy
	HierarchyAttr *int32
	// Hierarchy Level
	LevelAttr *uint32
	// Database Field
	DatabaseFieldAttr *bool
	// Member Property Count
	MappingCountAttr *uint32
	// Member Property Field
	MemberPropertyFieldAttr *bool
	// Shared Items
	SharedItems *CT_SharedItems
	// Field Group Properties
	FieldGroup *CT_FieldGroup
	// Member Properties Map
	MpMap []*CT_X
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_CacheField() *CT_CacheField {
	ret := &CT_CacheField{}
	return ret
}

func (m *CT_CacheField) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	if m.CaptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "caption"},
			Value: fmt.Sprintf("%v", *m.CaptionAttr)})
	}
	if m.PropertyNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "propertyName"},
			Value: fmt.Sprintf("%v", *m.PropertyNameAttr)})
	}
	if m.ServerFieldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "serverField"},
			Value: fmt.Sprintf("%v", *m.ServerFieldAttr)})
	}
	if m.UniqueListAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uniqueList"},
			Value: fmt.Sprintf("%v", *m.UniqueListAttr)})
	}
	if m.NumFmtIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "numFmtId"},
			Value: fmt.Sprintf("%v", *m.NumFmtIdAttr)})
	}
	if m.FormulaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "formula"},
			Value: fmt.Sprintf("%v", *m.FormulaAttr)})
	}
	if m.SqlTypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sqlType"},
			Value: fmt.Sprintf("%v", *m.SqlTypeAttr)})
	}
	if m.HierarchyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hierarchy"},
			Value: fmt.Sprintf("%v", *m.HierarchyAttr)})
	}
	if m.LevelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "level"},
			Value: fmt.Sprintf("%v", *m.LevelAttr)})
	}
	if m.DatabaseFieldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "databaseField"},
			Value: fmt.Sprintf("%v", *m.DatabaseFieldAttr)})
	}
	if m.MappingCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "mappingCount"},
			Value: fmt.Sprintf("%v", *m.MappingCountAttr)})
	}
	if m.MemberPropertyFieldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "memberPropertyField"},
			Value: fmt.Sprintf("%v", *m.MemberPropertyFieldAttr)})
	}
	e.EncodeToken(start)
	if m.SharedItems != nil {
		sesharedItems := xml.StartElement{Name: xml.Name{Local: "x:sharedItems"}}
		e.EncodeElement(m.SharedItems, sesharedItems)
	}
	if m.FieldGroup != nil {
		sefieldGroup := xml.StartElement{Name: xml.Name{Local: "x:fieldGroup"}}
		e.EncodeElement(m.FieldGroup, sefieldGroup)
	}
	if m.MpMap != nil {
		sempMap := xml.StartElement{Name: xml.Name{Local: "x:mpMap"}}
		e.EncodeElement(m.MpMap, sempMap)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CacheField) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
		if attr.Name.Local == "caption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CaptionAttr = &parsed
		}
		if attr.Name.Local == "propertyName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PropertyNameAttr = &parsed
		}
		if attr.Name.Local == "serverField" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ServerFieldAttr = &parsed
		}
		if attr.Name.Local == "uniqueList" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UniqueListAttr = &parsed
		}
		if attr.Name.Local == "numFmtId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.NumFmtIdAttr = &pt
		}
		if attr.Name.Local == "formula" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FormulaAttr = &parsed
		}
		if attr.Name.Local == "sqlType" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.SqlTypeAttr = &pt
		}
		if attr.Name.Local == "hierarchy" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.HierarchyAttr = &pt
		}
		if attr.Name.Local == "level" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.LevelAttr = &pt
		}
		if attr.Name.Local == "databaseField" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DatabaseFieldAttr = &parsed
		}
		if attr.Name.Local == "mappingCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.MappingCountAttr = &pt
		}
		if attr.Name.Local == "memberPropertyField" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MemberPropertyFieldAttr = &parsed
		}
	}
lCT_CacheField:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sharedItems":
				m.SharedItems = NewCT_SharedItems()
				if err := d.DecodeElement(m.SharedItems, &el); err != nil {
					return err
				}
			case "fieldGroup":
				m.FieldGroup = NewCT_FieldGroup()
				if err := d.DecodeElement(m.FieldGroup, &el); err != nil {
					return err
				}
			case "mpMap":
				tmp := NewCT_X()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.MpMap = append(m.MpMap, tmp)
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CacheField
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CacheField and its children
func (m *CT_CacheField) Validate() error {
	return m.ValidateWithPath("CT_CacheField")
}

// ValidateWithPath validates the CT_CacheField and its children, prefixing error messages with path
func (m *CT_CacheField) ValidateWithPath(path string) error {
	if m.SharedItems != nil {
		if err := m.SharedItems.ValidateWithPath(path + "/SharedItems"); err != nil {
			return err
		}
	}
	if m.FieldGroup != nil {
		if err := m.FieldGroup.ValidateWithPath(path + "/FieldGroup"); err != nil {
			return err
		}
	}
	for i, v := range m.MpMap {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/MpMap[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
