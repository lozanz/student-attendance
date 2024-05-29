package structs

type Attendance struct {
	ID        int64  `json:"id"`
	StudentID int    `json:"student_id"`
	Date      string `json:"date"`
	Status    string `json:"status"`
}
