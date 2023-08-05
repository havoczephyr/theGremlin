package logic

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func Suggest(args, userName, url, token string, port int) error {
	if strings.TrimSpace(args) == " " {
		return errors.New("empty suggestion")
	} else {
		requestUrl := url + ":" + strconv.Itoa(port) + "/api/createNew"

		requestBody, err := json.Marshal(map[string]interface{}{
			"body":   args,
			"author": userName,
		})
		req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(requestBody))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "theGremlin "+token)
	}

	return nil
}
