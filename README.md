```
house_system_backend
├─ .dockerignore  // docker忽略文件
├─ .vscode // 本工程使用vscode编辑器
│  ├─ launch.json  // 配置的debug运行配置
│  ├─ settings.json // 编辑器配置
│  └─ targets.log //debug日志
├─ Dockerfile    // docker配置
├─ README.md   // 程序说明
├─ config   // 应用配置
├─ controllers // 路由控制器
│  ├─ auth.go  // 鉴权
│  ├─ house.go // 房屋
│  ├─ order.go // 订单
│  ├─ post.go  // 帖
│  └─ user.go  // 用户
├─ go.mod   // 模块配置
├─ go.sum   // 是记录所依赖的项目的版本的锁定
├─ main.go  // 主应用入口
├─ model // 模型配置
│  ├─ house.go // 房屋
│  ├─ order.go // 订单
│  ├─ post.go  // 帖
│  ├─ response_data.go // 返回的范式数据模型
│  └─ user.go  // 用户
└─ utils // 实用工具
   ├─ Config.go // 读取应用配置文件
   └─ db.go // mysql数据库工具

```