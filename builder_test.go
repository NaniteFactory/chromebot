package chromebot_test

import (
	"testing"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/nanitefactory/chromebot"
)

func TestBuilder(t *testing.T) {
	c := chromebot.NewBuilder().
		WithFlags(chromebot.DefaultBuilderFlags).
		WithExecAllocatorOption(
			chromedp.DisableGPU,
			chromedp.Flag("headless", bFlagHeadless),
		).
		WithBrowserOption().
		WithContextOption().
		NewChrome()
	defer c.Close()
	c.AddNewTab()
	c.AddNewTab()
	time.Sleep(time.Second * 1)
}

func TestBuilderDeadline(t *testing.T) {
	c := chromebot.NewBuilder().WithFlags(func() chromebot.BuilderFlags {
		optFlags := chromebot.DefaultBuilderFlags
		optFlags.Headless = bFlagHeadless
		optFlags.MuteAudio = bFlagHeadless
		optFlags.HideScrollbars = bFlagHeadless
		return optFlags
	}()).WithDeadline(time.Now().Add(time.Second * 3)).NewChrome()
	c.AddNewTab()
	c.AddNewTab()
	c.AddNewTab()
	<-c.Dead()
	if nTabs := c.CountTabs(); nTabs != 4 {
		panic(nTabs)
	}
}
