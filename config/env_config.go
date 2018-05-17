package config

//import (
//	"github.com/joho/godotenv"
//	"path/filepath"
//	"runtime"
//	"fmt"
//)
//
//const envFile = "conf.env"
//
//// set app run path, notice : runtime.Caller param skip
//func init() {
//	// get path
//	_, file, _, _ := runtime.Caller(0)
//	currentPath := filepath.Join(file, ".."+string(filepath.Separator))
//	envFile := filepath.Join(currentPath, envFile)
//	// set env
//	if err := godotenv.Load(envFile); err != nil {
//		fmt.Println("godotenv.Load error : ", err.Error())
//	}
//}
