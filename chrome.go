package chromebot

import (
	"context"
	"errors"
	"sync"

	"github.com/chromedp/chromedp"
)

// Chrome represents a running Chrome browser.
type Chrome struct {
	finMu     sync.Mutex
	fin       chan struct{}
	exeCancel context.CancelFunc
	tabMu     sync.Mutex
	tabs      []*Tab // tab id ascending
}

// New is an alias for NewChrome.
func New(headless bool) *Chrome {
	return NewChrome(headless)
}

// NewChrome runs a new browser.
// Use `(*Chrome).Close()` to free the browser allocated by this.
func NewChrome(headless bool) *Chrome {
	exeCtx, exeCancel := chromedp.NewExecAllocator(context.Background(),
		func() []chromedp.ExecAllocatorOption {
			if headless {
				return append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Headless)
			}
			return append(chromedp.DefaultExecAllocatorOptions[:],
				chromedp.Flag("headless", false),
				chromedp.Flag("hide-scrollbars", false),
				chromedp.Flag("mute-audio", false),
			) // windowed foreground
		}()...)
	tabCtx, tabCancel := chromedp.NewContext(exeCtx)
	chromedp.Run(tabCtx) // ignore error
	ret := &Chrome{
		finMu:     sync.Mutex{},
		fin:       make(chan struct{}),
		exeCancel: exeCancel,
		tabMu:     sync.Mutex{},
		tabs:      []*Tab{newMainTab(tabCtx, tabCancel)},
	}
	go func() {
		<-tabCtx.Done() // when the main tab is dead
		defer func() {  // DCL
			ret.finMu.Lock()
			defer ret.finMu.Unlock()
			select {
			case <-ret.fin:
				return
			default: // free
				ret.exeCancel()
				close(ret.fin)
			}
		}()
	}()
	return ret
}

// Close gracefully closes the browser and all its tabs.
func (c *Chrome) Close() {
	defer func() { // DCL
		c.finMu.Lock()
		defer c.finMu.Unlock()
		select {
		case <-c.fin:
			return
		default:
			c.exeCancel()
			close(c.fin)
		}
	}()
}

// Dead returns a channel that's closed when work done on behalf of the browser should be dead.
// This notifies a user when this browser is closed.
func (c *Chrome) Dead() <-chan struct{} {
	return c.fin
}

// Do runs a cdproto command on this browser.
// You might consider using `(*Tab).Do()` instead, as this function cannot run page specific methods like `Page.navigate()`.
// func (c *Chrome) Do() domain.Domain {
// 	ctx := func() context.Context {
// 		c.tabMu.Lock()
// 		defer c.tabMu.Unlock()
// 		return c.tabs[0].ctx
// 	}()
// 	return domain.Do(cdp.WithExecutor(ctx, chromedp.FromContext(ctx).Browser))
// }

// Listen adds a callback function which will be called whenever a browser event is received on this browser.
//
// Usage:
// 	c.Listen(func(ev interface{}) {
// 		switch ev := ev.(type) {
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
func (c *Chrome) Listen(onBrowserEvent func(ev interface{})) error {
	select {
	case <-c.fin:
		return errors.New("the browser is dead")
	default:
	}
	c.tabMu.Lock()
	defer c.tabMu.Unlock()
	chromedp.ListenBrowser(c.tabs[0].ctx, onBrowserEvent)
	return nil
}

// CountTabs returns the number of tabs created by this Chrome.
func (c *Chrome) CountTabs() int {
	c.tabMu.Lock()
	defer c.tabMu.Unlock()
	return len(c.tabs)
}

// AddNewTab creates a new tab and add it to this Chrome.
// The function returns `nil` when the browser is dead.
// Panic if `chromedp.WithBrowserOption()` passed as an argument.
// Only tab options are accepted.
func (c *Chrome) AddNewTab(opts ...chromedp.ContextOption) (added *Tab) {
	select {
	case <-c.fin:
		return nil
	default:
	}
	c.tabMu.Lock()
	defer c.tabMu.Unlock()
	ctx, cancel := chromedp.NewContext(c.tabs[0].ctx, opts...)
	retTab := newExtraTab(ctx, cancel)
	c.tabs = append(c.tabs, retTab)
	chromedp.Run(retTab.ctx)
	return retTab
}

// Tab is a getter retrieving a tab. Returns nil upon exception.
// The `indexOfTab` is an integer within [0, N) where N is the number of tabs created by the browser.
// E.g. `(*Chrome).Tab(0)` will return the first tab of the browser. Tab(1) for the second one, Tab(2) for the third one and so on.
func (c *Chrome) Tab(indexOfTab int) *Tab {
	c.tabMu.Lock()
	defer c.tabMu.Unlock()
	if indexOfTab >= len(c.tabs) || indexOfTab < 0 {
		return nil
	}
	return c.tabs[indexOfTab]
}
