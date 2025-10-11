# 拉卡拉SDK

这是一个通用的拉卡拉SDK，提供了分账、交易、账户等功能的API接口封装，方便开发者快速集成拉卡拉支付服务。

## 目录结构

```
lklsdk/
├── client.go           # 核心客户端
├── split_ledger.go     # 分账基本功能
├── split_ledger_more.go # 分账扩展功能
├── trade.go            # 交易相关功能
├── account.go          # 账户相关功能
├── merge_pre.go        # 主扫合单交易功能
├── sdk.go              # SDK主入口
├── merchant.go         # 商户相关功能
└── uploadFile.go       # 文件上传相关功能
```

## 安装

```bash
# 将SDK引入到您的项目中
go get -u github.com/black1552/lkl_sdk
```

## 快速开始

### 初始化SDK

```go
import (
	"github.com/black1552/lkl_sdk/lklsdk"
	"github.com/black1552/lkl_sdk/model"
	"github.com/black1552/lkl_sdk/consts"
	"github.com/gogf/gf/v2/os/gctx"
)

// 创建配置JSON字符串
cfgJson := `{
	"public_key": "your_public_key",        // 公钥字符串
	"private_key": "your_private_key",      // 私钥字符串
	"app_id": "your_app_id",                // lakala应用ID
	"serial_no": "your_serial_no",          // 序列号
	"sub_app_id": "your_sub_app_id",        // 子应用ID 微信AppId
	"version": "3.0",                       // lakala版本号
	"account_type": "WECHAT",               // 账户类型
	"trans_type": "71",                     // 交易类型
	"notify_url": "your_notify_url"         // 回调地址
}`

// 初始化SDK（使用泛型指定响应类型）
sdk := lklsdk.NewSDK[model.ResponseType](gctx.New(), cfgJson)
```

## 功能模块

### 1. 主扫合单交易

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.MergePreorderResponse](gctx.New(), cfgJson)

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
sdk := lklsdk.NewSDK[model.ApplyLedgerMerResponse](gctx.New(), cfgJson)

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
	SplitRange:           consts.SPLIT_RANGE_ALL,                               // 分账范围（ALL全部交易分账，MARK标记交易分账）
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
sdk := lklsdk.NewSDK[model.TradeQueryResponse](gctx.New(), cfgJson)

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
sdk := lklsdk.NewSDK[model.OrderSplitLedgerResponse](gctx.New(), cfgJson)

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
	CalType:       consts.CAL_TYPE_AMOUNT,              // 分账计算类型（0-按金额，1-按比例）
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

### 5. 退款

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.RefundResponse](gctx.New(), cfgJson)

// 构建请求参数
refundReq := &model.RefundReqData{
	MerchantNo: "your_merchant_no",  // 商户号
	TermNo:     "your_term_no",      // 终端号
	OutTradeNo: "original_out_trade_no", // 原商户交易流水号
	OutRefundNo: "your_refund_out_trade_no", // 退款商户流水号
	RefundAmount: "100",  // 退款金额，单位为分
	TotalAmount: "100",   // 原交易总金额，单位为分
	RefundReason: "退款原因", // 退款原因
	NotifyUrl: "https://your-notify-url.com", // 回调地址
}

// 调用接口
refundResp, err := sdk.Refound(refundReq)
if err != nil {
	log.Printf("退款失败: %v\n", err)
}

// 使用SuccessOrFail方法判断请求是否成功
if !refundResp.SuccessOrFail() {
	log.Printf("退款失败: %s\n", refundResp.Msg)
}

// 处理成功响应
// 可以从refundResp.RespData中获取退款结果数据
```

### 6. 账户余额查询

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.BalanceQueryResponse](gctx.New(), cfgJson)

// 构建请求参数
balanceQueryReq := &model.BalanceQueryReqData{
	MerchantNo: config.MerchantNo,
	OrgNo:      "",    // 机构号
	PayNo:      "",    // 支付单号
	PayType:    consts.PAY_TYPE_CARD,    // 支付类型（CARD-卡支付，COUPON-优惠券支付，DISCOUNT-折扣支付，POINT-积分支付，MIX-混合支付）
	MgtFlag:    consts.MGT_FLAG_NO,    // 管理标志（NO-普通交易，YES-管理交易）
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

### 7. 账户提现

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.WithdrawResponse](gctx.New(), cfgJson)

// 构建请求参数
withdrawReq := &model.WithdrawReqData{
	MerchantNo: "your_merchant_no",  // 商户号
	TermNo:     "your_term_no",      // 终端号
	OutTradeNo: "your_withdraw_out_trade_no", // 提现商户流水号
	Amount:     "10000",             // 提现金额，单位为分
	Currency:   "CNY",               // 货币类型
	AccountType: "WECHAT",           // 账户类型
}

// 调用接口
withdrawResp, err := sdk.Withdraw(withdrawReq)
if err != nil {
	log.Printf("账户提现失败: %v\n", err)
}

// 使用SuccessOrFail方法判断请求是否成功
if !withdrawResp.SuccessOrFail() {
	log.Printf("账户提现失败: %s\n", withdrawResp.Msg)
}
```

### 8. 商户进件

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.MerchantApplyResponse](gctx.New(), cfgJson)

// 构建请求参数
merchantApplyReq := &model.MerchantApplyReqData{
	// 填写商户进件所需的各项参数
	MerchantName: "商户名称",
	ContactName:  "联系人姓名",
	ContactPhone: "联系电话",
	// ... 其他必要参数
}

// 调用接口
merchantApplyResp, err := sdk.AddMer(merchantApplyReq)
if err != nil {
	log.Printf("商户进件失败: %v\n", err)
}

// 使用SuccessOrFail方法判断请求是否成功
if !merchantApplyResp.SuccessOrFail() {
	log.Printf("商户进件失败: %s\n", merchantApplyResp.Msg)
}
```

### 9. 商户进件信息查询

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.QueryMerResponse](gctx.New(), cfgJson)

// 构建请求参数
queryMerReq := &model.QueryMerRequestData{
	MerchantNo: "your_merchant_no", // 商户号
	// ... 其他必要参数
}

// 调用接口
queryMerResp, err := sdk.QueryMerchant(queryMerReq)
if err != nil {
	log.Printf("商户进件信息查询失败: %v\n", err)
}

// 使用SuccessOrFail方法判断请求是否成功
if !queryMerResp.SuccessOrFail() {
	log.Printf("商户进件信息查询失败: %s\n", queryMerResp.Msg)
}
```

### 10. 商户进件信息校验

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.MerValidateResponse](gctx.New(), cfgJson)

// 构建请求参数
merValidateReq := &model.MerValidateRequestData{
	// 填写商户进件校验所需的各项参数
	MerchantName: "商户名称",
	// ... 其他必要参数
}

// 调用接口
merValidateResp, err := sdk.MerValidate(merValidateReq)
if err != nil {
	log.Printf("商户进件信息校验失败: %v\n", err)
}

// 使用SuccessOrFail方法判断请求是否成功
if !merValidateResp.SuccessOrFail() {
	log.Printf("商户进件信息校验失败: %s\n", merValidateResp.Msg)
}
```

### 11. 商户进件复议提交

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.ReConfSubmitResponse](gctx.New(), cfgJson)

// 构建请求参数
reConfSubmitReq := &model.ReConfSubmitRequestData{
	// 填写商户进件复议所需的各项参数
	MerchantNo: "your_merchant_no",
	// ... 其他必要参数
}

// 调用接口
reConfSubmitResp, err := sdk.ReconsiderSubmit(reConfSubmitReq)
if err != nil {
	log.Printf("商户进件复议提交失败: %v\n", err)
}

// 使用SuccessOrFail方法判断请求是否成功
if !reConfSubmitResp.SuccessOrFail() {
	log.Printf("商户进件复议提交失败: %s\n", reConfSubmitResp.Msg)
}
```

### 12. 聚合预下单

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.PreorderResponse](gctx.New(), cfgJson)

// 构建请求参数
preOrderReq := &model.PreorderReqData{
	MerchantNo:  "your_merchant_no",  // 商户号
	TermNo:      "your_term_no",      // 终端号
	OutTradeNo:  "your_out_trade_no", // 商户交易流水号
	AccountType: "WECHAT",            // 账户类型
	TransType:   "71",                // 交易类型
	Amount:      "100",               // 金额，单位为分
	Subject:     "测试订单",           // 订单标题
	NotifyUrl:   "https://your-notify-url.com", // 回调地址
}

// 调用接口
preOrderResp, err := sdk.PreOrder(preOrderReq)
if err != nil {
	log.Printf("聚合预下单失败: %v\n", err)
}

// 使用SuccessOrFail方法判断请求是否成功
if !preOrderResp.SuccessOrFail() {
	log.Printf("聚合预下单失败: %s\n", preOrderResp.Msg)
}
```

### 13. 文件上传查询

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.UploadFileResponse](gctx.New(), cfgJson)

// 构建请求参数
uploadFileReq := &model.UploadFileReqData{
	// 填写文件上传查询所需的各项参数
	// ...
}

// 调用接口
uploadFileResp, err := sdk.UploadFileQuery(uploadFileReq)
if err != nil {
	log.Printf("文件上传查询失败: %v\n", err)
}

// 使用SuccessOrFail方法判断请求是否成功
if !uploadFileResp.SuccessOrFail() {
	log.Printf("文件上传查询失败: %s\n", uploadFileResp.Msg)
}
```

### 14. 分账接收方创建申请

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.ApplyLedgerReceiverResponse](gctx.New(), cfgJson)

// 构建请求参数
applyLedgerReceiverReq := &model.ApplyLedgerReceiverReqData{
	MerchantNo:    "your_merchant_no",  // 商户号
	TermNo:        "your_term_no",      // 终端号
	ReceiverType:  "PERSON",            // 接收方类型
	ReceiverName:  "接收方名称",         // 接收方名称
	ReceiverNo:    "接收方编号",         // 接收方编号
	ReceiverAccNo: "接收方账号",         // 接收方账号
	// ... 其他必要参数
}

// 调用接口
applyLedgerReceiverResp, err := sdk.ApplyLedgerReceiver(applyLedgerReceiverReq)
if err != nil {
	log.Printf("分账接收方创建申请失败: %v\n", err)
}

// 使用SuccessOrFail方法判断请求是否成功
if !applyLedgerReceiverResp.SuccessOrFail() {
	log.Printf("分账接收方创建申请失败: %s\n", applyLedgerReceiverResp.Msg)
}
```

### 15. 分账关系绑定申请

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.ApplyBindResponse](gctx.New(), cfgJson)

// 构建请求参数
applyBindReq := &model.ApplyBindReqData{
	MerchantNo:   "your_merchant_no",  // 商户号
	TermNo:       "your_term_no",      // 终端号
	ReceiverType: "PERSON",            // 接收方类型
	ReceiverNo:   "接收方编号",         // 接收方编号
	BindType:     "1",                 // 绑定类型
	// ... 其他必要参数
}

// 调用接口
applyBindResp, err := sdk.ApplyBind(applyBindReq)
if err != nil {
	log.Printf("分账关系绑定申请失败: %v\n", err)
}

// 使用SuccessOrFail方法判断请求是否成功
if !applyBindResp.SuccessOrFail() {
	log.Printf("分账关系绑定申请失败: %s\n", applyBindResp.Msg)
}
```

### 16. 可分账金额查询

```go
// 初始化特定响应类型的SDK
sdk := lklsdk.NewSDK[model.SplitBalanceResponse](gctx.New(), cfgJson)

// 构建请求参数
splitBalanceReq := &model.SplitBalanceReqData{
	MerchantNo: "your_merchant_no",  // 商户号
	TermNo:     "your_term_no",      // 终端号
	// ... 其他必要参数
}

// 调用接口
splitBalanceResp, err := sdk.QuerySplitBalance(splitBalanceReq)
if err != nil {
	log.Printf("可分账金额查询失败: %v\n", err)
}

// 使用SuccessOrFail方法判断请求是否成功
if !splitBalanceResp.SuccessOrFail() {
	log.Printf("可分账金额查询失败: %s\n", splitBalanceResp.Msg)
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