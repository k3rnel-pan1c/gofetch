# HTTP Logging Server

A CLI HTTP server that logs all incoming requests in detail. Useful for capturing and inspecting HTTP traffic, including POST bodies and headers for exfiltrating secrets during penetration testing exercises.

## Capabilities

- Logs timestamp, client IP, HTTP method, path, headers, and body for every request
- Supports all HTTP methods (GET, POST, etc.)
- Logs to stdout by default, optionally to a file via `--log-file`

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--hostname` | `127.0.0.1` | Hostname to bind to |
| `--port` | `8080` | Port to listen on |
| `--log-file` | _(none)_ | Optional log file path |

## Usage

### Without Docker

Build:

```
cd src
go build -o gofetch .
```

Run on `0.0.0.0:4444`:

```
./gofetch --hostname 0.0.0.0 --port 4444
```

Log to a file:

```
./gofetch --hostname 0.0.0.0 --port 4444 --log-file httpd.log
```

### With Docker

Build:

```
docker build --rm -t gofetch .
```

Run on `0.0.0.0:4444`:

```
docker run --rm -it -p 4444:4444 gofetch /app/gofetch --hostname 0.0.0.0 --port 4444
```

Log to a file:

```
docker run --rm -it -p 4444:4444 -v "$PWD":/logs gofetch /app/gofetch --hostname 0.0.0.0 --port 4444 --log-file /logs/httpd.log
```

### Test

Send a test POST request:

```
curl -X POST http://127.0.0.1:4444/test -d "secret=password123"
```
