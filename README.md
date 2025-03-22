# SNMP Proxy

This is a simple service that uses HTTP request to get information from the device via SNMP.

## Get started

If you use Docker, you can build the container yourself:
```
git clone https://github.com/hramcovdv/snmp-proxy.git
cd snmp-proxy/
docker build --tag snmp-proxy:latest .
```

then run container:
```
docker run --name snmp-proxy -p 8080:8080 -d snmp-proxy:latest
```

## How to use

To perform a *GetRequest* use HTTP endpoint `/api/get` with POST requests:
```
curl -X POST \
-d 'oids=.1.3.6.1.2.1.1.1.0' \
-d 'oids=.1.3.6.1.2.1.1.5.0' \
-d 'target=192.168.0.1' \
-d 'community=public' \
http://localhost:8080/api/get
```

for *GetNextRequest* use `/api/walk`:
```
curl -X POST \
-d 'oids=.1.3.6.1.2.1.2.2.1.2' \
-d 'target=192.168.0.1' \
-d 'community=public' \
http://localhost:8080/api/walk
```

You can also use the web form to make requests http://localhost:8080/probe