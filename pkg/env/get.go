package env

import (
	"github.com/cockroachdb/errors"
	"os"
)

func Get(key string) (string, error) {
	v := os.Getenv(key)
	if len(v) == 0 {
		return "", errors.Newf("%s is not set", key)
	}

	return v, nil
}
