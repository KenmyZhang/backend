# quick start
##
    CREATE DATABASE test;

    CREATE TABLE `session` (
        `session_key` char(64) NOT NULL,
        `session_data` blob,
        `session_expiry` int(11) unsigned NOT NULL,
        PRIMARY KEY (`session_key`)
    ) ENGINE=MyISAM DEFAULT CHARSET=utf8;

# 当enablexsrf = true时,其中X-Xsrftoken头域的值是cookie第一个字段解码后的值
    // GetSecureCookie Get secure cookie from request by a given key.
    func (ctx *Context) GetSecureCookie(Secret, key string) (string, bool) {
        val := ctx.Input.Cookie(key)
        if val == "" {
            return "", false
        }

        parts := strings.SplitN(val, "|", 3)

        if len(parts) != 3 {
            return "", false
        }

        vs := parts[0]
        timestamp := parts[1]
        sig := parts[2]

        h := hmac.New(sha1.New, []byte(Secret))
        fmt.Fprintf(h, "%s%s", vs, timestamp)

        if fmt.Sprintf("%02x", h.Sum(nil)) != sig {
            return "", false
        }
        res, _ := base64.URLEncoding.DecodeString(vs)
        return string(res), true
    }

    var xsrf, xsrflist;
    xsrf = $.cookie("_xsrf");
    xsrflist = xsrf.split("|");
    args._xsrf = base64_decode(xsrflist[0]);

# parameters
### Get URL parameters
    GetString(key string) string
    GetStrings(key string) []string
    GetInt(key string) (int64, error)
    GetBool(key string) (bool, error)
    GetFloat(key string) (float64, error)
### Get Form parameters    
    ParseForm(obj interface{})
### Get request body parameters
     var ob models.Object
     json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
### usage
####Get URL param    
    func (this *MainController) Post() {
        jsoninfo := this.GetString("jsoninfo")
        ......
    }
####Get Form param 
    type user struct {
        Id    int         `form:"-"`
        Name  interface{} `form:"username"`
        Age   int         `form:"age"`
        Email string
    }
    func (this *MainController) Post() {
        u := user{}
        if err := this.ParseForm(&u); err != nil {
            //handle error
        }
    }    



# session
### method
    SetSession(name string, value interface{})
    GetSession(name string) interface{}
    DelSession(name string)
    SessionRegenerateID()
    DestroySession()
### attentions
    在使用 mysql 存储 session 信息的时候，需要事先在 mysql 创建表，建表语句如下
    CREATE TABLE `session` (
        `session_key` char(64) NOT NULL,
        `session_data` blob,
        `session_expiry` int(11) unsigned NOT NULL,
        PRIMARY KEY (`session_key`)
    ) ENGINE=MyISAM DEFAULT CHARSET=utf8;    
### usage 
    func (this *MainController) Get() {
        v := this.GetSession("asta")
        if v == nil {
            this.SetSession("asta", int(1))
        } else {
            this.SetSession("asta", v.(int)+1)
        }
        ......
    }


#  log
### set output
    beego.SetLogger("file", `{"filename":"logs/test.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
    beego.BeeLogger.DelLogger("console")
### level
    LevelEmergency
    LevelAlert
    LevelCritical
    LevelError
    LevelWarning
    LevelNotice
    LevelInformational
    LevelDebug
### set level
    beego.SetLevel(beego.LevelInformational)
### set file name and line number output
    beego.SetLogFuncCall(true)
### usage
    beego.Emergency("this is emergency")
    beego.Alert("this is alert")
    beego.Critical("this is critical")
    beego.Error("this is error")
    beego.Warning("this is warning")
    beego.Notice("this is notice")
    beego.Informational("this is informational")
    beego.Debug("this is debug")

# file upload
### method
    GetFile(key string) (multipart.File, *multipart.FileHeader, error)
    SaveToFile(fromfile, tofile string) error
### usage
    func (c *FormController) Post() {
        f, h, err := c.GetFile("uploadname")
        if err != nil {
            log.Fatal("getfile err ", err)
        }
        defer f.Close()
        c.SaveToFile("uploadname", "static/upload/" + h.Filename) 
    }

# API document auto created
###
    必须设置在 routers/router.go 中，文件的注释，最顶部：
    // @APIVersion 1.0.0
    // @Title mini-chat API
    // @Description easy to communicate with each other among team
    // @Contact 1027837952@qq.com.com
    //@TermsOfServiceUrl
    //@License
    //@LicenseUrl

    package routers

###
    // @Title Get Product list
    // @Description Get Product list by some info
    // @Success 200 {object} models.ZDTProduct.ProductList
    // @Param   category_id     query   int false       "category id"
    // @Failure 400 no enough input
    // @router /products [get]

###
    bee run -gendoc=true -downdoc=true

# Valid Form Parameters
###
    StructTag 可用的验证函数：

    Required 不为空，即各个类型要求不为其零值
    Min(min int) 最小值，有效类型：int，其他类型都将不能通过验证)
    Max(max int) 最大值，有效类型：int，其他类型都将不能通过验证
    Range(min, max int) 数值的范围，有效类型：int，他类型都将不能通过验证
    MinSize(min int) 最小长度，有效类型：string slice，其他类型都将不能通过验证
    MaxSize(max int) 最大长度，有效类型：string slice，其他类型都将不能通过验证
    Length(length int) 指定长度，有效类型：string slice，其他类型都将不能通过验证
    Alpha alpha字符，有效类型：string，其他类型都将不能通过验证
    Numeric 数字，有效类型：string，其他类型都将不能通过验证
    AlphaNumeric alpha 字符或数字，有效类型：string，其他类型都将不能通过验证
    Match(pattern string) 正则匹配，有效类型：string，其他类型都将被转成字符串再匹配(fmt.Sprintf(“%v”, obj).Match)
    AlphaDash alpha 字符或数字或横杠 -_，有效类型：string，其他类型都将不能通过验证
    Email 邮箱格式，有效类型：string，其他类型都将不能通过验证
    IP IP 格式，目前只支持 IPv4 格式验证，有效类型：string，其他类型都将不能通过验证
    Base64 base64 编码，有效类型：string，其他类型都将不能通过验证
    Mobile 手机号，有效类型：string，其他类型都将不能通过验证
    Tel 固定电话号，有效类型：string，其他类型都将不能通过验证
    Phone 手机号或固定电话号，有效类型：string，其他类型都将不能通过验证
    ZipCode 邮政编码，有效类型：string，其他类型都将不能通过验证





    