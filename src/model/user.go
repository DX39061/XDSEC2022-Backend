package model

type User struct {
	ID                 uint        `json:"id"`
	Admin              bool        `json:"admin"`
	Password           string      `json:"password"`
	NickName           string      `json:"nick-name"`
	Name               string      `json:"name"`
	Sex                string      `json:"sex"`
	Major              string      `json:"major"` // 学院/专业
	StudentID          string      `gorm:"unique" json:"student-id"`
	QQ                 string      `gorm:"unique" json:"qq"`
	Telephone          string      `gorm:"unique" json:"telephone"`
	Email              string      `gorm:"unique" json:"email"`
	Department         string      `json:"department"`          // 意向部门
	Direction          string      `json:"direction"`           // 学习方向
	LearnedTechnique   string      `json:"learned-technique"`   // 技术基础
	LearningExperience string      `json:"learning-experience"` // 学习经历
	HobbyAndAdvantage  string      `json:"hobby-and-advantage"` // 爱好特长
	Interviews         []Interview `gorm:"foreignKey:IntervieweeID" json:"-"`
}

type UserShort struct {
	ID       uint   `json:"id"`
	NickName string `json:"nick-name"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Major    string `json:"major"` // 学院/专业
}

type UserDetail struct {
	ID                 uint   `json:"id"`
	Admin              bool   `json:"admin"`
	NickName           string `json:"nick-name"`
	Name               string `json:"name"`
	Sex                string `json:"sex"`
	Major              string `json:"major"` // 学院/专业
	StudentID          string `gorm:"unique" json:"student-id"`
	Telephone          string `gorm:"unique" json:"telephone"`
	QQ                 string `gorm:"unique" json:"qq"`
	Email              string `gorm:"unique" json:"email"`
	Department         string `json:"department"`          // 意向部门
	Direction          string `json:"direction"`           // 学习方向
	LearnedTechnique   string `json:"learned-technique"`   // 技术基础
	LearningExperience string `json:"learning-experience"` // 学习经历
	HobbyAndAdvantage  string `json:"hobby-and-advantage"` //爱好特长
}

type UserCount struct {
	Total       int64 `json:"total"`
	Admin       int64 `json:"admin"`
	Male        int64 `json:"male"`
	Female      int64 `json:"female"`
	Web         int64 `json:"web"`
	Reverse     int64 `json:"reverse"`
	Pwn         int64 `json:"pwn"`
	Crypto      int64 `json:"crypto"`
	Misc        int64 `json:"misc"`
	Dev         int64 `json:"dev"`
	Secretariat int64 `json:"secretariat"`
	Technique   int64 `json:"technique"`
	Devops      int64 `json:"devops"`
	Publicity   int64 `json:"publicity"`
}
