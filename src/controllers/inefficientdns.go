package controllers

import (
	"github.com/kukkar/common-golang/globalconst"

	"github.com/gin-gonic/gin"
	"github.com/kukkar/common-golang/pkg/responsewriter"
	"github.com/kukkar/common-golang/pkg/utils"
	"github.com/kukkar/common-golang/pkg/utils/rError"
	appConf "github.com/kukkar/mta-hosting-optimizer/conf"
	bl "github.com/kukkar/mta-hosting-optimizer/src/mta-hosting-optimizer"
)

// GetInefficientHosts get unused Host of ips
// @Summary GetInefficientHosts get unused hosts on threshold
// @Produce json
// @Param threshhold query string false "threshhold"
// @Success 200 {object} ResUnusedIpHost
// @Router /v1/listunusedhost [get]
func GetInefficientHosts(c *gin.Context) {

	var rc utils.RequestContext
	if requestContext, ok := c.Get(globalconst.RequestContext); ok {
		rc = requestContext.(utils.RequestContext)
	}

	instance, err := bl.GetInterface(
		c.Request.Context(),
		bl.Config{
			StorageAdapter: "mysql",
			RC:             rc,
		})

	appConfig, err := appConf.GetAppConfig()
	if err != nil {
		err = rError.MiscError(c, err, "unable to load config")
		responsewriter.BuildResponse(c, "", err)
		return
	}

	threshhold := appConfig.ActiveIPCountThreshold

	if err != nil {
		err = rError.MiscError(c, err, "Unable to get instance")
		responsewriter.BuildResponse(c, "", err)
		return
	}
	data, err := instance.GetInefficientHosts(threshhold)
	res := make([]ResUnusedIpHost, 0)
	for _, eachHost := range data {
		res = append(res, ResUnusedIpHost{
			HostName: eachHost.HostName,
		})
	}

	responsewriter.BuildResponseWithBool(c, res, nil)
	return
}
