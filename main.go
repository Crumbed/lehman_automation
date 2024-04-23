package main;




import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "time"
    "strconv"
)


type Record struct {
    Student     string
    Assignment  string
}


func main() {
    fmt.Println("hello world")
    token := os.Args[1]
    client := http.Client {
        Timeout: 30 * time.Second,
    }
    canvas := NewCanvasApi(token, client)

    students, err := getStudents(canvas)
    if err != nil {
        fmt.Println(err)
        continue
    }

    records := getRecords(canvas, &students)
}


func getRecords(api *CanvasApi, students *map[string]User) []Record {
    records := make([]Record, 0)
    for _, s := range students {
        var gradeChanges []GradeEvent
        err := api.GradeChanges(&gradeChanges, s.Id, nil)
        if err != nil {
            fmt.Println(err)
            continue
        }

        for _, gradeChange := range gradeChanges {
            grade := strconv.Atoi(gradeChange.GradeAfter)
            if grade >= 60 { continue }
            var ass Assignment
            canvas.Assignment(
                &ass,
                gradeChange.Links.Course,
                gradeChange.Links.Assignment)

            records = append(records, Record{
                Student: s.SisId,
                Assignment: ass.Name,
            })
        }
    }

    return records
}


func getStudents(api *CanvasApi) (map[string]User, error) {
    var courses []Course
    err := api.Courses(&courses)
    if err != nil {
        return nil, err
    }

    students := make(map[string]User)
    for _, course := range courses {
        for _, enrollment := range course.Enrollments {
            if enrollment.Type != "StudentEnrollment" {
                continue
            }

            students[enrollment.User.SisId] = enrollment.User
        }
    }

    return students, nil
}


















