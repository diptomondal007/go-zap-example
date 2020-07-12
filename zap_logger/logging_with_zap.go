package zap_logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() func(){
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
    //encoderConfig := zap.NewProductionEncoderConfig()
	//encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	//	nanos := t.UnixNano()
	//	millis := nanos / int64(time.Millisecond)
	//	enc.AppendInt64(millis)
	//}
	//config.EncoderConfig = encoderConfig
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	return zap.ReplaceGlobals(logger)
}