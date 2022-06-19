package mock

import (
	"bingFood-Merchant/entity/prod/req"
	"bingFood-Merchant/utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

var prodReq = `{
    "prod": {
        "prodName": "【扬州炒饭】",
        "shopId": %v,
        "oriPrice": 0,
        "packingFee": 1,
        "price": 0,
        "pic": "",
        "imags": "",
        "description": "足时发酵，老坛炒饭",
        "content": "",
        "prodStatus": 0,
        "isAutoAdd": 0,
        "score": 1,
        "buyLimit": 0,
        "todaySoldOut": 0,
        "categoryId": 0,
        "properties": [
            {
                "propName": "麻辣",
                "shopId": %v,
            },
            {
                "propName": "香甜",
                "shopId": %v,
            }
        ],
        "skus": [
            {
                "skuName": "大份量%s",
                "price": 15,
                "weight": 1000,
                "sellStatus": 0,
                "stock": 100
            },
            {
                "skuName": "小份量%s",
                "price": 10,
                "weight": 1000,
                "sellStatus": 0,
                "stock": 50
            }
        ]
    },
    "shopId": %v
}`

func TestProd(t *testing.T) {
	for i := 1; i < 4; i++ {
		shopId := i
		req := fmt.Sprintf(prodReq, shopId, shopId, shopId, strconv.Itoa(shopId), strconv.Itoa(shopId), shopId)
		PostWithData(req)
	}
}

func PostWithData(data string) {
	client := &http.Client{}
	var d = strings.NewReader(data)
	req, err := http.NewRequest("POST", HostPort+"/shop/prod/add", d)
	fmt.Println(data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

func TestList(t *testing.T) {
	fmt.Println(utils.ToJsonString(req.ListProdReq{}))
}
