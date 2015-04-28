package math
// the package name above is nested in a folder "math". the name of the package and folder are the same.

// Finds the average of a series of numbers
// capital letter functions like "Average" means that other packages and programs are able to see this
// if we used a lowercase function name our main program will not be able to see it
// hiding functions allows us to freely change those later without worrying about breaking other programs
func Average(xs []float64) float64 {
  total := float64(0)
  for _, x := range xs {
    total += x
  }
  return total / float64(len(xs))
}