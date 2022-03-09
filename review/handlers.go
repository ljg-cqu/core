package main

import (
	"github.com/danielgtaylor/huma"
	"net/http"
)

func Create(ctx huma.Context, cReq CreateRequest) {
	_ = cReq
	ctx.WriteModel(http.StatusOK, &ResponseOk{})
}

func Cancel(ctx huma.Context, cReq CancelRequest) {
	_ = cReq
	ctx.WriteModel(http.StatusOK, &ResponseOk{})
}

func At(ctx huma.Context, aReq AtRequest) {
	_ = aReq
	ctx.WriteModel(http.StatusOK, &ResponseOk{})
}

func Comment(ctx huma.Context, cReq CommentRequest) {
	_ = cReq
	ctx.WriteModel(http.StatusOK, &ResponseOk{})
}

func Reply(ctx huma.Context, rReq ReplyRequest) {
	_ = rReq
	ctx.WriteModel(http.StatusOK, &ResponseOk{})
}
