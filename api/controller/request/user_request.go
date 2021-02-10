package request

type SignUp struct {
	Name        string `json:"name" validate:"required"`
	LoginId     string `json:"login_id"  validate:"required"`
	Password    string `json:"password" validate:"required"`
	MailAddress string `json:"mail_address" validate:"required,email"`
}

type SignIn struct {
	LoginId  string `json:"login_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}