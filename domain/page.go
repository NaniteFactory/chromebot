package domain

// Code generated by chromebot-domain-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/debugger"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/io"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/runtime"
)

// Page executes a cdproto command under Page domain.
type Page struct {
	ctxWithExecutor context.Context
}

// AddScriptToEvaluateOnNewDocument evaluates given script in every frame
// upon creation (before loading frame's scripts).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-addScriptToEvaluateOnNewDocument
//
// parameters:
//  - `source`
//  - `worldName`: This can be nil. (Optional) If specified, creates an isolated world with the given name and evaluates given script in it. This world name will be used as the ExecutionContextDescription::name when the corresponding event is emitted.
//
// returns:
//  - `retIdentifier`: Identifier of the added script.
func (doPage Page) AddScriptToEvaluateOnNewDocument(source string, worldName *string) (retIdentifier page.ScriptIdentifier, err error) {
	b := page.AddScriptToEvaluateOnNewDocument(source)
	if worldName != nil {
		b = b.WithWorldName(*worldName)
	}
	return b.Do(doPage.ctxWithExecutor)
}

// BringToFront brings page to front (activates tab).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-bringToFront
func (doPage Page) BringToFront() (err error) {
	b := page.BringToFront()
	return b.Do(doPage.ctxWithExecutor)
}

// CaptureScreenshot capture page screenshot.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-captureScreenshot
//
// parameters:
//  - `format`: This can be nil. (Optional) Image compression format (defaults to png).
//  - `quality`: This can be nil. (Optional) Compression quality from range [0..100] (jpeg only).
//  - `clip`: This can be nil. (Optional) Capture the screenshot of a given region only.
//  - `fromSurface`: This can be nil. (Optional) Capture the screenshot from the surface, rather than the view. Defaults to true.
//
// returns:
//  - `retData`: Base64-encoded image data.
func (doPage Page) CaptureScreenshot(format *page.CaptureScreenshotFormat, quality *int64, clip *page.Viewport, fromSurface *bool) (retData []byte, err error) {
	b := page.CaptureScreenshot()
	if format != nil {
		b = b.WithFormat(*format)
	}
	if quality != nil {
		b = b.WithQuality(*quality)
	}
	if clip != nil {
		b = b.WithClip(clip)
	}
	if fromSurface != nil {
		b = b.WithFromSurface(*fromSurface)
	}
	return b.Do(doPage.ctxWithExecutor)
}

// CaptureSnapshot returns a snapshot of the page as a string. For MHTML
// format, the serialization includes iframes, shadow DOM, external resources,
// and element-inline styles.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-captureSnapshot
//
// parameters:
//  - `format`: This can be nil. (Optional) Format (defaults to mhtml).
//
// returns:
//  - `retData`: Serialized page data.
func (doPage Page) CaptureSnapshot(format *page.CaptureSnapshotFormat) (retData string, err error) {
	b := page.CaptureSnapshot()
	if format != nil {
		b = b.WithFormat(*format)
	}
	return b.Do(doPage.ctxWithExecutor)
}

// CreateIsolatedWorld creates an isolated world for the given frame.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-createIsolatedWorld
//
// parameters:
//  - `frameID`: Id of the frame in which the isolated world should be created.
//  - `worldName`: This can be nil. (Optional) An optional name which is reported in the Execution Context.
//  - `grantUniveralAccess`: This can be nil. (Optional) Whether or not universal access should be granted to the isolated world. This is a powerful option, use with caution.
//
// returns:
//  - `retExecutionContextID`: Execution context of the isolated world.
func (doPage Page) CreateIsolatedWorld(frameID cdp.FrameID, worldName *string, grantUniveralAccess *bool) (retExecutionContextID runtime.ExecutionContextID, err error) {
	b := page.CreateIsolatedWorld(frameID)
	if worldName != nil {
		b = b.WithWorldName(*worldName)
	}
	if grantUniveralAccess != nil {
		b = b.WithGrantUniveralAccess(*grantUniveralAccess)
	}
	return b.Do(doPage.ctxWithExecutor)
}

// Disable disables page domain notifications.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-disable
func (doPage Page) Disable() (err error) {
	b := page.Disable()
	return b.Do(doPage.ctxWithExecutor)
}

// Enable enables page domain notifications.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-enable
func (doPage Page) Enable() (err error) {
	b := page.Enable()
	return b.Do(doPage.ctxWithExecutor)
}

// GetAppManifest [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-getAppManifest
//
// returns:
//  - `retURL`: Manifest location.
//  - `retErrors`
//  - `retData`: Manifest content.
//  - `retParsed`: Parsed manifest properties
func (doPage Page) GetAppManifest() (retURL string, retErrors []*page.AppManifestError, retData string, retParsed *page.AppManifestParsedProperties, err error) {
	b := page.GetAppManifest()
	return b.Do(doPage.ctxWithExecutor)
}

// GetInstallabilityErrors [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-getInstallabilityErrors
//
// returns:
//  - `retInstallabilityErrors`
func (doPage Page) GetInstallabilityErrors() (retInstallabilityErrors []*page.InstallabilityError, err error) {
	b := page.GetInstallabilityErrors()
	return b.Do(doPage.ctxWithExecutor)
}

// GetManifestIcons [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-getManifestIcons
//
// returns:
//  - `retPrimaryIcon`
func (doPage Page) GetManifestIcons() (retPrimaryIcon []byte, err error) {
	b := page.GetManifestIcons()
	return b.Do(doPage.ctxWithExecutor)
}

// GetFrameTree returns present frame tree structure.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-getFrameTree
//
// returns:
//  - `retFrameTree`: Present frame tree structure.
func (doPage Page) GetFrameTree() (retFrameTree *page.FrameTree, err error) {
	b := page.GetFrameTree()
	return b.Do(doPage.ctxWithExecutor)
}

// GetLayoutMetrics returns metrics relating to the layouting of the page,
// such as viewport bounds/scale.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-getLayoutMetrics
//
// returns:
//  - `retLayoutViewport`: Metrics relating to the layout viewport.
//  - `retVisualViewport`: Metrics relating to the visual viewport.
//  - `retContentSize`: Size of scrollable area.
func (doPage Page) GetLayoutMetrics() (retLayoutViewport *page.LayoutViewport, retVisualViewport *page.VisualViewport, retContentSize *dom.Rect, err error) {
	b := page.GetLayoutMetrics()
	return b.Do(doPage.ctxWithExecutor)
}

// GetNavigationHistory returns navigation history for the current page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-getNavigationHistory
//
// returns:
//  - `retCurrentIndex`: Index of the current navigation history entry.
//  - `retEntries`: Array of navigation history entries.
func (doPage Page) GetNavigationHistory() (retCurrentIndex int64, retEntries []*page.NavigationEntry, err error) {
	b := page.GetNavigationHistory()
	return b.Do(doPage.ctxWithExecutor)
}

// ResetNavigationHistory resets navigation history for the current page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-resetNavigationHistory
func (doPage Page) ResetNavigationHistory() (err error) {
	b := page.ResetNavigationHistory()
	return b.Do(doPage.ctxWithExecutor)
}

// GetResourceContent returns content of the given resource.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-getResourceContent
//
// parameters:
//  - `frameID`: Frame id to get resource for.
//  - `url`: URL of the resource to get content for.
//
// returns:
//  - `retContent`: Resource content.
func (doPage Page) GetResourceContent(frameID cdp.FrameID, url string) (retContent []byte, err error) {
	b := page.GetResourceContent(frameID, url)
	return b.Do(doPage.ctxWithExecutor)
}

// GetResourceTree returns present frame / resource tree structure.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-getResourceTree
//
// returns:
//  - `retFrameTree`: Present frame / resource tree structure.
func (doPage Page) GetResourceTree() (retFrameTree *page.FrameResourceTree, err error) {
	b := page.GetResourceTree()
	return b.Do(doPage.ctxWithExecutor)
}

// HandleJavaScriptDialog accepts or dismisses a JavaScript initiated dialog
// (alert, confirm, prompt, or onbeforeunload).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-handleJavaScriptDialog
//
// parameters:
//  - `accept`: Whether to accept or dismiss the dialog.
//  - `promptText`: This can be nil. (Optional) The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog.
func (doPage Page) HandleJavaScriptDialog(accept bool, promptText *string) (err error) {
	b := page.HandleJavaScriptDialog(accept)
	if promptText != nil {
		b = b.WithPromptText(*promptText)
	}
	return b.Do(doPage.ctxWithExecutor)
}

// Navigate navigates current page to the given URL.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-navigate
//
// parameters:
//  - `url`: URL to navigate the page to.
//  - `referrer`: This can be nil. (Optional) Referrer URL.
//  - `transitionType`: This can be nil. (Optional) Intended transition type.
//  - `frameID`: This can be nil. (Optional) Frame id to navigate, if not specified navigates the top frame.
//
// returns:
//  - `retFrameID`: Frame id that has navigated (or failed to navigate)
//  - `retLoaderID`: Loader identifier.
//  - `retErrorText`: User friendly error message, present if and only if navigation has failed.
func (doPage Page) Navigate(url string, referrer *string, transitionType *page.TransitionType, frameID *cdp.FrameID) (retFrameID cdp.FrameID, retLoaderID cdp.LoaderID, retErrorText string, err error) {
	b := page.Navigate(url)
	if referrer != nil {
		b = b.WithReferrer(*referrer)
	}
	if transitionType != nil {
		b = b.WithTransitionType(*transitionType)
	}
	if frameID != nil {
		b = b.WithFrameID(*frameID)
	}
	return b.Do(doPage.ctxWithExecutor)
}

// NavigateToHistoryEntry navigates current page to the given history entry.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-navigateToHistoryEntry
//
// parameters:
//  - `entryID`: Unique id of the entry to navigate to.
func (doPage Page) NavigateToHistoryEntry(entryID int64) (err error) {
	b := page.NavigateToHistoryEntry(entryID)
	return b.Do(doPage.ctxWithExecutor)
}

// PrintToPDF print page as PDF.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-printToPDF
//
// parameters:
//  - `landscape`: This can be nil. (Optional) Paper orientation. Defaults to false.
//  - `displayHeaderFooter`: This can be nil. (Optional) Display header and footer. Defaults to false.
//  - `printBackground`: This can be nil. (Optional) Print background graphics. Defaults to false.
//  - `scale`: This can be nil. (Optional) Scale of the webpage rendering. Defaults to 1.
//  - `paperWidth`: This can be nil. (Optional) Paper width in inches. Defaults to 8.5 inches.
//  - `paperHeight`: This can be nil. (Optional) Paper height in inches. Defaults to 11 inches.
//  - `marginTop`: This can be nil. (Optional) Top margin in inches. Defaults to 1cm (~0.4 inches).
//  - `marginBottom`: This can be nil. (Optional) Bottom margin in inches. Defaults to 1cm (~0.4 inches).
//  - `marginLeft`: This can be nil. (Optional) Left margin in inches. Defaults to 1cm (~0.4 inches).
//  - `marginRight`: This can be nil. (Optional) Right margin in inches. Defaults to 1cm (~0.4 inches).
//  - `pageRanges`: This can be nil. (Optional) Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means print all pages.
//  - `ignoreInvalidPageRanges`: This can be nil. (Optional) Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'. Defaults to false.
//  - `headerTemplate`: This can be nil. (Optional) HTML template for the print header. Should be valid HTML markup with following classes used to inject printing values into them: - date: formatted print date - title: document title - url: document location - pageNumber: current page number - totalPages: total pages in the document  For example, <span class=title></span> would generate span containing the title.
//  - `footerTemplate`: This can be nil. (Optional) HTML template for the print footer. Should use the same format as the headerTemplate.
//  - `preferCSSPageSize`: This can be nil. (Optional) Whether or not to prefer page size as defined by css. Defaults to false, in which case the content will be scaled to fit the paper size.
//  - `transferMode`: This can be nil. (Optional) return as stream
//
// returns:
//  - `retData`: Base64-encoded pdf data. Empty if |returnAsStream| is specified.
//  - `retStream`: A handle of the stream that holds resulting PDF data.
func (doPage Page) PrintToPDF(landscape *bool, displayHeaderFooter *bool, printBackground *bool, scale *float64, paperWidth *float64, paperHeight *float64, marginTop *float64, marginBottom *float64, marginLeft *float64, marginRight *float64, pageRanges *string, ignoreInvalidPageRanges *bool, headerTemplate *string, footerTemplate *string, preferCSSPageSize *bool, transferMode *page.PrintToPDFTransferMode) (retData []byte, retStream io.StreamHandle, err error) {
	b := page.PrintToPDF()
	if landscape != nil {
		b = b.WithLandscape(*landscape)
	}
	if displayHeaderFooter != nil {
		b = b.WithDisplayHeaderFooter(*displayHeaderFooter)
	}
	if printBackground != nil {
		b = b.WithPrintBackground(*printBackground)
	}
	if scale != nil {
		b = b.WithScale(*scale)
	}
	if paperWidth != nil {
		b = b.WithPaperWidth(*paperWidth)
	}
	if paperHeight != nil {
		b = b.WithPaperHeight(*paperHeight)
	}
	if marginTop != nil {
		b = b.WithMarginTop(*marginTop)
	}
	if marginBottom != nil {
		b = b.WithMarginBottom(*marginBottom)
	}
	if marginLeft != nil {
		b = b.WithMarginLeft(*marginLeft)
	}
	if marginRight != nil {
		b = b.WithMarginRight(*marginRight)
	}
	if pageRanges != nil {
		b = b.WithPageRanges(*pageRanges)
	}
	if ignoreInvalidPageRanges != nil {
		b = b.WithIgnoreInvalidPageRanges(*ignoreInvalidPageRanges)
	}
	if headerTemplate != nil {
		b = b.WithHeaderTemplate(*headerTemplate)
	}
	if footerTemplate != nil {
		b = b.WithFooterTemplate(*footerTemplate)
	}
	if preferCSSPageSize != nil {
		b = b.WithPreferCSSPageSize(*preferCSSPageSize)
	}
	if transferMode != nil {
		b = b.WithTransferMode(*transferMode)
	}
	return b.Do(doPage.ctxWithExecutor)
}

// Reload reloads given page optionally ignoring the cache.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-reload
//
// parameters:
//  - `ignoreCache`: This can be nil. (Optional) If true, browser cache is ignored (as if the user pressed Shift+refresh).
//  - `scriptToEvaluateOnLoad`: This can be nil. (Optional) If set, the script will be injected into all frames of the inspected page after reload. Argument will be ignored if reloading dataURL origin.
func (doPage Page) Reload(ignoreCache *bool, scriptToEvaluateOnLoad *string) (err error) {
	b := page.Reload()
	if ignoreCache != nil {
		b = b.WithIgnoreCache(*ignoreCache)
	}
	if scriptToEvaluateOnLoad != nil {
		b = b.WithScriptToEvaluateOnLoad(*scriptToEvaluateOnLoad)
	}
	return b.Do(doPage.ctxWithExecutor)
}

// RemoveScriptToEvaluateOnNewDocument removes given script from the list.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-removeScriptToEvaluateOnNewDocument
//
// parameters:
//  - `identifier`
func (doPage Page) RemoveScriptToEvaluateOnNewDocument(identifier page.ScriptIdentifier) (err error) {
	b := page.RemoveScriptToEvaluateOnNewDocument(identifier)
	return b.Do(doPage.ctxWithExecutor)
}

// ScreencastFrameAck acknowledges that a screencast frame has been received
// by the frontend.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-screencastFrameAck
//
// parameters:
//  - `sessionID`: Frame number.
func (doPage Page) ScreencastFrameAck(sessionID int64) (err error) {
	b := page.ScreencastFrameAck(sessionID)
	return b.Do(doPage.ctxWithExecutor)
}

// SearchInResource searches for given string in resource content.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-searchInResource
//
// parameters:
//  - `frameID`: Frame id for resource to search in.
//  - `url`: URL of the resource to search in.
//  - `query`: String to search for.
//  - `caseSensitive`: This can be nil. (Optional) If true, search is case sensitive.
//  - `isRegex`: This can be nil. (Optional) If true, treats string parameter as regex.
//
// returns:
//  - `retResult`: List of search matches.
func (doPage Page) SearchInResource(frameID cdp.FrameID, url string, query string, caseSensitive *bool, isRegex *bool) (retResult []*debugger.SearchMatch, err error) {
	b := page.SearchInResource(frameID, url, query)
	if caseSensitive != nil {
		b = b.WithCaseSensitive(*caseSensitive)
	}
	if isRegex != nil {
		b = b.WithIsRegex(*isRegex)
	}
	return b.Do(doPage.ctxWithExecutor)
}

// SetAdBlockingEnabled enable Chrome's experimental ad filter on all sites.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-setAdBlockingEnabled
//
// parameters:
//  - `enabled`: Whether to block ads.
func (doPage Page) SetAdBlockingEnabled(enabled bool) (err error) {
	b := page.SetAdBlockingEnabled(enabled)
	return b.Do(doPage.ctxWithExecutor)
}

// SetBypassCSP enable page Content Security Policy by-passing.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-setBypassCSP
//
// parameters:
//  - `enabled`: Whether to bypass page CSP.
func (doPage Page) SetBypassCSP(enabled bool) (err error) {
	b := page.SetBypassCSP(enabled)
	return b.Do(doPage.ctxWithExecutor)
}

// SetFontFamilies set generic font families.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-setFontFamilies
//
// parameters:
//  - `fontFamilies`: Specifies font families to set. If a font family is not specified, it won't be changed.
func (doPage Page) SetFontFamilies(fontFamilies *page.FontFamilies) (err error) {
	b := page.SetFontFamilies(fontFamilies)
	return b.Do(doPage.ctxWithExecutor)
}

// SetFontSizes set default font sizes.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-setFontSizes
//
// parameters:
//  - `fontSizes`: Specifies font sizes to set. If a font size is not specified, it won't be changed.
func (doPage Page) SetFontSizes(fontSizes *page.FontSizes) (err error) {
	b := page.SetFontSizes(fontSizes)
	return b.Do(doPage.ctxWithExecutor)
}

// SetDocumentContent sets given markup as the document's HTML.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-setDocumentContent
//
// parameters:
//  - `frameID`: Frame id to set HTML for.
//  - `html`: HTML content to set.
func (doPage Page) SetDocumentContent(frameID cdp.FrameID, html string) (err error) {
	b := page.SetDocumentContent(frameID, html)
	return b.Do(doPage.ctxWithExecutor)
}

// SetDownloadBehavior set the behavior when downloading a file.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-setDownloadBehavior
//
// parameters:
//  - `behavior`: Whether to allow all or deny all download requests, or use default Chrome behavior if available (otherwise deny).
//  - `downloadPath`: This can be nil. (Optional) The default path to save downloaded files to. This is required if behavior is set to 'allow'
func (doPage Page) SetDownloadBehavior(behavior page.SetDownloadBehaviorBehavior, downloadPath *string) (err error) {
	b := page.SetDownloadBehavior(behavior)
	if downloadPath != nil {
		b = b.WithDownloadPath(*downloadPath)
	}
	return b.Do(doPage.ctxWithExecutor)
}

// SetLifecycleEventsEnabled controls whether page will emit lifecycle
// events.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-setLifecycleEventsEnabled
//
// parameters:
//  - `enabled`: If true, starts emitting lifecycle events.
func (doPage Page) SetLifecycleEventsEnabled(enabled bool) (err error) {
	b := page.SetLifecycleEventsEnabled(enabled)
	return b.Do(doPage.ctxWithExecutor)
}

// StartScreencast starts sending each frame using the screencastFrame event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-startScreencast
//
// parameters:
//  - `format`: This can be nil. (Optional) Image compression format.
//  - `quality`: This can be nil. (Optional) Compression quality from range [0..100].
//  - `maxWidth`: This can be nil. (Optional) Maximum screenshot width.
//  - `maxHeight`: This can be nil. (Optional) Maximum screenshot height.
//  - `everyNthFrame`: This can be nil. (Optional) Send every n-th frame.
func (doPage Page) StartScreencast(format *page.ScreencastFormat, quality *int64, maxWidth *int64, maxHeight *int64, everyNthFrame *int64) (err error) {
	b := page.StartScreencast()
	if format != nil {
		b = b.WithFormat(*format)
	}
	if quality != nil {
		b = b.WithQuality(*quality)
	}
	if maxWidth != nil {
		b = b.WithMaxWidth(*maxWidth)
	}
	if maxHeight != nil {
		b = b.WithMaxHeight(*maxHeight)
	}
	if everyNthFrame != nil {
		b = b.WithEveryNthFrame(*everyNthFrame)
	}
	return b.Do(doPage.ctxWithExecutor)
}

// StopLoading force the page stop all navigations and pending resource
// fetches.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-stopLoading
func (doPage Page) StopLoading() (err error) {
	b := page.StopLoading()
	return b.Do(doPage.ctxWithExecutor)
}

// Crash crashes renderer on the IO thread, generates minidumps.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-crash
func (doPage Page) Crash() (err error) {
	b := page.Crash()
	return b.Do(doPage.ctxWithExecutor)
}

// Close tries to close page, running its beforeunload hooks, if any.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-close
func (doPage Page) Close() (err error) {
	b := page.Close()
	return b.Do(doPage.ctxWithExecutor)
}

// SetWebLifecycleState tries to update the web lifecycle state of the page.
// It will transition the page to the given state according to:
// https://github.com/WICG/web-lifecycle/.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-setWebLifecycleState
//
// parameters:
//  - `state`: Target lifecycle state
func (doPage Page) SetWebLifecycleState(state page.SetWebLifecycleStateState) (err error) {
	b := page.SetWebLifecycleState(state)
	return b.Do(doPage.ctxWithExecutor)
}

// StopScreencast stops sending each frame in the screencastFrame.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-stopScreencast
func (doPage Page) StopScreencast() (err error) {
	b := page.StopScreencast()
	return b.Do(doPage.ctxWithExecutor)
}

// SetProduceCompilationCache forces compilation cache to be generated for
// every subresource script.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-setProduceCompilationCache
//
// parameters:
//  - `enabled`
func (doPage Page) SetProduceCompilationCache(enabled bool) (err error) {
	b := page.SetProduceCompilationCache(enabled)
	return b.Do(doPage.ctxWithExecutor)
}

// AddCompilationCache seeds compilation cache for given url. Compilation
// cache does not survive cross-process navigation.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-addCompilationCache
//
// parameters:
//  - `url`
//  - `data`: Base64-encoded data
func (doPage Page) AddCompilationCache(url string, data string) (err error) {
	b := page.AddCompilationCache(url, data)
	return b.Do(doPage.ctxWithExecutor)
}

// ClearCompilationCache clears seeded compilation cache.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-clearCompilationCache
func (doPage Page) ClearCompilationCache() (err error) {
	b := page.ClearCompilationCache()
	return b.Do(doPage.ctxWithExecutor)
}

// GenerateTestReport generates a report for testing.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-generateTestReport
//
// parameters:
//  - `message`: Message to be displayed in the report.
//  - `group`: This can be nil. (Optional) Specifies the endpoint group to deliver the report to.
func (doPage Page) GenerateTestReport(message string, group *string) (err error) {
	b := page.GenerateTestReport(message)
	if group != nil {
		b = b.WithGroup(*group)
	}
	return b.Do(doPage.ctxWithExecutor)
}

// WaitForDebugger pauses page execution. Can be resumed using generic
// Runtime.runIfWaitingForDebugger.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-waitForDebugger
func (doPage Page) WaitForDebugger() (err error) {
	b := page.WaitForDebugger()
	return b.Do(doPage.ctxWithExecutor)
}

// SetInterceptFileChooserDialog intercept file chooser requests and transfer
// control to protocol clients. When file chooser interception is enabled,
// native file chooser dialog is not shown. Instead, a protocol event
// Page.fileChooserOpened is emitted.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#method-setInterceptFileChooserDialog
//
// parameters:
//  - `enabled`
func (doPage Page) SetInterceptFileChooserDialog(enabled bool) (err error) {
	b := page.SetInterceptFileChooserDialog(enabled)
	return b.Do(doPage.ctxWithExecutor)
}
