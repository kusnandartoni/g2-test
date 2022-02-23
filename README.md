# Test
with this repository, there is a file `data.json` that contains data users and movie

# Task
1. make a program to read that file with input user id as parameter
2. print array of **top four** object that contains movie title and count of users friend that watch that movie
3. you can use any programing language to complete this task as long as it can run by terminal

# Example
## Movies Collection
```javascript
  {
    "title": "The Shawshank Redemption",
    "duration": "PT142M",
    "actors": [ "Tim Robbins", "Morgan Freeman", "Bob Gunton" ],
    "ratings": [],
    "favorites": [66380, 7001, 9250, 34139],
    "watchlist": [15291, 51417, 62289, 6146, 71389, 93707]
  },
  {
    "title": "The Godfather",
    "duration": "PT175M",
    "actors": [ "Marlon Brando", "Al Pacino", "James Caan" ],
    "ratings": [],
    "favorites": [15291, 51417, 7001, 9250, 71389, 93707],
    "watchlist": [62289, 66380, 34139, 6146]
  },
  {
    "title": "The Dark Knight",
    "duration": "PT152M",
    "actors": [ "Christian Bale", "Heath Ledger", "Aaron Eckhart" ],
    "ratings": [],
    "favorites": [15291, 7001, 9250, 34139, 93707],
    "watchlist": [51417, 62289, 6146, 71389]
  },
  {
    "title": "Pulp Fiction",
    "duration": "PT154M",
    "actors": [ "John Travolta", "Uma Thurman", "Samuel L. Jackson" ],
    "ratings": [],
    "favorites": [15291, 51417, 62289, 66380, 71389, 93707],
    "watchlist": [7001, 9250, 34139, 6146]
  },
  {
    "title": "Schindler's List",
    "duration": "PT195M",
    "actors": [
      "Liam Neeson",
      "Ralph Fiennes",
      "Ben Kingsley"
    ],
    "ratings": [{"userId": 62289, "rating": 8},{"userId": 66380, "rating": 5},{"userId": 6146, "rating": 6},{"userId": 71389, "rating": 7}],
    "favorites": [62289, 66380, 6146, 71389],
    "watchlist": [15291, 51417, 7001, 9250, 93707]
  },
```

### Users collection

```javascript
    {
        "userId": 15291,
        "email": "Constantin_Kuhlman15@yahoo.com",
        "friends": [7001, 51417, 62289]
    },
    {
        "userId": 7001,
        "email": "Keven6@gmail.com",
        "friends": [15291, 51417, 62289, 66380]
    },
    {
        "userId": 51417,
        "email": "Margaretta82@gmail.com",
        "friends": [15291, 7001, 9250]
    },
    {
        "userId": 62289,
        "email": "Marquise.Borer@hotmail.com",
        "friends": [15291, 7001]
    }
```

# Example execution 
 *this example is using golang, `go run main.go [userId]`

> go run main.go 15291
### Example output
```json
[
  { "title": "Schindler's List", "watchedBy": 2 },
  { "title": "The Dark Knight", "watchedBy": 2 },
  { "title": "The Lord of the Rings: The Return of the King", "watchedBy": 2},
  { "title": "The Shawshank Redemption", "watchedBy": 2 }
]
```

> go run main.go 51417
### Example output
```json
[
  { "title": "Schindler's List", "watchedBy": 3 },
  { "title": "The Lord of the Rings: The Return of the King", "watchedBy": 3 },
  { "title": "Pulp Fiction", "watchedBy": 2 },
  { "title": "The Shawshank Redemption", "watchedBy": 1 }
]
```