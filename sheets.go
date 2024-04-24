package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"

    "golang.org/x/oauth2"
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
    PEmails     [3]string
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
    // The file token.json stores the user's access and refresh tokens, and is
    // created automatically when the authorization flow completes for the first
    // time.
    tokFile := "token.json"
    tok, err := tokenFromFile(tokFile)
    if err != nil {
        tok = getTokenFromWeb(config)
        saveToken(tokFile, tok)
    }
    return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
    authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
    fmt.Printf("Go to the following link in your browser then type the "+
        "authorization code: \n%v\n", authURL)

    var authCode string
    if _, err := fmt.Scan(&authCode); err != nil {
        log.Fatalf("Unable to read authorization code: %v", err)
    }

    tok, err := config.Exchange(context.TODO(), authCode)
    if err != nil {
        log.Fatalf("Unable to retrieve token from web: %v", err)
    }
    return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
    f, err := os.Open(file)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    tok := &oauth2.Token{}
    err = json.NewDecoder(f).Decode(tok)
    return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
    fmt.Printf("Saving credential file to: %s\n", path)
    f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
    if err != nil {
        log.Fatalf("Unable to cache oauth token: %v", err)
    }
    defer f.Close()
    json.NewEncoder(f).Encode(token)
}

func googleSheets() *sheets.Service {
    ctx := context.Background()
    b, err := os.ReadFile("client_secret.json")
    if err != nil {
        log.Fatalf("Unable to read client secret file: %v", err)
    }

    // If modifying these scopes, delete your previously saved token.json.
    config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
    if err != nil {
        log.Fatalf("Unable to parse client secret file to config: %v", err)
    }
    client := getClient(config)

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
            student := SynStudent{
                Classname: row[A].(string),
                SectionId: row[B].(string),
                Name: row[C].(string),
                Email: row[D].(string),
                ParentName: row[E].(string),
                PEmails: [3]string{row[F].(string), row[G].(string), row[H].(string)},
            }

            students[student.Name] = student
        }
    }

    return &students
}
































