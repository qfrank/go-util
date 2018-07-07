package goutil

import (
	"fmt"
	"github.com/pkg/errors"
)

func PanicIfErr(err error) {
	if err != nil {
		if _, ok := err.(fmt.Formatter); ok {
			panic(err)
		} else {
			panic(errors.WithStack(err))
		}
	}
}