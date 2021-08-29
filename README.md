# foundry
Generate fake data from schema

## Usage

```
$ make build
go build ./cmd/foundry

$ ./foundry gen -f spec.yaml
[{"enable":true,"id":74,"ipv4":"251.52.163.82","ipv6":"85b8:380c:44ae:63ba:2ba3:f4f0:eadb:8c5c","mail":"example@exmaple.com","msg":{"body":"ZICaW7gsqc","msg":{"body":"gIf33JMhNB"}},"name":"","uuid":"e8c66407-6b21-43e9-995e-7a4c699e8a8f"}

$ ./foundry gen -f spce.yaml -n 3
[{"enable":true,"id":39,"ipv4":"108.93.4.122","ipv6":"6a60:3ba7:43fa:f9a2:2e16:b80f:9bbd:8e81","mail":"example@exmaple.com","msg":{"body":"w6Kz4DVTIc","msg":{"body":"ctlV8UCnxD"}},"name":"","uuid":"aa575e94-8861-4720-8569-04d41210d177"},
{"enable":false,"id":73,"ipv4":"169.101.238.116","ipv6":"2872:ab3b:7bc:8559:7d87:7296:85bc:3790","mail":"example@exmaple.com","msg":{"body":"FCo1uvuIei","msg":{"body":"iuHabBvuYs"}},"name":"","uuid":"0081f577-5674-4462-a477-79ec64cff658"},
{"enable":false,"id":37,"ipv4":"21.135.56.195","ipv6":"85d0:4919:7250:b008:9249:21d8:8ac1:ec17","mail":"example@exmaple.com","msg":{"body":"pK85dDH91a","msg":{"body":"JHaa3mPVO5"}},"name":"","uuid":"646650ed-5306-4466-a8b6-6516cb0f0966"}]
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
