package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Client struct {
	Endpoint string `env:""`
}

func (c *Client) SetDefaults() {
	if c.Endpoint == "" {
		c.Endpoint = "http://127.0.0.1/v0/app/v0/object"
	}
}

func (c *Client) PresignPutURL(object string) (permlink string, s3link string, err error) {

	objectUrl := fmt.Sprintf("%s/%s", c.Endpoint, object)
	s3link, err = getLinks(objectUrl)

	return objectUrl, s3link, err
}

func (c *Client) PutFile(url string, file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	return putFile(url, data)
}

type PresignPutResponse struct {
	Status int
	Error  string
	Data   struct {
		PermanentiLink    string
		TemporaryRedirect string
	}
}

func getLinks(u string) (s3link string, err error) {
	resp, err := http.Post(u, "", nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	r := &PresignPutResponse{}
	err = json.Unmarshal(data, r)
	if err != nil {
		return
	}

	if r.Error != "" {
		return "", errors.New(r.Error)
	}

	return r.Data.TemporaryRedirect, nil
}

func putFile(s3link string, data []byte) error {
	req, _ := http.NewRequest("PUT", s3link, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "multipart/form-data")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
