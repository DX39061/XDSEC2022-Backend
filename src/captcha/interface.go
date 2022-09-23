package captcha

import (
	"XDSEC2022-Backend/src/config"
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type VerifyCaptchaResponse struct {
	Success bool `json:"success"`
}

func VerifyCaptcha(token string) (bool, error) {
	url := "https://recaptcha.net/recaptcha/api/siteverify?secret=" +
		config.CaptchaConfig.SecretKey + "&response=" + token
	client := &http.Client{}
	request, err := http.NewRequest("POST", url, &strings.Reader{})
	if err != nil {
		return false, err
	}
	resp, err := client.Do(request)
	if err != nil {
		return false, err
	}
	var body bytes.Buffer
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return false, err
	}
	var result VerifyCaptchaResponse
	err = json.Unmarshal(body.Bytes(), &result)
	if err != nil {
		return false, err
	}
	return result.Success, nil
}
