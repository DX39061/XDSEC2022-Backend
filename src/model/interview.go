package model

import (
	"gorm.io/gorm"
	"time"
)

type Interview struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdateAt  time.Time      `json:"update-at"`
	DeletedAt gorm.DeletedAt `json:"-"`

	Round         uint   `json:"round"`
	Pass          bool   `json:"pass"`
	Record        string `json:"record"`
	Interviewer   string `json:"interviewer"`
	Interviewee   string `json:"interviewee"`
	InterviewerID uint   `json:"interviewer-id"`
	IntervieweeID uint   `json:"interviewee-id"`
}

type InterviewShort struct {
	ID            uint      `json:"id"`
	UpdateAt      time.Time `json:"update-at"`
	Round         uint      `json:"round"`
	Pass          bool      `json:"pass"`
	Interviewer   string    `json:"interviewer"`
	Interviewee   string    `json:"interviewee"`
	IntervieweeID uint      `json:"interviewee-id"`
}

type InterviewDetail struct {
	ID          uint      `json:"id"`
	UpdateAt    time.Time `json:"update-at"`
	Round       uint      `json:"round"`
	Pass        bool      `json:"pass"`
	Interviewer string    `json:"interviewer"`
	Interviewee string    `json:"interviewee"`
	Record      string    `json:"record"`
}
