package zlog

import (
	"io"
	"os"
)

type Option func(*options)

type options struct {
	level Level
	caller bool
	output io.Writer
}

func newOptions() *options {
	return &options{
		level: InfoLevel,
		caller: true,
		output: os.Stdout,
	}
}

func AddLevel(l Level) Option {
	return func(o *options) {
		o.level = l
	}
}

func AddCaller(enabled bool) Option {
	return func(o *options) {
		o.caller = enabled
	}
}

func AddOutput(w io.Writer) Option {
	return func(o *options) {
		o.output = w
	}
}