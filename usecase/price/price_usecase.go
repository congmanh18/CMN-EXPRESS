package price

import (
	"context"
	mapper "express_be/mapper/req"

	error "express_be/core/err"
	model "express_be/model/req"

	"express_be/entity"
	"express_be/repository/price"
)

type PriceUsecase interface {
	CreateNewPrice(ctx context.Context, req model.PriceReq, adminID *string) *error.Err
	ReadAllPrice(ctx context.Context) ([]entity.BasicPrice, *error.Err)
	UpdatePrice(ctx context.Context, id *string, req model.PriceReq, adminID *string) *error.Err
	DeletePrice(ctx context.Context, id *string, adminID *string) *error.Err
}

type priceUsecaseImpl struct {
	priceRepo price.Repo
}

// CreatePrice implements PriceUsecase.
func (p *priceUsecaseImpl) CreateNewPrice(ctx context.Context, req model.PriceReq, adminID *string) *error.Err {
	new_price := mapper.PriceReqToBasicPrice(req)
	err := p.priceRepo.Create(ctx, new_price, adminID)
	if err != nil {
		return error.ErrInternalServer
	}

	return nil
}

// ReadAllPrice implements PriceUsecase.
func (p *priceUsecaseImpl) ReadAllPrice(ctx context.Context) ([]entity.BasicPrice, *error.Err) {
	prices, err := p.priceRepo.ReadAllPrice(ctx)
	if err != nil {
		return nil, error.ErrInternalServer
	}
	return prices, nil
}

// UpdatePrice implements PriceUsecase.
func (p *priceUsecaseImpl) UpdatePrice(ctx context.Context, id *string, req model.PriceReq, adminID *string) *error.Err {
	updated_price := mapper.PriceReqToBasicPrice(req)
	err := p.priceRepo.Update(ctx, id, updated_price, adminID)
	if err != nil {
		return error.ErrInternalServer
	}

	return nil
}

// DeletePrice implements PriceUsecase.
func (p *priceUsecaseImpl) DeletePrice(ctx context.Context, id *string, adminID *string) *error.Err {
	if err := p.priceRepo.Delete(ctx, id, adminID); err != nil {
		return error.ErrInternalServer
	}
	return nil
}

func NewPriceUsecase(priceRepo price.Repo) PriceUsecase {
	return &priceUsecaseImpl{priceRepo: priceRepo}
}
