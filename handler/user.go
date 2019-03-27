package handler

import (
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_member/model"
	"github.com/nclgh/lakawei_scaffold/utils"
	"github.com/nclgh/lakawei_scaffold/rpc/member"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

func CheckUserIdentity(req *member.CheckUserIdentityRequest) (rsp *member.CheckUserIdentityResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("CheckUserIdentity panic: %v, stack: %v", err, stacks)
		rsp = getCheckUserIdentityResponse(common.CodeFailed, "panic")
	})
	user, err := model.QueryUserByPwd(model.GetLakaweiDb(), req.UserName)
	if err != nil {
		logrus.Errorf("select user from mysql failed. username: %v, err: %v", req.UserName, err)
		return getCheckUserIdentityResponse(common.CodeFailed, "internal error")
	}
	if user == nil {
		return getCheckUserIdentityResponse(common.CodeSuccess, "用户不存在")
	}
	if user.Password != req.Password {
		return getCheckUserIdentityResponse(common.CodeSuccess, "密码错误")
	}
	rsp = getCheckUserIdentityResponse(common.CodeSuccess, "")
	rsp.UserId = user.Id
	return rsp
}

func getCheckUserIdentityResponse(code common.RspCode, msg string) *member.CheckUserIdentityResponse {
	rsp := &member.CheckUserIdentityResponse{
		UserId: 0,
		Code:   code,
		Msg:    msg,
	}
	return rsp
}
