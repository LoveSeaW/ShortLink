package handler

import (
	"net/http"

	"ShortLink/internal/logic"
	"ShortLink/internal/svc"
	"ShortLink/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ConverHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ConvertRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewConverLogic(r.Context(), svcCtx)
		resp, err := l.Conver(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
