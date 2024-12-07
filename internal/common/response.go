package common

import "github.com/gin-gonic/gin"


type BaseResponse struct {}

type response struct {
	Errno   int         `json:"errno"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	TraceID string      `json:"trace_id"`
}

func (base *BaseResponse) Response(c *gin.Context, err error, data interface{}) {
	if err != nil {
		base.Fail(c, err)		
	} else {
		base.Success(c, data)
	}
}

func (base *BaseResponse) Success(c *gin.Context, data interface{}) {

	// return base.Response(c, nil, data)
}

func (base *BaseResponse) Fail(c *gin.Context, err error) {
	// return base.Response(c, err, nil)
}