package repository

import (
	"github.com/MatchlockIsDAST/MatchlockCore/src/domain/model/request"
	"github.com/MatchlockIsDAST/MatchlockCore/src/domain/model/response"
)

type Response interface {
	Save(req *response.Response) (savesuccess bool, err error)
	Find(httpid string) (resp *response.Response)
}

type Request interface {
	Save(req *request.Request) (savesuccess bool, err error)
	Find(httpid string) (reqs *request.Request)
}

type HTTPPair interface {
	Lists(uuid string) (httpids []string, err error)
}
