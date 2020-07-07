package common

import "github.com/google/logger"
func ErrHandler(msg string,err error){
	if err != nil{
		logger.Errorln(err)
	}
}

func PanicErr(err error){
	if err !=nil{
		logger.Fatalln("程序挂了,",err)
	}
}