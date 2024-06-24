package data

import (
	"capstone-mikti/features/users"
	"capstone-mikti/helper/enkrip"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserData struct {
	db     *gorm.DB
	enkrip enkrip.HashInterface
}

func New(db *gorm.DB, enkrip enkrip.HashInterface) *UserData {
	return &UserData{
		db:     db,
		enkrip: enkrip,
	}
}

func (ud *UserData) Register(newData users.User) (*users.User, error) {
	var dbData = new(User)
	dbData.Username = newData.Username
	dbData.Email = newData.Email
	dbData.PhoneNumber = newData.PhoneNumber
	dbData.Password = newData.Password
	dbData.IsAdmin = newData.IsAdmin
	dbData.Status = newData.Status

	if err := ud.db.Create(dbData).Error; err != nil {
		logrus.Error("DATA : Register Error : ", err.Error())
		return nil, err
	}

	return &newData, nil
}

func (ud *UserData) Login(username, password string) (*users.User, error) {
	var dbData = new(User)
	dbData.Username = username

	var qry = ud.db.Where("username = ? AND status = ?", dbData.Username, true).First(dbData)

	var dataCount int64
	qry.Count(&dataCount)

	if dataCount == 0 {
		logrus.Error("DATA : Login Error : Data Not Found")
		return nil, errors.New("ERROR Data Not Found")
	}

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Login Get Error : ", err.Error())
		return nil, err
	}

	if err := ud.enkrip.Compare(dbData.Password, password); err != nil {
		logrus.Error("DATA : Incorrect Password")
		return nil, errors.New("ERROR Incorrect Password")
	}

	var result = new(users.User)
	result.ID = dbData.ID
	result.Username = dbData.Username
	result.Email = dbData.Email
	result.PhoneNumber = dbData.PhoneNumber
	result.IsAdmin = dbData.IsAdmin
	result.Status = dbData.Status

	return result, nil
}

func (ud *UserData) GetByID(id int) (users.User, error) {
	var listUser users.User
	var qry = ud.db.Table("users").Select("users.*").
		Where("users.id = ?", id).
		Where("users.status = ?", true).
		Scan(&listUser)

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error Get By ID : ", err.Error())
		return listUser, err
	}

	return listUser, nil
}

func (ud *UserData) GetByUsername(username string) (*users.User, error) {
	var dbData = new(User)
	dbData.Username = username

	var qry = ud.db.Where("username = ?", dbData.Username).First(dbData)

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error Get By Username : ", err.Error())
		return nil, err
	}

	var result = new(users.User)
	result.ID = dbData.ID
	result.Username = dbData.Username
	result.Email = dbData.Email
	result.PhoneNumber = dbData.PhoneNumber
	result.IsAdmin = dbData.IsAdmin
	result.Status = dbData.Status

	return result, nil
}

func (ud *UserData) InsertCode(username, code string) error {
	var newData = new(UserResetPass)
	newData.Username = username
	newData.Code = code
	newData.ExpiredAt = time.Now().Add(time.Minute * 10)

	_, err := ud.GetByCode(code)
	if err != nil {
		ud.DeleteCode(code)
	}

	if err := ud.db.Table("user_reset_passes").Create(newData).Error; err != nil {
		logrus.Error("DATA : Error Create User Reset Pass : ", err.Error())
		return err
	}

	return nil
}

func (ud *UserData) DeleteCode(code string) error {
	var deleteData = new(UserResetPass)

	if err := ud.db.Table("user_reset_passes").Where("code = ? ", code).Delete(deleteData).Error; err != nil {
		logrus.Error("DATA : Error Delete User Reset Pass : ", err.Error())
		return err
	}

	return nil
}

func (ud *UserData) GetByCode(code string) (*users.UserResetPass, error) {
	var dbData = new(UserResetPass)
	dbData.Code = code

	if err := ud.db.Table("user_reset_passes").Where("code = ?", dbData.Code).First(dbData).Error; err != nil {
		logrus.Error("DATA : Error Get User Reset Pass : ", err.Error())
	}

	var result = new(users.UserResetPass)
	result.Username = dbData.Username
	result.Code = dbData.Code
	result.ExpiredAt = dbData.ExpiredAt

	return result, nil
}

func (ud *UserData) ResetPassword(code, username, password string) error {
	if err := ud.db.Table("users").Where("username = ?", username).Update("password", password).Error; err != nil {
		logrus.Error("DATA : Error Update Password : ", err.Error())
		return err
	}

	checkData, _ := ud.GetByCode(code)
	if checkData.Code != "" {
		ud.DeleteCode(code)
	}

	return nil
}

func (ud *UserData) UpdateProfile(id int, newData users.UpdateProfile) (bool, error) {
	var qry = ud.db.Table("users").Where("id = ?", id).Updates(User{
		Username:    newData.Username,
		Email:       newData.Email,
		PhoneNumber: newData.PhoneNumber,
	})

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error Update Profile : ", err.Error())
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		logrus.Error("DATA : No Row Affected")
		return false, nil
	}

	return true, nil
}

func (ud *UserData) GetAll() ([]users.User, error) {
	var listUser = []users.User{}

	if err := ud.db.Find(&listUser).Error; err != nil {
		return nil, err
	}

	return listUser, nil
}

func (u *UserData) Activate(id int) (bool, error) {
	var qry = u.db.Table("users").Where("id = ?", id).Updates(User{Status: true})

	if err := qry.Error; err != nil {
		return false, err
	}

	return true, nil
}

func (u *UserData) Deactivate(id int) (bool, error) {
	var qry = u.db.Model(&User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"Status": false,
	})

	if err := qry.Error; err != nil {
		return false, err
	}

	return true, nil
}

func (u *UserData) UserDashboard() (users.UserDashboard, error) {
	var dashboardUser users.UserDashboard

	tUser, tUserBaru, tUserActive, tUserInactive := u.getTotalUser()

	dashboardUser.TotalUser = tUser
	dashboardUser.TotalUserBaru = tUserBaru
	dashboardUser.TotalUserActive = tUserActive
	dashboardUser.TotalUserInactive = tUserInactive

	return dashboardUser, nil
}

func (pdata *UserData) getTotalUser() (int, int, int, int) {
	var totalUser int64
	var totalUserBaru int64
	var totalUserActive int64
	var totalUserInactive int64

	var now = time.Now()
	var before = now.AddDate(0, 0, -7)

	var _ = pdata.db.Table("users").Count(&totalUser)
	var _ = pdata.db.Table("users").Where("created_at BETWEEN ? and ?", before, now).Count(&totalUserBaru)
	var _ = pdata.db.Table("users").Where("status = ?", true).Count(&totalUserActive)
	var _ = pdata.db.Table("users").Where("status = ?", false).Count(&totalUserInactive)

	totalUserInt := int(totalUser)
	totalUserBaruInt := int(totalUserBaru)
	totalUserActiveInt := int(totalUserActive)
	totalUserInactiveInt := int(totalUserInactive)

	return totalUserInt, totalUserBaruInt, totalUserActiveInt, totalUserInactiveInt
}

func (ud *UserData) InsertCodeVerification(username, code string) error {
	var newData = new(UserVerification)
	newData.Username = username
	newData.Code = code
	newData.ExpiredAt = time.Now().Add(time.Minute * 10)

	_, err := ud.GetByCodeVerification(code)
	if err != nil {
		ud.DeleteCodeVerfication(code)
	}

	if err := ud.db.Table("user_verifications").Create(newData).Error; err != nil {
		logrus.Error("DATA : Error Create User Verification : ", err.Error())
		return err
	}

	return nil
}

func (ud *UserData) DeleteCodeVerfication(code string) error {
	var deleteData = new(UserVerification)

	if err := ud.db.Table("user_verification").Where("code = ? ", code).Delete(deleteData).Error; err != nil {
		logrus.Error("DATA : Error Delete User Verification : ", err.Error())
		return err
	}

	return nil
}

func (ud *UserData) GetByCodeVerification(code string) (*users.UserVerification, error) {
	var dbData = new(UserVerification)
	dbData.Code = code

	if err := ud.db.Table("user_verifications").Where("code = ?", dbData.Code).First(dbData).Error; err != nil {
		logrus.Error("DATA : Error Get User Verifications : ", err.Error())
	}

	var result = new(users.UserVerification)
	result.Username = dbData.Username
	result.Code = dbData.Code
	result.ExpiredAt = dbData.ExpiredAt

	return result, nil
}

func (ud *UserData) UserVerification(code, username string) error {
	if err := ud.db.Table("users").Where("username = ?", username).Update("status", true).Error; err != nil {
		logrus.Error("DATA : Error Verification : ", err.Error())
		return err
	}

	checkData, _ := ud.GetByCodeVerification(code)
	if checkData.Code != "" {
		ud.DeleteCodeVerfication(code)
	}

	return nil
}
