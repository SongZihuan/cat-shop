# 猫猫超市（cat-shop）

后端与前端放在同一个 Git 仓库的两个目录里，按单仓（monorepo）管理，不使用 git submodule。

仓库地址：[github.com/SongZihuan/cat-shop

技术栈：Golang + gin + gorm + jwt + yaml ／ Vue 3 + Vite + TypeScript + Pinia + Element Plus。

## 目录

| 目录 | 内容 | 详细说明 |
| --- | --- | --- |
| `shop-backend/` | 后端 API 服务 | [shop-backend/README.md](./shop-backend/README.md) |
| `shop-frontend/` | 前台（用户端）与后台（管理员端） | [shop-frontend/README.md](./shop-frontend/README.md) |

## 运行

后端需要 MySQL 8.0 以上版本，配置文件写法见后端 README：

```shell
cd shop-backend
go run github.com/SongZihuan/cat-shop/shop-backend/src/cmd/version1 -config ./etc/config.yaml
```

前端（作者本人用 pnpm，用 npm 也可以）：

```shell
cd shop-frontend
npm install
npm run dev        # 使用模拟 API
npm run dev-real   # 连接真实后端
```

## 版本

两个子项目各自独立记版本：后端在 [shop-backend/VERSION](./shop-backend/VERSION)，前端在 `shop-frontend/package.json` 的 `version` 字段。

## 许可

[MIT License](./LICENSE)。`shop-backend/` 和 `shop-frontend/` 内的 `LICENSE` 与仓根内容一致。
