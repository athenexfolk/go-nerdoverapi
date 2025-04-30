package util

import (
	"errors"
	"regexp"
)

func ValidateSlug(input string) error {
	matched, err := regexp.MatchString("^[a-z](?:[a-z0-9-]*[a-z])?$", input)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("invalid slug")
	}
	return nil
}
