# id
generate id - [snowflake](https://github.com/sony/sonyflake), [pika](https://github.com/hopinc/pika), [stripe-id](https://clerk.dev/blog/generating-sortable-stripe-like-ids-with-segment-ksuids) mix

#### get the package
```shell
> go get github.com/beyazit/id
```

#### example
```go
func main() {
	generator := id.New([]*id.PrefixRecord{
		{
			Prefix:      "user",
			Description: "User ID",
			Secure:      false,
		},
		{
			Prefix:      "bearer",
			Description: "Bearer token",
			Secure:      true,
		},
	}, sonyflake.Settings{})

	fmt.Println(generator.Generate("user"))
	fmt.Println(generator.Generate("bearer"))
}
```
