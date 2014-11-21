package error

import (
    "fmt"
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
    Parent error
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

func Backtrace(err error) {
    if err == nil {
        return
    }
    fmt.Println(err.Error())
    switch err.(type) {
    case TError:
        e1 := err.(TError)
        Backtrace(e1.Parent)
    }
}

func (e *TError) Raise(debug bool, parent error, params ...string) TError {
    e1 := *e
    if debug {
        _, f, l, ok := runtime.Caller(1)
        if ok {
            e1.file = f
            e1.line = l
        }
    }
    e1.Params = append(e1.Params, params...)
    e1.Parent = parent
    return e1
}

func New(sev, code int, msg string) TError {
    return TError{Severity: sev, Code: code, msg: msg}
}

