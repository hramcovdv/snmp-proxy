package snmp

func Get(r *SnmpRequest) (res []SnmpResponse, err error) {
	params := newParams(r)

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

func Walk(r *SnmpRequest) (res []SnmpResponse, err error) {
	params := newParams(r)

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
