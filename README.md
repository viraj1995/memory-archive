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