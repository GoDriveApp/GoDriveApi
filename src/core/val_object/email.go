package val_object

import "regexp"

type Email struct {
	value string
}

const (
	regexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

func NewEmail(value string) (error, *Email) {
	if isValid(value) {
		return nil, &Email{value}
	}
	return nil, nil
}

func isValid(value string) bool {
	result, _ := regexp.MatchString(regexPattern, value)
	return result
}

func (e Email) IsValid() bool {
	return isValid(e.value)
}

func (e Email) Value() string {
	return e.value
}
