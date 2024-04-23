package main





type Course struct {
    Id          int             `json:"id"`
    Enrollments []Enrollment    `json:"enrollments"`
}

type Enrollment struct {
    AccountId   string  `json:"sis_account_id"`
    UserId      string  `json:"sis_user_id"`
    Type        string  `json:"type"`
    User        User    `json:"user"`
}

type User struct {
    Id          int     `json:"id"`
    Name        string  `json:"name"`
    SortName    string  `json:"sortable_name"`
    SisId       string  `json:"sis_user_id"`
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

type Assignment struct {
    Id          int     `json:"id"`
    Name        string  `json:"name"`
}






















