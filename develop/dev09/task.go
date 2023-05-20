package dev09

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

const u = "https://gobyexample.com/"

var dir = ""

func Main() {
	d, err := createDirectory()
	if err != nil {
		panic(err)
	}
	dir = d
	parse("")
}

func parse(relativePath string) {
	resp, err := http.Get(u + relativePath)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	tkn := html.NewTokenizer(bytes.NewReader(data))

	links := make([]string, 0)
	styles := make([]string, 0)

	for {
		tt := tkn.Next()
		if tt == html.ErrorToken {
			err = tkn.Err()
			if err == io.EOF {
				break
			}
		}
		if tt == html.StartTagToken {
			t := tkn.Token()
			if t.Data == "a" {
				v, ok := getAttribute(&t, "href")
				if ok {
					links = append(links, v)
				}
			}
			if t.Data == "link" {
				v, ok := getAttribute(&t, "rel")
				if ok && v == "stylesheet" {
					v, ok = getAttribute(&t, "href")
					if ok {
						styles = append(styles, v)
					}
				}
			}
		}
	}

	name := relativePath
	if name == "" {
		name = "index.html"
	}
	fmt.Println("relativePath:", relativePath, "name:", name)
	err = saveHtml(data, dir, name)
	if err != nil {
		if err == os.ErrExist {
			return
		}
		panic(err)
	}

	err = downloadAndSaveStyles(styles, dir)
	if err != nil {
		panic(err)
	}

	for _, link := range links {
		if link == "./" {
			continue
		}
		if !isUrl(link) || path.Base(u) == link {
			fmt.Println("LINK:", link)
			parse(link)
		}
	}
}

func getAttribute(token *html.Token, key string) (string, bool) {
	for _, a := range token.Attr {
		if a.Key == key {
			return a.Val, true
		}
	}
	return "", false
}

func createDirectory() (dir string, err error) {
	dir = path.Base(u)
	n := 0
	for {
		_, err = os.Stat(dir)
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				return "", err
			}
			break
		} else {
			n++
			dir = path.Base(u) + strconv.Itoa(n)
		}
	}
	return dir + "/", nil
}

func saveHtml(data []byte, dir, name string) error {
	_, err := os.Stat(dir + name)
	if os.IsNotExist(err) {
		return os.WriteFile(dir+name, data, 0644)
	}
	return os.ErrExist
}

func downloadAndSaveHtml(name, dir string) error {
	resp, err := http.Get(u + name)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = saveHtml(data, dir, name)
	if err != nil {
		return err
	}
	return nil
}

func downloadAndSaveStyles(styles []string, dir string) error {
	for _, s := range styles {
		resp, err := http.Get(u + s)
		if err != nil {
			return err
		}
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		_, err = os.Stat(dir + s)
		if os.IsNotExist(err) {
			err = os.WriteFile(dir+s, data, 0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func isUrl(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}
