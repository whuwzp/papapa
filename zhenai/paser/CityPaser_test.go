package paser

import (
	"testing"
	"fmt"
	"io/ioutil"

)

func TestCityPaser(t *testing.T) {
	contents, err := ioutil.ReadFile("CityListData.html")
	if err != nil{
		panic(err)
	}
	result := CityPaser(contents)

	fmt.Printf("%d items\n %d url\n",len(result.Items), len(result.Requests))
}