package dns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Olegas/dnstools/pkg/config"
	"github.com/Olegas/dnstools/pkg/keys"
)

func UpdateRecord(ip string, dnsZoneId string) bool {
	var request UpsertRecordsRequest

	addr := make([]string, 1)
	addr[0] = ip
	recordName, recordType, ttl := config.GetDNSRecord()
	record := DNSRecord{
		Name: recordName,
		Type: recordType,
		TTL:  ttl,
		Data: addr,
	}
	request.Replacements = append(request.Replacements, record)
	body, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("https://dns.api.cloud.yandex.net/dns/v1/zones/%s:upsertRecordSets", dnsZoneId)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", keys.GetIAMToken()))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}

	if resp.StatusCode != 200 {
		result, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s: %s", resp.Status, result)
		return false
	}

	return true
}
