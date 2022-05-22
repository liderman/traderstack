package apiclient

import (
	"context"
	"fmt"
	"github.com/liderman/traderstack/internal/domain"
	"github.com/liderman/traderstack/internal/grpc/gen/tinkoff/investapi"
	"github.com/liderman/traderstack/internal/grpc/mapper"
)

func (a *RealClient) GetShares() ([]*domain.Share, error) {
	resp, err := a.cInstrument.Shares(context.Background(), &investapi.InstrumentsRequest{
		InstrumentStatus: investapi.InstrumentStatus_INSTRUMENT_STATUS_BASE,
	})
	if err != nil {
		return nil, fmt.Errorf("error api Shares: %w", err)
	}

	return mapper.MapFromShares(resp.Instruments), nil
}
