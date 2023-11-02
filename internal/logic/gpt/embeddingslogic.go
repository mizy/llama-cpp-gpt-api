package gpt

import (
	"context"
	"net/http"

	"github.com/mizy/llama-cpp-gpt-api/internal/svc"
	"github.com/mizy/llama-cpp-gpt-api/internal/types"
	"github.com/mizy/llama-cpp-gpt-api/pkg/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmbeddingsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      *http.ResponseWriter
	r      *http.Request
}

func NewEmbeddingsLogic(ctx context.Context, svcCtx *svc.ServiceContext, w *http.ResponseWriter, r *http.Request) *EmbeddingsLogic {
	return &EmbeddingsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
		r:      r,
	}
}

func (l *EmbeddingsLogic) Embeddings(req *types.ReqEmbeddings) (resp *types.ResEmbeddings, err error) {

	embeds, err := model.LlamaInstance.Embeddings(req.Input)

	return &types.ResEmbeddings{
		Data: []types.Embedding{
			{
				Embedding: embeds,
				Index:     0,
			},
		},
	}, err
}
