# 拉卡拉SDK

这是一个通用的拉卡拉SDK，提供了分账、交易、账户等功能的API接口封装，方便开发者快速集成拉卡拉支付服务。

## 目录结构

```
lklsdk/
├── client.go     # 核心客户端
├── split_ledger.go     # 分账基本功能
├── split_ledger_more.go # 分账扩展功能
├── trade.go      # 交易相关功能
├── account.go    # 账户相关功能
├── merge_pre.go  # 主扫合单交易功能
└── sdk.go        # SDK主入口
```

## 安装

```bash
# 将SDK引入到您的项目中
go get -u github.com/black1552/lkl_sdk/lklsdk
```

## 快速开始

### 初始化SDK

```go
import (
	"github.com/black1552/lkl_sdk/lklsdk"
	"github.com/black1552/lkl_sdk/model"
	"github.com/gogf/gf/v2/os/gctx"
)

// 创建配置
config := &lklsdk.Config{
	AppID:        "your_app_id",        // 拉卡拉分配的AppID
	TermNo:       "your_term_no",       // 终端号
	MerchantNo:   "your_merchant_no",   // 商户号
	SettleMerNo:  "your_settle_mer_no", // 结算商户号
	SettleTermNo: "your_settle_term_no", // 结算终端号
	AccountType:  "WECHAT",              // 账户类型，如WECHAT、ALIPAY等
	TransType:    "71",                  // 交易类型
	Version:      "3.0",                 // API版本
	NotifyURL:    "your_notify_url",     // 回调URL
	SerialNo:     "your_mch_api_key",    // 商户API密钥
}

// 初始化SDK（使用泛型指定响应类型）
sdk := lklsdk.NewSDK[model.ResponseType](gctx.New(), config)
```

## 功能模块

### 1. 主扫合单交易

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.MergePreorderResponse](gctx.New(), config)

// 准备拆单信息
outSplitInfo := []*model.OutSplitInfo{
	{
		OutSubTradeNo: "子交易流水号1",
		MerchantNo:    config.MerchantNo,
		TermNo:        config.TermNo,
		Amount:        "100", // 1元
	},
	{
		OutSubTradeNo: "子交易流水号2",
		MerchantNo:    config.MerchantNo,
		TermNo:        config.TermNo,
		Amount:        "200", // 2元
	},
}

// 构建位置信息
locationInfo := &model.LocationInfo{
	RequestIp: "127.0.0.1",
}

// 构建请求参数
mergePreorderReq := &model.MergePreorderReqData{
	MerchantNo:   config.MerchantNo,
	TermNo:       config.TermNo,
	OutTradeNo:   "商户交易流水号",
	OutSplitInfo: outSplitInfo,
	AccountType:  "WECHAT",
	TransType:    "71",
	TotalAmount:  "300", // 3元
	LocationInfo: locationInfo,
	Subject:      "测试订单",
	NotifyUrl:    "https://your-notify-url.com",
}

// 调用接口
mergePreorderResp, err := sdk.MergePreOrder(mergePreorderReq)
if err != nil {
	log.Printf("主扫合单交易失败: %v\n", err)
}

// 使用SuccessOrFail方法判断请求是否成功
if !mergePreorderResp.SuccessOrFail() {
	log.Printf("主扫合单交易失败: %s\n", mergePreorderResp.Msg)
}
```

### 2. 商户分账业务开通申请

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.ApplyLedgerMerResponse](gctx.New(), config)

// 构建请求参数
applyLedgerReq := &model.ApplyLedgerMerReqData{
	Version:              "1.0",
	OrderNo:              "12345678901234567890123456789012", // 32位订单号
	OrgCode:              "123456789012",                      // 机构代码
	MerInnerNo:           "1234567821",                        // 拉卡拉内部商户号
	MerCupNo:             "",                                  // 银联商户号（与内部商户号二选一）
	ContactMobile:        "13311111111",                       // 联系手机号
	SplitLowestRatio:     3.51,                                 // 最低分账比例
	SplitEntrustFileName: "授权委托书.pdf",                     // 授权委托书文件名
	SplitEntrustFilePath: "path",                              // 授权委托书文件路径
	SplitRange:           "ALL",                               // 分账范围（ALL全部交易分账，MARK标记交易分账）
	RetUrl:               "notifyUrl.com",                     // 回调URL
}

// 调用接口
expectResp, err := sdk.ApplyLedgerMer(applyLedgerReq)
if err != nil {
	log.Printf("商户分账业务开通申请失败: %v\n", err)
}

// 使用SuccessOrFail方法判断请求是否成功
if !expectResp.SuccessOrFail() {
	log.Printf("商户分账业务开通申请失败: %s\n", expectResp.RetMsg)
}
```

### 3. 交易查询

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.TradeQueryResponse](gctx.New(), config)

// 构建请求参数
tradeQueryReq := &model.TradeQueryReqData{
	MerchantNo: config.MerchantNo,
	TermNo:     config.TermNo,
	OutTradeNo: "商户订单号", // 替换为实际的商户订单号
}

// 调用接口
tradeQueryResp, err := sdk.TradeQuery(tradeQueryReq)
if err != nil {
	log.Printf("交易查询失败: %v\n", err)
}

// 使用SuccessOrFail方法判断请求是否成功
if !tradeQueryResp.SuccessOrFail() {
	log.Printf("交易查询失败: %s\n", tradeQueryResp.Msg)
}
```

### 4. 订单分账

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.OrderSplitLedgerResponse](gctx.New(), config)

// 准备分账接收方数据
var recvDatas []*model.OrderSplitLedgerRecvDatas
// 可以向recvDatas数组中添加分账接收方信息

// 构建请求参数
splitLedgerReq := &model.OrderSplitLedgerReqData{
	MerchantNo:    config.MerchantNo,
	LogDate:       "",              // 交易日期，格式为yyyyMMdd
	LogNo:         "",              // 拉卡拉对账单流水号
	OutSeparateNo: "",              // 商户分账指令流水号
	TotalAmt:      "",              // 分账总金额，单位为分
	LklOrgNo:      "",              // 拉卡拉机构编号
	CalType:       "",              // 分账计算类型（0-按金额，1-按比例）
	NotifyUrl:     "",              // 回调地址
	RecvDatas:     recvDatas,        // 分账接收方数据
}

// 调用接口
splitLedgerResp, err := sdk.OrderSplitLedger(splitLedgerReq)
if err != nil {
	log.Printf("订单分账失败: %v\n", err)
}

// 使用SuccessOrFail方法判断请求是否成功
if !splitLedgerResp.SuccessOrFail() {
	log.Printf("订单分账失败: %s\n", splitLedgerResp.Msg)
}
```

### 5. 账户余额查询

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.BalanceQueryResponse](gctx.New(), config)

// 构建请求参数
balanceQueryReq := &model.BalanceQueryReqData{
	MerchantNo: config.MerchantNo,
	OrgNo:      "",    // 机构号
	PayNo:      "",    // 支付单号
	PayType:    "",    // 支付类型
	MgtFlag:    "",    // 管理标志
}

// 调用接口
balanceQueryResp, err := sdk.BalanceQuery(balanceQueryReq)
if err != nil {
	log.Printf("账户余额查询失败: %v\n", err)
}

// 使用SuccessOrFail方法判断请求是否成功
if !balanceQueryResp.SuccessOrFail() {
	log.Printf("账户余额查询失败: %s\n", balanceQueryResp.Msg)
}
```

## 错误处理

SDK使用两层错误处理机制，请确保同时检查网络错误和业务响应状态：

```go
// 1. 首先检查网络或SDK层面的错误
resp, err := sdk.SomeFunction(req)
if err != nil {
	log.Printf("调用接口失败: %v\n", err)
	// 处理网络错误或SDK内部错误
	return err
}

// 2. 然后使用SuccessOrFail()方法判断业务响应是否成功
if !resp.SuccessOrFail() {
	log.Printf("业务处理失败: %s\n", resp.Msg) // 或resp.RetMsg
	// 处理业务失败情况
	return errors.New("业务处理失败")
}

// 处理成功响应
// 使用resp获取返回的数据
```

## 注意事项

1. 请妥善保管您的AppID、商户号、密钥等敏感信息
2. 确保网络环境稳定，避免请求超时
3. 建议添加请求重试机制，处理网络波动
4. 请遵循拉卡拉的接口规范，不要频繁调用接口
5. 如遇到问题，请参考拉卡拉官方文档或联系技术支持