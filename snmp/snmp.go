package snmp

import (
	"time"

	"github.com/gosnmp/gosnmp"
)

func newParams(r *SnmpRequest) *gosnmp.GoSNMP {
	return &gosnmp.GoSNMP{
		Target:    r.Target,
		Port:      161,
		Community: r.Community,
		Version:   gosnmp.SnmpVersion(r.Version),
		Timeout:   5 * time.Second,
		Retries:   3,
	}
}

func Get(r *SnmpRequest) (resp []SnmpResponse, err error) {
	params := newParams(r)

	err = params.Connect()
	if err != nil {
		return resp, err
	}
	defer params.Conn.Close()

	result, err := params.Get([]string{r.Oid})
	if err != nil {
		return resp, err
	}

	for _, pdu := range result.Variables {
		resp = append(resp, getSnmpResponse(pdu))
	}

	return resp, nil
}

func Walk(r *SnmpRequest) (resp []SnmpResponse, err error) {
	params := newParams(r)

	err = params.Connect()
	if err != nil {
		return resp, err
	}
	defer params.Conn.Close()

	result, err := params.WalkAll(r.Oid)
	if err != nil {
		return resp, err
	}

	for _, pdu := range result {
		resp = append(resp, getSnmpResponse(pdu))
	}

	return resp, nil
}
