# Spark
A Logger library implementation in Golang.

## Installation
````
go get github.com/anjulapaulus/spark
````

## Implementation
````
package main

import "github.com/anjulapaulus/spark"

func main(){
    user := "anjula paulus"
	log := spark.NewLogger("INFO",true,true)
	log.PrefixedInfo("how are you",user)
}
````

## Future Implementations
* Batch Processing for logging