package work_flow

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/global"
	"project/model/common/response"
	modelWF "project/model/work_flow"
	WorkFlowReq "project/model/work_flow/request"
	"project/utils"
)

type AppApi struct {
}

// Empty
// @Tags App
// @Summary 返回空表单
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/json
// @Param appId query WorkFlowReq.EmptyApp true "应用id"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"null"}"
// @Router /app/empty [get]
func (f *AppApi) Empty(c *gin.Context) {
	var app WorkFlowReq.EmptyApp
	_ = c.ShouldBindQuery(&app)
	if err := utils.Verify(app, utils.EmptyAppVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := appService.GetAppEmpty(app.AppId)
	if err != nil {
		global.GSD_LOG.Error("获取空应用表单失败", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("该应用不存在", c)
	} else {
		global.GSD_LOG.Info("获取空应用表单成功", zap.Any("GetAppEmpty Success", data), utils.GetRequestID(c))
		response.OkWithData(data, c)
	}
}

// Create
// @Tags App
// @Summary 创建表单
// @Produce  application/json
// @Param data body uint true "创建人"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"创建表单成功"}"
// @Router /app/create [post]
func (f *AppApi) Create(c *gin.Context) {

}

// AddApp
// @Tags App
// @Summary 创建应用
// @Produce  application/json
// @Param data body uint true "创建人"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"创建表单成功"}"
// @Router /app/ [post]
func (f *AppApi) AddApp(c *gin.Context) {
	var addApp WorkFlowReq.AddApp
	_ = c.ShouldBindJSON(&addApp)
	if err := utils.Verify(addApp, utils.AddApp); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := appService.AddApp(modelWF.GzlApp{GSD_MODEL: global.GSD_MODEL{CreateBy: utils.GetUserID(c)}, Name: addApp.Name, Icon: addApp.Icon})
	if err != nil {
		global.GSD_LOG.Error("添加应用失败", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("操作失败", c)
	} else {
		response.OkWithMessage("添加应用成功", c)
	}
}
