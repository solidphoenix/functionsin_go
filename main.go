package main


import (
  "fmt"
  "time"
  "math"
)

func areaOfSquare(x float64) float64 {
  return math.Pow(x,2)
}

func perimeterOfSquare(ar float64) (float64,string,float64) {
  space := "   "
  return math.Pow(ar,2),space,(ar*4)
}

func joinTwoStrings(prefix string, suffix string) string {
  return prefix + suffix
}

func printOutDate() {
  t := time.Now()
  fmt.Println(t)
}

func waitForIt(message string) {
  defer fmt.Println("Done!")
  fmt.Println("Waiting")
  fmt.Println("Waiting")
  fmt.Println("Waiting")
  fmt.Println(message)
}

func main() {
  printOutDate()
  waitForIt(joinTwoStrings("Hi", " there"))
  a := 5.0
  fmt.Println("Area   Perimeter")
  fmt.Println(perimeterOfSquare(a))

  
}
