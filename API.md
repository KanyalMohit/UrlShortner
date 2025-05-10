# URL Shortener API Documentation

## Base URL
```
http://localhost:8080
```

## Endpoints

### 1. Create Short URL
Creates a new short URL for the given original URL.

```
POST /api/urls
```

#### Request
```json
{
    "original_url": "https://www.example.com/very/long/url"
}
```

#### Response
```json
{
    "short_url": "http://localhost:8080/abc123",
    "original_url": "https://www.example.com/very/long/url",
    "created_at": "2024-05-10T14:30:00Z"
}
```

#### Status Codes
- `201 Created`: URL successfully shortened
- `400 Bad Request`: Invalid request body
- `409 Conflict`: URL already exists
- `500 Internal Server Error`: Server error

### 2. Get URL Information
Retrieves information about a short URL.

```
GET /api/urls/{shortCode}
```

#### Parameters
- `shortCode`: The short code of the URL

#### Response
```json
{
    "short_url": "http://localhost:8080/abc123",
    "original_url": "https://www.example.com/very/long/url",
    "created_at": "2024-05-10T14:30:00Z"
}
```

#### Status Codes
- `200 OK`: URL found
- `400 Bad Request`: Invalid short code
- `404 Not Found`: URL not found
- `500 Internal Server Error`: Server error

### 3. Delete URL
Deletes a short URL.

```
DELETE /api/urls/{shortCode}
```

#### Parameters
- `shortCode`: The short code of the URL to delete

#### Response
- `204 No Content`: URL successfully deleted

#### Status Codes
- `204 No Content`: URL successfully deleted
- `400 Bad Request`: Invalid short code
- `404 Not Found`: URL not found
- `500 Internal Server Error`: Server error

### 4. Redirect to Original URL
Redirects to the original URL.

```
GET /{shortCode}
```

#### Parameters
- `shortCode`: The short code of the URL

#### Response
- Redirects to the original URL

#### Status Codes
- `301 Moved Permanently`: Redirect to original URL
- `400 Bad Request`: Invalid short code
- `404 Not Found`: URL not found
- `500 Internal Server Error`: Server error

## Error Responses
All error responses follow this format:
```json
{
    "error": "Error message description"
}
```

## Example Usage

### Using cURL

1. Create a short URL:
```bash
curl -X POST http://localhost:8080/api/urls \
  -H "Content-Type: application/json" \
  -d '{"original_url": "https://www.example.com/very/long/url"}'
```

2. Get URL information:
```bash
curl http://localhost:8080/api/urls/abc123
```

3. Delete a URL:
```bash
curl -X DELETE http://localhost:8080/api/urls/abc123
```

4. Visit short URL (in browser):
```
http://localhost:8080/abc123
```

### Using Postman

1. Create Short URL:
   - Method: POST
   - URL: http://localhost:8080/api/urls
   - Headers: Content-Type: application/json
   - Body:
   ```json
   {
       "original_url": "https://www.example.com/very/long/url"
   }
   ```

2. Get URL Info:
   - Method: GET
   - URL: http://localhost:8080/api/urls/{shortCode}

3. Delete URL:
   - Method: DELETE
   - URL: http://localhost:8080/api/urls/{shortCode}

4. Redirect:
   - Method: GET
   - URL: http://localhost:8080/{shortCode} 