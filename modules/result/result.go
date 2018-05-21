package result

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
)

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
}

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GetMsg(code int) string {

	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return ""
}

func New(code int, data interface{}) *Result {

	ret := new(Result)

	ret.Code = code
	ret.Msg = MsgFlags[code]
	ret.Data = data

	return ret
}

func Ok(data interface{}) *Result {

	ret := new(Result)

	ret.Code = SUCCESS
	ret.Msg = MsgFlags[SUCCESS]
	ret.Data = data

	return ret
}

func Error() *Result {

	ret := new(Result)

	ret.Code = ERROR
	ret.Msg = MsgFlags[ERROR]

	return ret
}
