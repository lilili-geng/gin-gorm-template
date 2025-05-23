<h1 align="center" style="margin: 30px 0 30px; font-weight: bold; font-size: 30px">gin-gorm-template</h1>
<h4 align="center">基于Go+Gin+Gorm实现的服务端脚手架</h4>

## 平台简介

- 本仓库为后端技术栈 [Gin](https://gin-gonic.com/zh-cn/docs) + [Gorm](https://gorm.io/zh_CN/docs/index.html) 的 `golang` 版本。
- 配套前端代码仓库地址  [Vue3-ts](https://github.com/lilili-geng/vue-ts-template)

## 后端运行

> **提示：** 运行前请先安装好 `go` 环境，版本 `1.21` 以上。

    # 克隆项目
    git clone https://github.com/lilili-geng/gin-gorm-template?tab=readme-ov-file

    # 进入项目目录
    cd gin-gorm-template

    # 修改配置文件
    application.yaml

    # 安装依赖
    go mod tidy

    # 启动服务
    go run main.go

    # 热更新启动服务
    fresh

## 前端运行

    # 调整 .env 文件
    VUE_APP_BASE_API = '/dev-api' 改为 VITE_APP_BASE_API = '/api'

    # 调整 vite.config.js 文件
    server: {
      port: 8000,
      open: false,
      proxy: {
        // https://cn.vitejs.dev/config/#server-proxy
        '/api': {
          target: 'http://localhost:3000',
          changeOrigin: true,
          rewrite: (p) => p.replace(/^\/api/, '/api')
        }
      }
    },

    # 安装依赖
    yarn

    # 启动服务
    yarn dev

## 后端打包

    # 打包
    go build main.go

## 内置功能

1.  用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2.  部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3.  岗位管理：配置系统用户所属担任职务。
4.  菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5.  角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6.  字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7.  参数管理：对系统动态配置常用参数。
8.  操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
9.  登录日志：系统登录日志记录查询包含登录异常。
