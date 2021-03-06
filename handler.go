package main

import (
	"github.com/nclgh/lakawei_member/handler"
	"github.com/nclgh/lakawei_scaffold/rpc/member"
)

type ServiceMember struct {
}

func (server *ServiceMember) CheckUserIdentity(req member.CheckUserIdentityRequest, res *member.CheckUserIdentityResponse) error {
	resp := handler.CheckUserIdentity(&req)
	*res = *resp
	return nil
}

func (server *ServiceMember) AddDepartment(req member.AddDepartmentRequest, res *member.AddDepartmentResponse) error {
	resp := handler.AddDepartment(&req)
	*res = *resp
	return nil
}

func (server *ServiceMember) DeleteDepartment(req member.DeleteDepartmentRequest, res *member.DeleteDepartmentResponse) error {
	resp := handler.DeleteDepartment(&req)
	*res = *resp
	return nil
}

func (server *ServiceMember) GetDepartmentByCode(req member.GetDepartmentByCodeRequest, res *member.GetDepartmentByCodeResponse) error {
	resp := handler.GetDepartmentByCode(&req)
	*res = *resp
	return nil
}

func (server *ServiceMember) QueryDepartment(req member.QueryDepartmentRequest, res *member.QueryDepartmentResponse) error {
	resp := handler.QueryDepartment(&req)
	*res = *resp
	return nil
}

func (server *ServiceMember) AddMember(req member.AddMemberRequest, res *member.AddMemberResponse) error {
	resp := handler.AddMember(&req)
	*res = *resp
	return nil
}

func (server *ServiceMember) DeleteMember(req member.DeleteMemberRequest, res *member.DeleteMemberResponse) error {
	resp := handler.DeleteMember(&req)
	*res = *resp
	return nil
}

func (server *ServiceMember) GetMemberByCode(req member.GetMemberByCodeRequest, res *member.GetMemberByCodeResponse) error {
	resp := handler.GetMemberByCode(&req)
	*res = *resp
	return nil
}

func (server *ServiceMember) QueryMember(req member.QueryMemberRequest, res *member.QueryMemberResponse) error {
	resp := handler.QueryMember(&req)
	*res = *resp
	return nil
}
