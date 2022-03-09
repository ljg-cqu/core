package main

import "time"

type ResponseOk struct {
	Code int    `json:"code" default:"200" example:"200" doc:"Code represents ok"`
	Msg  string `json:"msg" default:"ok" example:"ok" doc:"Message represents ok"`
}

type ResponseError struct {
	Code int    `json:"code" default:"400" example:"400" doc:"Error code"`
	Msg  string `json:"msg" default:"Bad Request" example:"Bad Request" doc:"Error message"`
}

type ResponseData struct {
	Code int    `json:"code" default:"200" example:"200" doc:"Normal code without error"`
	Msg  string `json:"msg" default:"ok" example:"ok" doc:"Normal message without error"`
	Data string `json:"data" example:"" doc:"Marshaled JSON object in string format"`
}

// ----------

type CreateRequest struct {
	ReviewNum    int    `json:"review_num"`
	TaskNum      int    `json:"task_num" `
	Subject      string `json:"subject"`
	CustomerName string `json:"customer_name"`
	ProjectName  string `json:"project_name"`

	ContractType     string    `json:"contract_type"`
	ContractEffectAt time.Time `json:"contract_effect_at"`
	ContractContent  string    `json:"contract_content, "`
}

type CreateResponse struct {
	ReviewNum string `json:"review_num" example:"review-001" doc:"Newly created contract review number"`
}

type CancelRequest struct {
	Code      int    `json:"code" default:"200" example:"200" doc:"Normal code without error"`
	Msg       string `json:"msg" default:"ok" example:"ok" doc:"Normal message without error"`
	SessionID string `json:"session_id "path:"session-id" example:"session-001" doc:"The session ID of contract review to be canceled"`
}

type CancelResponse struct {
	SessionID string `json:"session_id "example:"session-001" doc:"The session ID of contract review that has been canceled"`
}

type AtRequest struct {
	SessionID string   `json:"session_id"  example:"session-001" doc:"The session ID of contract review"`
	From      string   `json:"from" example:"user-001" doc:"User ID who wants to At someone something."`
	To        []string `json:"to" example:"user-002" example:"user-003" doc:"User ID(s) to AT"`
	Msg       string   `json:"msg" example:"各位，我刚刚发布了新的会审请求，清注意查阅。" doc:"Message to AT."`
}

type CommentRequest struct {
	SessionID string `json:"session_id"  example:"session-001" doc:"The session ID of contract review"`
	AuthorID  string `json:"author" example:"user-001" doc:"The person who comment in a contract review session"`
}

type ReplyRequest struct {
	AuthorID string `json:"author_id" example:"session-001" doc:"The person who "`
	Content  string `json:"content" example:"这条意见可行！" doc:"The comment content" `
}

type ReplyResponse struct {
	CommentID string `json:"comment_id" example:"cmt-001" doc:"The comment ID"`
}
