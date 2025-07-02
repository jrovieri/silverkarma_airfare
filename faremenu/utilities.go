package main

import "math"

func CalculateDistance(dstLng, dstLat, srcLng, srcLat float64) float64 {

	radSrcLat := float64(math.Pi * srcLat / 180)
	radDstLat := float64(math.Pi * dstLat / 180)

	theta := float64(srcLng - dstLng)
	radTheta := float64(math.Pi * theta / 180)

	distance := math.Sin(radSrcLat)*math.Sin(radDstLat) +
		math.Cos(radSrcLat)*math.Cos(radDstLat)*math.Cos(radTheta)

	if distance > 1 {
		distance = 1
	}

	distance = math.Acos(distance)
	distance = distance * 180 / math.Pi
	distance = distance * 60 * 1.1515
	distance = distance * 1.609344

	return distance
}
