package coll

import (
	"fmt"
)

var (
	Users       = NewColl("users")
	Likes       = NewChildColl("likes", Users)
	Posts       = NewChildColl("posts", Users)
	Comments    = NewChildColl("comments", Posts)
	PostsSeenBy = NewChildColl("seen-by", Posts)
)

func init() {
	if err := Open("/tmp/ada-test-data"); err != nil {
		panic(err)
	}
}

func ExampleKey() {
	fmt.Println(Users.Key("foo@bar.com"))
	// Output: users:foo@bar.com <nil>
}

func ExampleChildKey() {
	fmt.Println(Likes.Key("foo@bar.com", "123"))
	fmt.Println(Posts.Key("foo@bar.com", "314"))
	// Output: users:foo@bar.com:likes:123 <nil>
	// users:foo@bar.com:posts:314 <nil>
}

func ExampleGrandChildKey() {
	fmt.Println(Comments.Key("foo@bar.com", "314", "999888"))
	fmt.Println(PostsSeenBy.Key("foo@bar.com", "314", "123"))
	// Output: users:foo@bar.com:posts:314:comments:999888 <nil>
	// users:foo@bar.com:posts:314:seen-by:123 <nil>
}

func ExampleGetSet() {
	Users.Select("pi@number.com").Write("3.14")
	fmt.Println(Users.Select("pi@number.com").Read())

	Comments.Select("pi@number.com", "66055", "444111").Write("999.000")
	fmt.Println(Comments.Select("pi@number.com", "66055", "444111").Read())
	// Output: 3.14 <nil>
	// 999.000 <nil>
}

func ExampleIter() {
	Likes.Select("azer@roadbeats.com", "123").Write("999")
	Likes.Select("yo@roadbeats.com", "456").Write("888")
	Likes.Select("azer@roadbeats.com", "789").Write("777")

	iter := Users.Select("azer@roadbeats.com", "").Iter()

	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		fmt.Println(key)
		fmt.Println(value)
	}

	iter.Release()

	fmt.Println(iter.Error())
	// users:azer@roadbeats.com:likes:123
	// 999
	// users:azer@roadbeats.com:likes:789
	// 777
	// Output: <nil>
}
