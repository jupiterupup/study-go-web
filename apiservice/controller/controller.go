package controller

import (
	"encoding/json"
	"fire-press/api/types"
	"fire-press/apiservice/service"
	"fire-press/apiservice/util/viperhelper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yitter/idgenerator-go/idgen"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var configService service.ConfigService

func InitController() {
	installService()
	startService()
}

func saveConfig(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
	}()
	data, _ := ioutil.ReadAll(c.Request.Body)
	config := types.Config{}
	_ = json.Unmarshal(data, &config)
	config.Id = idgen.NextId()
	config.CreateTime = time.Now()
	config.ModifyTime = time.Now()
	configService.Save(config)
	c.JSON(http.StatusOK, config.Id)
}

func deleteConfig(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
	}()
	data := make(map[string]string)
	_ = c.BindJSON(&data)
	id := data["id"]
	if len(id) == 0 {
		c.JSON(http.StatusInternalServerError, "id不能为空")
		return
	}
	parseId, _ := strconv.ParseInt(id, 10, 64)
	deleteResult := configService.Delete(parseId)
	c.JSON(http.StatusOK, deleteResult)
}

func installService() {
	configService = service.NewConfigService()
}

func startService() {
	router := gin.Default()
	configG := router.Group("/configService")
	{
		configG.POST("/save", saveConfig)
		configG.POST("/delete", deleteConfig)
	}
	_ = router.Run(":" + viperhelper.GetLocalConfIfPresent("server.port"))
	fmt.Println("start service success")
}
