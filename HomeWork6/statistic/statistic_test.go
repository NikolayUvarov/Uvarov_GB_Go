package statistic

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}
var testssumm = []testpair{
	{[]float64{1, 2}, 1 + 2},
	{[]float64{1, 1, 1, 1, 1, 1}, 6},
	{[]float64{-1, 1}, 0},
}

func TestAverageSet(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestSummary(t *testing.T) {
	for _, pair := range testssumm {
		v := Summary(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestSummary2(t *testing.T) {
	if Summary([]float64{1, 2, 31}) != 34 {
		t.Error("Hello")
	}
}
