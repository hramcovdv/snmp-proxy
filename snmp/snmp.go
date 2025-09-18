package snmp

import (
	"context"
	"time"

	g "github.com/gosnmp/gosnmp"
)

type RequestFunc func(ctx context.Context, r *SnmpRequest) ([]SnmpResponse, error)

func newParams(ctx context.Context, r *SnmpRequest) *g.GoSNMP {
	return &g.GoSNMP{
		Target:    r.Target,
		Port:      161,
		Community: r.Community,
		Version:   g.SnmpVersion(r.Version),
		Timeout:   3 * time.Second,
		Retries:   1,
		Context:   ctx,
	}
}

func Get(ctx context.Context, r *SnmpRequest) (res []SnmpResponse, err error) {
	params := newParams(ctx, r)

	err = params.Connect()
	if err != nil {
		return res, err
	}
	defer params.Conn.Close()

	result, err := params.Get(r.Oids)
	if err != nil {
		return res, err
	}

	for _, pdu := range result.Variables {
		res = append(res, newSnmpResponse(pdu))
	}

	return res, nil
}

func Walk(ctx context.Context, r *SnmpRequest) (res []SnmpResponse, err error) {
	params := newParams(ctx, r)

	err = params.Connect()
	if err != nil {
		return res, err
	}
	defer params.Conn.Close()

	for _, oid := range r.Oids {
		results, err := params.WalkAll(oid)
		if err != nil {
			return res, err
		}

		for _, pdu := range results {
			res = append(res, newSnmpResponse(pdu))
		}
	}

	return res, nil
}
