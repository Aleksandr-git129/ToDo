package model

import (
	"errors"
	"strings"
	"unicode/utf8"
)

type UserDTO struct {
	Fullname    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

type PatchUserDTO struct {
	Fullname    *string `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (u *UserDTO) Validate() error {
	if u == nil {
		return errors.New("данные пользователя обязательны")
	}
	if err := validateFullname(u.Fullname); err != nil {
		return err
	}
	return validatePhoneNumber(u.PhoneNumber)
}

func (u *PatchUserDTO) Validate() error {
	if u == nil {
		return errors.New("данные пользователя обязательны")
	}
	if u.Fullname != nil {
		if err := validateFullname(*u.Fullname); err != nil {
			return err
		}
	}
	if u.PhoneNumber != nil {
		return validatePhoneNumber(*u.PhoneNumber)
	}
	return nil
}

func validateFullname(name string) error {
	name = strings.TrimSpace(name)
	nameLength := utf8.RuneCountInString(name)
	if nameLength < 3 || nameLength > 100 {
		return errors.New("full_name должен содержать от 3 до 100 символов")
	}
	return nil
}

func validatePhoneNumber(phone string) error {
	if phone == "" {
		return nil
	}
	phoneLength := utf8.RuneCountInString(phone)
	if phoneLength < 10 || phoneLength > 15 {
		return errors.New("phone_number должен содержать от 10 до 15 символов")
	}
	if !strings.HasPrefix(phone, "+") {
		return errors.New("phone_number должен начинаться со знака +")
	}
	for _, digit := range phone[1:] {
		if digit < '0' || digit > '9' {
			return errors.New("phone_number после знака + должен содержать только цифры")
		}
	}
	return nil
}
