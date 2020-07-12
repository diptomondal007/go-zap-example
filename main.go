package main

import (
	"go-zap-example/test_package"
	"go-zap-example/test_package_2"
	"go-zap-example/zap_logger"
	"go.uber.org/zap"
)

func main(){
	var aNumber int
	aNumber = 5
	_ = zap_logger.InitLogger()
	if aNumber ==5 {
		zap.L().Info("test logs")
	}else {
		zap.L().Debug("test log not 5")
	}

	var a test_package.A
	a.SetName("")
	a.SetName("test package name")

	var b test_package_2.B
	b.SetName("")
	b.SetName("test package 2 name")
}
