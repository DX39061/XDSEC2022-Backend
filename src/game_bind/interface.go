package game_bind

import (
	"XDSEC2022-Backend/src/config"
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GetUserDataOfGame(studentID string) (GameDataResponse, error) {
	client := &http.Client{}
	url := config.GameBindConfig.PlatformUrl +
		"?game=" + strconv.Itoa(config.GameBindConfig.GameID) +
		"&student_id=" + studentID
	request, err := http.NewRequest("GET", url, &strings.Reader{})
	if err != nil {
		return GameDataResponse{}, nil
	}
	request.Header.Add("Authorization", "Bearer "+config.GameBindConfig.AuthToken)
	resp, err := client.Do(request)
	if err != nil {
		return GameDataResponse{}, err
	}
	response, err := parseResponse(resp)
	if err != nil {
		return GameDataResponse{}, err
	}
	var data GameDataResponse
	if response["data"] == nil {
		return GameDataResponse{}, errors.New("record not found")
	}
	err = mapstructure.Decode(response["data"], &data)
	if err != nil {
		return GameDataResponse{}, err
	}
	return data, nil
}

func parseResponse(response *http.Response) (map[string]interface{}, error) {
	var result map[string]interface{}
	if response.Body == http.NoBody {
		return result, nil
	}
	body, err := io.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}

	return result, err
}
