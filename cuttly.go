package cuttly

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	Key     string
	BaseURL string
}

type Response struct {
	URL   `json:"url"`
	Stats `json:"stats"`
}

type URL struct {
	Status    int    `json:"status"`
	FullLink  string `json:"fullLink"`
	Date      string `json:"date"`
	ShortLink string `json:"shortLink"`
	Title     string `json:"title"`
}
type Stats struct {
	Status    int    `json:"status"`
	Clicks    int    `json:"clicks"`
	Date      string `json:"date"`
	Title     string `json:"title"`
	FullLink  string `json:"fullLink"`
	ShortLink string `json:"shortLink"`
	Facebook  int    `json:"facebook"`
	Twitter   int    `json:"twitter"`
	Pinterest int    `json:"pinterest"`
	Instagram int    `json:"instagram"`
	GPlus     int    `json:"googlePlus"`
	LinkedIn  int    `json:"linkedin"`
	Rest      int    `json:"rest"`
	//cutt.ly API most of the time supplies empty array for the below options
	Devices interface{}
	Refs    interface{}
}

func (c *Client) Shorten(longURL, name string) (URL, error) {
	v := url.Values{}
	v.Set("key", c.Key)
	v.Add("name", name)
	v.Add("short", longURL)

	r, err := http.Get(c.BaseURL + v.Encode())
	if err != nil {
		return URL{}, err
	}
	defer r.Body.Close()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return URL{}, err
	}

	var resp Response

	err = json.Unmarshal(content, &resp)

	if err != nil {
		return URL{}, err
	}

	err = checkErrorCode(resp.URL.Status, true)

	return resp.URL, err
}

func (c *Client) GetStats(shortURL string) (Stats, error) {
	v := url.Values{}
	v.Set("key", c.Key)
	v.Add("stats", shortURL)

	r, err := http.Get(c.BaseURL + v.Encode())
	if err != nil {
		return Stats{}, err
	}
	defer r.Body.Close()

	content, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return Stats{}, err
	}
	var response Response
	err = json.Unmarshal(content, &response)

	return response.Stats, err
}
