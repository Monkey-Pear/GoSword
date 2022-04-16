package work_flow

import (
	"errors"
	"gorm.io/gorm"
	"project/global"
	"project/model/work_flow"
	modelWF "project/model/work_flow"
	WorkFlowReq "project/model/work_flow/request"
	"time"
)

type TaskService struct {
}

// GetDynamic
// @author: [tanshaokang](https://github.com/worryfreet)
// @function: GetDynamic
// @description: 从mysql中获取流程动态数据
// @param: WorkFlowReq.Record
// @return: data []WorkFlowReq.Dynamic, err error
func (t TaskService) GetDynamic(applicantId, recordId int) (data []WorkFlowReq.Dynamic, err error) {
	db := global.GSD_DB.
		Model(modelWF.GzlTask{}).
		Joins("JOIN sys_users ON sys_users.id = ?", applicantId).
		Joins("JOIN gzl_records ON gzl_records.id = ?", recordId).
		Select("sys_users.username as Applicant", "gzl_tasks.created_at as InspectAt",
			"gzl_records.created_at as CreatedAt", "check_state as CheckState", "remarks as Remarks").
		Where("gzl_tasks.record_id = gzl_records.id")
	if err = db.Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	// 计算耗时, 格式化时间
	for i := 0; i < len(data); i++ {
		born := "2006-01-02 15:04:05"
		beginStr := data[i].CreatedAt.Format(born)
		endStr := data[i].InspectAt.Format(born)
		begin, _ := time.ParseInLocation(born, beginStr, time.Local)
		end, _ := time.ParseInLocation(born, endStr, time.Local)
		data[i].ConsumeTime = end.Unix() - begin.Unix()
		data[i].CreatedAtFormatStr = beginStr[:len(beginStr)-3]
		data[i].InspectAtFormatStr = endStr[:len(endStr)-3]
	}
	return
}

func (t *TaskService) GetScheduleList(InspectorId int) (err error, tasks []work_flow.GzlTask) {
	db := global.GSD_DB.Model(&work_flow.GzlTask{}) //查表GzlTask
	if err = db.Find(&tasks, "inspector = ?", InspectorId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { //如果待办为空，返回空
			return nil, nil
		} else {
			return
		}
	}
	return
}

//func (taskService *TaskService) GetHandleList(InspectorId int) (err error, tasks []work_flow.GzlTask) {
//	db := global.GSD_DB.Model(&work_flow.GzlTask{}) //查表GzlTask
//	if err = db.Find(&tasks, "inspector = ?", InspectorId).Error; err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) { //如果待办为空，返回空
//			return nil, nil
//		} else {
//			return
//		}
//	}
//	return
//}

func (t *TaskService) Inspect(task work_flow.GzlTask) (err error) {
	return global.GSD_DB.Updates(&task).Error
}

// GetReceive
// @author: [tanshaokang](https://github.com/worryfreet)
// @function: GetReceive
// @description: 从mysql中获取我收到的信息列表
// @param: WorkFlowReq.Record
// @return: data []WorkFlowReq.Dynamic, err error
func (t TaskService) GetReceive(userId int) (err error, tasks []modelWF.GzlTask) {
	// 1. 申请人姓名
	// 2. 审批人姓名
	// 3. 审批状态
	// 4. 应用名称
	// 5. 当前节点
	return
}
