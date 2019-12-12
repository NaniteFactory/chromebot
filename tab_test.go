package chromebot_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/chromedp/chromedp"
	"github.com/nanitefactory/chromebot"
)

func TestTab(t *testing.T) {
	c := chromebot.New(false)
	defer c.Close()
	c.AddNewTab()
	//
	wg := sync.WaitGroup{}
	doRunTest := func(iTab int) {
		defer wg.Add(-1)
		// Navigate
		if err := c.Tab(iTab).Run(
			chromedp.Navigate("https://godoc.org"),
			chromedp.WaitVisible("html"),
		); err != nil {
			panic(err)
		}
		// Another navigation method
		if _, _, strErr, err := c.Tab(iTab).Do().Page().Navigate("https://godoc.org", nil, nil, nil); err != nil {
			panic(fmt.Sprint(err, strErr))
		}
		// Get cookies
		cookies, err := c.Tab(iTab).Do().Network().GetAllCookies()
		if err != nil {
			panic(err)
		}
		// Print
		for i, cookie := range cookies {
			t.Log(i, cookie)
		}
	}
	c.Tab(0).Listen(func(interface{}) {})
	c.Tab(1).Listen(func(interface{}) {})
	wg.Add(2)
	go doRunTest(0)
	go doRunTest(1)
	wg.Wait()
	c.Tab(0).Close()
	c.Tab(1).Close()
	for i := 0; i < 10; i++ {
		select {
		case <-c.Tab(0).Dead():
		default:
			panic("Dead() not working well")
		}
	}
	for i := 0; i < 10; i++ {
		select {
		case <-c.Tab(1).Dead():
		default:
			panic("Dead() not working well")
		}
	}
}
