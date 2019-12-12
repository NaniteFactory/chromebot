package chromebot

import (
	"context"
	"errors"
	"sync"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/nanitefactory/chromebot/domain"
)

// Tab represents a running tab.
type Tab struct {
	finMu         sync.Mutex
	fin           chan struct{}
	ctx           context.Context
	cancel        context.CancelFunc
	closeStrategy func(ctx context.Context) // inject close method other than cancel()
}

func newTab(ctx context.Context, cancel context.CancelFunc, closeStrategy func(ctx context.Context)) *Tab {
	if closeStrategy == nil {
		closeStrategy = func(ctx context.Context) { cancel() }
	}
	ret := &Tab{
		finMu:         sync.Mutex{},
		fin:           make(chan struct{}),
		ctx:           ctx,
		cancel:        cancel,
		closeStrategy: closeStrategy,
	}
	go func() {
		<-ctx.Done()
		defer func() { // DCL
			ret.finMu.Lock()
			defer ret.finMu.Unlock()
			select {
			case <-ret.fin:
				return
			default:
				close(ret.fin)
			}
		}()
	}()
	return ret
}

func newMainTab(ctx context.Context, cancel context.CancelFunc) *Tab {
	return newTab(ctx, cancel, func(ctx context.Context) {
		domain.Do(cdp.WithExecutor(ctx, chromedp.FromContext(ctx).Target)).Page().Close()
	})
}

func newExtraTab(ctx context.Context, cancel context.CancelFunc) *Tab {
	return newTab(ctx, cancel, nil)
}

// Close closes this tab.
func (t *Tab) Close() {
	defer func() { // DCL
		t.finMu.Lock()
		defer t.finMu.Unlock()
		select {
		case <-t.fin:
			return
		default:
			t.closeStrategy(t.ctx)
			close(t.fin)
		}
	}()
}

// Dead returns a channel that's closed when work done on behalf of the tab should be dead.
// This notifies a user if and only if the browser is dead or `(*Tab).Close()` has been called.
func (t *Tab) Dead() <-chan struct{} {
	return t.fin
}

// Do runs a cdproto command on this tab.
func (t *Tab) Do() domain.Domain {
	return domain.Do(cdp.WithExecutor(t.ctx, chromedp.FromContext(t.ctx).Target))
}

// Run runs a chromedp action on this tab.
func (t *Tab) Run(actions ...chromedp.Action) error {
	return chromedp.Run(t.ctx, actions...)
}

// Listen adds a callback function which will be called whenever an event is received on this tab.
//
// Usage:
// 	t.Listen(func(ev interface{}) {
// 		switch ev := ev.(type) {
// 		case *page.EventJavascriptDialogOpening:
// 			// do something with ev
// 		case *runtime.EventConsoleAPICalled:
// 			// do something with ev
// 		case *runtime.EventExceptionThrown:
// 			// do something with ev
// 		}
// 	})
//
// Note that the function is called synchronously when handling events.
// The function should avoid blocking at all costs.
// For example, any Actions must be run via a separate goroutine.
func (t *Tab) Listen(fn func(ev interface{})) error {
	select {
	case <-t.ctx.Done():
		return errors.New("tab closed: " + t.ctx.Err().Error())
	case <-t.fin:
		return errors.New("tab closed")
	default:
	}
	chromedp.ListenTarget(t.ctx, fn)
	return nil
}
