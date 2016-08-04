package syService

import "log"

func SysInit(){
	initLog()	//初始化日志
}

func initLog(){
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

