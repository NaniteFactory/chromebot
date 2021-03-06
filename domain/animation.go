package domain

// Code generated by chromebot-domain-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/animation"
	"github.com/chromedp/cdproto/runtime"
)

// Animation executes a cdproto command under Animation domain.
type Animation struct {
	ctxWithExecutor context.Context
}

// Disable disables animation domain notifications.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Animation#method-disable
func (doAnimation Animation) Disable() (err error) {
	b := animation.Disable()
	return b.Do(doAnimation.ctxWithExecutor)
}

// Enable enables animation domain notifications.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Animation#method-enable
func (doAnimation Animation) Enable() (err error) {
	b := animation.Enable()
	return b.Do(doAnimation.ctxWithExecutor)
}

// GetCurrentTime returns the current time of the an animation.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Animation#method-getCurrentTime
//
// parameters:
//  - `id`: Id of animation.
//
// returns:
//  - `retCurrentTime`: Current time of the page.
func (doAnimation Animation) GetCurrentTime(id string) (retCurrentTime float64, err error) {
	b := animation.GetCurrentTime(id)
	return b.Do(doAnimation.ctxWithExecutor)
}

// GetPlaybackRate gets the playback rate of the document timeline.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Animation#method-getPlaybackRate
//
// returns:
//  - `retPlaybackRate`: Playback rate for animations on page.
func (doAnimation Animation) GetPlaybackRate() (retPlaybackRate float64, err error) {
	b := animation.GetPlaybackRate()
	return b.Do(doAnimation.ctxWithExecutor)
}

// ReleaseAnimations releases a set of animations to no longer be
// manipulated.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Animation#method-releaseAnimations
//
// parameters:
//  - `animations`: List of animation ids to seek.
func (doAnimation Animation) ReleaseAnimations(animations []string) (err error) {
	b := animation.ReleaseAnimations(animations)
	return b.Do(doAnimation.ctxWithExecutor)
}

// ResolveAnimation gets the remote object of the Animation.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Animation#method-resolveAnimation
//
// parameters:
//  - `animationID`: Animation id.
//
// returns:
//  - `retRemoteObject`: Corresponding remote object.
func (doAnimation Animation) ResolveAnimation(animationID string) (retRemoteObject *runtime.RemoteObject, err error) {
	b := animation.ResolveAnimation(animationID)
	return b.Do(doAnimation.ctxWithExecutor)
}

// SeekAnimations seek a set of animations to a particular time within each
// animation.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Animation#method-seekAnimations
//
// parameters:
//  - `animations`: List of animation ids to seek.
//  - `currentTime`: Set the current time of each animation.
func (doAnimation Animation) SeekAnimations(animations []string, currentTime float64) (err error) {
	b := animation.SeekAnimations(animations, currentTime)
	return b.Do(doAnimation.ctxWithExecutor)
}

// SetPaused sets the paused state of a set of animations.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Animation#method-setPaused
//
// parameters:
//  - `animations`: Animations to set the pause state of.
//  - `paused`: Paused state to set to.
func (doAnimation Animation) SetPaused(animations []string, paused bool) (err error) {
	b := animation.SetPaused(animations, paused)
	return b.Do(doAnimation.ctxWithExecutor)
}

// SetPlaybackRate sets the playback rate of the document timeline.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Animation#method-setPlaybackRate
//
// parameters:
//  - `playbackRate`: Playback rate for animations on page
func (doAnimation Animation) SetPlaybackRate(playbackRate float64) (err error) {
	b := animation.SetPlaybackRate(playbackRate)
	return b.Do(doAnimation.ctxWithExecutor)
}

// SetTiming sets the timing of an animation node.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Animation#method-setTiming
//
// parameters:
//  - `animationID`: Animation id.
//  - `duration`: Duration of the animation.
//  - `delay`: Delay of the animation.
func (doAnimation Animation) SetTiming(animationID string, duration float64, delay float64) (err error) {
	b := animation.SetTiming(animationID, duration, delay)
	return b.Do(doAnimation.ctxWithExecutor)
}
