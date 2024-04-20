package main


import (
    "fmt"
    "net/http"
    "bytes"
    "io"
)


const BASE string = "https://canvas.instructure.com/api/v1/"

type CanvasApi struct {
    auth    string
    Client  http.Client
}


func NewCanvasApi(auth string, client http.Client) *CanvasApi {
    return &CanvasApi {
        auth, 
        client, 
    }
}

func (api *CanvasApi) get(url string) (*http.Request, error) {
    reader := bytes.NewReader([]byte(""))
    req, err := http.NewRequest(http.MethodGet, url, reader)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.auth))
    return req, nil
}

func (api *CanvasApi) Courses() *io.ReadCloser {
    url := newUrlBuilder().
        Courses().
        String()
    req, err := api.get(url)
    if err != nil {
        fmt.Println(err)
        return nil
    }

    res, err := api.Client.Do(req)
    if err != nil {
        fmt.Println(err)
        return nil
    }

    return &res.Body
}


type urlBuilder struct {
    url []byte
}

func newUrlBuilder() *urlBuilder {
    return &urlBuilder {
        url: []byte(BASE),
    }
}

func (b *urlBuilder) Courses() *urlBuilder {
    b.url = append(b.url, []byte("courses/")...)
    return b
}

func (b *urlBuilder) Course(id string) *urlBuilder {
    b.url = append(b.url, []byte(fmt.Sprintf("%s/", id))...)
    return b
}




func (b *urlBuilder) String() string {
    return string(b.url)
}













