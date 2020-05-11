package main

import (
	"fmt"
	"runtime"
)

// TODO: These functions must be replaced with log package or another log library.

var base_str string = "%s:[%s]\n"

func MyPrintf(key string, val string) {
	fmt.Printf(base_str, key, val)
}

func MyPrintErr(err error) {
	pt, _, _, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pt)
	MyPrintf("ERROR", f.Name())
	fmt.Println(err)
}
