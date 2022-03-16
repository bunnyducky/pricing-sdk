package pricingsdk

import (
	"net/http"
	"os"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func getTestHttpClient(t *testing.T) *Client {
	logger := zerolog.New(os.Stdout)
	return NewClient("https://bonding-api.bunnyducky.com", http.DefaultClient, &logger)
	// return NewClient("https://staging.partyparrot.finance", http.DefaultClient, &logger)
}

func TestMain(t *testing.T) {
	pricing, err := getTestHttpClient(t).
		FetchPricing(solana.MustPublicKeyFromBase58("6fDbbhGpWLSUVb75o7npL2G17JqdXhX1uRNBsH5xmnUM"))
	assert.NoError(t, err)
	assert.Greater(t, pricing.ROI, float64(0))
	assert.Greater(t, pricing.BondingPrice, float64(0))
	assert.Greater(t, pricing.MarketPrice, float64(0))
	assert.Greater(t, pricing.PayoutAmount, float64(0))
	assert.Greater(t, pricing.MaxPayout, float64(0))
}
