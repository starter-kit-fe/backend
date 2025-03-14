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
type UserRoutes = []RouterItem

type RouterItem struct {
	Name       string       `json:"name"`
	Path       string       `json:"path"`
	Hidden     bool         `json:"hidden"`
	Redirect   string       `json:"redirect,omitempty"`
	Component  string       `json:"component"`
	AlwaysShow bool         `json:"alwaysShow,omitempty"`
	Meta       Meta         `json:"meta"`
	Children   []RouterItem `json:"children,omitempty"`
}

// Meta contains the metadata for menu items
type Meta struct {
	Title   string      `json:"title"`
	Icon    string      `json:"icon"`
	NoCache bool        `json:"noCache"`
	Link    interface{} `json:"link"`
}
