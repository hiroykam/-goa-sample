package services

import (
	"github.com/hiroykam/goa-sample/sample_error"
	"golang.org/x/crypto/bcrypt"
)

func Generate(password string) (string, *sample_error.SampleError) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", sample_error.NewSampleError(sample_error.InternalError, err.Error())
	}
	return string(hash), nil
}

func Confirm(hash, password string) *sample_error.SampleError {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return sample_error.NewSampleError(sample_error.WrongPassword, err.Error())
		}
		return sample_error.NewSampleError(sample_error.InternalError, err.Error())
	}
	return nil
}
