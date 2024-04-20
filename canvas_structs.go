package main





type Course struct {
    Id      int   `json:"id"`
}

type GradeLinks struct {
    Student     int     `json:"student"`
    Grader      int     `json:"grader"`
    Course      int     `json:"course"`
    Assignment  int     `json:"assignment"`
}

type GradeEvent struct {
    Id          int         `json:"id"`
    GradeBefore string      `json:"grade_before"`
    GradeAfter  string      `json:"grade_after"`
    RequestId   string      `json:"request_id"`
    Links       GradeLinks  `json:"links"`
}






















