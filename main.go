package main;



import (
    "fmt"
    "net/http"
    "bytes"
    "io/ioutil"
    "io"
    "time"
    "os"
)

var token string
var canvas_api string = "https://canvas.instructure.com/api/v1/"

func main() {
    fmt.Println("hello world")
    token = os.Args[1]
    fmt.Println(token)

    client := http.Client {
        Timeout: 30 * time.Second,
    }
    
    r := auth(client)
    if r == nil {
        return
    }

    bBody, _ := ioutil.ReadAll(*r)
    body := string(bBody)
    

    fmt.Println(body)
}

func auth(client http.Client) *io.ReadCloser {
    url := fmt.Sprintf("%scourses", canvas_api)
    reader := bytes.NewReader([]byte(""))
    req, err := http.NewRequest(http.MethodGet, url, reader)
    if err != nil {
        fmt.Println(err)
        return nil
    }

    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
    res, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return nil
    }

    return &res.Body
}
