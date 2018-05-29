package xml

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"net/url"
	"crypto/md5"
	"encoding/hex"
	"sort"
	"encoding/xml"
)

type epay struct {
	XMLName     xml.Name `xml:"epay"`
	Retcode     string   `xml:"retcode"`
	Date        string   `xml:"date"`
	Count       string   `xml:"count"`
	Money       string   `xml:"money"`
	Partnerno   string   `xml:"partnerno"`
	Items       []Items  `xml:"items"`
	//Description string   `xml:",innerxml"`
}

type Items struct {
	XMLName xml.Name `xml:"items"`
	Item    []Item   `xml:"item"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	PlatTransno string   `xml:"plat_transno"`
}

func main() {
	params := url.Values{}
	params.Set("service", "account_ download_service")
	params.Set("partner", "12038348")
	params.Set("order_date", "20180423")
	params.Set("order_type", "0001")

	sortedKeys := make([]string, 0)
	for k := range params {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	var signStrings string
	for _, k := range sortedKeys {
		value := fmt.Sprintf("%v", params[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value[1:len(value)-1] + "&"
		}
	}
	signStrings = signStrings[0:len(signStrings)-1] + "1272eba5fe89402fb5f9f980aa210d62"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	sign := hex.EncodeToString(cipherStr)
	params.Set("sign", sign)
	resp, _ := http.PostForm("https://pay.srbank.cn/epaygate/account_download.htm",
		params)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	epay := epay{}
	xml.Unmarshal(body, &epay)
	fmt.Println(epay.Items[0].Item[0].PlatTransno)
}
