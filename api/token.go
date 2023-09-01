package api

import (
	"errors"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
	"github.com/menghua6/token-net/db"
)

var (
	Forever  = time.Date(4000, time.December, 0, 0, 0, 0, 0, time.Local)
	Infinity = -1
)

func Token(limit string) (string, error) {
	expirationTime, count, err := explainLimit(limit)
	if err != nil {
		return "", err
	}
	token := generateToken()
	if err = db.CreateToken(token, expirationTime, count); err != nil {
		return "", err
	}
	return token, nil
}

func explainLimit(limit string) (time.Time, int, error) {
	if count, err := strconv.Atoi(limit); err == nil {
		//count
		return Forever, count, nil
	} else if strings.Contains(limit, "+") {
		//yyyy-mm-dd+hh:mm
		if len(limit) == 16 {
			if y, err := strconv.Atoi(limit[0:4]); err == nil {
				if m, err := strconv.Atoi(limit[5:7]); err == nil {
					if d, err := strconv.Atoi(limit[8:10]); err == nil {
						if h, err := strconv.Atoi(limit[11:13]); err == nil {
							if min, err := strconv.Atoi(limit[14:16]); err == nil {
								expirationTime := time.Date(y, time.Month(m), d, h, min, 0, 0, time.Local)
								return expirationTime, Infinity, nil
							}
						}
					}
				}
			}
		}
	} else {
		//YhZm
		if duration, err := time.ParseDuration(limit); err == nil {
			return time.Now().Add(duration), Infinity, nil
		}
	}
	return time.Time{}, 0, errors.New("limit formatting error")
}

func generateToken() string {
	uuidA := uuid.New().String()
	uuidA = strings.ReplaceAll(uuidA, "-", "")
	uuidB := uuid.New().String()
	uuidB = strings.ReplaceAll(uuidB, "-", "")
	return uuidA + uuidB
}
