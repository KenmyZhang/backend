
# GetPosts

## Url
    /v1/post/channel/:channel_id([A-Za-z0-9]+)/posts?per_page=2&page=0
    
### Method
    GET
    
### Request Payload

### Response Body
    [
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
        },
        {
            "id": "kc2ecs1p8jdcubw1bmwhhs",
            "create_at": 1507691546422,
            "update_at": 1507691546422,
            "delete_at": 0,
            "user_id": "bb5ccuh5zje1tcvza5ww3b",
            "channel_id": "jsm5vb35hicadac1c8cc3d",
            "message": "hello",
            "type": "",
            "last_picture_update": 0
        }
    ]

### Response Code
    HTTP/1.1 200 OK
    
### Example    
    curl -X GET "http://127.0.0.1:8099/v1/post/channel/jsm5vb35hicadac1c8cc3d/posts?per_page=2&page=0"  -i -b ~/go/src/github.com/KenmyZhang/mini-chat/cookie.file -i
        HTTP/1.1 200 OK
        Server: beegoServer:1.9.0
        Set-Cookie: mini-chat=f44a4658732b27d72fbd2ae9d03054e5; Path=/; Expires=Wed, 11 Oct 2017 07:09:02 GMT; Max-Age=3600; HttpOnly
        Date: Wed, 11 Oct 2017 06:09:02 GMT
        Content-Length: 447
        Content-Type: text/plain; charset=utf-8

    [
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
        },
        {
            "id": "kc2ecs1p8jdcubw1bmwhhs",
            "create_at": 1507691546422,
            "update_at": 1507691546422,
            "delete_at": 0,
            "user_id": "bb5ccuh5zje1tcvza5ww3b",
            "channel_id": "jsm5vb35hicadac1c8cc3d",
            "message": "hello",
            "type": "",
            "last_picture_update": 0
        }
    ]