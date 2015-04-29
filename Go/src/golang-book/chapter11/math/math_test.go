package math

import "testing"

// go test command will run this because function starts with word Test and takes type of *testing.T
func TestAverage(t *testing.T) {
  var v float64
  // testing average of 1 and 2
  v = Average([]float64{1,2})
  // we know the average of 1,2 is 1.5 so if it doesn't equal 1.5 we got a problem
  if v != 1.5 {
    t.Error("Expected 1.5, got ", v)
  }
}