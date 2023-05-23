package lazada

import (
	"context"
)

type GetMessageResponse struct {
	BaseResponse

	Data GetMessageResponseData `json:"data"`
}

type GetMessageResponseData struct {
	LastMessageID string             `json:"last_message_id"`
	MessageList   []MessagesListData `json:"message_list"`
	NextStartTime string             `json:"next_start_time"`
	HasMore       string             `json:"has_more"`
}

type MessagesListData struct {
	FromAccountType string `json:"from_account_type"`
	ProcessMsg      string `json:"process_msg"`
	SessionID       string `json:"session_id"`
	MessageID       string `json:"message_id"`
	Type            string `json:"type"`
	Content         string `json:"content"`
	ToAccountID     string `json:"to_account_id"`
	SendTime        string `json:"send_time"`
	AutoReply       string `json:"auto_reply"`
	ToAccountType   string `json:"to_account_type"`
	SiteID          string `json:"site_id"`
	TemplateID      string `json:"template_id"`
	FromAccountID   string `json:"from_account_id"`
	Status          string `json:"status"`
}

type GetMessageQueryParams struct {
	SessionID     string `url:"session_id"`
	StartTime     string `url:"start_time"`
	PageSize      int    `url:"page_size"`
	LastMessageID string `url:"last_message_id"`
}

type ChatService interface {
	GetMessageList(ctx context.Context, query GetMessageQueryParams) (GetMessageResponse, error)
}

type ChatServiceOp struct {
	sdk *LazadaClient
}

// func (lc *ChatServiceOp) GetMessageList(ctx context.Context, query GetMessageQueryParams) (*GetMessageResponse, error) {
// 	resp, err := lc.sdk.Client.R().
// 		SetResult(&Token{}).
// 		Get(ApiNames["RefreshToken"])

// 	if err != nil {
// 		return nil, fmt.Errorf("error making request: %w", err)
// 	}

// 	if resp.IsError() {
// 		return nil, fmt.Errorf("API request error: %v", resp.Status())
// 	}

// 	tokenResponse := &Token{}
// 	err = json.Unmarshal(resp.Body(), tokenResponse)
// 	if err != nil {
// 		return nil, fmt.Errorf("error unmarshalling response: %w", err)
// 	}

// 	return tokenResponse, nil
// }
