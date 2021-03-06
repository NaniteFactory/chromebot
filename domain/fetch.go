package domain

// Code generated by chromebot-domain-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/fetch"
	"github.com/chromedp/cdproto/io"
	"github.com/chromedp/cdproto/network"
)

// Fetch executes a cdproto command under Fetch domain.
type Fetch struct {
	ctxWithExecutor context.Context
}

// Disable disables the fetch domain.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-disable
func (doFetch Fetch) Disable() (err error) {
	b := fetch.Disable()
	return b.Do(doFetch.ctxWithExecutor)
}

// Enable enables issuing of requestPaused events. A request will be paused
// until client calls one of failRequest, fulfillRequest or
// continueRequest/continueWithAuth.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-enable
//
// parameters:
//  - `patterns`: This can be nil. (Optional) If specified, only requests matching any of these patterns will produce fetchRequested event and will be paused until clients response. If not set, all requests will be affected.
//  - `handleAuthRequests`: This can be nil. (Optional) If true, authRequired events will be issued and requests will be paused expecting a call to continueWithAuth.
func (doFetch Fetch) Enable(patterns []*fetch.RequestPattern, handleAuthRequests *bool) (err error) {
	b := fetch.Enable()
	if patterns != nil {
		b = b.WithPatterns(patterns)
	}
	if handleAuthRequests != nil {
		b = b.WithHandleAuthRequests(*handleAuthRequests)
	}
	return b.Do(doFetch.ctxWithExecutor)
}

// FailRequest causes the request to fail with specified reason.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-failRequest
//
// parameters:
//  - `requestID`: An id the client received in requestPaused event.
//  - `errorReason`: Causes the request to fail with the given reason.
func (doFetch Fetch) FailRequest(requestID fetch.RequestID, errorReason network.ErrorReason) (err error) {
	b := fetch.FailRequest(requestID, errorReason)
	return b.Do(doFetch.ctxWithExecutor)
}

// FulfillRequest provides response to the request.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-fulfillRequest
//
// parameters:
//  - `requestID`: An id the client received in requestPaused event.
//  - `responseCode`: An HTTP response code.
//  - `responseHeaders`: This can be nil. (Optional) Response headers.
//  - `binaryResponseHeaders`: This can be nil. (Optional) Alternative way of specifying response headers as a \0-separated series of name: value pairs. Prefer the above method unless you need to represent some non-UTF8 values that can't be transmitted over the protocol as text.
//  - `body`: This can be nil. (Optional) A response body.
//  - `responsePhrase`: This can be nil. (Optional) A textual representation of responseCode. If absent, a standard phrase matching responseCode is used.
func (doFetch Fetch) FulfillRequest(requestID fetch.RequestID, responseCode int64, responseHeaders []*fetch.HeaderEntry, binaryResponseHeaders *string, body *string, responsePhrase *string) (err error) {
	b := fetch.FulfillRequest(requestID, responseCode)
	if responseHeaders != nil {
		b = b.WithResponseHeaders(responseHeaders)
	}
	if binaryResponseHeaders != nil {
		b = b.WithBinaryResponseHeaders(*binaryResponseHeaders)
	}
	if body != nil {
		b = b.WithBody(*body)
	}
	if responsePhrase != nil {
		b = b.WithResponsePhrase(*responsePhrase)
	}
	return b.Do(doFetch.ctxWithExecutor)
}

// ContinueRequest continues the request, optionally modifying some of its
// parameters.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-continueRequest
//
// parameters:
//  - `requestID`: An id the client received in requestPaused event.
//  - `url`: This can be nil. (Optional) If set, the request url will be modified in a way that's not observable by page.
//  - `method`: This can be nil. (Optional) If set, the request method is overridden.
//  - `postData`: This can be nil. (Optional) If set, overrides the post data in the request.
//  - `headers`: This can be nil. (Optional) If set, overrides the request headrts.
func (doFetch Fetch) ContinueRequest(requestID fetch.RequestID, url *string, method *string, postData *string, headers []*fetch.HeaderEntry) (err error) {
	b := fetch.ContinueRequest(requestID)
	if url != nil {
		b = b.WithURL(*url)
	}
	if method != nil {
		b = b.WithMethod(*method)
	}
	if postData != nil {
		b = b.WithPostData(*postData)
	}
	if headers != nil {
		b = b.WithHeaders(headers)
	}
	return b.Do(doFetch.ctxWithExecutor)
}

// ContinueWithAuth continues a request supplying authChallengeResponse
// following authRequired event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-continueWithAuth
//
// parameters:
//  - `requestID`: An id the client received in authRequired event.
//  - `authChallengeResponse`: Response to  with an authChallenge.
func (doFetch Fetch) ContinueWithAuth(requestID fetch.RequestID, authChallengeResponse *fetch.AuthChallengeResponse) (err error) {
	b := fetch.ContinueWithAuth(requestID, authChallengeResponse)
	return b.Do(doFetch.ctxWithExecutor)
}

// GetResponseBody causes the body of the response to be received from the
// server and returned as a single string. May only be issued for a request that
// is paused in the Response stage and is mutually exclusive with
// takeResponseBodyForInterceptionAsStream. Calling other methods that affect
// the request or disabling fetch domain before body is received results in an
// undefined behavior.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-getResponseBody
//
// parameters:
//  - `requestID`: Identifier for the intercepted request to get body for.
//
// returns:
//  - `retBody`: Response body.
func (doFetch Fetch) GetResponseBody(requestID fetch.RequestID) (retBody []byte, err error) {
	b := fetch.GetResponseBody(requestID)
	return b.Do(doFetch.ctxWithExecutor)
}

// TakeResponseBodyAsStream returns a handle to the stream representing the
// response body. The request must be paused in the HeadersReceived stage. Note
// that after this command the request can't be continued as is -- client either
// needs to cancel it or to provide the response body. The stream only supports
// sequential read, IO.read will fail if the position is specified. This method
// is mutually exclusive with getResponseBody. Calling other methods that affect
// the request or disabling fetch domain before body is received results in an
// undefined behavior.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-takeResponseBodyAsStream
//
// parameters:
//  - `requestID`
//
// returns:
//  - `retStream`
func (doFetch Fetch) TakeResponseBodyAsStream(requestID fetch.RequestID) (retStream io.StreamHandle, err error) {
	b := fetch.TakeResponseBodyAsStream(requestID)
	return b.Do(doFetch.ctxWithExecutor)
}
