/**
 * This code is licensed under MIT license.
 * Please see LICENSE.md file for full license.
 */

package geohash_test

import (
	"math"
	"testing"

	. "github.com/tapglue/geohash"
)

func TestEncodeInt(t *testing.T) {
	var expected uint64 = 4064984913515641
	encodedVal := EncodeInt(37.8324, 112.5584, 52)
	if encodedVal != expected {
		t.Logf("Expected: %d Got: %d\n", encodedVal, expected)
		t.Fail()
	}
}

func TestDecodeInt(t *testing.T) {
	lat, lon, _, _ := DecodeInt(4064984913515641, 52)
	if math.Abs(37.8324-lat) >= 0.0001 {
		t.Logf("37.8324 - %.5f was >= 0.0001\n", lat)
		t.Fail()
	}
	if math.Abs(112.5584-lon) >= 0.0001 {
		t.Logf("112.5584 - %.5f was >= 0.0001\n", lon)
		t.Fail()
	}
}
