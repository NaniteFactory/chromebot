package domain

// Code generated by chromebot-domain-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/domdebugger"
	"github.com/chromedp/cdproto/runtime"
)

// DOMDebugger executes a cdproto command under DOMDebugger domain.
type DOMDebugger struct {
	ctxWithExecutor context.Context
}

// GetEventListeners returns event listeners of the given object.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger#method-getEventListeners
//
// parameters:
//  - `objectID`: Identifier of the object to return listeners for.
//  - `depth`: This can be nil. (Optional) The maximum depth at which Node children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
//  - `pierce`: This can be nil. (Optional) Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false). Reports listeners for all contexts if pierce is enabled.
//
// returns:
//  - `retListeners`: Array of relevant listeners.
func (doDOMDebugger DOMDebugger) GetEventListeners(objectID runtime.RemoteObjectID, depth *int64, pierce *bool) (retListeners []*domdebugger.EventListener, err error) {
	b := domdebugger.GetEventListeners(objectID)
	if depth != nil {
		b = b.WithDepth(*depth)
	}
	if pierce != nil {
		b = b.WithPierce(*pierce)
	}
	return b.Do(doDOMDebugger.ctxWithExecutor)
}

// RemoveDOMBreakpoint removes DOM breakpoint that was set using
// setDOMBreakpoint.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger#method-removeDOMBreakpoint
//
// parameters:
//  - `nodeID`: Identifier of the node to remove breakpoint from.
//  - `type`: Type of the breakpoint to remove.
func (doDOMDebugger DOMDebugger) RemoveDOMBreakpoint(nodeID cdp.NodeID, typeVal domdebugger.DOMBreakpointType) (err error) {
	b := domdebugger.RemoveDOMBreakpoint(nodeID, typeVal)
	return b.Do(doDOMDebugger.ctxWithExecutor)
}

// RemoveEventListenerBreakpoint removes breakpoint on particular DOM event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger#method-removeEventListenerBreakpoint
//
// parameters:
//  - `eventName`: Event name.
//  - `targetName`: This can be nil. (Optional) EventTarget interface name.
func (doDOMDebugger DOMDebugger) RemoveEventListenerBreakpoint(eventName string, targetName *string) (err error) {
	b := domdebugger.RemoveEventListenerBreakpoint(eventName)
	if targetName != nil {
		b = b.WithTargetName(*targetName)
	}
	return b.Do(doDOMDebugger.ctxWithExecutor)
}

// RemoveInstrumentationBreakpoint removes breakpoint on particular native
// event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger#method-removeInstrumentationBreakpoint
//
// parameters:
//  - `eventName`: Instrumentation name to stop on.
func (doDOMDebugger DOMDebugger) RemoveInstrumentationBreakpoint(eventName string) (err error) {
	b := domdebugger.RemoveInstrumentationBreakpoint(eventName)
	return b.Do(doDOMDebugger.ctxWithExecutor)
}

// RemoveXHRBreakpoint removes breakpoint from XMLHttpRequest.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger#method-removeXHRBreakpoint
//
// parameters:
//  - `url`: Resource URL substring.
func (doDOMDebugger DOMDebugger) RemoveXHRBreakpoint(url string) (err error) {
	b := domdebugger.RemoveXHRBreakpoint(url)
	return b.Do(doDOMDebugger.ctxWithExecutor)
}

// SetDOMBreakpoint sets breakpoint on particular operation with DOM.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger#method-setDOMBreakpoint
//
// parameters:
//  - `nodeID`: Identifier of the node to set breakpoint on.
//  - `type`: Type of the operation to stop upon.
func (doDOMDebugger DOMDebugger) SetDOMBreakpoint(nodeID cdp.NodeID, typeVal domdebugger.DOMBreakpointType) (err error) {
	b := domdebugger.SetDOMBreakpoint(nodeID, typeVal)
	return b.Do(doDOMDebugger.ctxWithExecutor)
}

// SetEventListenerBreakpoint sets breakpoint on particular DOM event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger#method-setEventListenerBreakpoint
//
// parameters:
//  - `eventName`: DOM Event name to stop on (any DOM event will do).
//  - `targetName`: This can be nil. (Optional) EventTarget interface name to stop on. If equal to "*" or not provided, will stop on any EventTarget.
func (doDOMDebugger DOMDebugger) SetEventListenerBreakpoint(eventName string, targetName *string) (err error) {
	b := domdebugger.SetEventListenerBreakpoint(eventName)
	if targetName != nil {
		b = b.WithTargetName(*targetName)
	}
	return b.Do(doDOMDebugger.ctxWithExecutor)
}

// SetInstrumentationBreakpoint sets breakpoint on particular native event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger#method-setInstrumentationBreakpoint
//
// parameters:
//  - `eventName`: Instrumentation name to stop on.
func (doDOMDebugger DOMDebugger) SetInstrumentationBreakpoint(eventName string) (err error) {
	b := domdebugger.SetInstrumentationBreakpoint(eventName)
	return b.Do(doDOMDebugger.ctxWithExecutor)
}

// SetXHRBreakpoint sets breakpoint on XMLHttpRequest.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger#method-setXHRBreakpoint
//
// parameters:
//  - `url`: Resource URL substring. All XHRs having this substring in the URL will get stopped upon.
func (doDOMDebugger DOMDebugger) SetXHRBreakpoint(url string) (err error) {
	b := domdebugger.SetXHRBreakpoint(url)
	return b.Do(doDOMDebugger.ctxWithExecutor)
}