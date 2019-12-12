
/*
Package domain helps executing cdproto commands.
*/
package domain

import (
	"context"
)

// Domain executes a cdproto command against a chromedp context.
type Domain struct {
	ctxWithExecutor context.Context // chromedp context
}

// Do makes a Domain that executes a cdproto command against a chromedp context.
// Set up arg `ctxWithExecutor` with `cdp.WithExecutor()` and `chromedp.FromContext()` beforehand.
func Do(ctxWithExecutor context.Context) Domain {
	return Domain{ctxWithExecutor}
}

// Accessibility executes a cdproto command under Accessibility domain.
func (d Domain) Accessibility() Accessibility {
	return Accessibility{d.ctxWithExecutor}
}

// Animation executes a cdproto command under Animation domain.
func (d Domain) Animation() Animation {
	return Animation{d.ctxWithExecutor}
}

// ApplicationCache executes a cdproto command under ApplicationCache domain.
func (d Domain) ApplicationCache() ApplicationCache {
	return ApplicationCache{d.ctxWithExecutor}
}

// Audits executes a cdproto command under Audits domain.
func (d Domain) Audits() Audits {
	return Audits{d.ctxWithExecutor}
}

// BackgroundService executes a cdproto command under BackgroundService domain.
func (d Domain) BackgroundService() BackgroundService {
	return BackgroundService{d.ctxWithExecutor}
}

// Browser executes a cdproto command under Browser domain.
func (d Domain) Browser() Browser {
	return Browser{d.ctxWithExecutor}
}

// CSS executes a cdproto command under CSS domain.
func (d Domain) CSS() CSS {
	return CSS{d.ctxWithExecutor}
}

// CacheStorage executes a cdproto command under CacheStorage domain.
func (d Domain) CacheStorage() CacheStorage {
	return CacheStorage{d.ctxWithExecutor}
}

// Cast executes a cdproto command under Cast domain.
func (d Domain) Cast() Cast {
	return Cast{d.ctxWithExecutor}
}

// DOM executes a cdproto command under DOM domain.
func (d Domain) DOM() DOM {
	return DOM{d.ctxWithExecutor}
}

// DOMDebugger executes a cdproto command under DOMDebugger domain.
func (d Domain) DOMDebugger() DOMDebugger {
	return DOMDebugger{d.ctxWithExecutor}
}

// DOMSnapshot executes a cdproto command under DOMSnapshot domain.
func (d Domain) DOMSnapshot() DOMSnapshot {
	return DOMSnapshot{d.ctxWithExecutor}
}

// DOMStorage executes a cdproto command under DOMStorage domain.
func (d Domain) DOMStorage() DOMStorage {
	return DOMStorage{d.ctxWithExecutor}
}

// Database executes a cdproto command under Database domain.
func (d Domain) Database() Database {
	return Database{d.ctxWithExecutor}
}

// Debugger executes a cdproto command under Debugger domain.
func (d Domain) Debugger() Debugger {
	return Debugger{d.ctxWithExecutor}
}

// DeviceOrientation executes a cdproto command under DeviceOrientation domain.
func (d Domain) DeviceOrientation() DeviceOrientation {
	return DeviceOrientation{d.ctxWithExecutor}
}

// Emulation executes a cdproto command under Emulation domain.
func (d Domain) Emulation() Emulation {
	return Emulation{d.ctxWithExecutor}
}

// Fetch executes a cdproto command under Fetch domain.
func (d Domain) Fetch() Fetch {
	return Fetch{d.ctxWithExecutor}
}

// HAR executes a cdproto command under HAR domain.
func (d Domain) HAR() HAR {
	return HAR{d.ctxWithExecutor}
}

// HeadlessExperimental executes a cdproto command under HeadlessExperimental domain.
func (d Domain) HeadlessExperimental() HeadlessExperimental {
	return HeadlessExperimental{d.ctxWithExecutor}
}

// HeapProfiler executes a cdproto command under HeapProfiler domain.
func (d Domain) HeapProfiler() HeapProfiler {
	return HeapProfiler{d.ctxWithExecutor}
}

// IO executes a cdproto command under IO domain.
func (d Domain) IO() IO {
	return IO{d.ctxWithExecutor}
}

// IndexedDB executes a cdproto command under IndexedDB domain.
func (d Domain) IndexedDB() IndexedDB {
	return IndexedDB{d.ctxWithExecutor}
}

// Input executes a cdproto command under Input domain.
func (d Domain) Input() Input {
	return Input{d.ctxWithExecutor}
}

// Inspector executes a cdproto command under Inspector domain.
func (d Domain) Inspector() Inspector {
	return Inspector{d.ctxWithExecutor}
}

// LayerTree executes a cdproto command under LayerTree domain.
func (d Domain) LayerTree() LayerTree {
	return LayerTree{d.ctxWithExecutor}
}

// Log executes a cdproto command under Log domain.
func (d Domain) Log() Log {
	return Log{d.ctxWithExecutor}
}

// Media executes a cdproto command under Media domain.
func (d Domain) Media() Media {
	return Media{d.ctxWithExecutor}
}

// Memory executes a cdproto command under Memory domain.
func (d Domain) Memory() Memory {
	return Memory{d.ctxWithExecutor}
}

// Network executes a cdproto command under Network domain.
func (d Domain) Network() Network {
	return Network{d.ctxWithExecutor}
}

// Overlay executes a cdproto command under Overlay domain.
func (d Domain) Overlay() Overlay {
	return Overlay{d.ctxWithExecutor}
}

// Page executes a cdproto command under Page domain.
func (d Domain) Page() Page {
	return Page{d.ctxWithExecutor}
}

// Performance executes a cdproto command under Performance domain.
func (d Domain) Performance() Performance {
	return Performance{d.ctxWithExecutor}
}

// Profiler executes a cdproto command under Profiler domain.
func (d Domain) Profiler() Profiler {
	return Profiler{d.ctxWithExecutor}
}

// Runtime executes a cdproto command under Runtime domain.
func (d Domain) Runtime() Runtime {
	return Runtime{d.ctxWithExecutor}
}

// Security executes a cdproto command under Security domain.
func (d Domain) Security() Security {
	return Security{d.ctxWithExecutor}
}

// ServiceWorker executes a cdproto command under ServiceWorker domain.
func (d Domain) ServiceWorker() ServiceWorker {
	return ServiceWorker{d.ctxWithExecutor}
}

// Storage executes a cdproto command under Storage domain.
func (d Domain) Storage() Storage {
	return Storage{d.ctxWithExecutor}
}

// SystemInfo executes a cdproto command under SystemInfo domain.
func (d Domain) SystemInfo() SystemInfo {
	return SystemInfo{d.ctxWithExecutor}
}

// Target executes a cdproto command under Target domain.
func (d Domain) Target() Target {
	return Target{d.ctxWithExecutor}
}

// Tethering executes a cdproto command under Tethering domain.
func (d Domain) Tethering() Tethering {
	return Tethering{d.ctxWithExecutor}
}

// Tracing executes a cdproto command under Tracing domain.
func (d Domain) Tracing() Tracing {
	return Tracing{d.ctxWithExecutor}
}

// WebAudio executes a cdproto command under WebAudio domain.
func (d Domain) WebAudio() WebAudio {
	return WebAudio{d.ctxWithExecutor}
}

// WebAuthn executes a cdproto command under WebAuthn domain.
func (d Domain) WebAuthn() WebAuthn {
	return WebAuthn{d.ctxWithExecutor}
}

