# Go Carithia 

No fuss. Just ride.

URL: `http://gocarinthia.cicika.info`

`Authorization: GoCarinthia auth_token`

`POST /user/register`

```json
    email: mail@address.com,
    name: Full Name,
    password: password,
    beacon_identifier: asdfasdf, //optional
```

Response:

`200 OK`

`POST /user/login`

```json
    email: mail@addres.com,
    password: password
```

Response:

```json
    id: 123,
    auth_token: token,
    type: passenger
```

`POST /trip/start` //depends if you want to start trip separately or on login, 2nd means one action less for the user

```json
    started_at: unix timestamp
```

Response: `200 OK`

`POST /trip/end` //if not posted by device, backend will handle it

```json
    ended_at: unix timestamp
```

Response: `200 OK`

`POST /trip/segment`

```json
    transport_station: (name or id from location database),
    transport_type: type from location database (can be empty),
    beacon_identifier: regiestered beacon,
    latitude: 44.28,
    longitude: 21.13, //don't care about decimals, send all
    created_at: unix timestamp)
    check_in_type: 0 for start, 1 for stop, 2 for end
```

`GET /user/:user_id/:beacon_identifier`

Response:

```json
     last_checkin_at: unix timestamp,
     beacon_identifier: string
```