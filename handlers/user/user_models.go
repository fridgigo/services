package user

type User struct {
	Email              string `json:"email"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Password           string `json:"password"`
	RepeatPassword     string `json:"repeat_password"`
	CreatedAt          int64  `json:"created_at"`
	UpdatedAt          int64  `json:"updated_at"`
	Deleted            bool   `json:"deleted"`
	Confirmed          bool   `json:"confirmed"`
	ConfirmationNumber int    `json:"confirmation_number"`
}
