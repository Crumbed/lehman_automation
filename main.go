package main;




import (
    "log"
    "fmt"
    "net/http"
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
    sheetssrv := googleSheets()
    gmailsrv := getGmail()

    for {
        students, err := getStudents(canvas)
        if err != nil {
            log.Fatalf("Unable to retrieve students: %v", err)
        }

        records := getRecords(canvas, &students)
        _ = records

        infoId := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
        synStudents := getStudentInfo(sheetssrv, infoId)


        for _, record := range records {
            synstudent, ok := (*synStudents)[record.Student]
            if !ok {
                fmt.Printf("Student %s not found\n", record.Student)
                continue
            }

            SendEmail(gmailsrv, synstudent)
        }
    }
}


func getRecords(api *CanvasApi, students *map[string]User) []Record {
    records := make([]Record, 0, 10)
    for _, s := range *students {
        var gradeChanges []GradeEvent
        err := api.GradeChanges(&gradeChanges, s.Id, nil)
        if err != nil {
            fmt.Println(err)
            continue
        }

        for _, gradeChange := range gradeChanges {
            grade, err := strconv.Atoi(gradeChange.GradeAfter)
            if err != nil {
                log.Fatalf("Unable to convert grade to int: %v", err)
            }
            if grade >= 60 { continue }
            var ass Assignment
            api.Assignment(
                &ass,
                gradeChange.Links.Course,
                gradeChange.Links.Assignment)

            records = append(records, Record{
                Student: s.Name,
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

            students[enrollment.User.Name] = enrollment.User
        }
    }

    return students, nil
}


















