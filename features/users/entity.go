package users

import (
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	IsAdmin     bool   `json:"is_admin"`
	Status      bool   `json:"status"`
}

type UserCredential struct {
	Username string         `json:"username"`
	Access   map[string]any `json:"token"`
}

type UserInfo struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UserResetPass struct {
	Username  string    `json:"username"`
	Code      string    `json:"code"`
	ExpiredAt time.Time `json:"expired_at"`
}

type UpdateProfile struct {
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type UserDashboard struct {
	TotalUser         int `json:"total_user"`
	TotalUserBaru     int `json:"total_new_user"`
	TotalUserActive   int `json:"total_active_user"`
	TotalUserInactive int `json:"total_inactive_user"`
}

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	ForgetPasswordWeb() echo.HandlerFunc
	ResetPassword() echo.HandlerFunc
	UpdateProfile() echo.HandlerFunc
	RefreshToken() echo.HandlerFunc
	Profile() echo.HandlerFunc

	GetUsers() echo.HandlerFunc
	ActivateUser() echo.HandlerFunc
	DeactivateUser() echo.HandlerFunc

	UserDashboard() echo.HandlerFunc
}

type UserServiceInterface interface {
	Register(newData User) (*User, error)
	Login(username string, password string) (*UserCredential, error)
	ForgetPasswordWeb(username string) error
	TokenResetVerify(code string) (*UserResetPass, error)
	ResetPassword(code, username, password string) error
	UpdateProfile(id int, newData UpdateProfile) (bool, error)
	Profile(id int) (*User, error)

	GetAll() ([]User, error)
	Activate(id int) (bool, error)
	Deactivate(id int) (bool, error)

	UserDashboard() (UserDashboard, error)
}

type UserDataInterface interface {
	Register(newData User) (*User, error)
	Login(username, password string) (*User, error)
	GetByID(id int) (User, error)
	GetByUsername(username string) (*User, error)
	InsertCode(username, code string) error
	DeleteCode(code string) error
	GetByCode(code string) (*UserResetPass, error)
	ResetPassword(code, username, password string) error
	UpdateProfile(id int, newData UpdateProfile) (bool, error)

	GetAll() ([]User, error)
	Activate(id int) (bool, error)
	Deactivate(id int) (bool, error)

	UserDashboard() (UserDashboard, error)
}
