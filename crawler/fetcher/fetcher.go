package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet,
		url, nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}
	defer resp.Body.Close()

	//s, err := httputil.DumpResponse(resp, true)
	//if err != nil {
	//	return nil, err
	//}

	bufReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bufReader)
	utf8Reader := transform.NewReader(bufReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}


func determineEncoding(r *bufio.Reader) encoding.Encoding   {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetch error:%v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
