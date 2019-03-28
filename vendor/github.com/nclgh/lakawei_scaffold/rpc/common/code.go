package common

type RspCode int64

const (
	CodeSuccess RspCode = 0
	CodeFailed  RspCode = 500
	CodeReqErr  RspCode = 1
)
