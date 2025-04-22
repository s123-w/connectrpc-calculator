package calculator

// Operation 操作类型枚举
type Operation int32

const (
	Operation_ADD      Operation = 0
	Operation_SUBTRACT Operation = 1
	Operation_MULTIPLY Operation = 2
	Operation_DIVIDE   Operation = 3
)

// CalculateRequest 计算请求
type CalculateRequest struct {
	FirstNumber  float64   `json:"firstNumber"`
	SecondNumber float64   `json:"secondNumber"`
	Operation    Operation `json:"operation"`
}

// CalculateResponse 计算响应
type CalculateResponse struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
} 