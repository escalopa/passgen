# passgen 🔐

PassGen: A CLI tool to generate secure passwords with specified options, including length, special characters, number of passwords, and output format and much more.

## Installation 🏗

### Using Go

You can install the tool using Go:

```bash
go install github.com/escalopa/passgen@latest
passgen --help
```

### Using Docker

```bash
docker pull dekuyo/passgen:latest
docker run --rm dekuyo/passgen:latest --help
```

## Usage

### Generating a Password

You can generate passwords using the `generate` command. The command takes various parameters:

| Flag            | Shorthand | Description                                                  | Default Value |
|-----------------|-----------|--------------------------------------------------------------|---------------|
| `--length`      | `-l`      | Length of the password                                       | 12            |
| `--use-signs`   | `-s`      | Include special characters in the password                  | false         |
| `--hash-times`  | `-t`      | Number of times to hash the password with Argon2            | 1             |
| `--num-passwords` | `-n`    | Number of passwords to generate                             | 1             |
| `--output`      | `-o`      | Output format (json, table, raw)                            | json          |
| `--custom-chars`| `-c`      | Custom characters to use for password generation            | ""            |

#### Examples

Generate a password of length 16:

```bash
passgen --length 16
```

Generate 5 passwords of length 16 with special characters and hash them 5 times:

```bash
passgen --length 16 --use-signs --hash-times 5 --num-passwords 5
```

Generate a password with custom characters and output in table format:

```bash
passgen --custom-chars "abcdef123$%" --output json
```

### Command Help

You can get help for the commands by using the `--help` flag:

```bash
task help
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull
