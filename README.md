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
	Likes       = Users.NewChild("likes", Users)
	Posts       = Users.NewChild("posts", Users)
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

Users.Select("azer@roadbeats.com").Attr("yo").Write("lo")
// SET users:azer@roadbeats.com:yo TO lo

Users.Select("azer@roadbats.com").Attr("yo").Delete()
// DELETE users:azer@roadbeats.com:yo
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
* New(name string) *Coll
* NewChild(name string, parent *Coll) *Coll
* Coll.Key(fields ...string) string
* Coll.Select(fields ...string) *ReadWrite
* Coll.NewChild(name string) *Coll
* ReadWrite.Read() (string, error)
* ReadWrite.ReadByte() ([]byte, error)
* ReadWrite.Write(value string) (string, error)
* ReadWrite.WriteByte(value []byte) ([]byte, error)
* ReadWrite.Delete() error
* ReadWrite.Iter() iterator.Iterator
* ReadWrite.Attr(name string) *ReadWrite

## Logging

You can enable logging by setting LOG=level-collection to output the logs.

```
$ LOG=level-collection go run my-program.go
```

More info about logging is at [github.com/azer/logger](http://github.com/azer/logger)
