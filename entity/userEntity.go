package entity

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Pin      string `json:"pin"`
	Cookies  string `json:"cookies"`
	OtpLink  string `json:"otpLink"`
}

type UserResponse struct {
	Message string `json:"message"`
	IsError bool   `json:"isError"`
	Data    any    `json:"data"`
}

type UserVerify struct {
	Otp string `json:"otp"`
}

type Denom struct {
	Label string `json:"label"`
	Pid   string `json:"pid"`
	Coin  string `json:"coin"`
	Rgid  string `json:"rgid"`
}

type Redeem struct {
	ID     string `json:"id"`
	Code   string `json:"code"`
	Status string `json:"status"`
}
