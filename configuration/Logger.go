package configuration

import (
	"github.com/rs/zerolog"
	"os"
)

func InitLogger() zerolog.Logger {
	runLogFile, _ := os.OpenFile(
		"myapp.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	output := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
	return zerolog.New(output).With().Timestamp().Caller().Logger()
}
