package main

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/ooclab/ga.sms.aliyun/cmd"
)

const programVersion = "0.0.1"

var (
	buildstamp = ""
	githash    = ""
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	})

	viper.Set("program.version", programVersion)
	viper.Set("program.buildstamp", buildstamp)
	viper.Set("program.githash", githash)
}

func main() {
	cmd.Execute()
}
