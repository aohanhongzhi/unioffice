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

type AG_AutoFormat struct {
	AutoFormatIdAttr            *uint32
	ApplyNumberFormatsAttr      *bool
	ApplyBorderFormatsAttr      *bool
	ApplyFontFormatsAttr        *bool
	ApplyPatternFormatsAttr     *bool
	ApplyAlignmentFormatsAttr   *bool
	ApplyWidthHeightFormatsAttr *bool
}

func NewAG_AutoFormat() *AG_AutoFormat {
	ret := &AG_AutoFormat{}
	return ret
}

func (m *AG_AutoFormat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.AutoFormatIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoFormatId"},
			Value: fmt.Sprintf("%v", *m.AutoFormatIdAttr)})
	}
	if m.ApplyNumberFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyNumberFormats"},
			Value: fmt.Sprintf("%v", *m.ApplyNumberFormatsAttr)})
	}
	if m.ApplyBorderFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyBorderFormats"},
			Value: fmt.Sprintf("%v", *m.ApplyBorderFormatsAttr)})
	}
	if m.ApplyFontFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyFontFormats"},
			Value: fmt.Sprintf("%v", *m.ApplyFontFormatsAttr)})
	}
	if m.ApplyPatternFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyPatternFormats"},
			Value: fmt.Sprintf("%v", *m.ApplyPatternFormatsAttr)})
	}
	if m.ApplyAlignmentFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyAlignmentFormats"},
			Value: fmt.Sprintf("%v", *m.ApplyAlignmentFormatsAttr)})
	}
	if m.ApplyWidthHeightFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyWidthHeightFormats"},
			Value: fmt.Sprintf("%v", *m.ApplyWidthHeightFormatsAttr)})
	}
	return nil
}

func (m *AG_AutoFormat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "autoFormatId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.AutoFormatIdAttr = &pt
		}
		if attr.Name.Local == "applyNumberFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyNumberFormatsAttr = &parsed
		}
		if attr.Name.Local == "applyBorderFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyBorderFormatsAttr = &parsed
		}
		if attr.Name.Local == "applyFontFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyFontFormatsAttr = &parsed
		}
		if attr.Name.Local == "applyPatternFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyPatternFormatsAttr = &parsed
		}
		if attr.Name.Local == "applyAlignmentFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyAlignmentFormatsAttr = &parsed
		}
		if attr.Name.Local == "applyWidthHeightFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyWidthHeightFormatsAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_AutoFormat: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_AutoFormat and its children
func (m *AG_AutoFormat) Validate() error {
	return m.ValidateWithPath("AG_AutoFormat")
}

// ValidateWithPath validates the AG_AutoFormat and its children, prefixing error messages with path
func (m *AG_AutoFormat) ValidateWithPath(path string) error {
	return nil
}
