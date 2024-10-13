package utils

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

func ValidateName(fieldName, name string, allowNonLetters bool, minLength, maxLength int) error {

	switch length := len(name); {
	case length < minLength:
		return errors.New(fieldName + " must be at least " + strconv.Itoa(minLength) + " characters long")
	case length > maxLength:
		return errors.New(fieldName + " cannot be longer than " + strconv.Itoa(maxLength) + " characters")
	}

	// If non-letters are not allowed, check the regex
	if !allowNonLetters {
		const nameRegex = `^[a-zA-Z]+$`
		re := regexp.MustCompile(nameRegex)

		if !re.MatchString(name) {
			return errors.New(fieldName + " must contain only letters")
		}
	}

	return nil
}

func ValidateTimeOrder(startTime, endTime time.Time) error {
	if endTime.Before(startTime) {
		return errors.New("end time must be after start time")
	}
	return nil
}
