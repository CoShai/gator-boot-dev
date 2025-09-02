
## Gator, built with Go, is a blog aggregator that pulls content from RSS feeds.

---

##  Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) 1.20+
- [PostgreSQL](https://www.postgresql.org/download/) 13+

### 1. Clone the Repository

```bash
git clone https://github.com/CoShai/gator-boot-dev.git
cd gator-boot-dev
```



### 2. Setup Environment

Copy ".gatorconfig.json" to ~/ and fill in your Postgres credentials:
```
{
  "db_url": "postgres://username:password@host:port/database?sslmode=disable"
}
```


### 3. Install or Build the App
```bash
go build
./gator
```
or

```bash
go install
gator
```

### 4. Available commands
Usage: cli <command> [args...]
```
gator register "name"
gator addfeed "exmaple.com/rss"
```


- help - show list of commands
```
"gator help"
```
- register - create user
```
"gator register <name>"
```
- login - login user
```
"gator login <name>"
```
- users - show registered users
```
"gator users"
```
- addfeed - add feed to database
```
"gator addfeed <name> <url>"
```
- follow - current user follow feed
```
"gator follow <url>"
```
- feeds - show feeds available
```
"gator feeds"
```
- following - show feeds current user follow
```
"gator following"
```
- agg - fetching feeds every time_between_reqs
```
"gator agg <time_between_reqs>"
```

