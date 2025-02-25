package common

type BizError struct {
	Err  error
	Code int64
	Msg  string
}

func (b *BizError) Error() string {
	return b.Err.Error()
}

func NewDefaultBizError(err error) *BizError {
	return &BizError{
		Err:  err,
		Code: -1,
		Msg:  "系统错误",
	}
}

func (b *BizError) WithErr(err error) *BizError {
	b.Err = err
	return b
}

func (b *BizError) WithCode(Code int64) *BizError {
	b.Code = Code
	return b
}

func (b *BizError) WithMsg(msg string) *BizError {
	b.Msg = msg
	return b
}
