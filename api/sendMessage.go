package api

import (
	"errors"
	"github.com/menghua6/token-net/db"
)

func SendMessage(token string, message string) (string, error) {
	tokenInfo, err := db.GetToken(token)
	if err != nil {
		return "", err
	}
	if !checkLimit(tokenInfo) {
		return "", errors.New("token has expired")
	}
	if tokenInfo.Count > 0 {
		tokenInfo.Count = tokenInfo.Count - 1
		err = db.UpdateToken(tokenInfo)
	}
	if err != nil {
		return "", err
	}
	err = db.CreateMessage(token, message)
	if err != nil {
		return "", err
	}
	messages, err := db.GetMessages(token)
	if err != nil {
		return "", err
	}
	messagesTran, err := messagesToString(tokenInfo, messages)
	if err != nil {
		return "", err
	}
	return messagesTran, nil
}
