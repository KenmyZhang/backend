# CreateUser

## Url
    /v1/user/create
    
### Method
    POST
    
### Request Payload
    {
        "username": "Beegzhang",
        "password": "123456",
        "email": "1023@qq.com",
        "age": 21,
        "phoneNum": "13544285662"
    }

### Response Body
    {
        "Id": "58w1b1sjkidsdmdphhcchs",
        "CreateAt": 1507263735423,
        "UpdateAt": 1507263735423,
        "DeleteAt": 0,
        "Username": "beegzhang",
        "Password": "$2a$10$0VssMCNyd8ZpdJd1DtddqeyKEhte7rHGYrpWHiRoNhlX18xih.Vp2",
        "AuthData": null,
        "AuthService": "",
        "Email": "1023@qq.com",
        "EmailVerified": false,
        "Age": 21,
        "Nickname": "",
        "Position": "",
        "Roles": "",
        "PhoneNum": "13544285662",
        "LastPasswordUpdate": 1507263735423,
        "LastPictureUpdate": 0,
        "FailedAttempts": 0
    }

### Response Code
    HTTP/1.1 200 OK
    
### Example    
     curl -X POST "http://127.0.0.1:8099/v1/user/create" -i -d '{"username": "Beegzhang","password":"123456","email":"1023@qq.com","age":21,"phoneNum":"13544285662"}'
        HTTP/1.1 200 OK
        Content-Length: 417
        Content-Type: text/plain; charset=utf-8
        Date: Fri, 06 Oct 2017 04:22:15 GMT
        Keep-Alive: timeout=38
        Server: beegoServer:1.9.0
        Set-Cookie: mini-chat=7a849d4441a7c16112220d72da834c1c; Path=/; Expires=Fri, 06 Oct 2017 05:22:15 GMT; Max-Age=3600; HttpOnly

    {
        "Id": "58w1b1sjkidsdmdphhcchs",
        "CreateAt": 1507263735423,
        "UpdateAt": 1507263735423,
        "DeleteAt": 0,
        "Username": "beegzhang",
        "Password": "$2a$10$0VssMCNyd8ZpdJd1DtddqeyKEhte7rHGYrpWHiRoNhlX18xih.Vp2",
        "AuthData": null,
        "AuthService": "",
        "Email": "1023@qq.com",
        "EmailVerified": false,
        "Age": 21,
        "Nickname": "",
        "Position": "",
        "Roles": "",
        "PhoneNum": "13544285662",
        "LastPasswordUpdate": 1507263735423,
        "LastPictureUpdate": 0,
        "FailedAttempts": 0
    }