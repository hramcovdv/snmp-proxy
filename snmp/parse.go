package snmp

import (
	"fmt"

	"github.com/gosnmp/gosnmp"
)

func getSnmpResponse(pdu gosnmp.SnmpPDU) (r SnmpResponse) {
	r.Oid = pdu.Name
	r.Type = pdu.Type.String()

	switch pdu.Type {
	case gosnmp.OctetString:
		bytes := pdu.Value.([]byte)
		r.Value = asString(bytes)
	case gosnmp.IPAddress:
		r.Value = fmt.Sprintf("%v", pdu.Value)
	default:
		r.Value = gosnmp.ToBigInt(pdu.Value)
	}

	return r
}

func asString(b []byte) string {
	s := string(b)

	for _, c := range s {
		if c > 127 { // if not ASCII
			return fmt.Sprintf("%x", s)
		}
	}

	return s
}
