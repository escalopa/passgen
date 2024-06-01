# passgen 🔐

PassGen: A CLI tool to generate secure passwords with specified options, including length, special characters, number of passwords, and output format and much more.

## Installation 🏗

### Using Go

You can install the tool using Go:

```bash
go install github.com/escalopa/passgen@v0.0.2
passgen generate -l 32
```

### Using Docker

```bash
docker pull dekuyo/passgen:v0.0.2
docker run --rm dekuyo/passgen:v0.0.2 generate -l 32
```

## Configuration

The application uses a configuration file named `.passgen` in one of the following formats: `.toml`, `.json`, `.yaml`, `.yml`.

Here are the available configuration options:

| Option      | Description                                           | Type   |
|-------------|-------------------------------------------------------|--------|
| length      | The length of the generated password                  | int    |
| iterations  | The number of passwords to generate                   | int    |
| characters  | The characters to use for password generation        | string |
| clipboard   | Whether to copy the generated password to clipboard  | bool   |

Here is an example of a `.passgen` configuration file:

```yaml
generate:
  length: 16
  iterations: 1
  characters: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  clipboard: true
```

## Milestones

- [x] Generate passwords with specified length & characters
- [ ] Fix merge config files (correct priority: flags > config file > default)
- [ ] Remove the hash & salt from output and just print the password

## Contributing

Contributions are welcome! Please open an issue or submit a pull
