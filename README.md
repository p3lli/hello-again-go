# Hello again, go

## It's good to be back

Removing the rust...

### Playing with net/http and gg

Let's try to build a simple API REST Endpoint: A single GET with a text query param.  
As a result, it will return a .png image containing the value of the queryparam.  

I am going to use `net/http` as a very simplistic web framework and [`gg`](https://github.com/fogleman/gg) to create the .png  

#### Test

```
go test ./...
```

#### Run

```
go run main.go
```

Try:
```
curl -O "http://localhost:8080/image?text=go" > go.png
```

![Go text](go.png)
