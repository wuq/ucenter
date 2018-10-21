package handler



import (
	"context"

	"github.com/micro/go-log"

	bind "tenno.ucenter/proto/bind"
//	c "tenno.ucenter/common"
//	"tenno.ucenter/model"
)

type Bind struct{}

//绑定
func (b *Bind) Bind(ctx context.Context, req *bind.BindRequest, rsp *bind.Response) error {
	log.Log("Received Bind.Bind request")

	return nil
}

//解绑
func (b *Bind) UnBind(ctx context.Context, req *bind.UnbindRequest, rsp *bind.Response) error {
	log.Log("Received Bind.UnBind request")

	return nil
}


//查询绑定列表
func (b *Bind) GetBindList(ctx context.Context, req *bind.GetBindListRequest, rsp *bind.Response) error {
	log.Log("Received Bind.GetBindList request")

	return nil
}
