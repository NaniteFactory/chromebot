package chromebot

import (
	"context"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
)

// Builder builds a Chrome instance.
type Builder struct {
	deadline *time.Time
	eOpts    []chromedp.ExecAllocatorOption
	bOpts    []chromedp.BrowserOption
	cOpts    []chromedp.ContextOption
}

// NewBuilder is a constructor.
func NewBuilder() *Builder {
	return &Builder{
		deadline: nil,
		eOpts:    []chromedp.ExecAllocatorOption{},
		bOpts:    []chromedp.BrowserOption{},
		cOpts:    []chromedp.ContextOption{},
	}
}

// BuilderFlags is a set of generic command line options to pass as flags to Chrome.
type BuilderFlags struct {
	NoFirstRun            bool
	NoDefaultBrowserCheck bool
	Headless              bool
	HideScrollbars        bool
	MuteAudio             bool
	//
	UserDataDir string // UserDataDir is the command line option to set the profile directory used by Chrome.
	ProxyServer string // ProxyServer is the command line option to set the outbound proxy server.
	WindowSize  struct {
		Width  int
		Height int
	}
	UserAgent  string
	NoSandbox  bool
	DisableGPU bool
	//
	DisableBackgroundNetworking         bool
	EnableFeatures                      string
	DisableBackgroundTimerThrottling    bool
	DisableBackgroundingOccludedWindows bool
	DisableBreakpad                     bool
	DisableClientSidePhishingDetection  bool
	DisableDefaultApps                  bool
	DisableDevShmUsage                  bool
	DisableExtensions                   bool
	DisableFeatures                     string
	DisableHangMonitor                  bool
	DisableIpcFloodingProtection        bool
	DisablePopupBlocking                bool
	DisablePromptOnRepost               bool
	DisableRendererBackgrounding        bool
	DisableSync                         bool
	ForceColorProfile                   string
	MetricsRecordingOnly                bool
	SafebrowsingDisableAutoUpdate       bool
	EnableAutomation                    bool
	PasswordStore                       string
	UseMockKeychain                     bool
}

// DefaultBuilderFlags is a set of generic command line options to pass as flags to Chrome.
var DefaultBuilderFlags = BuilderFlags{
	NoFirstRun:            true,
	NoDefaultBrowserCheck: true,
	Headless:              true,
	HideScrollbars:        true,
	MuteAudio:             true,
	//
	UserDataDir: "",
	ProxyServer: "",
	WindowSize: struct {
		Width  int
		Height int
	}{-1, -1},
	UserAgent:  "",
	NoSandbox:  false,
	DisableGPU: false,
	//
	DisableBackgroundNetworking:         true,
	EnableFeatures:                      "NetworkService,NetworkServiceInProcess",
	DisableBackgroundTimerThrottling:    true,
	DisableBackgroundingOccludedWindows: true,
	DisableBreakpad:                     true,
	DisableClientSidePhishingDetection:  true,
	DisableDefaultApps:                  true,
	DisableDevShmUsage:                  true,
	DisableExtensions:                   true,
	DisableFeatures:                     "site-per-process,TranslateUI,BlinkGenPropertyTrees",
	DisableHangMonitor:                  true,
	DisableIpcFloodingProtection:        true,
	DisablePopupBlocking:                true,
	DisablePromptOnRepost:               true,
	DisableRendererBackgrounding:        true,
	DisableSync:                         true,
	ForceColorProfile:                   "srgb",
	MetricsRecordingOnly:                true,
	SafebrowsingDisableAutoUpdate:       true,
	EnableAutomation:                    true,
	PasswordStore:                       "basic",
	UseMockKeychain:                     true,
}

// WithFlags sets a set of generic command line options to pass as flags to Chrome.
// Recommend use with `DefaultFlags`.
func (b *Builder) WithFlags(arg BuilderFlags) *Builder {
	b.eOpts = append(b.eOpts, []chromedp.ExecAllocatorOption{
		chromedp.Flag("no-first-run", arg.NoFirstRun),
		chromedp.Flag("no-default-browser-check", arg.NoDefaultBrowserCheck),
		chromedp.Flag("headless", arg.Headless),
		chromedp.Flag("hide-scrollbars", arg.HideScrollbars),
		chromedp.Flag("mute-audio", arg.MuteAudio),
		//
		chromedp.Flag("no-sandbox", arg.NoSandbox),
		chromedp.Flag("disable-gpu", arg.DisableGPU),
		//
		chromedp.Flag("disable-background-networking", arg.DisableBackgroundNetworking),
		chromedp.Flag("enable-features", arg.EnableFeatures),
		chromedp.Flag("disable-background-timer-throttling", arg.DisableBackgroundTimerThrottling),
		chromedp.Flag("disable-backgrounding-occluded-windows", arg.DisableBackgroundingOccludedWindows),
		chromedp.Flag("disable-breakpad", arg.DisableBreakpad),
		chromedp.Flag("disable-client-side-phishing-detection", arg.DisableClientSidePhishingDetection),
		chromedp.Flag("disable-default-apps", arg.DisableDefaultApps),
		chromedp.Flag("disable-dev-shm-usage", arg.DisableDevShmUsage),
		chromedp.Flag("disable-extensions", arg.DisableExtensions),
		chromedp.Flag("disable-features", arg.DisableFeatures),
		chromedp.Flag("disable-hang-monitor", arg.DisableHangMonitor),
		chromedp.Flag("disable-ipc-flooding-protection", arg.DisableIpcFloodingProtection),
		chromedp.Flag("disable-popup-blocking", arg.DisablePopupBlocking),
		chromedp.Flag("disable-prompt-on-repost", arg.DisablePromptOnRepost),
		chromedp.Flag("disable-renderer-backgrounding", arg.DisableRendererBackgrounding),
		chromedp.Flag("disable-sync", arg.DisableSync),
		chromedp.Flag("force-color-profile", arg.ForceColorProfile),
		chromedp.Flag("metrics-recording-only", arg.MetricsRecordingOnly),
		chromedp.Flag("safebrowsing-disable-auto-update", arg.SafebrowsingDisableAutoUpdate),
		chromedp.Flag("enable-automation", arg.EnableAutomation),
		chromedp.Flag("password-store", arg.PasswordStore),
		chromedp.Flag("use-mock-keychain", arg.UseMockKeychain),
	}...)
	if arg.UserDataDir != "" {
		b.eOpts = append(b.eOpts, chromedp.Flag("user-data-dir", arg.UserDataDir))
	}
	if arg.ProxyServer != "" {
		b.eOpts = append(b.eOpts, chromedp.Flag("proxy-server", arg.ProxyServer))
	}
	if arg.WindowSize.Width > 0 && arg.WindowSize.Height > 0 {
		b.eOpts = append(b.eOpts, chromedp.Flag("window-size", arg.WindowSize))
	}
	if arg.UserAgent != "" {
		b.eOpts = append(b.eOpts, chromedp.Flag("user-agent", arg.UserAgent))
	}
	return b
}

// WithExecAllocatorOption sets options.
// Generally unused.
func (b *Builder) WithExecAllocatorOption(opts ...chromedp.ExecAllocatorOption) *Builder {
	b.eOpts = append(b.eOpts, opts...)
	return b
}

// WithBrowserOption sets options.
// Generally unused.
func (b *Builder) WithBrowserOption(opts ...chromedp.BrowserOption) *Builder {
	b.bOpts = append(b.bOpts, opts...)
	return b
}

// WithContextOption sets options.
// Generally unused.
func (b *Builder) WithContextOption(opts ...chromedp.ContextOption) *Builder {
	b.cOpts = append(b.cOpts, opts...)
	return b
}

// WithDeadline sets the deadline.
// Generally unused.
func (b *Builder) WithDeadline(d time.Time) *Builder {
	b.deadline = &d
	return b
}

// NewChrome creates a Chrome instance.
func (b *Builder) NewChrome() *Chrome {
	NewWithOptions := func(
		deadline *time.Time,
		eOpts []chromedp.ExecAllocatorOption,
		bOpts []chromedp.BrowserOption,
		cOpts []chromedp.ContextOption,
	) *Chrome {
		if eOpts == nil {
			eOpts = []chromedp.ExecAllocatorOption{}
		}
		if bOpts == nil {
			bOpts = []chromedp.BrowserOption{}
		}
		if cOpts == nil {
			cOpts = []chromedp.ContextOption{}
		}
		baseCtx, baseCancel := context.Background(), context.CancelFunc(nil)
		if deadline != nil {
			baseCtx, baseCancel = context.WithDeadline(baseCtx, *deadline)
		}
		exeCtx, exeCancel := chromedp.NewExecAllocator(baseCtx, eOpts...)
		tabCtx, tabCancel := chromedp.NewContext(exeCtx, append(cOpts[:], chromedp.WithBrowserOption(bOpts...))...)
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
				default:
					ret.exeCancel()
					if baseCancel != nil { // if with deadline
						baseCancel() // free all contexts
					}
					close(ret.fin)
				}
			}()
		}()
		return ret
	}
	return NewWithOptions(b.deadline, b.eOpts, b.bOpts, b.cOpts)
}
