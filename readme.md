### Installation

* install `go`

```
# example using asdf

asdf plugin add golang
asdf install 1.21.4
```

* proxy to avoid CORS

```
npm install -g local-cors-proxy
lcp --proxyUrl http://localhost:8080/
```

* live code releading for go server is handled by `air`

```
# to install
go install github.com/cosmtrek/air@latest

# to start, go to directory with go server and run:
air

# if binary is not found you may need to add it to the PATH
# or if you're using tools like `asdf`:
asdf reshim
asdf exec air
```
