# ZSSN (Zombie Survival Social Network)

API for the zombie apocalypse.

## Building

Build local binary by running:

```
go build ./main.go
```

Build the docker image by running:

```
docker build -t <image name> .
```

## Deployment

Deploy locally by running:

```
docker compose up
```

## Testing

Run unit tests with:

```
go test ./...
```

## Usage

After running locally, the service can be used by accessing it via [swagger](http://localhost:8080/swagger/index.html),
or hitting it directly.

### Registering a survivor

A user is registered with *name*, *age*, *gender*, *location*, and an array of *items*.

The location must contain a *lat* and a *lon*.

Items is an array of item identifiers from the following list:

| Item         | Identifier |
|--------------|------------|
| Ammunition   | 1          |
| Medication   | 2          |
| Food         | 3          |
| Water        | 4          |

You might notice that the item identifer coincides with the item value - 
the was a shortcut I used to save time - in a pre-apocalypse world where we
have more development time, we might have separated identifiers from costs.

### Example

To register Rick Grimes with 3 Ammunition, 2 Food and 1 Medication, run:

```
curl -X 'POST' \
  'http://localhost:8080/api/v1/user/register' \  
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "age": 32,    
  "gender": "male",
  "items": [     
    1, 1, 1, 3, 3, 2
  ],
  "location": {
    "lat": 3.334433,
    "lon": 4.34534543
  },
  "name": "Rick Grimes"
}
'
```

The returned value contains the user id, which must be used when accessing
endpoints for this user.

### Example: Updating location

To update location, run, e.g.:

```
curl -X 'PUT' \ 
  'http://localhost:8080/api/v1/user/1/location' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "lat": 45.4545454,
  "lon": 55.5555555
}'
```

### Example: Flag survivor as infected

To flag another survivor as infected, run, e.g.:

```
curl -X 'POST' \
  'http://localhost:8080/api/v1/user/2/flag' \    
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "sender_id": 1
}'
```

### Example: Trade items

To trade 2 Ammunition for 1 Medication, run e.g.:

```
curl -X 'POST' \
  'http://localhost:8080/api/v1/user/2/flag' \    
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "sender_id": 0,
  "offered_items": [
    {
      "item": 1
    },
    {
      "item": 1
    }
  ],
  "recepient_id": 2,
  "requested_items": [
    {
      "item": 2
    }
  ]
}'
```

### Report: Percentage of infected survivors

To get percentage of of infected survivors, run:

```
curl -X 'GET' \
  'http://localhost:8080/api/v1/report/percentage?infected=true' \
  -H 'accept: application/json'
```

To get percentage of of non-infected survivors, run:

```
curl -X 'GET' \
  'http://localhost:8080/api/v1/report/percentage?infected=false' \
  -H 'accept: application/json'
```

### Report: Average amount of each kind of resource

To get average amount of each kind of resource per survivor (e.g. 5 waters per survivor), run:

```
curl -X 'GET' \
  'http://localhost:8080/api/v1/report/averages' \
  -H 'accept: application/json'
```

Items will be the item identifiers as shown above.

### Report: Points lost because of infected survivors

To get points lost because of infected survivors, run:

```
curl -X 'GET' \
  'http://localhost:8080/api/v1/report/lost' \    
  -H 'accept: application/json'
```

## Notes

Things that are missing because I had to run away from zombies, are
(include, but not limited to):

 * Unit tests for the db package.
 * E2E tests for the endpoints.
 * Tests on push in the form of a github action, which could run a test postgres service.
 * Better handling of failed database connection instead of panicing.

Some endpoints (`/flag` and `/trade`) trust that the sender id is valid, because
there is no authentication, if there were fewer zombie hordes then this could be remedied.

