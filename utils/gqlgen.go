package utils

import (
	"github.com/99designs/gqlgen/client"
)

func VarsOption(vars map[string]interface{}) client.Option {
	return func(bd *client.Request) {
		if bd.Variables == nil {
			bd.Variables = map[string]interface{}{}
		}
		bd.Variables = vars
	}
}
