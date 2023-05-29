package main

import (
	"context"
	"log"

	"github.com/gibs952/gdocsheets/client"
	"github.com/gibs952/gdocsheets/config"
	"github.com/gibs952/gdocsheets/gsheets"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	ctx := context.Background()
        filePath := "credentials.json"//modify the placehholder with the credentials file
        conf, err := config.NewConfigGoogle(filePath, "https://www.googleapis.com/auth/spreadsheets")//modify the scope will need to delete the previous token file
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }
		client := client.GetClient(conf, "token.json") //Change the Placeholder with the token file

        srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
        if err != nil {
                log.Fatalf("Unable to retrieve Sheets client: %v", err)
        }

		
	gw := &sheets.ValueRange{
		Range: "Sheet1!A1:D9",//see https://developers.google.com/sheets/api/guides/concepts#cell
		MajorDimension: "COLUMNS",// see https://developers.google.com/sheets/api/reference/rest/v4/Dimension
		Values: [][]interface{}{
			{"Melinda", "meraih", "mimpi"},
			{"kamu", "meraih", "pahala"},
			{true, false, true},
		// TODO: Add desired fields of the request body. All existing fields
		// will be replaced.
	}}
	//gsheets.Buat(srv, "1aYmUVf7ChU2eVBJY3584lja_SrMmDNtqSaMOeovJrqw",gw.Range, gw)
	gsheets.Baca(srv, "1aYmUVf7ChU2eVBJY3584lja_SrMmDNtqSaMOeovJrqw", gw.Range, "FORMATTED_VALUE","FORMATTED_STRING")
	//gsheets.Update(srv, "1aYmUVf7ChU2eVBJY3584lja_SrMmDNtqSaMOeovJrqw", gw.Range, gw)
	//gsheets.Hapus(srv, "1aYmUVf7ChU2eVBJY3584lja_SrMmDNtqSaMOeovJrqw", gw.Range)
}