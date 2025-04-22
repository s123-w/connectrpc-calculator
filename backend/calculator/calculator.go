package calculator

import (
	"context"
	"errors"
	"math"

	"github.com/bufbuild/connect-go"
)

// CalculatorServer 实现了我们在Proto中定义的CalculatorService
type CalculatorServer struct{}

// Calculate 实现了四则运算
func (s *CalculatorServer) Calculate(
	ctx context.Context,
	req *connect.Request[CalculateRequest],
) (*connect.Response[CalculateResponse], error) {
	input := req.Msg
	var result float64

	switch input.Operation {
	case Operation_ADD:
		result = input.FirstNumber + input.SecondNumber
	case Operation_SUBTRACT:
		result = input.FirstNumber - input.SecondNumber
	case Operation_MULTIPLY:
		result = input.FirstNumber * input.SecondNumber
	case Operation_DIVIDE:
		if input.SecondNumber == 0 {
			return connect.NewResponse(&CalculateResponse{
				Result: 0,
				Error:  "除数不能为零",
			}), nil
		}
		result = input.FirstNumber / input.SecondNumber
	default:
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("不支持的运算符"))
	}

	// 处理特殊情况：无限和NaN
	if math.IsInf(result, 0) || math.IsNaN(result) {
		return connect.NewResponse(&CalculateResponse{
			Result: 0,
			Error:  "计算结果为无限大或非数字",
		}), nil
	}

	return connect.NewResponse(&CalculateResponse{
		Result: result,
	}), nil
} 