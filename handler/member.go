package handler

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_member/model"
	"github.com/nclgh/lakawei_scaffold/utils"
	"github.com/nclgh/lakawei_scaffold/rpc/member"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

func tranMember(v *model.Member) *member.Member {
	return &member.Member{
		Code:           v.Code,
		Name:           v.Name,
		DepartmentCode: v.DepartmentCode,
	}
}

func batchTranMember(vs []*model.Member) map[string]*member.Member {
	ret := make(map[string]*member.Member)
	for _, v := range vs {
		ret[v.Code] = tranMember(v)
	}
	return ret
}

func rTranMember(v *member.Member) *model.Member {
	return &model.Member{
		Code:           v.Code,
		Name:           v.Name,
		DepartmentCode: v.DepartmentCode,
	}
}

func AddMember(req *member.AddMemberRequest) (rsp *member.AddMemberResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("AddMemberRequest panic: %v, stack: %v", err, stacks)
		rsp = getAddMemberRequestResponse(common.CodeFailed, "panic")
	})
	err := model.InsertMember(model.GetLakaweiDb(), req.Code, req.Name, req.DepartmentCode)
	if err != nil {
		logrus.Errorf("insert member into mysql failed. code: %v, err: %v", req.Code, err)
		return getAddMemberRequestResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getAddMemberRequestResponse(common.CodeSuccess, "")
	return rsp
}

func getAddMemberRequestResponse(code common.RspCode, msg string) *member.AddMemberResponse {
	rsp := &member.AddMemberResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}

func DeleteMember(req *member.DeleteMemberRequest) (rsp *member.DeleteMemberResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("DeleteMemberRequest panic: %v, stack: %v", err, stacks)
		rsp = getDeleteMemberRequestResponse(common.CodeFailed, "panic")
	})
	err := model.DeleteMember(model.GetLakaweiDb(), req.Code)
	if err != nil {
		logrus.Errorf("delete member from mysql failed. err: %v", err)
		return getDeleteMemberRequestResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getDeleteMemberRequestResponse(common.CodeSuccess, "")
	return rsp
}

func getDeleteMemberRequestResponse(code common.RspCode, msg string) *member.DeleteMemberResponse {
	rsp := &member.DeleteMemberResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}

func GetMemberByCode(req *member.GetMemberByCodeRequest) (rsp *member.GetMemberByCodeResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("GetMemberByCodeRequest panic: %v, stack: %v", err, stacks)
		rsp = getGetMemberByCodeRequestResponse(common.CodeFailed, "panic")
	})
	ret, err := model.GetMemberByCode(model.GetLakaweiDb(), req.Codes)
	if err != nil {
		logrus.Errorf("select member from mysql failed. err: %v", err)
		return getGetMemberByCodeRequestResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getGetMemberByCodeRequestResponse(common.CodeSuccess, "")
	rsp.Members = batchTranMember(ret)
	return rsp
}

func getGetMemberByCodeRequestResponse(code common.RspCode, msg string) *member.GetMemberByCodeResponse {
	rsp := &member.GetMemberByCodeResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}

func QueryMember(req *member.QueryMemberRequest) (rsp *member.QueryMemberResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("QueryMemberRequest panic: %v, stack: %v", err, stacks)
		rsp = getQueryMemberResponse(common.CodeFailed, "panic")
	})
	ret, cnt, err := model.QueryMember(model.GetLakaweiDb(), rTranMember(req.Member), req.Page, req.PageSize)
	if err != nil {
		logrus.Errorf("filter member from mysql failed. err: %v", err)
		return getQueryMemberResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getQueryMemberResponse(common.CodeSuccess, "")
	rsp.Members = batchTranMember(ret)
	rsp.TotalCount = cnt
	return rsp
}

func getQueryMemberResponse(code common.RspCode, msg string) *member.QueryMemberResponse {
	rsp := &member.QueryMemberResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}
