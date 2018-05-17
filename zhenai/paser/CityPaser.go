package paser

import (
	"github.com/whuwzp/papapa/deffer"
	"regexp"
	"fmt"
)

func CityPaser(all []byte) deffer.PaserResult {
	re := regexp.MustCompile(CityString)
	match := re.FindAllStringSubmatch(string(all), -1)

	result := deffer.PaserResult{}
	for _, m := range match {
		fmt.Printf("City: %s\t\t\t | URL: %s\n", m[2], m[1])
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, deffer.Request{string(m[1]), MenPaser})
	}

	return result
}