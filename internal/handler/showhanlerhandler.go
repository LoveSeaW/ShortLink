package handler

import (
	"net/http"

	"ShortLink/internal/logic"
	"ShortLink/internal/svc"
	"ShortLink/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShowHanlerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewShowHanlerLogic(r.Context(), svcCtx)
		resp, err := l.ShowHanler(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
