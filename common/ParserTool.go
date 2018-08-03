package common

import (
	"errors"
	"strconv"
)

func Parser_int32(content string, attachmentValue *int32) error {
	result, err := strconv.ParseInt(content, 10, 32)
	if err != nil {
		return err
	}
	*attachmentValue = int32(result)
	return nil
}
func Parser_int64(content string, attachmentValue *int64) error {
	result, err := strconv.ParseInt(content, 10, 64)
	if err != nil {
		return err
	}
	*attachmentValue = result
	return nil
}
func Parser_float32(content string, attachmentValue *float32) error {
	result, err := strconv.ParseFloat(content, 32)
	if err != nil {
		return err
	}
	*attachmentValue = float32(result)
	return nil
}
func Parser_float64(content string, attachmentValue *float64) error {
	result, err := strconv.ParseFloat(content, 64)
	if err != nil {
		return err
	}
	*attachmentValue = result
	return nil
}
func Parser_bool(content string, attachmentValue *bool) error {
	if content == "true" {
		*attachmentValue = true
	} else if content == "false" {
		*attachmentValue = true
	} else {
		return errors.New("can't parser bool " + content)
	}
	return nil
}
func Parser_string(content string, attachmentValue *string) error {

	*attachmentValue = content
	return nil
}

func CheckValueLimit_int32(value int32, limitMin string, limitMax string) error {
	var min int32 = 0
	var max int32 = 0
	if limitMin != "" {
		err := Parser_int32(limitMin, &min)
		if nil != err {
			return err
		}

		if value < min {
			return errors.New("error on check limit ")
		}
	}
	if limitMax != "" {
		err := Parser_int32(limitMax, &max)
		if nil != err {
			return err
		}
		if value > max {
			return errors.New("error on check limit ")
		}
	}
	return nil
}
func CheckValueLimit_int64(value int64, limitMin string, limitMax string) error {
	var min int64 = 0
	var max int64 = 0
	if limitMin != "" {
		err := Parser_int64(limitMin, &min)
		if nil != err {
			return err
		}

		if value < min {
			return errors.New("error on check limit ")
		}
	}
	if limitMax != "" {
		err := Parser_int64(limitMax, &max)
		if nil != err {
			return err
		}
		if value > max {
			return errors.New("error on check limit ")
		}
	}
	return nil
}
func CheckValueLimit_float32(value float32, limitMin string, limitMax string) error {
	var min float32 = 0
	var max float32 = 0
	if limitMin != "" {
		err := Parser_float32(limitMin, &min)
		if nil != err {
			return err
		}

		if value < min {
			return errors.New("error on check limit ")
		}
	}
	if limitMax != "" {
		err := Parser_float32(limitMax, &max)
		if nil != err {
			return err
		}
		if value > max {
			return errors.New("error on check limit ")
		}
	}
	return nil
}
func CheckValueLimit_float64(value float64, limitMin string, limitMax string) error {
	var min float64 = 0
	var max float64 = 0
	if limitMin != "" {
		err := Parser_float64(limitMin, &min)
		if nil != err {
			return err
		}

		if value < min {
			return errors.New("error on check limit ")
		}
	}
	if limitMax != "" {
		err := Parser_float64(limitMax, &max)
		if nil != err {
			return err
		}
		if value > max {
			return errors.New("error on check limit ")
		}
	}
	return nil
}
func IsTypeCanCheckLimit(fieldType string) bool {
	if fieldType == "int32" ||
		fieldType == "int64" ||
		fieldType == "float32" ||
		fieldType == "float64" {
		return true
	}
	return false
}
