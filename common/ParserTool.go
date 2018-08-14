package common

import (
	"errors"
	"image/color"
	"strconv"
	"strings"
	"time"
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
func Parser_color(contentStr string, attachmentValue *color.RGBA) error {
	content := strings.TrimPrefix(contentStr, "#")
	contentLen := strings.Count(content, "") - 1
	if contentLen != 6 && contentLen != 8 { // (r + g + b) * 2
		return errors.New("error on parser color")
	}

	color64, err := strconv.ParseInt(content, 16, 64) //字串到数据整型
	if err != nil {
		return err
	}
	if contentLen == 8 {
		red := color64 >> 24
		green := (color64 & 0x00FF0000) >> 16
		blue := (color64 & 0x0000FF00) >> 8
		alpha := color64 & 0x000000FF
		*attachmentValue = color.RGBA{uint8(red), uint8(green), uint8(blue), uint8(alpha)}
	} else {
		red := color64 >> 16
		green := (color64 & 0x00FF00) >> 8
		blue := color64 & 0x0000FF
		*attachmentValue = color.RGBA{uint8(red), uint8(green), uint8(blue), 0}
	}

	return nil
}
func Parser_dateTime(content string, attachmentValue *time.Time) error {
	parse_str_time, err := time.Parse("2006-01-02 15:04:05", content)
	if nil != err {
		return err
	}
	*attachmentValue = parse_str_time
	return nil
}
func CheckValueLimit_dateTime(value time.Time, limitMin string, limitMax string) error {
	var min time.Time
	var max time.Time
	if limitMin != "" {
		err := Parser_dateTime(limitMin, &min)
		if nil != err {
			return err
		}

		if value.Before(min) {
			return errors.New("error on check limit ")
		}
	}
	if limitMax != "" {
		err := Parser_dateTime(limitMax, &max)
		if nil != err {
			return err
		}
		if value.After(max) {
			return errors.New("error on check limit ")
		}
	}
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
