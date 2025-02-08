package order

import (
	"context"
	error "express_be/core/err"
	mapper "express_be/mapper/res"

	"express_be/model/res"
	"express_be/utils/geohash"
)

// FindNearestDeliveryPerson implements OrderUsecase.
func (o *orderUsecaseImpl) FindNearestDeliveryPerson(ctx context.Context, lat, lon *float64, status *string) (*res.BaseDeliveryPersonRes, *error.Err) {
	geoHashes, err := o.deliveryRepo.FetchOffDutyGeoHashes(ctx, status)
	if err != nil {
		return nil, error.ErrInternalServer
	}
	if len(geoHashes) == 0 {
		return nil, error.ErrInvalidFormat
	}
	nearestGeoHash := geohash.FindNearestGeohash(*lat, *lon, geoHashes)
	deliveryPerson, err := o.deliveryRepo.FindByGeohash(ctx, &nearestGeoHash)
	if err != nil {
		return nil, error.ErrInternalServer
	}
	nearestDeliveryPerson := mapper.BaseDeliveryPersonToRes(deliveryPerson)

	return &nearestDeliveryPerson, nil
}
