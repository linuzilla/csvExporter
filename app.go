package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		reader, err := New(arg)

		if err != nil {
			log.Fatal(err)
		}

		err = reader.Each(func(dataMap map[int]string) bool {
			dataSize := len(dataMap)

			for i := 0; i < dataSize; i++ {
				if value, ok := dataMap[i]; ok {
					fmt.Print(value)
				}
				if i < dataSize-1 {
					fmt.Print("\t")
				}
			}
			fmt.Print("\r\n")
			return true
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}
