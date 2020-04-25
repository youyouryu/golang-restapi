# golang-restapi
This repositry contains REST API server using MySQL written in golang.

## Reauirements
- golang
- docker
- docker-compose

## Run
```
git clone https://github.com/youyouryu/golang-restapi.git
cd golang-restapi
docker-compose up
```

## Usage
- POST /signup
```
$ curl localhost/signup -X POST -d '{"user_id": "{user_id}", "password": "{password}"}' --dump-header -
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 25 Apr 2020 09:26:59 GMT
Content-Length: 96

{"message":"Account successfully created","user":{"user_id":"{user_id}","nickname":"{user_id}"}}
```

- GET /users/{user_id}
```
$ curl localhost/users/{user_id} -X GET -H 'Authorization: Basic {user_id:password encoded with base64}'   --dump-header -
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 25 Apr 2020 09:28:38 GMT
Content-Length: 123

{"message":"User details by user_id","user":{"user_id":"{user_id}","nickname":"{nickname}","comment":"{comment}"}}
```

- PATCH /users/{user_id}
```
$ curl localhost/users/yuyamada -X PATCH -H 'Authorization: Basic {user_id:password encoded with base64}' -d '{"nickname": "{nickname}", "comment": "{comment}"}' --dump-header -
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 25 Apr 2020 09:39:43 GMT
Content-Length: 75

{"message":"User successfully updated","recipe":[{"comment":"{comment}"}]}
```

- DELETE /close
```
$ curl localhost/close -X POST -H 'Authorization: Basic {user_id:password encoded with base64}'   --dump-header -
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 25 Apr 2020 09:34:52 GMT
Content-Length: 51

{"message":"Account and user successfully removed"}
```
