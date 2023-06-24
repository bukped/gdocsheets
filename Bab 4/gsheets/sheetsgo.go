package gsheets

import (
	"fmt"
	"log"

	"google.golang.org/api/sheets/v4"
)

func Buat(srv *sheets.Service, IdSpreadsheet string, SheetRange string, gw *sheets.ValueRange) {

	// The ID of the spreadsheet to update.
	spreadsheetId := IdSpreadsheet // TODO: Update placeholder value.

	// The A1 notation of a range to search for a logical table of data.
	range2 := SheetRange // TODO: Update placeholder value.

	// How the input data should be interpreted.
	valueInputOption := "USER_ENTERED" // TODO: Update placeholder value.

	// How the input data should be inserted.
	insertDataOption := "INSERT_ROWS" // TODO: Update placeholder value.

	rb := gw

	resp, err := srv.Spreadsheets.Values.Append(spreadsheetId, range2, rb).ValueInputOption(valueInputOption).InsertDataOption(insertDataOption).Do()
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Change code below to process the `resp` object:
	fmt.Println(resp)
}

func Baca(srv *sheets.Service, IdSpreadhseet string, SheetRange string /*Using A1 Notation*/, valueouttype string, dateouttype string) {
	
	// The ID of the spreadsheet to retrieve data from.
	spreadsheetId := IdSpreadhseet // TODO: Update placeholder value.

	// The A1 notation of the values to retrieve.
	range2 := SheetRange // TODO: Update placeholder value.

	// How values should be represented in the output.
	// The default render option is ValueRenderOption.FORMATTED_VALUE.
	valueRenderOption := valueouttype
	// The default dateTime render option is [DateTimeRenderOption.SERIAL_NUMBER].
	dateTimeRenderOption := dateouttype // TODO: Update placeholder value.

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, range2).ValueRenderOption(valueRenderOption).DateTimeRenderOption(dateTimeRenderOption).Do()
	if err != nil {
			log.Fatal(err)
	}

	// TODO: Change code below to process the `resp` object:
	fmt.Println(resp)
}
func Update(srv *sheets.Service, IdSpreadsheet string, SheetRange string /*Using A1 Notation*/, gw *sheets.ValueRange) {
	spreadsheetId := IdSpreadsheet
	// The A1 notation of the values to update.
	range2 := SheetRange // TODO: Update placeholder value.
	// How the input data should be interpreted.
	valueInputOption := "USER_ENTERED"
	rb := gw
	resp, err := srv.Spreadsheets.Values.Update(spreadsheetId, range2, rb).ValueInputOption(valueInputOption).Do()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Change code below to process the `resp` object:
	fmt.Print(resp)
}

func Hapus(srv *sheets.Service, IdSpreadhseet string, SheetRange string) {
	// The ID of the spreadsheet to update.
	spreadsheetId := IdSpreadhseet // TODO: Update placeholder value.

	// The A1 notation of the values to clear.
	range2 := SheetRange // TODO: Update placeholder value.

	rb := &sheets.ClearValuesRequest{
		// TODO: Add desired fields of the request body.
	}

	resp, err := srv.Spreadsheets.Values.Clear(spreadsheetId, range2, rb).Do()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Change code below to process the `resp` object:
	fmt.Println(resp)
}