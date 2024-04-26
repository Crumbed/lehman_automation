package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"log"
	"os"

	"github.com/go-gomail/gomail"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)








func getGmail() *gmail.Service {
    ctx := context.Background()
    b, err := os.ReadFile("client_secret.json")
    if err != nil {
        log.Fatalf("Unable to read client secret file: %v", err)
    }

    // If modifying these scopes, delete your previously saved token.json.
    config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
    if err != nil {
        log.Fatalf("Unable to parse client secret file to config: %v", err)
    }
    client := getClient(config, "gmail_token.json")

    srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
    if err != nil {
        log.Fatalf("Unable to retrieve Gmail client: %v", err)
    }

    return srv

    /*
    user := "me"
    r, err := srv.Users.Labels.List(user).Do()
    if err != nil {
        log.Fatalf("Unable to retrieve labels: %v", err)
    credentialscredentials}
    if len(r.Labels) == 0 {
        fmt.Println("No labels found.")
        return
    }
    fmt.Println("Labels:")
    for _, l := range r.Labels {
        fmt.Printf("- %s\n", l.Name)
    }
    */
}


func SendEmail(srv *gmail.Service, student SynStudent) {
    var rec string
    for _, pemail := range student.PEmails {
        if pemail == nil { continue }
        rec = *pemail
        break
    }
    myprofile, err := srv.Users.GetProfile("me").Do()
    if err != nil {
        log.Fatalf("Error while getting user data: %v", err)
    }
    myemail := myprofile.EmailAddress
    msgpath := os.Getenv("AUTO_MSG")
    policypath := os.Getenv("POLICY")
    msgfile, err := os.Open(msgpath)
    if err != nil {
        log.Fatalf("Error while opening message file: %v", err)
    }
    defer msgfile.Close()
    msgbody, err := io.ReadAll(msgfile)
    if err != nil {
        log.Fatalf("Error while reading message file: %v", err)
    }

    var draft gmail.Draft
    msg := gomail.NewMessage()
    msg.SetHeader("From", myemail)
    msg.SetHeader("To", rec)
    msg.SetHeader("Subject", student.Name + " failing assignment")
    msg.SetBody("text/plain", string(msgbody))
    msg.Attach(policypath)

    buffer := new(bytes.Buffer)
    if _, err := msg.WriteTo(buffer); err != nil {
        log.Fatalf("Error while writing message to email buffer: %v", err)
    }
    var m gmail.Message
    m.Raw = base64.URLEncoding.EncodeToString(buffer.Bytes())
    draft.Message = &m
    res, err := srv.Users.Drafts.Create("me", &draft).Do()
    if err != nil {
        log.Fatalf("Failed to create draft: %v", err)
    }

    if _, err = srv.Users.Drafts.Send("me", res).Do(); err != nil {
        log.Fatalf("Failed to send draft: %v", err) 
    }
}




























