# Chromebot

[![GoDoc](https://godoc.org/github.com/nanitefactory/chromebot?status.svg)](https://godoc.org/github.com/nanitefactory/chromebot)
[![Build Status](https://travis-ci.org/NaniteFactory/chromebot.svg?branch=master)](https://travis-ci.org/NaniteFactory/chromebot)

Package `chromebot` provides a high-level interface to automate tasks in Chrome or Chromium. Work based on [chromedp](https://github.com/chromedp/chromedp).

## Installation

```Cmd
go get -v github.com/nanitefactory/chromebot
```

## Test

```Cmd
go test -v github.com/nanitefactory/chromebot
```

## Quick start

```Go
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
```
