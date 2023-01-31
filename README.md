# Supabase-Go 🚀

An unofficial client library for [Supabase](https://supabase.io/) written in Go.
It facilitates authentication and database querying without relying on GoTrue or any Postgresql library.

```go
func main() {
	ctx := context.Background()

	// Create a new Supabase client
	client := supabase.NewClient("<SUPABASE_URL>", "<SUPABASE_ANON_KEY>")

	// Sign in with email and password
	if err := client.SignIn(ctx, "<EMAIL>", "<PASSWORD>"); err != nil {
		log.Fatalf("Failed to sign in: %s", err)
	}

	// Query the database as the authenticated user
	var messages []Message
	if err := client.From("messages").Select("*").To(&messages).Execute(ctx); err != nil {
		log.Fatalf("Failed to retrieve messages: %s", err)
	}

	log.Printf("Successfully retrieved %d messages", len(messages))
}
```

## 🎉 Features

- Compatible with Supabase API 💻
- Fast and efficient 🚀
- Supports authentication 🔒
- Supports querying and filtering data 🔎

## 💻 Installation

To install Supabase-Go, simply run the following command:

```bash
go get github.com/francoisdtm/supabase-go
```

## 📚 Examples

Here are a few examples to get you started with Supabase-Go:

```go
... // TODO
```

## 🤝 Contributing

We welcome contributions to Supabase-Go! If you find a bug or want to request a
new feature, please open an issue. If you want to contribute code, fork the
repository and create a pull request.
