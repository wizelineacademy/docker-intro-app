package main

import (
	"fmt"
	"time"
)

func the_bug() {
  my_string := "I'm about to die"

  fmt.Println("hello from the bug...")

  time.Sleep(10 * time.Second)

  for i := 0; i < len(my_string) + 1; i += 1 {
    time.Sleep(500 * time.Millisecond)
      v := string(my_string[i])
      fmt.Println(v)
  }
}
