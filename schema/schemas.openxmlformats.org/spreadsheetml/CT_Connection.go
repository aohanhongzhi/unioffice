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

type CT_Connection struct {
	// Connection Id
	IdAttr uint32
	// Source Database File
	SourceFileAttr *string
	// Connection File
	OdcFileAttr *string
	// Keep Connection Open
	KeepAliveAttr *bool
	// Automatic Refresh Interval
	IntervalAttr *uint32
	// Connection Name
	NameAttr *string
	// Connection Description
	DescriptionAttr *string
	// Database Source Type
	TypeAttr *uint32
	// Reconnection Method
	ReconnectionMethodAttr *uint32
	// Last Refresh Version
	RefreshedVersionAttr uint8
	// Minimum Version Required for Refresh
	MinRefreshableVersionAttr *uint8
	// Save Password
	SavePasswordAttr *bool
	// New Connection
	NewAttr *bool
	// Deleted Connection
	DeletedAttr *bool
	// Only Use Connection File
	OnlyUseConnectionFileAttr *bool
	// Background Refresh
	BackgroundAttr *bool
	// Refresh on Open
	RefreshOnLoadAttr *bool
	// Save Data
	SaveDataAttr *bool
	// Reconnection Method
	CredentialsAttr ST_CredMethod
	// SSO Id
	SingleSignOnIdAttr *string
	// Database Properties
	DbPr *CT_DbPr
	// OLAP Properties
	OlapPr *CT_OlapPr
	// Web Query Properties
	WebPr *CT_WebPr
	// Text Import Settings
	TextPr *CT_TextPr
	// Query Parameters
	Parameters *CT_Parameters
	// Future Feature Data Storage
	ExtLst *CT_ExtensionList
}

func NewCT_Connection() *CT_Connection {
	ret := &CT_Connection{}
	return ret
}

func (m *CT_Connection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	if m.SourceFileAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sourceFile"},
			Value: fmt.Sprintf("%v", *m.SourceFileAttr)})
	}
	if m.OdcFileAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "odcFile"},
			Value: fmt.Sprintf("%v", *m.OdcFileAttr)})
	}
	if m.KeepAliveAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "keepAlive"},
			Value: fmt.Sprintf("%v", *m.KeepAliveAttr)})
	}
	if m.IntervalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "interval"},
			Value: fmt.Sprintf("%v", *m.IntervalAttr)})
	}
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.DescriptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "description"},
			Value: fmt.Sprintf("%v", *m.DescriptionAttr)})
	}
	if m.TypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"},
			Value: fmt.Sprintf("%v", *m.TypeAttr)})
	}
	if m.ReconnectionMethodAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "reconnectionMethod"},
			Value: fmt.Sprintf("%v", *m.ReconnectionMethodAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "refreshedVersion"},
		Value: fmt.Sprintf("%v", m.RefreshedVersionAttr)})
	if m.MinRefreshableVersionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minRefreshableVersion"},
			Value: fmt.Sprintf("%v", *m.MinRefreshableVersionAttr)})
	}
	if m.SavePasswordAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "savePassword"},
			Value: fmt.Sprintf("%v", *m.SavePasswordAttr)})
	}
	if m.NewAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "new"},
			Value: fmt.Sprintf("%v", *m.NewAttr)})
	}
	if m.DeletedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "deleted"},
			Value: fmt.Sprintf("%v", *m.DeletedAttr)})
	}
	if m.OnlyUseConnectionFileAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "onlyUseConnectionFile"},
			Value: fmt.Sprintf("%v", *m.OnlyUseConnectionFileAttr)})
	}
	if m.BackgroundAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "background"},
			Value: fmt.Sprintf("%v", *m.BackgroundAttr)})
	}
	if m.RefreshOnLoadAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "refreshOnLoad"},
			Value: fmt.Sprintf("%v", *m.RefreshOnLoadAttr)})
	}
	if m.SaveDataAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "saveData"},
			Value: fmt.Sprintf("%v", *m.SaveDataAttr)})
	}
	if m.CredentialsAttr != ST_CredMethodUnset {
		attr, err := m.CredentialsAttr.MarshalXMLAttr(xml.Name{Local: "credentials"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.SingleSignOnIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "singleSignOnId"},
			Value: fmt.Sprintf("%v", *m.SingleSignOnIdAttr)})
	}
	e.EncodeToken(start)
	if m.DbPr != nil {
		sedbPr := xml.StartElement{Name: xml.Name{Local: "x:dbPr"}}
		e.EncodeElement(m.DbPr, sedbPr)
	}
	if m.OlapPr != nil {
		seolapPr := xml.StartElement{Name: xml.Name{Local: "x:olapPr"}}
		e.EncodeElement(m.OlapPr, seolapPr)
	}
	if m.WebPr != nil {
		sewebPr := xml.StartElement{Name: xml.Name{Local: "x:webPr"}}
		e.EncodeElement(m.WebPr, sewebPr)
	}
	if m.TextPr != nil {
		setextPr := xml.StartElement{Name: xml.Name{Local: "x:textPr"}}
		e.EncodeElement(m.TextPr, setextPr)
	}
	if m.Parameters != nil {
		separameters := xml.StartElement{Name: xml.Name{Local: "x:parameters"}}
		e.EncodeElement(m.Parameters, separameters)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Connection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = uint32(parsed)
		}
		if attr.Name.Local == "sourceFile" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SourceFileAttr = &parsed
		}
		if attr.Name.Local == "odcFile" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OdcFileAttr = &parsed
		}
		if attr.Name.Local == "keepAlive" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.KeepAliveAttr = &parsed
		}
		if attr.Name.Local == "interval" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.IntervalAttr = &pt
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
		if attr.Name.Local == "description" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DescriptionAttr = &parsed
		}
		if attr.Name.Local == "type" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.TypeAttr = &pt
		}
		if attr.Name.Local == "reconnectionMethod" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ReconnectionMethodAttr = &pt
		}
		if attr.Name.Local == "refreshedVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			m.RefreshedVersionAttr = uint8(parsed)
		}
		if attr.Name.Local == "minRefreshableVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.MinRefreshableVersionAttr = &pt
		}
		if attr.Name.Local == "savePassword" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SavePasswordAttr = &parsed
		}
		if attr.Name.Local == "new" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NewAttr = &parsed
		}
		if attr.Name.Local == "deleted" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DeletedAttr = &parsed
		}
		if attr.Name.Local == "onlyUseConnectionFile" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OnlyUseConnectionFileAttr = &parsed
		}
		if attr.Name.Local == "background" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BackgroundAttr = &parsed
		}
		if attr.Name.Local == "refreshOnLoad" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RefreshOnLoadAttr = &parsed
		}
		if attr.Name.Local == "saveData" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SaveDataAttr = &parsed
		}
		if attr.Name.Local == "credentials" {
			m.CredentialsAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "singleSignOnId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SingleSignOnIdAttr = &parsed
		}
	}
lCT_Connection:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "dbPr":
				m.DbPr = NewCT_DbPr()
				if err := d.DecodeElement(m.DbPr, &el); err != nil {
					return err
				}
			case "olapPr":
				m.OlapPr = NewCT_OlapPr()
				if err := d.DecodeElement(m.OlapPr, &el); err != nil {
					return err
				}
			case "webPr":
				m.WebPr = NewCT_WebPr()
				if err := d.DecodeElement(m.WebPr, &el); err != nil {
					return err
				}
			case "textPr":
				m.TextPr = NewCT_TextPr()
				if err := d.DecodeElement(m.TextPr, &el); err != nil {
					return err
				}
			case "parameters":
				m.Parameters = NewCT_Parameters()
				if err := d.DecodeElement(m.Parameters, &el); err != nil {
					return err
				}
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
			break lCT_Connection
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Connection and its children
func (m *CT_Connection) Validate() error {
	return m.ValidateWithPath("CT_Connection")
}

// ValidateWithPath validates the CT_Connection and its children, prefixing error messages with path
func (m *CT_Connection) ValidateWithPath(path string) error {
	if err := m.CredentialsAttr.ValidateWithPath(path + "/CredentialsAttr"); err != nil {
		return err
	}
	if m.DbPr != nil {
		if err := m.DbPr.ValidateWithPath(path + "/DbPr"); err != nil {
			return err
		}
	}
	if m.OlapPr != nil {
		if err := m.OlapPr.ValidateWithPath(path + "/OlapPr"); err != nil {
			return err
		}
	}
	if m.WebPr != nil {
		if err := m.WebPr.ValidateWithPath(path + "/WebPr"); err != nil {
			return err
		}
	}
	if m.TextPr != nil {
		if err := m.TextPr.ValidateWithPath(path + "/TextPr"); err != nil {
			return err
		}
	}
	if m.Parameters != nil {
		if err := m.Parameters.ValidateWithPath(path + "/Parameters"); err != nil {
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
