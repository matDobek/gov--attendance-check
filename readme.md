### Installation

* install `go`

```
# example using asdf

asdf plugin add golang
asdf install 1.21.4
```

### Changelog

```go
//====================
// 30/11/2023
//====================
//
// * add predicates.Between(val, from, to)
// * add constraints package, for additional Type Constraints (e.g constraints.Number)
// * add enumerable package, allowing to Map and Filter over the same type
//

enumerable.New([]int{0, 1, 2}).Map(func(v int) int{
    return v + 1
}).Map(func(v int) int{
    return v * v
}).Filter(func(v int) bool{
    return v % 2 == 0
}).Do() // {4}

```
