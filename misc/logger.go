package misc

import (
	"fmt"
	"os"
	"strings"

	"github.com/micro/go-micro/v2/logger"
	zl "github.com/micro/go-plugins/logger/zerolog/v2"
	"github.com/rs/zerolog"
)

func Logger() logger.Logger {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	output := zerolog.ConsoleWriter{Out: os.Stdout}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}

	return zl.NewLogger(logger.WithOutput(output),
		logger.WithLevel(logger.DebugLevel))
}
