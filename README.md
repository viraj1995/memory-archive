# memory-archive
Create your personal journal in a CLI and private database

To be made into an openclaw skill.

CLI name: memarc (commands like memarc add, memarc retrieve)
Core commands: add, retrieve, list, tag, export
Storage: local encrypted database (pluggable backend)
Search: full-text + tag filters
Config: user profile + key management
Future: OpenClaw skill integration via command wrappers

## Implementation (brief)
- Language: Go or TypeScript
- CLI framework: Cobra (Go) or Commander (TS)
- Storage: local encrypted SQLite by default; optional Supabase backend later
- Search: SQLite FTS for local; SQL + tags for remote
- Config: user profile + encryption keys in a local config file

## Data Flow & Storage Strategy

**Raw ingestion (immediate):**
- `memarc add <text>` stores full entry with timestamp as single indexed record
- No processing or breakdown at write time
- Enables fast, offline capture

**Batch processing (scheduled):**
- Daily/weekly job extracts metadata: dates, names, keywords, amounts
- Auto-tags based on content analysis (birthdays, deadlines, tasks, etc.)
- Populates structured `metadata` table while keeping raw `content` intact
- Full-text search indexed for retrieval

**Retrieval:**
- Search across raw content + tags + extracted metadata
- Filter by type, tags, date range, or keyword
- Export preserves both raw and structured data

**Storage backends:**
- Local: SQLite (raw + FTS indexes)
- Remote (future): Supabase sync with client-side encryption

## Implementation Details

### Architecture Overview

The project follows **clean architecture** with layered separation of concerns:

```
cmd/memarc/
  └─ main.go           # Entry point, CLI setup (Cobra)
  
internal/
  ├─ cli/              # Command handlers (business logic)
  │  ├─ add.go
  │  └─ list.go
  ├─ models/           # Data structures
  │  └─ entry.go
  └─ storage/          # Data access layer (GORM)
     └─ db.go
```

**Key Design Decisions:**
1. **Layered approach** — separates concerns (CLI ↔ logic ↔ database)
2. **Dependency injection** — database passed to handlers, not globals
3. **GORM ORM** — handles schema migrations and SQL automatically
4. **Cobra CLI framework** — built-in help, flags, and subcommand structure

### Models (`internal/models/entry.go`)

```go
type Entry struct {
    ID        uint      // Primary key
    Content   string    // Raw text entry
    Type      string    // Category (personal, professional, study, etc.)
    Tags      string    // Comma-separated tags
    CreatedAt time.Time // Timestamp
    UpdatedAt time.Time // Update tracking
}
```

**Purpose:** Defines the data structure for memory entries with GORM ORM tags for database mapping.

### Storage Layer (`internal/storage/db.go`)

Responsibilities:
- **Initialization**: Creates directory structure, opens SQLite connection
- **Migrations**: Auto-creates tables via `AutoMigrate()`
- **CRUD operations**:
  - `CreateEntry()` — insert new entry
  - `ListEntries()` — fetch all entries
  - `GetEntry(id)` — fetch single entry
  - `DeleteEntry(id)` — remove entry

**Key detail:** Creates `~/.memarc/memarc.db` directory automatically if missing.

### CLI Commands (`internal/cli/`)

**add.go** — Handler for `memarc add`
- Accepts: positional text argument, `--type` flag, `--tags` flag
- Stores entry with current timestamp
- Example: `memarc add "text" --type personal --tags birthday,important`

**list.go** — Handler for `memarc list`
- Fetches all entries from database
- Formats output with ID, timestamp, type, truncated content, tags
- Returns error if no entries exist

### Main Entry Point (`cmd/memarc/main.go`)

**Flow:**
1. Parse flags and set database path (defaults to `~/.memarc/memarc.db`)
2. Initialize database connection in `PersistentPreRun` (runs before any command)
3. Create Cobra commands with wrappers that pass `*storage.DB` instance
4. Execute command and handle errors

**Why this structure:** Database is initialized once, passed to all commands, and gracefully handles initialization errors.

### Dependencies

- **Cobra** — CLI framework with flag parsing, help text, command structure
- **GORM** — ORM for database abstraction and migrations
- **SQLite driver** — embedded database, no external service needed

## Getting Started

### Prerequisites
- Go 1.21 or higher
- SQLite3 (included with Go SQLite driver)

### Build

```bash
# Navigate to project directory
cd memory-archive

# Download dependencies
go mod download
go mod tidy

# Build the binary
go build -o bin/memarc ./cmd/memarc

# Verify build succeeded
./bin/memarc --version  # or --help
```

### Run

**Add a memory entry:**
```bash
./bin/memarc add "My memory text" -t personal -g birthday,travel
./bin/memarc add "Completed code review" -t professional -g review,coding
./bin/memarc add "Learned about concurrency" -t study -g golang
```

**List all entries:**
```bash
./bin/memarc list
```

**Get help:**
```bash
./bin/memarc --help
./bin/memarc add --help
./bin/memarc list --help
```

**Custom database location:**
```bash
./bin/memarc --db /path/to/custom.db add "text"
./bin/memarc --db /path/to/custom.db list
```

### Database Location
SQLite database stored at: `~/.memarc/memarc.db` (auto-created)

### Project Structure
```
memory-archive/
├── bin/                    # Compiled binaries
│   └── memarc
├── cmd/                    # Entry points
│   └── memarc/
│       └── main.go         # CLI setup and command routing
├── internal/               # Private packages
│   ├── cli/               # Command handlers
│   │   ├── add.go
│   │   └── list.go
│   ├── models/            # Data structures
│   │   └── entry.go
│   └── storage/           # Database layer
│       └── db.go
├── go.mod                 # Dependency manifest
├── .gitignore             # Git ignore rules
└── README.md              # This file
```

## Future Enhancements

### Phase 1 (Immediate)
- `retrieve` command (search by ID or keyword)
- `delete` command (remove entries)
- `tag` command (add/remove tags from existing entries)
- Better filtering in list (by type, date range, tags)

### Phase 2 (Medium term)
- Batch processor for metadata extraction (separate goroutine/cron job)
- Full-text search support
- Export command (JSON/CSV)
- Configuration file support

### Phase 3 (Long term)
- Client-side encryption for sensitive data
- Supabase backend integration with sync
- OpenClaw skill wrapper
- Web dashboard for retrieval and export

## Resume Highlights

- **Designed & implemented** a full-stack CLI application in Go from scratch
- **Applied clean architecture principles** with layered separation (CLI → business logic → database)
- **Leveraged GORM ORM** for database abstraction and automatic schema migrations
- **Integrated Cobra framework** for professional-grade CLI with flags, help text, and subcommands
- **Built extensible foundation** ready for future features (encryption, cloud sync, batch processing)
- **Demonstrated solid Go practices** — error handling, dependency injection, organized project structure