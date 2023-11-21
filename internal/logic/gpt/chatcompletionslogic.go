package gpt

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/go-skynet/go-llama.cpp"
	"github.com/mizy/llama-cpp-gpt-api/internal/svc"
	"github.com/mizy/llama-cpp-gpt-api/internal/types"
	"github.com/mizy/llama-cpp-gpt-api/pkg/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatCompletionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      *http.ResponseWriter
	r      *http.Request
}

func NewChatCompletionsLogic(ctx context.Context, svcCtx *svc.ServiceContext, w *http.ResponseWriter, r *http.Request) *ChatCompletionsLogic {
	return &ChatCompletionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
		r:      r,
	}
}

func (l *ChatCompletionsLogic) ChatCompletions(req *types.ReqChatCompletion) (resp *types.ResChatCompletion, err error) {
	text := model.ConvertMessages2Text(req.Messages)
	w := *l.w
	if req.Stream {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

	}
	log.Print("start predict:\n", text)
	flusher, flusherOk := w.(http.Flusher)
	if !flusherOk {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return nil, nil
	}
	result, err := model.LlamaInstance.Predict(text, func(p *llama.PredictOptions) {
		if req.MaxTokens > 0 {
			p.Tokens = req.MaxTokens
		}
		if req.Temperature > 0 {
			p.Temperature = req.Temperature
		}
		if req.TopP > 0 {
			p.TopP = req.TopP
		}
		if req.Seed > 0 {
			p.Seed = req.Seed
		}
		if req.Stream {
			id := 0
			p.TokenCallback = func(token string) bool {
				data := &types.ResChatCompletion{
					ID: fmt.Sprint(id),
					Choices: []types.Choice{
						{

							Delta: types.Message{
								Content: token,
							},
							Index: 0,
						},
					},
				}
				json, err := json.Marshal(data)
				if err != nil {
					return false
				}
				id++
				fmt.Fprint(w, "data:"+string(json)+"\n\n")
				flusher.Flush()
				return true
			}
		}
		p.Threads = runtime.NumCPU()
		// p.Tokens = 512
	})
	log.Print("end predict", result)
	if !req.Stream {
		if err != nil {
			return nil, err
		}
		return &types.ResChatCompletion{
			ID: "",
			Choices: []types.Choice{
				{
					Message: types.Message{
						Content: result,
					},
					Index: 0,
				},
			},
		}, nil
	} else {
		fmt.Fprint(w, "[DONE]")
		flusher.Flush()
	}
	return nil, nil
}
