package logic

import (
	"ShortLink/internal/svc"
	"ShortLink/internal/types"
	"ShortLink/model"
	"ShortLink/pkg/base62"
	"ShortLink/pkg/connect"
	"ShortLink/pkg/md5"
	"ShortLink/pkg/urltool"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 将长连接转换为短链接
func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// 参数校验
	if ok := connect.Get(req.LongUrl); !ok {
		return nil, errors.New("无效链接：链接无法访问")
	}
	md5Value := md5.Sum([]byte(req.LongUrl))
	// 查询长连接是否已被转换
	u, err := l.svcCtx.ShortUrlModel.FindOneByMd5(l.ctx, sql.NullString{String: md5Value, Valid: true})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, fmt.Errorf("该链接已被转换为%s", u.Surl.String)
		}
		logx.Errorw("ShortUrlModel.FindOneByMd5 failed", logx.LogField{
			Key:   "err",
			Value: err.Error(),
		})
		return nil, err
	}
	basePath, err := urltool.GetBasePath(req.LongUrl)
	if err != nil {
		logx.Errorw("urltool.GetBasePath failed", logx.LogField{
			Key:   "err",
			Value: err.Error(),
		})
		return nil, err
	}
	_, err = l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{String: basePath, Valid: true})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, fmt.Errorf("该链接已经是短链了")
		}
		logx.Errorw("ShortUlrModel.FindOneBySurl failed", logx.LogField{
			Key:   "err",
			Value: err.Error(),
		})
	}

	var short string
	for {
		// 取号器，基于Mysql实现的发号器
		// 每一个转链请求， 使用 replace into 语句 插入一条数据，取出主键id作为号码
		seq, err := l.svcCtx.Sequence.Next()
		if err != nil {
			logx.Errorw("l.svcCtx.Sequence.Next() failed", logx.LogField{
				Key:   "err",
				Value: err.Error(),
			})
			return nil, err
		}
		fmt.Printf("id: %d\n", seq)

		short = base62.Int62ToString(seq)
		if _, ok := l.svcCtx.ShortUrlBlackList[short]; !ok {
			break
		}
		//fmt.Printf("short: %s\n", short)
	}

	_, err = l.svcCtx.ShortUrlModel.Insert(
		l.ctx,
		&model.ShortUrlMap{
			Lurl: sql.NullString{req.LongUrl, true},
			Md5:  sql.NullString{md5Value, true},
			Surl: sql.NullString{short, true},
		},
	)
	if err != nil {
		logx.Errorw("ShorUrlModel.Insert() failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}

	// 返回短域名加上短链接
	shortUrl := l.svcCtx.Config.ShortDomain + "/" + short
	return &types.ConvertResponse{ShortUrl: shortUrl}, nil
}
