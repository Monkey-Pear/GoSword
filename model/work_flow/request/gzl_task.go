package request

import "time"

type Task struct {
	State uint8
}
type Inspect struct {
	TaskId uint  `json:"taskId"`
	State  uint8 `json:"state"`
}

type Dynamic struct {
	CreatedAt          time.Time `json:"createdAt"`          // 创建时间
	InspectAt          time.Time `json:"inspectAt"`          // 审批时间
	CreatedAtFormatStr string    `json:"createdAtFormatStr"` //格式化后的创建时间 --> yyyy-mm-dd hh:mm
	InspectAtFormatStr string    `json:"inspectAtFormatStr"` //格式化后的审批时间 --> yyyy-mm-dd hh:mm
	ConsumeTime        int64     `json:"consumeTime"`        // 审批耗时
	Applicant          string    `json:"applicant"`          // 申请人
	CheckState         uint8     `json:"checkState"`         // 审批状态
	Remarks            string    `json:"remarks"`            // 备注
}
type Function struct {
	CreatedAt          time.Time  `json:"createdAt"`          // 创建时间
	Applicant          string     `json:"applicant"`          // 申请人
	Inspector           uint      `json:"inspector" gorm:"not null;comment:审批人"`
	AppName            string     `json:"appName"`           //应用名称
	CheckState         uint8      `json:"checkState" gorm:"not null;comment:审批状态(待审批1默认、审批通过2、审批拒绝3、或签已审核4)"`               //审批状态(待审批1默认、审批通过2、审批拒绝3、或签已审核4)
}