package work_flow

import (
	"github.com/gin-gonic/gin"
	WorkFlowReq "project/model/work_flow"
)

type RecordApi struct {
}

// Submit
// @Tags record
// @Summary 提交表单
// @Produce  application/json
// @Param data body WorkFlowReq.GzlRecord true "string"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"null"}"
// @Router /record/submit [post]
func (f *RecordApi) Submit(c *gin.Context) {
	var _ WorkFlowReq.GzlRecord
}

// Data
// @Tags record
// @Summary 返回之前填写过的表单数据
// @Produce  application/json
// @Param data body int true "string"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"null"}"
// @Router /record/data [get]
func (f *RecordApi) Data(c *gin.Context) {
	var _ WorkFlowReq.GzlAppUser
}
