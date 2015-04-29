package math

import "testing"

// set up a test struct to represent inputs and outputs for the function
type testpair struct {
  values []float64
  average float64
}

// this is a list of struct pairs
var tests = []testpair{
  { []float64{1,2}, 1.5},
  { []float64{1,1,1,1,1,1}, 1},
  { []float64{-1,1}, 0},
}

// go test command will run this because function starts with word Test and takes type of *testing.T
func TestAverage2(t *testing.T) {
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
