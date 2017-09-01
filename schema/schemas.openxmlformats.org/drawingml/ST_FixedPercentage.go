// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"fmt"
)

// ST_FixedPercentage is a union type
type ST_FixedPercentage struct {
	ST_FixedPercentageDecimal *int32
	ST_FixedPercentage        *ST_Percentage
}

func (m *ST_FixedPercentage) Validate() error {
	return m.ValidateWithPath("")
}

func (m *ST_FixedPercentage) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_FixedPercentageDecimal != nil {
		mems = append(mems, "ST_FixedPercentageDecimal")
	}
	if m.ST_FixedPercentage != nil {
		if err := m.ST_FixedPercentage.ValidateWithPath(path + "/ST_FixedPercentage"); err != nil {
			return err
		}
		mems = append(mems, "ST_FixedPercentage")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_FixedPercentage) String() string {
	if m.ST_FixedPercentageDecimal != nil {
		return fmt.Sprintf("%v", *m.ST_FixedPercentageDecimal)
	}
	if m.ST_FixedPercentage != nil {
		return m.ST_FixedPercentage.String()
	}
	return ""
}
