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

// ST_AdjAngle is a union type
type ST_AdjAngle struct {
	ST_Angle         *int32
	ST_GeomGuideName *string
}

func (m *ST_AdjAngle) Validate() error {
	return m.ValidateWithPath("")
}

func (m *ST_AdjAngle) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_Angle != nil {
		mems = append(mems, "ST_Angle")
	}
	if m.ST_GeomGuideName != nil {
		mems = append(mems, "ST_GeomGuideName")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_AdjAngle) String() string {
	if m.ST_Angle != nil {
		return fmt.Sprintf("%v", *m.ST_Angle)
	}
	if m.ST_GeomGuideName != nil {
		return fmt.Sprintf("%v", *m.ST_GeomGuideName)
	}
	return ""
}
