package service

import (
	"capstone-mikti/features/users"
	"capstone-mikti/helper/email"
	"capstone-mikti/helper/enkrip"
	"capstone-mikti/helper/jwt"
	"errors"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type UserService struct {
	data   users.UserDataInterface
	jwt    jwt.JWTInterface
	enkrip enkrip.HashInterface
	email  email.EmailInterface
}

func New(d users.UserDataInterface, j jwt.JWTInterface, e enkrip.HashInterface, em email.EmailInterface) *UserService {
	return &UserService{
		data:   d,
		jwt:    j,
		enkrip: e,
		email:  em,
	}
}

// Register implements users.UserServiceInterface.
func (u *UserService) Register(newData users.User) (*users.User, error) {
	_, err := u.data.GetByUsername(newData.Username)

	if err == nil {
		logrus.Error("Service : Username already registered")
		return nil, errors.New("ERROR Username already registered by another user")
	}

	hashPassword, err := u.enkrip.HashPassword(newData.Password)
	if err != nil {
		logrus.Error("Service : Error Hash Password : ", err.Error())
		return nil, errors.New("ERROR Hash Password")
	}

	newData.Password = hashPassword
	newData.IsAdmin = false
	newData.Status = true

	result, err := u.data.Register(newData)
	if err != nil {
		logrus.Error("Service : Error Register : ", err.Error())
		return nil, errors.New("ERROR Register")
	}

	return result, nil
}

// Login implements users.UserServiceInterface.
func (u *UserService) Login(username string, password string) (*users.UserCredential, error) {
	result, err := u.data.Login(username, password)

	if err != nil {
		if strings.Contains(err.Error(), "Incorrect Password") {
			return nil, errors.New("ERROR Incorrect Password")
		}
		if strings.Contains(err.Error(), "Not Found") {
			return nil, errors.New("ERROR User Not Found / User Inactive")
		}
		return nil, errors.New("ERROR Process Failed")
	}

	role := "user"

	if result.IsAdmin {
		role = "Admin"
	}

	tokenData := u.jwt.GenerateJWT(result.ID, result.Username, result.Email, result.PhoneNumber, role)

	if tokenData == nil {
		logrus.Error("Service : Generate Token Error : ")
		return nil, errors.New("ERROR Token Process Failed")
	}

	response := new(users.UserCredential)
	response.Username = result.Username
	response.Access = tokenData

	return response, nil
}

// TokenResetVerify implements users.UserServiceInterface.
func (u *UserService) TokenResetVerify(code string) (*users.UserResetPass, error) {
	result, err := u.data.GetByCode(code)
	if err != nil {
		logrus.Error("Service : Error Get By Code : ", err.Error())
		return nil, errors.New("ERROR Failed to verify token")
	}

	if result.ExpiredAt.Before(time.Now()) {
		return nil, errors.New("ERROR Token Expired")
	}

	return result, nil
}

// ForgetPasswordWeb implements users.UserServiceInterface.
func (u *UserService) ForgetPasswordWeb(username string) error {
	user, err := u.data.GetByUsername(username)
	if err != nil {
		logrus.Error("Service : Error Get By Username : ", err.Error())
		return errors.New("ERROR Data Not Found")
	}

	email := user.Email

	username = user.Username
	header, htmlBody, code := u.email.HTMLBody(username)

	if err := u.data.InsertCode(username, code); err != nil {
		logrus.Error("Service : Error Insert Code : ", err.Error())
		return errors.New("ERROR Insert Code Failed")
	}

	err = u.email.SendEmail(email, header, htmlBody)

	if err != nil {
		logrus.Error("Service : Error Send Email : ", err.Error())
		return errors.New("ERROR Send Email")
	}

	return nil
}

// ResetPassword implements users.UserServiceInterface.
func (u *UserService) ResetPassword(code string, username string, password string) error {
	hashPassword, err := u.enkrip.HashPassword(password)
	if err != nil {
		logrus.Error("Service : Hash Password Error : ", err.Error())
		return errors.New("ERROR Hash Password")
	}

	password = hashPassword

	if err := u.data.ResetPassword(code, username, password); err != nil {
		logrus.Error("Service : Reset Password Error : ", err.Error())
		return errors.New("ERROR Reset Password Failed")
	}

	return nil
}

// UpdateProfile implements users.UserServiceInterface.
func (u *UserService) UpdateProfile(id int, newData users.UpdateProfile) (bool, error) {
	res, err := u.data.UpdateProfile(id, newData)
	if err != nil {
		logrus.Error("Service : Update Profile Error : ", err.Error())
		return false, errors.New("ERROR Update Failed")
	}

	return res, nil
}

func (u *UserService) Profile(id int) (*users.User, error) {
	res, err := u.data.GetByID(id)

	if err != nil {
		logrus.Error("Service : Get Profile Error : ", err.Error())
		return nil, errors.New("ERROR Get Profile")
	}

	return &res, nil
}

func (u *UserService) GetAll() ([]users.User, error) {
	result, err := u.data.GetAll()

	if err != nil {
		logrus.Error("Service Error: ", err)
		return nil, errors.New("ERROR Get All Process Failed")
	}

	return result, nil
}

func (u *UserService) Activate(id int) (bool, error) {
	res, err := u.data.Activate(id)
	if err != nil {
		logrus.Error("Service Error : ", err.Error())
		return res, errors.New("ERROR Activate User")
	}
	return res, nil
}

func (u *UserService) Deactivate(id int) (bool, error) {
	res, err := u.data.Deactivate(id)
	if err != nil {
		logrus.Error("Service Error : ", err.Error())
		return res, errors.New("ERROR Deactivate User")
	}
	return res, nil
}

func (u *UserService) UserDashboard() (users.UserDashboard, error) {
	res, err := u.data.UserDashboard()

	if err != nil {
		logrus.Error("Service Error : ", err.Error())
		return res, errors.New("ERROR Process Failed")
	}

	return res, nil
}
