package pricingsdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

const ProdEndpoint = "https://bonding-api.bunnyducky.com"

func NewClient(host string, httpClient *http.Client, log *zerolog.Logger) *Client {
	return &Client{
		host:   host,
		client: httpClient,
		log:    log,
	}
}

type Client struct {
	host   string
	client *http.Client
	log    *zerolog.Logger
}

func (c *Client) get(path string, result interface{}) error {
	resp, err := c.client.Get(fmt.Sprintf("%s/api/v1/%s", c.host, strings.TrimLeft(path, "/")))
	if err != nil {
		return errors.Wrap(err, "http post err")
	}
	defer resp.Body.Close()

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	c.log.Debug().Str("path", path).Str("resp", string(respBodyBytes)).Msg("pricing sdk api response")
	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("none ok status: %d", resp.StatusCode)
	}

	err = json.Unmarshal(respBodyBytes, result)
	if err != nil {
		return errors.Wrap(err, "json unmarshal result failed")
	}
	return nil
}
