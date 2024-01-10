package logs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	
	"os"

	"go.uber.org/zap"
)

// cfg := zap.Config{
// }

// type Config struct {
// 	// Level is the minimum enabled logging level. Note that this is a dynamic
// 	// level, so calling Config.Level.SetLevel will atomically change the log
// 	// level of all loggers descended from this config.
// 	Level AtomicLevel `json:"level" yaml:"level"`
// 	// Development puts the logger in development mode, which changes the
// 	// behavior of DPanicLevel and takes stacktraces more liberally.
// 	Development bool `json:"development" yaml:"development"`
// 	// DisableCaller stops annotating logs with the calling function's file
// 	// name and line number. By default, all logs are annotated.
// 	DisableCaller bool `json:"disableCaller" yaml:"disableCaller"`
// 	// DisableStacktrace completely disables automatic stacktrace capturing. By
// 	// default, stacktraces are captured for WarnLevel and above logs in
// 	// development and ErrorLevel and above in production.
// 	DisableStacktrace bool `json:"disableStacktrace" yaml:"disableStacktrace"`
// 	// Sampling sets a sampling policy. A nil SamplingConfig disables sampling.
// 	Sampling *SamplingConfig `json:"sampling" yaml:"sampling"`
// 	// Encoding sets the logger's encoding. Valid values are "json" and
// 	// "console", as well as any third-party encodings registered via
// 	// RegisterEncoder.
// 	Encoding string `json:"encoding" yaml:"encoding"`
// 	// EncoderConfig sets options for the chosen encoder. See
// 	// zapcore.EncoderConfig for details.
// 	EncoderConfig zapcore.EncoderConfig `json:"encoderConfig" yaml:"encoderConfig"`
// 	// OutputPaths is a list of URLs or file paths to write logging output to.
// 	// See Open for details.
// 	OutputPaths []string `json:"outputPaths" yaml:"outputPaths"`
// 	// ErrorOutputPaths is a list of URLs to write internal logger errors to.
// 	// The default is standard error.
// 	//
// 	// Note that this setting only affects internal errors; for sample code that
// 	// sends error-level logs to a different location from info- and debug-level
// 	// logs, see the package-level AdvancedConfiguration example.
// 	ErrorOutputPaths []string `json:"errorOutputPaths" yaml:"errorOutputPaths"`
// 	// InitialFields is a collection of fields to add to the root logger.
// 	InitialFields map[string]interface{} `json:"initialFields" yaml:"initialFields"`
// }

func GetLogConfig() zap.Config {

	//Open our jsonFile
	jsonFile, err := os.Open("./Back/server/logs/config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	var cfg zap.Config
	// Read the file contents
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatalf("open json failed with ioutil: %v", err)
	}

	if err := json.Unmarshal(byteValue, &cfg); err != nil {
		panic(err)
	}

	return cfg
}

func GetLog(cfg zap.Config) *zap.Logger {

	logger_build,err := cfg.Build()

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	logger := zap.Must(logger_build,err)

	defer logger.Sync()

	logger.Info("logger construction succeeded")

	return logger
}
