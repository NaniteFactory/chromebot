package chromebot_test

import (
	"fmt"
	"testing"

	"github.com/chromedp/chromedp"
	"github.com/nanitefactory/chromebot"
)

func ExampleNew() {
	// Start a browser
	browser := chromebot.New(false)
	defer browser.Close()

	// Navigate
	browser.Tab(0).Run(
		chromedp.Navigate("https://godoc.org"),
		chromedp.WaitVisible("html"),
	)

	// Another navigation method
	// browser.Tab(0).Do().Page().Navigate("https://godoc.org", nil, nil, nil)

	// Get cookies
	cookies, err := browser.Tab(0).Do().Network().GetAllCookies()
	if err != nil {
		panic(err)
	}

	// Print
	for i, cookie := range cookies {
		fmt.Println(i, cookie)
	}

	// _Output:
	// 0 &{__utma 58978491.1884756898.1576137201.1576137201.1576137201.1 .godoc.org / 1.639209201e+09 60 false false false }
	// 1 &{__utmc 58978491 .godoc.org / -1 14 false false true }
	// 2 &{__utmz 58978491.1576137201.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none) .godoc.org / 1.591905201e+09 75 false false false }
	// 3 &{__utmt 1 .godoc.org / 1.576137801e+09 7 false false false }
	// 4 &{__utmb 58978491.1.10.1576137201 .godoc.org / 1.576139001e+09 30 false false false }
}

func TestExampleNew(t *testing.T) {
	// Start a browser
	browser := chromebot.New(bFlagHeadless)
	defer browser.Close()
	// Navigate
	if err := browser.Tab(0).Run(
		chromedp.Navigate("https://godoc.org"),
		chromedp.WaitVisible("html"),
	); err != nil {
		panic(err)
	}
	// Another navigation method
	if _, _, strErr, err := browser.Tab(0).Do().Page().Navigate("https://godoc.org", nil, nil, nil); err != nil {
		panic(fmt.Sprint(err, strErr))
	}
	// Get cookies
	cookies, err := browser.Tab(0).Do().Network().GetAllCookies()
	if err != nil {
		panic(err)
	}
	// Print
	for i, cookie := range cookies {
		t.Log(i, cookie)
	}
}
