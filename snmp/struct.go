package snmp

type SnmpRequest struct {
	Oid       string `json:"oid"`
	Target    string `json:"target"`
	Community string `json:"community"`
}

type SnmpResponse struct {
	Oid   string      `json:"oid"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}
