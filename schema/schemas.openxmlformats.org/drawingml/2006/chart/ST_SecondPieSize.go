// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"fmt"
)

// ST_SecondPieSize is a union type
type ST_SecondPieSize struct {
	ST_SecondPieSizePercent *string
	ST_SecondPieSizeUShort  *uint16
}

func (m *ST_SecondPieSize) Validate() error {
	return m.ValidateWithPath("")
}

func (m *ST_SecondPieSize) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_SecondPieSizePercent != nil {
		mems = append(mems, "ST_SecondPieSizePercent")
	}
	if m.ST_SecondPieSizeUShort != nil {
		mems = append(mems, "ST_SecondPieSizeUShort")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_SecondPieSize) String() string {
	if m.ST_SecondPieSizePercent != nil {
		return fmt.Sprintf("%v", *m.ST_SecondPieSizePercent)
	}
	if m.ST_SecondPieSizeUShort != nil {
		return fmt.Sprintf("%v", *m.ST_SecondPieSizeUShort)
	}
	return ""
}
