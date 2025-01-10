package entity

import "express_be/core/record"

type Area struct {
	record.BaseEntity
	Name                   *string  `json:"name"`
	GeoHash                *string  `json:"geo_hash"`
	City                   *string  `json:"city"`
	District               *string  `json:"district"`
	Ward                   *string  `json:"ward"`
	CenterLat              *float64 `json:"center_lat"`
	CenterLng              *float64 `json:"center_lng"`
	RadiusKm               *float64 `json:"radius_km"`
	Type                   *string  `json:"type"`
	CustomerCount          *int     `json:"customer_count"`
	DeliveryPersonCapacity *int     `json:"delivery_person_capacity"`
}
