# server part

本程序运行于服务器端，用于辅助客户端之间通过网络传输

## Configuration

1. 服务端ip与监听端口
2. 数据库连接
3. 静态文件托管目录

## dev tips

1. go proxy:

```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## Sync API 示例（Apifox）

`GET /api/v1/sync?space_id=<ULID>&last_pulled_at=0` 响应中，核心关系表示例（仅 `created` 含数据，无时间戳字段）：

```json
{
  "changes": {
    "users": {
      "created": [{ "id": "01ARZ3NDEKTSV4RRFFQ69G5FAV", "nickname": "Alice" }],
      "updated": [],
      "deleted": []
    },
    "spaces": {
      "created": [{ "id": "01ARZ3NDEKTSV4RRFFQ69G5FAW", "name": "Trip" }],
      "updated": [],
      "deleted": []
    },
    "space_members": {
      "created": [
        {
          "id": "01ARZ3NDEKTSV4RRFFQ69G5FAW_01ARZ3NDEKTSV4RRFFQ69G5FAV",
          "space_id": "01ARZ3NDEKTSV4RRFFQ69G5FAW",
          "user_id": "01ARZ3NDEKTSV4RRFFQ69G5FAV"
        }
      ],
      "updated": [],
      "deleted": []
    }
  },
  "timestamp": 1710000000000
}
```

`POST /api/v1/sync?last_pulled_at=<T>` 请求体片段：核心关系表只带必要字段（不要带 `created_at` / `updated_at` / `deleted_at`）；普通数据表仍带完整同步字段。

```json
{
  "users": {
    "created": [],
    "updated": [{ "id": "01ARZ3NDEKTSV4RRFFQ69G5FAV", "nickname": "Aly" }],
    "deleted": []
  },
  "photos": {
    "created": [],
    "updated": [
      {
        "id": "01ARZ3NDEKTSV4RRFFQ69G5FAX",
        "space_id": "01ARZ3NDEKTSV4RRFFQ69G5FAW",
        "uploader_id": "01ARZ3NDEKTSV4RRFFQ69G5FAV",
        "remote_url": "https://example/1.jpg",
        "post_id": "01ARZ3NDEKTSV4RRFFQ69G5FAY",
        "shoted_at": 0,
        "created_at": 1710000000000,
        "updated_at": 1710000000001,
        "deleted_at": 0
      }
    ],
    "deleted": []
  }
}
```

