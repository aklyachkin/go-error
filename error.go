package error

import (
    "runtime"
    "strconv"
)

type TError struct {
    Severity int
    Code int
    msg string
    file string
    line int
    Params []string
}

const (
    S_FATAL = iota
    S_ERROR
    S_WARNING
    S_INFO
)

var s = [...]string{
    S_FATAL: "Fatal",
    S_ERROR: "Error",
    S_WARNING: "Warning",
    S_INFO: "Info",
}

func (e TError) Error() string {
    s := s[e.Severity] + " "  + strconv.Itoa(e.Code) + ": " + e.msg
    if e.file != "" && e.line != 0 {
        s = s + " in " + e.file + ":" + strconv.Itoa(e.line)
    }
    for _, v := range e.Params {
        s = s + " " + v
    }
    return s
}

func (e *TError) Raise(debug bool, params ...string) TError {
    if debug {
        _, f, l, ok := runtime.Caller(1)
        if ok {
            e.file = f
            e.line = l
        }
    }
    e.Params = append(e.Params, params...)
    return *e
}

func New(sev, code int, msg string) TError {
    return TError{Severity: sev, Code: code, msg: msg}
}

