package test_package_2

import "go.uber.org/zap"

type B struct {
	Name string
}

func(b *B)SetName(name string){
	if name == ""{
		zap.L().Warn("no name is entered")
		return
	}
	b.Name = name
	zap.L().Info(b.Name)
}

