package dto

type (
	UserTotpRequest struct {
		OtpCode     string `json:"otp_code" form:"otp_code" binding:"required"`
		AccountName string `json:"account_name" form:"account_name" binding:"required"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserRegisterRequest struct {
		UserLoginRequest
		FullName string `json:"full_name" form:"full_name" binding:"required"`
	}

	UserResponse struct {
		ID      string `json:"id"`
		Name    string `json:"name,omitempty"`
		Email   string `json:"email,omitempty"`
		Picture string `json:"picture,omitempty"`
	}
)
