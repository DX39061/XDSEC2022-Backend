package repository

import (
	"XDSEC2022-Backend/src/logger"
	"XDSEC2022-Backend/src/model"
	"XDSEC2022-Backend/src/utility"
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var UserInitFlag = false

func init() {
	RegisterModel(model.User{})
}

func GetUserCount() (model.UserCount, error) {
	var count model.UserCount
	err := Database.Model(&model.User{}).Count(&count.Total).Error
	if err != nil {
		return model.UserCount{}, err
	}
	err = Database.Model(&model.User{}).Where("direction = ?", "web").Count(&count.Web).Error
	if err != nil {
		return model.UserCount{}, err
	}
	err = Database.Model(&model.User{}).Where("direction = ?", "reverse").Count(&count.Reverse).Error
	if err != nil {
		return model.UserCount{}, err
	}
	err = Database.Model(&model.User{}).Where("direction = ?", "pwn").Count(&count.Pwn).Error
	if err != nil {
		return model.UserCount{}, err
	}
	err = Database.Model(&model.User{}).Where("direction = ?", "crypto").Count(&count.Crypto).Error
	if err != nil {
		return model.UserCount{}, err
	}
	err = Database.Model(&model.User{}).Where("direction = ?", "misc").Count(&count.Misc).Error
	if err != nil {
		return model.UserCount{}, err
	}
	err = Database.Model(&model.User{}).Where("direction = ?", "dev").Count(&count.Dev).Error
	if err != nil {
		return model.UserCount{}, err
	}
	err = Database.Model(&model.User{}).Where("department = ?", "secretariat").Count(&count.Secretariat).Error
	if err != nil {
		return model.UserCount{}, err
	}
	err = Database.Model(&model.User{}).Where("department = ?", "devops").Count(&count.Devops).Error
	if err != nil {
		return model.UserCount{}, err
	}
	err = Database.Model(&model.User{}).Where("department = ?", "technique").Count(&count.Technique).Error
	if err != nil {
		return model.UserCount{}, err
	}
	err = Database.Model(&model.User{}).Where("department = ?", "publicity").Count(&count.Publicity).Error
	if err != nil {
		return model.UserCount{}, err
	}
	err = Database.Model(&model.User{}).Where("sex = ?", "男").Count(&count.Male).Error
	if err != nil {
		return model.UserCount{}, err
	}
	err = Database.Model(&model.User{}).Where("sex = ?", "女").Count(&count.Female).Error
	if err != nil {
		return model.UserCount{}, err
	}
	err = Database.Model(&model.User{}).Where("admin = ?", true).Count(&count.Admin).Error
	if err != nil {
		return model.UserCount{}, err
	}
	return count, nil
}

func SearchUsers(keyword string) ([]model.UserShort, error) {
	var users []model.User
	payload, args := utility.ConstructPayload(keyword, &model.User{})
	err := Database.Model(&model.User{}).
		Where(payload, args...).
		Order("id asc").
		Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	var usersShort []model.UserShort
	err = copier.Copy(&usersShort, users)
	if err != nil {
		return nil, err
	}
	return usersShort, nil
}

func GetUserList(limit int, skip int) ([]model.UserShort, error) {
	var users []model.User
	err := Database.Model(&model.User{}).
		Offset(skip).
		Limit(limit).
		Order("id asc").
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	var UsersShort []model.UserShort
	err = copier.Copy(&UsersShort, users)
	if err != nil {
		return nil, err
	}
	return UsersShort, nil
}

func GetUserDetail(id uint) (model.UserDetail, error) {
	var user model.User
	err := Database.Model(&model.User{}).
		Where("id = ?", id).
		First(&user).Error
	if err != nil {
		return model.UserDetail{}, err
	}
	var userDetail model.UserDetail
	err = copier.Copy(&userDetail, user)
	if err != nil {
		return model.UserDetail{}, err
	}
	return userDetail, nil
}

func GetUserByAccountString(account string) (model.User, error) {
	var user model.User
	if err := Database.Model(&model.User{}).
		Where("student_id = ? or telephone = ? or email = ?", account, account, account).
		First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func ValidateUserPassword(account string, password string) (model.User, error) {
	user, err := GetUserByAccountString(account)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.WarnFmt("Non-exist user login: %s", account)
		} else {
			logger.WarnFmt("Failed to get user by account: %s", account)
		}
		return model.User{}, errors.New("account or password is incorrect")
	}
	if !utility.CheckPasswordHash(password, user.Password) {
		logger.WarnFmt("Wrong password login: %s", account)
		return model.User{}, errors.New("account or password is incorrect")
	}
	return user, nil
}

func GetUserByID(id uint) (model.User, error) {
	var user model.User
	err := Database.Model(&model.User{}).
		Where("id = ?", id).
		First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user *model.User) error {
	if UserInitFlag {
		user.Admin = true
		UserInitFlag = false
	}
	return Database.Create(user).Error
}
func UpdateUser(user *model.User) error {
	return Database.Save(user).Error
}

func DeleteUser(userID uint) error {
	return Database.Delete(&model.User{}, userID).Error
}
