


Test:

```sh
$ curl --location --request GET 'http://localhost:8842'
```

response:

```json

[{
    "Id":1,
    "Name":"GO ARCH",
    "Completed":true
}]
```


```sh
$ curl --location --request GET 'http://localhost:8842/?id=1'
```

response:

```json

{
    "Id":1,
    "Name":"GO ARCH",
    "Completed":true
}
```