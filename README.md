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
  "domain": "koustavmondal.me"
}
```

- Response you will get :-

```json
{
  "status": "true",
  "error": "nil",
  "name_server": ["dorthy.ns.cloudflare.com.", "seth.ns.cloudflare.com."]
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
  "domain": "koustavmondal.me"
}
```

- Response you will get :-

```json
{
  "status": "",
  "error": "",
  "txt_record": [
    "v=spf1 redirect=_spf.facebook.com",
    "google-site-verification=A2WZWCNQHrGV_TWwKh6KHY90tY0SHZo_RnyMJoDaG0s",
    "google-site-verification=sK6uY9x7eaMoEMfn3OILqwTFYgaNp4llmguKI-C3_iA",
    "google-site-verification=wdH5DTJTc9AYNwVunSVFeK0hYDGUIEOGb-RReU6pJlY"
  ]
}
```
