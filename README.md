## level-collection

A minimalistic organizer for your LevelDB database.

## Install

```
$ go get github.com/azer/level-collection
```

## Usage

Defining collections;

```go
import "github.com/azer/level-collection"

var (
	Users       = coll.New("users")
	Likes       = coll.NewChild("likes", Users)
	Posts       = coll.NewChild("posts", Users)
)

func init () {
  if err := coll.Open("/tmp/some-data"); err != nil {
    panic(err)
  }
}
```

Read, write and delete:

```go
Users.Select("azer@roadbeats.com").Write("foo") // or WriteByte([]byte("..."))
// SET users:azer@roadbeats.com TO foo

Users.Select("azer@roadbeats.com").Read() // or ReadByte()
// => foo

Posts.Select("azer@roadbeats.com", "this is a title").Write("and this is the post")
// SET users:azer@roadbeats.com:posts:this is a title TO and this is is the post

Posts.Select("azer@roadbeats.com", "yo").Write("lo")
Posts.Select("azer@roadbats.com", "yo").Delete()
```

Iterating:

```go
iter := Users.Select("azer@roadbeats.com", "").Iter()

for iter.Next() {
  fmt.Println(iter.Key())
  fmt.Println(iter.Value())
}

iter.Release()

if err := iter.Error(); err != nil {
  panic(err)
}
```

See `coll_test.go` for more info.

## API

* Open(path string)
* Set(key, value []byte)
* Get(key []byte)
* Delete(key []byte)
* NewColl(name string) *Coll
* NewChildColl(name string, parent *Coll) *Coll
* Coll.Key(fields ...string) string
* Coll.Select(fields ...string) *ReadWrite
* ReadWrite.Read() (string, error)
* ReadWrite.ReadByte() ([]byte, error)
* ReadWrite.Write(value string) (string, error)
* ReadWrite.WriteByte(value []byte) ([]byte, error)
* ReadWrite.Delete() error
* ReadWrite.Iter() iterator.Iterator

## Logging

You can enable logging by setting LOG=level-collection to output the logs.

```
$ LOG=level-collection go run my-program.go
```

More info about logging is at [github.com/azer/logger](http://github.com/azer/logger)