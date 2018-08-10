package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

type myTransport struct {
	// Uncomment this if you want to capture the transport
	// CapturedTransport http.RoundTripper
}

func get_rawtx(data string) string {
	var tx *types.Transaction
	rawtx, err := hex.DecodeString(data)
	if err != nil {
		return ""
	}
	rlp.DecodeBytes(rawtx, &tx)
	return tx.String()
}

func get_balance(data string) string {
	return data
}

func report(buf []byte, elapsed time.Duration, body []byte) {
	var result map[string]interface{}
	json.Unmarshal(buf, &result)

	method := result["method"]
	params := result["params"]
	value := params.([]interface{})
	fmt.Println(method, value)
	/*
		for key, value := range result {
			// Each value is an interface{} type, that is type asserted as a string
			fmt.Println(key, value)
		}
	*/
	info := ""
	if method == "eth_sendRawTransaction" && len(value) > 0 {
		tx := value[0].(string)
		if tx[:2] == "0x" || tx[:2] == "0X" {
			info = get_rawtx(tx[2:])
		}
	}
	log.Println("----------------New Request-------------------------\n",
		string(buf), info,
		"\n\tResponse Time", elapsed.Nanoseconds(),
		"\n\tResponse Body------\n", string(body),
		"\n------------------------------End----------------")

}

func (t *myTransport) RoundTrip(request *http.Request) (*http.Response, error) {

	buf, _ := ioutil.ReadAll(request.Body)
	//rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	request.Body = rdr2 // OK since rdr2 implements the

	start := time.Now()
	response, err := http.DefaultTransport.RoundTrip(request)

	/*
		for k, v := range request.Header {
			fmt.Printf("Header field %q, Value %q\n", k, v)
		} */
	if err != nil {
		print(string(buf), "\n\ncame in error resp here", err)
		return nil, err //Server is not reachable. Server not working
	}

	elapsed := time.Since(start)
	body, err := httputil.DumpResponse(response, true)
	if err != nil {
		print(string(buf), "\n\nerror in dumb response")
		// copying the response body did not work
		return nil, err
	}
	report(buf, elapsed, body)

	return response, err
}
