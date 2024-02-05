# wordle-go

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)

wordle-go is a simple and engaging implementation of the popular word-guessing game, Wordle, written in Go. Challenge yourself to guess the hidden word within a limited number of attempts and improve your word-guessing skills!

## API Routes

### Check

- **Method:** `GET`
- **Endpoint:** `/api/check`
- **Description:** Check if a word is correct.
- **Query Parameters:**
  - `word` (string): The word to be checked.
- **Responses:**
  - Code 200: Success. Returns the result of the check.
  - Code 400: Bad request. Returns details of the error.
- **example:** `/api/check?word=hello`

  ```json
  {
    "check": [
      { "letter": "h", "check": "correct" },
      { "letter": "e", "check": "correct" },
      { "letter": "l", "check": "wrong" },
      { "letter": "l", "check": "wrong" },
      { "letter": "o", "check": "elsewhere" }
    ],
    "success": false
  }
  ```

### Attempts

- **Method:** `GET`
- **Endpoint:** `/api/attempts/:date`
- **Description:** Get user attempts on a specific date.
- **Path Parameters:**
  - `date` (string): Date in the format "YYYY-MM-DD".
- **Responses:**
  - Code 200: Success. Returns user attempts on the specified date.
  - Code 400: Bad request. Returns details of the error.
- **example:** `/api/attempts/2024-01-22`
  ```json
  {
    "attempts": [
      {
        "IP": "::1",
        "Attempt": [
          { "letter": "s", "check": "wrong" },
          { "letter": "h", "check": "wrong" },
          { "letter": "i", "check": "wrong" },
          { "letter": "f", "check": "wrong" },
          { "letter": "t", "check": "elsewhere" }
        ],
        "Success": false,
        "Date": "2024-01-22T13:24:02.11768545-03:00"
      },
      {
        "IP": "::1",
        "Attempt": [
          { "letter": "t", "check": "correct" },
          { "letter": "w", "check": "correct" },
          { "letter": "i", "check": "wrong" },
          { "letter": "s", "check": "wrong" },
          { "letter": "t", "check": "wrong" }
        ],
        "Success": false,
        "Date": "2024-01-22T13:24:14.035498653-03:00"
      },
      {
        "IP": "::1",
        "Attempt": [
          { "letter": "t", "check": "correct" },
          { "letter": "w", "check": "correct" },
          { "letter": "e", "check": "correct" },
          { "letter": "a", "check": "correct" },
          { "letter": "k", "check": "correct" }
        ],
        "Success": true,
        "Date": "2024-01-22T13:24:22.785897983-03:00"
      }
    ]
  }
  ```
