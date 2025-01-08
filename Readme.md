## Steps to install

### Step 1: Clone the Repository
```bash
git clone https://github.com/yourusername/your-repo.git
cd your-repo
```

### Step 2: Install dependencies
```bash
    go mod tidy
```

### Step 3: Create .env file and have following variables in it
```bash
    DB_HOST=your-db-host
    DB_USERNAME=your-db-user
    DB_PASSWORD=your-db-password
    DB_NAME=your-db-name
```

### Step 4: Run the project
```bash
    go run main.go
```

It will start server port 8000

## Test Api:
### POST `/api/v1/operation`

This API endpoint is used to schedule a job with a given timestamp and URL.

**Request Body:**
```json
{
    "url": "string",
    "timestamp": 1736332532
}
```

**Response**
```json
    {
        "success": "bool",
        "message": "string",
        "data": "any",
        "error": "any"
    }
```