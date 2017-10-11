# CreateDirectChannel

## Url
    /v1/channel/create
    
### Method
    POST
    
### Request Payload
    {
        "userIds":["8bd4bw4dhbdi3ci8zddc34"]
    }

### Response Body
    {
        "Id": "jsm5vb35hicadac1c8cc3d",
        "CreateAt": 1507688491527,
        "UpdateAt": 1507688491527,
        "DeleteAt": 0,
        "Type": "D",
        "DisplayName": "",
        "Name": "8bd4bw4dhbdi3ci8zddc34__bb5ccuh5zje1tcvza5ww3b",
        "Header": "",
        "Purpose": "",
        "LastPostAt": 0,
        "TotalMsgCount": 0,
        "CreatorId": "bb5ccuh5zje1tcvza5ww3b"
    }

### Response Code
    HTTP/1.1 200 OK
    
### Example    
    curl -X POST "http://127.0.0.1:8099/v1/channel/create" -i -d '{"userIds":["8bd4bw4dhbdi3ci8zddc34"]}' -b ~/go/src/github.com/KenmyZhang/mini-chat/cookie.file
        HTTP/1.1 200 OK
        Server: beegoServer:1.9.0
        Date: Wed, 11 Oct 2017 02:21:31 GMT
        Content-Length: 273
        Content-Type: text/plain; charset=utf-8

    {
        "Id": "jsm5vb35hicadac1c8cc3d",
        "CreateAt": 1507688491527,
        "UpdateAt": 1507688491527,
        "DeleteAt": 0,
        "Type": "D",
        "DisplayName": "",
        "Name": "8bd4bw4dhbdi3ci8zddc34__bb5ccuh5zje1tcvza5ww3b",
        "Header": "",
        "Purpose": "",
        "LastPostAt": 0,
        "TotalMsgCount": 0,
        "CreatorId": "bb5ccuh5zje1tcvza5ww3b"
    }