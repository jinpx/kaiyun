package caller

import "hypers/internal/keyfunc"

func Caller() {
	kf := new(keyfunc.KeyFunction)
	kf.Key = "read_20231207.exe"
	kf.CallExecute()
}
