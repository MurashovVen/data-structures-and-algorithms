package collections

import "errors"

var (
	ErrEmpty = errors.New("коллекция пуста")
	NotFound = errors.New("не найден")
)
