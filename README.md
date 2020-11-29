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
curl "http://localhost:8080/image?text=go" --output go.png
```

![Go image](go.png)

### Docker

Muli-stage build is used. See `docker-files` directory.  
Use `./build.sh` to automatically build the docker image. 

#### `golangcli-lint`
During docker image building, before the tests run  
I have added `golangcli-lint` for the first time  
as code quality measurement.  

Still studying configurations... sorry
