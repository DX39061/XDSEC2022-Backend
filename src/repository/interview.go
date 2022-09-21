package repository

import (
	"XDSEC2022-Backend/src/model"
	"github.com/jinzhu/copier"
)

func init() {
	RegisterModel(model.Interview{})
}

func GetInterviewCount() (int64, error) {
	var count int64
	err := Database.Model(&model.Interview{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetInterviewList(limit int, skip int) ([]model.InterviewShort, error) {
	var interviews []model.Interview
	err := Database.Model(&model.Interview{}).
		Offset(skip).
		Limit(limit).
		Order("id asc").
		Find(&interviews).Error
	if err != nil {
		return nil, err
	}
	var interviewsShort []model.InterviewShort
	err = copier.Copy(&interviewsShort, interviews)
	if err != nil {
		return nil, err
	}
	return interviewsShort, nil
}

func GetInterviewDetail(id uint) (model.InterviewDetail, error) {
	var interview model.Interview
	err := Database.Model(&model.Interview{}).
		Where("id = ?", id).
		First(&interview).Error
	if err != nil {
		return model.InterviewDetail{}, err
	}
	var interviewDetail model.InterviewDetail
	err = copier.Copy(&interviewDetail, interview)
	if err != nil {
		return model.InterviewDetail{}, err
	}
	return interviewDetail, nil
}

func GetInterviewDetailOfUser(userID uint) ([]model.InterviewDetail, error) {
	var interviews []model.Interview
	err := Database.Model(&model.Interview{}).
		Where("interviewee_id = ?", userID).
		Order("created_at asc").
		Find(&interviews).Error
	if err != nil {
		return nil, err
	}
	var interviewsDetail []model.InterviewDetail
	err = copier.Copy(&interviewsDetail, interviews)
	if err != nil {
		return nil, err
	}
	return interviewsDetail, nil
}

func GetInterview(id uint) (model.Interview, error) {
	var interview model.Interview
	err := Database.Model(&model.Interview{}).
		Where("id = ?", id).
		First(&interview).Error
	if err != nil {
		return model.Interview{}, err
	}
	return interview, nil
}

func CreateInterview(interview *model.Interview) error {
	return Database.Create(interview).Error
}

func UpdateInterview(interview *model.Interview) error {
	return Database.Save(interview).Error
}

func DeleteInterview(id uint) error {
	return Database.Delete(&model.Interview{}, id).Error
}
