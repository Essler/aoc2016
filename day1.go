package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main(input string) {
  
  curr_dir := 'N'
	vector := [2]int{0, 0}
	
	moves := strings.Split(input, ", ")
	for _, move := range moves {
	  
	  move_array := strings.Split(move,"")
	  distance, err := strconv.Atoi(move_array[1])
	  if (err != nil) { fmt.Println(err) }
	  
	  switch move_array[0] {
	    
	    case "R":
	      switch curr_dir {
	        case 'N':
	          curr_dir = 'E'
	          vector[0] += distance
	        case 'E':
	          curr_dir = 'S'
	          vector[1] -= distance
	        case 'S':
	          curr_dir = 'W'
	          vector[0] -= distance
	        case 'W':
	          curr_dir = 'N'
	          vector[1] += distance
	        default:
	          panic("Non cardinal direction.")
	      }
	      
	    case "L":
	      switch curr_dir {
	        case 'N':
	          curr_dir = 'W'
	          vector[0] -= distance
	        case 'E':
	          curr_dir = 'N'
	          vector[1] += distance
	        case 'S':
	          curr_dir = 'E'
	          vector[0] += distance
	        case 'W':
	          curr_dir = 'S'
	          vector[1] -= distance
	        default:
	          panic("Non cardinal direction.")
	      }
	      
	    default:
	      panic("Not R or L.")
	  }
	}
	
	distance := intAbs(vector[0]) + intAbs(vector[1])
	fmt.Println(distance)
}

func intAbs(i int) int {
  
  if (i < 0) { i = -i }
  return i
}
