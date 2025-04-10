package goenvlib

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var mut sync.Mutex

type envRegistered struct {
	updater func()
	addr    interface{}
}

var envMap = make(map[string]envRegistered)

func ReloadEnv() {
	mut.Lock()
	defer mut.Unlock()
	for _, reg := range envMap {
		reg.updater()
	}
}

func getenv[T any](env string, defaultVal T, converter func(string) (T, error)) (T, error) {
	ret := os.Getenv(env)
	if ret == "" {
		return defaultVal, nil
	}
	return converter(ret)
}

func mustGetenv[T any](env string, defaultVal T, converter func(string) (T, error)) *T {
	mut.Lock()
	defer mut.Unlock()

	reg, ok := envMap[env]
	if ok {
		return reg.addr.(*T)
	}
	var ret T
	updater := func() {
		var err error
		ret, err = getenv(env, defaultVal, converter)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting environment variable %s: %v, using default value %v\n", env, err, defaultVal)
			ret = defaultVal
		}
	}
	updater()
	reg.updater = updater
	reg.addr = &ret
	envMap[env] = reg
	return &ret
}

func GetenvString(env string, def string) *string {
	return mustGetenv(env, def, func(s string) (string, error) {
		return s, nil
	})
}

func GetenvInt(env string, def int) *int {
	return mustGetenv(env, def, func(s string) (int, error) {
		return strconv.Atoi(s)
	})
}

func GetenvFloat64(env string, def float64) *float64 {
	return mustGetenv(env, def, func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	})
}

func GetenvBool(env string, def bool) *bool {
	return mustGetenv(env, def, func(s string) (bool, error) {
		return strconv.ParseBool(s)
	})
}

func GetenvStringSlice(env string, def []string) *[]string {
	return mustGetenv(env, def, func(s string) ([]string, error) {
		return strings.Split(s, ","), nil
	})
}

func GetenvIntSlice(env string, def []int) *[]int {
	return mustGetenv(env, def, func(s string) ([]int, error) {
		strarray := strings.Split(s, ",")
		intarray := make([]int, len(strarray))
		var err error
		for i, str := range strarray {
			intarray[i], err = strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
		}
		return intarray, nil
	})
}

func GetenvFloat64Slice(env string, def []float64) *[]float64 {
	return mustGetenv(env, def, func(s string) ([]float64, error) {
		strarray := strings.Split(s, ",")
		floatarray := make([]float64, len(strarray))
		var err error
		for i, str := range strarray {
			floatarray[i], err = strconv.ParseFloat(str, 64)
			if err != nil {
				return nil, err
			}
		}
		return floatarray, nil
	})
}

func GetenvBoolSlice(env string, def []bool) *[]bool {
	return mustGetenv(env, def, func(s string) ([]bool, error) {
		strarray := strings.Split(s, ",")
		boolarray := make([]bool, len(strarray))
		var err error
		for i, str := range strarray {
			boolarray[i], err = strconv.ParseBool(str)
			if err != nil {
				return nil, err
			}
		}
		return boolarray, nil
	})
}
