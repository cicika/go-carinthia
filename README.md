# Go Carithia 

No fuss. Just ride.

URL: `http://gocarinthia.cicika.info`

`Authorization: GoCarinthia auth_token`

`PUT /user/register/`

```json
    Email: mail@address.com,
    Name: Full Name,
    Password: password,
    BeaconIdentifier: asdfasdf, //optional
```

Response:

`200 OK`

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

`POST /trip/start/` //depends if you want to start trip separately or on login, 2nd means one action less for the user

```json
    StartedAt: unix timestamp
```

Response: `200 OK`

`POST /trip/end/` //if not posted by device, backend will handle it

```json
    EndedAt: unix timestamp
```

Response: `200 OK`

`POST /trip/segment/`

```json
    TransportStation: (name or id from location database),
    TransportType: type from location database (can be empty),
    BeaconIdentifier: regiestered beacon,
    Latitude: 44.28,
    Longitude: 21.13, //don't care about decimals, send all
    CreatedAt: unix timestamp)
    CheckInType: 0 for start, 1 for stop, 2 for end
```

`GET /user/:user_id/:beacon_identifier`

Response:

```json
     LastCheckinAt: unix timestamp,
     BeaconIdentifier: string
```

`POST /user/payment/method/`

```json
    CardNumner: string
    CardHolder: string
    Expiration: string
    Cvv: integer
```