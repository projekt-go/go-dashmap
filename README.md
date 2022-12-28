# go-dashmap
Simple concurrent thread-safe hash map for Go

## Installation
```
go get -u github.com/projekt-go/dashmap
```

## Usage
use it like a normal map, but it's also safe for concurrent usage.
```go
mp := dashmap.New[string, int]()

mp.Put("fish", 2)

val, ok := mp.Get("fish")

fmt.Println(val) // 2

for entry := range mp.Entries() {
  fmt.Printf("key: %s, value: %d\n", entry.Key, entry.Value)
}
```
