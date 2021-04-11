package cuttly

import "errors"

func checkErrorCode(errorCode int, isURL bool) error {
	switch errorCode {
	case 0:
		return errors.New("this shortened link does not exist")
	case 1:
		return errors.New("the shortened link comes from the domain that shortens the link, i.e. the link has already been shortened")
	case 2:
		return errors.New("the entered link is not a link")
	case 3:
		return errors.New("the preferred link name is already taken")
	case 4:
		return errors.New("invalid API key")
	case 5:
		return errors.New("the link has not passed the validation. Includes invalid characters")
	case 6:
		return errors.New("the link provided is from a blocked domain")
	}

	return nil
}
