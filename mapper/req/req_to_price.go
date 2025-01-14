package req

import (
	"express_be/core/pointer"
	"express_be/core/record"
	model "express_be/model/req"
	"express_be/repository/price/entity"

	"github.com/google/uuid"
)

func PriceReqToBasicPrice(req model.PriceReq) *entity.BasicPrice {
	return &entity.BasicPrice{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		Region:    req.Region,
		Area:      req.Area,
		BasePrice: req.BasePrice,
	}
}
