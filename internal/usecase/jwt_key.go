package usecase

import (
	"errors"
	"fmt"
)

func (uc *Usecase) jwtGetSignKey(userID string, refresh string) ([]byte, error) {
	conn := uc.redis.Client()
	defer conn.Close()

	reply, err := conn.Do("GET", formatKey(userID, refresh))
	if err != nil {
		return nil, err
	}

	r, ok := reply.([]byte)
	if !ok {
		return nil, errors.New("reply its not a []byte")
	}
	return r, err
}

func (uc *Usecase) jwtSaveSignKey(userID string, refresh string, signKey string, exp int) error {
	conn := uc.redis.Client()
	defer conn.Close()

	_, err := conn.Do("SET", formatKey(userID, refresh), []byte(signKey), "EX", exp)
	return err
}

func (uc *Usecase) jwtDeleteSignKey(userID string, refresh string) error {
	conn := uc.redis.Client()
	defer conn.Close()

	_, err := conn.Do("DEL", formatKey(userID, refresh))
	return err
}

func formatKey(userID, refresh string) string {
	return fmt.Sprintf("%s-%s", userID, refresh)
}
