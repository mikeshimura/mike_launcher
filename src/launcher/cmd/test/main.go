package main

import (
"launcher/util"
)

func main(){
	err := util.Unzip("./test/package.zip","./test")
	if err != nil {
		panic(err)
	} 
}