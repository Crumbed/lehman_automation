package main;




import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "time"
)


func main() {
    fmt.Println("hello world")
    token := os.Args[1]
    client := http.Client {
        Timeout: 30 * time.Second,
    }
    canvas := NewCanvasApi(token, client)

    for true {
        students, err := getStudents(canvas)
        if err != nil {
            fmt.Println(err)
            continue
        }

        for _, student := range students {
            var gradeChanges []GradeEvent
            err := canvas.GradeChanges(&gradeChanges, student.Id, nil)
            if err != nil {
                fmt.Println(err)
                continue
            }

        }
    }
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


















