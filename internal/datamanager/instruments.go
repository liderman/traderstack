package datamanager

import (
	"errors"
	"fmt"
	"github.com/liderman/traderstack/internal/apiclient"
	"github.com/liderman/traderstack/internal/domain"
	"github.com/patrickmn/go-cache"
	"strings"
	"time"
)

type Instruments struct {
	client apiclient.ApiClient
	cache  *cache.Cache
}

func NewInstruments(client apiclient.ApiClient, cache *cache.Cache) *Instruments {
	return &Instruments{
		client: client,
		cache:  cache,
	}
}

func (i *Instruments) GetShares() ([]*domain.Share, error) {
	if v, ok := i.cache.Get("shares"); ok {
		return v.([]*domain.Share), nil
	}

	resp, err := i.client.GetShares()
	if err != nil {
		return nil, fmt.Errorf("error api GetShares: %w", err)
	}

	i.cache.Set("shares", resp, time.Minute*15)

	return resp, nil
}

func (i *Instruments) GetShareByFigi(figi string) (*domain.Share, error) {
	shares, err := i.GetShares()
	if err != nil {
		return nil, err
	}

	for _, v := range shares {
		if v.Figi == figi {
			return v, nil
		}
	}

	return nil, errors.New("share is not exist")
}

func (i *Instruments) SearchShares(text string, limit int) ([]*domain.Share, error) {
	loText := strings.ToLower(text)
	allShares, err := i.GetShares()
	if err != nil {
		return nil, err
	}

	total := 0
	rating1 := []*domain.Share{}
	rating2 := []*domain.Share{}
	rating3 := []*domain.Share{}
	for _, v := range allShares {
		if total >= limit {
			break
		}
		ticker := strings.ToLower(v.Ticker)
		figi := strings.ToLower(v.Figi)
		name := strings.ToLower(v.Name)

		if ticker == loText || figi == loText || name == loText {
			rating1 = append(rating1, v)
			total++
			continue
		}

		if strings.HasPrefix(name, loText) || strings.HasPrefix(ticker, loText) {
			rating2 = append(rating2, v)
			total++
			continue
		}

		if strings.Contains(name, loText) {
			rating3 = append(rating3, v)
			total++
			continue
		}
	}

	ret := rating1
	ret = append(ret, rating2...)
	ret = append(ret, rating3...)
	return ret, nil
}
