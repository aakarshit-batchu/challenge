package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"challenge/models"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bytes"
)

var _ = Describe("Challenge", func() {
	var (
		url,successStatus string
		response models.Response
	)
	BeforeEach(func() {
		url, successStatus = "http://10.71.3.11:8080/v1/order", "Success"
	})

        Context("Order Items", func() {
                BeforeEach(func() {
			var jsonStr = []byte(`{"name":"aakarshit","address":"601","phonenumber":9790848677,"order":[{"item":"Idly","quantity":2},{"item":"Dosa","quantity":3}]}`)
                        req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
                        client := &http.Client{}
                        resp, _ := client.Do(req)
                        bodyBytes, _ := ioutil.ReadAll(resp.Body)
                        defer resp.Body.Close()
                        _ = json.Unmarshal(bodyBytes, &response)
        })
                It("Return Status should be Success", func() {
                        Expect(response.Status).Should(Equal(successStatus))
                })
        })

})
