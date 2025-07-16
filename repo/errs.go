package repo

import "errors"

var (
	invalidName     = errors.New("недопустимое имя")
	dirNotSpecified = errors.New("директория не указана")
)
