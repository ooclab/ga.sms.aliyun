package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ooclab/ga.sms.aliyun/serve"
)

func init() {
	serveCmd.Flags().Int("port", 3000, "Port to run server on")
	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))

	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve ARGS",
	Short: "start supervisor serve",
	Run:   serve.Run,
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindEnv("ACCESS_KEY_ID")
		viper.BindEnv("ACCESS_KEY_SECRET")
	},
}
