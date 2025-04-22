'use client';

import { useState, useEffect } from 'react';
import { calculateRequest, Operation } from './api/connectrpc/calculator';

export default function Calculator() {
  const [firstNumber, setFirstNumber] = useState<number | ''>('');
  const [secondNumber, setSecondNumber] = useState<number | ''>('');
  const [operation, setOperation] = useState<Operation>(Operation.ADD);
  const [result, setResult] = useState<number | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [serverStatus, setServerStatus] = useState<'checking' | 'online' | 'offline'>('checking');

  // 检查服务器状态
  useEffect(() => {
    const checkServerStatus = async () => {
      try {
        const controller = new AbortController();
        const timeoutId = setTimeout(() => controller.abort(), 2000);
        
        await fetch('http://localhost:8080/calculator.CalculatorService/', {
          method: 'OPTIONS',
          signal: controller.signal
        });
        
        clearTimeout(timeoutId);
        setServerStatus('online');
      } catch (err) {
        setServerStatus('offline');
      }
    };
    
    checkServerStatus();
    // 每10秒检查一次服务器状态
    const intervalId = setInterval(checkServerStatus, 10000);
    
    return () => clearInterval(intervalId);
  }, []);

  // 操作符映射
  const operationSymbols = {
    [Operation.ADD]: '+',
    [Operation.SUBTRACT]: '-',
    [Operation.MULTIPLY]: '×',
    [Operation.DIVIDE]: '÷',
  };

  // 处理计算
  const handleCalculate = async () => {
    setIsLoading(true);
    setError(null);
    
    try {
      const first = typeof firstNumber === 'string' ? parseFloat(firstNumber) || 0 : firstNumber;
      const second = typeof secondNumber === 'string' ? parseFloat(secondNumber) || 0 : secondNumber;
      
      const response = await calculateRequest({
        firstNumber: first,
        secondNumber: second,
        operation,
      });
      
      if (response.error) {
        setError(response.error);
        setResult(null);
      } else {
        setResult(response.result);
      }
    } catch (err) {
      setError('连接服务器失败: ' + (err instanceof Error ? err.message : String(err)));
      setResult(null);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <main className="min-h-screen flex items-center justify-center p-4 md:p-8">
      <div className="calculator-container">
        <h1 className="calculator-title">ConnectRPC 计算器</h1>
        
        {/* 服务器状态指示器 */}
        <div className="flex justify-center mb-4">
          {serverStatus === 'checking' && (
            <span className="px-3 py-1 text-xs bg-gray-100 text-gray-600 rounded-full">检查服务器状态...</span>
          )}
          {serverStatus === 'online' && (
            <span className="px-3 py-1 text-xs bg-green-100 text-green-600 rounded-full flex items-center">
              <span className="w-2 h-2 bg-green-500 rounded-full mr-1 animate-pulse"></span>
              服务器在线
            </span>
          )}
          {serverStatus === 'offline' && (
            <span className="px-3 py-1 text-xs bg-red-100 text-red-600 rounded-full flex items-center">
              <span className="w-2 h-2 bg-red-500 rounded-full mr-1"></span>
              服务器离线
            </span>
          )}
        </div>
        
        <div className="space-y-4">
          <div className="input-group">
            <label className="input-label">第一个数字</label>
            <input
              type="number"
              value={firstNumber}
              onChange={(e) => setFirstNumber(e.target.value === '' ? '' : Number(e.target.value))}
              className="form-input"
              placeholder="输入第一个数字"
            />
          </div>
          
          <div className="input-group">
            <label className="input-label">操作</label>
            <select
              value={operation}
              onChange={(e) => setOperation(Number(e.target.value) as Operation)}
              className="form-select"
            >
              <option value={Operation.ADD}>加法 (+)</option>
              <option value={Operation.SUBTRACT}>减法 (-)</option>
              <option value={Operation.MULTIPLY}>乘法 (×)</option>
              <option value={Operation.DIVIDE}>除法 (÷)</option>
            </select>
          </div>
          
          <div className="input-group">
            <label className="input-label">第二个数字</label>
            <input
              type="number"
              value={secondNumber}
              onChange={(e) => setSecondNumber(e.target.value === '' ? '' : Number(e.target.value))}
              className="form-input"
              placeholder="输入第二个数字"
            />
          </div>
          
          <button
            onClick={handleCalculate}
            disabled={isLoading || serverStatus === 'offline'}
            className="calculate-btn"
          >
            {isLoading ? '计算中...' : '计算'}
          </button>
          
          {result !== null && !error && (
            <div className="result-box result-success">
              <p>
                <span className="font-medium">
                  {firstNumber === '' ? '0' : firstNumber} {operationSymbols[operation]} {secondNumber === '' ? '0' : secondNumber} = 
                </span>
                <span className="result-value">{result}</span>
              </p>
            </div>
          )}
          
          {error && (
            <div className="result-box result-error">
              <p className="error-message">{error}</p>
            </div>
          )}
        </div>
      </div>
    </main>
  );
}
