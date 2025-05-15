package env

import (
	"github.com/cockroachdb/errors"
	"strconv"
)

type mutantType interface {
	int | bool | []byte
}

// MuGet Мутирует строку из env в переданный через дженерики тип
func MuGet[T mutantType](key string) (T, error) {
	var mutant T
	var rawValue string
	var err error

	if rawValue, err = Get(key); err != nil {
		return mutant, err
	}

	switch any(mutant).(type) {
	case int:
		i, err := strconv.Atoi(rawValue)
		if err != nil {
			return mutant, err
		}
		return any(i).(T), nil
	case bool:
		parseBool, err := strconv.ParseBool(rawValue)
		if err != nil {
			return mutant, err
		}

		return any(parseBool).(T), nil
	case []byte:
		return any([]byte(rawValue)).(T), nil
	default:
		return mutant, errors.Newf("not support type: %T", any(mutant))
	}
}
