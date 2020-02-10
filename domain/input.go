package domain

// Code generated by chromebot-domain-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/input"
)

// Input executes a cdproto command under Input domain.
type Input struct {
	ctxWithExecutor context.Context
}

// DispatchKeyEvent dispatches a key event to the page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#method-dispatchKeyEvent
//
// parameters:
//  - `type`: Type of the key event.
//  - `modifiers`: This can be nil. (Optional) Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
//  - `timestamp`: This can be nil. (Optional) Time at which the event occurred.
//  - `text`: This can be nil. (Optional) Text as generated by processing a virtual key code with a keyboard layout. Not needed for for keyUp and rawKeyDown events (default: "")
//  - `unmodifiedText`: This can be nil. (Optional) Text that would have been generated by the keyboard if no modifiers were pressed (except for shift). Useful for shortcut (accelerator) key handling (default: "").
//  - `keyIdentifier`: This can be nil. (Optional) Unique key identifier (e.g., 'U+0041') (default: "").
//  - `code`: This can be nil. (Optional) Unique DOM defined string value for each physical key (e.g., 'KeyA') (default: "").
//  - `key`: This can be nil. (Optional) Unique DOM defined string value describing the meaning of the key in the context of active modifiers, keyboard layout, etc (e.g., 'AltGr') (default: "").
//  - `windowsVirtualKeyCode`: This can be nil. (Optional) Windows virtual key code (default: 0).
//  - `nativeVirtualKeyCode`: This can be nil. (Optional) Native virtual key code (default: 0).
//  - `autoRepeat`: This can be nil. (Optional) Whether the event was generated from auto repeat (default: false).
//  - `isKeypad`: This can be nil. (Optional) Whether the event was generated from the keypad (default: false).
//  - `isSystemKey`: This can be nil. (Optional) Whether the event was a system key event (default: false).
//  - `location`: This can be nil. (Optional) Whether the event was from the left or right side of the keyboard. 1=Left, 2=Right (default: 0).
func (doInput Input) DispatchKeyEvent(typeVal input.KeyType, modifiers *input.Modifier, timestamp *input.TimeSinceEpoch, text *string, unmodifiedText *string, keyIdentifier *string, code *string, key *string, windowsVirtualKeyCode *int64, nativeVirtualKeyCode *int64, autoRepeat *bool, isKeypad *bool, isSystemKey *bool, location *int64) (err error) {
	b := input.DispatchKeyEvent(typeVal)
	if modifiers != nil {
		b = b.WithModifiers(*modifiers)
	}
	if timestamp != nil {
		b = b.WithTimestamp(timestamp)
	}
	if text != nil {
		b = b.WithText(*text)
	}
	if unmodifiedText != nil {
		b = b.WithUnmodifiedText(*unmodifiedText)
	}
	if keyIdentifier != nil {
		b = b.WithKeyIdentifier(*keyIdentifier)
	}
	if code != nil {
		b = b.WithCode(*code)
	}
	if key != nil {
		b = b.WithKey(*key)
	}
	if windowsVirtualKeyCode != nil {
		b = b.WithWindowsVirtualKeyCode(*windowsVirtualKeyCode)
	}
	if nativeVirtualKeyCode != nil {
		b = b.WithNativeVirtualKeyCode(*nativeVirtualKeyCode)
	}
	if autoRepeat != nil {
		b = b.WithAutoRepeat(*autoRepeat)
	}
	if isKeypad != nil {
		b = b.WithIsKeypad(*isKeypad)
	}
	if isSystemKey != nil {
		b = b.WithIsSystemKey(*isSystemKey)
	}
	if location != nil {
		b = b.WithLocation(*location)
	}
	return b.Do(doInput.ctxWithExecutor)
}

// InsertText this method emulates inserting text that doesn't come from a
// key press, for example an emoji keyboard or an IME.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#method-insertText
//
// parameters:
//  - `text`: The text to insert.
func (doInput Input) InsertText(text string) (err error) {
	b := input.InsertText(text)
	return b.Do(doInput.ctxWithExecutor)
}

// DispatchMouseEvent dispatches a mouse event to the page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#method-dispatchMouseEvent
//
// parameters:
//  - `type`: Type of the mouse event.
//  - `x`: X coordinate of the event relative to the main frame's viewport in CSS pixels.
//  - `y`: Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
//  - `modifiers`: This can be nil. (Optional) Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
//  - `timestamp`: This can be nil. (Optional) Time at which the event occurred.
//  - `button`: This can be nil. (Optional) Mouse button (default: "none").
//  - `buttons`: This can be nil. (Optional) A number indicating which buttons are pressed on the mouse when a mouse event is triggered. Left=1, Right=2, Middle=4, Back=8, Forward=16, None=0.
//  - `clickCount`: This can be nil. (Optional) Number of times the mouse button was clicked (default: 0).
//  - `deltaX`: This can be nil. (Optional) X delta in CSS pixels for mouse wheel event (default: 0).
//  - `deltaY`: This can be nil. (Optional) Y delta in CSS pixels for mouse wheel event (default: 0).
//  - `pointerType`: This can be nil. (Optional) Pointer type (default: "mouse").
func (doInput Input) DispatchMouseEvent(typeVal input.MouseType, x float64, y float64, modifiers *input.Modifier, timestamp *input.TimeSinceEpoch, button *input.MouseButton, buttons *int64, clickCount *int64, deltaX *float64, deltaY *float64, pointerType *input.DispatchMouseEventPointerType) (err error) {
	b := input.DispatchMouseEvent(typeVal, x, y)
	if modifiers != nil {
		b = b.WithModifiers(*modifiers)
	}
	if timestamp != nil {
		b = b.WithTimestamp(timestamp)
	}
	if button != nil {
		b = b.WithButton(*button)
	}
	if buttons != nil {
		b = b.WithButtons(*buttons)
	}
	if clickCount != nil {
		b = b.WithClickCount(*clickCount)
	}
	if deltaX != nil {
		b = b.WithDeltaX(*deltaX)
	}
	if deltaY != nil {
		b = b.WithDeltaY(*deltaY)
	}
	if pointerType != nil {
		b = b.WithPointerType(*pointerType)
	}
	return b.Do(doInput.ctxWithExecutor)
}

// DispatchTouchEvent dispatches a touch event to the page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#method-dispatchTouchEvent
//
// parameters:
//  - `type`: Type of the touch event. TouchEnd and TouchCancel must not contain any touch points, while TouchStart and TouchMove must contains at least one.
//  - `touchPoints`: Active touch points on the touch device. One event per any changed point (compared to previous touch event in a sequence) is generated, emulating pressing/moving/releasing points one by one.
//  - `modifiers`: This can be nil. (Optional) Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
//  - `timestamp`: This can be nil. (Optional) Time at which the event occurred.
func (doInput Input) DispatchTouchEvent(typeVal input.TouchType, touchPoints []*input.TouchPoint, modifiers *input.Modifier, timestamp *input.TimeSinceEpoch) (err error) {
	b := input.DispatchTouchEvent(typeVal, touchPoints)
	if modifiers != nil {
		b = b.WithModifiers(*modifiers)
	}
	if timestamp != nil {
		b = b.WithTimestamp(timestamp)
	}
	return b.Do(doInput.ctxWithExecutor)
}

// EmulateTouchFromMouseEvent emulates touch event from the mouse event
// parameters.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#method-emulateTouchFromMouseEvent
//
// parameters:
//  - `type`: Type of the mouse event.
//  - `x`: X coordinate of the mouse pointer in DIP.
//  - `y`: Y coordinate of the mouse pointer in DIP.
//  - `button`: Mouse button. Only "none", "left", "right" are supported.
//  - `timestamp`: This can be nil. (Optional) Time at which the event occurred (default: current time).
//  - `deltaX`: This can be nil. (Optional) X delta in DIP for mouse wheel event (default: 0).
//  - `deltaY`: This can be nil. (Optional) Y delta in DIP for mouse wheel event (default: 0).
//  - `modifiers`: This can be nil. (Optional) Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
//  - `clickCount`: This can be nil. (Optional) Number of times the mouse button was clicked (default: 0).
func (doInput Input) EmulateTouchFromMouseEvent(typeVal input.MouseType, x int64, y int64, button input.MouseButton, timestamp *input.TimeSinceEpoch, deltaX *float64, deltaY *float64, modifiers *input.Modifier, clickCount *int64) (err error) {
	b := input.EmulateTouchFromMouseEvent(typeVal, x, y, button)
	if timestamp != nil {
		b = b.WithTimestamp(timestamp)
	}
	if deltaX != nil {
		b = b.WithDeltaX(*deltaX)
	}
	if deltaY != nil {
		b = b.WithDeltaY(*deltaY)
	}
	if modifiers != nil {
		b = b.WithModifiers(*modifiers)
	}
	if clickCount != nil {
		b = b.WithClickCount(*clickCount)
	}
	return b.Do(doInput.ctxWithExecutor)
}

// SetIgnoreInputEvents ignores input events (useful while auditing page).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#method-setIgnoreInputEvents
//
// parameters:
//  - `ignore`: Ignores input events processing when set to true.
func (doInput Input) SetIgnoreInputEvents(ignore bool) (err error) {
	b := input.SetIgnoreInputEvents(ignore)
	return b.Do(doInput.ctxWithExecutor)
}

// SynthesizePinchGesture synthesizes a pinch gesture over a time period by
// issuing appropriate touch events.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#method-synthesizePinchGesture
//
// parameters:
//  - `x`: X coordinate of the start of the gesture in CSS pixels.
//  - `y`: Y coordinate of the start of the gesture in CSS pixels.
//  - `scaleFactor`: Relative scale factor after zooming (>1.0 zooms in, <1.0 zooms out).
//  - `relativeSpeed`: This can be nil. (Optional) Relative pointer speed in pixels per second (default: 800).
//  - `gestureSourceType`: This can be nil. (Optional) Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
func (doInput Input) SynthesizePinchGesture(x float64, y float64, scaleFactor float64, relativeSpeed *int64, gestureSourceType *input.GestureSourceType) (err error) {
	b := input.SynthesizePinchGesture(x, y, scaleFactor)
	if relativeSpeed != nil {
		b = b.WithRelativeSpeed(*relativeSpeed)
	}
	if gestureSourceType != nil {
		b = b.WithGestureSourceType(*gestureSourceType)
	}
	return b.Do(doInput.ctxWithExecutor)
}

// SynthesizeScrollGesture synthesizes a scroll gesture over a time period by
// issuing appropriate touch events.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#method-synthesizeScrollGesture
//
// parameters:
//  - `x`: X coordinate of the start of the gesture in CSS pixels.
//  - `y`: Y coordinate of the start of the gesture in CSS pixels.
//  - `xDistance`: This can be nil. (Optional) The distance to scroll along the X axis (positive to scroll left).
//  - `yDistance`: This can be nil. (Optional) The distance to scroll along the Y axis (positive to scroll up).
//  - `xOverscroll`: This can be nil. (Optional) The number of additional pixels to scroll back along the X axis, in addition to the given distance.
//  - `yOverscroll`: This can be nil. (Optional) The number of additional pixels to scroll back along the Y axis, in addition to the given distance.
//  - `preventFling`: This can be nil. (Optional) Prevent fling (default: true).
//  - `speed`: This can be nil. (Optional) Swipe speed in pixels per second (default: 800).
//  - `gestureSourceType`: This can be nil. (Optional) Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
//  - `repeatCount`: This can be nil. (Optional) The number of times to repeat the gesture (default: 0).
//  - `repeatDelayMs`: This can be nil. (Optional) The number of milliseconds delay between each repeat. (default: 250).
//  - `interactionMarkerName`: This can be nil. (Optional) The name of the interaction markers to generate, if not empty (default: "").
func (doInput Input) SynthesizeScrollGesture(x float64, y float64, xDistance *float64, yDistance *float64, xOverscroll *float64, yOverscroll *float64, preventFling *bool, speed *int64, gestureSourceType *input.GestureSourceType, repeatCount *int64, repeatDelayMs *int64, interactionMarkerName *string) (err error) {
	b := input.SynthesizeScrollGesture(x, y)
	if xDistance != nil {
		b = b.WithXDistance(*xDistance)
	}
	if yDistance != nil {
		b = b.WithYDistance(*yDistance)
	}
	if xOverscroll != nil {
		b = b.WithXOverscroll(*xOverscroll)
	}
	if yOverscroll != nil {
		b = b.WithYOverscroll(*yOverscroll)
	}
	if preventFling != nil {
		b = b.WithPreventFling(*preventFling)
	}
	if speed != nil {
		b = b.WithSpeed(*speed)
	}
	if gestureSourceType != nil {
		b = b.WithGestureSourceType(*gestureSourceType)
	}
	if repeatCount != nil {
		b = b.WithRepeatCount(*repeatCount)
	}
	if repeatDelayMs != nil {
		b = b.WithRepeatDelayMs(*repeatDelayMs)
	}
	if interactionMarkerName != nil {
		b = b.WithInteractionMarkerName(*interactionMarkerName)
	}
	return b.Do(doInput.ctxWithExecutor)
}

// SynthesizeTapGesture synthesizes a tap gesture over a time period by
// issuing appropriate touch events.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#method-synthesizeTapGesture
//
// parameters:
//  - `x`: X coordinate of the start of the gesture in CSS pixels.
//  - `y`: Y coordinate of the start of the gesture in CSS pixels.
//  - `duration`: This can be nil. (Optional) Duration between touchdown and touchup events in ms (default: 50).
//  - `tapCount`: This can be nil. (Optional) Number of times to perform the tap (e.g. 2 for double tap, default: 1).
//  - `gestureSourceType`: This can be nil. (Optional) Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
func (doInput Input) SynthesizeTapGesture(x float64, y float64, duration *int64, tapCount *int64, gestureSourceType *input.GestureSourceType) (err error) {
	b := input.SynthesizeTapGesture(x, y)
	if duration != nil {
		b = b.WithDuration(*duration)
	}
	if tapCount != nil {
		b = b.WithTapCount(*tapCount)
	}
	if gestureSourceType != nil {
		b = b.WithGestureSourceType(*gestureSourceType)
	}
	return b.Do(doInput.ctxWithExecutor)
}
