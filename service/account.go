package service

import (
	"encoding/json"
	"fmt"
	"github.com/AbnerEarl/goutils/files"
	"github.com/AbnerEarl/goutils/utils"
	"sso-demo/config"
	"sso-demo/model"
	"time"
	"unicode"
)

func AccountLogin(username, password string) (string, error) {
	md5 := files.String2Md5(password)
	account := model.AccountModel{Username: username}
	err := model.DB.RetrieveFirstByFind(&account)
	fmt.Println(account)
	if err != nil {
		return "", err
	}
	if account.Status != 1 {
		return "", config.ErrUserStatus
	}
	if account.Password != md5 {
		return "", config.ErrVerifyLogin
	}
	account.UpdatedAt = time.Now()
	account.LastLogin = time.Now()
	model.DB.UpdateById(&account)
	data, err := json.Marshal(account)
	if err != nil {
		return "", err
	}
	result, err := files.EncryptByAes(data)
	if err != nil {
		return "", err
	}
	return result, nil
}

func RepeatLogin(username string) (string, error) {
	account := model.AccountModel{Username: username}
	err := model.DB.RetrieveFirstByFind(&account)
	fmt.Println(account)
	if err != nil {
		return "", err
	}
	if account.Status != 1 {
		return "", config.ErrUserStatus
	}
	account.UpdatedAt = time.Now()
	account.LastLogin = time.Now()
	model.DB.UpdateById(&account)
	data, err := json.Marshal(account)
	if err != nil {
		return "", err
	}
	result, err := files.EncryptByAes(data)
	if err != nil {
		return "", err
	}
	return result, nil
}

func CheckPassword(password string) error {
	if len(password) < 10 || utils.StrLength(password) > 20 {
		return config.ErrPasswordLength
	}
	var digit, upper, lower, special bool
	for _, s := range password {
		if unicode.IsLower(s) {
			lower = true
		} else if unicode.IsUpper(s) {
			upper = true
		} else if unicode.IsDigit(s) {
			digit = true
		} else {
			special = true
		}
	}
	if digit && upper && lower && special {
		return nil
	}
	return config.ErrPasswordDisable
}

func AllAccountMap() (map[uint64]model.AccountModel, error) {
	bys, _, err := model.DB.RetrieveByWhereBytes(config.DefaultMaxLimit, 0, &model.AccountModel{}, "", "", nil)
	if err != nil {
		return nil, err
	}
	var accountList []model.AccountModel
	err = json.Unmarshal(bys, &accountList)
	if err != nil {
		return nil, err
	}
	accountMap := map[uint64]model.AccountModel{}
	for _, account := range accountList {
		accountMap[account.Id] = account
	}
	return accountMap, nil
}
