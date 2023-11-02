package gpt

import (
	"net/http"

	"github.com/mizy/llama-cpp-gpt-api/internal/logic/gpt"
	"github.com/mizy/llama-cpp-gpt-api/internal/svc"
	"github.com/mizy/llama-cpp-gpt-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatCompletionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqChatCompletion
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := gpt.NewChatCompletionsLogic(r.Context(), svcCtx, &w, r)
		resp, err := l.ChatCompletions(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else if resp != nil {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
