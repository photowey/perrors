/*
 * Copyright © 2022 photowey (photowey@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package perrors

import (
	`fmt`
	`io`
	`runtime`

	"github.com/pkg/errors"
)

// Acknowledgments: https://github.com/pkg/errors

const (
	defaultSkip = 3  // default skip depth
	depth       = 32 // default stack depth
	single      = 1  // [x]
)

var _ StackError = (*perror)(nil)
var _ fmt.Formatter = (*perror)(nil)

// StackError an error wrap with the caller stack info.
type StackError interface {
	t()
	error // std error
}

type perror struct {
	message string
	stack   []uintptr
}

func (pe *perror) t() {}

func (pe *perror) Error() string {
	return pe.message
}

func (pe *perror) Format(s fmt.State, verb rune) {
	_, _ = io.WriteString(s, pe.message)
	_, _ = io.WriteString(s, "\n")

	for _, ps := range pe.stack {
		_, _ = fmt.Fprintf(s, "%+v\n", errors.Frame(ps))
	}
}

func New(message string, skipd ...int) StackError {
	skip := defaultSkip
	switch len(skipd) {
	case single:
		skip = skipd[0]
	}

	return &perror{message: message, stack: Callers(skip)}
}

func Errorf(template string, args ...any) StackError {
	return &perror{
		message: fmt.Sprintf(template, args...),
		stack:   Callers(), // default skipd
	}
}

func Errorsf(skipd int, template string, args ...any) StackError {
	return &perror{
		message: fmt.Sprintf(template, args...),
		stack:   Callers(skipd), // given skipd
	}
}

// Wrap with some extra message into err
func Wrap(err error, message string) StackError {
	if err == nil {
		return nil
	}

	return Wraps(err, message, defaultSkip+1)
}

func Wrapf(err error, template string, args ...any) StackError {
	if err == nil {
		return nil
	}
	message := template
	if len(args) > 0 {
		message = fmt.Sprintf(template, args...)
	}

	return Wraps(err, message, defaultSkip+1)
}

func Wrapsf(err error, skipd int, template string, args ...any) StackError {
	if err == nil {
		return nil
	}
	message := template
	if len(args) > 0 {
		message = fmt.Sprintf(template, args...)
	}

	return Wraps(err, message, skipd+1) // given skipd
}

func Wraps(err error, message string, skipd ...int) StackError {
	if err == nil {
		return nil
	}

	pe, ok := err.(*perror)
	if !ok {
		skip := defaultSkip
		switch len(skipd) {
		case single:
			skip = skipd[0]
		}
		return &perror{
			message: fmt.Sprintf("%s -> %s", message, err.Error()),
			stack:   Callers(skip), // given skipd
		}
	}
	pe.message = fmt.Sprintf("%s -> %s", message, pe.message)

	return pe
}

func Stack(err error, skipd ...int) StackError {
	if err == nil {
		return nil
	}

	if pe, ok := err.(*perror); ok {
		return pe
	}
	skip := defaultSkip
	switch len(skipd) {
	case single:
		skip = skipd[0]
	}

	return &perror{
		message: err.Error(),
		stack:   Callers(skip), // given skipd
	}
}

func Callers(skipd ...int) []uintptr {
	skip := defaultSkip
	switch len(skipd) {
	case single:
		skip = skipd[0]
	}
	var pcs [depth]uintptr
	length := runtime.Callers(skip, pcs[:])

	return pcs[:length]
}
