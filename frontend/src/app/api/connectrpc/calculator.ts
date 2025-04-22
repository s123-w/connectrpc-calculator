// 与后端通信的协议定义

// 操作类型枚举
export enum Operation {
  ADD = 0,
  SUBTRACT = 1,
  MULTIPLY = 2,
  DIVIDE = 3,
}

// 计算请求接口
export interface CalculateRequest {
  firstNumber: number;
  secondNumber: number;
  operation: Operation;
}

// 计算响应接口
export interface CalculateResponse {
  result: number;
  error?: string;
}

// 创建客户端函数
export async function calculateRequest(request: CalculateRequest): Promise<CalculateResponse> {
  const response = await fetch('http://localhost:8080/calculator.CalculatorService/Calculate', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Connect-Protocol-Version': '1',
    },
    body: JSON.stringify(request),
  });

  const data = await response.json();
  
  // ConnectRPC响应通常包含在data对象中
  if (data.error) {
    return {
      result: 0,
      error: data.error.message || '未知错误',
    };
  }

  return data;
} 