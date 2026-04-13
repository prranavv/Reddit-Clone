<p align="center">
  <img src="logo.svg" alt="Reddit Clone Logo" width="180"/>
</p>

<h1 align="center">Reddit Clone</h1>
<p align="center"><strong>Full-Stack Social Platform Built with Go, HTMX & PostgreSQL</strong></p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go" />
  <img src="https://img.shields.io/badge/HTMX-1.9-blue?style=flat-square" />
  <img src="https://img.shields.io/badge/PostgreSQL-16-4169E1?style=flat-square&logo=postgresql" />
  <img src="https://img.shields.io/badge/chi--router-✓-00D18C?style=flat-square" />
  <img src="https://img.shields.io/badge/Deployed-Linode-green?style=flat-square" />
  <img src="https://img.shields.io/badge/License-MIT-green?style=flat-square" />
</p>

<p align="center">
  A server-rendered Reddit-style social platform built with Go's standard library and chi-router on the backend, HTMX for dynamic interactions on the frontend, and PostgreSQL for persistence. No heavy JavaScript frameworks — just HTML, CSS, and HTMX.
</p>

<p align="center">
  <a href="https://172-235-29-203.ip.linodeusercontent.com/"><strong>Live Demo →</strong></a>
</p>

---

## The Motivation

Most modern web apps default to React or a similar SPA framework, even when server-rendered HTML would be simpler, faster, and more maintainable. This project explores the alternative — a full-featured social platform built with Go templates and HTMX, where the server owns the rendering and the browser handles minimal interactivity. The result is a fast, lightweight app with almost zero client-side JavaScript.

## The Stack

| Layer | Technology | Why |
|:---:|---|---|
| **Backend** | Go + chi-router | High performance, minimal dependencies, expressive routing |
| **Frontend** | HTML + CSS + HTMX | Server-rendered pages with dynamic partial updates — no React, no build step |
| **Database** | PostgreSQL | Reliable relational storage for users, posts, comments, and votes |
| **Deployment** | Linode | Self-hosted VPS |

---

## Features

- **User authentication** — registration, login, and session management
- **Posts and comments** — create, read, and interact with threaded content
- **Voting** — upvote and downvote posts and comments
- **Server-rendered UI** — Go templates render full pages, HTMX handles partial updates without full page reloads
- **Structured logging** — JSON-formatted logs with Go's `slog` package
- **Database migrations** — schema management via Soda (Pop)
- **Zero frontend build step** — no npm, no webpack, no bundler — just plain HTML/CSS served directly

---

## Architecture

```
┌──────────────────────────────────────────────────────────────┐
│                         Browser                              │
│                                                              │
│   ┌─────────────────────────────────────────────────────┐    │
│   │  HTML + CSS + HTMX                                  │    │
│   │  Server-rendered pages                              │    │
│   │  HTMX swaps for dynamic updates (no full reload)    │    │
│   └────────────────────────┬────────────────────────────┘    │
└────────────────────────────┼─────────────────────────────────┘
                             │  HTTP (GET/POST + HTMX partials)
┌────────────────────────────┴─────────────────────────────────┐
│                    Go Backend (chi-router)                    │
│                                                              │
│  ┌──────────┐  ┌───────────┐  ┌──────────┐  ┌────────────┐  │
│  │  Routes  │  │ Handlers  │  │Templates │  │ Middleware  │  │
│  │          │  │           │  │          │  │            │  │
│  │ chi      │  │ Auth      │  │ Go html/ │  │ Session    │  │
│  │ router   │  │ Posts     │  │ template │  │ Logging    │  │
│  │          │  │ Comments  │  │ partials │  │ Recovery   │  │
│  │          │  │ Votes     │  │          │  │            │  │
│  └──────────┘  └───────────┘  └──────────┘  └────────────┘  │
│                                                              │
└────────────────────────────┬─────────────────────────────────┘
                             │  SQL
┌────────────────────────────┴─────────────────────────────────┐
│                      PostgreSQL                              │
│                                                              │
│  Users  │  Posts  │  Comments  │  Votes  │  Sessions         │
└──────────────────────────────────────────────────────────────┘
```

---

## Quick Start

### Prerequisites

- Go 1.21+
- PostgreSQL 14+
- [Soda CLI](https://gobuffalo.io/documentation/database/soda/) (for migrations)

### 1. Clone and configure

```bash
git clone https://github.com/prranavv/Reddit-Clon.git
cd Reddit-Clon
```

Create a `database.yml` with your Postgres credentials (see `database.yml.example` for the format).

Create a `.env` file:

```env
DATABASE_URL=host=localhost port=5432 dbname=dbname user=user password=pwd sslmode=disable
```

### 2. Run migrations

```bash
soda migrate up
```

### 3. Build and run

```bash
go build -o reddit cmd/web/*.go
./reddit
```

If everything is working, you'll see:

```json
{
  "time": "2024-01-30T12:39:57.558014061+05:30",
  "level": "INFO",
  "msg": "Connected to Database"
}
{
  "time": "2024-01-30T12:39:57.558167688+05:30",
  "level": "INFO",
  "msg": "Server is running on port 8080"
}
```

---

## Tech Stack

| Component | Technology | Purpose |
|---|---|---|
| **Language** | Go | Backend logic, templating, HTTP handling |
| **Router** | chi-router | Lightweight, composable HTTP routing with middleware support |
| **Frontend** | HTMX | Declarative partial page updates via HTML attributes |
| **Templating** | Go `html/template` | Server-side HTML rendering |
| **Database** | PostgreSQL | Relational data storage |
| **Migrations** | Soda (Pop) | Database schema versioning and migrations |
| **Logging** | Go `slog` | Structured JSON logging |
| **Deployment** | Linode VPS | Self-hosted production environment |

---

## FAQ's

**"Why Go instead of Node/Python/Ruby?"**
> Go compiles to a single static binary with no runtime dependencies, starts in milliseconds, and handles concurrent requests natively with goroutines. For a server-rendered app where the backend does all the heavy lifting, it's hard to beat.

**"Why HTMX instead of React?"**
> HTMX lets you build dynamic UIs by returning HTML fragments from the server instead of JSON. The server stays in control of rendering, there's no client-side state management, no build step, and no JavaScript bundle to ship. For a content-driven app like Reddit, this is simpler and faster.

**"Why chi-router instead of Gin or Echo?"**
> Chi is composable and middleware-friendly while staying close to Go's `net/http` standard library. It doesn't force you into a framework — it's just a router, which means less magic and more control.

**"Why PostgreSQL?"**
> Reddit-style apps are inherently relational — users have posts, posts have comments, comments have votes. Postgres handles these relationships naturally with foreign keys, joins, and indexes, and it's battle-tested at scale.

---

## Disclaimer

This is a personal learning project built to explore server-rendered web development with Go and HTMX. It is not intended as a production-grade social platform. Authentication, input validation, and security hardening have not been thoroughly audited.

---

## License

MIT — see [LICENSE](LICENSE) for details.

---

<p align="center">
  <sub>Built by <a href="https://github.com/prranavv">prranavv</a></sub>
</p>
