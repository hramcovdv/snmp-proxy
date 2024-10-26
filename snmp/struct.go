package snmp

type SnmpRequest struct {
	Oid       string `schema:"oid,required"`
	Target    string `schema:"target,required"`
	Community string `schema:"community,default:public"`
	Version   int    `schema:"version,default:1"`
}

type SnmpResponse struct {
	Oid   string      `json:"oid"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type RequestFunc func(r *SnmpRequest) ([]SnmpResponse, error)
