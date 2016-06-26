# Go Carithia 

## No fuss! Just ride :)

## API Documentation

URL: `http://gocarinthia.cicika.info`

`Authorization: GoCarinthia <auth_token>`

### User registration

`PUT /user/`

Request:
```json
    Email: mail@address.com,
    Name: Full Name,
    Password: string,
    BeaconIdentifier: string, //optional
```

Response: 
`200 OK`
```json
    AuthToken: <token>,
    UserId: integer
```

## User login

Request:

`POST /user/login/`

```json
    Email: mail@addres.com,
    Password: password
```

Response:

```json
    Id: 123,
    AuthToken: token,
    Type: passenger
```

### Start new trip

Request:

`POST /trip/start/` 

```json
    StartedAt: Unix timestamp
```

Response: 

`200 OK`

### End trip

Request:

`POST /trip/end/` 

If not posted by device, backend will handle it.

```json
    EndedAt: unix timestamp
```

Response: `200 OK`

### Record check in point (trip segment)

Request:

`POST /trip/segment/`

```json
    TransportStation: (name or id from location database),
    TransportType: type from location database (can be empty),
    BeaconIdentifier: registered beacon,
    Latitude: 44.28,
    Longitude: 21.13, //don't care about decimals, send all
    CreatedAt: unix timestamp)
    CheckInType: 0 for start, 1 for stop, 2 for end
```
### Passenger check (for controller's device)

`GET /user/:user_id/:beacon_identifier`

Response:

```json
    LastCheckinAt: unix timestamp,
    BeaconIdentifier: string
```

### CC details

Request:

`POST /user/payment/method/`

```json
    CardNumner: string
    CardHolder: string
    Expiration: string
    Cvv: integer
```

Response:
`200 Ok`