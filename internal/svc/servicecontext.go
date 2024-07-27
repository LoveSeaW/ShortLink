package svc

import (
	"ShortLink/internal/config"
	"ShortLink/model"
	"ShortLink/sequence"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	ShortUrlModel     model.ShortUrlMapModel
	Sequence          sequence.Sequence
	ShortUrlBlackList map[string]struct{}
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	m := make(map[string]struct{}, len(c.ShortUrlBlackList))
	for _, v := range c.ShortUrlBlackList {
		m[v] = struct{}{}
	}
	return &ServiceContext{
		Config:            c,
		ShortUrlModel:     model.NewShortUrlMapModel(conn),
		Sequence:          sequence.NewMysql(c.Sequence.DSN),
		ShortUrlBlackList: m,
	}
}
