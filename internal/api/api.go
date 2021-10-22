package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	UNLISTED     = "1"
	EXPIRE_IN    = "1H"
	PASTE_OPTION = "paste"
	POST_URL     = "https://pastebin.com/api/api_post.php"
	GET_URL      = "https://pastebin.com/raw/"
)

func Send(ApiKey string, message string) string {
	resp, err := http.PostForm(POST_URL, url.Values{
		"api_dev_key":       {ApiKey},
		"api_paste_code":    {message},
		"api_paste_private": {UNLISTED},
		"api_paste_expire":  {EXPIRE_IN},
		"api_option":        {PASTE_OPTION},
	})
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
