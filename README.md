# origin

A template application built on [pathless](https://github.com/timefactoryio/pathless).

```go
package main

import (
    "os"

    "github.com/timefactoryio/pathless"
)

func main() {
    p := pathless.NewPathless(os.Getenv("ORIGIN"), os.Getenv("CIRCUIT"))
    p.Home(os.Getenv("LOGO"), os.Getenv("TITLE"))
    p.Text(os.Getenv("TEXT"))
    p.Slides(os.Getenv("SLIDES"))
    p.Serve()
}
```

## Overview

Serves two endpoints:
- `:1000` — HTML shell
- `:1001` — wire gateway (binary API)

Content types:
- **Home**: logo and heading
- **Text**: markdown rendered from a URL or file
- **Slides**: image slideshow fetched from a URL or local directory

## Environment Variables

| Variable  | Description                                               |
| --------- | --------------------------------------------------------- |
| `ORIGIN`  | Bare domain for the HTML shell (e.g. `example.com`)       |
| `CIRCUIT` | Bare domain for the wire gateway (e.g. `api.example.com`) |
| `LOGO`    | URL or path to SVG logo                                   |
| `TITLE`   | Heading text                                              |
| `TEXT`    | URL or file path for markdown content                     |
| `SLIDES`  | URL or file path for slides payload                       |

Leave `ORIGIN` and `CIRCUIT` empty for dev — defaults to `localhost` with CORS `*`.

## Quick Start

```bash
docker compose -f docker-compose.dev.yml up
```

Runs on `http://localhost:1000`.

## Production

```yaml
environment:
  - ORIGIN=example.com
  - CIRCUIT=api.example.com
  - LOGO=https://your-bucket.s3.example.com/logo.svg
  - TITLE=your title
  - TEXT=https://raw.githubusercontent.com/your-org/repo/main/README.md
  - SLIDES=https://your-bucket.s3.example.com/slides
```