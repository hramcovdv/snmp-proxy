package snmp

import (
	"context"
	"fmt"
	"net"
	"time"

	g "github.com/gosnmp/gosnmp"
)

type RequestFunc func(ctx context.Context, r *SnmpRequest) ([]SnmpResponse, error)

type SnmpRequest struct {
	Oids      []string `schema:"oids,required"`
	Target    string   `schema:"target,required"`
	Community string   `schema:"community,default:public"`
	Version   int      `schema:"version,default:1"`
}

type SnmpResponse struct {
	Oid   string `json:"oid"`
	Type  string `json:"type"`
	Value any    `json:"value"`
}

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

func newSnmpResponse(pdu g.SnmpPDU) (r SnmpResponse) {
	r.Oid = pdu.Name
	r.Type = pdu.Type.String()

	switch pdu.Type {
	case g.OctetString:
		r.Value = asString(pdu.Value)
	case g.IPAddress:
		r.Value = asIpAddr(pdu.Value)
	default:
		r.Value = g.ToBigInt(pdu.Value)
	}

	return r
}

func asString(v any) string {
	s := string(v.([]byte))

	for _, c := range s {
		if c > 127 { // not ASCII
			return fmt.Sprintf("%x", v)
		}
	}

	return s
}

func asIpAddr(v any) string {
	b := v.([]byte)

	return net.IP(b).String()
}
