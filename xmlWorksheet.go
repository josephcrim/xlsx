package xlsx

import (
	"encoding/xml"
	"strings"
)

// xlsxWorksheet directly maps the worksheet element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxWorksheet struct {
	XMLName         xml.Name                 `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main worksheet"`
	SheetPr         xlsxSheetPr              `xml:"sheetPr"`
	Dimension       xlsxDimension            `xml:"dimension"`
	SheetViews      xlsxSheetViews           `xml:"sheetViews"`
	SheetFormatPr   xlsxSheetFormatPr        `xml:"sheetFormatPr"`
	Cols            *xlsxCols                `xml:"cols,omitempty"`
	SheetData       xlsxSheetData            `xml:"sheetData"`
	DataValidations *xlsxCellDataValidations `xml:"dataValidations"`
	AutoFilter      *xlsxAutoFilter          `xml:"autoFilter,omitempty"`
	MergeCells      *xlsxMergeCells          `xml:"mergeCells,omitempty"`
	PrintOptions    xlsxPrintOptions         `xml:"printOptions"`
	PageMargins     xlsxPageMargins          `xml:"pageMargins"`
	PageSetUp       xlsxPageSetUp            `xml:"pageSetup"`
	HeaderFooter    xlsxHeaderFooter         `xml:"headerFooter"`
}

// xlsxHeaderFooter directly maps the headerFooter element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxHeaderFooter struct {
	DifferentFirst   bool            `xml:"differentFirst,attr"`
	DifferentOddEven bool            `xml:"differentOddEven,attr"`
	OddHeader        []xlsxOddHeader `xml:"oddHeader"`
	OddFooter        []xlsxOddFooter `xml:"oddFooter"`
}

// xlsxOddHeader directly maps the oddHeader element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxOddHeader struct {
	Content string `xml:",chardata"`
}

// xlsxOddFooter directly maps the oddFooter element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxOddFooter struct {
	Content string `xml:",chardata"`
}

// xlsxPageSetUp directly maps the pageSetup element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxPageSetUp struct {
	PaperSize          string  `xml:"paperSize,attr"`
	Scale              int     `xml:"scale,attr"`
	FirstPageNumber    int     `xml:"firstPageNumber,attr"`
	FitToWidth         int     `xml:"fitToWidth,attr"`
	FitToHeight        int     `xml:"fitToHeight,attr"`
	PageOrder          string  `xml:"pageOrder,attr"`
	Orientation        string  `xml:"orientation,attr"`
	UsePrinterDefaults bool    `xml:"usePrinterDefaults,attr"`
	BlackAndWhite      bool    `xml:"blackAndWhite,attr"`
	Draft              bool    `xml:"draft,attr"`
	CellComments       string  `xml:"cellComments,attr"`
	UseFirstPageNumber bool    `xml:"useFirstPageNumber,attr"`
	HorizontalDPI      float32 `xml:"horizontalDpi,attr"`
	VerticalDPI        float32 `xml:"verticalDpi,attr"`
	Copies             int     `xml:"copies,attr"`
}

// xlsxPrintOptions directly maps the printOptions element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxPrintOptions struct {
	Headings           bool `xml:"headings,attr"`
	GridLines          bool `xml:"gridLines,attr"`
	GridLinesSet       bool `xml:"gridLinesSet,attr"`
	HorizontalCentered bool `xml:"horizontalCentered,attr"`
	VerticalCentered   bool `xml:"verticalCentered,attr"`
}

// xlsxPageMargins directly maps the pageMargins element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxPageMargins struct {
	Left   float64 `xml:"left,attr"`
	Right  float64 `xml:"right,attr"`
	Top    float64 `xml:"top,attr"`
	Bottom float64 `xml:"bottom,attr"`
	Header float64 `xml:"header,attr"`
	Footer float64 `xml:"footer,attr"`
}

// xlsxSheetFormatPr directly maps the sheetFormatPr element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxSheetFormatPr struct {
	DefaultColWidth  float64 `xml:"defaultColWidth,attr,omitempty"`
	DefaultRowHeight float64 `xml:"defaultRowHeight,attr"`
	OutlineLevelCol  uint8   `xml:"outlineLevelCol,attr,omitempty"`
	OutlineLevelRow  uint8   `xml:"outlineLevelRow,attr,omitempty"`
}

// xlsxSheetViews directly maps the sheetViews element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxSheetViews struct {
	SheetView []xlsxSheetView `xml:"sheetView"`
}

// xlsxSheetView directly maps the sheetView element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxSheetView struct {
	WindowProtection        bool            `xml:"windowProtection,attr"`
	ShowFormulas            bool            `xml:"showFormulas,attr"`
	ShowGridLines           bool            `xml:"showGridLines,attr"`
	ShowRowColHeaders       bool            `xml:"showRowColHeaders,attr"`
	ShowZeros               bool            `xml:"showZeros,attr"`
	RightToLeft             bool            `xml:"rightToLeft,attr"`
	TabSelected             bool            `xml:"tabSelected,attr"`
	ShowOutlineSymbols      bool            `xml:"showOutlineSymbols,attr"`
	DefaultGridColor        bool            `xml:"defaultGridColor,attr"`
	View                    string          `xml:"view,attr"`
	TopLeftCell             string          `xml:"topLeftCell,attr"`
	ColorId                 int             `xml:"colorId,attr"`
	ZoomScale               float64         `xml:"zoomScale,attr"`
	ZoomScaleNormal         float64         `xml:"zoomScaleNormal,attr"`
	ZoomScalePageLayoutView float64         `xml:"zoomScalePageLayoutView,attr"`
	WorkbookViewId          int             `xml:"workbookViewId,attr"`
	Pane                    *xlsxPane       `xml:"pane"`
	Selection               []xlsxSelection `xml:"selection"`
}

// xlsxSelection directly maps the selection element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxSelection struct {
	Pane         string `xml:"pane,attr"`
	ActiveCell   string `xml:"activeCell,attr"`
	ActiveCellId int    `xml:"activeCellId,attr"`
	SQRef        string `xml:"sqref,attr"`
}

// xlsxSelection directly maps the selection element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxPane struct {
	XSplit      float64 `xml:"xSplit,attr"`
	YSplit      float64 `xml:"ySplit,attr"`
	TopLeftCell string  `xml:"topLeftCell,attr"`
	ActivePane  string  `xml:"activePane,attr"`
	State       string  `xml:"state,attr"` // Either "split" or "frozen"
}

// xlsxSheetPr directly maps the sheetPr element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxSheetPr struct {
	FilterMode  bool              `xml:"filterMode,attr"`
	PageSetUpPr []xlsxPageSetUpPr `xml:"pageSetUpPr"`
}

// xlsxPageSetUpPr directly maps the pageSetupPr element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxPageSetUpPr struct {
	FitToPage bool `xml:"fitToPage,attr"`
}

// xlsxCols directly maps the cols element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxCols struct {
	Col []xlsxCol `xml:"col"`
}

// xlsxCol directly maps the col element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxCol struct {
	Collapsed    bool    `xml:"collapsed,attr"`
	Hidden       bool    `xml:"hidden,attr"`
	Max          int     `xml:"max,attr"`
	Min          int     `xml:"min,attr"`
	Style        int     `xml:"style,attr"`
	Width        float64 `xml:"width,attr"`
	CustomWidth  bool    `xml:"customWidth,attr,omitempty"`
	OutlineLevel uint8   `xml:"outlineLevel,attr,omitempty"`
}

// xlsxDimension directly maps the dimension element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxDimension struct {
	Ref string `xml:"ref,attr"`
}

// xlsxSheetData directly maps the sheetData element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxSheetData struct {
	XMLName xml.Name  `xml:"sheetData"`
	Row     []xlsxRow `xml:"row"`
}

// xlsxCellDataValidations  excel cell data validation
type xlsxCellDataValidations struct {
	DataValidattion []*xlsxCellDataValidation `xml:"dataValidation"`
	Count           int                       `xml:"count,attr"`
}

// xlsxCellDataValidation
// A single item of data validation defined on a range of the worksheet.
// The list validation type would more commonly be called "a drop down box."
type xlsxCellDataValidation struct {
	// A boolean value indicating whether the data validation allows the use of empty or blank
	//entries. 1 means empty entries are OK and do not violate the validation constraints.
	AllowBlank bool `xml:"allowBlank,attr,omitempty"`
	// A boolean value indicating whether to display the input prompt message.
	ShowInputMessage bool `xml:"showInputMessage,attr,omitempty"`
	// A boolean value indicating whether to display the error alert message when an invalid
	// value has been entered, according to the criteria specified.
	ShowErrorMessage bool `xml:"showErrorMessage,attr,omitempty"`
	// The style of error alert used for this data validation.
	// warning, infomation, or stop
	// Stop will prevent the user from entering data that does not pass validation.
	ErrorStyle *string `xml:"errorStyle,attr"`
	// Title bar text of error alert.
	ErrorTitle *string `xml:"errorTitle,attr"`
	// The relational operator used with this data validation.
	// The possible values for this can be equal, notEqual, lessThan, etc.
	// This only applies to certain validation types.
	Operator string `xml:"operator,attr,omitempty"`
	// Message text of error alert.
	Error *string `xml:"error,attr"`
	// Title bar text of input prompt.
	PromptTitle *string `xml:"promptTitle,attr"`
	// Message text of input prompt.
	Prompt *string `xml:"prompt,attr"`
	// The type of data validation.
	// none, custom, date, decimal, list, textLength, time, whole
	Type string `xml:"type,attr"`
	// Range over which data validation is applied.
	// Cell or range, eg: A1 OR A1:A20
	Sqref string `xml:"sqref,attr,omitempty"`
	// The first formula in the Data Validation dropdown. It is used as a bounds for 'between' and
	// 'notBetween' relational operators, and the only formula used for other relational operators
	// (equal, notEqual, lessThan, lessThanOrEqual, greaterThan, greaterThanOrEqual), or for custom
	// or list type data validation. The content can be a formula or a constant or a list series (comma separated values).
	Formula1 string `xml:"formula1"`
	// The second formula in the DataValidation dropdown. It is used as a bounds for 'between' and
	// 'notBetween' relational operators only.
	Formula2 string `xml:"formula2,omitempty"`
	minRow   int    //`xml:"-"`
	maxRow   int    //`xml:"-"`
	//minCol         int     `xml:"-"` //spare
	//maxCol         int     `xml:"-"` //spare
}

// xlsxRow directly maps the row element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxRow struct {
	R            int     `xml:"r,attr"`
	Spans        string  `xml:"spans,attr,omitempty"`
	Hidden       bool    `xml:"hidden,attr,omitempty"`
	C            []xlsxC `xml:"c"`
	Ht           string  `xml:"ht,attr,omitempty"`
	CustomHeight bool    `xml:"customHeight,attr,omitempty"`
	OutlineLevel uint8   `xml:"outlineLevel,attr,omitempty"`
}

type xlsxAutoFilter struct {
	Ref string `xml:"ref,attr"`
}

type xlsxMergeCell struct {
	Ref string `xml:"ref,attr"` // ref: horiz "A1:C1", vert "B3:B6", both  "D3:G4"
}

type xlsxMergeCells struct {
	XMLName xml.Name        //`xml:"mergeCells,omitempty"`
	Count   int             `xml:"count,attr,omitempty"`
	Cells   []xlsxMergeCell `xml:"mergeCell,omitempty"`
}

// Return the cartesian extent of a merged cell range from its origin
// cell (the closest merged cell to the to left of the sheet.
func (mc *xlsxMergeCells) getExtent(cellRef string) (int, int, error) {
	if mc == nil {
		return 0, 0, nil
	}
	for _, cell := range mc.Cells {
		if strings.HasPrefix(cell.Ref, cellRef+cellRangeChar) {
			parts := strings.Split(cell.Ref, cellRangeChar)
			startx, starty, err := GetCoordsFromCellIDString(parts[0])
			if err != nil {
				return -1, -1, err
			}
			endx, endy, err := GetCoordsFromCellIDString(parts[1])
			if err != nil {
				return -2, -2, err
			}
			return endx - startx, endy - starty, nil
		}
	}
	return 0, 0, nil
}

// xlsxC directly maps the c element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxC struct {
	R  string  `xml:"r,attr"`           // Cell ID, e.g. A1
	S  int     `xml:"s,attr,omitempty"` // Style reference.
	T  string  `xml:"t,attr,omitempty"` // Type.
	F  *xlsxF  `xml:"f,omitempty"`      // Formula
	V  string  `xml:"v,omitempty"`      // Value
	Is *xlsxSI `xml:"is,omitempty"`     // Inline String.
}

// xlsxF directly maps the f element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxF struct {
	Content string `xml:",chardata"`
	T       string `xml:"t,attr,omitempty"`   // Formula type
	Ref     string `xml:"ref,attr,omitempty"` // Shared formula ref
	Si      int    `xml:"si,attr,omitempty"`  // Shared formula index
}

// Create a new XLSX Worksheet with default values populated.
// Strictly for internal use only!
func newXlsxWorksheet() (worksheet *xlsxWorksheet) {
	worksheet = &xlsxWorksheet{}
	worksheet.SheetPr.FilterMode = false
	worksheet.SheetPr.PageSetUpPr = make([]xlsxPageSetUpPr, 1)
	worksheet.SheetPr.PageSetUpPr[0] = xlsxPageSetUpPr{FitToPage: false}
	worksheet.SheetViews.SheetView = make([]xlsxSheetView, 1)
	worksheet.SheetViews.SheetView[0] = xlsxSheetView{
		ColorId:                 64,
		DefaultGridColor:        true,
		RightToLeft:             false,
		Selection:               make([]xlsxSelection, 1),
		ShowFormulas:            false,
		ShowGridLines:           true,
		ShowOutlineSymbols:      true,
		ShowRowColHeaders:       true,
		ShowZeros:               true,
		TabSelected:             false,
		TopLeftCell:             "A1",
		View:                    "normal",
		WindowProtection:        false,
		WorkbookViewId:          0,
		ZoomScale:               100,
		ZoomScaleNormal:         100,
		ZoomScalePageLayoutView: 100}
	worksheet.SheetViews.SheetView[0].Selection[0] = xlsxSelection{
		Pane:         "topLeft",
		ActiveCell:   "A1",
		ActiveCellId: 0,
		SQRef:        "A1"}
	worksheet.SheetFormatPr.DefaultRowHeight = 12.85
	worksheet.PrintOptions.Headings = false
	worksheet.PrintOptions.GridLines = false
	worksheet.PrintOptions.GridLinesSet = true
	worksheet.PrintOptions.HorizontalCentered = false
	worksheet.PrintOptions.VerticalCentered = false
	worksheet.PageMargins.Left = 0.7875
	worksheet.PageMargins.Right = 0.7875
	worksheet.PageMargins.Top = 1.05277777777778
	worksheet.PageMargins.Bottom = 1.05277777777778
	worksheet.PageMargins.Header = 0.7875
	worksheet.PageMargins.Footer = 0.7875
	worksheet.PageSetUp.PaperSize = "9"
	worksheet.PageSetUp.Scale = 100
	worksheet.PageSetUp.FirstPageNumber = 1
	worksheet.PageSetUp.FitToWidth = 1
	worksheet.PageSetUp.FitToHeight = 1
	worksheet.PageSetUp.PageOrder = "downThenOver"
	worksheet.PageSetUp.Orientation = "portrait"
	worksheet.PageSetUp.UsePrinterDefaults = false
	worksheet.PageSetUp.BlackAndWhite = false
	worksheet.PageSetUp.Draft = false
	worksheet.PageSetUp.CellComments = "none"
	worksheet.PageSetUp.UseFirstPageNumber = true
	worksheet.PageSetUp.HorizontalDPI = 300
	worksheet.PageSetUp.VerticalDPI = 300
	worksheet.PageSetUp.Copies = 1
	worksheet.HeaderFooter.OddHeader = make([]xlsxOddHeader, 1)
	worksheet.HeaderFooter.OddHeader[0] = xlsxOddHeader{Content: `&C&"Times New Roman,Regular"&12&A`}
	worksheet.HeaderFooter.OddFooter = make([]xlsxOddFooter, 1)
	worksheet.HeaderFooter.OddFooter[0] = xlsxOddFooter{Content: `&C&"Times New Roman,Regular"&12Page &P`}

	return
}
