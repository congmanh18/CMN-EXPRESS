package geohash

import (
	"testing"
)

func TestCreateGeohash(t *testing.T) {
	latitude := 10.7769
	longitude := 106.7009
	precision := 5

	expected := "w3gvk"
	result := CreateGeohash(latitude, longitude, precision)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestHaversine(t *testing.T) {
	lat1, lon1 := 10.7769, 106.7009
	lat2, lon2 := 10.7790, 106.7020

	expected := 0.229 // Khoảng cách dự kiến (km)
	result := Haversine(lat1, lon1, lat2, lon2)

	if result < expected-0.1 || result > expected+0.1 {
		t.Errorf("Expected distance around %.2f but got %.2f", expected, result)
	}
}

func TestFindNearestGeohash(t *testing.T) {
	lat := 10.7769
	lon := 106.7009
	geohashes := []string{
		"9q8yy", "9q8yz", "9q8y0", "9q8y1",
	}

	expected := "9q8y0"
	result := FindNearestGeohash(lat, lon, geohashes)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
