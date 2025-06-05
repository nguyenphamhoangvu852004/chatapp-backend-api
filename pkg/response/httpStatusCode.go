package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
)

// message

var message = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "param invalid",
}
