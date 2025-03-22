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
			H4(Text("Get Request")),
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
		P(oidsInput(".1.3.6.1.2.1.1.1.0")),
		P(targetInput("127.0.0.1")),
		P(communityInput("public")),
		P(versionSelect("1")),
		P(Button(Text("Submit"), Attr("type", "submit"))),
	)
}

func oidsInput(value string) Node {
	return Input(
		Attr("type", "text"),
		Attr("name", "oids"),
		Attr("placeholder", "OID"),
		Attr("required", ""),
		Attr("value", value),
	)
}

func targetInput(value string) Node {
	return Input(
		Attr("type", "text"),
		Attr("name", "target"),
		Attr("placeholder", "Target"),
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

func versionSelect(value string) Node {
	return Select(
		Attr("name", "version"),
		Option(Text("Version 1"), Attr("value", "1")),
		Option(Text("Version 2c"), Attr("value", "2")),
		Attr("value", value))
}
