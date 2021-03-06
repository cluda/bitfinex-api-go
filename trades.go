package bitfinex

import (
	"net/url"
	"strconv"
	"strings"
)

type TradesService struct {
	client *Client
}

type Trade struct {
	Price     string
	Amount    string
	Exchange  string
	Type      string
	Timestamp int64
	TradeId   int64 `json:"tid,int"`
}

// All returns trades where first is newest and last is oldest
func (s *TradesService) All(pair string, timestamp int64, limitTrades int) ([]Trade, error) {
	pair = strings.ToUpper(pair)

	params := url.Values{}
	if timestamp != 0 {
		params.Add("timestamp", strconv.FormatInt(timestamp, 10))
	}
	if limitTrades != 0 {
		params.Add("limit_trades", strconv.Itoa(limitTrades))
	}
	req, err := s.client.newRequest("GET", "trades/"+pair, params)
	if err != nil {
		return nil, err
	}

	var v []Trade

	_, err = s.client.do(req, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
