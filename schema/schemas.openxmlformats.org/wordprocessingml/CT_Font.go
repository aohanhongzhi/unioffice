// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_Font struct {
	// Primary Font Name
	NameAttr string
	// Alternate Names for Font
	AltName *CT_String
	// Panose-1 Typeface Classification Number
	Panose1 *CT_Panose
	// Character Set Supported By Font
	Charset *CT_Charset
	// Font Family
	Family *CT_FontFamily
	// Raster or Vector Font
	NotTrueType *CT_OnOff
	// Font Pitch
	Pitch *CT_Pitch
	// Supported Unicode Subranges and Code Pages
	Sig *CT_FontSig
	// Regular Font Style Embedding
	EmbedRegular *CT_FontRel
	// Bold Style Font Style Embedding
	EmbedBold *CT_FontRel
	// Italic Font Style Embedding
	EmbedItalic *CT_FontRel
	// Bold Italic Font Style Embedding
	EmbedBoldItalic *CT_FontRel
}

func NewCT_Font() *CT_Font {
	ret := &CT_Font{}
	return ret
}

func (m *CT_Font) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	e.EncodeToken(start)
	if m.AltName != nil {
		sealtName := xml.StartElement{Name: xml.Name{Local: "w:altName"}}
		e.EncodeElement(m.AltName, sealtName)
	}
	if m.Panose1 != nil {
		sepanose1 := xml.StartElement{Name: xml.Name{Local: "w:panose1"}}
		e.EncodeElement(m.Panose1, sepanose1)
	}
	if m.Charset != nil {
		secharset := xml.StartElement{Name: xml.Name{Local: "w:charset"}}
		e.EncodeElement(m.Charset, secharset)
	}
	if m.Family != nil {
		sefamily := xml.StartElement{Name: xml.Name{Local: "w:family"}}
		e.EncodeElement(m.Family, sefamily)
	}
	if m.NotTrueType != nil {
		senotTrueType := xml.StartElement{Name: xml.Name{Local: "w:notTrueType"}}
		e.EncodeElement(m.NotTrueType, senotTrueType)
	}
	if m.Pitch != nil {
		sepitch := xml.StartElement{Name: xml.Name{Local: "w:pitch"}}
		e.EncodeElement(m.Pitch, sepitch)
	}
	if m.Sig != nil {
		sesig := xml.StartElement{Name: xml.Name{Local: "w:sig"}}
		e.EncodeElement(m.Sig, sesig)
	}
	if m.EmbedRegular != nil {
		seembedRegular := xml.StartElement{Name: xml.Name{Local: "w:embedRegular"}}
		e.EncodeElement(m.EmbedRegular, seembedRegular)
	}
	if m.EmbedBold != nil {
		seembedBold := xml.StartElement{Name: xml.Name{Local: "w:embedBold"}}
		e.EncodeElement(m.EmbedBold, seembedBold)
	}
	if m.EmbedItalic != nil {
		seembedItalic := xml.StartElement{Name: xml.Name{Local: "w:embedItalic"}}
		e.EncodeElement(m.EmbedItalic, seembedItalic)
	}
	if m.EmbedBoldItalic != nil {
		seembedBoldItalic := xml.StartElement{Name: xml.Name{Local: "w:embedBoldItalic"}}
		e.EncodeElement(m.EmbedBoldItalic, seembedBoldItalic)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Font) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
	}
lCT_Font:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "altName":
				m.AltName = NewCT_String()
				if err := d.DecodeElement(m.AltName, &el); err != nil {
					return err
				}
			case "panose1":
				m.Panose1 = NewCT_Panose()
				if err := d.DecodeElement(m.Panose1, &el); err != nil {
					return err
				}
			case "charset":
				m.Charset = NewCT_Charset()
				if err := d.DecodeElement(m.Charset, &el); err != nil {
					return err
				}
			case "family":
				m.Family = NewCT_FontFamily()
				if err := d.DecodeElement(m.Family, &el); err != nil {
					return err
				}
			case "notTrueType":
				m.NotTrueType = NewCT_OnOff()
				if err := d.DecodeElement(m.NotTrueType, &el); err != nil {
					return err
				}
			case "pitch":
				m.Pitch = NewCT_Pitch()
				if err := d.DecodeElement(m.Pitch, &el); err != nil {
					return err
				}
			case "sig":
				m.Sig = NewCT_FontSig()
				if err := d.DecodeElement(m.Sig, &el); err != nil {
					return err
				}
			case "embedRegular":
				m.EmbedRegular = NewCT_FontRel()
				if err := d.DecodeElement(m.EmbedRegular, &el); err != nil {
					return err
				}
			case "embedBold":
				m.EmbedBold = NewCT_FontRel()
				if err := d.DecodeElement(m.EmbedBold, &el); err != nil {
					return err
				}
			case "embedItalic":
				m.EmbedItalic = NewCT_FontRel()
				if err := d.DecodeElement(m.EmbedItalic, &el); err != nil {
					return err
				}
			case "embedBoldItalic":
				m.EmbedBoldItalic = NewCT_FontRel()
				if err := d.DecodeElement(m.EmbedBoldItalic, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Font
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Font and its children
func (m *CT_Font) Validate() error {
	return m.ValidateWithPath("CT_Font")
}

// ValidateWithPath validates the CT_Font and its children, prefixing error messages with path
func (m *CT_Font) ValidateWithPath(path string) error {
	if m.AltName != nil {
		if err := m.AltName.ValidateWithPath(path + "/AltName"); err != nil {
			return err
		}
	}
	if m.Panose1 != nil {
		if err := m.Panose1.ValidateWithPath(path + "/Panose1"); err != nil {
			return err
		}
	}
	if m.Charset != nil {
		if err := m.Charset.ValidateWithPath(path + "/Charset"); err != nil {
			return err
		}
	}
	if m.Family != nil {
		if err := m.Family.ValidateWithPath(path + "/Family"); err != nil {
			return err
		}
	}
	if m.NotTrueType != nil {
		if err := m.NotTrueType.ValidateWithPath(path + "/NotTrueType"); err != nil {
			return err
		}
	}
	if m.Pitch != nil {
		if err := m.Pitch.ValidateWithPath(path + "/Pitch"); err != nil {
			return err
		}
	}
	if m.Sig != nil {
		if err := m.Sig.ValidateWithPath(path + "/Sig"); err != nil {
			return err
		}
	}
	if m.EmbedRegular != nil {
		if err := m.EmbedRegular.ValidateWithPath(path + "/EmbedRegular"); err != nil {
			return err
		}
	}
	if m.EmbedBold != nil {
		if err := m.EmbedBold.ValidateWithPath(path + "/EmbedBold"); err != nil {
			return err
		}
	}
	if m.EmbedItalic != nil {
		if err := m.EmbedItalic.ValidateWithPath(path + "/EmbedItalic"); err != nil {
			return err
		}
	}
	if m.EmbedBoldItalic != nil {
		if err := m.EmbedBoldItalic.ValidateWithPath(path + "/EmbedBoldItalic"); err != nil {
			return err
		}
	}
	return nil
}
