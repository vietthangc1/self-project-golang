// Package errgroup provides synchronization, error propagation, and Context
// cancelation for groups of goroutines working on subtasks of a common task.
package errgroup

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"sync"
)

var ErrPanicRecovered = errors.New("panic recovered")

// A Group is a collection of goroutines working on subtasks that are part of
// the same overall task.
//
// A zero Group is valid and does not cancel on error.
type Group struct {
	cancel func()

	wg sync.WaitGroup

	errOnce sync.Once
	err     error
}

// WithContext returns a new Group and an associated Context derived from ctx.
//
// The derived Context is canceled the first time a function passed to Go
// returns a non-nil error or the first time Wait returns, whichever occurs
// first.
func WithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{cancel: cancel}, ctx
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (g *Group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}

// Go calls the given function in a new goroutine.
//
// The first call to return a non-nil error cancels the group; its error will be
// returned by Wait.
func (g *Group) Go(f func() error) {
	g.wg.Add(1)

	go func() {
		closeOnErr := func(e error) {
			g.errOnce.Do(func() {
				g.err = e
				if g.cancel != nil {
					g.cancel()
				}
			})
		}

		defer func() {
			if r := recover(); r != nil {
				stack := string(debug.Stack())
				if e, ok := r.(error); ok {
					closeOnErr(fmt.Errorf("%w, stack %s", e, stack))
					return
				}
				err := fmt.Errorf("%v: %w, stack: %s", r, ErrPanicRecovered, stack)
				closeOnErr(err)
			}
			g.wg.Done()
		}()

		if err := f(); err != nil {
			closeOnErr(err)
		}
	}()
}
