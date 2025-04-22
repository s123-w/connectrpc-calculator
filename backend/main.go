package main

import (
	"calculator/calculator"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	calculatorServer := &calculator.CalculatorServer{}
	path, handler := calculator.NewCalculatorServiceHandler(calculatorServer)

	mux := http.NewServeMux()
	mux.Handle(path, handler)

	// 添加CORS中间件
	corsHandler := withCORS(mux)

	// 配置服务器
	server := &http.Server{
		Addr:              ":8080",
		Handler:           h2c.NewHandler(corsHandler, &http2.Server{}),
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Println("计算器服务启动在: http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}

// CORS中间件
func withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置必要的CORS头
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Connect-Protocol-Version")
		w.Header().Set("Access-Control-Max-Age", "86400") // 24小时

		// 处理预检请求
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// 处理常规请求
		h.ServeHTTP(w, r)
	})
} 