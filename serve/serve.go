package serve

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Run run cobra subcommand
func Run(cmd *cobra.Command, args []string) {

	port := viper.GetInt("port")

	accessKeyID := viper.GetString("ACCESS_KEY_ID")
	if len(accessKeyID) == 0 {
		logrus.Errorf("can not found GA_SMS_ACCESS_KEY_ID")
		return
	}

	accessKeySecret := viper.GetString("ACCESS_KEY_SECRET")
	if len(accessKeySecret) == 0 {
		logrus.Errorf("can not found GA_SMS_ACCESS_KEY_SECRET")
		return
	}

	api, err := NewAPI(accessKeyID, accessKeySecret)
	if err != nil {
		logrus.Errorf("create api failed: %s", err)
		return
	}

	getServeEngine(api).Run(fmt.Sprintf(":%d", port))
}

func getServeEngine(api *API) *gin.Engine {
	r := gin.Default()

	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.StaticFile("/", "api.yml")
	r.GET("/ping", api.Ping)
	r.GET("/_health", api.Health)

	// TODO: drop
	gwAuthorized := r.Group("/sms")
	gwAuthorized.Use(gwAuthMiddleware())
	{
		gwAuthorized.GET("/template/:template_code", api.QuerySmsTemplate)
		gwAuthorized.POST("/send", api.SendSms)
	}

	return r
}
