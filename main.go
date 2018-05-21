// Downloads all photos from woz.org :D

package main

import "fmt"
import "net/http"
import "io/ioutil"
import "strings"
import "strconv"

import "bytes"

import "golang.org/x/net/html"

var count int

func main() {
	resp, err := http.Get("http://woz.org/photos")
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	z := html.NewTokenizer(bytes.NewReader(b))

	resp.Body.Close()

	var Found []string

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			tn := z.Token()
			isAnchor := tn.Data == "a"
			if isAnchor {
				for _, a := range tn.Attr {
					if a.Key == "href" {
						if strings.Index(a.Val, "/photos/") == 0 {
							for _, v := range Found {
								if strings.Contains(v, a.Val) {
									continue
								}
							}
							Found = append(Found, a.Val)
							Download(a.Val)
						}
					}
				}
			} else {
			}
		}
	}

	fmt.Println(string(b))
}

func DownloadDownload(a string) {
	resp, err := http.Get("http://woz.org" + a)
	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp.Body.Close()

	z := html.NewTokenizer(bytes.NewReader(b))

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return

		case html.StartTagToken:
			tn := z.Token()
			isAnchor := tn.Data == "img"

			if isAnchor {
				for _, c := range tn.Attr {
					if c.Key == "src" {
						if strings.Index(c.Val, "http") == 0 {
							resp, err := http.Get(c.Val)
							if err != nil {
								continue
							}
							defer resp.Body.Close()

							b, err := ioutil.ReadAll(resp.Body)
							if err != nil {
								continue
							}

							fmt.Println(b)
							err = ioutil.WriteFile("woz"+strconv.Itoa(count)+".jpeg", b, 0644)
							count++
							if err != nil {
								panic(err)
							}

						}
					}
				}
			}
		}
	}
}

func Download(a string) {
	resp, err := http.Get("http://woz.org" + a)
	fmt.Println(a)
	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp.Body.Close()

	z := html.NewTokenizer(bytes.NewReader(b))
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			tn := z.Token()
			isAnchor := tn.Data == "a"

			if isAnchor {
				for _, b := range tn.Attr {
					if b.Key == "href" {
						if strings.Index(b.Val, "/photos/") == 0 {
							fmt.Println(b.Val)
							DownloadDownload(b.Val)
						}

					}
				}
			}

		}
	}
}
