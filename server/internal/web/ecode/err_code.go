//go:generate stringer -type=ErrCode -linecomment -output err_code_string.go

package ecode

type ErrCode uint

// feature 10000 - 10699
const (
	InternalServerError ErrCode = 10500
	NotFound            ErrCode = 10404
	MethodNotAllowed    ErrCode = 10405
	BadRequest          ErrCode = 10400
	TooManyRequests     ErrCode = 10429
	Unauthorized        ErrCode = 10401
	BadGateway          ErrCode = 10502
)

// request 10600 - 10699
// 10600-10610 特殊固定值
const (
	WaitQuery         ErrCode = iota + 10600 // 等待查询结果
	ClientUUID        ErrCode = iota + 10610 // 客户端标识错误
	AppID                                 // 应用标识错误
	InvalidCaptcha                        // 验证码错误
	InvalidCiphertext                     // 密文识别错误
)

// auth 10700 - 10799
const (
	LoginFailed       ErrCode = iota + 10700 // 登录失败
	UserNotExist                          // 用户不存在
	UsernameWrong                         // 账号错误
	UserChipEmpty                         // 代币额度为空
	UserChipNotEnough                     // 代币余额不足
)

// internal 20100 - 20199
// 数据库读取写入错误之类问题,不适合外部展示
const (
	DatabaseError ErrCode = iota + 20100 // 数据操作失败
	EventError                        // 事件处理失败
	ModelNotFound
)

// external 20100 - 20199
// 阿里云账户缺钱之类的问题,不适合外部展示
const (
	AliError ErrCode = iota + 20200 // 云服务错误
)