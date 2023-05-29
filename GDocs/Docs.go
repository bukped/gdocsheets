package gdocs

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/api/docs/v1"
)

func BuatFile(namaFile string) error {
	file, err := os.Create(namaFile)
	if err != nil {
		return fmt.Errorf("gagal membuat file: %v", err)
	}
	defer file.Close()

	// Lakukan operasi lain terkait pembuatan file

	return nil
}

// BuatDirektori adalah fungsi untuk membuat direktori baru
func BuatDirektori(namaDirektori string) error {
	err := os.Mkdir(namaDirektori, 0755)
	if err != nil {
		return fmt.Errorf("gagal membuat direktori: %v", err)
	}
	return nil
}

func DeleteDocument(srv *docs.Service,documentID, credentialsFile string) error {
	// ctx := context.Background()

	// srv, err := docs.NewService(ctx, option.WithCredentialsFile(credentialsFile))
	// if err != nil {
	// 	return fmt.Errorf("gagal membuat layanan Google Docs: %v", err)
	// }

	req := docs.BatchUpdateDocumentRequest{
		Requests: []*docs.Request{
			{
				DeleteContentRange: &docs.DeleteContentRangeRequest{
					Range: &docs.Range{
						StartIndex: 1, // Menghapus semua isi dokumen
					},
				},
			},
		},
	}

	_,err := srv.Documents.BatchUpdate(documentID, &req).Do()
	if err != nil {
		return fmt.Errorf("gagal menghapus dokumen: %v", err)
	}

	return nil
}

func Filebaru(srv *docs.Service, docsID string) {
	// ctx := context.Background()
	// b, err := os.ReadFile("credentials.json")
	// if err != nil {
	// 	log.Fatalf("Unable to read client secret file: %v", err)
	// }

	// // If modifying these scopes, delete your previously saved token.json.
	// config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/documents.readonly")
	// if err != nil {
	// 	log.Fatalf("Unable to parse client secret file to config: %v", err)
	// }
	// client := client.GetClient(config)

	// srv, err := docs.NewService(ctx, option.WithHTTPClient(client))
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve Docs client: %v", err)
	// }

	// Prints the title of the requested doc:
	// https://docs.google.com/document/d/195j9eDD3ccgjQRttHhJPymLJUCOUjs-jmwTrekvdjFE/edit
	docId := docsID
	doc, err := srv.Documents.Get(docId).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from document: %v", err)
	}
	fmt.Printf("The title of the doc is: %s\n", doc.Title)
}
