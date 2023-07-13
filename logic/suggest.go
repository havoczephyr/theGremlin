package logic

import (
	"errors"
	"strings"
)

func Suggest(args string, url string, port int) error {
	if strings.TrimSpace(args) == " " {
		return errors.New("empty suggestion")
	}

	return nil
}
