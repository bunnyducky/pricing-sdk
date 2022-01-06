package pricingsdk

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func getTestHttpClient(t *testing.T) *Client {
	sLogger := zap.S()

	client := &Client{
		Host:   "https://staging.partyparrot.finance",
		Client: http.DefaultClient,
		Logger: sLogger,
	}

	return client
}

func TestMain(t *testing.T) {
	pricing, err := getTestHttpClient(t).
		FetchPricing("dCN5mwZbDeWCHfp9NF7tv9VVHPmjPSKLpUnKc4WC8bJ")
	assert.NoError(t, err)
	assert.Greater(t, pricing.ROI, float64(0))
	assert.Greater(t, pricing.BondingPrice, float64(0))
	assert.Greater(t, pricing.MarketPrice, float64(0))
	//assert.Greater(t, pricing.MaxPayout, float64(0)) // TODO temporality
}
