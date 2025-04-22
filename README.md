# ConnectRPC 计算器应用

这是一个使用Go和Next.js构建的全栈计算器应用，通过ConnectRPC进行前后端通信。

![应用截图](screenshots/calculator-screenshot.jpg)

## 功能特点

- 前端：使用Next.js和React构建的现代UI
- 后端：基于Go语言的高性能服务
- 通信：采用ConnectRPC协议而非传统REST API
- 操作：支持基础的加减乘除运算
- 界面：响应式设计，适配不同设备屏幕

## 项目结构

```
calculator-app/
├── backend/           # Go后端
│   ├── calculator/    # 计算器服务实现
│   ├── main.go        # 服务器入口
│   ├── buf.yaml       # Protobuf配置
├── frontend/          # Next.js前端
│   ├── src/           # 源代码
│   │   ├── app/       # Next.js应用
│   │   │   ├── api/connectrpc/  # ConnectRPC客户端
│   │   │   ├── page.tsx  # 计算器页面
```

## 安装依赖

### 后端

1. 安装Go (1.18+): https://golang.org/doc/install
2. 安装Protobuf工具:

```bash
go install github.com/bufbuild/buf/cmd/buf@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
```

3. 安装Go依赖:

```bash
cd backend
go mod tidy
```

4. 生成ConnectRPC代码:

```bash
cd backend
buf generate
```

### 前端

1. 安装Node.js (18+): https://nodejs.org/
2. 安装依赖:

```bash
cd frontend
npm install
```

## 运行应用

### 后端

```bash
cd backend
go run main.go
```

服务器将在 http://localhost:8080 上启动。

### 前端

```bash
cd frontend
npm run dev
```

前端应用将在 http://localhost:3000 上启动。

## 使用说明

1. 在第一个输入框中输入数字
2. 从下拉菜单中选择操作符（加、减、乘、除）
3. 在第二个输入框中输入数字
4. 点击"计算"按钮
5. 结果将显示在下方

## 技术栈

- **后端**: Go, ConnectRPC
- **前端**: Next.js, React, TypeScript, Tailwind CSS
- **通信**: ConnectRPC协议

## 贡献

欢迎提交问题和Pull Request。

## 许可证

MIT 