# Go Blog 后端

这是 Go Blog 的 HTTP API 服务，使用 Go、Gin、GORM 和 Elasticsearch 客户端编写。后端负责用户与权限、文章和评论、图片上传、网站配置等功能。前端静态文件由 Vite 开发服务器或 Nginx 等静态服务器提供；Gin 服务负责 API 和上传文件访问。

## 后端如何启动

程序入口是 `main.go`。启动时大致按以下顺序工作：

1. 从当前工作目录读取 `config.yaml`，并创建日志记录器。
2. 初始化 JWT 相关配置和本地黑名单缓存。
3. 初始化 MySQL（GORM）、Redis 和 Elasticsearch 客户端。
4. 如果命令行带有维护参数，执行数据库迁移、索引操作、数据导入导出或管理员创建，然后退出。
5. 没有维护参数时，注册定时任务并启动 Gin HTTP 服务。

这一点对初学者很重要：**即使命令是 `-sql`、`-es` 或 `-admin`，程序也会先读取配置并初始化各项依赖。**本地开发时请先启动 MySQL、Redis 和 Elasticsearch，并检查 `config.yaml` 中的连接地址。

## 环境准备

- Go 1.27.1。
- MySQL。项目使用 GORM 自动创建和更新表结构，但 MySQL 数据库本身需要先创建。
- Redis。当前初始化逻辑会连接 Redis 并执行 `PING`。
- Elasticsearch 8.19.0（本地开发建议版本）。Go Elasticsearch 客户端依赖版本见 `go.mod`。
- Docker 可用于启动本地依赖服务，但不是运行 Go 程序的必需条件。

以下命令仅适用于本地开发，会关闭 Elasticsearch 安全认证并使用简单的 MySQL 密码。不要将这些设置用于可从公网访问的服务器：

```bash
docker run -d --name mysql -p 127.0.0.1:3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=blog_db mysql:8.0
docker run -d --name redis -p 127.0.0.1:6379:6379 redis:7
docker run -d --name es -p 127.0.0.1:9200:9200 -e discovery.type=single-node -e xpack.security.enabled=false -e xpack.security.http.ssl.enabled=false -e ES_JAVA_OPTS="-Xms512m -Xmx512m" docker.elastic.co/elasticsearch/elasticsearch:8.19.0
```

如果这些容器已经创建但停止了，可以用 `docker start mysql redis es` 启动。容器名 `mysql` 也被当前 `-sql-export` 命令使用。

## 配置 `config.yaml`

配置文件是本目录下的 [`config.yaml`](config.yaml)（从最大仓库根目录看是 `server/config.yaml`）。程序按相对路径读取它，因此执行下面的命令前要先进入 `server` 目录。配置文件包含数据库地址、密钥等信息；请按本机环境填写，不要把生产密钥复制到公开仓库或 README。

| 配置段 | 需要关注的字段 | 作用 |
| --- | --- | --- |
| `mysql` | `host`、`port`、`db_name`、`username`、`password`、`config` | MySQL 连接。`db_name` 对应的数据库必须先存在。 |
| `redis` | `address`、`password`、`db` | Redis 地址、密码和逻辑数据库编号。没有密码时按本地 Redis 的配置留空。 |
| `es` | `url`、`username`、`password` | Elasticsearch 地址和认证信息。关闭本地 ES 安全认证时用户名、密码通常留空。 |
| `system` | `host`、`port`、`env`、`router_prefix`、`sessions_secret`、`oss_type` | 监听地址、端口、Gin 模式、API 前缀、会话密钥和存储方式。 |
| `jwt` | 两个密钥、`access_token_expiry_time`、`refresh_token_expiry_time`、`issuer` | JWT 签发与过期设置。时长支持 `d`、`h`、`m`、`s`，例如 `15m`、`30d`。 |
| `upload` | `path`、`size` | 本地上传目录和图片大小上限。使用本地存储时，后端会通过该路径提供上传文件。 |
| `website` | 标题、标语、站点名、地址、联系方式等 | 网站公开信息；创建管理员时会使用 `website.name` 和 `website.address` 填充管理员资料。 |
| `email`、`qq`、`qiniu`、`gaode` | 对应服务的地址、开关和密钥 | 邮箱验证码、QQ 登录、七牛云存储和高德功能；不使用某项功能时不需要配置它的真实凭据。 |
| `zap`、`captcha` | 日志轮转、控制台输出和验证码参数 | 日志输出及验证码图片设置。 |

本地前端默认示例使用 API 前缀 `/api` 和后端端口 `8080`。请确保 `system.router_prefix`、`system.port` 与 `web/.env.local` 中的 `VITE_BASE_API`、`VITE_SERVER_URL` 相匹配。更换 API 前缀时，也要检查 Vite 和生产反向代理的路由规则。

下面只展示与上面 Docker 命令对应的本地连接字段。请修改现有 `config.yaml` 中的对应字段，并保留其余验证码、JWT、网站和日志设置；不要用这个片段覆盖整份配置。这里的 `root` 密码只适用于上述本地 MySQL 容器：

```yaml
mysql:
  host: 127.0.0.1
  port: 3306
  db_name: blog_db
  username: root
  password: root
  config: charset=utf8mb4&parseTime=True&loc=Local

redis:
  address: 127.0.0.1:6379
  password: ""
  db: 0

es:
  url: http://127.0.0.1:9200
  username: ""
  password: ""

system:
  host: 0.0.0.0
  port: 8080
  env: debug
  router_prefix: /api
  sessions_secret: 请替换为本地随机字符串
  oss_type: local
```

## 初始化并运行

如果你在最大仓库根目录，先进入 `server`；如果终端已经在本目录，就跳过 `cd server`：

```bash
cd server
go mod download
```

首次运行时，按顺序创建 MySQL 表结构、Elasticsearch 索引，并创建管理员：

```bash
go run . -sql
go run . -es
go run . -admin
```

- `-sql` 使用 GORM `AutoMigrate` 创建或更新应用所需的 MySQL 表；它不会替你创建 MySQL 数据库。
- `-es` 创建文章索引 `article_index` 及其字段映射。若索引已经存在，程序会询问是否删除并重建；回答 `y` 会删除该索引中的现有数据。没有明确需要时不要确认删除。
- `-admin` 会交互式要求输入邮箱和密码。密码要求 8 到 20 个字符；管理员昵称和地址取自 `website.name` 与 `website.address`。

初始化完成后，启动 API 服务：

```bash
go run .
```

服务监听地址由 `system.host` 和 `system.port` 决定，API 路径由 `system.router_prefix` 决定。例如配置为 `0.0.0.0:8080` 和 `/api` 时，可以在本机访问：

```text
http://127.0.0.1:8080/api/article/tags
```

前端开发服务器会代理 `/api` 和 `/uploads` 请求。生产部署时，应由反向代理分别提供前端 `dist` 文件，并将 API 与上传路径转发到本服务。

## 命令行维护操作

一次只运行一个维护参数：

| 命令 | 作用 |
| --- | --- |
| `go run . -sql` | 创建或更新 MySQL 表结构。 |
| `go run . -es` | 创建文章索引和映射；已有索引时会询问是否删除重建。 |
| `go run . -admin` | 交互式创建管理员。 |
| `go run . -es-export` | 导出 `article_index` 到当前目录下的 `es_日期.json` 文件。 |
| `go run . -es-import=./es_日期.json` | 从导出 JSON 导入 Elasticsearch。**此操作会自动删除并重建现有文章索引。** |
| `go run . -sql-export` | 导出 SQL 文件；当前实现通过 `docker exec mysql ...` 调用容器，因此要求 MySQL 容器名为 `mysql`。 |
| `go run . -sql-import=./backup.sql` | 执行 SQL 文件中的语句。导入逻辑较简单，正式恢复数据库前请先备份并确认文件兼容。 |

Elasticsearch 文章索引中，`category` 和 `tags` 使用 Keyword 类型，搜索时按完整值筛选；标题、摘要和正文使用全文搜索字段。修改索引映射后，旧索引不会自动变更，需要备份数据并重建或重新导入索引。

## 代码目录

| 目录 | 职责 |
| --- | --- |
| `api` | Gin HTTP 处理函数：读取请求参数、调用业务层并组织响应。 |
| `router` | 按功能注册 API 路径，并将路径连接到处理函数。 |
| `middleware` | JWT 登录验证、管理员权限、登录记录、日志和异常恢复。 |
| `service` | 主要业务逻辑，例如文章、用户、评论、配置和 Elasticsearch 操作。 |
| `model/request`、`model/response` | API 请求与响应的数据结构。 |
| `model/database` | GORM 使用的 MySQL 表模型。 |
| `model/elasticsearch` | Elasticsearch 文档结构和索引映射。 |
| `config` | `config.yaml` 各配置段对应的 Go 结构体。 |
| `global` | 全局配置、日志、数据库、Redis、Elasticsearch 客户端及缓存实例。 |
| `initialize`、`core` | 初始化依赖、路由、日志和 HTTP 服务。 |
| `flag` | 数据库迁移、导入导出、索引和管理员等命令行操作。 |
| `task` | 定时任务，例如同步文章浏览量、热榜和日历数据。 |
| `utils` | 分页、JWT、上传、图片、时间等通用工具。 |

新增接口时，可以先在 `router` 注册路径，再在 `api` 处理输入和响应，把具体业务放入 `service`。公开接口、登录用户接口和管理员接口由不同的路由组挂载；不要只依赖前端隐藏页面来保护操作。

## 常见问题

- **提示找不到 `config.yaml`**：先切换到 `server` 目录，再运行 `go run .`。
- **启动时数据库或 Redis 报错**：检查服务是否运行、地址是否能从当前运行环境访问、账号密码和数据库名是否正确。
- **前端请求 404**：检查后端 `system.router_prefix`、前端 `VITE_BASE_API` 以及 Vite/Nginx 代理路径是否一致。
- **文章标签筛选与旧数据不匹配**：确认 Elasticsearch 使用当前 `article_index` 映射。重建前先导出数据；`-es-import` 会替换现有索引。
- **图片能上传但浏览器打不开**：检查 `system.oss_type`、`upload.path`、本地文件权限，以及前端代理或生产服务器对 `/uploads` 的转发。
