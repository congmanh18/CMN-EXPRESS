package geohash

import (
	"math"

	"github.com/mmcloughlin/geohash"
)

func CreateGeohash(latitude, longitude float64, precision int) string {
	return geohash.EncodeWithPrecision(latitude, longitude, uint(precision))
}

func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371.0 // Bán kính trái đất (km)

	// Chuyển đổi các tọa độ từ độ sang radian
	lat1Rad := deg2rad(lat1)
	lon1Rad := deg2rad(lon1)
	lat2Rad := deg2rad(lat2)
	lon2Rad := deg2rad(lon2)

	// Sự khác biệt giữa các tọa độ
	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	// Áp dụng công thức Haversine
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Tính khoảng cách
	return R * c
}

// deg2rad chuyển đổi độ sang radian
func deg2rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func FindNearestGeohash(targetLat, targetLon float64, geohashes []string) string {
	var nearestGeohash string
	minDistance := math.MaxFloat64

	for _, hash := range geohashes {
		// Tính tọa độ trung tâm của geohash
		lat, lon := geohash.Decode(hash)
		distance := Haversine(targetLat, targetLon, lat, lon)

		if distance < minDistance {
			minDistance = distance
			nearestGeohash = hash
		}
	}

	return nearestGeohash
}
