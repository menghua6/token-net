package api

import (
	"errors"
	"strconv"
	"time"
	"github.com/menghua6/token-net/db"
	"github.com/menghua6/token-net/entity"
)

func CheckMessages(token string) (string, error) {
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

func messagesToString(token *entity.Token, messages []entity.Message) (string, error) {
	mms := "口令过期时间：" + token.ExpirationTime.String()[:19] + "\n"
	if token.Count >= 0 {
		mms += "剩余可访问次数：" + strconv.Itoa(token.Count) + "\n"
	}
	for i := 0; i < len(messages); i++ {
		mms += "\n" + messages[i].CreatedAt.String()[:19] + "\n"
		mms += messages[i].Message + "\n"
	}
	return mms, nil
}

func checkLimit(token *entity.Token) bool {
	if time.Now().After(token.ExpirationTime) || token.Count == 0 {
		return false
	}
	return true
}
