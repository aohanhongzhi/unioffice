// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chartDrawing

import "baliance.com/gooxml"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_ShapeNonVisual", NewCT_ShapeNonVisual)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Shape", NewCT_Shape)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_ConnectorNonVisual", NewCT_ConnectorNonVisual)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Connector", NewCT_Connector)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_PictureNonVisual", NewCT_PictureNonVisual)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Picture", NewCT_Picture)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_GraphicFrameNonVisual", NewCT_GraphicFrameNonVisual)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_GraphicFrame", NewCT_GraphicFrame)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_GroupShapeNonVisual", NewCT_GroupShapeNonVisual)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_GroupShape", NewCT_GroupShape)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Marker", NewCT_Marker)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_RelSizeAnchor", NewCT_RelSizeAnchor)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_AbsSizeAnchor", NewCT_AbsSizeAnchor)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Drawing", NewCT_Drawing)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "EG_ObjectChoices", NewEG_ObjectChoices)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "EG_Anchor", NewEG_Anchor)
}
