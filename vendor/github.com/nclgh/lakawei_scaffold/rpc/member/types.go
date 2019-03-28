package member

import (
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

type CheckUserIdentityRequest struct {
	UserName string
	Password string
}

type CheckUserIdentityResponse struct {
	UserId int64
	Code   common.RspCode
	Msg    string
}

// Department
type Department struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type AddDepartmentRequest struct {
	Code string
	Name string
}

type AddDepartmentResponse struct {
	Code common.RspCode
	Msg  string
}

type DeleteDepartmentRequest struct {
	Code string
}

type DeleteDepartmentResponse struct {
	Code common.RspCode
	Msg  string
}

type GetDepartmentByCodeRequest struct {
	Codes []string
}

type GetDepartmentByCodeResponse struct {
	Departments map[string]*Department
	Code        common.RspCode
	Msg         string
}

type QueryDepartmentRequest struct {
	Department *Department
	Page       int64
	PageSize   int64
}

type QueryDepartmentResponse struct {
	Departments map[string]*Department
	TotalCount  int64
	Code        common.RspCode
	Msg         string
}

// Member
type Member struct {
	Code           string `json:"code"`
	Name           string `json:"name"`
	DepartmentCode string `json:"department_code"`
}

type AddMemberRequest struct {
	Code           string
	Name           string
	DepartmentCode string
}

type AddMemberResponse struct {
	Code common.RspCode
	Msg  string
}

type DeleteMemberRequest struct {
	Code string
}

type DeleteMemberResponse struct {
	Code common.RspCode
	Msg  string
}

type GetMemberByCodeRequest struct {
	Codes []string
}

type GetMemberByCodeResponse struct {
	Members map[string]*Member
	Code    common.RspCode
	Msg     string
}

type QueryMemberRequest struct {
	Member   *Member
	Page     int64
	PageSize int64
}

type QueryMemberResponse struct {
	Members    map[string]*Member
	TotalCount int64
	Code       common.RspCode
	Msg        string
}
