# DNS-Lookup API

> Made using Go and Fiber

### Prerequisite before running the application

> Create .env in root dir of the project and copy all the variables from .env.example and past in .env and assign the values of all the variables.

### How to run the application

- To run the application in development mode:

```shell
go run .
```

- To make a build of the application:

```shell
go build -o dns
```

- To make a build of the application for windows:

```shell
go build -o dns.exe
```

<br>

### How to Get NS Record

- Request sent :-

```json
{
  "domain": "google.com"
}
```

- Response you will get :-

```json
{
  "status": "true",
  "error": "nil",
  "name_server": [
    "ns2.google.com.",
    "ns3.google.com.",
    "ns4.google.com.",
    "ns1.google.com."
  ]
}
```

<br>

### How to Get CNAME Record

- Request sent :-

```json
{
  "domain": "m.google.com"
}
```

- Response you will get :-

```json
{
  "status": "true",
  "error": "nil",
  "cname": "mobile.l.google.com."
}
```

<br>

### How to Get TXT Record

- Request sent :-

```json
{
  "domain": "google.com"
}
```

- Response you will get :-

```json
{
  "status": "true",
  "error": "nil",
  "txt_record": [
    "docusign=1b0a6754-49b1-4db5-8540-d2c12664b289",
    "google-site-verification=TV9-DBe4R80X4v0M4U_bd_J9cpOJM0nikft0jAgjmsQ",
    "v=spf1 include:_spf.google.com ~all",
    "onetrust-domain-verification=de01ed21f2fa4d8781cbc3ffb89cf4ef",
    "docusign=05958488-4752-4ef2-95eb-aa7ba8a3bd0e",
    "facebook-domain-verification=22rm551cu4k0ab0bxsw536tlds4h95",
    "globalsign-smime-dv=CDYX+XFHUw2wml6/Gb8+59BsH31KzUr6c1l2BPvqKX8=",
    "MS=E4A68B9AB2BB9670BCE15412F62916164C0B20BB",
    "google-site-verification=wD8N7i1JTNTkezJ49swvWW48f8_9xveREV4oB-0Hf5o",
    "webexdomainverification.8YX6G=6e6922db-e3e6-4a36-904e-a805c28087fa",
    "atlassian-domain-verification=5YjTmWmjI92ewqkx2oXmBaD60Td9zWon9r6eakvHX6B77zzkFQto8PQ9QsKnbf4I",
    "apple-domain-verification=30afIBcvSuDV2PLX"
  ]
}
```

<br>

### How to Get MX Record

- Request sent :-

```json
{
  "domain": "google.com"
}
```

- Response you will get :-

```json
{
  "status": "true",
  "error": "nil",
  "mx": ["smtp.google.com."]
}
```
