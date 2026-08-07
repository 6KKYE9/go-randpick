# go-randpick

从候选列表里随机挑 N 项，支持去重。零依赖，用 `crypto/rand` 做随机源。

## 装

```
go build -o go-randpick .
```

## 用

```
go-randpick -n 3 alice bob carol dave eve     # 随机挑 3 个
go-randpick -n 3 -no-dup alice bob carol      # 不重复，每人最多中一次
go-randpick -no-dup -sorted  < names.txt      # 从文件读，挑完排序
cat items.txt | go-randpick -n 2              # 管道
```

## 选项

| 选项 | 说明 |
|---|---|
| `-n N` | 挑几项，默认 1 |
| `-no-dup` | 不重复挑同一项。超过总数会报错 |
| `-sorted` | 输出前按字母序排 |

## 说明

- 去重模式用 Fisher-Yates 洗牌取前 N 项，保证公平。
- 不去重模式允许挑出的数量超过候选数，会重复出现。
- 随机源用的是 `crypto/rand`，不是 `math/rand`。
- 从标准输入读时，空行自动跳过。
