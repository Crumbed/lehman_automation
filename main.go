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

    res := canvas.Courses()
    if res == nil {
        return
    }

    byteBody, err := ioutil.ReadAll(*res)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(byteBody))
}




func getFailingStudents(api *CanvasApi) {
    
}


















