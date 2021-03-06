package domain

// Code generated by chromebot-domain-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cast"
)

// Cast executes a cdproto command under Cast domain.
type Cast struct {
	ctxWithExecutor context.Context
}

// Enable starts observing for sinks that can be used for tab mirroring, and
// if set, sinks compatible with |presentationUrl| as well. When sinks are
// found, a |sinksUpdated| event is fired. Also starts observing for issue
// messages. When an issue is added or removed, an |issueUpdated| event is
// fired.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Cast#method-enable
//
// parameters:
//  - `presentationURL`
func (doCast Cast) Enable(presentationURL *string) (err error) {
	b := cast.Enable()
	if presentationURL != nil {
		b = b.WithPresentationURL(*presentationURL)
	}
	return b.Do(doCast.ctxWithExecutor)
}

// Disable stops observing for sinks and issues.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Cast#method-disable
func (doCast Cast) Disable() (err error) {
	b := cast.Disable()
	return b.Do(doCast.ctxWithExecutor)
}

// SetSinkToUse sets a sink to be used when the web page requests the browser
// to choose a sink via Presentation API, Remote Playback API, or Cast SDK.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Cast#method-setSinkToUse
//
// parameters:
//  - `sinkName`
func (doCast Cast) SetSinkToUse(sinkName string) (err error) {
	b := cast.SetSinkToUse(sinkName)
	return b.Do(doCast.ctxWithExecutor)
}

// StartTabMirroring starts mirroring the tab to the sink.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Cast#method-startTabMirroring
//
// parameters:
//  - `sinkName`
func (doCast Cast) StartTabMirroring(sinkName string) (err error) {
	b := cast.StartTabMirroring(sinkName)
	return b.Do(doCast.ctxWithExecutor)
}

// StopCasting stops the active Cast session on the sink.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Cast#method-stopCasting
//
// parameters:
//  - `sinkName`
func (doCast Cast) StopCasting(sinkName string) (err error) {
	b := cast.StopCasting(sinkName)
	return b.Do(doCast.ctxWithExecutor)
}
