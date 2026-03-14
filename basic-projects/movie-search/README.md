# Movie Search with Database Caching

CLI tool to search movies using OMDb API with SQLite caching to avoid redundant API calls.

---

## Setup

### 1. Get API Key
- Visit: http://www.omdbapi.com/apikey.aspx
- Select FREE tier (1000 requests/day)
- Verify email and get your API key


### 2. Configure
Create `config.json`:
```json
{
  "omdb_api_key": "YOUR_API_KEY_HERE",
  "db_name": "movies.db"
}
```

### 3. Install Dependencies
```bash
go mod tidy
```

---

## Usage

### Search Movie
```bash
go run . --title=matrix
```

**First search:** Fetches from API and caches to database  
**Subsequent searches:** Returns from cache (faster, no API call)

### Examples
```bash
go run . --title=inception
go run . --title="the dark knight"
go run . --title=MATRIX  # Case-insensitive
```

---

## How It Works

**Flow:**
```
1. Check SQLite database for movie
2. If found → Return cached data 
3. If not found → Fetch from OMDb API
4. Save to database for future
5. Return movie data
```

**Database:** Automatically created at `data/movies.db` on first run

---

## Project Structure

```
movie-search/
├── main.go           # Entry point, CLI flags
├── movies.go         # API client, config loading
├── db-sqlite.go      # Database operations, caching logic
├── config.json       # API key (gitignored)
└── data/
    └── movies.db     # SQLite database (auto-created)
```

---

## Key Learnings

- **Config Management:** JSON config with `json.Unmarshal`
- **HTTP Client:** `net/http` with URL parameters
- **JSON Decoding:** Stream decoding with `json.NewDecoder`
- **SQLite:** Database operations with `database/sql`
- **Prepared Statements:** SQL injection prevention with `?` placeholders
- **Caching Strategy:** Check cache before external API call
- **Error Handling:** Graceful error propagation

---

## Notes

- Database and table created automatically on first run
- No need to manually create `data/` directory
- API key stored in `config.json` (keep secure, don't commit)
- Cache persists between runs
