package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type nextYearItem struct {
	mo      time.Month
	d, h, m int
	tests   nextYearTests
}

func (n nextYearItem) Name() string {
	return fmt.Sprintf("%s %dd %dh %dm", n.mo, n.d, n.h, n.m)
}

func (n nextYearItem) F() func(t time.Time) time.Time {
	return NextYear(n.mo, n.d, n.h, n.m)
}

type nextYearTests []struct{ in, out time.Time }

var nextYearTestData = []nextYearItem{
	{
		mo: time.December,
		d:  25,
		h:  9,
		m:  0,
		tests: nextYearTests{
			{
				time.Date(2000, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2000, time.December, 25, 9, 0, 0, 0, time.UTC),
			},
		},
	},
}

func TestNextYear(t *testing.T) {
	for _, i := range nextYearTestData {
		t.Run(i.Name(), func(t *testing.T) {
			f := i.F()
			for _, j := range i.tests {
				assert.Equal(t, j.out, f(j.in))
			}
		})
	}
}
