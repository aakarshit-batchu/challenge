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
		url, successStatus = "http://0.0.0.0:8080/v1/inventory", "Success"
	})
	Context("List Inventory", func() {
		BeforeEach(func() {
	                req, _ := http.NewRequest("GET", url, nil)
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

        Context("Add items to Inventory", func() {
                BeforeEach(func() {
			var jsonStr = []byte(`[{"item":"Dosa","price":40,"category":"SouthIndian"},{"item":"Idly","price":30,"category":"SouthIndian"}]`)
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

        Context("Add items to Inventory", func() {
                BeforeEach(func() {
                        var jsonStr = []byte(`[{"item":"Dosa","price":65,"category":"SouthIndian"},{"item":"Poori","price":40,"category":"SouthIndian"}]`)
                        req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
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

        Context("Add items to Inventory", func() {
                BeforeEach(func() {
                        var jsonStr = []byte(`[{"item":"Poori"}]`)
                        req, _ := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
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
