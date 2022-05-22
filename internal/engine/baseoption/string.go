package baseoption

import (
	"errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

type String struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Desc         string  `json:"desc"`
	BaseType     string  `json:"baseType"`
	ExtendedType string  `json:"extendedType"`
	Default      string  `json:"default"`
	Required     bool    `json:"required"`
	Regexp       *string `json:"regexp,omitempty"`
	MinLen       *int    `json:"minLen,omitempty"`
	MaxLen       *int    `json:"maxLen,omitempty"`
}

func NewString(id string) *String {
	return &String{
		Id:       id,
		BaseType: "string",
	}
}

func (s *String) Validate(value any) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("bad value type")
	}

	if s.MinLen != nil && utf8.RuneCountInString(v) < *s.MinLen {
		return fmt.Errorf("minimum lenght value %d", *s.MinLen)
	}

	if s.MaxLen != nil && utf8.RuneCountInString(v) > *s.MaxLen {
		return fmt.Errorf("maximum lenght value %d", *s.MaxLen)
	}

	if s.Regexp != nil {
		ok, err := regexp.MatchString(*s.Regexp, v)
		if err != nil {
			return errors.New("internal error")
		}

		if !ok {
			return errors.New("incorrect value format")
		}
	}

	return nil
}
