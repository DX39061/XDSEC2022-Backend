package model

type User struct {
	ID                 uint   `json:"id"`
	Admin              bool   `json:"admin"`
	Password           string `json:"password"`
	Name               string `json:"name"`
	Sex                string `json:"sex"`
	Major              string `json:"major"` // 学院/专业
	StudentID          string `gorm:"unique" json:"student-id"`
	Telephone          string `gorm:"unique" json:"telephone"`
	Email              string `gorm:"unique" json:"email"`
	Department         string `json:"department"`          // 意向部门
	Direction          string `json:"direction"`           // 学习方向
	LearnedTechnique   string `json:"learned-technique"`   // 技术基础
	LearningExperience string `json:"learning-experience"` // 学习经历
	HobbyAndAdvantage  string `json:"hobby-and-advantage"` // 爱好特长
}

type UserShort struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Major string `json:"major"` // 学院/专业
}

type UserDetail struct {
	ID                 uint   `json:"id"`
	Admin              bool   `json:"admin"`
	Name               string `json:"name"`
	Sex                string `json:"sex"`
	Major              string `json:"major"` // 学院/专业
	StudentID          string `gorm:"unique" json:"student-id"`
	Telephone          string `gorm:"unique" json:"telephone"`
	Email              string `gorm:"unique" json:"email"`
	Department         string `json:"department"`          // 意向部门
	Direction          string `json:"direction"`           // 学习方向
	LearnedTechnique   string `json:"learned-technique"`   // 技术基础
	LearningExperience string `json:"learning-experience"` // 学习经历
	HobbyAndAdvantage  string `json:"hobby-and-advantage"` //爱好特长
}