package math_convey

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAverage(t *testing.T) {
	convey.Convey("Average should work correctly", t, func() {
		convey.Convey("Average should be calculated correctly", func() {
			average := Average([]float64{1, 2})
			convey.So(average, convey.ShouldEqual, 1.5)
		})
	})
}
