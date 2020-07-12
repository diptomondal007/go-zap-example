package test_package

import "go.uber.org/zap"

type A struct {
	Name string
}

func(a *A)SetName(name string){
	if name == ""{
		zap.L().Warn("no name is entered")
		return
	}
	a.Name = name
	zap.L().Info(a.Name)
}
