package utils

import (
	"ampl/src/config"
	"fmt"
	"strings"
)

func IsRelease() bool {
	fmt.Println(config.Config.Env)
	return strings.EqualFold(config.Config.Env, PROD)
}
