package main

// go内置的net/http包十分优秀，提供了HTTP客户端和服务端的实现

// 1.net/http 介绍 提供了HTTP客户端和服务端的实现
// HTTP协议：超文本传输协议，所有的www文件都必须遵守这个标准。设计HTTP最初的目的是为了提供一种发布和接收HTML页面的方法

// 2.Http客户端 Get、Head、Post、PostForm函数发出HTTP/HTTPS请求
//resp, err := http.Get("http://example.com/")
//resp, err := http.Post("http://example.com/", "image/jpeg", &but)
//resp, err := http.PostForm("http://example.com/form", url.Values{"key": {"Value"}, "id": {"123"}})

// 程序在使用完response后必须关闭回复的主体
//resp, err := http.Get("http://example.com")
//if err != nil{
//	// handle error
//}
//
//defer resp.Body.Close()
//
//body, err := ioutil.ReadAll(resp.Body)

//func main() {
//	resp, err := http.Get("http://www.liwenzhou.com/")
//	if err != nil {
//		fmt.Printf("get failed, err:%v\n", err)
//		return
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil{
//		fmt.Printf("read from resp.Body failed, err:%v\n", err)
//		return
//	}
//	fmt.Print(string(body))
//}

// 2.1.带参数的Get请求：GET请求的参数需要使用GO语言内置的net/url这个标准库来处理
//func main() {
//	apiUrl := "http://127.0.0.1:9090/get"
//	data := url.Values{}
//	data.Set("name", "小明")
//	data.Set("age", "18")
//	u, err := url.ParseRequestURI(apiUrl)
//	if err != nil{
//		fmt.Printf("parse url requestURL failed, err:%v\n", err)
//	}
//	u.RawQuery = data.Encode() // URL encode
//	fmt.Println(u.String())
//	resp, err := http.Get(u.String())
//	if err != nil{
//		fmt.Printf("Post failed, err:%v\n", err)
//		return
//	}
//	defer resp.Body.Close()
//	b, err := ioutil.ReadAll(resp.Body)
//	if err != nil{
//		fmt.Printf("get resp failed, err:%v\n", err)
//		return
//	}
//	fmt.Println(string(b))
//}
//
//// 2.1.对应的Server端HandlerFunc如下：
//func getHandler(w http.ResponseWriter, r *http.Request){
//	defer r.Body.Close()
//	data := r.URL.Query()
//	fmt.Println(data.Get("name"))
//	fmt.Println(data.Get("age"))
//	answer := `{"status": "ok"}`
//	w.Write([]byte(answer))
//}

// 2.2.Post请求示例
// net/http post demo
//func main() {
//	url := "http://127.0.0.1:9090/post"
//	contentType := "application/json"
//	data := `{"name":"小明", "age": "18"}`
//	resp, err := http.Post(url, contentType, strings.NewReader(data))
//	if err != nil{
//		fmt.Printf("post failed, err:%v\n", err)
//		return
//	}
//	defer resp.Body.Close()
//	b, err := ioutil.ReadAll(resp.Body)
//	if err != nil{
//		fmt.Printf("get resp failed, err: %v\n", err)
//		return
//	}
//	fmt.Println(string(b))
//}
//
//// 2.2.server端HandlerFunc如下：
//func postHandler(w http.ResponseWriter, r *http.Request){
//	defer r.Body.Close()
//	// 2.2.1.请求类型是application/x-www-form-urlencoded时解析form数据
//	r.ParseForm()
//	fmt.Println(r.PostForm)	// 打印form数据
//	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
//	// 2.2.2.请求类型是application/json时从r.Body读取数据
//	b, err := ioutil.ReadAll(r.Body)
//	if err != nil{
//		fmt.Printf("read request.Body failed, err:%v\n", err)
//		return
//	}
//	fmt.Println(string(b))
//	answer := `{"status": "ok"}`
//	w.Write([]byte(answer))
//}

// 2.3.自定义Client 要管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client：
//client := &http.Client{
//	CheckRedirect: redirectPolicyFunc,
//}
//resp, err := client.Get("http://example.com")
//req, err := http.NewRequest("GET", "http://example.com", nil)
//req.Header.Add("If-None-Macth", `W/"wyzzy"`)
//resp. err := client.Do(req)
//
//// 2.3.自定义Transport
//// 要管理代理，TLS配置、keep-alive、压缩和其他设置，创建一个Transpost
//tr := &http.Transport{
//	TLSClientConfig: &tls.Config{RootCAs: pool},
//	DisableCompression: true
//}
//
//client := &http.Client{Transport: tr}
//resp, err := client.Get("https://example.com")
//// Client和Transpost类型都可以安全的被多个goroutine同时使用，出于效率考虑，应该一次建立、重复使用

// 3.服务端
// 默认的server
// ListenAndServer使用指定的监听地址和处理器启动一个HTTP服务端。处理器参数通常是nil，这表示采用包变量DefaultServerMux作为处理器
// Handle和HandleFunc函数可以向DefaultServereMux添加处理器
//http.Handle("/foo", fooHandler)
//http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request){
//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
//})
//log.Fatal(http.ListenAndServe(":8080", nil))

// 3.1.默认的server示例
//func SayHello(w http.ResponseWriter, r *http.Request){
//	fmt.Fprintln(w, "Hello 沙河！")
//}
//
//func main() {
//	http.HandleFunc("/", SayHello)
//	err := http.ListenAndServe(":9090", nil)
//	if err != nil{
//		fmt.Printf("http server failed, err:%v\n", err)
//		return
//	}
//}

// 3.2.自定义server
// 管理服务端的行为，可以创建一个自定义的Server：
//s := &http.Server{
//	Addr: 				":8080",
//	Handler:			myHandler,
//	ReadTimeOut:		10 * time.Second,
//	WriteTimeOut:		10 * time.Second,
//	MaxHeaderBytes: 	1 << 20,
//}
//log.Fatal(s.ListenAndServer())

