package handlers

type H map[string]any

type MyErr struct {
	Msg  string
	Errs []error
}

func Err(msg string, errs ...error) MyErr {
	return MyErr{msg, errs}
}
func ErrParam(errs ...error) MyErr {
	return MyErr{"参数不正确", errs}
}

const fail = 1
const ok = 0
