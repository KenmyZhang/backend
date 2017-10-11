# CreatePost

## Url
    /v1/post/create
    
### Method
    POST
    
### Request Payload
    {
        "channel_id": "jsm5vb35hicadac1c8cc3d",
        "message": "hello"
    }

### Response Body
    {
        "id": "abv81ewkibcbuvc1p3b1pw",
        "create_at": 1507691601762,
        "update_at": 1507691601762,
        "delete_at": 0,
        "user_id": "bb5ccuh5zje1tcvza5ww3b",
        "channel_id": "jsm5vb35hicadac1c8cc3d",
        "message": "hello",
        "type": "",
        "last_picture_update": 0
    }

### Response Code
    HTTP/1.1 200 OK
    
### Example    
    curl -X POST "http://127.0.0.1:8099/v1/post/create" -d '{"channel_id":"jsm5vb35hicadac1c8cc3d","message":"hello"}' -i -b ~/go/src/github.com/KenmyZhang/mini-chat/cookie.file
        HTTP/1.1 200 OK
        Server: beegoServer:1.9.0
        Date: Wed, 11 Oct 2017 03:13:21 GMT
        Content-Length: 222
        Content-Type: text/plain; charset=utf-8

    {
        "id": "abv81ewkibcbuvc1p3b1pw",
        "create_at": 1507691601762,
        "update_at": 1507691601762,
        "delete_at": 0,
        "user_id": "bb5ccuh5zje1tcvza5ww3b",
        "channel_id": "jsm5vb35hicadac1c8cc3d",
        "message": "hello",
        "type": "",
        "last_picture_update": 0
    }