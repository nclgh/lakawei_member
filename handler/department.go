package handler

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_member/model"
	"github.com/nclgh/lakawei_scaffold/utils"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
	"github.com/nclgh/lakawei_scaffold/rpc/member"
)

func tranDepartment(md *model.Department) *member.Department {
	return &member.Department{
		Code: md.Code,
		Name: md.Name,
	}
}

func batchTranDepartment(mds []*model.Department) map[string]*member.Department {
	ret := make(map[string]*member.Department)
	for _, v := range mds {
		ret[v.Code] = tranDepartment(v)
	}
	return ret
}

func rTranDepartment(d *member.Department) *model.Department {
	return &model.Department{
		Code: d.Code,
		Name: d.Name,
	}
}

func AddDepartment(req *member.AddDepartmentRequest) (rsp *member.AddDepartmentResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("AddDepartment panic: %v, stack: %v", err, stacks)
		rsp = getAddDepartmentResponse(common.CodeFailed, "panic")
	})
	err := model.InsertDepartment(model.GetLakaweiDb(), req.Code, req.Name)
	if err != nil {
		logrus.Errorf("insert department into mysql failed. code: %v, err: %v", req.Code, err)
		return getAddDepartmentResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	return getAddDepartmentResponse(common.CodeSuccess, "")
}

func getAddDepartmentResponse(code common.RspCode, msg string) *member.AddDepartmentResponse {
	rsp := &member.AddDepartmentResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}

func DeleteDepartment(req *member.DeleteDepartmentRequest) (rsp *member.DeleteDepartmentResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("DeleteDepartment panic: %v, stack: %v", err, stacks)
		rsp = getDeleteDepartmentResponse(common.CodeFailed, "panic")
	})
	err := model.DeleteDepartment(model.GetLakaweiDb(), req.Code)
	if err != nil {
		logrus.Errorf("delete department from mysql failed. err: %v", err)
		return getDeleteDepartmentResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	return getDeleteDepartmentResponse(common.CodeSuccess, "")
}

func getDeleteDepartmentResponse(code common.RspCode, msg string) *member.DeleteDepartmentResponse {
	rsp := &member.DeleteDepartmentResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}

func GetDepartmentByCode(req *member.GetDepartmentByCodeRequest) (rsp *member.GetDepartmentByCodeResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("GetDepartmentByCode panic: %v, stack: %v", err, stacks)
		rsp = getGetDepartmentByCodeResponse(common.CodeFailed, "panic")
	})
	ret, err := model.GetDepartmentByCode(model.GetLakaweiDb(), req.Codes)
	if err != nil {
		logrus.Errorf("select department from mysql failed. err: %v", err)
		return getGetDepartmentByCodeResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getGetDepartmentByCodeResponse(common.CodeSuccess, "")
	rsp.Departments = batchTranDepartment(ret)
	return rsp
}

func getGetDepartmentByCodeResponse(code common.RspCode, msg string) *member.GetDepartmentByCodeResponse {
	rsp := &member.GetDepartmentByCodeResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}

func QueryDepartment(req *member.QueryDepartmentRequest) (rsp *member.QueryDepartmentResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("QueryDepartment panic: %v, stack: %v", err, stacks)
		rsp = getQueryDepartmentResponse(common.CodeFailed, "panic")
	})
	ret, cnt, err := model.QueryDepartment(model.GetLakaweiDb(), rTranDepartment(req.Department), req.Page, req.PageSize)
	if err != nil {
		logrus.Errorf("filter department from mysql failed. err: %v", err)
		return getQueryDepartmentResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getQueryDepartmentResponse(common.CodeSuccess, "")
	rsp.Departments = batchTranDepartment(ret)
	rsp.TotalCount = cnt
	return rsp
}

func getQueryDepartmentResponse(code common.RspCode, msg string) *member.QueryDepartmentResponse {
	rsp := &member.QueryDepartmentResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}
