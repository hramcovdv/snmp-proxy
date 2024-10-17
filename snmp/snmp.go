package snmp

import (
	"time"

	"github.com/gosnmp/gosnmp"
)

func newParams(target, community string) *gosnmp.GoSNMP {
	return &gosnmp.GoSNMP{
		Target:    target,
		Port:      161,
		Community: community,
		Version:   gosnmp.Version2c,
		Timeout:   time.Second * 5,
		Retries:   3,
	}
}

func Get(r *SnmpRequest) (resp []SnmpResponse, err error) {
	params := newParams(r.Target, r.Community)

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
	params := newParams(r.Target, r.Community)

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
