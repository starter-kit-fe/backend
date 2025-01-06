package dto

type GoogleSigninRequest struct {
	AccessToken string `uri:"access_token" binding:"required,min=10"`
}
type EmailExistsQueryRequest struct {
	Email string `uri:"email" binding:"required,email"`
}
type EmailCodeRequest struct {
	Token string `json:"token" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type SignupRequest struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Code     string `json:"code" binding:"required,min=4"`
	Email    string `json:"email" binding:"required,email"`
}

type SigninRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Token    string `json:"token" binding:"required"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	NickName  string `json:"nickName"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
	Gender    uint   `json:"gender"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
