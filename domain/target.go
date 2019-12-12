package domain

// Code generated by chromebot-domain-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/target"
)

// Target executes a cdproto command under Target domain.
type Target struct {
	ctxWithExecutor context.Context
}

// ActivateTarget activates (focuses) the target.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-activateTarget
//
// parameters:
//  - `targetID`
func (doTarget Target) ActivateTarget(targetID target.ID) (err error) {
	b := target.ActivateTarget(targetID)
	return b.Do(doTarget.ctxWithExecutor)
}

// AttachToTarget attaches to the target with given id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-attachToTarget
//
// parameters:
//  - `targetID`
//  - `flatten`: This can be nil. (Optional) Enables "flat" access to the session via specifying sessionId attribute in the commands. We plan to make this the default, deprecate non-flattened mode, and eventually retire it. See crbug.com/991325.
//
// returns:
//  - `retSessionID`: Id assigned to the session.
func (doTarget Target) AttachToTarget(targetID target.ID, flatten *bool) (retSessionID target.SessionID, err error) {
	b := target.AttachToTarget(targetID)
	if flatten != nil {
		b = b.WithFlatten(*flatten)
	}
	return b.Do(doTarget.ctxWithExecutor)
}

// AttachToBrowserTarget attaches to the browser target, only uses flat
// sessionId mode.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-attachToBrowserTarget
//
// returns:
//  - `retSessionID`: Id assigned to the session.
func (doTarget Target) AttachToBrowserTarget() (retSessionID target.SessionID, err error) {
	b := target.AttachToBrowserTarget()
	return b.Do(doTarget.ctxWithExecutor)
}

// CloseTarget closes the target. If the target is a page that gets closed
// too.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-closeTarget
//
// parameters:
//  - `targetID`
//
// returns:
//  - `retSuccess`
func (doTarget Target) CloseTarget(targetID target.ID) (retSuccess bool, err error) {
	b := target.CloseTarget(targetID)
	return b.Do(doTarget.ctxWithExecutor)
}

// ExposeDevToolsProtocol inject object to the target's main frame that
// provides a communication channel with browser target. Injected object will be
// available as window[bindingName]. The object has the follwing API: -
// binding.send(json) - a method to send messages over the remote debugging
// protocol - binding.onmessage = json => handleMessage(json) - a callback that
// will be called for the protocol notifications and command responses.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-exposeDevToolsProtocol
//
// parameters:
//  - `targetID`
//  - `bindingName`: This can be nil. (Optional) Binding name, 'cdp' if not specified.
func (doTarget Target) ExposeDevToolsProtocol(targetID target.ID, bindingName *string) (err error) {
	b := target.ExposeDevToolsProtocol(targetID)
	if bindingName != nil {
		b = b.WithBindingName(*bindingName)
	}
	return b.Do(doTarget.ctxWithExecutor)
}

// CreateBrowserContext creates a new empty BrowserContext. Similar to an
// incognito profile but you can have more than one.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-createBrowserContext
//
// returns:
//  - `retBrowserContextID`: The id of the context created.
func (doTarget Target) CreateBrowserContext() (retBrowserContextID target.BrowserContextID, err error) {
	b := target.CreateBrowserContext()
	return b.Do(doTarget.ctxWithExecutor)
}

// GetBrowserContexts returns all browser contexts created with
// Target.createBrowserContext method.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-getBrowserContexts
//
// returns:
//  - `retBrowserContextIds`: An array of browser context ids.
func (doTarget Target) GetBrowserContexts() (retBrowserContextIds []target.BrowserContextID, err error) {
	b := target.GetBrowserContexts()
	return b.Do(doTarget.ctxWithExecutor)
}

// CreateTarget creates a new page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-createTarget
//
// parameters:
//  - `url`: The initial URL the page will be navigated to.
//  - `width`: This can be nil. (Optional) Frame width in DIP (headless chrome only).
//  - `height`: This can be nil. (Optional) Frame height in DIP (headless chrome only).
//  - `browserContextID`: This can be nil. (Optional) The browser context to create the page in.
//  - `enableBeginFrameControl`: This can be nil. (Optional) Whether BeginFrames for this target will be controlled via DevTools (headless chrome only, not supported on MacOS yet, false by default).
//  - `newWindow`: This can be nil. (Optional) Whether to create a new Window or Tab (chrome-only, false by default).
//  - `background`: This can be nil. (Optional) Whether to create the target in background or foreground (chrome-only, false by default).
//
// returns:
//  - `retTargetID`: The id of the page opened.
func (doTarget Target) CreateTarget(url string, width *int64, height *int64, browserContextID *target.BrowserContextID, enableBeginFrameControl *bool, newWindow *bool, background *bool) (retTargetID target.ID, err error) {
	b := target.CreateTarget(url)
	if width != nil {
		b = b.WithWidth(*width)
	}
	if height != nil {
		b = b.WithHeight(*height)
	}
	if browserContextID != nil {
		b = b.WithBrowserContextID(*browserContextID)
	}
	if enableBeginFrameControl != nil {
		b = b.WithEnableBeginFrameControl(*enableBeginFrameControl)
	}
	if newWindow != nil {
		b = b.WithNewWindow(*newWindow)
	}
	if background != nil {
		b = b.WithBackground(*background)
	}
	return b.Do(doTarget.ctxWithExecutor)
}

// DetachFromTarget detaches session with given id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-detachFromTarget
//
// parameters:
//  - `sessionID`: This can be nil. (Optional) Session to detach.
func (doTarget Target) DetachFromTarget(sessionID *target.SessionID) (err error) {
	b := target.DetachFromTarget()
	if sessionID != nil {
		b = b.WithSessionID(*sessionID)
	}
	return b.Do(doTarget.ctxWithExecutor)
}

// DisposeBrowserContext deletes a BrowserContext. All the belonging pages
// will be closed without calling their beforeunload hooks.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-disposeBrowserContext
//
// parameters:
//  - `browserContextID`
func (doTarget Target) DisposeBrowserContext(browserContextID target.BrowserContextID) (err error) {
	b := target.DisposeBrowserContext(browserContextID)
	return b.Do(doTarget.ctxWithExecutor)
}

// GetTargetInfo returns information about a target.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-getTargetInfo
//
// parameters:
//  - `targetID`
//
// returns:
//  - `retTargetInfo`
func (doTarget Target) GetTargetInfo(targetID *target.ID) (retTargetInfo *target.Info, err error) {
	b := target.GetTargetInfo()
	if targetID != nil {
		b = b.WithTargetID(*targetID)
	}
	return b.Do(doTarget.ctxWithExecutor)
}

// GetTargets retrieves a list of available targets.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-getTargets
//
// returns:
//  - `retTargetInfos`: The list of targets.
func (doTarget Target) GetTargets() (retTargetInfos []*target.Info, err error) {
	b := target.GetTargets()
	return b.Do(doTarget.ctxWithExecutor)
}

// SetAutoAttach controls whether to automatically attach to new targets
// which are considered to be related to this one. When turned on, attaches to
// all existing related targets as well. When turned off, automatically detaches
// from all currently attached targets.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-setAutoAttach
//
// parameters:
//  - `autoAttach`: Whether to auto-attach to related targets.
//  - `waitForDebuggerOnStart`: Whether to pause new targets when attaching to them. Use Runtime.runIfWaitingForDebugger to run paused targets.
//  - `flatten`: This can be nil. (Optional) Enables "flat" access to the session via specifying sessionId attribute in the commands. We plan to make this the default, deprecate non-flattened mode, and eventually retire it. See crbug.com/991325.
//  - `windowOpen`: This can be nil. (Optional) Auto-attach to the targets created via window.open from current target.
func (doTarget Target) SetAutoAttach(autoAttach bool, waitForDebuggerOnStart bool, flatten *bool, windowOpen *bool) (err error) {
	b := target.SetAutoAttach(autoAttach, waitForDebuggerOnStart)
	if flatten != nil {
		b = b.WithFlatten(*flatten)
	}
	if windowOpen != nil {
		b = b.WithWindowOpen(*windowOpen)
	}
	return b.Do(doTarget.ctxWithExecutor)
}

// SetDiscoverTargets controls whether to discover available targets and
// notify via targetCreated/targetInfoChanged/targetDestroyed events.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-setDiscoverTargets
//
// parameters:
//  - `discover`: Whether to discover available targets.
func (doTarget Target) SetDiscoverTargets(discover bool) (err error) {
	b := target.SetDiscoverTargets(discover)
	return b.Do(doTarget.ctxWithExecutor)
}

// SetRemoteLocations enables target discovery for the specified locations,
// when setDiscoverTargets was set to true.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-setRemoteLocations
//
// parameters:
//  - `locations`: List of remote locations.
func (doTarget Target) SetRemoteLocations(locations []*target.RemoteLocation) (err error) {
	b := target.SetRemoteLocations(locations)
	return b.Do(doTarget.ctxWithExecutor)
}