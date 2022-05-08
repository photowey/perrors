package perrors_test

import (
	`errors`
	`log`
	`testing`

	`github.com/photowey/perrors`
)

func TestNew(t *testing.T) {
	type args struct {
		message string
		depth   []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test perrors New()-depth default",
			args: args{
				message: "hello world",
				depth:   []int{},
			},
		},
		{
			name: "Test perrors New()-depth 0",
			args: args{
				message: "hello world",
				depth:   []int{0},
			},
		},
		{
			name: "Test perrors New()-depth 1",
			args: args{
				message: "hello world",
				depth:   []int{1},
			},
		},
		{
			name: "Test perrors New()-depth 2",
			args: args{
				message: "hello world",
				depth:   []int{2},
			},
		},
		{
			name: "Test perrors New()-depth 3",
			args: args{
				message: "hello world",
				depth:   []int{3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := perrors.New(tt.args.message, tt.args.depth...)
			log.Printf("the New() error info :%v", err)
		})
	}
}

func TestErrorf(t *testing.T) {
	type args struct {
		template string
		args     []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test perrors Errorf()",
			args: args{
				template: "hello %s",
				args:     []any{"world"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := perrors.Errorf(tt.args.template, tt.args.args...)
			log.Printf("the Errorf() error info :%v", err)
		})
	}
}

func TestErrordf(t *testing.T) {
	type args struct {
		depth    int
		template string
		args     []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test perrors Errorsf()",
			args: args{
				depth:    2,
				template: "hello %s",
				args:     []any{"world"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := perrors.Errorsf(tt.args.depth, tt.args.template, tt.args.args...)
			log.Printf("the Errorsf() error info :%v", err)
		})
	}
}

func TestWrap(t *testing.T) {
	type args struct {
		err     error
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test perrors Wrap()",
			args: args{
				err:     errors.New("std error"),
				message: "hello error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := perrors.Wrap(tt.args.err, tt.args.message)
			log.Printf("the Wrap() error info :%v", err)
		})
	}
}

func TestWrapf(t *testing.T) {
	type args struct {
		err      error
		template string
		args     []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test perrors Wrapf()",
			args: args{
				err:      errors.New("std error"),
				template: "hello %s",
				args:     []any{"error"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := perrors.Wrapf(tt.args.err, tt.args.template, tt.args.args...)
			log.Printf("the Wrapf() error info :%v", err)
		})
	}
}

func TestWrapsf(t *testing.T) {
	type args struct {
		err      error
		skipd    int
		template string
		args     []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test perrors Wrapsf()",
			args: args{
				err:      errors.New("std error"),
				skipd:    2,
				template: "hello %s",
				args:     []any{"error"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := perrors.Wrapsf(tt.args.err, tt.args.skipd, tt.args.template, tt.args.args...)
			log.Printf("the Wrapsf() error info :%v", err)
		})
	}
}

func TestWraps(t *testing.T) {
	type args struct {
		err     error
		message string
		skipd   []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test perrors Wraps()",
			args: args{
				err:     errors.New("std error"),
				message: "hello error",
				skipd:   []int{2}, // 3 2
			},
		},
		{
			name: "Test perrors Wraps()",
			args: args{
				err:     errors.New("std error"),
				message: "hello error",
				skipd:   []int{2, 3}, // -> default 3
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := perrors.Wraps(tt.args.err, tt.args.message, tt.args.skipd...)
			log.Printf("the Wraps() error info :%v", err)
		})
	}
}

func TestStack(t *testing.T) {
	type args struct {
		err   error
		skipd []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test perrors Stack()",
			args: args{
				err:   errors.New("std error"),
				skipd: []int{2, 3}, // -> default 3
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := perrors.Stack(tt.args.err, tt.args.skipd...)
			log.Printf("the Stack() error info :%v", err)
		})
	}
}
