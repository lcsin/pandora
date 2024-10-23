package message

// ErrCode 错误码
//go:generate stringer -type=ErrCode -linecomment
type ErrCode int64

const (
	// OK 请求成功
	OK ErrCode = 0 // ok
	// Failed 请求失败
	Failed ErrCode = -1 // unknown error

	// BadRequest 参数错误
	BadRequest ErrCode = -4 // 参数错误
	// FileUploadFailed 文件上传失败
	FileUploadFailed ErrCode = -5 // 文件上传失败

	// UserNotFound 用户不存在
	UserNotFound ErrCode = -1001 // 用户不存在或密码错误
	// UserExisted 用户已存在
	UserExisted ErrCode = -1002 // 用户已存在

	// MusicNotFound 音乐不存在
	MusicNotFound ErrCode = -1101 // 音乐不存在
)
