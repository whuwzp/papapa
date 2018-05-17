package fetcher

import (
	"bufio"
	"errors"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func UrlFetcher(url string) ([]byte, error) {
	resp, err := http.Get(
		url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//fmt.Println("error! statuscode = ", resp.StatusCode)
		return nil, errors.New("error! statuscode...")
	}

	reader := bufio.NewReader(resp.Body)
	utf8Reader := transform.NewReader(reader, DetermineEncoding(reader).NewDecoder())

	R, e := ioutil.ReadAll(utf8Reader)
	//fmt.Println(string(R))
	return R, e
}

func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
