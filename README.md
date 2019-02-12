# JWT Output Tool (JOT)

A simple utility for outputting the payload of a JSON Web Token.

## How to build it

```bash
$ make build
```

## Usage

```bash
$ ./bin/jot eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c

{
 "iat": 1516239022,
 "name": "John Doe",
 "sub": "1234567890"
}
```
