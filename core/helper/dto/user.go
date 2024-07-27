package dto

type (
	UserLoginRequest struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserRegisterRequest struct {
		UserLoginRequest
		FullName    string `json:"full_name" form:"full_name" binding:"required"`
	}

	UserResponse struct {
		ID      string `json:"id"`
		Name    string `json:"name,omitempty"`
		Email   string `json:"email,omitempty"`
		Picture string `json:"picture,omitempty"`
	}
)
