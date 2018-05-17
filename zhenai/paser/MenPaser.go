package paser

import (
	"github.com/whuwzp/papapa/deffer"
	"fmt"
	"regexp"
)

func MenPaser(all []byte) deffer.PaserResult {
	//find people
	re1 := regexp.MustCompile(MenString)
	match1 := re1.FindAllStringSubmatch(string(all), -1)

	result := deffer.PaserResult{}
	for _, m := range match1 {
		//fmt.Printf("Name: %s\t\t\t | URL: %s\n", m[2], m[1])
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, deffer.Request{string(m[1]), ManPaser})
	}

	//find next page
	re2 := regexp.MustCompile(NextMen)
	match2 := re2.FindAllStringSubmatch(string(all), -1)
	for _, m := range match2 {
		fmt.Println("next page:" , m[1])
		result.Items = append(result.Items, nil)
		result.Requests = append(result.Requests, deffer.Request{string(m[1]), MenPaser})
	}

	//if len(match2) <= 0{
	//	fmt.Println(string(all))
	//}

	return result
}

//<a href="http://album.zhenai.com/u/105942278" target="_blank">天地</a>