package logic

import (
	"context"

	"ShortLink/internal/svc"
	"ShortLink/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowHanlerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowHanlerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowHanlerLogic {
	return &ShowHanlerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowHanlerLogic) ShowHanler(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
