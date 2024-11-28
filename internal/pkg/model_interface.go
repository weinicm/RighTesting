package main

// ModelAdapter 定义模型适配器接口
type ModelAdapter interface {
	CallModel(input string) (string, error)
}

// ModelRequest 定义请求结构体
type ModelRequest struct {
	Input string `json:"input"`
}

// ModelResponse 定义响应结构体
type ModelResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}
