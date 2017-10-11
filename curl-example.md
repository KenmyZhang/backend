
  

    curl -X POST http://127.0.0.1:8099/v1/file/upload -i -F "uploadname=@/home/kenmy/Downloads/wx_sample.php"
    HTTP/1.1 200 OK
    Content-Length: 7
    Content-Type: text/plain; charset=utf-8
    Date: Fri, 06 Oct 2017 05:09:09 GMT
    Keep-Alive: timeout=38
    Server: beegoServer:1.9.0
    Set-Cookie: mini-chat=17bd7ca43486b1843348e96a1c0c656e; Path=/; Expires=Fri, 06 Oct 2017 06:09:09 GMT; Max-Age=3600; HttpOnly

    success    

    curl -X POST http://127.0.0.1:8099/v1/user/logout -i 
    HTTP/1.1 200 OK
    Content-Length: 7
    Content-Type: text/plain; charset=utf-8
    Date: Fri, 06 Oct 2017 05:11:07 GMT
    Keep-Alive: timeout=38
    Server: beegoServer:1.9.0
    Set-Cookie: mini-chat=98b71e82385ad391eafe60ee0cb82715; Path=/; Expires=Fri, 06 Oct 2017 06:11:07 GMT; Max-Age=3600; HttpOnly
    Set-Cookie: mini-chat=; Path=/; Expires=Fri, 06 Oct 2017 05:11:07 GMT; Max-Age=0; HttpOnly

    success    