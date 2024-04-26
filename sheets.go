package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)


type SynStudent struct {
    Classname   string
    SectionId   string
    Name        string
    Email       string
    ParentName  string
    PEmails     [3]*string
}


func googleSheets() *sheets.Service {
    ctx := context.Background()
    b, err := os.ReadFile("client_secret.json")
    if err != nil {
        log.Fatalf("Unable to read client secret file: %v", err)
    }

    // If modifying these scopes, delete your previously saved token.json.
    config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
    if err != nil {
        log.Fatalf("Unable to parse client secret file to config: %v", err)
    }
    client := getClient(config, "sheets_token.json")

    srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
    if err != nil {
        log.Fatalf("Unable to retrieve Sheets client: %v", err)
    }

    return srv

    /*
    // Prints the names and majors of students in a sample spreadsheet:
    // https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
    spreadsheetId := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
    readRange := "Class Data!A2:E"
    resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
    if err != nil {
        log.Fatalf("Unable to retrieve data from sheet: %v", err)
    }

    if len(resp.Values) == 0 {
        fmt.Println("No data found.")
    } else {
        fmt.Println("Name, Major:")
        for _, row := range resp.Values {
            // Print columns A and E, which correspond to indices 0 and 4.
            fmt.Printf("%s, %s\n", row[0], row[4])
        }
    }
    */
}

const A int = 0
const B int = 1
const C int = 2
const D int = 3
const E int = 4
const F int = 5
const G int = 6
const H int = 7

func getStudentInfo(srv *sheets.Service, sheetId string) *map[string]SynStudent {
    readRange := "Sheet1!A:H"
    resp, err := srv.Spreadsheets.Values.Get(sheetId, readRange).Do()
    if err != nil {
        log.Fatalf("Unable to retrieve data from sheet: %v", err)
    }

    students := make(map[string]SynStudent)
    if len(resp.Values) == 0 {
        fmt.Println("No data found.")
    } else {
        fmt.Println("Name, Major:")
        for _, row := range resp.Values {
            var pemails [3]*string
            for i := F; i <= G; i++ {
                email := row[i].(string)
                var emailPtr *string
                if row[i].(string) != "" {
                    emailPtr = &email
                }
                
                pemails[i-F] = emailPtr
            }
            student := SynStudent {
                Classname: row[A].(string),
                SectionId: row[B].(string),
                Name: row[C].(string),
                Email: row[D].(string),
                ParentName: row[E].(string),
                PEmails: pemails,
            }

            students[student.Name] = student
        }
    }

    return &students
}

func updateContactLog(
    srv *sheets.Service,
    record *Record,
    parentName string,
    parentEmail string,
) {
    id := os.Getenv("CONTACT_LOG_ID")
    values := [][]interface{} {
        { record.Student, record.Assignment, parentName, parentEmail, time.Now().UTC().Format("MM/DD/YYYY") },
    }

    srv.Spreadsheets.Values.Append(id, "Sheet1", &sheets.ValueRange {
        MajorDimension: "ROWS",
        Values: values,
    }).Do()
}






























