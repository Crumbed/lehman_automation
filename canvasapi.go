package main


import (
    "fmt"
    "net/http"
    "bytes"
    "io/ioutil"
    "encoding/json"
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

func (api *CanvasApi) send(req *http.Request) ([]byte, error) {
    res, err := api.Client.Do(req)
    if err != nil {
        return nil, err
    }

    bytes, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }

    return bytes, nil
}

func (api *CanvasApi) Courses(courses *[]Course) error {
    url := fmt.Sprintf("%s/courses", BASE)
    req, err := api.get(url)
    if err != nil {
        return err
    }

    bytes, err  := api.send(req)
    if err != nil {
        return err
    }

    err = json.Unmarshal(bytes, courses)
    if err != nil {
        return err
    }
    return nil
}

func (api *CanvasApi) GradeChanges(gradeChanges *[]GradeEvent, studentId int, startTime *string) error {
    url := fmt.Sprintf("%s/audit/grade_chage/students/%d", BASE, studentId)
    if startTime != nil {
        url = fmt.Sprintf("%s?start_time=%s", url, *startTime)
    }
    
    req, err := api.get(url)
    if err != nil {
        return err
    }

    bytes, err := api.send(req)
    if err != nil {
        return err
    }

    err = json.Unmarshal(bytes, gradeChanges)
    if err != nil {
        return err
    }

    return nil
}

func (api *CanvasApi) Assignment(assignment *Assignment, courseId int, assignmentId int) error {
    url := fmt.Sprintf("%s/courses/%d/assignments/%d", BASE, courseId, assignmentId)
    req, err := api.get(url)
    if err != nil {
        return err
    }

    bytes, err := api.send(req)
    if err != nil {
        return err
    }

    err = json.Unmarshal(bytes, assignment)
    if err != nil {
        return err
    }
    
    return nil
}













