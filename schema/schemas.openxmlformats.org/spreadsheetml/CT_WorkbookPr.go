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

type CT_WorkbookPr struct {
	// Date 1904
	Date1904Attr *bool
	// Show Objects
	ShowObjectsAttr ST_Objects
	// Show Border Unselected Table
	ShowBorderUnselectedTablesAttr *bool
	// Filter Privacy
	FilterPrivacyAttr *bool
	// Prompted Solutions
	PromptedSolutionsAttr *bool
	// Show Ink Annotations
	ShowInkAnnotationAttr *bool
	// Create Backup File
	BackupFileAttr *bool
	// Save External Link Values
	SaveExternalLinkValuesAttr *bool
	// Update Links Behavior
	UpdateLinksAttr ST_UpdateLinks
	// Code Name
	CodeNameAttr *string
	// Hide Pivot Field List
	HidePivotFieldListAttr *bool
	// Show Pivot Chart Filter
	ShowPivotChartFilterAttr *bool
	// Allow Refresh Query
	AllowRefreshQueryAttr *bool
	// Publish Items
	PublishItemsAttr *bool
	// Check Compatibility On Save
	CheckCompatibilityAttr *bool
	// Auto Compress Pictures
	AutoCompressPicturesAttr *bool
	// Refresh all Connections on Open
	RefreshAllConnectionsAttr *bool
	// Default Theme Version
	DefaultThemeVersionAttr *uint32
}

func NewCT_WorkbookPr() *CT_WorkbookPr {
	ret := &CT_WorkbookPr{}
	return ret
}

func (m *CT_WorkbookPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Date1904Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "date1904"},
			Value: fmt.Sprintf("%v", *m.Date1904Attr)})
	}
	if m.ShowObjectsAttr != ST_ObjectsUnset {
		attr, err := m.ShowObjectsAttr.MarshalXMLAttr(xml.Name{Local: "showObjects"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ShowBorderUnselectedTablesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showBorderUnselectedTables"},
			Value: fmt.Sprintf("%v", *m.ShowBorderUnselectedTablesAttr)})
	}
	if m.FilterPrivacyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "filterPrivacy"},
			Value: fmt.Sprintf("%v", *m.FilterPrivacyAttr)})
	}
	if m.PromptedSolutionsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "promptedSolutions"},
			Value: fmt.Sprintf("%v", *m.PromptedSolutionsAttr)})
	}
	if m.ShowInkAnnotationAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showInkAnnotation"},
			Value: fmt.Sprintf("%v", *m.ShowInkAnnotationAttr)})
	}
	if m.BackupFileAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "backupFile"},
			Value: fmt.Sprintf("%v", *m.BackupFileAttr)})
	}
	if m.SaveExternalLinkValuesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "saveExternalLinkValues"},
			Value: fmt.Sprintf("%v", *m.SaveExternalLinkValuesAttr)})
	}
	if m.UpdateLinksAttr != ST_UpdateLinksUnset {
		attr, err := m.UpdateLinksAttr.MarshalXMLAttr(xml.Name{Local: "updateLinks"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CodeNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "codeName"},
			Value: fmt.Sprintf("%v", *m.CodeNameAttr)})
	}
	if m.HidePivotFieldListAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidePivotFieldList"},
			Value: fmt.Sprintf("%v", *m.HidePivotFieldListAttr)})
	}
	if m.ShowPivotChartFilterAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showPivotChartFilter"},
			Value: fmt.Sprintf("%v", *m.ShowPivotChartFilterAttr)})
	}
	if m.AllowRefreshQueryAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "allowRefreshQuery"},
			Value: fmt.Sprintf("%v", *m.AllowRefreshQueryAttr)})
	}
	if m.PublishItemsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "publishItems"},
			Value: fmt.Sprintf("%v", *m.PublishItemsAttr)})
	}
	if m.CheckCompatibilityAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "checkCompatibility"},
			Value: fmt.Sprintf("%v", *m.CheckCompatibilityAttr)})
	}
	if m.AutoCompressPicturesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoCompressPictures"},
			Value: fmt.Sprintf("%v", *m.AutoCompressPicturesAttr)})
	}
	if m.RefreshAllConnectionsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "refreshAllConnections"},
			Value: fmt.Sprintf("%v", *m.RefreshAllConnectionsAttr)})
	}
	if m.DefaultThemeVersionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "defaultThemeVersion"},
			Value: fmt.Sprintf("%v", *m.DefaultThemeVersionAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_WorkbookPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "date1904" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.Date1904Attr = &parsed
		}
		if attr.Name.Local == "showObjects" {
			m.ShowObjectsAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "showBorderUnselectedTables" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowBorderUnselectedTablesAttr = &parsed
		}
		if attr.Name.Local == "filterPrivacy" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FilterPrivacyAttr = &parsed
		}
		if attr.Name.Local == "promptedSolutions" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PromptedSolutionsAttr = &parsed
		}
		if attr.Name.Local == "showInkAnnotation" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowInkAnnotationAttr = &parsed
		}
		if attr.Name.Local == "backupFile" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BackupFileAttr = &parsed
		}
		if attr.Name.Local == "saveExternalLinkValues" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SaveExternalLinkValuesAttr = &parsed
		}
		if attr.Name.Local == "updateLinks" {
			m.UpdateLinksAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "codeName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CodeNameAttr = &parsed
		}
		if attr.Name.Local == "hidePivotFieldList" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HidePivotFieldListAttr = &parsed
		}
		if attr.Name.Local == "showPivotChartFilter" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowPivotChartFilterAttr = &parsed
		}
		if attr.Name.Local == "allowRefreshQuery" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AllowRefreshQueryAttr = &parsed
		}
		if attr.Name.Local == "publishItems" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PublishItemsAttr = &parsed
		}
		if attr.Name.Local == "checkCompatibility" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CheckCompatibilityAttr = &parsed
		}
		if attr.Name.Local == "autoCompressPictures" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoCompressPicturesAttr = &parsed
		}
		if attr.Name.Local == "refreshAllConnections" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RefreshAllConnectionsAttr = &parsed
		}
		if attr.Name.Local == "defaultThemeVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DefaultThemeVersionAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_WorkbookPr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_WorkbookPr and its children
func (m *CT_WorkbookPr) Validate() error {
	return m.ValidateWithPath("CT_WorkbookPr")
}

// ValidateWithPath validates the CT_WorkbookPr and its children, prefixing error messages with path
func (m *CT_WorkbookPr) ValidateWithPath(path string) error {
	if err := m.ShowObjectsAttr.ValidateWithPath(path + "/ShowObjectsAttr"); err != nil {
		return err
	}
	if err := m.UpdateLinksAttr.ValidateWithPath(path + "/UpdateLinksAttr"); err != nil {
		return err
	}
	return nil
}
