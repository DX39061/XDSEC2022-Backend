package email

import (
	"XDSEC2022-Backend/src/config"
	"gopkg.in/gomail.v2"
)

func SendVerifyCode(url string, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.EmailConfig.Email_from)
	m.SetHeader("To", url)
	m.SetHeader("Subject", "验证您的邮箱地址")
	m.SetBody("text/html", "验证您的邮箱地址，这是您的邮箱验证码: "+code+" ，有效期五分钟，请注意不要向他人透露!")
	d := gomail.NewDialer(
		config.EmailConfig.Smtp_url,
		config.EmailConfig.Smtp_port,
		config.EmailConfig.Email_from,
		config.EmailConfig.Email_password,
	)
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
