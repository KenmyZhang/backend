# UploadFile

## Url
    /v1/file/upload
    
### Method
    POST
    
### Request Payload
    -F "uploadname=@/home/accurme/下载/vue.js" -i 

### Response Body
    success

### Response Code
    HTTP/1.1 200 OK
    
### Example    
    curl -X POST "http://127.0.0.1:8099/v1/file/upload" -F "uploadname=@/home/accurme/下载/vue.js" -i 
HTTP/1.1 100 Continue

    HTTP/1.1 200 OK
    Server: beegoServer:1.9.0
    Set-Cookie: mini-chat=c796bcf39605d739a23d7ffafecb94f3; Path=/; Expires=Wed, 11 Oct 2017 03:35:34 GMT; Max-Age=3600; HttpOnly
    Date: Wed, 11 Oct 2017 02:35:34 GMT
    Content-Length: 7
    Content-Type: text/plain; charset=utf-8

    success