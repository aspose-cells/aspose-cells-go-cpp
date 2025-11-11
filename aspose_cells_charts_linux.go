// +build linux

/* ----------------------------------------------------------------
 * Copyright (c) 2001-2025 Aspose Pty Ltd. All Rights Reserved.
 * Powered by Aspose.Cells.
 * ---------------------------------------------------------------*/


package asposecells

// #cgo CXXFLAGS: -std=c++11
// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -Wl,-rpath,"${SRCDIR}/lib/linux_x86_64" -L"${SRCDIR}/lib/linux_x86_64" -lAspose.Cells.CWrapper
// #include <CellsFunctionMap.h>
import "C"
import (
	"fmt"  
 	
	"errors"	
	"runtime"
	"unsafe" 
)

/**************Enum AxisType *****************/

// Represents the axis type.
type AxisType int32

const(
// Category axis
AxisType_Category AxisType = 0 

// Value axis
AxisType_Value AxisType = 1 

// Series axis
AxisType_Series AxisType = 2 
)

func Int32ToAxisType(value int32)(AxisType ,error){
	switch value {
		case 0:  return AxisType_Category, nil  
		case 1:  return AxisType_Value, nil  
		case 2:  return AxisType_Series, nil  
		default:
			return 0 ,fmt.Errorf("invalid AxisType value: %d", value)
	}
}

/**************Enum BackgroundMode *****************/

// Represents the display mode of the background.
type BackgroundMode int32

const(
// Automatic
BackgroundMode_Automatic BackgroundMode = 0 

// Opaque
BackgroundMode_Opaque BackgroundMode = 1 

// Transparent
BackgroundMode_Transparent BackgroundMode = 2 
)

func Int32ToBackgroundMode(value int32)(BackgroundMode ,error){
	switch value {
		case 0:  return BackgroundMode_Automatic, nil  
		case 1:  return BackgroundMode_Opaque, nil  
		case 2:  return BackgroundMode_Transparent, nil  
		default:
			return 0 ,fmt.Errorf("invalid BackgroundMode value: %d", value)
	}
}

/**************Enum Bar3DShapeType *****************/

// Represents the shape used with the 3-D bar or column chart.
type Bar3DShapeType int32

const(
// Box
Bar3DShapeType_Box Bar3DShapeType = 0 

// PyramidToPoint
Bar3DShapeType_PyramidToPoint Bar3DShapeType = 1 

// PyramidToMax
Bar3DShapeType_PyramidToMax Bar3DShapeType = 2 

// Cylinder
Bar3DShapeType_Cylinder Bar3DShapeType = 3 

// ConeToPoint
Bar3DShapeType_ConeToPoint Bar3DShapeType = 4 

// ConeToMax
Bar3DShapeType_ConeToMax Bar3DShapeType = 5 
)

func Int32ToBar3DShapeType(value int32)(Bar3DShapeType ,error){
	switch value {
		case 0:  return Bar3DShapeType_Box, nil  
		case 1:  return Bar3DShapeType_PyramidToPoint, nil  
		case 2:  return Bar3DShapeType_PyramidToMax, nil  
		case 3:  return Bar3DShapeType_Cylinder, nil  
		case 4:  return Bar3DShapeType_ConeToPoint, nil  
		case 5:  return Bar3DShapeType_ConeToMax, nil  
		default:
			return 0 ,fmt.Errorf("invalid Bar3DShapeType value: %d", value)
	}
}

/**************Enum BubbleSizeRepresents *****************/

// Represents what the bubble size represents on a bubble chart.
type BubbleSizeRepresents int32

const(
// Represents the value of <see cref="Series.BubbleSizes"/> is area of the bubble.
BubbleSizeRepresents_SizeIsArea BubbleSizeRepresents = 0 

// Represents the value of <see cref="Series.BubbleSizes"/> is width of the bubble.
BubbleSizeRepresents_SizeIsWidth BubbleSizeRepresents = 1 
)

func Int32ToBubbleSizeRepresents(value int32)(BubbleSizeRepresents ,error){
	switch value {
		case 0:  return BubbleSizeRepresents_SizeIsArea, nil  
		case 1:  return BubbleSizeRepresents_SizeIsWidth, nil  
		default:
			return 0 ,fmt.Errorf("invalid BubbleSizeRepresents value: %d", value)
	}
}

/**************Enum CategoryType *****************/

// Represents the category axis type.
type CategoryType int32

const(
// AutomaticScale
CategoryType_AutomaticScale CategoryType = 0 

// CategoryScale
CategoryType_CategoryScale CategoryType = 1 

// TimeScale
CategoryType_TimeScale CategoryType = 2 
)

func Int32ToCategoryType(value int32)(CategoryType ,error){
	switch value {
		case 0:  return CategoryType_AutomaticScale, nil  
		case 1:  return CategoryType_CategoryScale, nil  
		case 2:  return CategoryType_TimeScale, nil  
		default:
			return 0 ,fmt.Errorf("invalid CategoryType value: %d", value)
	}
}

/**************Enum ChartColorPaletteType *****************/

// Enumerates all Monochromatic Palettes used in Excel chart.
type ChartColorPaletteType int32

const(
// accent1 theme color gradient, dark to light.
ChartColorPaletteType_MonochromaticPalette1 ChartColorPaletteType = 1 

// accent2 theme color gradient, dark to light.
ChartColorPaletteType_MonochromaticPalette2 ChartColorPaletteType = 2 

// accent3 theme color gradient, dark to light.
ChartColorPaletteType_MonochromaticPalette3 ChartColorPaletteType = 3 

// accent4 theme color gradient, dark to light.
ChartColorPaletteType_MonochromaticPalette4 ChartColorPaletteType = 4 

// accent5 theme color gradient, dark to light.
ChartColorPaletteType_MonochromaticPalette5 ChartColorPaletteType = 5 

// accent6 theme color gradient, dark to light.
ChartColorPaletteType_MonochromaticPalette6 ChartColorPaletteType = 6 

// accent7 theme color gradient.
ChartColorPaletteType_MonochromaticPalette7 ChartColorPaletteType = 7 

// accent1 theme color gradient, light to dark.
ChartColorPaletteType_MonochromaticPalette8 ChartColorPaletteType = 8 

// accent2 theme color gradient, light to dark.
ChartColorPaletteType_MonochromaticPalette9 ChartColorPaletteType = 9 

// accent3 theme color gradient, light to dark.
ChartColorPaletteType_MonochromaticPalette10 ChartColorPaletteType = 10 

// accent4 theme color gradient, light to dark.
ChartColorPaletteType_MonochromaticPalette11 ChartColorPaletteType = 11 

// accent5 theme color gradient, light to dark.
ChartColorPaletteType_MonochromaticPalette12 ChartColorPaletteType = 12 

// accent6 theme color gradient, light to dark.
ChartColorPaletteType_MonochromaticPalette13 ChartColorPaletteType = 13 
)

func Int32ToChartColorPaletteType(value int32)(ChartColorPaletteType ,error){
	switch value {
		case 1:  return ChartColorPaletteType_MonochromaticPalette1, nil  
		case 2:  return ChartColorPaletteType_MonochromaticPalette2, nil  
		case 3:  return ChartColorPaletteType_MonochromaticPalette3, nil  
		case 4:  return ChartColorPaletteType_MonochromaticPalette4, nil  
		case 5:  return ChartColorPaletteType_MonochromaticPalette5, nil  
		case 6:  return ChartColorPaletteType_MonochromaticPalette6, nil  
		case 7:  return ChartColorPaletteType_MonochromaticPalette7, nil  
		case 8:  return ChartColorPaletteType_MonochromaticPalette8, nil  
		case 9:  return ChartColorPaletteType_MonochromaticPalette9, nil  
		case 10:  return ChartColorPaletteType_MonochromaticPalette10, nil  
		case 11:  return ChartColorPaletteType_MonochromaticPalette11, nil  
		case 12:  return ChartColorPaletteType_MonochromaticPalette12, nil  
		case 13:  return ChartColorPaletteType_MonochromaticPalette13, nil  
		default:
			return 0 ,fmt.Errorf("invalid ChartColorPaletteType value: %d", value)
	}
}

/**************Enum ChartLineFormattingType *****************/

// Represents line format type of chart line.
type ChartLineFormattingType int32

const(
// Represents automatic formatting type.
ChartLineFormattingType_Automatic ChartLineFormattingType = 0 

// Represents solid formatting type.
ChartLineFormattingType_Solid ChartLineFormattingType = 1 

// Represents none formatting type.
ChartLineFormattingType_None ChartLineFormattingType = 2 

// Gradient
ChartLineFormattingType_Gradient ChartLineFormattingType = 3 
)

func Int32ToChartLineFormattingType(value int32)(ChartLineFormattingType ,error){
	switch value {
		case 0:  return ChartLineFormattingType_Automatic, nil  
		case 1:  return ChartLineFormattingType_Solid, nil  
		case 2:  return ChartLineFormattingType_None, nil  
		case 3:  return ChartLineFormattingType_Gradient, nil  
		default:
			return 0 ,fmt.Errorf("invalid ChartLineFormattingType value: %d", value)
	}
}

/**************Enum ChartMarkerType *****************/

// Represents the marker style in a line chart, scatter chart, or radar chart.
type ChartMarkerType int32

const(
// Automatic markers.
ChartMarkerType_Automatic ChartMarkerType = 0 

// Circular markers.
ChartMarkerType_Circle ChartMarkerType = 1 

// Long bar markers
ChartMarkerType_Dash ChartMarkerType = 2 

// Diamond-shaped markers.
ChartMarkerType_Diamond ChartMarkerType = 3 

// Short bar markers.
ChartMarkerType_Dot ChartMarkerType = 4 

// No markers.
ChartMarkerType_None ChartMarkerType = 5 

// Square markers with a plus sign.
ChartMarkerType_SquarePlus ChartMarkerType = 6 

// Square markers.
ChartMarkerType_Square ChartMarkerType = 7 

// Square markers with an asterisk.
ChartMarkerType_SquareStar ChartMarkerType = 8 

// Triangular markers.
ChartMarkerType_Triangle ChartMarkerType = 9 

// Square markers with an X.
ChartMarkerType_SquareX ChartMarkerType = 10 

// Picture
ChartMarkerType_Picture ChartMarkerType = 11 
)

func Int32ToChartMarkerType(value int32)(ChartMarkerType ,error){
	switch value {
		case 0:  return ChartMarkerType_Automatic, nil  
		case 1:  return ChartMarkerType_Circle, nil  
		case 2:  return ChartMarkerType_Dash, nil  
		case 3:  return ChartMarkerType_Diamond, nil  
		case 4:  return ChartMarkerType_Dot, nil  
		case 5:  return ChartMarkerType_None, nil  
		case 6:  return ChartMarkerType_SquarePlus, nil  
		case 7:  return ChartMarkerType_Square, nil  
		case 8:  return ChartMarkerType_SquareStar, nil  
		case 9:  return ChartMarkerType_Triangle, nil  
		case 10:  return ChartMarkerType_SquareX, nil  
		case 11:  return ChartMarkerType_Picture, nil  
		default:
			return 0 ,fmt.Errorf("invalid ChartMarkerType value: %d", value)
	}
}

/**************Enum ChartSplitType *****************/

// Represents the way the two sections of either a pie of pie chart or a bar of pie chart are split.
type ChartSplitType int32

const(
// Represents the data points shall be split between the pie
// and the second chart by putting the last Split Position
// of the data points in the second chart
ChartSplitType_Position ChartSplitType = 0 

// Represents the data points shall be split between the pie
// and the second chart by putting the data points with
// value less than Split Position in the second chart.
ChartSplitType_Value ChartSplitType = 1 

// Represents the data points shall be split between the pie
// and the second chart by putting the points with
// percentage less than Split Position percent in the
// second chart.
ChartSplitType_PercentValue ChartSplitType = 2 

// Represents the data points shall be split between the pie
// and the second chart according to the Custom Split
// values.
ChartSplitType_Custom ChartSplitType = 3 

// Represents the data points shall be split using the default
// mechanism for this chart type.
ChartSplitType_Auto ChartSplitType = 4 
)

func Int32ToChartSplitType(value int32)(ChartSplitType ,error){
	switch value {
		case 0:  return ChartSplitType_Position, nil  
		case 1:  return ChartSplitType_Value, nil  
		case 2:  return ChartSplitType_PercentValue, nil  
		case 3:  return ChartSplitType_Custom, nil  
		case 4:  return ChartSplitType_Auto, nil  
		default:
			return 0 ,fmt.Errorf("invalid ChartSplitType value: %d", value)
	}
}

/**************Enum ChartTextDirectionType *****************/

// Represents the text direction type of the chart.
type ChartTextDirectionType int32

const(
// Horizontal direction type.
ChartTextDirectionType_Horizontal ChartTextDirectionType = 0 

// Vertical direction type.
ChartTextDirectionType_Vertical ChartTextDirectionType = 1 

// Rotate 90 angle.
ChartTextDirectionType_Rotate90 ChartTextDirectionType = 2 

// Rotate 270 angle.
ChartTextDirectionType_Rotate270 ChartTextDirectionType = 3 

// Stacked text.
ChartTextDirectionType_Stacked ChartTextDirectionType = 4 
)

func Int32ToChartTextDirectionType(value int32)(ChartTextDirectionType ,error){
	switch value {
		case 0:  return ChartTextDirectionType_Horizontal, nil  
		case 1:  return ChartTextDirectionType_Vertical, nil  
		case 2:  return ChartTextDirectionType_Rotate90, nil  
		case 3:  return ChartTextDirectionType_Rotate270, nil  
		case 4:  return ChartTextDirectionType_Stacked, nil  
		default:
			return 0 ,fmt.Errorf("invalid ChartTextDirectionType value: %d", value)
	}
}

/**************Enum ChartType *****************/

// Enumerates all chart types used in Excel.
type ChartType int32

const(
// Represents Area Chart.
ChartType_Area ChartType = 0 

// Represents Stacked Area Chart.
ChartType_AreaStacked ChartType = 1 

// Represents 100% Stacked Area Chart.
ChartType_Area100PercentStacked ChartType = 2 

// Represents 3D Area Chart.
ChartType_Area3D ChartType = 3 

// Represents 3D Stacked Area Chart.
ChartType_Area3DStacked ChartType = 4 

// Represents 3D 100% Stacked Area Chart.
ChartType_Area3D100PercentStacked ChartType = 5 

// Represents Bar Chart: Clustered Bar Chart.
ChartType_Bar ChartType = 6 

// Represents Stacked Bar Chart.
ChartType_BarStacked ChartType = 7 

// Represents 100% Stacked Bar Chart.
ChartType_Bar100PercentStacked ChartType = 8 

// Represents 3D Clustered Bar Chart.
ChartType_Bar3DClustered ChartType = 9 

// Represents 3D Stacked Bar Chart.
ChartType_Bar3DStacked ChartType = 10 

// Represents 3D 100% Stacked Bar Chart.
ChartType_Bar3D100PercentStacked ChartType = 11 

// Represents Bubble Chart.
ChartType_Bubble ChartType = 12 

// Represents 3D Bubble Chart.
ChartType_Bubble3D ChartType = 13 

// Represents Column Chart: Clustered Column Chart.
ChartType_Column ChartType = 14 

// Represents Stacked Column Chart.
ChartType_ColumnStacked ChartType = 15 

// Represents 100% Stacked Column Chart.
ChartType_Column100PercentStacked ChartType = 16 

// Represents 3D Column Chart.
ChartType_Column3D ChartType = 17 

// Represents 3D Clustered Column Chart.
ChartType_Column3DClustered ChartType = 18 

// Represents 3D Stacked Column Chart.
ChartType_Column3DStacked ChartType = 19 

// Represents 3D 100% Stacked Column Chart.
ChartType_Column3D100PercentStacked ChartType = 20 

// Represents Cone Chart.
ChartType_Cone ChartType = 21 

// Represents Stacked Cone Chart.
ChartType_ConeStacked ChartType = 22 

// Represents 100% Stacked Cone Chart.
ChartType_Cone100PercentStacked ChartType = 23 

// Represents Conical Bar Chart.
ChartType_ConicalBar ChartType = 24 

// Represents Stacked Conical Bar Chart.
ChartType_ConicalBarStacked ChartType = 25 

// Represents 100% Stacked Conical Bar Chart.
ChartType_ConicalBar100PercentStacked ChartType = 26 

// Represents 3D Conical Column Chart.
ChartType_ConicalColumn3D ChartType = 27 

// Represents Cylinder Chart.
ChartType_Cylinder ChartType = 28 

// Represents Stacked Cylinder Chart.
ChartType_CylinderStacked ChartType = 29 

// Represents 100% Stacked Cylinder Chart.
ChartType_Cylinder100PercentStacked ChartType = 30 

// Represents Cylindrical Bar Chart.
ChartType_CylindricalBar ChartType = 31 

// Represents Stacked Cylindrical Bar Chart.
ChartType_CylindricalBarStacked ChartType = 32 

// Represents 100% Stacked Cylindrical Bar Chart.
ChartType_CylindricalBar100PercentStacked ChartType = 33 

// Represents 3D Cylindrical Column Chart.
ChartType_CylindricalColumn3D ChartType = 34 

// Represents Doughnut Chart.
ChartType_Doughnut ChartType = 35 

// Represents Exploded Doughnut Chart.
ChartType_DoughnutExploded ChartType = 36 

// Represents Line Chart.
ChartType_Line ChartType = 37 

// Represents Stacked Line Chart.
ChartType_LineStacked ChartType = 38 

// Represents 100% Stacked Line Chart.
ChartType_Line100PercentStacked ChartType = 39 

// Represents Line Chart with data markers.
ChartType_LineWithDataMarkers ChartType = 40 

// Represents Stacked Line Chart with data markers.
ChartType_LineStackedWithDataMarkers ChartType = 41 

// Represents 100% Stacked Line Chart with data markers.
ChartType_Line100PercentStackedWithDataMarkers ChartType = 42 

// Represents 3D Line Chart.
ChartType_Line3D ChartType = 43 

// Represents Pie Chart.
ChartType_Pie ChartType = 44 

// Represents 3D Pie Chart.
ChartType_Pie3D ChartType = 45 

// Represents Pie of Pie Chart.
ChartType_PiePie ChartType = 46 

// Represents Exploded Pie Chart.
ChartType_PieExploded ChartType = 47 

// Represents 3D Exploded Pie Chart.
ChartType_Pie3DExploded ChartType = 48 

// Represents Bar of Pie Chart.
ChartType_PieBar ChartType = 49 

// Represents Pyramid Chart.
ChartType_Pyramid ChartType = 50 

// Represents Stacked Pyramid Chart.
ChartType_PyramidStacked ChartType = 51 

// Represents 100% Stacked Pyramid Chart.
ChartType_Pyramid100PercentStacked ChartType = 52 

// Represents Pyramid Bar Chart.
ChartType_PyramidBar ChartType = 53 

// Represents Stacked Pyramid Bar Chart.
ChartType_PyramidBarStacked ChartType = 54 

// Represents 100% Stacked Pyramid Bar Chart.
ChartType_PyramidBar100PercentStacked ChartType = 55 

// Represents 3D Pyramid Column Chart.
ChartType_PyramidColumn3D ChartType = 56 

// Represents Radar Chart.
ChartType_Radar ChartType = 57 

// Represents Radar Chart with data markers.
ChartType_RadarWithDataMarkers ChartType = 58 

// Represents Filled Radar Chart.
ChartType_RadarFilled ChartType = 59 

// Represents Scatter Chart.
ChartType_Scatter ChartType = 60 

// Represents Scatter Chart connected by curves, with data markers.
ChartType_ScatterConnectedByCurvesWithDataMarker ChartType = 61 

// Represents Scatter Chart connected by curves, without data markers.
ChartType_ScatterConnectedByCurvesWithoutDataMarker ChartType = 62 

// Represents Scatter Chart connected by lines, with data markers.
ChartType_ScatterConnectedByLinesWithDataMarker ChartType = 63 

// Represents Scatter Chart connected by lines, without data markers.
ChartType_ScatterConnectedByLinesWithoutDataMarker ChartType = 64 

// Represents High-Low-Close Stock Chart.
ChartType_StockHighLowClose ChartType = 65 

// Represents Open-High-Low-Close Stock Chart.
ChartType_StockOpenHighLowClose ChartType = 66 

// Represents Volume-High-Low-Close Stock Chart.
ChartType_StockVolumeHighLowClose ChartType = 67 

// Represents Volume-Open-High-Low-Close Stock Chart.
ChartType_StockVolumeOpenHighLowClose ChartType = 68 

// Represents Surface Chart: 3D Surface Chart.
ChartType_Surface3D ChartType = 69 

// Represents Wireframe 3D Surface Chart.
ChartType_SurfaceWireframe3D ChartType = 70 

// Represents Contour Chart.
ChartType_SurfaceContour ChartType = 71 

// Represents Wireframe Contour Chart.
ChartType_SurfaceContourWireframe ChartType = 72 

// The series is laid out as box and whisker.
ChartType_BoxWhisker ChartType = 73 

// The series is laid out as a funnel.
ChartType_Funnel ChartType = 74 

// The series is laid out as pareto lines.
ChartType_ParetoLine ChartType = 75 

// The series is laid out as a sunburst.
ChartType_Sunburst ChartType = 76 

// The series is laid out as a treemap.
ChartType_Treemap ChartType = 77 

// The series is laid out as a waterfall.
ChartType_Waterfall ChartType = 78 

// The series is laid out as a histogram.
ChartType_Histogram ChartType = 79 

// The series is laid out as a region map.
ChartType_Map ChartType = 80 

// The series is laid out as a radial historgram. It is used only for rendering
ChartType_RadialHistogram ChartType = 81 
)

func Int32ToChartType(value int32)(ChartType ,error){
	switch value {
		case 0:  return ChartType_Area, nil  
		case 1:  return ChartType_AreaStacked, nil  
		case 2:  return ChartType_Area100PercentStacked, nil  
		case 3:  return ChartType_Area3D, nil  
		case 4:  return ChartType_Area3DStacked, nil  
		case 5:  return ChartType_Area3D100PercentStacked, nil  
		case 6:  return ChartType_Bar, nil  
		case 7:  return ChartType_BarStacked, nil  
		case 8:  return ChartType_Bar100PercentStacked, nil  
		case 9:  return ChartType_Bar3DClustered, nil  
		case 10:  return ChartType_Bar3DStacked, nil  
		case 11:  return ChartType_Bar3D100PercentStacked, nil  
		case 12:  return ChartType_Bubble, nil  
		case 13:  return ChartType_Bubble3D, nil  
		case 14:  return ChartType_Column, nil  
		case 15:  return ChartType_ColumnStacked, nil  
		case 16:  return ChartType_Column100PercentStacked, nil  
		case 17:  return ChartType_Column3D, nil  
		case 18:  return ChartType_Column3DClustered, nil  
		case 19:  return ChartType_Column3DStacked, nil  
		case 20:  return ChartType_Column3D100PercentStacked, nil  
		case 21:  return ChartType_Cone, nil  
		case 22:  return ChartType_ConeStacked, nil  
		case 23:  return ChartType_Cone100PercentStacked, nil  
		case 24:  return ChartType_ConicalBar, nil  
		case 25:  return ChartType_ConicalBarStacked, nil  
		case 26:  return ChartType_ConicalBar100PercentStacked, nil  
		case 27:  return ChartType_ConicalColumn3D, nil  
		case 28:  return ChartType_Cylinder, nil  
		case 29:  return ChartType_CylinderStacked, nil  
		case 30:  return ChartType_Cylinder100PercentStacked, nil  
		case 31:  return ChartType_CylindricalBar, nil  
		case 32:  return ChartType_CylindricalBarStacked, nil  
		case 33:  return ChartType_CylindricalBar100PercentStacked, nil  
		case 34:  return ChartType_CylindricalColumn3D, nil  
		case 35:  return ChartType_Doughnut, nil  
		case 36:  return ChartType_DoughnutExploded, nil  
		case 37:  return ChartType_Line, nil  
		case 38:  return ChartType_LineStacked, nil  
		case 39:  return ChartType_Line100PercentStacked, nil  
		case 40:  return ChartType_LineWithDataMarkers, nil  
		case 41:  return ChartType_LineStackedWithDataMarkers, nil  
		case 42:  return ChartType_Line100PercentStackedWithDataMarkers, nil  
		case 43:  return ChartType_Line3D, nil  
		case 44:  return ChartType_Pie, nil  
		case 45:  return ChartType_Pie3D, nil  
		case 46:  return ChartType_PiePie, nil  
		case 47:  return ChartType_PieExploded, nil  
		case 48:  return ChartType_Pie3DExploded, nil  
		case 49:  return ChartType_PieBar, nil  
		case 50:  return ChartType_Pyramid, nil  
		case 51:  return ChartType_PyramidStacked, nil  
		case 52:  return ChartType_Pyramid100PercentStacked, nil  
		case 53:  return ChartType_PyramidBar, nil  
		case 54:  return ChartType_PyramidBarStacked, nil  
		case 55:  return ChartType_PyramidBar100PercentStacked, nil  
		case 56:  return ChartType_PyramidColumn3D, nil  
		case 57:  return ChartType_Radar, nil  
		case 58:  return ChartType_RadarWithDataMarkers, nil  
		case 59:  return ChartType_RadarFilled, nil  
		case 60:  return ChartType_Scatter, nil  
		case 61:  return ChartType_ScatterConnectedByCurvesWithDataMarker, nil  
		case 62:  return ChartType_ScatterConnectedByCurvesWithoutDataMarker, nil  
		case 63:  return ChartType_ScatterConnectedByLinesWithDataMarker, nil  
		case 64:  return ChartType_ScatterConnectedByLinesWithoutDataMarker, nil  
		case 65:  return ChartType_StockHighLowClose, nil  
		case 66:  return ChartType_StockOpenHighLowClose, nil  
		case 67:  return ChartType_StockVolumeHighLowClose, nil  
		case 68:  return ChartType_StockVolumeOpenHighLowClose, nil  
		case 69:  return ChartType_Surface3D, nil  
		case 70:  return ChartType_SurfaceWireframe3D, nil  
		case 71:  return ChartType_SurfaceContour, nil  
		case 72:  return ChartType_SurfaceContourWireframe, nil  
		case 73:  return ChartType_BoxWhisker, nil  
		case 74:  return ChartType_Funnel, nil  
		case 75:  return ChartType_ParetoLine, nil  
		case 76:  return ChartType_Sunburst, nil  
		case 77:  return ChartType_Treemap, nil  
		case 78:  return ChartType_Waterfall, nil  
		case 79:  return ChartType_Histogram, nil  
		case 80:  return ChartType_Map, nil  
		case 81:  return ChartType_RadialHistogram, nil  
		default:
			return 0 ,fmt.Errorf("invalid ChartType value: %d", value)
	}
}

/**************Enum CrossType *****************/

// Represents the axis cross type.
type CrossType int32

const(
// Microsoft Excel sets the axis crossing point.
CrossType_Automatic CrossType = 0 

// The axis crosses at the maximum value.
CrossType_Maximum CrossType = 1 

// The axis crosses at the minimum value.
CrossType_Minimum CrossType = 2 

// The axis crosses at the custom value.
CrossType_Custom CrossType = 3 
)

func Int32ToCrossType(value int32)(CrossType ,error){
	switch value {
		case 0:  return CrossType_Automatic, nil  
		case 1:  return CrossType_Maximum, nil  
		case 2:  return CrossType_Minimum, nil  
		case 3:  return CrossType_Custom, nil  
		default:
			return 0 ,fmt.Errorf("invalid CrossType value: %d", value)
	}
}

/**************Enum DataLabelsSeparatorType *****************/

// Represents the separator type of DataLabels.
type DataLabelsSeparatorType int32

const(
// Represents automatic separator
DataLabelsSeparatorType_Auto DataLabelsSeparatorType = 0 

// Represents space(" ")
DataLabelsSeparatorType_Space DataLabelsSeparatorType = 1 

// Represents comma(",")
DataLabelsSeparatorType_Comma DataLabelsSeparatorType = 2 

// Represents semicolon(";")
DataLabelsSeparatorType_Semicolon DataLabelsSeparatorType = 3 

// Represents period(".")
DataLabelsSeparatorType_Period DataLabelsSeparatorType = 4 

// Represents newline("\n")
DataLabelsSeparatorType_NewLine DataLabelsSeparatorType = 5 

// Represents custom separator
DataLabelsSeparatorType_Custom DataLabelsSeparatorType = 6 
)

func Int32ToDataLabelsSeparatorType(value int32)(DataLabelsSeparatorType ,error){
	switch value {
		case 0:  return DataLabelsSeparatorType_Auto, nil  
		case 1:  return DataLabelsSeparatorType_Space, nil  
		case 2:  return DataLabelsSeparatorType_Comma, nil  
		case 3:  return DataLabelsSeparatorType_Semicolon, nil  
		case 4:  return DataLabelsSeparatorType_Period, nil  
		case 5:  return DataLabelsSeparatorType_NewLine, nil  
		case 6:  return DataLabelsSeparatorType_Custom, nil  
		default:
			return 0 ,fmt.Errorf("invalid DataLabelsSeparatorType value: %d", value)
	}
}

/**************Enum DisplayUnitType *****************/

// Represents the type of display unit of chart's axis.
type DisplayUnitType int32

const(
// Display unit is None.
DisplayUnitType_None DisplayUnitType = 0 

// Specifies the values on the chart shall be divided by 100.
DisplayUnitType_Hundreds DisplayUnitType = 1 

// Specifies the values on the chart shall be divided by 1,000.
DisplayUnitType_Thousands DisplayUnitType = 2 

// Specifies the values on the chart shall be divided by 10,000.
DisplayUnitType_TenThousands DisplayUnitType = 3 

// Specifies the values on the chart shall be divided by 100,000.
DisplayUnitType_HundredThousands DisplayUnitType = 4 

// Specifies the values on the chart shall be divided by 1,000,000.
DisplayUnitType_Millions DisplayUnitType = 5 

// Specifies the values on the chart shall be divided by 10,000,000.
DisplayUnitType_TenMillions DisplayUnitType = 6 

// Specifies the values on the chart shall be divided by 100,000,000.
DisplayUnitType_HundredMillions DisplayUnitType = 7 

// Specifies the values on the chart shall be divided by 1,000,000,000.
DisplayUnitType_Billions DisplayUnitType = 8 

// Specifies the values on the chart shall be divided by 1,000,000,000,000.
DisplayUnitType_Trillions DisplayUnitType = 9 

// The values on the chart shall be divided by 0.01.
DisplayUnitType_Percentage DisplayUnitType = 10 

// specifies a custom value for the display unit.
DisplayUnitType_Cust DisplayUnitType = 11 

// specifies a custom value for the display unit.
DisplayUnitType_Custom DisplayUnitType = 12 
)

func Int32ToDisplayUnitType(value int32)(DisplayUnitType ,error){
	switch value {
		case 0:  return DisplayUnitType_None, nil  
		case 1:  return DisplayUnitType_Hundreds, nil  
		case 2:  return DisplayUnitType_Thousands, nil  
		case 3:  return DisplayUnitType_TenThousands, nil  
		case 4:  return DisplayUnitType_HundredThousands, nil  
		case 5:  return DisplayUnitType_Millions, nil  
		case 6:  return DisplayUnitType_TenMillions, nil  
		case 7:  return DisplayUnitType_HundredMillions, nil  
		case 8:  return DisplayUnitType_Billions, nil  
		case 9:  return DisplayUnitType_Trillions, nil  
		case 10:  return DisplayUnitType_Percentage, nil  
		case 11:  return DisplayUnitType_Cust, nil  
		case 12:  return DisplayUnitType_Custom, nil  
		default:
			return 0 ,fmt.Errorf("invalid DisplayUnitType value: %d", value)
	}
}

/**************Enum ErrorBarDisplayType *****************/

// Represents error bar display type.
type ErrorBarDisplayType int32

const(
// Both
ErrorBarDisplayType_Both ErrorBarDisplayType = 0 

// Minus
ErrorBarDisplayType_Minus ErrorBarDisplayType = 1 

// None
ErrorBarDisplayType_None ErrorBarDisplayType = 2 

// Plus
ErrorBarDisplayType_Plus ErrorBarDisplayType = 3 
)

func Int32ToErrorBarDisplayType(value int32)(ErrorBarDisplayType ,error){
	switch value {
		case 0:  return ErrorBarDisplayType_Both, nil  
		case 1:  return ErrorBarDisplayType_Minus, nil  
		case 2:  return ErrorBarDisplayType_None, nil  
		case 3:  return ErrorBarDisplayType_Plus, nil  
		default:
			return 0 ,fmt.Errorf("invalid ErrorBarDisplayType value: %d", value)
	}
}

/**************Enum ErrorBarType *****************/

// Represents error bar amount type.
type ErrorBarType int32

const(
// InnerCustom value type.
ErrorBarType_Custom ErrorBarType = 0 

// Fixed value type.
ErrorBarType_FixedValue ErrorBarType = 1 

// Percentage type
ErrorBarType_Percent ErrorBarType = 2 

// Standard deviation type.
ErrorBarType_StDev ErrorBarType = 3 

// Standard error type.
ErrorBarType_StError ErrorBarType = 4 
)

func Int32ToErrorBarType(value int32)(ErrorBarType ,error){
	switch value {
		case 0:  return ErrorBarType_Custom, nil  
		case 1:  return ErrorBarType_FixedValue, nil  
		case 2:  return ErrorBarType_Percent, nil  
		case 3:  return ErrorBarType_StDev, nil  
		case 4:  return ErrorBarType_StError, nil  
		default:
			return 0 ,fmt.Errorf("invalid ErrorBarType value: %d", value)
	}
}

/**************Enum FormattingType *****************/

// Represents the type of formatting applied to an <see cref="Area"/> object or a <see cref="Line"/> object.
type FormattingType int32

const(
// Represents automatic formatting type.
FormattingType_Automatic FormattingType = 0 

// Represents custom formatting type.
FormattingType_Custom FormattingType = 2 

// Represents none formatting type.
FormattingType_None FormattingType = 1 
)

func Int32ToFormattingType(value int32)(FormattingType ,error){
	switch value {
		case 0:  return FormattingType_Automatic, nil  
		case 2:  return FormattingType_Custom, nil  
		case 1:  return FormattingType_None, nil  
		default:
			return 0 ,fmt.Errorf("invalid FormattingType value: %d", value)
	}
}

/**************Enum LabelPositionType *****************/

// Represents data label position type.
type LabelPositionType int32

const(
// Applies only to bar, 2d/3d pie charts
LabelPositionType_Center LabelPositionType = 0 

// Applies only to bar, 2d/3d pie charts
LabelPositionType_InsideBase LabelPositionType = 1 

// Applies only to bar charts
LabelPositionType_InsideEnd LabelPositionType = 2 

// Applies only to bar, 2d/3d pie charts
LabelPositionType_OutsideEnd LabelPositionType = 3 

// Applies only to line charts
LabelPositionType_Above LabelPositionType = 4 

// Applies only to line charts
LabelPositionType_Below LabelPositionType = 5 

// Applies only to line charts
LabelPositionType_Left LabelPositionType = 6 

// Applies only to line charts
LabelPositionType_Right LabelPositionType = 7 

// Applies only to 2d/3d pie charts
LabelPositionType_BestFit LabelPositionType = 8 

// User moved the data labels, Only for reading chart from template file.
LabelPositionType_Moved LabelPositionType = 9 
)

func Int32ToLabelPositionType(value int32)(LabelPositionType ,error){
	switch value {
		case 0:  return LabelPositionType_Center, nil  
		case 1:  return LabelPositionType_InsideBase, nil  
		case 2:  return LabelPositionType_InsideEnd, nil  
		case 3:  return LabelPositionType_OutsideEnd, nil  
		case 4:  return LabelPositionType_Above, nil  
		case 5:  return LabelPositionType_Below, nil  
		case 6:  return LabelPositionType_Left, nil  
		case 7:  return LabelPositionType_Right, nil  
		case 8:  return LabelPositionType_BestFit, nil  
		case 9:  return LabelPositionType_Moved, nil  
		default:
			return 0 ,fmt.Errorf("invalid LabelPositionType value: %d", value)
	}
}

/**************Enum LegendPositionType *****************/

// Enumerates the legend position types.
type LegendPositionType int32

const(
// Displays the legend to the bottom of the chart's plot area.
LegendPositionType_Bottom LegendPositionType = 0 

// Displays the legend to the corner of the chart's plot area.
LegendPositionType_Corner LegendPositionType = 1 

// Displays the legend to the left of the chart's plot area.
LegendPositionType_Left LegendPositionType = 4 

// Represents that the legend is not docked.
LegendPositionType_NotDocked LegendPositionType = 7 

// Displays the legend to the right of the chart's plot area.
LegendPositionType_Right LegendPositionType = 3 

// Displays the legend to the top of the chart's plot area.
LegendPositionType_Top LegendPositionType = 2 
)

func Int32ToLegendPositionType(value int32)(LegendPositionType ,error){
	switch value {
		case 0:  return LegendPositionType_Bottom, nil  
		case 1:  return LegendPositionType_Corner, nil  
		case 4:  return LegendPositionType_Left, nil  
		case 7:  return LegendPositionType_NotDocked, nil  
		case 3:  return LegendPositionType_Right, nil  
		case 2:  return LegendPositionType_Top, nil  
		default:
			return 0 ,fmt.Errorf("invalid LegendPositionType value: %d", value)
	}
}

/**************Enum MapChartLabelLayout *****************/

// Represents the layout of map chart's labels.
type MapChartLabelLayout int32

const(
// Only best fit.
MapChartLabelLayout_BestFitOnly MapChartLabelLayout = 0 

// Shows all labels.
MapChartLabelLayout_ShowAll MapChartLabelLayout = 1 

// No labels.
MapChartLabelLayout_None MapChartLabelLayout = 2 
)

func Int32ToMapChartLabelLayout(value int32)(MapChartLabelLayout ,error){
	switch value {
		case 0:  return MapChartLabelLayout_BestFitOnly, nil  
		case 1:  return MapChartLabelLayout_ShowAll, nil  
		case 2:  return MapChartLabelLayout_None, nil  
		default:
			return 0 ,fmt.Errorf("invalid MapChartLabelLayout value: %d", value)
	}
}

/**************Enum MapChartProjectionType *****************/

// Represents projection type of the map chart.
type MapChartProjectionType int32

const(
// Automatic
MapChartProjectionType_Automatic MapChartProjectionType = 0 

// Mercator
MapChartProjectionType_Mercator MapChartProjectionType = 1 

// Miller
MapChartProjectionType_Miller MapChartProjectionType = 2 

// Albers
MapChartProjectionType_Albers MapChartProjectionType = 3 
)

func Int32ToMapChartProjectionType(value int32)(MapChartProjectionType ,error){
	switch value {
		case 0:  return MapChartProjectionType_Automatic, nil  
		case 1:  return MapChartProjectionType_Mercator, nil  
		case 2:  return MapChartProjectionType_Miller, nil  
		case 3:  return MapChartProjectionType_Albers, nil  
		default:
			return 0 ,fmt.Errorf("invalid MapChartProjectionType value: %d", value)
	}
}

/**************Enum MapChartRegionType *****************/

// Represents the region type of the map chart.
type MapChartRegionType int32

const(
// Automatic
MapChartRegionType_Automatic MapChartRegionType = 0 

// Only Data.
MapChartRegionType_DataOnly MapChartRegionType = 1 

// Country region list.
MapChartRegionType_CountryRegionList MapChartRegionType = 2 

// World.
MapChartRegionType_World MapChartRegionType = 3 
)

func Int32ToMapChartRegionType(value int32)(MapChartRegionType ,error){
	switch value {
		case 0:  return MapChartRegionType_Automatic, nil  
		case 1:  return MapChartRegionType_DataOnly, nil  
		case 2:  return MapChartRegionType_CountryRegionList, nil  
		case 3:  return MapChartRegionType_World, nil  
		default:
			return 0 ,fmt.Errorf("invalid MapChartRegionType value: %d", value)
	}
}

/**************Enum PlotDataByType *****************/

// Represents the type of data plot by row or column.
type PlotDataByType int32

const(
// By row.
PlotDataByType_Row PlotDataByType = 0 

// By column.
PlotDataByType_Column PlotDataByType = 1 
)

func Int32ToPlotDataByType(value int32)(PlotDataByType ,error){
	switch value {
		case 0:  return PlotDataByType_Row, nil  
		case 1:  return PlotDataByType_Column, nil  
		default:
			return 0 ,fmt.Errorf("invalid PlotDataByType value: %d", value)
	}
}

/**************Enum PlotEmptyCellsType *****************/

// Represents all plot empty cells type of a chart.
type PlotEmptyCellsType int32

const(
// Not plotted(leave gap)
PlotEmptyCellsType_NotPlotted PlotEmptyCellsType = 0 

// Zero
PlotEmptyCellsType_Zero PlotEmptyCellsType = 1 

// Interpolated
PlotEmptyCellsType_Interpolated PlotEmptyCellsType = 2 
)

func Int32ToPlotEmptyCellsType(value int32)(PlotEmptyCellsType ,error){
	switch value {
		case 0:  return PlotEmptyCellsType_NotPlotted, nil  
		case 1:  return PlotEmptyCellsType_Zero, nil  
		case 2:  return PlotEmptyCellsType_Interpolated, nil  
		default:
			return 0 ,fmt.Errorf("invalid PlotEmptyCellsType value: %d", value)
	}
}

/**************Enum QuartileCalculationType *****************/

// Represents quartile calculation methods.
type QuartileCalculationType int32

const(
// The quartile calculation includes the median when splitting the dataset into quartiles.
QuartileCalculationType_Exclusive QuartileCalculationType = 0 

// The quartile calculation excludes the median when splitting the dataset into quartiles.
QuartileCalculationType_Inclusive QuartileCalculationType = 1 
)

func Int32ToQuartileCalculationType(value int32)(QuartileCalculationType ,error){
	switch value {
		case 0:  return QuartileCalculationType_Exclusive, nil  
		case 1:  return QuartileCalculationType_Inclusive, nil  
		default:
			return 0 ,fmt.Errorf("invalid QuartileCalculationType value: %d", value)
	}
}

/**************Enum SparklineAxisMinMaxType *****************/

// Represents the minimum and maximum value types for the sparkline vertical axis.
type SparklineAxisMinMaxType int32

const(
// Automatic for each sparkline.
SparklineAxisMinMaxType_AutoIndividual SparklineAxisMinMaxType = 0 

// Same for all sparklines in the group.
SparklineAxisMinMaxType_Group SparklineAxisMinMaxType = 1 

// Custom value for sparkline.
SparklineAxisMinMaxType_Custom SparklineAxisMinMaxType = 2 
)

func Int32ToSparklineAxisMinMaxType(value int32)(SparklineAxisMinMaxType ,error){
	switch value {
		case 0:  return SparklineAxisMinMaxType_AutoIndividual, nil  
		case 1:  return SparklineAxisMinMaxType_Group, nil  
		case 2:  return SparklineAxisMinMaxType_Custom, nil  
		default:
			return 0 ,fmt.Errorf("invalid SparklineAxisMinMaxType value: %d", value)
	}
}

/**************Enum SparklinePresetStyleType *****************/

// Represents the preset style types for sparkline.
type SparklinePresetStyleType int32

const(
// Style 1
SparklinePresetStyleType_Style1 SparklinePresetStyleType = 0 

// Style 2
SparklinePresetStyleType_Style2 SparklinePresetStyleType = 1 

// Style 3
SparklinePresetStyleType_Style3 SparklinePresetStyleType = 2 

// Style 4
SparklinePresetStyleType_Style4 SparklinePresetStyleType = 3 

// Style 5
SparklinePresetStyleType_Style5 SparklinePresetStyleType = 4 

// Style 6
SparklinePresetStyleType_Style6 SparklinePresetStyleType = 5 

// Style 7
SparklinePresetStyleType_Style7 SparklinePresetStyleType = 6 

// Style 8
SparklinePresetStyleType_Style8 SparklinePresetStyleType = 7 

// Style 9
SparklinePresetStyleType_Style9 SparklinePresetStyleType = 8 

// Style 10
SparklinePresetStyleType_Style10 SparklinePresetStyleType = 9 

// Style 11
SparklinePresetStyleType_Style11 SparklinePresetStyleType = 10 

// Style 12
SparklinePresetStyleType_Style12 SparklinePresetStyleType = 11 

// Style 13
SparklinePresetStyleType_Style13 SparklinePresetStyleType = 12 

// Style 14
SparklinePresetStyleType_Style14 SparklinePresetStyleType = 13 

// Style 15
SparklinePresetStyleType_Style15 SparklinePresetStyleType = 14 

// Style 16
SparklinePresetStyleType_Style16 SparklinePresetStyleType = 15 

// Style 17
SparklinePresetStyleType_Style17 SparklinePresetStyleType = 16 

// Style 18
SparklinePresetStyleType_Style18 SparklinePresetStyleType = 17 

// Style 19
SparklinePresetStyleType_Style19 SparklinePresetStyleType = 18 

// Style 20
SparklinePresetStyleType_Style20 SparklinePresetStyleType = 19 

// Style 21
SparklinePresetStyleType_Style21 SparklinePresetStyleType = 20 

// Style 22
SparklinePresetStyleType_Style22 SparklinePresetStyleType = 21 

// Style 23
SparklinePresetStyleType_Style23 SparklinePresetStyleType = 22 

// Style 24
SparklinePresetStyleType_Style24 SparklinePresetStyleType = 23 

// Style 25
SparklinePresetStyleType_Style25 SparklinePresetStyleType = 24 

// Style 26
SparklinePresetStyleType_Style26 SparklinePresetStyleType = 25 

// Style 27
SparklinePresetStyleType_Style27 SparklinePresetStyleType = 26 

// Style 28
SparklinePresetStyleType_Style28 SparklinePresetStyleType = 27 

// Style 29
SparklinePresetStyleType_Style29 SparklinePresetStyleType = 28 

// Style 30
SparklinePresetStyleType_Style30 SparklinePresetStyleType = 29 

// Style 31
SparklinePresetStyleType_Style31 SparklinePresetStyleType = 30 

// Style 32
SparklinePresetStyleType_Style32 SparklinePresetStyleType = 31 

// Style 33
SparklinePresetStyleType_Style33 SparklinePresetStyleType = 32 

// Style 34
SparklinePresetStyleType_Style34 SparklinePresetStyleType = 33 

// Style 35
SparklinePresetStyleType_Style35 SparklinePresetStyleType = 34 

// Style 36
SparklinePresetStyleType_Style36 SparklinePresetStyleType = 35 

// No preset style.
SparklinePresetStyleType_Custom SparklinePresetStyleType = 36 
)

func Int32ToSparklinePresetStyleType(value int32)(SparklinePresetStyleType ,error){
	switch value {
		case 0:  return SparklinePresetStyleType_Style1, nil  
		case 1:  return SparklinePresetStyleType_Style2, nil  
		case 2:  return SparklinePresetStyleType_Style3, nil  
		case 3:  return SparklinePresetStyleType_Style4, nil  
		case 4:  return SparklinePresetStyleType_Style5, nil  
		case 5:  return SparklinePresetStyleType_Style6, nil  
		case 6:  return SparklinePresetStyleType_Style7, nil  
		case 7:  return SparklinePresetStyleType_Style8, nil  
		case 8:  return SparklinePresetStyleType_Style9, nil  
		case 9:  return SparklinePresetStyleType_Style10, nil  
		case 10:  return SparklinePresetStyleType_Style11, nil  
		case 11:  return SparklinePresetStyleType_Style12, nil  
		case 12:  return SparklinePresetStyleType_Style13, nil  
		case 13:  return SparklinePresetStyleType_Style14, nil  
		case 14:  return SparklinePresetStyleType_Style15, nil  
		case 15:  return SparklinePresetStyleType_Style16, nil  
		case 16:  return SparklinePresetStyleType_Style17, nil  
		case 17:  return SparklinePresetStyleType_Style18, nil  
		case 18:  return SparklinePresetStyleType_Style19, nil  
		case 19:  return SparklinePresetStyleType_Style20, nil  
		case 20:  return SparklinePresetStyleType_Style21, nil  
		case 21:  return SparklinePresetStyleType_Style22, nil  
		case 22:  return SparklinePresetStyleType_Style23, nil  
		case 23:  return SparklinePresetStyleType_Style24, nil  
		case 24:  return SparklinePresetStyleType_Style25, nil  
		case 25:  return SparklinePresetStyleType_Style26, nil  
		case 26:  return SparklinePresetStyleType_Style27, nil  
		case 27:  return SparklinePresetStyleType_Style28, nil  
		case 28:  return SparklinePresetStyleType_Style29, nil  
		case 29:  return SparklinePresetStyleType_Style30, nil  
		case 30:  return SparklinePresetStyleType_Style31, nil  
		case 31:  return SparklinePresetStyleType_Style32, nil  
		case 32:  return SparklinePresetStyleType_Style33, nil  
		case 33:  return SparklinePresetStyleType_Style34, nil  
		case 34:  return SparklinePresetStyleType_Style35, nil  
		case 35:  return SparklinePresetStyleType_Style36, nil  
		case 36:  return SparklinePresetStyleType_Custom, nil  
		default:
			return 0 ,fmt.Errorf("invalid SparklinePresetStyleType value: %d", value)
	}
}

/**************Enum SparklineType *****************/

// Represents the sparkline types.
type SparklineType int32

const(
// Line sparkline.
SparklineType_Line SparklineType = 0 

// Column sparkline.
SparklineType_Column SparklineType = 1 

// Win/Loss sparkline.
SparklineType_Stacked SparklineType = 2 
)

func Int32ToSparklineType(value int32)(SparklineType ,error){
	switch value {
		case 0:  return SparklineType_Line, nil  
		case 1:  return SparklineType_Column, nil  
		case 2:  return SparklineType_Stacked, nil  
		default:
			return 0 ,fmt.Errorf("invalid SparklineType value: %d", value)
	}
}

/**************Enum TickLabelAlignmentType *****************/

// Represents the text alignment type for the tick labels on the axis
type TickLabelAlignmentType int32

const(
// Represents the text shall be centered.
TickLabelAlignmentType_Center TickLabelAlignmentType = 0 

// Represents the text shall be left justified.
TickLabelAlignmentType_Left TickLabelAlignmentType = 1 

// Represents the text shall be right justified.
TickLabelAlignmentType_Right TickLabelAlignmentType = 2 
)

func Int32ToTickLabelAlignmentType(value int32)(TickLabelAlignmentType ,error){
	switch value {
		case 0:  return TickLabelAlignmentType_Center, nil  
		case 1:  return TickLabelAlignmentType_Left, nil  
		case 2:  return TickLabelAlignmentType_Right, nil  
		default:
			return 0 ,fmt.Errorf("invalid TickLabelAlignmentType value: %d", value)
	}
}

/**************Enum TickLabelPositionType *****************/

// Represents the position type of tick-mark labels on the specified axis.
type TickLabelPositionType int32

const(
// Position type is high.
TickLabelPositionType_High TickLabelPositionType = 0 

// Position type is low.
TickLabelPositionType_Low TickLabelPositionType = 1 

// Position type is next to axis.
TickLabelPositionType_NextToAxis TickLabelPositionType = 2 

// Position type is none.
TickLabelPositionType_None TickLabelPositionType = 3 
)

func Int32ToTickLabelPositionType(value int32)(TickLabelPositionType ,error){
	switch value {
		case 0:  return TickLabelPositionType_High, nil  
		case 1:  return TickLabelPositionType_Low, nil  
		case 2:  return TickLabelPositionType_NextToAxis, nil  
		case 3:  return TickLabelPositionType_None, nil  
		default:
			return 0 ,fmt.Errorf("invalid TickLabelPositionType value: %d", value)
	}
}

/**************Enum TickMarkType *****************/

// Represents the tick mark type for the specified axis.
type TickMarkType int32

const(
// Tick mark type is Cross.
TickMarkType_Cross TickMarkType = 0 

// Tick mark type is Inside.
TickMarkType_Inside TickMarkType = 1 

// Tick mark type is None.
TickMarkType_None TickMarkType = 2 

// Tick mark type is Outside
TickMarkType_Outside TickMarkType = 3 
)

func Int32ToTickMarkType(value int32)(TickMarkType ,error){
	switch value {
		case 0:  return TickMarkType_Cross, nil  
		case 1:  return TickMarkType_Inside, nil  
		case 2:  return TickMarkType_None, nil  
		case 3:  return TickMarkType_Outside, nil  
		default:
			return 0 ,fmt.Errorf("invalid TickMarkType value: %d", value)
	}
}

/**************Enum TimeUnit *****************/

// Represents the base unit for the category axis.
type TimeUnit int32

const(
// Days
TimeUnit_Days TimeUnit = 0 

// Months
TimeUnit_Months TimeUnit = 1 

// Years
TimeUnit_Years TimeUnit = 2 
)

func Int32ToTimeUnit(value int32)(TimeUnit ,error){
	switch value {
		case 0:  return TimeUnit_Days, nil  
		case 1:  return TimeUnit_Months, nil  
		case 2:  return TimeUnit_Years, nil  
		default:
			return 0 ,fmt.Errorf("invalid TimeUnit value: %d", value)
	}
}

/**************Enum TrendlineType *****************/

// Represents the trendline type.
type TrendlineType int32

const(
// Exponential
TrendlineType_Exponential TrendlineType = 0 

// Linear
TrendlineType_Linear TrendlineType = 1 

// Logarithmic
TrendlineType_Logarithmic TrendlineType = 2 

// MovingAverage
TrendlineType_MovingAverage TrendlineType = 3 

// Polynomial
TrendlineType_Polynomial TrendlineType = 4 

// Power
TrendlineType_Power TrendlineType = 5 
)

func Int32ToTrendlineType(value int32)(TrendlineType ,error){
	switch value {
		case 0:  return TrendlineType_Exponential, nil  
		case 1:  return TrendlineType_Linear, nil  
		case 2:  return TrendlineType_Logarithmic, nil  
		case 3:  return TrendlineType_MovingAverage, nil  
		case 4:  return TrendlineType_Polynomial, nil  
		case 5:  return TrendlineType_Power, nil  
		default:
			return 0 ,fmt.Errorf("invalid TrendlineType value: %d", value)
	}
}
// Class Axis 

// Encapsulates the object that represents an axis of chart.
type Axis struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *Axis) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Axis_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="Area"/>.
// Returns:
//   Area  
func (instance *Axis) GetArea()  (*Area,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Axis_GetArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Area{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteArea) 

	return result, nil 
}
// Indicates whether the min value is automatically assigned.
// Returns:
//   bool  
func (instance *Axis) IsAutomaticMinValue()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Axis_IsAutomaticMinValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the min value is automatically assigned.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Axis) SetIsAutomaticMinValue(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Axis_SetIsAutomaticMinValue"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the minimum value on the value axis.
// Returns:
//   Object  
func (instance *Axis) GetMinValue()  (*Object,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Axis_GetMinValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Object{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteObject) 

	return result, nil 
}
// Represents the minimum value on the value axis.
// Parameters:
//   value - Object 
// Returns:
//   void  
func (instance *Axis) SetMinValue(value *Object)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("Axis_SetMinValue"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the max value is automatically assigned.
// Returns:
//   bool  
func (instance *Axis) IsAutomaticMaxValue()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Axis_IsAutomaticMaxValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the max value is automatically assigned.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Axis) SetIsAutomaticMaxValue(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Axis_SetIsAutomaticMaxValue"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the maximum value on the value axis.
// Returns:
//   Object  
func (instance *Axis) GetMaxValue()  (*Object,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Axis_GetMaxValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Object{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteObject) 

	return result, nil 
}
// Represents the maximum value on the value axis.
// Parameters:
//   value - Object 
// Returns:
//   void  
func (instance *Axis) SetMaxValue(value *Object)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("Axis_SetMaxValue"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the major unit of the axis is automatically assigned.
// Returns:
//   bool  
func (instance *Axis) IsAutomaticMajorUnit()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Axis_IsAutomaticMajorUnit"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the major unit of the axis is automatically assigned.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Axis) SetIsAutomaticMajorUnit(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Axis_SetIsAutomaticMajorUnit"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the major units for the axis.
// Returns:
//   float64  
func (instance *Axis) GetMajorUnit()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Axis_GetMajorUnit"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the major units for the axis.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Axis) SetMajorUnit(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Axis_SetMajorUnit"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the minor unit of the axis is automatically assigned.
// Returns:
//   bool  
func (instance *Axis) IsAutomaticMinorUnit()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Axis_IsAutomaticMinorUnit"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the minor unit of the axis is automatically assigned.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Axis) SetIsAutomaticMinorUnit(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Axis_SetIsAutomaticMinorUnit"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the minor units for the axis.
// Returns:
//   float64  
func (instance *Axis) GetMinorUnit()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Axis_GetMinorUnit"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the minor units for the axis.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Axis) SetMinorUnit(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Axis_SetMinorUnit"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the appearance of an Axis.
// Returns:
//   Line  
func (instance *Axis) GetAxisLine()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Axis_GetAxisLine"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Represents the type of major tick mark for the specified axis.
// Returns:
//   int32  
func (instance *Axis) GetMajorTickMark()  (TickMarkType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Axis_GetMajorTickMark"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTickMarkType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the type of major tick mark for the specified axis.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Axis) SetMajorTickMark(value TickMarkType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Axis_SetMajorTickMark"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of minor tick mark for the specified axis.
// Returns:
//   int32  
func (instance *Axis) GetMinorTickMark()  (TickMarkType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Axis_GetMinorTickMark"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTickMarkType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the type of minor tick mark for the specified axis.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Axis) SetMinorTickMark(value TickMarkType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Axis_SetMinorTickMark"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the position of tick-mark labels on the specified axis.
// Returns:
//   int32  
func (instance *Axis) GetTickLabelPosition()  (TickLabelPositionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Axis_GetTickLabelPosition"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTickLabelPositionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the position of tick-mark labels on the specified axis.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Axis) SetTickLabelPosition(value TickLabelPositionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Axis_SetTickLabelPosition"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the point on the value axis where the category axis crosses it.
// Returns:
//   float64  
func (instance *Axis) GetCrossAt()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Axis_GetCrossAt"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the point on the value axis where the category axis crosses it.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Axis) SetCrossAt(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Axis_SetCrossAt"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the <see cref="CrossType"/> on the specified axis where the other axis crosses.
// Returns:
//   int32  
func (instance *Axis) GetCrossType()  (CrossType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Axis_GetCrossType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToCrossType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the <see cref="CrossType"/> on the specified axis where the other axis crosses.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Axis) SetCrossType(value CrossType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Axis_SetCrossType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the logarithmic base. Default value is 10.Only applies for Excel2007.
// Returns:
//   float64  
func (instance *Axis) GetLogBase()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Axis_GetLogBase"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the logarithmic base. Default value is 10.Only applies for Excel2007.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Axis) SetLogBase(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Axis_SetLogBase"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents if the value axis scale type is logarithmic or not.
// Returns:
//   bool  
func (instance *Axis) IsLogarithmic()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Axis_IsLogarithmic"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents if the value axis scale type is logarithmic or not.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Axis) SetIsLogarithmic(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Axis_SetIsLogarithmic"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents if Microsoft Excel plots data points from last to first.
// Returns:
//   bool  
func (instance *Axis) IsPlotOrderReversed()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Axis_IsPlotOrderReversed"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents if Microsoft Excel plots data points from last to first.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Axis) SetIsPlotOrderReversed(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Axis_SetIsPlotOrderReversed"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents if the value axis crosses the category axis between categories.
// Returns:
//   bool  
func (instance *Axis) GetAxisBetweenCategories()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Axis_GetAxisBetweenCategories"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents if the value axis crosses the category axis between categories.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Axis) SetAxisBetweenCategories(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Axis_SetAxisBetweenCategories"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns a <see cref="TickLabels"/> object that represents the tick-mark labels for the specified axis.
// Returns:
//   TickLabels  
func (instance *Axis) GetTickLabels()  (*TickLabels,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Axis_GetTickLabels"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TickLabels{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTickLabels) 

	return result, nil 
}
// Represents the number of categories or series between tick-mark labels. Applies only to category and series axes.
// Returns:
//   int32  
func (instance *Axis) GetTickLabelSpacing()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Axis_GetTickLabelSpacing"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the number of categories or series between tick-mark labels. Applies only to category and series axes.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Axis) SetTickLabelSpacing(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Axis_SetTickLabelSpacing"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the spacing of tick label is automatic
// Returns:
//   bool  
func (instance *Axis) IsAutoTickLabelSpacing()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Axis_IsAutoTickLabelSpacing"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the spacing of tick label is automatic
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Axis) SetIsAutoTickLabelSpacing(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Axis_SetIsAutoTickLabelSpacing"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the number of categories or series between tick marks. Applies only to category and series axes.
// Returns:
//   int32  
func (instance *Axis) GetTickMarkSpacing()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Axis_GetTickMarkSpacing"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the number of categories or series between tick marks. Applies only to category and series axes.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Axis) SetTickMarkSpacing(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Axis_SetTickMarkSpacing"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the unit label for the specified axis.
// Returns:
//   int32  
func (instance *Axis) GetDisplayUnit()  (DisplayUnitType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Axis_GetDisplayUnit"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToDisplayUnitType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the unit label for the specified axis.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Axis) SetDisplayUnit(value DisplayUnitType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Axis_SetDisplayUnit"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies a custom value for the display unit.
// Returns:
//   float64  
func (instance *Axis) GetCustomDisplayUnit()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Axis_GetCustomDisplayUnit"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies a custom value for the display unit.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Axis) SetCustomDisplayUnit(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Axis_SetCustomDisplayUnit"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents a unit label on an axis in the specified chart.
// Unit labels are useful for charting large values— for example, in the millions or billions.
// Returns:
//   DisplayUnitLabel  
func (instance *Axis) GetDisplayUnitLabel()  (*DisplayUnitLabel,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Axis_GetDisplayUnitLabel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DisplayUnitLabel{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDisplayUnitLabel) 

	return result, nil 
}
// Represents if the display unit label is shown on the specified axis.
// Returns:
//   bool  
func (instance *Axis) IsDisplayUnitLabelShown()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Axis_IsDisplayUnitLabelShown"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents if the display unit label is shown on the specified axis.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Axis) SetIsDisplayUnitLabelShown(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Axis_SetIsDisplayUnitLabelShown"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the axis' title.
// Returns:
//   Title  
func (instance *Axis) GetTitle()  (*Title,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Axis_GetTitle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Title{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTitle) 

	return result, nil 
}
// Represents the category axis type.
// Returns:
//   int32  
func (instance *Axis) GetCategoryType()  (CategoryType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Axis_GetCategoryType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToCategoryType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the category axis type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Axis) SetCategoryType(value CategoryType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Axis_SetCategoryType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the base unit scale for the category axis.
// Returns:
//   int32  
func (instance *Axis) GetBaseUnitScale()  (TimeUnit,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Axis_GetBaseUnitScale"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTimeUnit(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the base unit scale for the category axis.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Axis) SetBaseUnitScale(value TimeUnit)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Axis_SetBaseUnitScale"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the major unit scale for the category axis.
// Returns:
//   int32  
func (instance *Axis) GetMajorUnitScale()  (TimeUnit,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Axis_GetMajorUnitScale"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTimeUnit(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the major unit scale for the category axis.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Axis) SetMajorUnitScale(value TimeUnit)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Axis_SetMajorUnitScale"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the major unit scale for the category axis.
// Returns:
//   int32  
func (instance *Axis) GetMinorUnitScale()  (TimeUnit,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Axis_GetMinorUnitScale"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTimeUnit(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the major unit scale for the category axis.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Axis) SetMinorUnitScale(value TimeUnit)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Axis_SetMinorUnitScale"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents if the axis is visible.
// Returns:
//   bool  
func (instance *Axis) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Axis_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents if the axis is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Axis) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Axis_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents major gridlines on a chart axis.
// Returns:
//   Line  
func (instance *Axis) GetMajorGridLines()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Axis_GetMajorGridLines"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Represents minor gridlines on a chart axis.
// Returns:
//   Line  
func (instance *Axis) GetMinorGridLines()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Axis_GetMinorGridLines"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Indicates whether the labels shall be shown as multi level.
// Returns:
//   bool  
func (instance *Axis) GetHasMultiLevelLabels()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Axis_GetHasMultiLevelLabels"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the labels shall be shown as multi level.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Axis) SetHasMultiLevelLabels(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Axis_SetHasMultiLevelLabels"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the labels of the axis after call Chart.Calculate() method.
// Returns:
//   []string  
func (instance *Axis) GetAxisTexts()  ([]string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZY(C.CString("Axis_GetAxisTexts"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result:= make([]string, CGoReturnPtr.column_length)
	for i := 0; i < int(CGoReturnPtr.column_length); i++ {
	   offset := uintptr(C.size_t(i)) * uintptr(CGoReturnPtr.size)
	   cObject := *(*C.char)(unsafe.Pointer( uintptr( unsafe.Pointer(CGoReturnPtr.return_value)) + offset))
	   goObject :=string(cObject)
	   result[i] = goObject
	}
	 

	return result, nil 
}
// Represents bins on a chart(Histogram/Pareto) axis
// Returns:
//   AxisBins  
func (instance *Axis) GetBins()  (*AxisBins,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Axis_GetBins"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &AxisBins{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteAxisBins) 

	return result, nil 
}



func DeleteAxis(axis *Axis){
	runtime.SetFinalizer(axis, nil)
	C.Delete_CObject(C.CString("Delete_Axis"),axis.ptr)
	axis.ptr = nil
}

// Class AxisBins 

// Represents axis bins
type AxisBins struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *AxisBins) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("AxisBins_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether grouping data by category
// Returns:
//   bool  
func (instance *AxisBins) IsByCategory()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("AxisBins_IsByCategory"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether grouping data by category
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *AxisBins) SetIsByCategory(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("AxisBins_SetIsByCategory"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the axis bins are automatic.
// Returns:
//   bool  
func (instance *AxisBins) IsAutomatic()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("AxisBins_IsAutomatic"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the axis bins are automatic.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *AxisBins) SetIsAutomatic(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("AxisBins_SetIsAutomatic"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of axis bin
// Returns:
//   float64  
func (instance *AxisBins) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("AxisBins_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of axis bin
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *AxisBins) SetWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("AxisBins_SetWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or set the count of axis bins
// Returns:
//   int32  
func (instance *AxisBins) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("AxisBins_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or set the count of axis bins
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *AxisBins) SetCount(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("AxisBins_SetCount"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or set the overflow of axis bins
// Returns:
//   float64  
func (instance *AxisBins) GetOverflow()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("AxisBins_GetOverflow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or set the overflow of axis bins
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *AxisBins) SetOverflow(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("AxisBins_SetOverflow"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or set the underflow of axis bins
// Returns:
//   float64  
func (instance *AxisBins) GetUnderflow()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("AxisBins_GetUnderflow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or set the underflow of axis bins
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *AxisBins) SetUnderflow(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("AxisBins_SetUnderflow"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteAxisBins(axisbins *AxisBins){
	runtime.SetFinalizer(axisbins, nil)
	C.Delete_CObject(C.CString("Delete_AxisBins"),axisbins.ptr)
	axisbins.ptr = nil
}

// Class Chart 

// Encapsulates the object that represents a single Excel chart.
type Chart struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *Chart) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Chart_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the builtin style.
// Returns:
//   int32  
func (instance *Chart) GetStyle()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Chart_GetStyle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the builtin style.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Chart) SetStyle(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Chart_SetStyle"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the chartShape;
// Returns:
//   ChartShape  
func (instance *Chart) GetChartObject()  (*ChartShape,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetChartObject"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ChartShape{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteChartShape) 

	return result, nil 
}
// Indicates whether hide the pivot chart field buttons only when the chart is PivotChart.
// Returns:
//   bool  
func (instance *Chart) GetHidePivotFieldButtons()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Chart_GetHidePivotFieldButtons"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether hide the pivot chart field buttons only when the chart is PivotChart.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Chart) SetHidePivotFieldButtons(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Chart_SetHidePivotFieldButtons"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the pivot controls that appear on the chart
// Returns:
//   PivotOptions  
func (instance *Chart) GetPivotOptions()  (*PivotOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetPivotOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotOptions) 

	return result, nil 
}
// The source is the data of the pivotTable.
// If PivotSource is not empty ,the chart is PivotChart.
// Returns:
//   string  
func (instance *Chart) GetPivotSource()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Chart_GetPivotSource"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The source is the data of the pivotTable.
// If PivotSource is not empty ,the chart is PivotChart.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Chart) SetPivotSource(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Chart_SetPivotSource"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns whether the cell refered by the chart.
// Parameters:
//   sheetIndex - int32 
//   rowIndex - int32 
//   columnIndex - int32 
// Returns:
//   bool  
func (instance *Chart) IsCellReferedByChart(sheetindex int32, rowindex int32, columnindex int32)  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMO(C.CString("Chart_IsCellReferedByChart"), instance.ptr, C.int(sheetindex), C.int(rowindex), C.int(columnindex))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Detects if a chart's data source has changed.
// Returns:
//   bool  
func (instance *Chart) IsChartDataChanged()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Chart_IsChartDataChanged"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets whether plot by row or column.
// Returns:
//   int32  
func (instance *Chart) GetPlotBy()  (PlotDataByType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Chart_GetPlotBy"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPlotDataByType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Refreshes chart's data from pivot table.
// Returns:
//   void  
func (instance *Chart) RefreshPivotData()  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("Chart_RefreshPivotData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets  how to plot the empty cells.
// Returns:
//   int32  
func (instance *Chart) GetPlotEmptyCellsType()  (PlotEmptyCellsType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Chart_GetPlotEmptyCellsType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPlotEmptyCellsType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets  how to plot the empty cells.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Chart) SetPlotEmptyCellsType(value PlotEmptyCellsType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Chart_SetPlotEmptyCellsType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether plot visible cells only.
// Returns:
//   bool  
func (instance *Chart) GetPlotVisibleCellsOnly()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Chart_GetPlotVisibleCellsOnly"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether plot visible cells only.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Chart) SetPlotVisibleCellsOnly(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Chart_SetPlotVisibleCellsOnly"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether displaying #N/A as blank value.
// Returns:
//   bool  
func (instance *Chart) GetDisplayNaAsBlank()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Chart_GetDisplayNaAsBlank"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether displaying #N/A as blank value.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Chart) SetDisplayNaAsBlank(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Chart_SetDisplayNaAsBlank"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the name of the chart.
// Returns:
//   string  
func (instance *Chart) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Chart_GetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of the chart.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Chart) SetName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Chart_SetName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if Microsoft Excel resizes the chart to match the size of the chart sheet window.
// Returns:
//   bool  
func (instance *Chart) GetSizeWithWindow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Chart_GetSizeWithWindow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if Microsoft Excel resizes the chart to match the size of the chart sheet window.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Chart) SetSizeWithWindow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Chart_SetSizeWithWindow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the worksheet which contains this chart.
// Returns:
//   Worksheet  
func (instance *Chart) GetWorksheet()  (*Worksheet,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetWorksheet"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Returns all drawing shapes in this chart.
// Returns:
//   ShapeCollection  
func (instance *Chart) GetShapes()  (*ShapeCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetShapes"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ShapeCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteShapeCollection) 

	return result, nil 
}
// Gets and sets the printed chart size.
// Returns:
//   int32  
func (instance *Chart) GetPrintSize()  (PrintSizeType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Chart_GetPrintSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPrintSizeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the printed chart size.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Chart) SetPrintSize(value PrintSizeType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Chart_SetPrintSize"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Change chart type with preset template.
// Parameters:
//   data - []byte 
// Returns:
//   void  
func (instance *Chart) ChangeTemplate(data []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("Chart_ChangeTemplate"), instance.ptr, unsafe.Pointer(&data[0]), C.int( len(data)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets a chart's type.
// Returns:
//   int32  
func (instance *Chart) GetType()  (ChartType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Chart_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToChartType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets a chart's type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Chart) SetType(value ChartType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Chart_SetType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Moves the chart to a specified location.
// Parameters:
//   upperLeftRow - int32 
//   upperLeftColumn - int32 
//   lowerRightRow - int32 
//   lowerRightColumn - int32 
// Returns:
//   void  
func (instance *Chart) Move(upperleftrow int32, upperleftcolumn int32, lowerrightrow int32, lowerrightcolumn int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZCD(C.CString("Chart_Move"), instance.ptr, C.int(upperleftrow), C.int(upperleftcolumn), C.int(lowerrightrow), C.int(lowerrightcolumn))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets a <see cref="SeriesCollection"/> collection representing the data series in the chart.
// Returns:
//   SeriesCollection  
func (instance *Chart) GetNSeries()  (*SeriesCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetNSeries"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SeriesCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSeriesCollection) 

	return result, nil 
}
// Gets a <see cref="SeriesCollection"/> collection representing the data series that are filtered in the chart.
// Returns:
//   SeriesCollection  
func (instance *Chart) GetFilteredNSeries()  (*SeriesCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetFilteredNSeries"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SeriesCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSeriesCollection) 

	return result, nil 
}
// Gets the chart's title.
// Returns:
//   Title  
func (instance *Chart) GetTitle()  (*Title,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetTitle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Title{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTitle) 

	return result, nil 
}
// Gets the chart's sub-title.
// Only for ODS format file.
// Returns:
//   Title  
func (instance *Chart) GetSubTitle()  (*Title,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetSubTitle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Title{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTitle) 

	return result, nil 
}
// Gets the chart's plot area which includes axis tick labels.
// Returns:
//   PlotArea  
func (instance *Chart) GetPlotArea()  (*PlotArea,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetPlotArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PlotArea{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePlotArea) 

	return result, nil 
}
// Gets the chart area in the worksheet.
// Returns:
//   ChartArea  
func (instance *Chart) GetChartArea()  (*ChartArea,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetChartArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ChartArea{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteChartArea) 

	return result, nil 
}
// Gets the chart's X axis.
// Returns:
//   Axis  
func (instance *Chart) GetCategoryAxis()  (*Axis,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetCategoryAxis"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Axis{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteAxis) 

	return result, nil 
}
// Gets the chart's Y axis.
// Returns:
//   Axis  
func (instance *Chart) GetValueAxis()  (*Axis,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetValueAxis"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Axis{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteAxis) 

	return result, nil 
}
// Gets the chart's second Y axis.
// Returns:
//   Axis  
func (instance *Chart) GetSecondValueAxis()  (*Axis,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetSecondValueAxis"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Axis{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteAxis) 

	return result, nil 
}
// Gets the chart's second X axis.
// Returns:
//   Axis  
func (instance *Chart) GetSecondCategoryAxis()  (*Axis,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetSecondCategoryAxis"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Axis{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteAxis) 

	return result, nil 
}
// Gets the chart's series axis.
// Returns:
//   Axis  
func (instance *Chart) GetSeriesAxis()  (*Axis,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetSeriesAxis"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Axis{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteAxis) 

	return result, nil 
}
// Gets the chart legend.
// Returns:
//   Legend  
func (instance *Chart) GetLegend()  (*Legend,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetLegend"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Legend{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLegend) 

	return result, nil 
}
// Represents the chart data table.
// Returns:
//   ChartDataTable  
func (instance *Chart) GetChartDataTable()  (*ChartDataTable,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetChartDataTable"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ChartDataTable{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteChartDataTable) 

	return result, nil 
}
// Gets or sets a value indicating whether the chart legend will be displayed. Default is true.
// Returns:
//   bool  
func (instance *Chart) GetShowLegend()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Chart_GetShowLegend"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets a value indicating whether the chart legend will be displayed. Default is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Chart) SetShowLegend(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Chart_SetShowLegend"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets a value indicating whether the chart area is rectangular cornered.
// Default is true.
// Returns:
//   bool  
func (instance *Chart) IsRectangularCornered()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Chart_IsRectangularCornered"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets a value indicating whether the chart area is rectangular cornered.
// Default is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Chart) SetIsRectangularCornered(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Chart_SetIsRectangularCornered"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets a value indicating whether the chart displays a data table.
// Returns:
//   bool  
func (instance *Chart) GetShowDataTable()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Chart_GetShowDataTable"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets a value indicating whether the chart displays a data table.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Chart) SetShowDataTable(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Chart_SetShowDataTable"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the angle of the first pie-chart or doughnut-chart slice, in degrees (clockwise from vertical).
// Applies only to pie, 3-D pie, and doughnut charts, 0 to 360.
// Returns:
//   int32  
func (instance *Chart) GetFirstSliceAngle()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Chart_GetFirstSliceAngle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the angle of the first pie-chart or doughnut-chart slice, in degrees (clockwise from vertical).
// Applies only to pie, 3-D pie, and doughnut charts, 0 to 360.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Chart) SetFirstSliceAngle(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Chart_SetFirstSliceAngle"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the space between bar or column clusters, as a percentage of the bar or column width.
// The value of this property must be between 0 and 500.
// Returns:
//   int32  
func (instance *Chart) GetGapWidth()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Chart_GetGapWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the space between bar or column clusters, as a percentage of the bar or column width.
// The value of this property must be between 0 and 500.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Chart) SetGapWidth(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Chart_SetGapWidth"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the distance between the data series in a 3-D chart, as a percentage of the marker width.
// The value of this property must be between 0 and 500.
// Returns:
//   int32  
func (instance *Chart) GetGapDepth()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Chart_GetGapDepth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the distance between the data series in a 3-D chart, as a percentage of the marker width.
// The value of this property must be between 0 and 500.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Chart) SetGapDepth(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Chart_SetGapDepth"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Calculates the custom position of plot area, axes if the position of them are auto assigned.
// Returns:
//   void  
func (instance *Chart) Calculate()  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("Chart_Calculate"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Calculates the custom position of plot area, axes if the position of them are auto assigned, with Chart Calculate Options.
// Parameters:
//   calculateOptions - ChartCalculateOptions 
// Returns:
//   void  
func (instance *Chart) Calculate_ChartCalculateOptions(calculateoptions *ChartCalculateOptions)  error {
	
	var calculateoptions_ptr unsafe.Pointer = nil
	if calculateoptions != nil {
	  calculateoptions_ptr =calculateoptions.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("Chart_Calculate_ChartCalculateOptions"), instance.ptr, calculateoptions_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns a <see cref="Floor"/> object that represents the walls of a 3-D chart.
// Returns:
//   Floor  
func (instance *Chart) GetFloor()  (*Floor,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetFloor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Floor{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFloor) 

	return result, nil 
}
// Returns a <see cref="Walls"/> object that represents the walls of a 3-D chart.
// Returns:
//   Walls  
func (instance *Chart) GetWalls()  (*Walls,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetWalls"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Walls{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWalls) 

	return result, nil 
}
// Returns a <see cref="Walls"/> object that represents the back wall of a 3-D chart.
// Returns:
//   Walls  
func (instance *Chart) GetBackWall()  (*Walls,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetBackWall"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Walls{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWalls) 

	return result, nil 
}
// Returns a <see cref="Walls"/> object that represents the side wall of a 3-D chart.
// Returns:
//   Walls  
func (instance *Chart) GetSideWall()  (*Walls,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetSideWall"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Walls{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWalls) 

	return result, nil 
}
// True if gridlines are drawn two-dimensionally on a 3-D chart.
// Returns:
//   bool  
func (instance *Chart) GetWallsAndGridlines2D()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Chart_GetWallsAndGridlines2D"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if gridlines are drawn two-dimensionally on a 3-D chart.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Chart) SetWallsAndGridlines2D(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Chart_SetWallsAndGridlines2D"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the rotation of the 3-D chart view (the rotation of the plot area around the z-axis, in degrees).
// Returns:
//   int32  
func (instance *Chart) GetRotationAngle()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Chart_GetRotationAngle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the rotation of the 3-D chart view (the rotation of the plot area around the z-axis, in degrees).
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Chart) SetRotationAngle(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Chart_SetRotationAngle"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the elevation of the 3-D chart view, in degrees.
// Returns:
//   int32  
func (instance *Chart) GetElevation()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Chart_GetElevation"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the elevation of the 3-D chart view, in degrees.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Chart) SetElevation(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Chart_SetElevation"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the chart axes are at right angles. Applies only for 3-D charts(except Column3D and 3-D Pie Charts).
// Returns:
//   bool  
func (instance *Chart) GetRightAngleAxes()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Chart_GetRightAngleAxes"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the chart axes are at right angles. Applies only for 3-D charts(except Column3D and 3-D Pie Charts).
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Chart) SetRightAngleAxes(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Chart_SetRightAngleAxes"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if Microsoft Excel scales a 3-D chart so that it's closer in size to the equivalent 2-D chart.
// The RightAngleAxes property must be True.
// Returns:
//   bool  
func (instance *Chart) GetAutoScaling()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Chart_GetAutoScaling"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if Microsoft Excel scales a 3-D chart so that it's closer in size to the equivalent 2-D chart.
// The RightAngleAxes property must be True.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Chart) SetAutoScaling(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Chart_SetAutoScaling"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the height of a 3-D chart as a percentage of the chart width (between 5 and 500 percent).
// Returns:
//   int32  
func (instance *Chart) GetHeightPercent()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Chart_GetHeightPercent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the height of a 3-D chart as a percentage of the chart width (between 5 and 500 percent).
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Chart) SetHeightPercent(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Chart_SetHeightPercent"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the perspective for the 3-D chart view. Must be between 0 and 100.
// This property is ignored if the RightAngleAxes property is True.
// Returns:
//   int16  
func (instance *Chart) GetPerspective()  (int16,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZHT(C.CString("Chart_GetPerspective"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int16(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the perspective for the 3-D chart view. Must be between 0 and 100.
// This property is ignored if the RightAngleAxes property is True.
// Parameters:
//   value - int16 
// Returns:
//   void  
func (instance *Chart) SetPerspective(value int16)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMP(C.CString("Chart_SetPerspective"), instance.ptr, C.short(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the chart is a 3d chart.
// Returns:
//   bool  
func (instance *Chart) GetIs3D()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Chart_GetIs3D"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the depth of a 3-D chart as a percentage of the chart width (between 20 and 2000 percent).
// Returns:
//   int32  
func (instance *Chart) GetDepthPercent()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Chart_GetDepthPercent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the depth of a 3-D chart as a percentage of the chart width (between 20 and 2000 percent).
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Chart) SetDepthPercent(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Chart_SetDepthPercent"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Creates the chart image and saves it to a file.
// The extension of the file name determines the format of the image.
// Parameters:
//   imageFile - string 
// Returns:
//   void  
func (instance *Chart) ToImage_String(imagefile string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Chart_ToImage_String"), instance.ptr, C.CString(imagefile))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Creates the chart image and saves it to a file in the specified image type.
// Parameters:
//   imageFile - string 
//   imageType - int32 
// Returns:
//   void  
func (instance *Chart) ToImage_String_ImageType(imagefile string, imagetype ImageType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZJD(C.CString("Chart_ToImage_String_ImageType"), instance.ptr, C.CString(imagefile), C.int( int32(imagetype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Creates the chart image and saves it to a stream in the Jpeg format.
// Parameters:
//   jpegQuality - int64 
// Returns:
//   []byte  
func (instance *Chart) ToImage_Int64(jpegquality int64)  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMQ(C.CString("Chart_ToImage_Long"), instance.ptr, C.long(jpegquality))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Creates the chart image and saves it to a stream in the specified format.
// Parameters:
//   imageType - int32 
// Returns:
//   []byte  
func (instance *Chart) ToImage_ImageType(imagetype ImageType)  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZJE(C.CString("Chart_ToImage_ImageType"), instance.ptr, C.int( int32(imagetype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Saves the chart to a pdf file.
// Parameters:
//   fileName - string 
// Returns:
//   void  
func (instance *Chart) ToPdf_String(filename string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Chart_ToPdf_String"), instance.ptr, C.CString(filename))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Saves the chart to a pdf file.
// Parameters:
//   fileName - string 
//   desiredPageWidth - float32 
//   desiredPageHeight - float32 
//   hAlignmentType - int32 
//   vAlignmentType - int32 
// Returns:
//   void  
func (instance *Chart) ToPdf_String_Float_Float_PageLayoutAlignmentType_PageLayoutAlignmentType(filename string, desiredpagewidth float32, desiredpageheight float32, halignmenttype PageLayoutAlignmentType, valignmenttype PageLayoutAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMR(C.CString("Chart_ToPdf_String_Floating_Floating_PageLayoutAlignmentType_PageLayoutAlignmentType"), instance.ptr, C.CString(filename), C.float(desiredpagewidth), C.float(desiredpageheight), C.int( int32(halignmenttype)), C.int( int32(valignmenttype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Creates the chart pdf and saves it to a stream.
// Returns:
//   []byte  
func (instance *Chart) ToPdf()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("Chart_ToPdf"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Creates the chart pdf and saves it to a stream.
// Parameters:
//   desiredPageWidth - float32 
//   desiredPageHeight - float32 
//   hAlignmentType - int32 
//   vAlignmentType - int32 
// Returns:
//   []byte  
func (instance *Chart) ToPdf_Float_Float_PageLayoutAlignmentType_PageLayoutAlignmentType(desiredpagewidth float32, desiredpageheight float32, halignmenttype PageLayoutAlignmentType, valignmenttype PageLayoutAlignmentType)  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMS(C.CString("Chart_ToPdf_Floating_Floating_PageLayoutAlignmentType_PageLayoutAlignmentType"), instance.ptr, C.float(desiredpagewidth), C.float(desiredpageheight), C.int( int32(halignmenttype)), C.int( int32(valignmenttype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Creates the chart image and saves it to a file.
// The extension of the file name determines the format of the image.
// Parameters:
//   imageFile - string 
//   options - ImageOrPrintOptions 
// Returns:
//   void  
func (instance *Chart) ToImage_String_ImageOrPrintOptions(imagefile string, options *ImageOrPrintOptions)  error {
	
	var options_ptr unsafe.Pointer = nil
	if options != nil {
	  options_ptr =options.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZAQ(C.CString("Chart_ToImage_String_ImageOrPrintOptions"), instance.ptr, C.CString(imagefile), options_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Creates the chart image and saves it to a stream in the specified format.
// Parameters:
//   options - ImageOrPrintOptions 
// Returns:
//   []byte  
func (instance *Chart) ToImage_ImageOrPrintOptions(options *ImageOrPrintOptions)  ([]byte,  error)  {
	
	var options_ptr unsafe.Pointer = nil
	if options != nil {
	  options_ptr =options.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZET(C.CString("Chart_ToImage_ImageOrPrintOptions"), instance.ptr, options_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets actual size of chart in unit of pixels.
// Returns:
//   []int32_t  
func (instance *Chart) GetActualSize()  ([]int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZY(C.CString("Chart_GetActualSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result:= make([]int32, CGoReturnPtr.column_length)
	for i := 0; i < int(CGoReturnPtr.column_length); i++ {
	   offset := uintptr(C.size_t(i)) * uintptr(CGoReturnPtr.size)
	   cObject := *(*C.int)(unsafe.Pointer( uintptr( unsafe.Pointer(CGoReturnPtr.return_value)) + offset))
	   goObject :=int32(cObject)
	   result[i] = goObject
	}
	 

	return result, nil 
}
// Represents the way the chart is attached to the cells below it.
// Returns:
//   int32  
func (instance *Chart) GetPlacement()  (PlacementType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Chart_GetPlacement"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPlacementType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the way the chart is attached to the cells below it.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Chart) SetPlacement(value PlacementType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Chart_SetPlacement"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the page setup description in this chart.
// Returns:
//   PageSetup  
func (instance *Chart) GetPageSetup()  (*PageSetup,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetPageSetup"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PageSetup{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePageSetup) 

	return result, nil 
}
// Returns which axes exist on the chart.
// Parameters:
//   aixsType - int32 
//   isPrimary - bool 
// Returns:
//   bool  
func (instance *Chart) HasAxis(aixstype AxisType, isprimary bool)  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMT(C.CString("Chart_HasAxis"), instance.ptr, C.int( int32(aixstype)), C.bool(isprimary))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Switches row/column.
// Returns:
//   bool  
func (instance *Chart) SwitchRowColumn()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Chart_SwitchRowColumn"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the data source range of the chart.
// Returns:
//   string  
func (instance *Chart) GetChartDataRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Chart_GetChartDataRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies data range for a chart.
// Parameters:
//   area - string 
//   isVertical - bool 
// Returns:
//   void  
func (instance *Chart) SetChartDataRange(area string, isvertical bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAL(C.CString("Chart_SetChartDataRange"), instance.ptr, C.CString(area), C.bool(isvertical))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the line.
// Returns:
//   Line  
func (instance *Chart) GetLine()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Chart_GetLine"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}



func DeleteChart(chart *Chart){
	runtime.SetFinalizer(chart, nil)
	C.Delete_CObject(C.CString("Delete_Chart"),chart.ptr)
	chart.ptr = nil
}

// Class ChartArea 

// Encapsulates the object that represents the chart area in the worksheet.
type ChartArea struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ChartFrame 
func NewChartArea(src *ChartFrame) ( *ChartArea, error) {
	chartarea := &ChartArea{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_ChartArea"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		chartarea.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(chartarea, DeleteChartArea)
		return chartarea, nil
	} else {
		chartarea.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return chartarea, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ChartArea) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartArea_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or gets the horizontal offset from its upper left corner column, in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *ChartArea) GetXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartArea_GetXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or gets the horizontal offset from its upper left corner column, in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ChartArea) SetXRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ChartArea_SetXRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or gets the vertical offset from its upper left corner row, in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *ChartArea) GetYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartArea_GetYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or gets the vertical offset from its upper left corner row, in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ChartArea) SetYRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ChartArea_SetYRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the vertical offset from its lower right corner row, in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *ChartArea) GetHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartArea_GetHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the vertical offset from its lower right corner row, in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ChartArea) SetHeightRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ChartArea_SetHeightRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the horizontal offset from its lower right corner column, in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *ChartArea) GetWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartArea_GetWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the horizontal offset from its lower right corner column, in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ChartArea) SetWidthRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ChartArea_SetWidthRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets a <see cref="Font"/> object of the specified chartarea object.
// Returns:
//   Font  
func (instance *ChartArea) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartArea_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.
// False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
// Returns:
//   bool  
func (instance *ChartArea) IsInnerMode()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartArea_IsInnerMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.
// False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartArea) SetIsInnerMode(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartArea_SetIsInnerMode"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the display mode of the background
// Returns:
//   int32  
func (instance *ChartArea) GetBackgroundMode()  (BackgroundMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ChartArea_GetBackgroundMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBackgroundMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the display mode of the background
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartArea) SetBackgroundMode(value BackgroundMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ChartArea_SetBackgroundMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the frame has a shadow.
// Returns:
//   bool  
func (instance *ChartArea) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartArea_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the frame has a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartArea) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartArea_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="ShapeProperties"/> object.
// Returns:
//   ShapePropertyCollection  
func (instance *ChartArea) GetShapeProperties()  (*ShapePropertyCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartArea_GetShapeProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ShapePropertyCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteShapePropertyCollection) 

	return result, nil 
}
// Indicates whether default position(DefaultX, DefaultY, DefaultWidth and DefaultHeight) are set.
// Returns:
//   bool  
func (instance *ChartArea) IsDefaultPosBeSet()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartArea_IsDefaultPosBeSet"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents x of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *ChartArea) GetDefaultXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartArea_GetDefaultXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents y of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *ChartArea) GetDefaultYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartArea_GetDefaultYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents width of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *ChartArea) GetDefaultWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartArea_GetDefaultWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents height of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *ChartArea) GetDefaultHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartArea_GetDefaultHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="Line">border</see>.
// Returns:
//   Line  
func (instance *ChartArea) GetBorder()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartArea_GetBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Gets the <see cref="Area">area</see>.
// Returns:
//   Area  
func (instance *ChartArea) GetArea()  (*Area,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartArea_GetArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Area{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteArea) 

	return result, nil 
}
// Gets and sets the options of the text.
// Returns:
//   TextOptions  
func (instance *ChartArea) GetTextOptions()  (*TextOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartArea_GetTextOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TextOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTextOptions) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Returns:
//   bool  
func (instance *ChartArea) GetAutoScaleFont()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartArea_GetAutoScaleFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartArea) SetAutoScaleFont(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartArea_SetAutoScaleFont"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the chart frame is automatic sized.
// Returns:
//   bool  
func (instance *ChartArea) IsAutomaticSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartArea_IsAutomaticSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the chart frame is automatic sized.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartArea) SetIsAutomaticSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartArea_SetIsAutomaticSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Pixel.
// Returns:
//   int32  
func (instance *ChartArea) GetXPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartArea_GetXPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartArea) SetXPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ChartArea_SetXPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Pixel.
// Returns:
//   int32  
func (instance *ChartArea) GetYPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartArea_GetYPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartArea) SetYPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ChartArea_SetYPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of frame in units of Pixel.
// Returns:
//   int32  
func (instance *ChartArea) GetWidthPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartArea_GetWidthPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of frame in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartArea) SetWidthPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ChartArea_SetWidthPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the height of frame in units of Pixel.
// Returns:
//   int32  
func (instance *ChartArea) GetHeightPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartArea_GetHeightPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the height of frame in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartArea) SetHeightPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ChartArea_SetHeightPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Set position of the frame to automatic
// Returns:
//   void  
func (instance *ChartArea) SetPositionAuto()  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("ChartArea_SetPositionAuto"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *ChartArea) ToChartFrame() *ChartFrame {
	parentClass := &ChartFrame{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteChartArea(chartarea *ChartArea){
	runtime.SetFinalizer(chartarea, nil)
	C.Delete_CObject(C.CString("Delete_ChartArea"),chartarea.ptr)
	chartarea.ptr = nil
}

// Class ChartCalculateOptions 

// Represents the options for calculating chart.
type ChartCalculateOptions struct {
	ptr unsafe.Pointer
}

// Creates the options for calculating chart.
func NewChartCalculateOptions() ( *ChartCalculateOptions, error) {
	chartcalculateoptions := &ChartCalculateOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_ChartCalculateOptions"),)
	if CGoReturnPtr.error_no == 0 {
		chartcalculateoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(chartcalculateoptions, DeleteChartCalculateOptions)
		return chartcalculateoptions, nil
	} else {
		chartcalculateoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return chartcalculateoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ChartCalculateOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartCalculateOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Whether update all data points when performing the chart calculation. Default: False.
// When you want to get the value for each data point in the chart specifically, set it to true.
// If this parameter is set to True, the new data points may be generated when chart is calculated. This could make the Excel file larger.
// Returns:
//   bool  
func (instance *ChartCalculateOptions) GetUpdateAllPoints()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartCalculateOptions_GetUpdateAllPoints"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Whether update all data points when performing the chart calculation. Default: False.
// When you want to get the value for each data point in the chart specifically, set it to true.
// If this parameter is set to True, the new data points may be generated when chart is calculated. This could make the Excel file larger.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartCalculateOptions) SetUpdateAllPoints(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartCalculateOptions_SetUpdateAllPoints"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteChartCalculateOptions(chartcalculateoptions *ChartCalculateOptions){
	runtime.SetFinalizer(chartcalculateoptions, nil)
	C.Delete_CObject(C.CString("Delete_ChartCalculateOptions"),chartcalculateoptions.ptr)
	chartcalculateoptions.ptr = nil
}

// Class ChartCollection 

// Encapsulates a collection of <see cref="Chart"/> objects.
type ChartCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ChartCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a chart to the collection.
// Parameters:
//   type - int32 
//   left - int32 
//   top - int32 
//   width - int32 
//   height - int32 
// Returns:
//   int32  
func (instance *ChartCollection) AddFloatingChart(type_ ChartType, left int32, top int32, width int32, height int32)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMU(C.CString("ChartCollection_AddFloatingChart"), instance.ptr, C.int( int32(type_)), C.int(left), C.int(top), C.int(width), C.int(height))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a chart to the collection.
// Parameters:
//   type - int32 
//   upperLeftRow - int32 
//   upperLeftColumn - int32 
//   lowerRightRow - int32 
//   lowerRightColumn - int32 
// Returns:
//   int32  
func (instance *ChartCollection) Add_ChartType_Int_Int_Int_Int(type_ ChartType, upperleftrow int32, upperleftcolumn int32, lowerrightrow int32, lowerrightcolumn int32)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMU(C.CString("ChartCollection_Add_ChartType_Integer_Integer_Integer_Integer"), instance.ptr, C.int( int32(type_)), C.int(upperleftrow), C.int(upperleftcolumn), C.int(lowerrightrow), C.int(lowerrightcolumn))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a chart with preset template.
// Parameters:
//   data - []byte 
//   dataRange - string 
//   isVertical - bool 
//   topRow - int32 
//   leftColumn - int32 
//   rightRow - int32 
//   bottomColumn - int32 
// Returns:
//   int32  
func (instance *ChartCollection) Add_Stream_String_Bool_Int_Int_Int_Int(data []byte, datarange string, isvertical bool, toprow int32, leftcolumn int32, rightrow int32, bottomcolumn int32)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMV(C.CString("ChartCollection_Add_Stream_String_Boolean_Integer_Integer_Integer_Integer"), instance.ptr, unsafe.Pointer(&data[0]), C.int( len(data)), C.CString(datarange), C.bool(isvertical), C.int(toprow), C.int(leftcolumn), C.int(rightrow), C.int(bottomcolumn))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a chart to the collection.
// Parameters:
//   type - int32 
//   dataRange - string 
//   isVertical - bool 
//   topRow - int32 
//   leftColumn - int32 
//   rightRow - int32 
//   bottomColumn - int32 
// Returns:
//   int32  
func (instance *ChartCollection) Add_ChartType_String_Bool_Int_Int_Int_Int(type_ ChartType, datarange string, isvertical bool, toprow int32, leftcolumn int32, rightrow int32, bottomcolumn int32)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMW(C.CString("ChartCollection_Add_ChartType_String_Boolean_Integer_Integer_Integer_Integer"), instance.ptr, C.int( int32(type_)), C.CString(datarange), C.bool(isvertical), C.int(toprow), C.int(leftcolumn), C.int(rightrow), C.int(bottomcolumn))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="Chart"/> element at the specified index.
// Parameters:
//   index - int32 
// Returns:
//   Chart  
func (instance *ChartCollection) Get_Int(index int32)  (*Chart,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("ChartCollection_Get_Integer"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Chart{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteChart) 

	return result, nil 
}
// Gets the chart by the name.
// Parameters:
//   name - string 
// Returns:
//   Chart  
func (instance *ChartCollection) Get_String(name string)  (*Chart,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBU(C.CString("ChartCollection_Get_String"), instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Chart{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteChart) 

	return result, nil 
}
// Remove a chart at the specific index.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *ChartCollection) RemoveAt(index int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ChartCollection_RemoveAt"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Clear all charts.
// Returns:
//   void  
func (instance *ChartCollection) Clear()  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("ChartCollection_Clear"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns:
//   int32  
func (instance *ChartCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteChartCollection(chartcollection *ChartCollection){
	runtime.SetFinalizer(chartcollection, nil)
	C.Delete_CObject(C.CString("Delete_ChartCollection"),chartcollection.ptr)
	chartcollection.ptr = nil
}

// Class ChartDataTable 

// Represents a chart data table.
type ChartDataTable struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ChartDataTable) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartDataTable_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets a <see cref="Font"/> object which represents the font setting of the specified chart data table.
// Returns:
//   Font  
func (instance *ChartDataTable) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartDataTable_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes.
// The default value is True.
// Returns:
//   bool  
func (instance *ChartDataTable) GetAutoScaleFont()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartDataTable_GetAutoScaleFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes.
// The default value is True.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartDataTable) SetAutoScaleFont(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartDataTable_SetAutoScaleFont"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the display mode of the background
// Returns:
//   int32  
func (instance *ChartDataTable) GetBackgroundMode()  (BackgroundMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ChartDataTable_GetBackgroundMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBackgroundMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the display mode of the background
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartDataTable) SetBackgroundMode(value BackgroundMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ChartDataTable_SetBackgroundMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the chart data table has horizontal cell borders
// Returns:
//   bool  
func (instance *ChartDataTable) GetHasHorizontalBorder()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartDataTable_GetHasHorizontalBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the chart data table has horizontal cell borders
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartDataTable) SetHasHorizontalBorder(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartDataTable_SetHasHorizontalBorder"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the chart data table has vertical cell borders
// Returns:
//   bool  
func (instance *ChartDataTable) GetHasVerticalBorder()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartDataTable_GetHasVerticalBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the chart data table has vertical cell borders
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartDataTable) SetHasVerticalBorder(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartDataTable_SetHasVerticalBorder"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the chart data table has outline borders
// Returns:
//   bool  
func (instance *ChartDataTable) GetHasOutlineBorder()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartDataTable_GetHasOutlineBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the chart data table has outline borders
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartDataTable) SetHasOutlineBorder(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartDataTable_SetHasOutlineBorder"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the data label legend key is visible.
// Returns:
//   bool  
func (instance *ChartDataTable) GetShowLegendKey()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartDataTable_GetShowLegendKey"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the data label legend key is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartDataTable) SetShowLegendKey(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartDataTable_SetShowLegendKey"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns a Border object that represents the border of the object
// Returns:
//   Line  
func (instance *ChartDataTable) GetBorder()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartDataTable_GetBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}



func DeleteChartDataTable(chartdatatable *ChartDataTable){
	runtime.SetFinalizer(chartdatatable, nil)
	C.Delete_CObject(C.CString("Delete_ChartDataTable"),chartdatatable.ptr)
	chartdatatable.ptr = nil
}

// Class ChartFrame 

// Encapsulates the object that represents the frame object in a chart.
type ChartFrame struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ChartFrame) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartFrame_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.
// False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
// Returns:
//   bool  
func (instance *ChartFrame) IsInnerMode()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartFrame_IsInnerMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.
// False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartFrame) SetIsInnerMode(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartFrame_SetIsInnerMode"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the display mode of the background
// Returns:
//   int32  
func (instance *ChartFrame) GetBackgroundMode()  (BackgroundMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ChartFrame_GetBackgroundMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBackgroundMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the display mode of the background
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartFrame) SetBackgroundMode(value BackgroundMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ChartFrame_SetBackgroundMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the frame has a shadow.
// Returns:
//   bool  
func (instance *ChartFrame) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartFrame_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the frame has a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartFrame) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartFrame_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="ShapeProperties"/> object.
// Returns:
//   ShapePropertyCollection  
func (instance *ChartFrame) GetShapeProperties()  (*ShapePropertyCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartFrame_GetShapeProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ShapePropertyCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteShapePropertyCollection) 

	return result, nil 
}
// Indicates whether default position(DefaultX, DefaultY, DefaultWidth and DefaultHeight) are set.
// Returns:
//   bool  
func (instance *ChartFrame) IsDefaultPosBeSet()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartFrame_IsDefaultPosBeSet"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents x of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *ChartFrame) GetDefaultXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartFrame_GetDefaultXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents y of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *ChartFrame) GetDefaultYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartFrame_GetDefaultYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents width of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *ChartFrame) GetDefaultWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartFrame_GetDefaultWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents height of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *ChartFrame) GetDefaultHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartFrame_GetDefaultHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="Line">border</see>.
// Returns:
//   Line  
func (instance *ChartFrame) GetBorder()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartFrame_GetBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Gets the <see cref="Area">area</see>.
// Returns:
//   Area  
func (instance *ChartFrame) GetArea()  (*Area,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartFrame_GetArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Area{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteArea) 

	return result, nil 
}
// Gets and sets the options of the text.
// Returns:
//   TextOptions  
func (instance *ChartFrame) GetTextOptions()  (*TextOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartFrame_GetTextOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TextOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTextOptions) 

	return result, nil 
}
// Gets a <see cref="Font"/> object of the specified ChartFrame object.
// Returns:
//   Font  
func (instance *ChartFrame) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartFrame_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Returns:
//   bool  
func (instance *ChartFrame) GetAutoScaleFont()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartFrame_GetAutoScaleFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartFrame) SetAutoScaleFont(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartFrame_SetAutoScaleFont"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the chart frame is automatic sized.
// Returns:
//   bool  
func (instance *ChartFrame) IsAutomaticSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartFrame_IsAutomaticSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the chart frame is automatic sized.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartFrame) SetIsAutomaticSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartFrame_SetIsAutomaticSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the x coordinate of the upper left corner in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *ChartFrame) GetXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartFrame_GetXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the x coordinate of the upper left corner in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ChartFrame) SetXRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ChartFrame_SetXRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the y coordinate of the upper left corner in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *ChartFrame) GetYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartFrame_GetYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the y coordinate of the upper left corner in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ChartFrame) SetYRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ChartFrame_SetYRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of frame in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *ChartFrame) GetWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartFrame_GetWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of frame in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ChartFrame) SetWidthRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ChartFrame_SetWidthRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the height of frame in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *ChartFrame) GetHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartFrame_GetHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the height of frame in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ChartFrame) SetHeightRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ChartFrame_SetHeightRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Pixel.
// Returns:
//   int32  
func (instance *ChartFrame) GetXPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartFrame_GetXPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartFrame) SetXPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ChartFrame_SetXPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Pixel.
// Returns:
//   int32  
func (instance *ChartFrame) GetYPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartFrame_GetYPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartFrame) SetYPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ChartFrame_SetYPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of frame in units of Pixel.
// Returns:
//   int32  
func (instance *ChartFrame) GetWidthPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartFrame_GetWidthPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of frame in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartFrame) SetWidthPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ChartFrame_SetWidthPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the height of frame in units of Pixel.
// Returns:
//   int32  
func (instance *ChartFrame) GetHeightPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartFrame_GetHeightPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the height of frame in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartFrame) SetHeightPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ChartFrame_SetHeightPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Set position of the frame to automatic
// Returns:
//   void  
func (instance *ChartFrame) SetPositionAuto()  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("ChartFrame_SetPositionAuto"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteChartFrame(chartframe *ChartFrame){
	runtime.SetFinalizer(chartframe, nil)
	C.Delete_CObject(C.CString("Delete_ChartFrame"),chartframe.ptr)
	chartframe.ptr = nil
}

// Class ChartGlobalizationSettings 

// Represents the globalization settings for chart.
type ChartGlobalizationSettings struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewChartGlobalizationSettings() ( *ChartGlobalizationSettings, error) {
	chartglobalizationsettings := &ChartGlobalizationSettings{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_ChartGlobalizationSettings"),)
	if CGoReturnPtr.error_no == 0 {
		chartglobalizationsettings.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(chartglobalizationsettings, DeleteChartGlobalizationSettings)
		return chartglobalizationsettings, nil
	} else {
		chartglobalizationsettings.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return chartglobalizationsettings, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ChartGlobalizationSettings) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartGlobalizationSettings_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of Series in the Chart.
// Returns:
//   string  
func (instance *ChartGlobalizationSettings) GetSeriesName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ChartGlobalizationSettings_GetSeriesName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of Chart Title.
// Returns:
//   string  
func (instance *ChartGlobalizationSettings) GetChartTitleName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ChartGlobalizationSettings_GetChartTitleName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of increase for Legend.
// Returns:
//   string  
func (instance *ChartGlobalizationSettings) GetLegendIncreaseName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ChartGlobalizationSettings_GetLegendIncreaseName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of Decrease for Legend.
// Returns:
//   string  
func (instance *ChartGlobalizationSettings) GetLegendDecreaseName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ChartGlobalizationSettings_GetLegendDecreaseName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of Total for Legend.
// Returns:
//   string  
func (instance *ChartGlobalizationSettings) GetLegendTotalName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ChartGlobalizationSettings_GetLegendTotalName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of Title for Axis.
// Returns:
//   string  
func (instance *ChartGlobalizationSettings) GetAxisTitleName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ChartGlobalizationSettings_GetAxisTitleName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of "Other" labels for Chart.
// Returns:
//   string  
func (instance *ChartGlobalizationSettings) GetOtherName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ChartGlobalizationSettings_GetOtherName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the Name of Axis Unit.
// Parameters:
//   type - int32 
// Returns:
//   string  
func (instance *ChartGlobalizationSettings) GetAxisUnitName(type_ DisplayUnitType)  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAM(C.CString("ChartGlobalizationSettings_GetAxisUnitName"), instance.ptr, C.int( int32(type_)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteChartGlobalizationSettings(chartglobalizationsettings *ChartGlobalizationSettings){
	runtime.SetFinalizer(chartglobalizationsettings, nil)
	C.Delete_CObject(C.CString("Delete_ChartGlobalizationSettings"),chartglobalizationsettings.ptr)
	chartglobalizationsettings.ptr = nil
}

// Class ChartPoint 

// Represents a single point in a series in a chart.
type ChartPoint struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ChartPoint) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartPoint_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// The distance of an open pie slice from the center of the pie chart is expressed as a percentage of the pie diameter.
// Returns:
//   int32  
func (instance *ChartPoint) GetExplosion()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartPoint_GetExplosion"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// The distance of an open pie slice from the center of the pie chart is expressed as a percentage of the pie diameter.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartPoint) SetExplosion(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ChartPoint_SetExplosion"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the chartpoint has a shadow.
// Returns:
//   bool  
func (instance *ChartPoint) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartPoint_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the chartpoint has a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartPoint) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartPoint_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="Line"> border</see>.
// Returns:
//   Line  
func (instance *ChartPoint) GetBorder()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartPoint_GetBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Gets the <see cref="Area"> area</see>.
// Returns:
//   Area  
func (instance *ChartPoint) GetArea()  (*Area,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartPoint_GetArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Area{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteArea) 

	return result, nil 
}
// Gets the <see cref="Marker"> marker</see>.
// Returns:
//   Marker  
func (instance *ChartPoint) GetMarker()  (*Marker,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartPoint_GetMarker"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Marker{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteMarker) 

	return result, nil 
}
// Returns a <see cref="DataLabels"/> object that represents the data label associated with this chart point.
// Returns:
//   DataLabels  
func (instance *ChartPoint) GetDataLabels()  (*DataLabels,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartPoint_GetDataLabels"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DataLabels{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDataLabels) 

	return result, nil 
}
// Gets or sets the Y value of the chart point.
// Returns:
//   Object  
func (instance *ChartPoint) Get_YValue()  (*Object,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartPoint_Get_YValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Object{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteObject) 

	return result, nil 
}
// Gets or sets the Y value of the chart point.
// Parameters:
//   value - Object 
// Returns:
//   void  
func (instance *ChartPoint) SetYValue(value *Object)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ChartPoint_SetYValue"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets Y value type of the chart point.
// Returns:
//   int32  
func (instance *ChartPoint) GetYValueType()  (CellValueType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ChartPoint_GetYValueType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToCellValueType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the X value of the chart point.
// Returns:
//   Object  
func (instance *ChartPoint) GetXValue()  (*Object,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartPoint_GetXValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Object{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteObject) 

	return result, nil 
}
// Gets or sets the X value of the chart point.
// Parameters:
//   value - Object 
// Returns:
//   void  
func (instance *ChartPoint) SetXValue(value *Object)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ChartPoint_SetXValue"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets X value type of the chart point.
// Returns:
//   int32  
func (instance *ChartPoint) GetXValueType()  (CellValueType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ChartPoint_GetXValueType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToCellValueType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the <see cref="ShapePropertyCollection"/> object that holds the visual shape properties of the ChartPoint.
// Returns:
//   ShapePropertyCollection  
func (instance *ChartPoint) GetShapeProperties()  (*ShapePropertyCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartPoint_GetShapeProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ShapePropertyCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteShapePropertyCollection) 

	return result, nil 
}
// Gets or sets a value indicates whether this data points is in the second pie or bar
// on a pie of pie or bar of pie chart
// Returns:
//   bool  
func (instance *ChartPoint) IsInSecondaryPlot()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartPoint_IsInSecondaryPlot"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets a value indicates whether this data points is in the second pie or bar
// on a pie of pie or bar of pie chart
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartPoint) SetIsInSecondaryPlot(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartPoint_SetIsInSecondaryPlot"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the x coordinate of the upper left corner in units of 1/4000 of chart's width after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *ChartPoint) GetShapeX()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartPoint_GetShapeX"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the y coordinate of the upper left corner in units of 1/4000 of chart's height after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *ChartPoint) GetShapeY()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartPoint_GetShapeY"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the width in units of 1/4000 of chart's width after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *ChartPoint) GetShapeWidth()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartPoint_GetShapeWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the height in units of 1/4000 of chart's height after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *ChartPoint) GetShapeHeight()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartPoint_GetShapeHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the x coordinate of the upper left corner in units of pixels after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *ChartPoint) GetShapeXPx()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartPoint_GetShapeXPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the y coordinate of the upper left corner in units of pixels after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *ChartPoint) GetShapeYPx()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartPoint_GetShapeYPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the width in units of pixels after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *ChartPoint) GetShapeWidthPx()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartPoint_GetShapeWidthPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the height in units of pixels after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *ChartPoint) GetShapeHeightPx()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartPoint_GetShapeHeightPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the width of border in units of pixels after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *ChartPoint) GetBorderWidthPx()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartPoint_GetBorderWidthPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the radius of bubble, pie or doughnut in units of pixels after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *ChartPoint) GetRadiusPx()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartPoint_GetRadiusPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the inner radius of doughnut slice in units of pixels after calls Chart.Calculate() method.
// Applies to Doughnut chart.
// Returns:
//   int32  
func (instance *ChartPoint) GetDoughnutInnerRadius()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartPoint_GetDoughnutInnerRadius"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the starting angle for the pie section, measured in degrees clockwise from the x-axis after calls Chart.Calculate() method.
// Applies to Pie chart.
// Returns:
//   float32  
func (instance *ChartPoint) GetStartAngle()  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAN(C.CString("ChartPoint_GetStartAngle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the ending angle for the pie section, measured in degrees clockwise from the x-axis after calls Chart.Calculate() method.
// Applies to Pie chart.
// Returns:
//   float32  
func (instance *ChartPoint) GetEndAngle()  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAN(C.CString("ChartPoint_GetEndAngle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the x coordinate of starting point for the pie section after calls Chart.Calculate() method.
// Applies to Pie and Doughnut  chart.
// Returns:
//   float32  
func (instance *ChartPoint) GetArcStartPointXPx()  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAN(C.CString("ChartPoint_GetArcStartPointXPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the y coordinate of starting point for the pie section after calls Chart.Calculate() method.
// Applies to Pie and Doughnut  chart.
// Returns:
//   float32  
func (instance *ChartPoint) GetArcStartPointYPx()  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAN(C.CString("ChartPoint_GetArcStartPointYPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the x coordinate of ending point for the pie section after calls Chart.Calculate() method.
// Applies to Pie and Doughnut  chart.
// Returns:
//   float32  
func (instance *ChartPoint) GetArcEndPointXPx()  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAN(C.CString("ChartPoint_GetArcEndPointXPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the y coordinate of ending point for the pie section after calls Chart.Calculate() method.
// Applies to Pie and Doughnut chart.
// Returns:
//   float32  
func (instance *ChartPoint) GetArcEndPointYPx()  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAN(C.CString("ChartPoint_GetArcEndPointYPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the x coordinate of starting point for the pie section after calls Chart.Calculate() method.
// Applies to Doughnut chart.
// Returns:
//   float32  
func (instance *ChartPoint) GetInnerArcStartPointXPx()  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAN(C.CString("ChartPoint_GetInnerArcStartPointXPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the y coordinate of starting point for the pie section after calls Chart.Calculate() method.
// Applies to Doughnut chart.
// Returns:
//   float32  
func (instance *ChartPoint) GetInnerArcStartPointYPx()  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAN(C.CString("ChartPoint_GetInnerArcStartPointYPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the x coordinate of ending point for the pie section after calls Chart.Calculate() method.
// Applies to Doughnut chart.
// Returns:
//   float32  
func (instance *ChartPoint) GetInnerArcEndPointXPx()  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAN(C.CString("ChartPoint_GetInnerArcEndPointXPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the y coordinate of ending point for the pie section after calls Chart.Calculate() method.
// Applies to Doughnut chart.
// Returns:
//   float32  
func (instance *ChartPoint) GetInnerArcEndPointYPx()  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAN(C.CString("ChartPoint_GetInnerArcEndPointYPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the number of top points after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *ChartPoint) GetTopPointCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartPoint_GetTopPointCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets x-coordinate of the top point of shape after calls Chart.Calculate() method.
// Applies 3D charts: Column3D, Bar3D, Cone, Cylinder, Pyramid and Area3D
// Parameters:
//   index - int32 
// Returns:
//   float32  
func (instance *ChartPoint) GetTopPointXPx(index int32)  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMX(C.CString("ChartPoint_GetTopPointXPx"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets y-coordinate of the top point of shape after calls Chart.Calculate() method.
// Applies 3D charts: Column3D, Bar3D, Cone, Cylinder, Pyramid and Area3D
// Parameters:
//   index - int32 
// Returns:
//   float32  
func (instance *ChartPoint) GetTopPointYPx(index int32)  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMX(C.CString("ChartPoint_GetTopPointYPx"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the number of bottom points  after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *ChartPoint) GetBottomPointCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartPoint_GetBottomPointCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets x-coordinate of the bottom point of shape after calls Chart.Calculate() method.
// Applies 3D charts: Column3D, Bar3D, Cone, Cylinder, Pyramid
// Parameters:
//   index - int32 
// Returns:
//   float32  
func (instance *ChartPoint) GetBottomPointXPx(index int32)  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMX(C.CString("ChartPoint_GetBottomPointXPx"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets y-coordinate of the bottom point of shape  after calls Chart.Calculate() method.
// Applies 3D charts: Column3D, Bar3D, Cone, Cylinder, Pyramid
// Parameters:
//   index - int32 
// Returns:
//   float32  
func (instance *ChartPoint) GetBottomPointYPx(index int32)  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMX(C.CString("ChartPoint_GetBottomPointYPx"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the number of the points on category axis after calls Chart.Calculate() method. Only applies to area chart.
// Returns:
//   int32  
func (instance *ChartPoint) GetOnCategoryAxisPointCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartPoint_GetOnCategoryAxisPointCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets x-coordinate of the point on category axis after calls Chart.Calculate() method. Only applies to Area chart.
// Parameters:
//   index - int32 
// Returns:
//   float32  
func (instance *ChartPoint) GetOnCategoryAxisPointXPx(index int32)  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMX(C.CString("ChartPoint_GetOnCategoryAxisPointXPx"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets y-coordinate of the point on category axis after calls Chart.Calculate() method. Only applies to Area chart.
// Parameters:
//   index - int32 
// Returns:
//   float32  
func (instance *ChartPoint) GetOnCategoryAxisPointYPx(index int32)  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMX(C.CString("ChartPoint_GetOnCategoryAxisPointYPx"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteChartPoint(chartpoint *ChartPoint){
	runtime.SetFinalizer(chartpoint, nil)
	C.Delete_CObject(C.CString("Delete_ChartPoint"),chartpoint.ptr)
	chartpoint.ptr = nil
}

// Class ChartPointCollection 

// Represents a collection that contains all the points in one series.
type ChartPointCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ChartPointCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartPointCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns an enumerator for the entire <see cref="ChartPointCollection"/>.
// Returns:
//   unsafe.Pointer  
func (instance *ChartPointCollection) GetEnumerator()  (*ChartPointEnumerator,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAW(C.CString("ChartPointCollection_GetEnumerator"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ChartPointEnumerator{}
	result.ptr = CGoReturnPtr.return_value
	runtime.SetFinalizer(result, DeleteChartPointEnumerator)
	 

	return result, nil 
}
// Remove all setting of the chart points.
// Returns:
//   void  
func (instance *ChartPointCollection) Clear()  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("ChartPointCollection_Clear"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes point at the index of the series..
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *ChartPointCollection) RemoveAt(index int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ChartPointCollection_RemoveAt"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the count of the chart point.
// Returns:
//   int32  
func (instance *ChartPointCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartPointCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="ChartPoint"/> element at the specified index in the series.
// Parameters:
//   index - int32 
// Returns:
//   ChartPoint  
func (instance *ChartPointCollection) Get(index int32)  (*ChartPoint,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("ChartPointCollection_Get"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ChartPoint{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteChartPoint) 

	return result, nil 
}



func DeleteChartPointCollection(chartpointcollection *ChartPointCollection){
	runtime.SetFinalizer(chartpointcollection, nil)
	C.Delete_CObject(C.CString("Delete_ChartPointCollection"),chartpointcollection.ptr)
	chartpointcollection.ptr = nil
}

// Class ChartTextFrame 

// Encapsulates the object that represents the frame object which contains text.
type ChartTextFrame struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ChartFrame 
func NewChartTextFrame(src *ChartFrame) ( *ChartTextFrame, error) {
	charttextframe := &ChartTextFrame{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_ChartTextFrame"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		charttextframe.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(charttextframe, DeleteChartTextFrame)
		return charttextframe, nil
	} else {
		charttextframe.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return charttextframe, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ChartTextFrame) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartTextFrame_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this data labels is deleted.
// Returns:
//   bool  
func (instance *ChartTextFrame) IsDeleted()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartTextFrame_IsDeleted"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this data labels is deleted.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartTextFrame) SetIsDeleted(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartTextFrame_SetIsDeleted"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the text horizontal alignment.
// Returns:
//   int32  
func (instance *ChartTextFrame) GetTextHorizontalAlignment()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ChartTextFrame_GetTextHorizontalAlignment"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the text horizontal alignment.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartTextFrame) SetTextHorizontalAlignment(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ChartTextFrame_SetTextHorizontalAlignment"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the text vertical alignment of text.
// Returns:
//   int32  
func (instance *ChartTextFrame) GetTextVerticalAlignment()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ChartTextFrame_GetTextVerticalAlignment"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the text vertical alignment of text.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartTextFrame) SetTextVerticalAlignment(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ChartTextFrame_SetTextVerticalAlignment"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents text rotation angle.
// Returns:
//   int32  
func (instance *ChartTextFrame) GetRotationAngle()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartTextFrame_GetRotationAngle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents text rotation angle.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartTextFrame) SetRotationAngle(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ChartTextFrame_SetRotationAngle"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the text of the chart is automatically rotated.
// Returns:
//   bool  
func (instance *ChartTextFrame) IsAutomaticRotation()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartTextFrame_IsAutomaticRotation"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns a Characters object that represents a range of characters within the text.
// Parameters:
//   startIndex - int32 
//   length - int32 
// Returns:
//   FontSetting  
func (instance *ChartTextFrame) Characters(startindex int32, length int32)  (*FontSetting,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBG(C.CString("ChartTextFrame_Characters"), instance.ptr, C.int(startindex), C.int(length))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &FontSetting{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFontSetting) 

	return result, nil 
}
// Represents text reading order.
// Returns:
//   int32  
func (instance *ChartTextFrame) GetReadingOrder()  (TextDirectionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ChartTextFrame_GetReadingOrder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextDirectionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents text reading order.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartTextFrame) SetReadingOrder(value TextDirectionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ChartTextFrame_SetReadingOrder"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets whether a shape should be auto-fit to fully contain the text described within it. Auto-fitting is
// when text within a shape is scaled in order to contain all the text inside.
// Returns:
//   bool  
func (instance *ChartTextFrame) IsResizeShapeToFitText()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartTextFrame_IsResizeShapeToFitText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets whether a shape should be auto-fit to fully contain the text described within it. Auto-fitting is
// when text within a shape is scaled in order to contain all the text inside.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartTextFrame) SetIsResizeShapeToFitText(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartTextFrame_SetIsResizeShapeToFitText"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates the text is auto generated.
// Returns:
//   bool  
func (instance *ChartTextFrame) IsAutoText()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartTextFrame_IsAutoText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates the text is auto generated.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartTextFrame) SetIsAutoText(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartTextFrame_SetIsAutoText"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the text of a frame's title.
// Returns:
//   string  
func (instance *ChartTextFrame) GetText()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ChartTextFrame_GetText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the text of a frame's title.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ChartTextFrame) SetText(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("ChartTextFrame_SetText"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets a reference to the worksheet.
// Returns:
//   string  
func (instance *ChartTextFrame) GetLinkedSource()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ChartTextFrame_GetLinkedSource"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets a reference to the worksheet.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ChartTextFrame) SetLinkedSource(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("ChartTextFrame_SetLinkedSource"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the direction of text.
// Returns:
//   int32  
func (instance *ChartTextFrame) GetDirectionType()  (ChartTextDirectionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ChartTextFrame_GetDirectionType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToChartTextDirectionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the direction of text.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartTextFrame) SetDirectionType(value ChartTextDirectionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ChartTextFrame_SetDirectionType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets a value indicating whether the text is wrapped.
// Returns:
//   bool  
func (instance *ChartTextFrame) IsTextWrapped()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartTextFrame_IsTextWrapped"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets a value indicating whether the text is wrapped.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartTextFrame) SetIsTextWrapped(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartTextFrame_SetIsTextWrapped"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.
// False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
// Returns:
//   bool  
func (instance *ChartTextFrame) IsInnerMode()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartTextFrame_IsInnerMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.
// False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartTextFrame) SetIsInnerMode(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartTextFrame_SetIsInnerMode"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the display mode of the background
// Returns:
//   int32  
func (instance *ChartTextFrame) GetBackgroundMode()  (BackgroundMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ChartTextFrame_GetBackgroundMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBackgroundMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the display mode of the background
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartTextFrame) SetBackgroundMode(value BackgroundMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ChartTextFrame_SetBackgroundMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the frame has a shadow.
// Returns:
//   bool  
func (instance *ChartTextFrame) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartTextFrame_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the frame has a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartTextFrame) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartTextFrame_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="ShapeProperties"/> object.
// Returns:
//   ShapePropertyCollection  
func (instance *ChartTextFrame) GetShapeProperties()  (*ShapePropertyCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartTextFrame_GetShapeProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ShapePropertyCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteShapePropertyCollection) 

	return result, nil 
}
// Indicates whether default position(DefaultX, DefaultY, DefaultWidth and DefaultHeight) are set.
// Returns:
//   bool  
func (instance *ChartTextFrame) IsDefaultPosBeSet()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartTextFrame_IsDefaultPosBeSet"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents x of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *ChartTextFrame) GetDefaultXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartTextFrame_GetDefaultXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents y of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *ChartTextFrame) GetDefaultYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartTextFrame_GetDefaultYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents width of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *ChartTextFrame) GetDefaultWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartTextFrame_GetDefaultWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents height of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *ChartTextFrame) GetDefaultHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartTextFrame_GetDefaultHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="Line">border</see>.
// Returns:
//   Line  
func (instance *ChartTextFrame) GetBorder()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartTextFrame_GetBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Gets the <see cref="Area">area</see>.
// Returns:
//   Area  
func (instance *ChartTextFrame) GetArea()  (*Area,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartTextFrame_GetArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Area{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteArea) 

	return result, nil 
}
// Gets and sets the options of the text.
// Returns:
//   TextOptions  
func (instance *ChartTextFrame) GetTextOptions()  (*TextOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartTextFrame_GetTextOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TextOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTextOptions) 

	return result, nil 
}
// Gets a <see cref="Font"/> object of the specified ChartFrame object.
// Returns:
//   Font  
func (instance *ChartTextFrame) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ChartTextFrame_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Returns:
//   bool  
func (instance *ChartTextFrame) GetAutoScaleFont()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartTextFrame_GetAutoScaleFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartTextFrame) SetAutoScaleFont(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartTextFrame_SetAutoScaleFont"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the chart frame is automatic sized.
// Returns:
//   bool  
func (instance *ChartTextFrame) IsAutomaticSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ChartTextFrame_IsAutomaticSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the chart frame is automatic sized.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ChartTextFrame) SetIsAutomaticSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ChartTextFrame_SetIsAutomaticSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the x coordinate of the upper left corner in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *ChartTextFrame) GetXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartTextFrame_GetXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the x coordinate of the upper left corner in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ChartTextFrame) SetXRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ChartTextFrame_SetXRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the y coordinate of the upper left corner in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *ChartTextFrame) GetYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartTextFrame_GetYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the y coordinate of the upper left corner in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ChartTextFrame) SetYRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ChartTextFrame_SetYRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of frame in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *ChartTextFrame) GetWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartTextFrame_GetWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of frame in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ChartTextFrame) SetWidthRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ChartTextFrame_SetWidthRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the height of frame in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *ChartTextFrame) GetHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ChartTextFrame_GetHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the height of frame in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ChartTextFrame) SetHeightRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ChartTextFrame_SetHeightRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Pixel.
// Returns:
//   int32  
func (instance *ChartTextFrame) GetXPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartTextFrame_GetXPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartTextFrame) SetXPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ChartTextFrame_SetXPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Pixel.
// Returns:
//   int32  
func (instance *ChartTextFrame) GetYPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartTextFrame_GetYPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartTextFrame) SetYPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ChartTextFrame_SetYPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of frame in units of Pixel.
// Returns:
//   int32  
func (instance *ChartTextFrame) GetWidthPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartTextFrame_GetWidthPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of frame in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartTextFrame) SetWidthPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ChartTextFrame_SetWidthPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the height of frame in units of Pixel.
// Returns:
//   int32  
func (instance *ChartTextFrame) GetHeightPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ChartTextFrame_GetHeightPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the height of frame in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ChartTextFrame) SetHeightPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ChartTextFrame_SetHeightPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Set position of the frame to automatic
// Returns:
//   void  
func (instance *ChartTextFrame) SetPositionAuto()  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("ChartTextFrame_SetPositionAuto"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *ChartTextFrame) ToChartFrame() *ChartFrame {
	parentClass := &ChartFrame{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteChartTextFrame(charttextframe *ChartTextFrame){
	runtime.SetFinalizer(charttextframe, nil)
	C.Delete_CObject(C.CString("Delete_ChartTextFrame"),charttextframe.ptr)
	charttextframe.ptr = nil
}

// Class DataLabels 

// Encapsulates a collection of all the DataLabel objects for the specified Series.
type DataLabels struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ChartTextFrame 
func NewDataLabels(src *ChartTextFrame) ( *DataLabels, error) {
	datalabels := &DataLabels{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_DataLabels"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		datalabels.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(datalabels, DeleteDataLabels)
		return datalabels, nil
	} else {
		datalabels.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return datalabels, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DataLabels) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="Line">border</see>.
// Returns:
//   Line  
func (instance *DataLabels) GetBorder()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DataLabels_GetBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Gets the <see cref="Area">area</see>.
// Returns:
//   Area  
func (instance *DataLabels) GetArea()  (*Area,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DataLabels_GetArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Area{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteArea) 

	return result, nil 
}
// Indicates the text is auto generated.
// Returns:
//   bool  
func (instance *DataLabels) IsAutoText()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_IsAutoText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates the text is auto generated.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetIsAutoText(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetIsAutoText"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the direction of text.
// Returns:
//   int32  
func (instance *DataLabels) GetDirectionType()  (ChartTextDirectionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DataLabels_GetDirectionType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToChartTextDirectionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the direction of text.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataLabels) SetDirectionType(value ChartTextDirectionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DataLabels_SetDirectionType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the text of data label.
// Returns:
//   string  
func (instance *DataLabels) GetText()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataLabels_GetText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the text of data label.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DataLabels) SetText(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DataLabels_SetText"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets a value indicating whether the text is wrapped.
// Returns:
//   bool  
func (instance *DataLabels) IsTextWrapped()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_IsTextWrapped"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets a value indicating whether the text is wrapped.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetIsTextWrapped(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetIsTextWrapped"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the display mode of the background
// Returns:
//   int32  
func (instance *DataLabels) GetBackgroundMode()  (BackgroundMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DataLabels_GetBackgroundMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBackgroundMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the display mode of the background
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataLabels) SetBackgroundMode(value BackgroundMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DataLabels_SetBackgroundMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents a specified chart's data label values display behavior. True displays the values. False to hide.
// Returns:
//   bool  
func (instance *DataLabels) GetShowValue()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_GetShowValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents a specified chart's data label values display behavior. True displays the values. False to hide.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetShowValue(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetShowValue"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether showing cell range as the data labels.
// Returns:
//   bool  
func (instance *DataLabels) GetShowCellRange()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_GetShowCellRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether showing cell range as the data labels.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetShowCellRange(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetShowCellRange"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents a specified chart's data label percentage value display behavior. True displays the percentage value. False to hide.
// Returns:
//   bool  
func (instance *DataLabels) GetShowPercentage()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_GetShowPercentage"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents a specified chart's data label percentage value display behavior. True displays the percentage value. False to hide.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetShowPercentage(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetShowPercentage"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents a specified chart's data label percentage value display behavior. True displays the percentage value. False to hide.
// Returns:
//   bool  
func (instance *DataLabels) GetShowBubbleSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_GetShowBubbleSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents a specified chart's data label percentage value display behavior. True displays the percentage value. False to hide.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetShowBubbleSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetShowBubbleSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents a specified chart's data label category name display behavior.True to display the category name for the data labels on a chart. False to hide.
// Returns:
//   bool  
func (instance *DataLabels) GetShowCategoryName()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_GetShowCategoryName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents a specified chart's data label category name display behavior.True to display the category name for the data labels on a chart. False to hide.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetShowCategoryName(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetShowCategoryName"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the series name displays for the data labels on a chart.
// True to show the series name. False to hide.
// Returns:
//   bool  
func (instance *DataLabels) GetShowSeriesName()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_GetShowSeriesName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the series name displays for the data labels on a chart.
// True to show the series name. False to hide.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetShowSeriesName(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetShowSeriesName"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents a specified chart's data label legend key display behavior.
// True if the data label legend key is visible.
// Returns:
//   bool  
func (instance *DataLabels) GetShowLegendKey()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_GetShowLegendKey"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents a specified chart's data label legend key display behavior.
// True if the data label legend key is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetShowLegendKey(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetShowLegendKey"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the format string for the DataLabels object.
// Returns:
//   string  
func (instance *DataLabels) Get_NumberFormat()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataLabels_Get_NumberFormat"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the format string for the DataLabels object.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DataLabels) SetNumberFormat(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DataLabels_SetNumberFormat"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the built-in number format.
// Returns:
//   int32  
func (instance *DataLabels) GetNumber()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DataLabels_GetNumber"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the built-in number format.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataLabels) SetNumber(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("DataLabels_SetNumber"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the number format is linked to the cells
// (so that the number format changes in the labels when it changes in the cells).
// Returns:
//   bool  
func (instance *DataLabels) GetNumberFormatLinked()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_GetNumberFormatLinked"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the number format is linked to the cells
// (so that the number format changes in the labels when it changes in the cells).
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetNumberFormatLinked(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetNumberFormatLinked"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Apply the font of the datalabels to all child nodes.
// Returns:
//   void  
func (instance *DataLabels) ApplyFont()  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("DataLabels_ApplyFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the font of the DataLabels;
// Returns:
//   Font  
func (instance *DataLabels) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DataLabels_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Gets or sets the separator type used for the data labels on a chart.
// Returns:
//   int32  
func (instance *DataLabels) GetSeparatorType()  (DataLabelsSeparatorType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DataLabels_GetSeparatorType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToDataLabelsSeparatorType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the separator type used for the data labels on a chart.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataLabels) SetSeparatorType(value DataLabelsSeparatorType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DataLabels_SetSeparatorType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the separator value used for the data labels on a chart.
// Returns:
//   string  
func (instance *DataLabels) GetSeparatorValue()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataLabels_GetSeparatorValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the separator value used for the data labels on a chart.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DataLabels) SetSeparatorValue(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DataLabels_SetSeparatorValue"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the position of the data label.
// Returns:
//   int32  
func (instance *DataLabels) GetPosition()  (LabelPositionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DataLabels_GetPosition"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToLabelPositionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the position of the data label.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataLabels) SetPosition(value LabelPositionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DataLabels_SetPosition"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the datalabels display never overlap. (For Pie chart)
// Returns:
//   bool  
func (instance *DataLabels) IsNeverOverlap()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_IsNeverOverlap"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the datalabels display never overlap. (For Pie chart)
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetIsNeverOverlap(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetIsNeverOverlap"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets  shape type of data label.
// Returns:
//   int32  
func (instance *DataLabels) GetShapeType()  (DataLabelShapeType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DataLabels_GetShapeType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToDataLabelShapeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets  shape type of data label.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataLabels) SetShapeType(value DataLabelShapeType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DataLabels_SetShapeType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.
// False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
// Returns:
//   bool  
func (instance *DataLabels) IsInnerMode()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_IsInnerMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.
// False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetIsInnerMode(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetIsInnerMode"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the frame has a shadow.
// Returns:
//   bool  
func (instance *DataLabels) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the frame has a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="ShapeProperties"/> object.
// Returns:
//   ShapePropertyCollection  
func (instance *DataLabels) GetShapeProperties()  (*ShapePropertyCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DataLabels_GetShapeProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ShapePropertyCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteShapePropertyCollection) 

	return result, nil 
}
// Indicates whether default position(DefaultX, DefaultY, DefaultWidth and DefaultHeight) are set.
// Returns:
//   bool  
func (instance *DataLabels) IsDefaultPosBeSet()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_IsDefaultPosBeSet"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents x of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *DataLabels) GetDefaultXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("DataLabels_GetDefaultXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents y of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *DataLabels) GetDefaultYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("DataLabels_GetDefaultYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents width of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *DataLabels) GetDefaultWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("DataLabels_GetDefaultWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents height of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *DataLabels) GetDefaultHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("DataLabels_GetDefaultHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the options of the text.
// Returns:
//   TextOptions  
func (instance *DataLabels) GetTextOptions()  (*TextOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DataLabels_GetTextOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TextOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTextOptions) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Returns:
//   bool  
func (instance *DataLabels) GetAutoScaleFont()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_GetAutoScaleFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetAutoScaleFont(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetAutoScaleFont"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the chart frame is automatic sized.
// Returns:
//   bool  
func (instance *DataLabels) IsAutomaticSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_IsAutomaticSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the chart frame is automatic sized.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetIsAutomaticSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetIsAutomaticSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the x coordinate of the upper left corner in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *DataLabels) GetXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("DataLabels_GetXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the x coordinate of the upper left corner in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *DataLabels) SetXRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("DataLabels_SetXRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the y coordinate of the upper left corner in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *DataLabels) GetYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("DataLabels_GetYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the y coordinate of the upper left corner in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *DataLabels) SetYRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("DataLabels_SetYRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of frame in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *DataLabels) GetWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("DataLabels_GetWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of frame in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *DataLabels) SetWidthRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("DataLabels_SetWidthRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the height of frame in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *DataLabels) GetHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("DataLabels_GetHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the height of frame in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *DataLabels) SetHeightRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("DataLabels_SetHeightRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Pixel.
// Returns:
//   int32  
func (instance *DataLabels) GetXPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DataLabels_GetXPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataLabels) SetXPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("DataLabels_SetXPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Pixel.
// Returns:
//   int32  
func (instance *DataLabels) GetYPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DataLabels_GetYPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataLabels) SetYPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("DataLabels_SetYPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of frame in units of Pixel.
// Returns:
//   int32  
func (instance *DataLabels) GetWidthPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DataLabels_GetWidthPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of frame in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataLabels) SetWidthPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("DataLabels_SetWidthPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the height of frame in units of Pixel.
// Returns:
//   int32  
func (instance *DataLabels) GetHeightPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DataLabels_GetHeightPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the height of frame in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataLabels) SetHeightPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("DataLabels_SetHeightPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Set position of the frame to automatic
// Returns:
//   void  
func (instance *DataLabels) SetPositionAuto()  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("DataLabels_SetPositionAuto"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this data labels is deleted.
// Returns:
//   bool  
func (instance *DataLabels) IsDeleted()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_IsDeleted"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this data labels is deleted.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetIsDeleted(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetIsDeleted"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the text horizontal alignment.
// Returns:
//   int32  
func (instance *DataLabels) GetTextHorizontalAlignment()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DataLabels_GetTextHorizontalAlignment"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the text horizontal alignment.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataLabels) SetTextHorizontalAlignment(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DataLabels_SetTextHorizontalAlignment"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the text vertical alignment of text.
// Returns:
//   int32  
func (instance *DataLabels) GetTextVerticalAlignment()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DataLabels_GetTextVerticalAlignment"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the text vertical alignment of text.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataLabels) SetTextVerticalAlignment(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DataLabels_SetTextVerticalAlignment"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents text rotation angle.
// Returns:
//   int32  
func (instance *DataLabels) GetRotationAngle()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DataLabels_GetRotationAngle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents text rotation angle.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataLabels) SetRotationAngle(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("DataLabels_SetRotationAngle"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the text of the chart is automatically rotated.
// Returns:
//   bool  
func (instance *DataLabels) IsAutomaticRotation()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_IsAutomaticRotation"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns a Characters object that represents a range of characters within the text.
// Parameters:
//   startIndex - int32 
//   length - int32 
// Returns:
//   FontSetting  
func (instance *DataLabels) Characters(startindex int32, length int32)  (*FontSetting,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBG(C.CString("DataLabels_Characters"), instance.ptr, C.int(startindex), C.int(length))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &FontSetting{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFontSetting) 

	return result, nil 
}
// Represents text reading order.
// Returns:
//   int32  
func (instance *DataLabels) GetReadingOrder()  (TextDirectionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DataLabels_GetReadingOrder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextDirectionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents text reading order.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataLabels) SetReadingOrder(value TextDirectionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DataLabels_SetReadingOrder"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets whether a shape should be auto-fit to fully contain the text described within it. Auto-fitting is
// when text within a shape is scaled in order to contain all the text inside.
// Returns:
//   bool  
func (instance *DataLabels) IsResizeShapeToFitText()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataLabels_IsResizeShapeToFitText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets whether a shape should be auto-fit to fully contain the text described within it. Auto-fitting is
// when text within a shape is scaled in order to contain all the text inside.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataLabels) SetIsResizeShapeToFitText(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataLabels_SetIsResizeShapeToFitText"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets a reference to the worksheet.
// Returns:
//   string  
func (instance *DataLabels) GetLinkedSource()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataLabels_GetLinkedSource"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets a reference to the worksheet.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DataLabels) SetLinkedSource(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DataLabels_SetLinkedSource"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *DataLabels) ToChartTextFrame() *ChartTextFrame {
	parentClass := &ChartTextFrame{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *DataLabels) ToChartFrame() *ChartFrame {
	parentClass := &ChartFrame{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteDataLabels(datalabels *DataLabels){
	runtime.SetFinalizer(datalabels, nil)
	C.Delete_CObject(C.CString("Delete_DataLabels"),datalabels.ptr)
	datalabels.ptr = nil
}

// Class DisplayUnitLabel 

// Represents the display unit label.
type DisplayUnitLabel struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ChartTextFrame 
func NewDisplayUnitLabel(src *ChartTextFrame) ( *DisplayUnitLabel, error) {
	displayunitlabel := &DisplayUnitLabel{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_DisplayUnitLabel"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		displayunitlabel.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(displayunitlabel, DeleteDisplayUnitLabel)
		return displayunitlabel, nil
	} else {
		displayunitlabel.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return displayunitlabel, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DisplayUnitLabel) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DisplayUnitLabel_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the text of display unit label.
// Returns:
//   string  
func (instance *DisplayUnitLabel) GetText()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DisplayUnitLabel_GetText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the text of display unit label.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetText(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DisplayUnitLabel_SetText"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets a <see cref="Font"/> object of the specified ChartFrame object.
// Returns:
//   Font  
func (instance *DisplayUnitLabel) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DisplayUnitLabel_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Returns:
//   bool  
func (instance *DisplayUnitLabel) GetAutoScaleFont()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DisplayUnitLabel_GetAutoScaleFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetAutoScaleFont(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DisplayUnitLabel_SetAutoScaleFont"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.
// False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
// Returns:
//   bool  
func (instance *DisplayUnitLabel) IsInnerMode()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DisplayUnitLabel_IsInnerMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.
// False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetIsInnerMode(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DisplayUnitLabel_SetIsInnerMode"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the display mode of the background
// Returns:
//   int32  
func (instance *DisplayUnitLabel) GetBackgroundMode()  (BackgroundMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DisplayUnitLabel_GetBackgroundMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBackgroundMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the display mode of the background
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetBackgroundMode(value BackgroundMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DisplayUnitLabel_SetBackgroundMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the frame has a shadow.
// Returns:
//   bool  
func (instance *DisplayUnitLabel) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DisplayUnitLabel_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the frame has a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DisplayUnitLabel_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="ShapeProperties"/> object.
// Returns:
//   ShapePropertyCollection  
func (instance *DisplayUnitLabel) GetShapeProperties()  (*ShapePropertyCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DisplayUnitLabel_GetShapeProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ShapePropertyCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteShapePropertyCollection) 

	return result, nil 
}
// Indicates whether default position(DefaultX, DefaultY, DefaultWidth and DefaultHeight) are set.
// Returns:
//   bool  
func (instance *DisplayUnitLabel) IsDefaultPosBeSet()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DisplayUnitLabel_IsDefaultPosBeSet"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents x of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *DisplayUnitLabel) GetDefaultXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("DisplayUnitLabel_GetDefaultXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents y of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *DisplayUnitLabel) GetDefaultYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("DisplayUnitLabel_GetDefaultYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents width of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *DisplayUnitLabel) GetDefaultWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("DisplayUnitLabel_GetDefaultWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents height of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *DisplayUnitLabel) GetDefaultHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("DisplayUnitLabel_GetDefaultHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="Line">border</see>.
// Returns:
//   Line  
func (instance *DisplayUnitLabel) GetBorder()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DisplayUnitLabel_GetBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Gets the <see cref="Area">area</see>.
// Returns:
//   Area  
func (instance *DisplayUnitLabel) GetArea()  (*Area,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DisplayUnitLabel_GetArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Area{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteArea) 

	return result, nil 
}
// Gets and sets the options of the text.
// Returns:
//   TextOptions  
func (instance *DisplayUnitLabel) GetTextOptions()  (*TextOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DisplayUnitLabel_GetTextOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TextOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTextOptions) 

	return result, nil 
}
// Indicates whether the chart frame is automatic sized.
// Returns:
//   bool  
func (instance *DisplayUnitLabel) IsAutomaticSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DisplayUnitLabel_IsAutomaticSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the chart frame is automatic sized.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetIsAutomaticSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DisplayUnitLabel_SetIsAutomaticSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the x coordinate of the upper left corner in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *DisplayUnitLabel) GetXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("DisplayUnitLabel_GetXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the x coordinate of the upper left corner in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetXRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("DisplayUnitLabel_SetXRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the y coordinate of the upper left corner in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *DisplayUnitLabel) GetYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("DisplayUnitLabel_GetYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the y coordinate of the upper left corner in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetYRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("DisplayUnitLabel_SetYRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of frame in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *DisplayUnitLabel) GetWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("DisplayUnitLabel_GetWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of frame in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetWidthRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("DisplayUnitLabel_SetWidthRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the height of frame in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *DisplayUnitLabel) GetHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("DisplayUnitLabel_GetHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the height of frame in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetHeightRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("DisplayUnitLabel_SetHeightRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Pixel.
// Returns:
//   int32  
func (instance *DisplayUnitLabel) GetXPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DisplayUnitLabel_GetXPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetXPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("DisplayUnitLabel_SetXPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Pixel.
// Returns:
//   int32  
func (instance *DisplayUnitLabel) GetYPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DisplayUnitLabel_GetYPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetYPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("DisplayUnitLabel_SetYPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of frame in units of Pixel.
// Returns:
//   int32  
func (instance *DisplayUnitLabel) GetWidthPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DisplayUnitLabel_GetWidthPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of frame in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetWidthPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("DisplayUnitLabel_SetWidthPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the height of frame in units of Pixel.
// Returns:
//   int32  
func (instance *DisplayUnitLabel) GetHeightPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DisplayUnitLabel_GetHeightPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the height of frame in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetHeightPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("DisplayUnitLabel_SetHeightPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Set position of the frame to automatic
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetPositionAuto()  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("DisplayUnitLabel_SetPositionAuto"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this data labels is deleted.
// Returns:
//   bool  
func (instance *DisplayUnitLabel) IsDeleted()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DisplayUnitLabel_IsDeleted"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this data labels is deleted.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetIsDeleted(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DisplayUnitLabel_SetIsDeleted"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the text horizontal alignment.
// Returns:
//   int32  
func (instance *DisplayUnitLabel) GetTextHorizontalAlignment()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DisplayUnitLabel_GetTextHorizontalAlignment"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the text horizontal alignment.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetTextHorizontalAlignment(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DisplayUnitLabel_SetTextHorizontalAlignment"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the text vertical alignment of text.
// Returns:
//   int32  
func (instance *DisplayUnitLabel) GetTextVerticalAlignment()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DisplayUnitLabel_GetTextVerticalAlignment"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the text vertical alignment of text.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetTextVerticalAlignment(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DisplayUnitLabel_SetTextVerticalAlignment"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents text rotation angle.
// Returns:
//   int32  
func (instance *DisplayUnitLabel) GetRotationAngle()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DisplayUnitLabel_GetRotationAngle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents text rotation angle.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetRotationAngle(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("DisplayUnitLabel_SetRotationAngle"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the text of the chart is automatically rotated.
// Returns:
//   bool  
func (instance *DisplayUnitLabel) IsAutomaticRotation()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DisplayUnitLabel_IsAutomaticRotation"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns a Characters object that represents a range of characters within the text.
// Parameters:
//   startIndex - int32 
//   length - int32 
// Returns:
//   FontSetting  
func (instance *DisplayUnitLabel) Characters(startindex int32, length int32)  (*FontSetting,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBG(C.CString("DisplayUnitLabel_Characters"), instance.ptr, C.int(startindex), C.int(length))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &FontSetting{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFontSetting) 

	return result, nil 
}
// Represents text reading order.
// Returns:
//   int32  
func (instance *DisplayUnitLabel) GetReadingOrder()  (TextDirectionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DisplayUnitLabel_GetReadingOrder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextDirectionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents text reading order.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetReadingOrder(value TextDirectionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DisplayUnitLabel_SetReadingOrder"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets whether a shape should be auto-fit to fully contain the text described within it. Auto-fitting is
// when text within a shape is scaled in order to contain all the text inside.
// Returns:
//   bool  
func (instance *DisplayUnitLabel) IsResizeShapeToFitText()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DisplayUnitLabel_IsResizeShapeToFitText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets whether a shape should be auto-fit to fully contain the text described within it. Auto-fitting is
// when text within a shape is scaled in order to contain all the text inside.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetIsResizeShapeToFitText(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DisplayUnitLabel_SetIsResizeShapeToFitText"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates the text is auto generated.
// Returns:
//   bool  
func (instance *DisplayUnitLabel) IsAutoText()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DisplayUnitLabel_IsAutoText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates the text is auto generated.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetIsAutoText(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DisplayUnitLabel_SetIsAutoText"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets a reference to the worksheet.
// Returns:
//   string  
func (instance *DisplayUnitLabel) GetLinkedSource()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DisplayUnitLabel_GetLinkedSource"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets a reference to the worksheet.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetLinkedSource(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DisplayUnitLabel_SetLinkedSource"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the direction of text.
// Returns:
//   int32  
func (instance *DisplayUnitLabel) GetDirectionType()  (ChartTextDirectionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DisplayUnitLabel_GetDirectionType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToChartTextDirectionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the direction of text.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetDirectionType(value ChartTextDirectionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DisplayUnitLabel_SetDirectionType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets a value indicating whether the text is wrapped.
// Returns:
//   bool  
func (instance *DisplayUnitLabel) IsTextWrapped()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DisplayUnitLabel_IsTextWrapped"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets a value indicating whether the text is wrapped.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DisplayUnitLabel) SetIsTextWrapped(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DisplayUnitLabel_SetIsTextWrapped"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *DisplayUnitLabel) ToChartTextFrame() *ChartTextFrame {
	parentClass := &ChartTextFrame{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *DisplayUnitLabel) ToChartFrame() *ChartFrame {
	parentClass := &ChartFrame{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteDisplayUnitLabel(displayunitlabel *DisplayUnitLabel){
	runtime.SetFinalizer(displayunitlabel, nil)
	C.Delete_CObject(C.CString("Delete_DisplayUnitLabel"),displayunitlabel.ptr)
	displayunitlabel.ptr = nil
}

// Class DropBars 

// Represents the up/down bars in a chart.
type DropBars struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DropBars) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DropBars_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the border <see cref="Line"/>.
// Returns:
//   Line  
func (instance *DropBars) GetBorder()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DropBars_GetBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Gets the <see cref="Area"/>.
// Returns:
//   Area  
func (instance *DropBars) GetArea()  (*Area,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DropBars_GetArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Area{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteArea) 

	return result, nil 
}



func DeleteDropBars(dropbars *DropBars){
	runtime.SetFinalizer(dropbars, nil)
	C.Delete_CObject(C.CString("Delete_DropBars"),dropbars.ptr)
	dropbars.ptr = nil
}

// Class ErrorBar 

// Represents error bar of data series.
type ErrorBar struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Line 
func NewErrorBar(src *Line) ( *ErrorBar, error) {
	errorbar := &ErrorBar{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_ErrorBar"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		errorbar.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(errorbar, DeleteErrorBar)
		return errorbar, nil
	} else {
		errorbar.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return errorbar, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ErrorBar) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ErrorBar_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents error bar amount type.
// Returns:
//   int32  
func (instance *ErrorBar) GetType()  (ErrorBarType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ErrorBar_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToErrorBarType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents error bar amount type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ErrorBar) SetType(value ErrorBarType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ErrorBar_SetType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the display type of error bar.
// Returns:
//   int32  
func (instance *ErrorBar) GetDisplayType()  (ErrorBarDisplayType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ErrorBar_GetDisplayType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToErrorBarDisplayType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the display type of error bar.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ErrorBar) SetDisplayType(value ErrorBarDisplayType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ErrorBar_SetDisplayType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents amount of error bar.
// Returns:
//   float64  
func (instance *ErrorBar) GetAmount()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ErrorBar_GetAmount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents amount of error bar.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ErrorBar) SetAmount(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ErrorBar_SetAmount"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates if formatting error bars with a T-top.
// Returns:
//   bool  
func (instance *ErrorBar) GetShowMarkerTTop()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ErrorBar_GetShowMarkerTTop"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates if formatting error bars with a T-top.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ErrorBar) SetShowMarkerTTop(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ErrorBar_SetShowMarkerTTop"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents positive error amount when error bar type is Custom.
// Returns:
//   string  
func (instance *ErrorBar) GetPlusValue()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ErrorBar_GetPlusValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents positive error amount when error bar type is Custom.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ErrorBar) SetPlusValue(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("ErrorBar_SetPlusValue"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents negative error amount when error bar type is Custom.
// Returns:
//   string  
func (instance *ErrorBar) GetMinusValue()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ErrorBar_GetMinusValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents negative error amount when error bar type is Custom.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ErrorBar) SetMinusValue(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("ErrorBar_SetMinusValue"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the compound line type
// Returns:
//   int32  
func (instance *ErrorBar) GetCompoundType()  (MsoLineStyle,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ErrorBar_GetCompoundType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMsoLineStyle(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the compound line type
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ErrorBar) SetCompoundType(value MsoLineStyle)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ErrorBar_SetCompoundType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the dash line type
// Returns:
//   int32  
func (instance *ErrorBar) GetDashType()  (MsoLineDashStyle,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ErrorBar_GetDashType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMsoLineDashStyle(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the dash line type
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ErrorBar) SetDashType(value MsoLineDashStyle)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ErrorBar_SetDashType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the ending caps.
// Returns:
//   int32  
func (instance *ErrorBar) GetCapType()  (LineCapType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ErrorBar_GetCapType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToLineCapType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the ending caps.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ErrorBar) SetCapType(value LineCapType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ErrorBar_SetCapType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the joining caps.
// Returns:
//   int32  
func (instance *ErrorBar) GetJoinType()  (LineJoinType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ErrorBar_GetJoinType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToLineJoinType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the joining caps.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ErrorBar) SetJoinType(value LineJoinType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ErrorBar_SetJoinType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies an arrowhead for the begin of a line.
// Returns:
//   int32  
func (instance *ErrorBar) GetBeginType()  (MsoArrowheadStyle,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ErrorBar_GetBeginType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMsoArrowheadStyle(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies an arrowhead for the begin of a line.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ErrorBar) SetBeginType(value MsoArrowheadStyle)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ErrorBar_SetBeginType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies an arrowhead for the end of a line.
// Returns:
//   int32  
func (instance *ErrorBar) GetEndType()  (MsoArrowheadStyle,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ErrorBar_GetEndType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMsoArrowheadStyle(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies an arrowhead for the end of a line.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ErrorBar) SetEndType(value MsoArrowheadStyle)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ErrorBar_SetEndType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the length of the arrowhead for the begin of a line.
// Returns:
//   int32  
func (instance *ErrorBar) GetBeginArrowLength()  (MsoArrowheadLength,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ErrorBar_GetBeginArrowLength"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMsoArrowheadLength(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the length of the arrowhead for the begin of a line.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ErrorBar) SetBeginArrowLength(value MsoArrowheadLength)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ErrorBar_SetBeginArrowLength"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the length of the arrowhead for the end of a line.
// Returns:
//   int32  
func (instance *ErrorBar) GetEndArrowLength()  (MsoArrowheadLength,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ErrorBar_GetEndArrowLength"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMsoArrowheadLength(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the length of the arrowhead for the end of a line.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ErrorBar) SetEndArrowLength(value MsoArrowheadLength)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ErrorBar_SetEndArrowLength"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the width of the arrowhead for the begin of a line.
// Returns:
//   int32  
func (instance *ErrorBar) GetBeginArrowWidth()  (MsoArrowheadWidth,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ErrorBar_GetBeginArrowWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMsoArrowheadWidth(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the width of the arrowhead for the begin of a line.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ErrorBar) SetBeginArrowWidth(value MsoArrowheadWidth)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ErrorBar_SetBeginArrowWidth"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the width of the arrowhead for the end of a line.
// Returns:
//   int32  
func (instance *ErrorBar) GetEndArrowWidth()  (MsoArrowheadWidth,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ErrorBar_GetEndArrowWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMsoArrowheadWidth(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the width of the arrowhead for the end of a line.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ErrorBar) SetEndArrowWidth(value MsoArrowheadWidth)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ErrorBar_SetEndArrowWidth"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the theme color.
// Returns:
//   ThemeColor  
func (instance *ErrorBar) GetThemeColor()  (*ThemeColor,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ErrorBar_GetThemeColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ThemeColor{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteThemeColor) 

	return result, nil 
}
// Gets and sets the theme color.
// Parameters:
//   value - ThemeColor 
// Returns:
//   void  
func (instance *ErrorBar) SetThemeColor(value *ThemeColor)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ErrorBar_SetThemeColor"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the <see cref="Color"/> of the line.
// Returns:
//   Color  
func (instance *ErrorBar) GetColor()  (*Color,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAD(C.CString("ErrorBar_GetColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Color{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Represents the <see cref="Color"/> of the line.
// Parameters:
//   value - Color 
// Returns:
//   void  
func (instance *ErrorBar) SetColor(value *Color)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAE(C.CString("ErrorBar_SetColor"), instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the degree of transparency of the line as a value from 0.0 (opaque) through 1.0 (clear).
// Returns:
//   float64  
func (instance *ErrorBar) GetTransparency()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ErrorBar_GetTransparency"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the degree of transparency of the line as a value from 0.0 (opaque) through 1.0 (clear).
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ErrorBar) SetTransparency(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ErrorBar_SetTransparency"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the style of the line.
// Returns:
//   int32  
func (instance *ErrorBar) GetStyle()  (LineType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ErrorBar_GetStyle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToLineType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the style of the line.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ErrorBar) SetStyle(value LineType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ErrorBar_SetStyle"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the <see cref="WeightType"/> of the line.
// Returns:
//   int32  
func (instance *ErrorBar) GetWeight()  (WeightType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ErrorBar_GetWeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToWeightType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the <see cref="WeightType"/> of the line.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ErrorBar) SetWeight(value WeightType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ErrorBar_SetWeight"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the weight of the line in unit of points.
// Returns:
//   float64  
func (instance *ErrorBar) GetWeightPt()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ErrorBar_GetWeightPt"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the weight of the line in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ErrorBar) SetWeightPt(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ErrorBar_SetWeightPt"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the weight of the line in unit of pixels.
// Returns:
//   float64  
func (instance *ErrorBar) GetWeightPx()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ErrorBar_GetWeightPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the weight of the line in unit of pixels.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ErrorBar) SetWeightPx(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ErrorBar_SetWeightPx"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets format type.
// Returns:
//   int32  
func (instance *ErrorBar) GetFormattingType()  (ChartLineFormattingType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ErrorBar_GetFormattingType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToChartLineFormattingType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets format type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ErrorBar) SetFormattingType(value ChartLineFormattingType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ErrorBar_SetFormattingType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the color of line is automatic assigned.
// Returns:
//   bool  
func (instance *ErrorBar) IsAutomaticColor()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ErrorBar_IsAutomaticColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents whether the line is visible.
// Returns:
//   bool  
func (instance *ErrorBar) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ErrorBar_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents whether the line is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ErrorBar) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ErrorBar_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this line style is auto assigned.
// Returns:
//   bool  
func (instance *ErrorBar) IsAuto()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ErrorBar_IsAuto"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this line style is auto assigned.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ErrorBar) SetIsAuto(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ErrorBar_SetIsAuto"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents gradient fill.
// Returns:
//   GradientFill  
func (instance *ErrorBar) GetGradientFill()  (*GradientFill,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ErrorBar_GetGradientFill"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &GradientFill{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteGradientFill) 

	return result, nil 
}


func (instance *ErrorBar) ToLine() *Line {
	parentClass := &Line{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteErrorBar(errorbar *ErrorBar){
	runtime.SetFinalizer(errorbar, nil)
	C.Delete_CObject(C.CString("Delete_ErrorBar"),errorbar.ptr)
	errorbar.ptr = nil
}

// Class Floor 

// Encapsulates the object that represents the floor of a 3-D chart.
type Floor struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Area 
func NewFloor(src *Area) ( *Floor, error) {
	floor := &Floor{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_Floor"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		floor.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(floor, DeleteFloor)
		return floor, nil
	} else {
		floor.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return floor, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *Floor) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Floor_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the border <see cref="Line"/>.
// Returns:
//   Line  
func (instance *Floor) GetBorder()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Floor_GetBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Gets or sets the border <see cref="Line"/>.
// Parameters:
//   value - Line 
// Returns:
//   void  
func (instance *Floor) SetBorder(value *Line)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("Floor_SetBorder"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the background <see cref="Color"/> of the <see cref="Area"/>.
// Returns:
//   Color  
func (instance *Floor) GetBackgroundColor()  (*Color,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAD(C.CString("Floor_GetBackgroundColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Color{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets or sets the background <see cref="Color"/> of the <see cref="Area"/>.
// Parameters:
//   value - Color 
// Returns:
//   void  
func (instance *Floor) SetBackgroundColor(value *Color)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAE(C.CString("Floor_SetBackgroundColor"), instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the foreground <see cref="Color"/>.
// Returns:
//   Color  
func (instance *Floor) GetForegroundColor()  (*Color,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAD(C.CString("Floor_GetForegroundColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Color{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets or sets the foreground <see cref="Color"/>.
// Parameters:
//   value - Color 
// Returns:
//   void  
func (instance *Floor) SetForegroundColor(value *Color)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAE(C.CString("Floor_SetForegroundColor"), instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the formatting of the area.
// Returns:
//   int32  
func (instance *Floor) GetFormatting()  (FormattingType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Floor_GetFormatting"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToFormattingType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the formatting of the area.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Floor) SetFormatting(value FormattingType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Floor_SetFormatting"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// If the property is true and the value of chart point is a negative number,
// the foreground color and background color will be exchanged.
// Returns:
//   bool  
func (instance *Floor) GetInvertIfNegative()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Floor_GetInvertIfNegative"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// If the property is true and the value of chart point is a negative number,
// the foreground color and background color will be exchanged.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Floor) SetInvertIfNegative(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Floor_SetInvertIfNegative"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents a <see cref="FillFormat"/> object that contains fill formatting properties for the specified chart or shape.
// Returns:
//   FillFormat  
func (instance *Floor) GetFillFormat()  (*FillFormat,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Floor_GetFillFormat"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &FillFormat{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFillFormat) 

	return result, nil 
}
// Returns or sets the degree of transparency of the area as a value from 0.0 (opaque) through 1.0 (clear).
// Returns:
//   float64  
func (instance *Floor) GetTransparency()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Floor_GetTransparency"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the degree of transparency of the area as a value from 0.0 (opaque) through 1.0 (clear).
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Floor) SetTransparency(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Floor_SetTransparency"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *Floor) ToArea() *Area {
	parentClass := &Area{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteFloor(floor *Floor){
	runtime.SetFinalizer(floor, nil)
	C.Delete_CObject(C.CString("Delete_Floor"),floor.ptr)
	floor.ptr = nil
}

// Class Legend 

// Encapsulates the object that represents the chart legend.
type Legend struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ChartTextFrame 
func NewLegend(src *ChartTextFrame) ( *Legend, error) {
	legend := &Legend{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_Legend"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		legend.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(legend, DeleteLegend)
		return legend, nil
	} else {
		legend.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return legend, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *Legend) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Legend_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the legend position type.
// Returns:
//   int32  
func (instance *Legend) GetPosition()  (LegendPositionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Legend_GetPosition"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToLegendPositionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the legend position type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Legend) SetPosition(value LegendPositionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Legend_SetPosition"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets a collection of all the LegendEntry objects in the specified chart legend.
// Setting the legend entries of the surface chart is not supported.
// So it will return null if the chart type is surface chart type.
// Returns:
//   LegendEntryCollection  
func (instance *Legend) GetLegendEntries()  (*LegendEntryCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Legend_GetLegendEntries"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &LegendEntryCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLegendEntryCollection) 

	return result, nil 
}
// Gets the labels of the legend entries after call Chart.Calculate() method.
// Returns:
//   []string  
func (instance *Legend) GetLegendLabels()  ([]string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZY(C.CString("Legend_GetLegendLabels"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result:= make([]string, CGoReturnPtr.column_length)
	for i := 0; i < int(CGoReturnPtr.column_length); i++ {
	   offset := uintptr(C.size_t(i)) * uintptr(CGoReturnPtr.size)
	   cObject := *(*C.char)(unsafe.Pointer( uintptr( unsafe.Pointer(CGoReturnPtr.return_value)) + offset))
	   goObject :=string(cObject)
	   result[i] = goObject
	}
	 

	return result, nil 
}
// Gets or sets whether showing the legend without overlapping the chart.
// Returns:
//   bool  
func (instance *Legend) IsOverLay()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Legend_IsOverLay"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets whether showing the legend without overlapping the chart.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Legend) SetIsOverLay(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Legend_SetIsOverLay"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.
// False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
// Returns:
//   bool  
func (instance *Legend) IsInnerMode()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Legend_IsInnerMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.
// False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Legend) SetIsInnerMode(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Legend_SetIsInnerMode"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the display mode of the background
// Returns:
//   int32  
func (instance *Legend) GetBackgroundMode()  (BackgroundMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Legend_GetBackgroundMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBackgroundMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the display mode of the background
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Legend) SetBackgroundMode(value BackgroundMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Legend_SetBackgroundMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the frame has a shadow.
// Returns:
//   bool  
func (instance *Legend) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Legend_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the frame has a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Legend) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Legend_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="ShapeProperties"/> object.
// Returns:
//   ShapePropertyCollection  
func (instance *Legend) GetShapeProperties()  (*ShapePropertyCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Legend_GetShapeProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ShapePropertyCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteShapePropertyCollection) 

	return result, nil 
}
// Indicates whether default position(DefaultX, DefaultY, DefaultWidth and DefaultHeight) are set.
// Returns:
//   bool  
func (instance *Legend) IsDefaultPosBeSet()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Legend_IsDefaultPosBeSet"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents x of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *Legend) GetDefaultXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Legend_GetDefaultXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents y of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *Legend) GetDefaultYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Legend_GetDefaultYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents width of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *Legend) GetDefaultWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Legend_GetDefaultWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents height of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *Legend) GetDefaultHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Legend_GetDefaultHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="Line">border</see>.
// Returns:
//   Line  
func (instance *Legend) GetBorder()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Legend_GetBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Gets the <see cref="Area">area</see>.
// Returns:
//   Area  
func (instance *Legend) GetArea()  (*Area,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Legend_GetArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Area{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteArea) 

	return result, nil 
}
// Gets and sets the options of the text.
// Returns:
//   TextOptions  
func (instance *Legend) GetTextOptions()  (*TextOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Legend_GetTextOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TextOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTextOptions) 

	return result, nil 
}
// Gets a <see cref="Font"/> object of the specified ChartFrame object.
// Returns:
//   Font  
func (instance *Legend) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Legend_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Returns:
//   bool  
func (instance *Legend) GetAutoScaleFont()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Legend_GetAutoScaleFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Legend) SetAutoScaleFont(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Legend_SetAutoScaleFont"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the chart frame is automatic sized.
// Returns:
//   bool  
func (instance *Legend) IsAutomaticSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Legend_IsAutomaticSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the chart frame is automatic sized.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Legend) SetIsAutomaticSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Legend_SetIsAutomaticSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the x coordinate of the upper left corner in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *Legend) GetXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Legend_GetXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the x coordinate of the upper left corner in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Legend) SetXRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Legend_SetXRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the y coordinate of the upper left corner in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *Legend) GetYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Legend_GetYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the y coordinate of the upper left corner in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Legend) SetYRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Legend_SetYRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of frame in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *Legend) GetWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Legend_GetWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of frame in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Legend) SetWidthRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Legend_SetWidthRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the height of frame in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *Legend) GetHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Legend_GetHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the height of frame in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Legend) SetHeightRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Legend_SetHeightRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Pixel.
// Returns:
//   int32  
func (instance *Legend) GetXPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Legend_GetXPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Legend) SetXPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Legend_SetXPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Pixel.
// Returns:
//   int32  
func (instance *Legend) GetYPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Legend_GetYPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Legend) SetYPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Legend_SetYPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of frame in units of Pixel.
// Returns:
//   int32  
func (instance *Legend) GetWidthPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Legend_GetWidthPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of frame in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Legend) SetWidthPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Legend_SetWidthPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the height of frame in units of Pixel.
// Returns:
//   int32  
func (instance *Legend) GetHeightPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Legend_GetHeightPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the height of frame in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Legend) SetHeightPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Legend_SetHeightPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Set position of the frame to automatic
// Returns:
//   void  
func (instance *Legend) SetPositionAuto()  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("Legend_SetPositionAuto"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this data labels is deleted.
// Returns:
//   bool  
func (instance *Legend) IsDeleted()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Legend_IsDeleted"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this data labels is deleted.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Legend) SetIsDeleted(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Legend_SetIsDeleted"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the text horizontal alignment.
// Returns:
//   int32  
func (instance *Legend) GetTextHorizontalAlignment()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Legend_GetTextHorizontalAlignment"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the text horizontal alignment.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Legend) SetTextHorizontalAlignment(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Legend_SetTextHorizontalAlignment"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the text vertical alignment of text.
// Returns:
//   int32  
func (instance *Legend) GetTextVerticalAlignment()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Legend_GetTextVerticalAlignment"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the text vertical alignment of text.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Legend) SetTextVerticalAlignment(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Legend_SetTextVerticalAlignment"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents text rotation angle.
// Returns:
//   int32  
func (instance *Legend) GetRotationAngle()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Legend_GetRotationAngle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents text rotation angle.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Legend) SetRotationAngle(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Legend_SetRotationAngle"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the text of the chart is automatically rotated.
// Returns:
//   bool  
func (instance *Legend) IsAutomaticRotation()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Legend_IsAutomaticRotation"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns a Characters object that represents a range of characters within the text.
// Parameters:
//   startIndex - int32 
//   length - int32 
// Returns:
//   FontSetting  
func (instance *Legend) Characters(startindex int32, length int32)  (*FontSetting,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBG(C.CString("Legend_Characters"), instance.ptr, C.int(startindex), C.int(length))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &FontSetting{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFontSetting) 

	return result, nil 
}
// Represents text reading order.
// Returns:
//   int32  
func (instance *Legend) GetReadingOrder()  (TextDirectionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Legend_GetReadingOrder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextDirectionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents text reading order.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Legend) SetReadingOrder(value TextDirectionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Legend_SetReadingOrder"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets whether a shape should be auto-fit to fully contain the text described within it. Auto-fitting is
// when text within a shape is scaled in order to contain all the text inside.
// Returns:
//   bool  
func (instance *Legend) IsResizeShapeToFitText()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Legend_IsResizeShapeToFitText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets whether a shape should be auto-fit to fully contain the text described within it. Auto-fitting is
// when text within a shape is scaled in order to contain all the text inside.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Legend) SetIsResizeShapeToFitText(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Legend_SetIsResizeShapeToFitText"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates the text is auto generated.
// Returns:
//   bool  
func (instance *Legend) IsAutoText()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Legend_IsAutoText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates the text is auto generated.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Legend) SetIsAutoText(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Legend_SetIsAutoText"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the text of a frame's title.
// Returns:
//   string  
func (instance *Legend) GetText()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Legend_GetText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the text of a frame's title.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Legend) SetText(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Legend_SetText"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets a reference to the worksheet.
// Returns:
//   string  
func (instance *Legend) GetLinkedSource()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Legend_GetLinkedSource"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets a reference to the worksheet.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Legend) SetLinkedSource(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Legend_SetLinkedSource"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the direction of text.
// Returns:
//   int32  
func (instance *Legend) GetDirectionType()  (ChartTextDirectionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Legend_GetDirectionType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToChartTextDirectionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the direction of text.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Legend) SetDirectionType(value ChartTextDirectionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Legend_SetDirectionType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets a value indicating whether the text is wrapped.
// Returns:
//   bool  
func (instance *Legend) IsTextWrapped()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Legend_IsTextWrapped"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets a value indicating whether the text is wrapped.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Legend) SetIsTextWrapped(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Legend_SetIsTextWrapped"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *Legend) ToChartTextFrame() *ChartTextFrame {
	parentClass := &ChartTextFrame{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *Legend) ToChartFrame() *ChartFrame {
	parentClass := &ChartFrame{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteLegend(legend *Legend){
	runtime.SetFinalizer(legend, nil)
	C.Delete_CObject(C.CString("Delete_Legend"),legend.ptr)
	legend.ptr = nil
}

// Class LegendEntry 

// Represents a legend entry in a chart legend.
type LegendEntry struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *LegendEntry) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LegendEntry_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets whether the legend entry is deleted.
// Returns:
//   bool  
func (instance *LegendEntry) IsDeleted()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LegendEntry_IsDeleted"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets whether the legend entry is deleted.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LegendEntry) SetIsDeleted(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("LegendEntry_SetIsDeleted"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets a <see cref="Font"/> object of the specified ChartFrame object.
// Returns:
//   Font  
func (instance *LegendEntry) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LegendEntry_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Gets or sets no fill of the text.
// Returns:
//   bool  
func (instance *LegendEntry) IsTextNoFill()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LegendEntry_IsTextNoFill"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets no fill of the text.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LegendEntry) SetIsTextNoFill(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("LegendEntry_SetIsTextNoFill"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the text in the object changes font size when the object size changes.
// The default value is True.
// Returns:
//   bool  
func (instance *LegendEntry) GetAutoScaleFont()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LegendEntry_GetAutoScaleFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes.
// The default value is True.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LegendEntry) SetAutoScaleFont(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("LegendEntry_SetAutoScaleFont"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the display mode of the background
// Returns:
//   int32  
func (instance *LegendEntry) GetBackgroundMode()  (BackgroundMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("LegendEntry_GetBackgroundMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBackgroundMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the display mode of the background
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LegendEntry) SetBackgroundMode(value BackgroundMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("LegendEntry_SetBackgroundMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteLegendEntry(legendentry *LegendEntry){
	runtime.SetFinalizer(legendentry, nil)
	C.Delete_CObject(C.CString("Delete_LegendEntry"),legendentry.ptr)
	legendentry.ptr = nil
}

// Class LegendEntryCollection 

// Represents a collection of all the <see cref="LegendEntry"/> objects in the specified chart legend.
type LegendEntryCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *LegendEntryCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LegendEntryCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="LegendEntry"/> element at the specified index.
// Parameters:
//   index - int32 
// Returns:
//   LegendEntry  
func (instance *LegendEntryCollection) Get(index int32)  (*LegendEntry,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("LegendEntryCollection_Get"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &LegendEntry{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLegendEntry) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *LegendEntryCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("LegendEntryCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteLegendEntryCollection(legendentrycollection *LegendEntryCollection){
	runtime.SetFinalizer(legendentrycollection, nil)
	C.Delete_CObject(C.CString("Delete_LegendEntryCollection"),legendentrycollection.ptr)
	legendentrycollection.ptr = nil
}

// Class Marker 

// Represents the marker in a line chart, scatter chart, or radar chart.
type Marker struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *Marker) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Marker_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="Line">border</see>.
// Returns:
//   Line  
func (instance *Marker) GetBorder()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Marker_GetBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Gets the <see cref="Area">area</see>.
// Returns:
//   Area  
func (instance *Marker) GetArea()  (*Area,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Marker_GetArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Area{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteArea) 

	return result, nil 
}
// Represents the marker style. Applies to line chart, scatter chart, or radar chart.
// Returns:
//   int32  
func (instance *Marker) GetMarkerStyle()  (ChartMarkerType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Marker_GetMarkerStyle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToChartMarkerType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the marker style. Applies to line chart, scatter chart, or radar chart.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Marker) SetMarkerStyle(value ChartMarkerType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Marker_SetMarkerStyle"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the marker size in unit of points. Applies to line chart, scatter chart, or radar chart.
// Returns:
//   int32  
func (instance *Marker) GetMarkerSize()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Marker_GetMarkerSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the marker size in unit of points. Applies to line chart, scatter chart, or radar chart.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Marker) SetMarkerSize(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Marker_SetMarkerSize"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the marker size in unit of pixels. Applies to line chart, scatter chart, or radar chart.
// Returns:
//   int32  
func (instance *Marker) GetMarkerSizePx()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Marker_GetMarkerSizePx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the marker size in unit of pixels. Applies to line chart, scatter chart, or radar chart.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Marker) SetMarkerSizePx(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Marker_SetMarkerSizePx"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the marker foreground color in a line chart, scatter chart, or radar chart.
// Returns:
//   Color  
func (instance *Marker) GetForegroundColor()  (*Color,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAD(C.CString("Marker_GetForegroundColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Color{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Represents the marker foreground color in a line chart, scatter chart, or radar chart.
// Parameters:
//   value - Color 
// Returns:
//   void  
func (instance *Marker) SetForegroundColor(value *Color)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAE(C.CString("Marker_SetForegroundColor"), instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the marker foreground color set type.
// Returns:
//   int32  
func (instance *Marker) GetForegroundColorSetType()  (FormattingType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Marker_GetForegroundColorSetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToFormattingType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the marker foreground color set type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Marker) SetForegroundColorSetType(value FormattingType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Marker_SetForegroundColorSetType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the marker background color in a line chart, scatter chart, or radar chart.
// Returns:
//   Color  
func (instance *Marker) GetBackgroundColor()  (*Color,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAD(C.CString("Marker_GetBackgroundColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Color{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Represents the marker background color in a line chart, scatter chart, or radar chart.
// Parameters:
//   value - Color 
// Returns:
//   void  
func (instance *Marker) SetBackgroundColor(value *Color)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAE(C.CString("Marker_SetBackgroundColor"), instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the marker background color set type.
// Returns:
//   int32  
func (instance *Marker) GetBackgroundColorSetType()  (FormattingType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Marker_GetBackgroundColorSetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToFormattingType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the marker background color set type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Marker) SetBackgroundColorSetType(value FormattingType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Marker_SetBackgroundColorSetType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteMarker(marker *Marker){
	runtime.SetFinalizer(marker, nil)
	C.Delete_CObject(C.CString("Delete_Marker"),marker.ptr)
	marker.ptr = nil
}

// Class PivotOptions 

// Represents a complex type that specifies the pivot controls that appear on the chart
type PivotOptions struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PivotOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies whether a control for each PivotTable field on the PivotTable page axis
// of the source PivotTable appears on the chart when dropZonesVisible is set to true.
// Returns:
//   bool  
func (instance *PivotOptions) GetDropZoneFilter()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PivotOptions_GetDropZoneFilter"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies whether a control for each PivotTable field on the PivotTable page axis
// of the source PivotTable appears on the chart when dropZonesVisible is set to true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotOptions) SetDropZoneFilter(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PivotOptions_SetDropZoneFilter"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies whether a control for each PivotTable field on the PivotTable row axis
// of the source PivotTable appears on the chart when dropZonesVisible is set to true.
// Returns:
//   bool  
func (instance *PivotOptions) GetDropZoneCategories()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PivotOptions_GetDropZoneCategories"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies whether a control for each PivotTable field on the PivotTable row axis
// of the source PivotTable appears on the chart when dropZonesVisible is set to true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotOptions) SetDropZoneCategories(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PivotOptions_SetDropZoneCategories"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies whether a control for each PivotTable field on the PivotTable data axis
// of the source PivotTable appears on the chart when dropZonesVisible is set to true.
// Returns:
//   bool  
func (instance *PivotOptions) GetDropZoneData()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PivotOptions_GetDropZoneData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies whether a control for each PivotTable field on the PivotTable data axis
// of the source PivotTable appears on the chart when dropZonesVisible is set to true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotOptions) SetDropZoneData(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PivotOptions_SetDropZoneData"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies whether a control for each PivotTable field on the PivotTable column axis
// of the source PivotTable appears on the chart when dropZonesVisible is set to true.
// Returns:
//   bool  
func (instance *PivotOptions) GetDropZoneSeries()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PivotOptions_GetDropZoneSeries"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies whether a control for each PivotTable field on the PivotTable column axis
// of the source PivotTable appears on the chart when dropZonesVisible is set to true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotOptions) SetDropZoneSeries(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PivotOptions_SetDropZoneSeries"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies whether any pivot controls can appear on the pivot chart.
// Returns:
//   bool  
func (instance *PivotOptions) GetDropZonesVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PivotOptions_GetDropZonesVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies whether any pivot controls can appear on the pivot chart.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotOptions) SetDropZonesVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PivotOptions_SetDropZonesVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeletePivotOptions(pivotoptions *PivotOptions){
	runtime.SetFinalizer(pivotoptions, nil)
	C.Delete_CObject(C.CString("Delete_PivotOptions"),pivotoptions.ptr)
	pivotoptions.ptr = nil
}

// Class PlotArea 

// Encapsulates the object that represents the plot area in a chart.
type PlotArea struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ChartFrame 
func NewPlotArea(src *ChartFrame) ( *PlotArea, error) {
	plotarea := &PlotArea{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_PlotArea"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		plotarea.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(plotarea, DeletePlotArea)
		return plotarea, nil
	} else {
		plotarea.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return plotarea, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PlotArea) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PlotArea_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or gets the x coordinate of the upper left corner of plot-area bounding box in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *PlotArea) GetXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("PlotArea_GetXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or gets the x coordinate of the upper left corner of plot-area bounding box in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *PlotArea) SetXRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("PlotArea_SetXRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or gets the y coordinate of the upper top corner  of plot-area bounding box in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *PlotArea) GetYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("PlotArea_GetYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or gets the y coordinate of the upper top corner  of plot-area bounding box in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *PlotArea) SetYRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("PlotArea_SetYRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the height of plot-area bounding box in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *PlotArea) GetHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("PlotArea_GetHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the height of plot-area bounding box in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *PlotArea) SetHeightRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("PlotArea_SetHeightRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of plot-area bounding box in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *PlotArea) GetWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("PlotArea_GetWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of plot-area bounding box in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *PlotArea) SetWidthRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("PlotArea_SetWidthRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or gets the x coordinate of the upper top corner of plot area in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *PlotArea) GetInnerXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("PlotArea_GetInnerXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or gets the x coordinate of the upper top corner of plot area in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *PlotArea) SetInnerXRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("PlotArea_SetInnerXRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or gets the x coordinate of the upper top corner of plot area in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *PlotArea) GetInnerYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("PlotArea_GetInnerYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or gets the x coordinate of the upper top corner of plot area in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *PlotArea) SetInnerYRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("PlotArea_SetInnerYRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the height of plot area in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *PlotArea) GetInnerHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("PlotArea_GetInnerHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the height of plot area in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *PlotArea) SetInnerHeightRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("PlotArea_SetInnerHeightRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width  of plot area in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *PlotArea) GetInnerWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("PlotArea_GetInnerWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width  of plot area in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *PlotArea) SetInnerWidthRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("PlotArea_SetInnerWidthRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Set position of the plot area to automatic
// Returns:
//   void  
func (instance *PlotArea) SetPositionAuto()  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("PlotArea_SetPositionAuto"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the plot area is automatic sized.
// Returns:
//   bool  
func (instance *PlotArea) IsAutomaticSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PlotArea_IsAutomaticSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the plot area is automatic sized.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PlotArea) SetIsAutomaticSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PlotArea_SetIsAutomaticSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.
// False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
// Returns:
//   bool  
func (instance *PlotArea) IsInnerMode()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PlotArea_IsInnerMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.
// False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PlotArea) SetIsInnerMode(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PlotArea_SetIsInnerMode"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the display mode of the background
// Returns:
//   int32  
func (instance *PlotArea) GetBackgroundMode()  (BackgroundMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("PlotArea_GetBackgroundMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBackgroundMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the display mode of the background
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PlotArea) SetBackgroundMode(value BackgroundMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("PlotArea_SetBackgroundMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the frame has a shadow.
// Returns:
//   bool  
func (instance *PlotArea) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PlotArea_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the frame has a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PlotArea) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PlotArea_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="ShapeProperties"/> object.
// Returns:
//   ShapePropertyCollection  
func (instance *PlotArea) GetShapeProperties()  (*ShapePropertyCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("PlotArea_GetShapeProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ShapePropertyCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteShapePropertyCollection) 

	return result, nil 
}
// Indicates whether default position(DefaultX, DefaultY, DefaultWidth and DefaultHeight) are set.
// Returns:
//   bool  
func (instance *PlotArea) IsDefaultPosBeSet()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PlotArea_IsDefaultPosBeSet"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents x of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *PlotArea) GetDefaultXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("PlotArea_GetDefaultXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents y of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *PlotArea) GetDefaultYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("PlotArea_GetDefaultYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents width of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *PlotArea) GetDefaultWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("PlotArea_GetDefaultWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents height of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *PlotArea) GetDefaultHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("PlotArea_GetDefaultHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="Line">border</see>.
// Returns:
//   Line  
func (instance *PlotArea) GetBorder()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("PlotArea_GetBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Gets the <see cref="Area">area</see>.
// Returns:
//   Area  
func (instance *PlotArea) GetArea()  (*Area,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("PlotArea_GetArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Area{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteArea) 

	return result, nil 
}
// Gets and sets the options of the text.
// Returns:
//   TextOptions  
func (instance *PlotArea) GetTextOptions()  (*TextOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("PlotArea_GetTextOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TextOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTextOptions) 

	return result, nil 
}
// Gets a <see cref="Font"/> object of the specified ChartFrame object.
// Returns:
//   Font  
func (instance *PlotArea) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("PlotArea_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Returns:
//   bool  
func (instance *PlotArea) GetAutoScaleFont()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PlotArea_GetAutoScaleFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PlotArea) SetAutoScaleFont(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PlotArea_SetAutoScaleFont"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Pixel.
// Returns:
//   int32  
func (instance *PlotArea) GetXPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("PlotArea_GetXPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PlotArea) SetXPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("PlotArea_SetXPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Pixel.
// Returns:
//   int32  
func (instance *PlotArea) GetYPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("PlotArea_GetYPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PlotArea) SetYPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("PlotArea_SetYPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of frame in units of Pixel.
// Returns:
//   int32  
func (instance *PlotArea) GetWidthPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("PlotArea_GetWidthPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of frame in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PlotArea) SetWidthPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("PlotArea_SetWidthPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the height of frame in units of Pixel.
// Returns:
//   int32  
func (instance *PlotArea) GetHeightPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("PlotArea_GetHeightPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the height of frame in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PlotArea) SetHeightPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("PlotArea_SetHeightPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *PlotArea) ToChartFrame() *ChartFrame {
	parentClass := &ChartFrame{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeletePlotArea(plotarea *PlotArea){
	runtime.SetFinalizer(plotarea, nil)
	C.Delete_CObject(C.CString("Delete_PlotArea"),plotarea.ptr)
	plotarea.ptr = nil
}

// Class Series 

// Encapsulates the object that represents a single data series in a chart.
type Series struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *Series) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Series_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the series is selected or filtered.True represents this series is filtered, and it will not be displayed on the chart.
// Returns:
//   bool  
func (instance *Series) IsFiltered()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Series_IsFiltered"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the series is selected or filtered.True represents this series is filtered, and it will not be displayed on the chart.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Series) SetIsFiltered(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Series_SetIsFiltered"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the properties of layout.
// Returns:
//   SeriesLayoutProperties  
func (instance *Series) GetLayoutProperties()  (*SeriesLayoutProperties,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetLayoutProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SeriesLayoutProperties{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSeriesLayoutProperties) 

	return result, nil 
}
// Moves the series up or down.
// Parameters:
//   count - int32 
// Returns:
//   void  
func (instance *Series) Move(count int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Series_Move"), instance.ptr, C.int(count))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the collection of points in a series in a chart.
// Returns:
//   ChartPointCollection  
func (instance *Series) GetPoints()  (*ChartPointCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetPoints"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ChartPointCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteChartPointCollection) 

	return result, nil 
}
// Represents the background area of Series object.
// Returns:
//   Area  
func (instance *Series) GetArea()  (*Area,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Area{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteArea) 

	return result, nil 
}
// Represents border of Series object.
// Returns:
//   Line  
func (instance *Series) GetBorder()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Gets or sets the name of the data series.
// Returns:
//   string  
func (instance *Series) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Series_GetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the name of the data series.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Series) SetName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Series_SetName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the series's name that displays on the chart graph.
// Returns:
//   string  
func (instance *Series) GetDisplayName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Series_GetDisplayName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the number of the data values.
// Returns:
//   int32  
func (instance *Series) GetCountOfDataValues()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Series_GetCountOfDataValues"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the data source is vertical.
// Returns:
//   bool  
func (instance *Series) IsVerticalValues()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Series_IsVerticalValues"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the Y values of this chart series.
// Returns:
//   string  
func (instance *Series) GetValues()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Series_GetValues"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the Y values of this chart series.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Series) SetValues(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Series_SetValues"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents format code of Values's NumberList.
// Returns:
//   string  
func (instance *Series) GetValuesFormatCode()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Series_GetValuesFormatCode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents format code of Values's NumberList.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Series) SetValuesFormatCode(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Series_SetValuesFormatCode"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents format code of X Values's NumberList.
// Returns:
//   string  
func (instance *Series) GetXValuesFormatCode()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Series_GetXValuesFormatCode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents format code of X Values's NumberList.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Series) SetXValuesFormatCode(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Series_SetXValuesFormatCode"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the x values of the chart series.
// Returns:
//   string  
func (instance *Series) GetXValues()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Series_GetXValues"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the x values of the chart series.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Series) SetXValues(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Series_SetXValues"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the bubble sizes values of the chart series.
// Returns:
//   string  
func (instance *Series) GetBubbleSizes()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Series_GetBubbleSizes"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the bubble sizes values of the chart series.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Series) SetBubbleSizes(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Series_SetBubbleSizes"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns all the trendlines of this series.
// Returns:
//   TrendlineCollection  
func (instance *Series) GetTrendLines()  (*TrendlineCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetTrendLines"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TrendlineCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTrendlineCollection) 

	return result, nil 
}
// Represents curve smoothing.
// True if curve smoothing is turned on for the line chart or scatter chart.
// Applies only to line and scatter connected by lines charts.
// Returns:
//   bool  
func (instance *Series) GetSmooth()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Series_GetSmooth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents curve smoothing.
// True if curve smoothing is turned on for the line chart or scatter chart.
// Applies only to line and scatter connected by lines charts.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Series) SetSmooth(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Series_SetSmooth"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the series has a shadow.
// Returns:
//   bool  
func (instance *Series) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Series_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the series has a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Series) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Series_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the series has a three-dimensional appearance.
// Applies only to bubble charts.
// Returns:
//   bool  
func (instance *Series) GetHas3DEffect()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Series_GetHas3DEffect"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the series has a three-dimensional appearance.
// Applies only to bubble charts.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Series) SetHas3DEffect(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Series_SetHas3DEffect"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the 3D shape type used with the 3-D bar or column chart.
// Returns:
//   int32  
func (instance *Series) GetBar3DShapeType()  (Bar3DShapeType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Series_GetBar3DShapeType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBar3DShapeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the 3D shape type used with the 3-D bar or column chart.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Series) SetBar3DShapeType(value Bar3DShapeType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Series_SetBar3DShapeType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the DataLabels object for the specified ASeries.
// Returns:
//   DataLabels  
func (instance *Series) GetDataLabels()  (*DataLabels,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetDataLabels"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DataLabels{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDataLabels) 

	return result, nil 
}
// Gets or sets a data series' type.
// Returns:
//   int32  
func (instance *Series) GetType()  (ChartType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Series_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToChartType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets a data series' type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Series) SetType(value ChartType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Series_SetType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="Marker">marker</see>.
// Returns:
//   Marker  
func (instance *Series) GetMarker()  (*Marker,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetMarker"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Marker{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteMarker) 

	return result, nil 
}
// Indicates if this series is plotted on second value axis.
// Returns:
//   bool  
func (instance *Series) GetPlotOnSecondAxis()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Series_GetPlotOnSecondAxis"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates if this series is plotted on second value axis.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Series) SetPlotOnSecondAxis(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Series_SetPlotOnSecondAxis"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents X direction error bar of the series.
// Returns:
//   ErrorBar  
func (instance *Series) GetXErrorBar()  (*ErrorBar,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetXErrorBar"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ErrorBar{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteErrorBar) 

	return result, nil 
}
// Represents Y direction error bar of the series.
// Returns:
//   ErrorBar  
func (instance *Series) GetYErrorBar()  (*ErrorBar,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetYErrorBar"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ErrorBar{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteErrorBar) 

	return result, nil 
}
// True if the line chart has high-low lines.
// Applies only to line charts.
// Returns:
//   bool  
func (instance *Series) GetHasHiLoLines()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Series_GetHasHiLoLines"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the line chart has high-low lines.
// Applies only to line charts.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Series) SetHasHiLoLines(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Series_SetHasHiLoLines"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns a HiLoLines object that represents the high-low lines for a series on a line chart.
// Applies only to line charts.
// Returns:
//   Line  
func (instance *Series) GetHiLoLines()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetHiLoLines"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// True if a stacked column chart or bar chart has series lines or
// if a Pie of Pie chart or Bar of Pie chart has connector lines between the two sections.
// Applies only to stacked column charts, bar charts, Pie of Pie charts, or Bar of Pie charts.
// Returns:
//   bool  
func (instance *Series) GetHasSeriesLines()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Series_GetHasSeriesLines"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if a stacked column chart or bar chart has series lines or
// if a Pie of Pie chart or Bar of Pie chart has connector lines between the two sections.
// Applies only to stacked column charts, bar charts, Pie of Pie charts, or Bar of Pie charts.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Series) SetHasSeriesLines(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Series_SetHasSeriesLines"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns a SeriesLines object that represents the series lines for a stacked bar chart or a stacked column chart.
// Applies only to stacked bar and stacked column charts.
// Returns:
//   Line  
func (instance *Series) GetSeriesLines()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetSeriesLines"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// True if the chart has drop lines.
// Applies only to line chart or area charts.
// Returns:
//   bool  
func (instance *Series) GetHasDropLines()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Series_GetHasDropLines"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the chart has drop lines.
// Applies only to line chart or area charts.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Series) SetHasDropLines(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Series_SetHasDropLines"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns a <see cref="Line"/> object that represents the drop lines for a series on the line chart or area chart.
// Applies only to line chart or area charts.
// Returns:
//   Line  
func (instance *Series) GetDropLines()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetDropLines"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// True if a line chart has up and down bars.
// Applies only to line charts.
// Returns:
//   bool  
func (instance *Series) GetHasUpDownBars()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Series_GetHasUpDownBars"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if a line chart has up and down bars.
// Applies only to line charts.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Series) SetHasUpDownBars(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Series_SetHasUpDownBars"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns an DropBars object that represents the up bars on a line chart.
// Applies only to line charts.
// Returns:
//   DropBars  
func (instance *Series) GetUpBars()  (*DropBars,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetUpBars"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DropBars{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDropBars) 

	return result, nil 
}
// Returns a <see cref="DropBars"/> object that represents the down bars on a line chart.
// Applies only to line charts.
// Returns:
//   DropBars  
func (instance *Series) GetDownBars()  (*DropBars,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetDownBars"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DropBars{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDropBars) 

	return result, nil 
}
// Represents if the color of points is varied.
// The chart must contain only one series.
// Returns:
//   bool  
func (instance *Series) IsColorVaried()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Series_IsColorVaried"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents if the color of points is varied.
// The chart must contain only one series.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Series) SetIsColorVaried(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Series_SetIsColorVaried"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the space between bar or column clusters, as a percentage of the bar or column width.
// The value of this property must be between 0 and 500.
// Returns:
//   int16  
func (instance *Series) GetGapWidth()  (int16,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZHT(C.CString("Series_GetGapWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int16(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the space between bar or column clusters, as a percentage of the bar or column width.
// The value of this property must be between 0 and 500.
// Parameters:
//   value - int16 
// Returns:
//   void  
func (instance *Series) SetGapWidth(value int16)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMP(C.CString("Series_SetGapWidth"), instance.ptr, C.short(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the angle of the first pie-chart or doughnut-chart slice, in degrees (clockwise from vertical).
// Applies only to pie, 3-D pie, and doughnut charts, 0 to 360.
// Returns:
//   int16  
func (instance *Series) GetFirstSliceAngle()  (int16,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZHT(C.CString("Series_GetFirstSliceAngle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int16(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the angle of the first pie-chart or doughnut-chart slice, in degrees (clockwise from vertical).
// Applies only to pie, 3-D pie, and doughnut charts, 0 to 360.
// Parameters:
//   value - int16 
// Returns:
//   void  
func (instance *Series) SetFirstSliceAngle(value int16)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMP(C.CString("Series_SetFirstSliceAngle"), instance.ptr, C.short(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies how bars and columns are positioned.
// Can be a value between – 100 and 100.
// Applies only to 2-D bar and 2-D column charts.
// Returns:
//   int16  
func (instance *Series) GetOverlap()  (int16,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZHT(C.CString("Series_GetOverlap"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int16(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies how bars and columns are positioned.
// Can be a value between – 100 and 100.
// Applies only to 2-D bar and 2-D column charts.
// Parameters:
//   value - int16 
// Returns:
//   void  
func (instance *Series) SetOverlap(value int16)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMP(C.CString("Series_SetOverlap"), instance.ptr, C.short(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the size of the secondary section of either a pie of pie chart or a bar of pie chart,
// as a percentage of the size of the primary pie.
// Can be a value from 5 to 200.
// Returns:
//   int16  
func (instance *Series) GetSecondPlotSize()  (int16,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZHT(C.CString("Series_GetSecondPlotSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int16(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the size of the secondary section of either a pie of pie chart or a bar of pie chart,
// as a percentage of the size of the primary pie.
// Can be a value from 5 to 200.
// Parameters:
//   value - int16 
// Returns:
//   void  
func (instance *Series) SetSecondPlotSize(value int16)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMP(C.CString("Series_SetSecondPlotSize"), instance.ptr, C.short(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets a value that how to determine which data points are in the second pie or bar on a pie of pie or bar of
// pie chart.
// Returns:
//   int32  
func (instance *Series) GetSplitType()  (ChartSplitType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Series_GetSplitType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToChartSplitType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Returns or sets a value that how to determine which data points are in the second pie or bar on a pie of pie or bar of
// pie chart.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Series) SetSplitType(value ChartSplitType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Series_SetSplitType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets a value that shall be used to determine which data points are in the second pie or bar on
// a pie of pie or bar of pie chart.
// Returns:
//   float64  
func (instance *Series) GetSplitValue()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Series_GetSplitValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets a value that shall be used to determine which data points are in the second pie or bar on
// a pie of pie or bar of pie chart.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Series) SetSplitValue(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Series_SetSplitValue"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the threshold value is automatic.
// Returns:
//   bool  
func (instance *Series) IsAutoSplit()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Series_IsAutoSplit"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the scale factor for bubbles in the specified chart group.
// It can be an integer value from 0 (zero) to 300,
// corresponding to a percentage of the default size.
// Applies only to bubble charts.
// Returns:
//   int32  
func (instance *Series) GetBubbleScale()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Series_GetBubbleScale"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the scale factor for bubbles in the specified chart group.
// It can be an integer value from 0 (zero) to 300,
// corresponding to a percentage of the default size.
// Applies only to bubble charts.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Series) SetBubbleScale(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Series_SetBubbleScale"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets what the bubble size represents on a bubble chart.
// Returns:
//   int32  
func (instance *Series) GetSizeRepresents()  (BubbleSizeRepresents,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Series_GetSizeRepresents"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBubbleSizeRepresents(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets what the bubble size represents on a bubble chart.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Series) SetSizeRepresents(value BubbleSizeRepresents)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Series_SetSizeRepresents"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if negative bubbles are shown for the chart group. Valid only for bubble charts.
// Returns:
//   bool  
func (instance *Series) GetShowNegativeBubbles()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Series_GetShowNegativeBubbles"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if negative bubbles are shown for the chart group. Valid only for bubble charts.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Series) SetShowNegativeBubbles(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Series_SetShowNegativeBubbles"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the size of the hole in a doughnut chart group.
// The hole size is expressed as a percentage of the chart size, between 10 and 90 percent.
// Returns:
//   int32  
func (instance *Series) GetDoughnutHoleSize()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Series_GetDoughnutHoleSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the size of the hole in a doughnut chart group.
// The hole size is expressed as a percentage of the chart size, between 10 and 90 percent.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Series) SetDoughnutHoleSize(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Series_SetDoughnutHoleSize"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The distance of an open pie slice from the center of the pie chart is expressed as a percentage of the pie diameter.
// Returns:
//   int32  
func (instance *Series) GetExplosion()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Series_GetExplosion"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// The distance of an open pie slice from the center of the pie chart is expressed as a percentage of the pie diameter.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Series) SetExplosion(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Series_SetExplosion"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if a radar chart has category axis labels. Applies only to radar charts.
// Returns:
//   bool  
func (instance *Series) GetHasRadarAxisLabels()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Series_GetHasRadarAxisLabels"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if a radar chart has category axis labels. Applies only to radar charts.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Series) SetHasRadarAxisLabels(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Series_SetHasRadarAxisLabels"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the series has leader lines.
// Returns:
//   bool  
func (instance *Series) GetHasLeaderLines()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Series_GetHasLeaderLines"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the series has leader lines.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Series) SetHasLeaderLines(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Series_SetHasLeaderLines"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents leader lines on a chart. Leader lines connect data labels to data points.
// This object isn’t a collection; there’s no object that represents a single leader line.
// Returns:
//   Line  
func (instance *Series) GetLeaderLines()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetLeaderLines"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Gets the legend entry according to this series.
// Returns:
//   LegendEntry  
func (instance *Series) GetLegendEntry()  (*LegendEntry,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetLegendEntry"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &LegendEntry{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLegendEntry) 

	return result, nil 
}
// Gets the <see cref="ShapePropertyCollection"/> object that holds the visual shape properties of the Series.
// Returns:
//   ShapePropertyCollection  
func (instance *Series) GetShapeProperties()  (*ShapePropertyCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Series_GetShapeProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ShapePropertyCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteShapePropertyCollection) 

	return result, nil 
}



func DeleteSeries(series *Series){
	runtime.SetFinalizer(series, nil)
	C.Delete_CObject(C.CString("Delete_Series"),series.ptr)
	series.ptr = nil
}

// Class SeriesCollection 

// Encapsulates a collection of <see cref="Series"/> objects.
type SeriesCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SeriesCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SeriesCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="Series"/> element at the specified index.
// Parameters:
//   index - int32 
// Returns:
//   Series  
func (instance *SeriesCollection) Get(index int32)  (*Series,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("SeriesCollection_Get"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Series{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSeries) 

	return result, nil 
}
// Gets the <see cref="Series"/> element by order.
// Parameters:
//   order - int32 
// Returns:
//   Series  
func (instance *SeriesCollection) GetSeriesByOrder(order int32)  (*Series,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("SeriesCollection_GetSeriesByOrder"), instance.ptr, C.int(order))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Series{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSeries) 

	return result, nil 
}
// Remove at a series at the specific index.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *SeriesCollection) RemoveAt(index int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("SeriesCollection_RemoveAt"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the range of category Axis values.
// It can be a range of cells (such as, "d1:e10"),
// or a sequence of values (such as,"{2,6,8,10}").
// Returns:
//   string  
func (instance *SeriesCollection) GetCategoryData()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("SeriesCollection_GetCategoryData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the range of category Axis values.
// It can be a range of cells (such as, "d1:e10"),
// or a sequence of values (such as,"{2,6,8,10}").
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *SeriesCollection) SetCategoryData(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("SeriesCollection_SetCategoryData"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the range of second category Axis values.
// It can be a range of cells (such as, "d1:e10"),
// or a sequence of values (such as,"{2,6,8,10}").
// Only effects when some ASerieses plot on the second axis.
// Returns:
//   string  
func (instance *SeriesCollection) GetSecondCategoryData()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("SeriesCollection_GetSecondCategoryData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the range of second category Axis values.
// It can be a range of cells (such as, "d1:e10"),
// or a sequence of values (such as,"{2,6,8,10}").
// Only effects when some ASerieses plot on the second axis.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *SeriesCollection) SetSecondCategoryData(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("SeriesCollection_SetSecondCategoryData"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Directly changes the orders of the two series.
// Parameters:
//   sourceIndex - int32 
//   destIndex - int32 
// Returns:
//   void  
func (instance *SeriesCollection) SwapSeries(sourceindex int32, destindex int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZCH(C.CString("SeriesCollection_SwapSeries"), instance.ptr, C.int(sourceindex), C.int(destindex))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Sets the name of all the serieses in the chart.
// Parameters:
//   startIndex - int32 
//   area - string 
//   isVertical - bool 
// Returns:
//   void  
func (instance *SeriesCollection) SetSeriesNames(startindex int32, area string, isvertical bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMY(C.CString("SeriesCollection_SetSeriesNames"), instance.ptr, C.int(startindex), C.CString(area), C.bool(isvertical))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Adds the <see cref="Series"/> collection to a chart.
// Parameters:
//   area - string 
//   isVertical - bool 
// Returns:
//   int32  
func (instance *SeriesCollection) AddR1C1(area string, isvertical bool)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZNA(C.CString("SeriesCollection_AddR1C1"), instance.ptr, C.CString(area), C.bool(isvertical))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds the <see cref="Series"/> collection to a chart.
// Parameters:
//   area - string 
//   isVertical - bool 
// Returns:
//   int32  
func (instance *SeriesCollection) Add_String_Bool(area string, isvertical bool)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZNA(C.CString("SeriesCollection_Add_String_Boolean"), instance.ptr, C.CString(area), C.bool(isvertical))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds the <see cref="Series"/> collection to a chart.
// Parameters:
//   area - string 
//   isVertical - bool 
//   checkLabels - bool 
// Returns:
//   int32  
func (instance *SeriesCollection) Add_String_Bool_Bool(area string, isvertical bool, checklabels bool)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZNB(C.CString("SeriesCollection_Add_String_Boolean_Boolean"), instance.ptr, C.CString(area), C.bool(isvertical), C.bool(checklabels))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents if the color of points is varied.
// Returns:
//   bool  
func (instance *SeriesCollection) IsColorVaried()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SeriesCollection_IsColorVaried"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents if the color of points is varied.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SeriesCollection) SetIsColorVaried(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SeriesCollection_SetIsColorVaried"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Clears the collection
// Returns:
//   void  
func (instance *SeriesCollection) Clear()  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("SeriesCollection_Clear"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Set Monochromatic Palette for chart series.
// Parameters:
//   type - int32 
// Returns:
//   void  
func (instance *SeriesCollection) ChangeColors(type_ ChartColorPaletteType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("SeriesCollection_ChangeColors"), instance.ptr, C.int( int32(type_)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns:
//   int32  
func (instance *SeriesCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SeriesCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSeriesCollection(seriescollection *SeriesCollection){
	runtime.SetFinalizer(seriescollection, nil)
	C.Delete_CObject(C.CString("Delete_SeriesCollection"),seriescollection.ptr)
	seriescollection.ptr = nil
}

// Class SeriesLayoutProperties 

// Represents the properties of series layout.
type SeriesLayoutProperties struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewSeriesLayoutProperties() ( *SeriesLayoutProperties, error) {
	serieslayoutproperties := &SeriesLayoutProperties{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_SeriesLayoutProperties"),)
	if CGoReturnPtr.error_no == 0 {
		serieslayoutproperties.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(serieslayoutproperties, DeleteSeriesLayoutProperties)
		return serieslayoutproperties, nil
	} else {
		serieslayoutproperties.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return serieslayoutproperties, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SeriesLayoutProperties) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SeriesLayoutProperties_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether showing connector lines between data points.
// Returns:
//   bool  
func (instance *SeriesLayoutProperties) GetShowConnectorLines()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SeriesLayoutProperties_GetShowConnectorLines"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether showing connector lines between data points.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SeriesLayoutProperties) SetShowConnectorLines(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SeriesLayoutProperties_SetShowConnectorLines"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether showing the line connecting all mean points.
// Returns:
//   bool  
func (instance *SeriesLayoutProperties) GetShowMeanLine()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SeriesLayoutProperties_GetShowMeanLine"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether showing the line connecting all mean points.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SeriesLayoutProperties) SetShowMeanLine(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SeriesLayoutProperties_SetShowMeanLine"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether showing outlier data points.
// Returns:
//   bool  
func (instance *SeriesLayoutProperties) GetShowOutlierPoints()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SeriesLayoutProperties_GetShowOutlierPoints"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether showing outlier data points.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SeriesLayoutProperties) SetShowOutlierPoints(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SeriesLayoutProperties_SetShowOutlierPoints"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether showing markers denoting the mean.
// Returns:
//   bool  
func (instance *SeriesLayoutProperties) GetShowMeanMarker()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SeriesLayoutProperties_GetShowMeanMarker"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether showing markers denoting the mean.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SeriesLayoutProperties) SetShowMeanMarker(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SeriesLayoutProperties_SetShowMeanMarker"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether showing non-outlier data points.
// Returns:
//   bool  
func (instance *SeriesLayoutProperties) GetShowInnerPoints()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SeriesLayoutProperties_GetShowInnerPoints"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether showing non-outlier data points.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SeriesLayoutProperties) SetShowInnerPoints(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SeriesLayoutProperties_SetShowInnerPoints"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the index of a subtotal data point.
// Returns:
//   []int32_t  
func (instance *SeriesLayoutProperties) GetSubtotals()  ([]int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZY(C.CString("SeriesLayoutProperties_GetSubtotals"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result:= make([]int32, CGoReturnPtr.column_length)
	for i := 0; i < int(CGoReturnPtr.column_length); i++ {
	   offset := uintptr(C.size_t(i)) * uintptr(CGoReturnPtr.size)
	   cObject := *(*C.int)(unsafe.Pointer( uintptr( unsafe.Pointer(CGoReturnPtr.return_value)) + offset))
	   goObject :=int32(cObject)
	   result[i] = goObject
	}
	 

	return result, nil 
}
// Represents the index of a subtotal data point.
// Parameters:
//   value - []int32_t 
// Returns:
//   void  
func (instance *SeriesLayoutProperties) SetSubtotals(value []int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAJ(C.CString("SeriesLayoutProperties_SetSubtotals"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the statistical properties for the series.
// Returns:
//   int32  
func (instance *SeriesLayoutProperties) GetQuartileCalculation()  (QuartileCalculationType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("SeriesLayoutProperties_GetQuartileCalculation"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToQuartileCalculationType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the statistical properties for the series.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SeriesLayoutProperties) SetQuartileCalculation(value QuartileCalculationType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("SeriesLayoutProperties_SetQuartileCalculation"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the layout of map labels.
// Returns:
//   int32  
func (instance *SeriesLayoutProperties) GetMapLabelLayout()  (MapChartLabelLayout,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("SeriesLayoutProperties_GetMapLabelLayout"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMapChartLabelLayout(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the layout of map labels.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SeriesLayoutProperties) SetMapLabelLayout(value MapChartLabelLayout)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("SeriesLayoutProperties_SetMapLabelLayout"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the interval is closed on the left side.
// Returns:
//   bool  
func (instance *SeriesLayoutProperties) IsIntervalLeftClosed()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SeriesLayoutProperties_IsIntervalLeftClosed"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the interval is closed on the left side.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SeriesLayoutProperties) SetIsIntervalLeftClosed(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SeriesLayoutProperties_SetIsIntervalLeftClosed"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the region type of the map.
// Returns:
//   int32  
func (instance *SeriesLayoutProperties) GetMapChartRegionType()  (MapChartRegionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("SeriesLayoutProperties_GetMapChartRegionType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMapChartRegionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the region type of the map.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SeriesLayoutProperties) SetMapChartRegionType(value MapChartRegionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("SeriesLayoutProperties_SetMapChartRegionType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the projection type of the map.
// Returns:
//   int32  
func (instance *SeriesLayoutProperties) GetMapChartProjectionType()  (MapChartProjectionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("SeriesLayoutProperties_GetMapChartProjectionType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMapChartProjectionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the projection type of the map.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SeriesLayoutProperties) SetMapChartProjectionType(value MapChartProjectionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("SeriesLayoutProperties_SetMapChartProjectionType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteSeriesLayoutProperties(serieslayoutproperties *SeriesLayoutProperties){
	runtime.SetFinalizer(serieslayoutproperties, nil)
	C.Delete_CObject(C.CString("Delete_SeriesLayoutProperties"),serieslayoutproperties.ptr)
	serieslayoutproperties.ptr = nil
}

// Class Sparkline 

// A sparkline represents a tiny chart or graphic in a worksheet cell that provides a visual representation of data.
type Sparkline struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *Sparkline) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Sparkline_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the data range of the sparkline.
// Returns:
//   string  
func (instance *Sparkline) GetDataRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Sparkline_GetDataRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the data range of the sparkline.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Sparkline) SetDataRange(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Sparkline_SetDataRange"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the row index of the sparkline.
// Returns:
//   int32  
func (instance *Sparkline) GetRow()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Sparkline_GetRow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the column index of the sparkline.
// Returns:
//   int32  
func (instance *Sparkline) GetColumn()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Sparkline_GetColumn"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Converts a sparkline to an image.
// Parameters:
//   fileName - string 
//   options - ImageOrPrintOptions 
// Returns:
//   void  
func (instance *Sparkline) ToImage_String_ImageOrPrintOptions(filename string, options *ImageOrPrintOptions)  error {
	
	var options_ptr unsafe.Pointer = nil
	if options != nil {
	  options_ptr =options.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZAQ(C.CString("Sparkline_ToImage_String_ImageOrPrintOptions"), instance.ptr, C.CString(filename), options_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Converts a sparkline to an image.
// Parameters:
//   options - ImageOrPrintOptions 
// Returns:
//   []byte  
func (instance *Sparkline) ToImage_ImageOrPrintOptions(options *ImageOrPrintOptions)  ([]byte,  error)  {
	
	var options_ptr unsafe.Pointer = nil
	if options != nil {
	  options_ptr =options.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZET(C.CString("Sparkline_ToImage_ImageOrPrintOptions"), instance.ptr, options_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}



func DeleteSparkline(sparkline *Sparkline){
	runtime.SetFinalizer(sparkline, nil)
	C.Delete_CObject(C.CString("Delete_Sparkline"),sparkline.ptr)
	sparkline.ptr = nil
}

// Class SparklineCollection 

// Encapsulates a collection of <see cref="Sparkline"/> objects.
type SparklineCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SparklineCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SparklineCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="Sparkline"/> element at the specified index.
// Parameters:
//   index - int32 
// Returns:
//   Sparkline  
func (instance *SparklineCollection) Get(index int32)  (*Sparkline,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("SparklineCollection_Get"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Sparkline{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSparkline) 

	return result, nil 
}
// Add a sparkline.
// Parameters:
//   dataRange - string 
//   row - int32 
//   column - int32 
// Returns:
//   int32  
func (instance *SparklineCollection) Add(datarange string, row int32, column int32)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZGS(C.CString("SparklineCollection_Add"), instance.ptr, C.CString(datarange), C.int(row), C.int(column))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Removes the sparkline
// Parameters:
//   o - Object 
// Returns:
//   void  
func (instance *SparklineCollection) Remove(o *Object)  error {
	
	var o_ptr unsafe.Pointer = nil
	if o != nil {
	  o_ptr =o.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("SparklineCollection_Remove"), instance.ptr, o_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns:
//   int32  
func (instance *SparklineCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SparklineCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSparklineCollection(sparklinecollection *SparklineCollection){
	runtime.SetFinalizer(sparklinecollection, nil)
	C.Delete_CObject(C.CString("Delete_SparklineCollection"),sparklinecollection.ptr)
	sparklinecollection.ptr = nil
}

// Class SparklineGroup 

// <see cref="Sparkline"/> is organized into sparkline group. A SparklineGroup contains a variable number of sparkline items.
// A sparkline group specifies the type, display settings and axis settings for the sparklines.
type SparklineGroup struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SparklineGroup) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SparklineGroup_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Resets the data range and location range of the sparkline group.
// This method will clear original sparkline items in the group and creates new sparkline items for the new ranges.
// Parameters:
//   dataRange - string 
//   isVertical - bool 
//   locationRange - CellArea 
// Returns:
//   void  
func (instance *SparklineGroup) ResetRanges(datarange string, isvertical bool, locationrange *CellArea)  error {
	
	var locationrange_ptr unsafe.Pointer = nil
	if locationrange != nil {
	  locationrange_ptr =locationrange.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZNC(C.CString("SparklineGroup_ResetRanges"), instance.ptr, C.CString(datarange), C.bool(isvertical), locationrange_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the preset style type of the sparkline group.
// Returns:
//   int32  
func (instance *SparklineGroup) GetPresetStyle()  (SparklinePresetStyleType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("SparklineGroup_GetPresetStyle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSparklinePresetStyleType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the preset style type of the sparkline group.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SparklineGroup) SetPresetStyle(value SparklinePresetStyleType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("SparklineGroup_SetPresetStyle"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the collection of <see cref="Sparkline"/> object.
// Returns:
//   SparklineCollection  
func (instance *SparklineGroup) GetSparklines()  (*SparklineCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("SparklineGroup_GetSparklines"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SparklineCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSparklineCollection) 

	return result, nil 
}
// Indicates the sparkline type of the sparkline group.
// Returns:
//   int32  
func (instance *SparklineGroup) GetType()  (SparklineType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("SparklineGroup_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSparklineType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates the sparkline type of the sparkline group.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SparklineGroup) SetType(value SparklineType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("SparklineGroup_SetType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates how to plot empty cells.
// Returns:
//   int32  
func (instance *SparklineGroup) GetPlotEmptyCellsType()  (PlotEmptyCellsType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("SparklineGroup_GetPlotEmptyCellsType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPlotEmptyCellsType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates how to plot empty cells.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SparklineGroup) SetPlotEmptyCellsType(value PlotEmptyCellsType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("SparklineGroup_SetPlotEmptyCellsType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show data in hidden rows and columns.
// Returns:
//   bool  
func (instance *SparklineGroup) GetDisplayHidden()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SparklineGroup_GetDisplayHidden"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show data in hidden rows and columns.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SparklineGroup) SetDisplayHidden(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SparklineGroup_SetDisplayHidden"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to highlight the highest points of data in the sparkline group.
// Returns:
//   bool  
func (instance *SparklineGroup) GetShowHighPoint()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SparklineGroup_GetShowHighPoint"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to highlight the highest points of data in the sparkline group.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SparklineGroup) SetShowHighPoint(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SparklineGroup_SetShowHighPoint"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the color of the highest points of data in the sparkline group.
// Returns:
//   CellsColor  
func (instance *SparklineGroup) GetHighPointColor()  (*CellsColor,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("SparklineGroup_GetHighPointColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellsColor{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellsColor) 

	return result, nil 
}
// Gets and sets the color of the highest points of data in the sparkline group.
// Parameters:
//   value - CellsColor 
// Returns:
//   void  
func (instance *SparklineGroup) SetHighPointColor(value *CellsColor)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("SparklineGroup_SetHighPointColor"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to highlight the lowest points of data in the sparkline group.
// Returns:
//   bool  
func (instance *SparklineGroup) GetShowLowPoint()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SparklineGroup_GetShowLowPoint"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to highlight the lowest points of data in the sparkline group.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SparklineGroup) SetShowLowPoint(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SparklineGroup_SetShowLowPoint"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the color of the lowest points of data in the sparkline group.
// Returns:
//   CellsColor  
func (instance *SparklineGroup) GetLowPointColor()  (*CellsColor,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("SparklineGroup_GetLowPointColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellsColor{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellsColor) 

	return result, nil 
}
// Gets and sets the color of the lowest points of data in the sparkline group.
// Parameters:
//   value - CellsColor 
// Returns:
//   void  
func (instance *SparklineGroup) SetLowPointColor(value *CellsColor)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("SparklineGroup_SetLowPointColor"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to highlight the negative values on the sparkline group with a different color or marker.
// Returns:
//   bool  
func (instance *SparklineGroup) GetShowNegativePoints()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SparklineGroup_GetShowNegativePoints"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to highlight the negative values on the sparkline group with a different color or marker.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SparklineGroup) SetShowNegativePoints(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SparklineGroup_SetShowNegativePoints"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the color of the negative values on the sparkline group.
// Returns:
//   CellsColor  
func (instance *SparklineGroup) GetNegativePointsColor()  (*CellsColor,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("SparklineGroup_GetNegativePointsColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellsColor{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellsColor) 

	return result, nil 
}
// Gets and sets the color of the negative values on the sparkline group.
// Parameters:
//   value - CellsColor 
// Returns:
//   void  
func (instance *SparklineGroup) SetNegativePointsColor(value *CellsColor)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("SparklineGroup_SetNegativePointsColor"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to highlight the first point of data in the sparkline group.
// Returns:
//   bool  
func (instance *SparklineGroup) GetShowFirstPoint()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SparklineGroup_GetShowFirstPoint"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to highlight the first point of data in the sparkline group.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SparklineGroup) SetShowFirstPoint(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SparklineGroup_SetShowFirstPoint"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the color of the first point of data in the sparkline group.
// Returns:
//   CellsColor  
func (instance *SparklineGroup) GetFirstPointColor()  (*CellsColor,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("SparklineGroup_GetFirstPointColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellsColor{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellsColor) 

	return result, nil 
}
// Gets and sets the color of the first point of data in the sparkline group.
// Parameters:
//   value - CellsColor 
// Returns:
//   void  
func (instance *SparklineGroup) SetFirstPointColor(value *CellsColor)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("SparklineGroup_SetFirstPointColor"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to highlight the last point of data in the sparkline group.
// Returns:
//   bool  
func (instance *SparklineGroup) GetShowLastPoint()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SparklineGroup_GetShowLastPoint"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to highlight the last point of data in the sparkline group.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SparklineGroup) SetShowLastPoint(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SparklineGroup_SetShowLastPoint"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the color of the last point of data in the sparkline group.
// Returns:
//   CellsColor  
func (instance *SparklineGroup) GetLastPointColor()  (*CellsColor,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("SparklineGroup_GetLastPointColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellsColor{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellsColor) 

	return result, nil 
}
// Gets and sets the color of the last point of data in the sparkline group.
// Parameters:
//   value - CellsColor 
// Returns:
//   void  
func (instance *SparklineGroup) SetLastPointColor(value *CellsColor)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("SparklineGroup_SetLastPointColor"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to highlight each point in each line sparkline in the sparkline group.
// Returns:
//   bool  
func (instance *SparklineGroup) GetShowMarkers()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SparklineGroup_GetShowMarkers"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to highlight each point in each line sparkline in the sparkline group.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SparklineGroup) SetShowMarkers(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SparklineGroup_SetShowMarkers"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the color of points in each line sparkline in the sparkline group.
// Returns:
//   CellsColor  
func (instance *SparklineGroup) GetMarkersColor()  (*CellsColor,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("SparklineGroup_GetMarkersColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellsColor{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellsColor) 

	return result, nil 
}
// Gets and sets the color of points in each line sparkline in the sparkline group.
// Parameters:
//   value - CellsColor 
// Returns:
//   void  
func (instance *SparklineGroup) SetMarkersColor(value *CellsColor)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("SparklineGroup_SetMarkersColor"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the color of the sparklines in the sparkline group.
// Returns:
//   CellsColor  
func (instance *SparklineGroup) GetSeriesColor()  (*CellsColor,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("SparklineGroup_GetSeriesColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellsColor{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellsColor) 

	return result, nil 
}
// Gets and sets the color of the sparklines in the sparkline group.
// Parameters:
//   value - CellsColor 
// Returns:
//   void  
func (instance *SparklineGroup) SetSeriesColor(value *CellsColor)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("SparklineGroup_SetSeriesColor"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the plot data is right to left.
// Returns:
//   bool  
func (instance *SparklineGroup) GetPlotRightToLeft()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SparklineGroup_GetPlotRightToLeft"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the plot data is right to left.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SparklineGroup) SetPlotRightToLeft(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SparklineGroup_SetPlotRightToLeft"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the line weight in each line sparkline in the sparkline group, in the unit of points.
// Returns:
//   float64  
func (instance *SparklineGroup) GetLineWeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("SparklineGroup_GetLineWeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the line weight in each line sparkline in the sparkline group, in the unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *SparklineGroup) SetLineWeight(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("SparklineGroup_SetLineWeight"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the color of the horizontal axis in the sparkline group.
// Returns:
//   CellsColor  
func (instance *SparklineGroup) GetHorizontalAxisColor()  (*CellsColor,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("SparklineGroup_GetHorizontalAxisColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellsColor{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellsColor) 

	return result, nil 
}
// Gets and sets the color of the horizontal axis in the sparkline group.
// Parameters:
//   value - CellsColor 
// Returns:
//   void  
func (instance *SparklineGroup) SetHorizontalAxisColor(value *CellsColor)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("SparklineGroup_SetHorizontalAxisColor"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show the sparkline horizontal axis.
// The horizontal axis appears if the sparkline has data that crosses the zero axis.
// Returns:
//   bool  
func (instance *SparklineGroup) GetShowHorizontalAxis()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SparklineGroup_GetShowHorizontalAxis"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show the sparkline horizontal axis.
// The horizontal axis appears if the sparkline has data that crosses the zero axis.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SparklineGroup) SetShowHorizontalAxis(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SparklineGroup_SetShowHorizontalAxis"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the range that contains the date values for the sparkline data.
// Returns:
//   string  
func (instance *SparklineGroup) GetHorizontalAxisDateRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("SparklineGroup_GetHorizontalAxisDateRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the range that contains the date values for the sparkline data.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *SparklineGroup) SetHorizontalAxisDateRange(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("SparklineGroup_SetHorizontalAxisDateRange"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the vertical axis maximum value type.
// Returns:
//   int32  
func (instance *SparklineGroup) GetVerticalAxisMaxValueType()  (SparklineAxisMinMaxType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("SparklineGroup_GetVerticalAxisMaxValueType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSparklineAxisMinMaxType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the vertical axis maximum value type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SparklineGroup) SetVerticalAxisMaxValueType(value SparklineAxisMinMaxType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("SparklineGroup_SetVerticalAxisMaxValueType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the custom maximum value for the vertical axis.
// Returns:
//   float64  
func (instance *SparklineGroup) GetVerticalAxisMaxValue()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("SparklineGroup_GetVerticalAxisMaxValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the custom maximum value for the vertical axis.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *SparklineGroup) SetVerticalAxisMaxValue(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("SparklineGroup_SetVerticalAxisMaxValue"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the vertical axis minimum value type.
// Returns:
//   int32  
func (instance *SparklineGroup) GetVerticalAxisMinValueType()  (SparklineAxisMinMaxType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("SparklineGroup_GetVerticalAxisMinValueType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSparklineAxisMinMaxType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the vertical axis minimum value type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SparklineGroup) SetVerticalAxisMinValueType(value SparklineAxisMinMaxType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("SparklineGroup_SetVerticalAxisMinValueType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the custom minimum value for the vertical axis.
// Returns:
//   float64  
func (instance *SparklineGroup) GetVerticalAxisMinValue()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("SparklineGroup_GetVerticalAxisMinValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the custom minimum value for the vertical axis.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *SparklineGroup) SetVerticalAxisMinValue(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("SparklineGroup_SetVerticalAxisMinValue"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteSparklineGroup(sparklinegroup *SparklineGroup){
	runtime.SetFinalizer(sparklinegroup, nil)
	C.Delete_CObject(C.CString("Delete_SparklineGroup"),sparklinegroup.ptr)
	sparklinegroup.ptr = nil
}

// Class SparklineGroupCollection 

// Encapsulates a collection of <see cref="SparklineGroup"/> objects.
type SparklineGroupCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SparklineGroupCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SparklineGroupCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="SparklineGroup"/> element at the specified index.
// Parameters:
//   index - int32 
// Returns:
//   SparklineGroup  
func (instance *SparklineGroupCollection) Get(index int32)  (*SparklineGroup,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("SparklineGroupCollection_Get"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SparklineGroup{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSparklineGroup) 

	return result, nil 
}
// Adds an <see cref="SparklineGroup"/> with a <see cref="Sparkline"/> to the collection.
// Parameters:
//   type - int32 
// Returns:
//   int32  
func (instance *SparklineGroupCollection) Add_SparklineType(type_ SparklineType)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZGM(C.CString("SparklineGroupCollection_Add_SparklineType"), instance.ptr, C.int( int32(type_)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds an <see cref="SparklineGroup"/> with <see cref="Sparkline"/> to the collection.
// Parameters:
//   type - int32 
//   dataRange - string 
//   isVertical - bool 
//   locationRange - CellArea 
// Returns:
//   int32  
func (instance *SparklineGroupCollection) Add_SparklineType_String_Bool_CellArea(type_ SparklineType, datarange string, isvertical bool, locationrange *CellArea)  (int32,  error)  {
	
	var locationrange_ptr unsafe.Pointer = nil
	if locationrange != nil {
	  locationrange_ptr =locationrange.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZND(C.CString("SparklineGroupCollection_Add_SparklineType_String_Boolean_CellArea"), instance.ptr, C.int( int32(type_)), C.CString(datarange), C.bool(isvertical), locationrange_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Clears the sparklines that is inside an area of cells.
// Parameters:
//   cellArea - CellArea 
// Returns:
//   void  
func (instance *SparklineGroupCollection) ClearSparklines(cellarea *CellArea)  error {
	
	var cellarea_ptr unsafe.Pointer = nil
	if cellarea != nil {
	  cellarea_ptr =cellarea.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("SparklineGroupCollection_ClearSparklines"), instance.ptr, cellarea_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Clears the sparkline groups that overlaps an area of cells.
// Parameters:
//   cellArea - CellArea 
// Returns:
//   void  
func (instance *SparklineGroupCollection) ClearSparklineGroups(cellarea *CellArea)  error {
	
	var cellarea_ptr unsafe.Pointer = nil
	if cellarea != nil {
	  cellarea_ptr =cellarea.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("SparklineGroupCollection_ClearSparklineGroups"), instance.ptr, cellarea_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns:
//   int32  
func (instance *SparklineGroupCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SparklineGroupCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSparklineGroupCollection(sparklinegroupcollection *SparklineGroupCollection){
	runtime.SetFinalizer(sparklinegroupcollection, nil)
	C.Delete_CObject(C.CString("Delete_SparklineGroupCollection"),sparklinegroupcollection.ptr)
	sparklinegroupcollection.ptr = nil
}

// Class TickLabelItem 

// Represents a tick label in the chart.
type TickLabelItem struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TickLabelItem) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TickLabelItem_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// X coordinates of Ticklabel item in ratio of chart width.
// Returns:
//   float64  
func (instance *TickLabelItem) GetX()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TickLabelItem_GetX"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Y coordinates of Ticklabel item in ratio of chart height.
// Returns:
//   float64  
func (instance *TickLabelItem) GetY()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TickLabelItem_GetY"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Width of Ticklabel item in ratio of chart width.
// Returns:
//   float64  
func (instance *TickLabelItem) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TickLabelItem_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Height of Ticklabel item in ratio of chart height.
// Returns:
//   float64  
func (instance *TickLabelItem) GetHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TickLabelItem_GetHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteTickLabelItem(ticklabelitem *TickLabelItem){
	runtime.SetFinalizer(ticklabelitem, nil)
	C.Delete_CObject(C.CString("Delete_TickLabelItem"),ticklabelitem.ptr)
	ticklabelitem.ptr = nil
}

// Class TickLabels 

// Represents the tick-mark labels associated with tick marks on a chart axis.
type TickLabels struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TickLabels) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TickLabels_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns a <see cref="Font"/> object that represents the font of the specified TickLabels object.
// Returns:
//   Font  
func (instance *TickLabels) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TickLabels_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Returns:
//   bool  
func (instance *TickLabels) GetAutoScaleFont()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TickLabels_GetAutoScaleFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TickLabels) SetAutoScaleFont(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TickLabels_SetAutoScaleFont"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the display mode of the background
// Returns:
//   int32  
func (instance *TickLabels) GetBackgroundMode()  (BackgroundMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TickLabels_GetBackgroundMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBackgroundMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the display mode of the background
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TickLabels) SetBackgroundMode(value BackgroundMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TickLabels_SetBackgroundMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents text rotation angle in clockwise.
// Returns:
//   int32  
func (instance *TickLabels) GetRotationAngle()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("TickLabels_GetRotationAngle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents text rotation angle in clockwise.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TickLabels) SetRotationAngle(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("TickLabels_SetRotationAngle"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the rotation angle is automatic
// Returns:
//   bool  
func (instance *TickLabels) IsAutomaticRotation()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TickLabels_IsAutomaticRotation"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the rotation angle is automatic
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TickLabels) SetIsAutomaticRotation(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TickLabels_SetIsAutomaticRotation"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the format string for the TickLabels object.
// Returns:
//   string  
func (instance *TickLabels) Get_NumberFormat()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("TickLabels_Get_NumberFormat"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the format string for the TickLabels object.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *TickLabels) SetNumberFormat(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("TickLabels_SetNumberFormat"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the format number for the TickLabels object.
// Returns:
//   int32  
func (instance *TickLabels) GetNumber()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("TickLabels_GetNumber"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the format number for the TickLabels object.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TickLabels) SetNumber(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("TickLabels_SetNumber"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the number format is linked to the cells
// (so that the number format changes in the labels when it changes in the cells).
// Returns:
//   bool  
func (instance *TickLabels) GetNumberFormatLinked()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TickLabels_GetNumberFormatLinked"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the number format is linked to the cells
// (so that the number format changes in the labels when it changes in the cells).
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TickLabels) SetNumberFormatLinked(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TickLabels_SetNumberFormatLinked"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the display number format of tick labels.
// Returns:
//   string  
func (instance *TickLabels) GetDisplayNumberFormat()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("TickLabels_GetDisplayNumberFormat"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the distance of labels from the category axis.
// Only for category (x) axis.
// Returns:
//   int32  
func (instance *TickLabels) GetOffset()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("TickLabels_GetOffset"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the distance of labels from the category axis.
// Only for category (x) axis.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TickLabels) SetOffset(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("TickLabels_SetOffset"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents text reading order.
// Returns:
//   int32  
func (instance *TickLabels) GetReadingOrder()  (TextDirectionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TickLabels_GetReadingOrder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextDirectionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents text reading order.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TickLabels) SetReadingOrder(value TextDirectionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TickLabels_SetReadingOrder"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the direction of text.
// Returns:
//   int32  
func (instance *TickLabels) GetDirectionType()  (ChartTextDirectionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TickLabels_GetDirectionType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToChartTextDirectionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the direction of text.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TickLabels) SetDirectionType(value ChartTextDirectionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TickLabels_SetDirectionType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the display tick labels of the axis.
// Returns:
//   []TickLabelItem  
func (instance *TickLabels) GetTickLabelItems()  ([]TickLabelItem,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZY(C.CString("TickLabels_GetTickLabelItems"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result:= make([]TickLabelItem, CGoReturnPtr.column_length)
	for i := 0; i < int(CGoReturnPtr.column_length); i++ {
	   offset := uintptr(C.size_t(i)) * uintptr(CGoReturnPtr.size)
	   goObject := &TickLabelItem{}
	   goObject.ptr =unsafe.Pointer(uintptr( unsafe.Pointer(CGoReturnPtr.return_value)) + offset)
	   result[i] = *goObject
	}
	 

	return result, nil 
}
// Gets and sets the text alignment for the tick labels on the axis.
// Returns:
//   int32  
func (instance *TickLabels) GetAlignmentType()  (TickLabelAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TickLabels_GetAlignmentType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTickLabelAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the text alignment for the tick labels on the axis.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TickLabels) SetAlignmentType(value TickLabelAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TickLabels_SetAlignmentType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteTickLabels(ticklabels *TickLabels){
	runtime.SetFinalizer(ticklabels, nil)
	C.Delete_CObject(C.CString("Delete_TickLabels"),ticklabels.ptr)
	ticklabels.ptr = nil
}

// Class Title 

// Encapsulates the object that represents the title of chart or axis.
type Title struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ChartTextFrame 
func NewTitle(src *ChartTextFrame) ( *Title, error) {
	title := &Title{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_Title"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		title.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(title, DeleteTitle)
		return title, nil
	} else {
		title.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return title, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *Title) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Title_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the text of display unit label.
// Returns:
//   string  
func (instance *Title) GetText()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Title_GetText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the text of display unit label.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Title) SetText(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Title_SetText"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents whether the title is visible.
// Returns:
//   bool  
func (instance *Title) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Title_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents whether the title is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Title) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Title_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Fraction of the chart area.
// X In Pixels = XRatioToChart * Chart.ChartObject.Width;
// Returns:
//   float64  
func (instance *Title) GetXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Title_GetXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Fraction of the chart area.
// X In Pixels = XRatioToChart * Chart.ChartObject.Width;
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Title) SetXRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Title_SetXRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Fraction of the chart area.
// Y In Pixels = YRatioToChart * Chart.ChartObject.Width;
// Returns:
//   float64  
func (instance *Title) GetYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Title_GetYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Fraction of the chart area.
// Y In Pixels = YRatioToChart * Chart.ChartObject.Width;
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Title) SetYRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Title_SetYRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents overlay centered title on chart without resizing chart.
// Returns:
//   bool  
func (instance *Title) GetOverLay()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Title_GetOverLay"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents overlay centered title on chart without resizing chart.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Title) SetOverLay(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Title_SetOverLay"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets rich text formatting of this Title.
// Returns:
//   []FontSetting  
func (instance *Title) Characters()  ([]FontSetting,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZY(C.CString("Title_Characters"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result:= make([]FontSetting, CGoReturnPtr.column_length)
	for i := 0; i < int(CGoReturnPtr.column_length); i++ {
	   offset := uintptr(C.size_t(i)) * uintptr(CGoReturnPtr.size)
	   goObject := &FontSetting{}
	   goObject.ptr =unsafe.Pointer(uintptr( unsafe.Pointer(CGoReturnPtr.return_value)) + offset)
	   result[i] = *goObject
	}
	 

	return result, nil 
}
// Returns a Characters object that represents a range of characters within the text.
// Parameters:
//   startIndex - int32 
//   length - int32 
// Returns:
//   FontSetting  
func (instance *Title) Characters_Int_Int(startindex int32, length int32)  (*FontSetting,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBG(C.CString("Title_Characters_Integer_Integer"), instance.ptr, C.int(startindex), C.int(length))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &FontSetting{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFontSetting) 

	return result, nil 
}
// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.
// False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
// Returns:
//   bool  
func (instance *Title) IsInnerMode()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Title_IsInnerMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.
// False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Title) SetIsInnerMode(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Title_SetIsInnerMode"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the display mode of the background
// Returns:
//   int32  
func (instance *Title) GetBackgroundMode()  (BackgroundMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Title_GetBackgroundMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBackgroundMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the display mode of the background
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Title) SetBackgroundMode(value BackgroundMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Title_SetBackgroundMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the frame has a shadow.
// Returns:
//   bool  
func (instance *Title) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Title_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the frame has a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Title) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Title_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="ShapeProperties"/> object.
// Returns:
//   ShapePropertyCollection  
func (instance *Title) GetShapeProperties()  (*ShapePropertyCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Title_GetShapeProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ShapePropertyCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteShapePropertyCollection) 

	return result, nil 
}
// Indicates whether default position(DefaultX, DefaultY, DefaultWidth and DefaultHeight) are set.
// Returns:
//   bool  
func (instance *Title) IsDefaultPosBeSet()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Title_IsDefaultPosBeSet"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents x of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *Title) GetDefaultXRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Title_GetDefaultXRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents y of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *Title) GetDefaultYRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Title_GetDefaultYRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents width of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *Title) GetDefaultWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Title_GetDefaultWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents height of default position in units of Fraction of the chart area.
// Returns:
//   float64  
func (instance *Title) GetDefaultHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Title_GetDefaultHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="Line">border</see>.
// Returns:
//   Line  
func (instance *Title) GetBorder()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Title_GetBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Gets the <see cref="Area">area</see>.
// Returns:
//   Area  
func (instance *Title) GetArea()  (*Area,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Title_GetArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Area{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteArea) 

	return result, nil 
}
// Gets and sets the options of the text.
// Returns:
//   TextOptions  
func (instance *Title) GetTextOptions()  (*TextOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Title_GetTextOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TextOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTextOptions) 

	return result, nil 
}
// Gets a <see cref="Font"/> object of the specified ChartFrame object.
// Returns:
//   Font  
func (instance *Title) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Title_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Returns:
//   bool  
func (instance *Title) GetAutoScaleFont()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Title_GetAutoScaleFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the text in the object changes font size when the object size changes. The default value is True.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Title) SetAutoScaleFont(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Title_SetAutoScaleFont"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the chart frame is automatic sized.
// Returns:
//   bool  
func (instance *Title) IsAutomaticSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Title_IsAutomaticSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the chart frame is automatic sized.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Title) SetIsAutomaticSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Title_SetIsAutomaticSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of frame in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *Title) GetWidthRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Title_GetWidthRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of frame in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Title) SetWidthRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Title_SetWidthRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the height of frame in units of ratio of the chart area.
// Returns:
//   float64  
func (instance *Title) GetHeightRatioToChart()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Title_GetHeightRatioToChart"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the height of frame in units of ratio of the chart area.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Title) SetHeightRatioToChart(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Title_SetHeightRatioToChart"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Pixel.
// Returns:
//   int32  
func (instance *Title) GetXPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Title_GetXPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the x coordinate of the upper left corner in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Title) SetXPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Title_SetXPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Pixel.
// Returns:
//   int32  
func (instance *Title) GetYPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Title_GetYPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the y coordinate of the upper left corner in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Title) SetYPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Title_SetYPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of frame in units of Pixel.
// Returns:
//   int32  
func (instance *Title) GetWidthPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Title_GetWidthPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of frame in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Title) SetWidthPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Title_SetWidthPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the height of frame in units of Pixel.
// Returns:
//   int32  
func (instance *Title) GetHeightPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Title_GetHeightPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the height of frame in units of Pixel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Title) SetHeightPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Title_SetHeightPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Set position of the frame to automatic
// Returns:
//   void  
func (instance *Title) SetPositionAuto()  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("Title_SetPositionAuto"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this data labels is deleted.
// Returns:
//   bool  
func (instance *Title) IsDeleted()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Title_IsDeleted"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this data labels is deleted.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Title) SetIsDeleted(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Title_SetIsDeleted"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the text horizontal alignment.
// Returns:
//   int32  
func (instance *Title) GetTextHorizontalAlignment()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Title_GetTextHorizontalAlignment"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the text horizontal alignment.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Title) SetTextHorizontalAlignment(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Title_SetTextHorizontalAlignment"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the text vertical alignment of text.
// Returns:
//   int32  
func (instance *Title) GetTextVerticalAlignment()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Title_GetTextVerticalAlignment"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the text vertical alignment of text.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Title) SetTextVerticalAlignment(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Title_SetTextVerticalAlignment"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents text rotation angle.
// Returns:
//   int32  
func (instance *Title) GetRotationAngle()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Title_GetRotationAngle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents text rotation angle.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Title) SetRotationAngle(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Title_SetRotationAngle"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the text of the chart is automatically rotated.
// Returns:
//   bool  
func (instance *Title) IsAutomaticRotation()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Title_IsAutomaticRotation"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents text reading order.
// Returns:
//   int32  
func (instance *Title) GetReadingOrder()  (TextDirectionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Title_GetReadingOrder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextDirectionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents text reading order.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Title) SetReadingOrder(value TextDirectionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Title_SetReadingOrder"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets whether a shape should be auto-fit to fully contain the text described within it. Auto-fitting is
// when text within a shape is scaled in order to contain all the text inside.
// Returns:
//   bool  
func (instance *Title) IsResizeShapeToFitText()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Title_IsResizeShapeToFitText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets whether a shape should be auto-fit to fully contain the text described within it. Auto-fitting is
// when text within a shape is scaled in order to contain all the text inside.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Title) SetIsResizeShapeToFitText(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Title_SetIsResizeShapeToFitText"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates the text is auto generated.
// Returns:
//   bool  
func (instance *Title) IsAutoText()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Title_IsAutoText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates the text is auto generated.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Title) SetIsAutoText(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Title_SetIsAutoText"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets a reference to the worksheet.
// Returns:
//   string  
func (instance *Title) GetLinkedSource()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Title_GetLinkedSource"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets a reference to the worksheet.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Title) SetLinkedSource(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Title_SetLinkedSource"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the direction of text.
// Returns:
//   int32  
func (instance *Title) GetDirectionType()  (ChartTextDirectionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Title_GetDirectionType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToChartTextDirectionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the direction of text.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Title) SetDirectionType(value ChartTextDirectionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Title_SetDirectionType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets a value indicating whether the text is wrapped.
// Returns:
//   bool  
func (instance *Title) IsTextWrapped()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Title_IsTextWrapped"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets a value indicating whether the text is wrapped.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Title) SetIsTextWrapped(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Title_SetIsTextWrapped"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *Title) ToChartTextFrame() *ChartTextFrame {
	parentClass := &ChartTextFrame{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *Title) ToChartFrame() *ChartFrame {
	parentClass := &ChartFrame{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteTitle(title *Title){
	runtime.SetFinalizer(title, nil)
	C.Delete_CObject(C.CString("Delete_Title"),title.ptr)
	title.ptr = nil
}

// Class Trendline 

// Represents a trendline in a chart.
type Trendline struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Line 
func NewTrendline(src *Line) ( *Trendline, error) {
	trendline := &Trendline{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_Trendline"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		trendline.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(trendline, DeleteTrendline)
		return trendline, nil
	} else {
		trendline.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return trendline, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *Trendline) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Trendline_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns if Microsoft Excel automatically determines the name of the trendline.
// Returns:
//   bool  
func (instance *Trendline) IsNameAuto()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Trendline_IsNameAuto"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns if Microsoft Excel automatically determines the name of the trendline.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Trendline) SetIsNameAuto(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Trendline_SetIsNameAuto"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns the trendline type.
// Returns:
//   int32  
func (instance *Trendline) GetType()  (TrendlineType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Trendline_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTrendlineType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Returns the name of the trendline.
// Returns:
//   string  
func (instance *Trendline) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Trendline_GetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns the name of the trendline.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Trendline) SetName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Trendline_SetName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the trendline order (an integer greater than 1) when the trendline type is Polynomial.
// The order must be between 2 and 6.
// Returns:
//   int32  
func (instance *Trendline) GetOrder()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Trendline_GetOrder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the trendline order (an integer greater than 1) when the trendline type is Polynomial.
// The order must be between 2 and 6.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Trendline) SetOrder(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Trendline_SetOrder"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the period for the moving-average trendline.
// Returns:
//   int32  
func (instance *Trendline) GetPeriod()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Trendline_GetPeriod"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the period for the moving-average trendline.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Trendline) SetPeriod(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Trendline_SetPeriod"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the number of periods (or units on a scatter chart) that the trendline extends forward.
// The number of periods must be greater than or equal to zero.
// Returns:
//   float64  
func (instance *Trendline) GetForward()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Trendline_GetForward"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the number of periods (or units on a scatter chart) that the trendline extends forward.
// The number of periods must be greater than or equal to zero.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Trendline) SetForward(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Trendline_SetForward"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the number of periods (or units on a scatter chart) that the trendline extends backward.
// The number of periods must be greater than or equal to zero.
// If the chart type is column ,the number of periods must be between 0 and 0.5
// Returns:
//   float64  
func (instance *Trendline) GetBackward()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Trendline_GetBackward"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the number of periods (or units on a scatter chart) that the trendline extends backward.
// The number of periods must be greater than or equal to zero.
// If the chart type is column ,the number of periods must be between 0 and 0.5
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Trendline) SetBackward(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Trendline_SetBackward"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents if the equation for the trendline is displayed on the chart (in the same data label as the R-squared value). Setting this property to True automatically turns on data labels.
// Returns:
//   bool  
func (instance *Trendline) GetDisplayEquation()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Trendline_GetDisplayEquation"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents if the equation for the trendline is displayed on the chart (in the same data label as the R-squared value). Setting this property to True automatically turns on data labels.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Trendline) SetDisplayEquation(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Trendline_SetDisplayEquation"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents if the R-squared value of the trendline is displayed on the chart (in the same data label as the equation). Setting this property to True automatically turns on data labels.
// Returns:
//   bool  
func (instance *Trendline) GetDisplayRSquared()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Trendline_GetDisplayRSquared"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents if the R-squared value of the trendline is displayed on the chart (in the same data label as the equation). Setting this property to True automatically turns on data labels.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Trendline) SetDisplayRSquared(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Trendline_SetDisplayRSquared"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the point where the trendline crosses the value axis.
// Returns:
//   float64  
func (instance *Trendline) GetIntercept()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Trendline_GetIntercept"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the point where the trendline crosses the value axis.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Trendline) SetIntercept(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Trendline_SetIntercept"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the DataLabels object for the specified series.
// Returns:
//   DataLabels  
func (instance *Trendline) GetDataLabels()  (*DataLabels,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Trendline_GetDataLabels"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DataLabels{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDataLabels) 

	return result, nil 
}
// Gets the legend entry according to this trendline
// Returns:
//   LegendEntry  
func (instance *Trendline) GetLegendEntry()  (*LegendEntry,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Trendline_GetLegendEntry"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &LegendEntry{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLegendEntry) 

	return result, nil 
}
// Specifies the compound line type
// Returns:
//   int32  
func (instance *Trendline) GetCompoundType()  (MsoLineStyle,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Trendline_GetCompoundType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMsoLineStyle(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the compound line type
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Trendline) SetCompoundType(value MsoLineStyle)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Trendline_SetCompoundType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the dash line type
// Returns:
//   int32  
func (instance *Trendline) GetDashType()  (MsoLineDashStyle,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Trendline_GetDashType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMsoLineDashStyle(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the dash line type
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Trendline) SetDashType(value MsoLineDashStyle)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Trendline_SetDashType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the ending caps.
// Returns:
//   int32  
func (instance *Trendline) GetCapType()  (LineCapType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Trendline_GetCapType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToLineCapType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the ending caps.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Trendline) SetCapType(value LineCapType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Trendline_SetCapType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the joining caps.
// Returns:
//   int32  
func (instance *Trendline) GetJoinType()  (LineJoinType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Trendline_GetJoinType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToLineJoinType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the joining caps.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Trendline) SetJoinType(value LineJoinType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Trendline_SetJoinType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies an arrowhead for the begin of a line.
// Returns:
//   int32  
func (instance *Trendline) GetBeginType()  (MsoArrowheadStyle,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Trendline_GetBeginType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMsoArrowheadStyle(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies an arrowhead for the begin of a line.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Trendline) SetBeginType(value MsoArrowheadStyle)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Trendline_SetBeginType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies an arrowhead for the end of a line.
// Returns:
//   int32  
func (instance *Trendline) GetEndType()  (MsoArrowheadStyle,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Trendline_GetEndType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMsoArrowheadStyle(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies an arrowhead for the end of a line.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Trendline) SetEndType(value MsoArrowheadStyle)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Trendline_SetEndType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the length of the arrowhead for the begin of a line.
// Returns:
//   int32  
func (instance *Trendline) GetBeginArrowLength()  (MsoArrowheadLength,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Trendline_GetBeginArrowLength"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMsoArrowheadLength(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the length of the arrowhead for the begin of a line.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Trendline) SetBeginArrowLength(value MsoArrowheadLength)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Trendline_SetBeginArrowLength"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the length of the arrowhead for the end of a line.
// Returns:
//   int32  
func (instance *Trendline) GetEndArrowLength()  (MsoArrowheadLength,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Trendline_GetEndArrowLength"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMsoArrowheadLength(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the length of the arrowhead for the end of a line.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Trendline) SetEndArrowLength(value MsoArrowheadLength)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Trendline_SetEndArrowLength"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the width of the arrowhead for the begin of a line.
// Returns:
//   int32  
func (instance *Trendline) GetBeginArrowWidth()  (MsoArrowheadWidth,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Trendline_GetBeginArrowWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMsoArrowheadWidth(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the width of the arrowhead for the begin of a line.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Trendline) SetBeginArrowWidth(value MsoArrowheadWidth)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Trendline_SetBeginArrowWidth"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the width of the arrowhead for the end of a line.
// Returns:
//   int32  
func (instance *Trendline) GetEndArrowWidth()  (MsoArrowheadWidth,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Trendline_GetEndArrowWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMsoArrowheadWidth(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the width of the arrowhead for the end of a line.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Trendline) SetEndArrowWidth(value MsoArrowheadWidth)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Trendline_SetEndArrowWidth"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the theme color.
// Returns:
//   ThemeColor  
func (instance *Trendline) GetThemeColor()  (*ThemeColor,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Trendline_GetThemeColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ThemeColor{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteThemeColor) 

	return result, nil 
}
// Gets and sets the theme color.
// Parameters:
//   value - ThemeColor 
// Returns:
//   void  
func (instance *Trendline) SetThemeColor(value *ThemeColor)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("Trendline_SetThemeColor"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the <see cref="Color"/> of the line.
// Returns:
//   Color  
func (instance *Trendline) GetColor()  (*Color,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAD(C.CString("Trendline_GetColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Color{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Represents the <see cref="Color"/> of the line.
// Parameters:
//   value - Color 
// Returns:
//   void  
func (instance *Trendline) SetColor(value *Color)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAE(C.CString("Trendline_SetColor"), instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the degree of transparency of the line as a value from 0.0 (opaque) through 1.0 (clear).
// Returns:
//   float64  
func (instance *Trendline) GetTransparency()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Trendline_GetTransparency"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the degree of transparency of the line as a value from 0.0 (opaque) through 1.0 (clear).
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Trendline) SetTransparency(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Trendline_SetTransparency"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the style of the line.
// Returns:
//   int32  
func (instance *Trendline) GetStyle()  (LineType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Trendline_GetStyle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToLineType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the style of the line.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Trendline) SetStyle(value LineType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Trendline_SetStyle"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the <see cref="WeightType"/> of the line.
// Returns:
//   int32  
func (instance *Trendline) GetWeight()  (WeightType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Trendline_GetWeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToWeightType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the <see cref="WeightType"/> of the line.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Trendline) SetWeight(value WeightType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Trendline_SetWeight"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the weight of the line in unit of points.
// Returns:
//   float64  
func (instance *Trendline) GetWeightPt()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Trendline_GetWeightPt"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the weight of the line in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Trendline) SetWeightPt(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Trendline_SetWeightPt"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the weight of the line in unit of pixels.
// Returns:
//   float64  
func (instance *Trendline) GetWeightPx()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Trendline_GetWeightPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the weight of the line in unit of pixels.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Trendline) SetWeightPx(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Trendline_SetWeightPx"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets format type.
// Returns:
//   int32  
func (instance *Trendline) GetFormattingType()  (ChartLineFormattingType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Trendline_GetFormattingType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToChartLineFormattingType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets format type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Trendline) SetFormattingType(value ChartLineFormattingType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Trendline_SetFormattingType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the color of line is automatic assigned.
// Returns:
//   bool  
func (instance *Trendline) IsAutomaticColor()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Trendline_IsAutomaticColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents whether the line is visible.
// Returns:
//   bool  
func (instance *Trendline) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Trendline_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents whether the line is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Trendline) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Trendline_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this line style is auto assigned.
// Returns:
//   bool  
func (instance *Trendline) IsAuto()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Trendline_IsAuto"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this line style is auto assigned.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Trendline) SetIsAuto(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Trendline_SetIsAuto"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents gradient fill.
// Returns:
//   GradientFill  
func (instance *Trendline) GetGradientFill()  (*GradientFill,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Trendline_GetGradientFill"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &GradientFill{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteGradientFill) 

	return result, nil 
}


func (instance *Trendline) ToLine() *Line {
	parentClass := &Line{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteTrendline(trendline *Trendline){
	runtime.SetFinalizer(trendline, nil)
	C.Delete_CObject(C.CString("Delete_Trendline"),trendline.ptr)
	trendline.ptr = nil
}

// Class TrendlineCollection 

// Represents a collection of all the <see cref="Trendline"/> objects for the specified data series.
type TrendlineCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TrendlineCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TrendlineCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a <see cref="Trendline"/> object to this collection with specified type.
// Parameters:
//   type - int32 
// Returns:
//   int32  
func (instance *TrendlineCollection) Add_TrendlineType(type_ TrendlineType)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZGM(C.CString("TrendlineCollection_Add_TrendlineType"), instance.ptr, C.int( int32(type_)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a <see cref="Trendline"/> object to this collection with specified type and name.
// Parameters:
//   type - int32 
//   name - string 
// Returns:
//   int32  
func (instance *TrendlineCollection) Add_TrendlineType_String(type_ TrendlineType, name string)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZEX(C.CString("TrendlineCollection_Add_TrendlineType_String"), instance.ptr, C.int( int32(type_)), C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets a <see cref="Trendline"/> object by its index.
// Parameters:
//   index - int32 
// Returns:
//   Trendline  
func (instance *TrendlineCollection) Get(index int32)  (*Trendline,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("TrendlineCollection_Get"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Trendline{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTrendline) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *TrendlineCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("TrendlineCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteTrendlineCollection(trendlinecollection *TrendlineCollection){
	runtime.SetFinalizer(trendlinecollection, nil)
	C.Delete_CObject(C.CString("Delete_TrendlineCollection"),trendlinecollection.ptr)
	trendlinecollection.ptr = nil
}

// Class Walls 

// Encapsulates the object that represents the walls of a 3-D chart.
type Walls struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Floor 
func NewWalls(src *Floor) ( *Walls, error) {
	walls := &Walls{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_Walls"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		walls.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(walls, DeleteWalls)
		return walls, nil
	} else {
		walls.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return walls, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *Walls) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Walls_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the x coordinate of the left-bottom corner of Wall center in units of 1/4000 of chart's width after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *Walls) GetCenterX()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Walls_GetCenterX"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the y coordinate of the left-bottom corner of Wall center in units of 1/4000 of chart's height after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *Walls) GetCenterY()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Walls_GetCenterY"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the width of left to right in units of 1/4000 of chart's width after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *Walls) GetWidth()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Walls_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the depth front to back in units of 1/4000 of chart's width after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *Walls) GetDepth()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Walls_GetDepth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the height of top to bottom in units of 1/4000 of chart's height after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *Walls) GetHeight()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Walls_GetHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the x coordinate of the left-bottom corner of Wall center in units of pixels after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *Walls) GetCenterXPx()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Walls_GetCenterXPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the y coordinate of the left-bottom corner of Wall center in units of pixels after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *Walls) GetCenterYPx()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Walls_GetCenterYPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the width of left to right in units of pixels after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *Walls) GetWidthPx()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Walls_GetWidthPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the depth front to back in units of pixels after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *Walls) GetDepthPx()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Walls_GetDepthPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the height of top to bottom in units of pixels after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *Walls) GetHeightPx()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Walls_GetHeightPx"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the number of cube points after calls Chart.Calculate() method.
// Returns:
//   int32  
func (instance *Walls) GetCubePointCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Walls_GetCubePointCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets x-coordinate of the apex point of walls cube after calls Chart.Calculate() method.
// The number of apex points of walls cube is eight
// Parameters:
//   index - int32 
// Returns:
//   float32  
func (instance *Walls) GetCubePointXPx(index int32)  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMX(C.CString("Walls_GetCubePointXPx"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets y-coordinate of the apex point of walls cube after calls Chart.Calculate() method.
// The number of apex points of walls cube is eight.
// Parameters:
//   index - int32 
// Returns:
//   float32  
func (instance *Walls) GetCubePointYPx(index int32)  (float32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMX(C.CString("Walls_GetCubePointYPx"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the background <see cref="Color"/> of the <see cref="Area"/>.
// Returns:
//   Color  
func (instance *Walls) GetBackgroundColor()  (*Color,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAD(C.CString("Walls_GetBackgroundColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Color{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets or sets the background <see cref="Color"/> of the <see cref="Area"/>.
// Parameters:
//   value - Color 
// Returns:
//   void  
func (instance *Walls) SetBackgroundColor(value *Color)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAE(C.CString("Walls_SetBackgroundColor"), instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the foreground <see cref="Color"/>.
// Returns:
//   Color  
func (instance *Walls) GetForegroundColor()  (*Color,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAD(C.CString("Walls_GetForegroundColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Color{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets or sets the foreground <see cref="Color"/>.
// Parameters:
//   value - Color 
// Returns:
//   void  
func (instance *Walls) SetForegroundColor(value *Color)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAE(C.CString("Walls_SetForegroundColor"), instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the formatting of the area.
// Returns:
//   int32  
func (instance *Walls) GetFormatting()  (FormattingType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Walls_GetFormatting"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToFormattingType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the formatting of the area.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Walls) SetFormatting(value FormattingType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Walls_SetFormatting"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// If the property is true and the value of chart point is a negative number,
// the foreground color and background color will be exchanged.
// Returns:
//   bool  
func (instance *Walls) GetInvertIfNegative()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Walls_GetInvertIfNegative"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// If the property is true and the value of chart point is a negative number,
// the foreground color and background color will be exchanged.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Walls) SetInvertIfNegative(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Walls_SetInvertIfNegative"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents a <see cref="FillFormat"/> object that contains fill formatting properties for the specified chart or shape.
// Returns:
//   FillFormat  
func (instance *Walls) GetFillFormat()  (*FillFormat,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Walls_GetFillFormat"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &FillFormat{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFillFormat) 

	return result, nil 
}
// Returns or sets the degree of transparency of the area as a value from 0.0 (opaque) through 1.0 (clear).
// Returns:
//   float64  
func (instance *Walls) GetTransparency()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Walls_GetTransparency"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the degree of transparency of the area as a value from 0.0 (opaque) through 1.0 (clear).
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Walls) SetTransparency(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Walls_SetTransparency"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the border <see cref="Line"/>.
// Returns:
//   Line  
func (instance *Walls) GetBorder()  (*Line,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Walls_GetBorder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Line{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLine) 

	return result, nil 
}
// Gets or sets the border <see cref="Line"/>.
// Parameters:
//   value - Line 
// Returns:
//   void  
func (instance *Walls) SetBorder(value *Line)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("Walls_SetBorder"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *Walls) ToFloor() *Floor {
	parentClass := &Floor{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *Walls) ToArea() *Area {
	parentClass := &Area{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteWalls(walls *Walls){
	runtime.SetFinalizer(walls, nil)
	C.Delete_CObject(C.CString("Delete_Walls"),walls.ptr)
	walls.ptr = nil
}
