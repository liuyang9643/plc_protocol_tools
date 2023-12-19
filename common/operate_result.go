package common

import (
	"fmt"
	"strconv"
)

type OperateResult struct {
	ErrorCode int
	Message   string
	IsSuccess bool
	Content   interface{}
	Content1  interface{}
	Content2  interface{}
	Content3  interface{}
	Content4  interface{}
	Content5  interface{}
}

func NewOperateResult(err int, msg string) *OperateResult {
	return &OperateResult{
		ErrorCode: err,
		Message:   msg,
		IsSuccess: false,
	}
}

func (o *OperateResult) String() string {
	return fmt.Sprintf("OperateResult IsSuccess:%v ErrorCode:%d Message:%s", o.IsSuccess, o.ErrorCode, o.Message)
}

func (o *OperateResult) ToString() string {
	return "Language.ErrorCode:" + strconv.Itoa(o.ErrorCode) + "\r\nLanguage.TextDescription:" + o.Message
}

func (o *OperateResult) CopyErrorFromOther(result *OperateResult) {
	if result != nil {
		o.ErrorCode = result.ErrorCode
		o.Message = result.Message
	}
}

func CreateFailedResult(result *OperateResult) *OperateResult {
	failed := NewOperateResult(0, "")
	if result != nil {
		failed.ErrorCode = result.ErrorCode
		failed.Message = result.Message
	}
	return failed
}

func CreateSuccessResult(Content1, Content2, Content3, Content4, Content5 interface{}) *OperateResult {
	success := NewOperateResult(0, "")
	success.IsSuccess = true
	success.Message = "SuccessText"
	success.Content = Content1
	success.Content1 = Content1
	success.Content2 = Content2
	success.Content3 = Content3
	success.Content4 = Content4
	success.Content5 = Content5
	return success
}
