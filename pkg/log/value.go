package log

import (
	"context"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	defaultDepth = 3
	// DefaultCaller is a Valuer that returns the file and line.
	DefaultCaller = Caller(defaultDepth)

	// DefaultTimestamp is a Valuer that returns the current wallclock time.
	DefaultTimestamp = Timestamp(time.RFC3339)
)

type Valuer func(ctx context.Context) interface{}

// Caller returns returns a Valuer that returns a pkg/file:line description of the caller.
func Caller(depth int) Valuer {
	return func(context.Context) interface{} {
		_, file, line, _ := runtime.Caller(depth)
		if strings.LastIndex(file, "/log/filter.go") > 0 {
			depth++
			_, file, line, _ = runtime.Caller(depth)
		}
		if strings.LastIndex(file, "/log/helper.go") > 0 {
			depth++
			_, file, line, _ = runtime.Caller(depth)
		}
		idx := strings.LastIndexByte(file, '/')
		return file[idx+1:] + ":" + strconv.Itoa(line)
	}
}

func Caller2(depth int) Valuer {
	return func(context.Context) interface{} {
		_, file, line, _ := runtime.Caller(depth)
		if strings.LastIndex(file, "/log/filter.go") > 0 {
			depth++
			_, file, line, _ = runtime.Caller(depth)
		}
		if strings.LastIndex(file, "/log/helper.go") > 0 {
			depth++
			_, file, line, _ = runtime.Caller(depth)
		}
		idx := strings.LastIndex(file, "/app")
		return file[idx+1:] + ":" + strconv.Itoa(line)
	}
}

func bindValues(ctx context.Context, keyvals []interface{}) {
	for i := 1; i < len(keyvals); i += 2 {
		if v, ok := keyvals[i].(Valuer); ok {
			keyvals[i] = v(ctx)
		}
	}
}

func containsValuer(keyvals []interface{}) bool {
	for i := 1; i < len(keyvals); i += 2 {
		if _, ok := keyvals[i].(Valuer); ok {
			return true
		}
	}
	return false
}

// Timestamp returns a timestamp Valuer with a custom time format.
func Timestamp(layout string) Valuer {
	return func(context.Context) interface{} {
		return time.Now().Format(layout)
	}
}
