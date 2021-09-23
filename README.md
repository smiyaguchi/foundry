# foundry
Generate fake data from schema

## Usage

```
$ make build
go build ./cmd/foundry

$ ./foundry gen -f spec.yaml
{"enable":false,"id":44,"ipv4":"178.23.106.51","ipv6":"fa75:6b15:7e13:ed00:6bd6:ae4b:d3dd:dd0a","mail":"example@exmaple.com","msg":{"body":"crNDAcT8KO","msg":{"body":"AGzk2CxrLU"}},"name":"","uuid":"a5b50fb2-2a42-4ead-83f4-ef23ea68258a"}

$ ./foundry gen -f spce.yaml -n 3
{"enable":true,"id":6,"ipv4":"12.224.105.95","ipv6":"d165:48d9:1364:7ec9:83b0:b035:6aee:a5b5","mail":"example@exmaple.com","msg":{"body":"cY69zbXyFB","msg":{"body":"2SbpAXwdpd"}},"name":"","uuid":"a65dd941-6bba-448c-8512-371b9023dad5"}{"enable":true,"id":84,"ipv4":"25.24.132.40","ipv6":"e03a:42a7:69e0:6c76:d22c:f269:4e2b:1ff8","mail":"example@exmaple.com","msg":{"body":"w3hBq3Fnx6","msg":{"body":"li31kSutUb"}},"name":"","uuid":"7eafb85a-9938-4590-9c36-5d4fc8d796df"}{"enable":false,"id":16,"ipv4":"188.247.192.19","ipv6":"ebd9:4005:62f7:6816:552b:25fb:a48d:4236","mail":"example@exmaple.com","msg":{"body":"33uUhEUbgM","msg":{"body":"JQyqZeUmV8"}},"name":"","uuid":"1ecf6196-95fd-4b28-b278-f7cfa235574a"}
```

## Spec file
The spec file is a file that defines the structure of the fake data. 
By selecting a generator, you can generate various types of data. It is also possible to set fixed values.

### Example

```
schema:
  id:
    type: int
  name:
    type: string
    value: ""
  mail:
    type: string
    value: example@exmaple.com
  ipv4:
    type: string
    gen: ipv4
  ipv6:
    type: string
    gen: ipv6
  uuid:
    type: string
    gen: uuid
  msg:
    schema:
      body:
        type: string
        msg:
          schema:
            body:
              type: string
  enable:
    type: bool
```

### support generator
The generator generates various types of fake data. The following types of generators are supported.
If not specified, the default generator will be used.

|name|type|description|
|:---|:--:|:----------|
|ipv4|`string`|generate random ipv4 address.|
|ipv6|`string`|generate random ipv6 address.|
|random|`int`|generate random num.|
|uuid|`string`|generate uuid|
