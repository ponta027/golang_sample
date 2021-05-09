# Networking


## httpget

httpget.go に誤字あり


```
resp.Close()
```
↓

```
resp.Body.Close()
```

https://golang.org/pkg/net/http/#Response






