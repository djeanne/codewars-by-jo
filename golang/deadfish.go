package kata

import "strings"

func Parse(data string) []int {
  steps := strings.Split(data, "")    
  output := make([]int, 0)

  x := 0

  for _, step := range steps {
    switch(step) {
    case "i":
      x++
    case "d":
      x--
    case "s":
      x = x * x
    case "o":
      output = append(output, x)
    default:
      continue  
      }
  } 
  return output
}
