package main

import (

	"github.com/whuwzp/papapa/deffer"
	"github.com/whuwzp/papapa/fetcher"
	"fmt"
	"github.com/whuwzp/papapa/zhenai/paser"
)

const InitUrl  =   "http://city.zhenai.com/"

func main() {
	Run(deffer.Request{InitUrl, paser.CityPaser})
}

func Run(seeds ...deffer.Request) {
	var REQS []deffer.Request
	REQS = append(REQS, seeds...)

	for len(REQS) > 0 {
		r := REQS[0]
		REQS = REQS[1:]

		body, err := fetcher.UrlFetcher(r.Url)
		if err != nil {
			fmt.Printf("fetching %s, error: %v", r.Url, err)
			continue
		}

		result := r.PaserFunc(body)

		REQS = append(REQS, result.Requests...)
		//fmt.Println(REQS)
		//for _, item := range result.Items{
		//	fmt.Println("items", item)
		//}

	}

}
