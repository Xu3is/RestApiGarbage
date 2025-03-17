package restapigarbage

type Course struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	AgeGroup      string  `json:"age_group"`
	PricePerClass float64 `json:"price_per_class"`
	Duration      int     `json:"duration"`
}

type Trainer struct {
	ID             int    `json:"id"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	Specialization int    `json:"specialization"`
	Address        string `json:"address"`
}

type Lesson struct {
	ID        int    `json:"id"`
	CourseID  int    `json:"course_id"`
	TrainerID int    `json:"trainer_id"`
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type Student struct {
	ID               int    `json:"id"`
	FullName         string `json:"full_name"`
	PhoneNumber      string `json:"phone_number"`
	Age              int    `json:"age"`
	RegistrationDate string `json:"registration_date"`
	CourseID         int    `json:"course_id"`
}

type Payment struct {
	ID              int `json:"id"`
	StudentID       int `json:"student_id"`
	CourseID        int `json:"course_id"`
	AttendedLessons int `json:"attended_lessons"`
	PaidLessons     int `json:"paid_lessons"`
}
