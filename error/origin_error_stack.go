package main

import (
	"fmt"
	"github.com/pkg/errors"
	_ "github.com/pkg/errors"
	"os"
	"runtime"
	"strings"
)

func callers() []uintptr {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])
	st := pcs[0:n]
	return st
}

type trace struct {
	m string
	s []uintptr
}

func (e *trace) Error() string {
	var b strings.Builder
	b.WriteString(e.m)
	b.WriteString("\n\n")
	b.WriteString("Traceback:")
	for _, pc := range e.s {
		fn := runtime.FuncForPC(pc)
		b.WriteString("\n")
		f, n := fn.FileLine(pc)
		b.WriteString(fmt.Sprintf("%s:%d", f, n))
	}
	return b.String()
}

// NewTrace creats a simple traceable error.
func NewTrace(message string) error {
	return &trace{m: message, s: callers()}
}

func f() error {
	return NewTrace("ooops")
}

func parseArgs(args []string) error {
	if len(args) < 3 {
		return errors.Errorf("at least 3 arguments.")
	}
	return nil
}

func main() {
	//fmt.Println(f())
	err := parseArgs(os.Args[1:])
	//fmt.Printf("%v\n", err)
	//之所以会有这个效果是因为 pkg/errors 实现了 Formatter 接口:
	//pkg/errors 堆栈跟踪不是没有运行时开销的, 官方给出的指标是每个操作大约 1000-3000 ns.
	//如果超过了你的性能预期, 可以定制成调试模式启用.
	//实现 Formatter 的一个好处是, 我们可以如使用标准库般始终如一, 没有多余的任何 func 调用.
	fmt.Printf("%+v\n", err)
}
