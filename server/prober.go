package server

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func probePage() Node {
	return HTML(
		Head(
			Title("SNMP Probe"),
		),
		Body(
			H4(Text("SNMP Probe")),
			probeForm("/api/get", "post"),
			Hr(),
			H4(Text("Walk Request")),
			probeForm("/api/walk", "post"),
		),
	)
}

func probeForm(action, method string) Node {
	return Form(
		Attr("action", action),
		Attr("method", method),
		P(oidInput("")),
		P(hostnameInput("127.0.0.1")),
		P(communityInput("public")),
		P(Button(Text("Submit"))),
	)
}

func oidInput(value string) Node {
	return Input(
		Attr("type", "text"),
		Attr("name", "oid"),
		Attr("placeholder", "OID"),
		Attr("required", ""),
		Attr("value", value),
	)
}

func hostnameInput(value string) Node {
	return Input(
		Attr("type", "text"),
		Attr("name", "hostname"),
		Attr("placeholder", "Hostname"),
		Attr("required", ""),
		Attr("value", value),
	)
}

func communityInput(value string) Node {
	return Input(
		Attr("type", "text"),
		Attr("name", "community"),
		Attr("placeholder", "Community"),
		Attr("required", ""),
		Attr("value", value),
	)
}
