package users

type UserModelValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
	} `json:"user"`
	// userModel UserModel `json:"-"`
}


func NewUserModelValidator() UserModelValidator{
	userModelValidator := UserModelValidator{}
	return userModelValidator
}