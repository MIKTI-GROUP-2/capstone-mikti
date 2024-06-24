package handler

import (
	"capstone-mikti/features/users"
	"capstone-mikti/helper"
	"capstone-mikti/helper/jwt"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	service users.UserServiceInterface
	jwt     jwt.JWTInterface
}

func NewHandler(s users.UserServiceInterface, j jwt.JWTInterface) *UserHandler {
	return &UserHandler{
		service: s,
		jwt:     j,
	}
}

// Register implements users.UserHandlerInterface.
func (u *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RegisterInput)
		if err := c.Bind(&input); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input", nil))
		}

		isValid, errors := helper.ValidateJSON(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		if !helper.ValidatePassword(input.Password) {
			errPass := []string{"Password must contain a combination letters, symbols, and numbers"}
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errPass))
		}

		var serviceInput = new(users.User)
		serviceInput.Email = input.Email
		serviceInput.PhoneNumber = input.PhoneNumber
		serviceInput.Username = input.Username
		serviceInput.Password = input.Password

		result, err := u.service.Register(*serviceInput)

		if err != nil {
			if strings.Contains(err.Error(), "Username already registered by another user") {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Username Already Registered", nil))
			}
			c.Logger().Info("Handler : Register Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Register Process Failed", nil))
		}

		verificationCode := u.service.UserVerificationCode(input.Username, input.Email)

		if verificationCode != nil {
			logrus.Error("Handler : Send Email Error")
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Send Email Process Failed", nil))
		}

		var response = new(RegisterResponse)
		response.Email = result.Email
		response.Username = result.Username
		response.PhoneNumber = result.PhoneNumber

		return c.JSON(http.StatusCreated, helper.FormatResponse("Register Success, Please check your email to verification ", response))
	}
}

// Login implements users.UserHandlerInterface.
func (u *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginInput)

		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input", nil))
		}

		result, err := u.service.Login(input.Username, input.Password)

		if err != nil {
			c.Logger().Error("Handler : Login Failed : ", err.Error())
			if strings.Contains(err.Error(), "Not Found") {
				return c.JSON(http.StatusNotFound, helper.FormatResponse("User Not Found / User Inactive", nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Login Process Error", nil))
		}

		var response = new(LoginResponse)
		response.Username = result.Username
		response.Token = result.Access

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Login", response))
	}
}

// ForgetPasswordWeb implements users.UserHandlerInterface.
func (u *UserHandler) ForgetPasswordWeb() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(ForgetPasswordInput)

		if err := c.Bind(&input); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input", nil))
		}

		isValid, err := helper.ValidateJSON(input)
		if !isValid {
			c.Logger().Info("Handler : Bind Input Error : ", err)
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Validation Error", err))
		}

		result := u.service.ForgetPasswordWeb(input.Username)

		if result != nil {
			logrus.Error("Handler : Send Email Error")
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Send Email Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Send Reset Code to Email", nil))
	}
}

// ResetPassword implements users.UserHandlerInterface.
func (u *UserHandler) ResetPassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		var token = c.QueryParam("token_reset_password")
		if token == "" {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Token Not Found", nil))
		}

		dataToken, err := u.service.TokenResetVerify(token)
		if err != nil {
			c.Logger().Info("Handler : Token Reset Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Token Reset Verifi Error", nil))
		}

		_, err = u.service.TokenResetVerify(token)
		if err != nil {
			c.Logger().Info("Handler : Token Verify Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Token Expired", nil))
		}

		var input = new(ResetPasswordInput)
		if err := c.Bind(&input); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input", nil))
		}

		isValid, errorMsg := helper.ValidateJSON(input)
		if !isValid {
			c.Logger().Info("Handler : Validate Json Error")
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Validation Error", errorMsg))
		}

		if input.Password != input.PasswordConfirm {
			c.Logger().Info("Handler : Password Not Match")
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Password Not Match", nil))
		}

		if !helper.ValidatePassword(input.Password) {
			errPass := []string{"Password must contain a combination letters, symbols, and numbers"}
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errPass))
		}

		result := u.service.ResetPassword(dataToken.Code, dataToken.Username, input.Password)

		if result != nil {
			c.Logger().Info("Handler : Reset Password Error : ", result.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Reset Password Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success to Reset Password", result))
	}
}

// UpdateProfile implements users.UserHandlerInterface.
func (u *UserHandler) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		ext, err := u.jwt.ExtractToken(c)

		if err != nil {
			c.Logger().Error("Handler : Extract Token Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Extract Token Error", nil))
		}

		id := ext.ID
		var input = new(UpdateProfile)
		if err := c.Bind(&input); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input", nil))
		}

		var serviceUpdate = new(users.UpdateProfile)
		serviceUpdate.Username = input.Username
		serviceUpdate.PhoneNumber = input.PhoneNumber
		serviceUpdate.Email = input.Email

		result, err := u.service.UpdateProfile(int(id), *serviceUpdate)
		if err != nil {
			c.Logger().Info("Handler : Update Profile Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Update Profile Error", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Update Profile", result))
	}
}

// RefreshToken implements users.UserHandlerInterface.
func (u *UserHandler) RefreshToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RefreshTokenInput)
		if err := c.Bind(&input); err != nil {
			c.Logger().Info("Helper : Bind Refresh Token Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input", nil))
		}

		var currentToken = u.jwt.GetCurrentToken(c)

		result, err := u.jwt.RefreshJWT(input.Token, currentToken)

		if err != nil {
			c.Logger().Info("Handler : Refresh JWT Process Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Refresh JWT Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Refresh Token", result))
	}
}

func (u *UserHandler) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		ext, err := u.jwt.ExtractToken(c)

		if err != nil {
			c.Logger().Error("Handler : Extract Token Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Extract Token Error", nil))
		}

		id := ext.ID

		res, err := u.service.Profile(int(id))
		if err != nil {
			c.Logger().Error("Handler : Get Profile Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get Profile Error", nil))
		}

		var response = new(UserInfo)
		response.Email = res.Email
		response.Username = res.Username
		response.PhoneNumber = res.PhoneNumber
		response.Role = ext.Role
		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Profile", response))
	}
}

func (u *UserHandler) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := u.jwt.ValidateRole(c); !err {
			c.Logger().Info("Handler : Unauthorized Access : ", errors.New("you have no permission to access this feature"))
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Restricted Access", nil))
		}

		result, err := u.service.GetAll()

		if err != nil {
			c.Logger().Error("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get All Process Error", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get All Data", result))
	}
}

// ActivateUser handles deleting an existing user.
func (u *UserHandler) ActivateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := u.jwt.ValidateRole(c); !err {
			c.Logger().Info("Handler : Unauthorized Access : ", errors.New("you have no permission to access this feature"))
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Restricted Access", nil))
		}

		id := c.Param("id")
		userID, err := strconv.Atoi(id)
		if err != nil {
			c.Logger().Info("Handler : Invalid ID : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User ID", nil))
		}

		res, err := u.service.Activate(userID)
		if err != nil {
			c.Logger().Info("Handler : Activate User Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Activate User Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Activate User", res))
	}
}

// DeactivateUser handles deleting an existing user.
func (u *UserHandler) DeactivateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := u.jwt.ValidateRole(c); !err {
			c.Logger().Info("Handler : Unauthorized Access : ", errors.New("you have no permission to access this feature"))
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Restricted Access", nil))
		}

		id := c.Param("id")
		userID, err := strconv.Atoi(id)
		if err != nil {
			c.Logger().Info("Handler : Invalid ID : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User ID", nil))
		}

		res, err := u.service.Deactivate(userID)
		if err != nil {
			c.Logger().Info("Handler : Deactivate User Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Deactivate User Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Deactivate User", res))
	}
}

func (u *UserHandler) UserDashboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := u.jwt.ValidateRole(c); !err {
			c.Logger().Info("Handler : Unauthorized Access : ", errors.New("you have no permission to access this feature"))
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Restricted Access", nil))
		}

		res, err := u.service.UserDashboard()

		if err != nil {
			c.Logger().Error("Handler: Callback process error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Dashboard User Error", nil))
		}

		var response = new(DashboardResponse)
		response.TotalUser = res.TotalUser
		response.TotalUserBaru = res.TotalUserBaru
		response.TotalUserActive = res.TotalUserActive
		response.TotalUserInactive = res.TotalUserInactive

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Patient", response))
	}
}

// ResetPassword implements users.UserHandlerInterface.
func (u *UserHandler) UserVerification() echo.HandlerFunc {
	return func(c echo.Context) error {
		var token = c.QueryParam("token_verification")
		if token == "" {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Token Not Found", nil))
		}

		dataToken, err := u.service.TokenVerificationResetVerify(token)
		if err != nil {
			c.Logger().Info("Handler : Token Reset Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Token Reset Verifi Error", nil))
		}

		_, err = u.service.TokenVerificationResetVerify(token)
		if err != nil {
			c.Logger().Info("Handler : Token Verify Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Token Expired", nil))
		}

		result := u.service.UserVerification(dataToken.Code, dataToken.Username)

		if result != nil {
			c.Logger().Info("Handler : Verification User Error : ", result.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Verification User Error", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success to verification, enable to login", result))
	}
}
