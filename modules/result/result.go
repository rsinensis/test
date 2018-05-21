package result

import "encoding/json"

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

func (r *Result) Json() string {

	b, err := json.Marshal(r)

	if err != nil {
		return ""
	}

	return string(b)
}

func (r *Result) Jsonp(callback string) string {

	b, err := json.Marshal(r)

	if err != nil {
		return ""
	}

	return callback + "(" + string(b) + ")"
}
