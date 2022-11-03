package dns

type DNSRecord struct {
	Name string   `json:"name"`
	Type string   `json:"type"`
	TTL  string   `json:"ttl"`
	Data []string `json:"data"`
}

type DNSRecordsResoonse struct {
	Records   []DNSRecord `json:"recordSets"`
	PageToken *string     `json:"pageToken"`
}

type UpsertRecordsRequest struct {
	Replacements []DNSRecord `json:"replacements"`
}
