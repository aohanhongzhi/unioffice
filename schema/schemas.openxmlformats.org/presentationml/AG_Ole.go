// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type AG_Ole struct {
	SpidAttr       *string
	NameAttr       *string
	ShowAsIconAttr *bool
	IdAttr         *string
	ImgWAttr       *int32
	ImgHAttr       *int32
}

func NewAG_Ole() *AG_Ole {
	ret := &AG_Ole{}
	return ret
}

func (m *AG_Ole) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.SpidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spid"},
			Value: fmt.Sprintf("%v", *m.SpidAttr)})
	}
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.ShowAsIconAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showAsIcon"},
			Value: fmt.Sprintf("%v", *m.ShowAsIconAttr)})
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.ImgWAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "imgW"},
			Value: fmt.Sprintf("%v", *m.ImgWAttr)})
	}
	if m.ImgHAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "imgH"},
			Value: fmt.Sprintf("%v", *m.ImgHAttr)})
	}
	return nil
}

func (m *AG_Ole) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "spid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SpidAttr = &parsed
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
		if attr.Name.Local == "showAsIcon" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowAsIconAttr = &parsed
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
		}
		if attr.Name.Local == "imgW" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.ImgWAttr = &pt
		}
		if attr.Name.Local == "imgH" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.ImgHAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_Ole: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_Ole and its children
func (m *AG_Ole) Validate() error {
	return m.ValidateWithPath("AG_Ole")
}

// ValidateWithPath validates the AG_Ole and its children, prefixing error messages with path
func (m *AG_Ole) ValidateWithPath(path string) error {
	if m.ImgWAttr != nil {
		if *m.ImgWAttr < 0 {
			return fmt.Errorf("%s/m.ImgWAttr must be >= 0 (have %v)", path, *m.ImgWAttr)
		}
	}
	if m.ImgHAttr != nil {
		if *m.ImgHAttr < 0 {
			return fmt.Errorf("%s/m.ImgHAttr must be >= 0 (have %v)", path, *m.ImgHAttr)
		}
	}
	return nil
}
