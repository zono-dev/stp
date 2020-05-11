package main

import (
	"os"
	"reflect"
	"strconv"
)

// MyEnv stores environment values.
type MyEnv struct {
	Vlog    bool   `env:"VLOG"`
	Prefix  string `env:"PREFIX"`
	ResizeX int    `env:"RESIZEX"`
	ResizeY int    `env:"RESIZEY"`
	Dtable  string `env:"DTABLE"`
}

// checkEnv checks the key in environment values.
func checkEnv(key string) (bool, string) {
	tmp := os.Getenv(key)
	if tmp == "" {
		return false, ""
	}
	return true, tmp
}

// setValue sets environment values in MyEnv struct.
// It use checkEnv to check the same name `env` tag and a environment value.
func setValue(env *MyEnv) {
	rtenv := reflect.TypeOf(*env)
	rvenv := reflect.ValueOf(env)

	for i := 0; i < rtenv.NumField(); i++ {
		// os.Getenv by struct tag name
		MyPrintf("Tag", rtenv.Field(i).Tag.Get("env"))
		if b, s := checkEnv(rtenv.Field(i).Tag.Get("env")); b {
			// Get value by field name
			v := rvenv.Elem().FieldByName(rtenv.Field(i).Name)

			// Check typeof the field of env and Set the value
			if _, ok := v.Interface().(bool); ok {
				if s == "true" {
					v.SetBool(true)
				} else {
					v.SetBool(false)
				}
			}

			if _, ok := v.Interface().(int); ok {
				if int_v, err := strconv.ParseInt(s, 10, 64); err == nil {
					v.SetInt(int_v)
				}
			}

			if _, ok := v.Interface().(string); ok {
				v.SetString(s)
			}

		}
	}
}

// Setup checks environment value.
func Setup() MyEnv {
	defenv := MyEnv{Vlog: false, Prefix: "resized-", ResizeX: 900, ResizeY: 900, Dtable: "ImageList"}
	env := MyEnv{}
	env = defenv
	setValue(&env)
	return env
}
