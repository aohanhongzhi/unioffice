// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"fmt"
)

// ST_PrSetCustVal is a union type
type ST_PrSetCustVal struct {
	ST_Percentage *string
	Int32         *int32
}

func (m *ST_PrSetCustVal) Validate() error {
	return m.ValidateWithPath("")
}

func (m *ST_PrSetCustVal) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_Percentage != nil {
		mems = append(mems, "ST_Percentage")
	}
	if m.Int32 != nil {
		mems = append(mems, "Int32")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_PrSetCustVal) String() string {
	if m.ST_Percentage != nil {
		return fmt.Sprintf("%v", *m.ST_Percentage)
	}
	if m.Int32 != nil {
		return fmt.Sprintf("%v", *m.Int32)
	}
	return ""
}
