package main

// 手动引入microErrors，防止重名
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"oldboyedu-go/my-micro/demo/src/share/config"
	"oldboyedu-go/my-micro/demo/src/share/utils/path"
	"strconv"

	"github.com/micro/go-micro/cmd"
)

func main() {
	// 创建handler
	mux := http.NewServeMux()
	// 创建 handleRPC() 方法，作为处理器
	// 自定义的serverMux对象传到server对象中
	// 第一个参数是请求路径，第二个参数是函数类型
	mux.HandleFunc("/", handleRPC)
	log.Println("Listen on :8888")
	// 监听8888端口，绑定handler处理http请求
	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		fmt.Println(err)
	}
}

// 实现handler，并实现跨域处理
// 参数根据HandleFunc中要求的写
func handleRPC(w http.ResponseWriter, r *http.Request) {
	log.Println("handleRPC ....")
	// 1.处理正常请求
	if r.URL.Path == "/" {
		_, err := w.Write([]byte("server ..."))
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	// 2.处理跨域请求
	// Get("Origin")处理跨域问题
	// Get("Origin")会得到具体路径，例如http://localhost:4200
	if origin := r.Header.Get("Origin"); true {
		// 设置允许的跨域访问路径为取到的路径
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	// 设置跨越请求允许的请求方式
	//w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	// 设置跨越请求允许的数据格式
	w.Header().Set("Access-Control-Allow-Headers",
		"Content-Type, Content-Length, Accept-Encoding, X-Token, X-Client")
	// 设置跨越请求是否可携带证书
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	// 创建 handleJSONRPC() 方法
	// 具体的rpc请求处理
	handleJSONRPC(w, r)
	return
}

// 具体的rpc请求处理
func handleJSONRPC(w http.ResponseWriter, r *http.Request) {
	// 处理请求路径，得到具体服务和方法
	// 将url转换为service和method
	// 这里使用了工具类中的 pathToReceiver()
	service, method := path.PathToReceiver(config.Namespace, r.URL.Path)
	// service最终为 service:com.class.cinema.user
	// method最终为 method:UserService.SelectUser
	log.Println("service:" + service)
	log.Println("method:" + method)
	// 读取请求体
	br, _ := ioutil.ReadAll(r.Body)
	// 封装request
	request := json.RawMessage(br)
	// 调用服务
	var response json.RawMessage
	req := (*cmd.DefaultOptions().Client).NewJsonRequest(service, method, &request)
	ctx := path.RequestToContext(r)
	err := (*cmd.DefaultOptions().Client).Call(ctx, req, &response)
	// make the call
	if err != nil {
		fmt.Println(err)
		return
	}
	// 编码json
	b, _ := response.MarshalJSON()
	// 设置响应头
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	// 写出
	_, err = w.Write(b)
	if err != nil {
		fmt.Println(err)
	}
}
