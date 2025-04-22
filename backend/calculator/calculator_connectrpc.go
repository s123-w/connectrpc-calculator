package calculator

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/bufbuild/connect-go"
)

// 实现 CalculatorServiceHandler 接口
const CalculatorServiceName = "calculator.CalculatorService"

var _ http.Handler = &CalculatorServiceHandler{}

// CalculatorServiceHandler 是 CalculatorService 的处理器
type CalculatorServiceHandler struct {
	srv CalculatorServiceInterface
	mux *http.ServeMux
}

// CalculatorServiceInterface 包含所有 CalculatorService 的方法
type CalculatorServiceInterface interface {
	Calculate(context.Context, *connect.Request[CalculateRequest]) (*connect.Response[CalculateResponse], error)
}

// NewCalculatorServiceHandler 创建新的处理器
func NewCalculatorServiceHandler(srv CalculatorServiceInterface) (string, http.Handler) {
	path := "/" + CalculatorServiceName + "/"
	handler := &CalculatorServiceHandler{
		srv: srv,
		mux: http.NewServeMux(),
	}
	handler.mux.Handle(path+"Calculate", http.HandlerFunc(handler.HandleCalculate))
	return path, handler
}

// ServeHTTP 实现 http.Handler 接口
func (h *CalculatorServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// HandleCalculate 处理计算请求
func (h *CalculatorServiceHandler) HandleCalculate(w http.ResponseWriter, r *http.Request) {
	var req CalculateRequest
	var err error

	// 解析JSON请求
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if err = json.Unmarshal(body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 调用实际处理函数
	connectReq := connect.NewRequest(&req)
	connectResp, err := h.srv.Calculate(r.Context(), connectReq)
	if err != nil {
		if connectErr, ok := err.(*connect.Error); ok {
			code := int(connectErr.Code())
			http.Error(w, connectErr.Message(), code)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// 返回JSON响应
	w.Header().Set("Content-Type", "application/json")
	respData, err := json.Marshal(connectResp.Msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	_, err = w.Write(respData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
} 