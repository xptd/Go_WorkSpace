package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

/*
1. get https://www.csdn.net/` respose context
2.find tool-bar script
3.find toolbar url and param
4.get tool-bar response context
5.

*/
const buflen int = 512
const urlCsdn string = `https://www.csdn.net/`
const cookie string = `uuid_tt_dd=10_36636175020-1602319481864-892790; UN=weixin_42158070; Hm_ct_6bcd52f51e9b3dce32bec4a3997715ac=6525*1*10_36636175020-1602319481864-892790!5744*1*weixin_42158070; Hm_lvt_e5ef47b9f471504959267fd614d579cd=1603867504,1603867793; UserName=weixin_42158070; UserInfo=ff13651810db4458ab491ef5944af8aa; UserToken=ff13651810db4458ab491ef5944af8aa; UserNick=%40%E5%B0%8F%E8%B7%91%E5%A0%82%E7%9A%84; AU=FEE; BT=1618366949729; p_uid=U010000; ssxmod_itna=YqIxyDnDgADQeD5iQDXiao0Kommu3qKq==Wb+vDloxWqPGzDAxn40iDtrrwDObtAaqYKWom0Ei=7=CgPv+W6nWPDHxY=DUr7GP3D445GwD0eG+DD4DW0x03DoxGAysrqY3DcxYQDmxitDzq=DG3wbidULlqGRD7QDAwGBmi4nD0tDIqGX8ruKiqDB9juKDYPaeqiDxBQD7krnzdLKx0P3KDxa0p+sGG44na4q87DPG05su0GQ8x4pHG+QiAPPiAxWA+yTB1ioU/9TeD=; ssxmod_itna2=YqIxyDnDgADQeD5iQDXiao0Kommu3qKq==Wb+D6pRhD0vmrq03zOubegUD6YohqYvzDfOKeCOSrB7RzqfhuGd84aSkK=KZDbwAmpOvZpMQzicWIv2hn03hBTujvSKhW7IX5uprz3C4EV0MKzQ0E4r+Xq=eNVUhNIBONpUk=vGSQlDS6ERe6SQqYeGYUUIOGHCufl0S+lf83dP3WjK7GFP4QAAheHIP54Crb859QnWSW/WzD0WF7nC+E+h88pG4bYRd5CnFdCzORGMU5YrUvqxmXKhCfKEWU+PXfLUdFgUvvkD65mRXgAmKsndvnAVoIcrj4REZBUQr5gYU9Y3Ce5j0wEgttaja6i3UttO=fA5R1rYzrGBdjiq4e40jROl37zLRiNCxTonD+DhktDV7e4nDUYowiT74hQYWCzpTE6pKKSxaNFPXrrqGWpBK/c2p1QQgchj23S+QaQVYbz6o2WGr/Kh3oip4DQFe43Qf43fxrtW1PSN4bc9MQWDDFqD+cDxD==; Hm_up_6bcd52f51e9b3dce32bec4a3997715ac=%7B%22islogin%22%3A%7B%22value%22%3A%221%22%2C%22scope%22%3A1%7D%2C%22isonline%22%3A%7B%22value%22%3A%221%22%2C%22scope%22%3A1%7D%2C%22isvip%22%3A%7B%22value%22%3A%220%22%2C%22scope%22%3A1%7D%2C%22uid_%22%3A%7B%22value%22%3A%22weixin_42158070%22%2C%22scope%22%3A1%7D%7D; __gads=ID=24bdad66bc0a66e6-22234eb368c70034:T=1618540502:RT=1618540502:S=ALNI_MaUXhMg_4P0AMGsC4Sf3v0__ukifw; log_Id_click=775; csrfToken=IehVc7xUA4iemp7VlMUPSNy7; c_pref=; c_ref=https%3A//www.baidu.com/link; c_first_page=https%3A//www.csdn.net/; c_first_ref=www.baidu.com; c_segment=6; dc_sid=a36df5f9edd58aac7f21b153052715f2; Hm_lvt_6bcd52f51e9b3dce32bec4a3997715ac=1618536238,1618537687,1618540503,1618801746; dc_session_id=10_1618808803147.671678; c_page_id=default; dc_tos=qrsowj; log_Id_pv=1811; Hm_lpvt_6bcd52f51e9b3dce32bec4a3997715ac=1618808852; announcement-new=%7B%22isLogin%22%3Atrue%2C%22announcementUrl%22%3A%22https%3A%2F%2Fblog.csdn.net%2Fblogdevteam%2Farticle%2Fdetails%2F112280974%3Futm_source%3Dgonggao_0107%22%2C%22announcementCount%22%3A0%2C%22announcementExpire%22%3A3600000%7D; log_Id_view=5753`

func getRes(target string) (ret string, rerr error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", target, nil)
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	// req.Header.Set("accept-encoding", "gzip, deflate, br")
	// req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	// req.Header.Set("cookie", cookie)
	// req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
	// req.Header.Set("authority", "www.csdn.net")
	// req.Header.Set("path", "/")
	// req.Header.Set("cache-control", "max-age=0")
	// req.Header.Set("scheme", "https")

	if err != nil {
		rerr = err
		return
	}
	res, err := client.Do(req)
	if err != nil {
		rerr = err
		return
	}
	defer func() {
		if err = res.Body.Close(); err != nil {
			rerr = err
			return
		}
	}()
	body := res.Body
	//fmt.Println(res.Header.Get("Content-encoding"))
	//fmt.Println(res.Header.Get("Content-type"))
	// if res.Header.Get("Content-Encoding") == "gzip" {
	// 	body, err = gzip.NewReader(body)
	// 	if err != nil {
	// 		return "", err
	// 	}

	// }
	reader := bufio.NewReader(body)
	var buf []byte = make([]byte, buflen)

	for {
		cnt, err := reader.Read(buf)
		if cnt == 0 {
			break
		}
		if err != nil && err != io.EOF {
			rerr = err
			return
		}

		ret += string(buf[:cnt])
	}
	return ret, rerr
}
func findToolbarJs(context string) string {
	//(?s:(.*?))
	//<script src="https://g.csdnimg.cn/common/csdn-toolbar/csdn-toolbar.js"></script>
	/*
		1.<div>(?s:(.*?))</div>
			1.1:(.*?):(pattern):匹配pattern并获取这一匹配
			1.2:(?s)即Singleline(单行模式)
		2.`<script src="[a-zA-z]+://[^\s]*csdn-toolbar.js"></script>`
	*/
	reg := regexp.MustCompile(`<script src="[a-zA-z]+://[^\s]*csdn-toolbar.js"></script>`)
	tmp := reg.FindString(context)
	arr := strings.Split(tmp, "src=")
	tmp_1 := strings.Split(arr[1], ">")
	return strings.ReplaceAll(tmp_1[0], "\"", "")
}

func saveFile(fpath string, context string) (cnt int, serr error) {
	if fpath == "" {
		serr = fmt.Errorf("fpath err:%s", fpath)
		return
	}
	file, err := os.Create(fpath)
	if err != nil {
		serr = err
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			serr = err
			return
		}
	}()
	cnt, serr = file.WriteString(context)
	return cnt, serr
}
func getToolBarRes(jsPath string) (ret string, rerr error) {

	clinet := &http.Client{}
	fmt.Println(jsPath)
	req, err := http.NewRequest("GET", jsPath, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
	req.Header.Set("accept-encoding", "gzip, deflate, br")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("authority", "g.csdnimg.cn")

	res, err := clinet.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		if err = res.Body.Close(); err != nil {
			rerr = err
		}
	}()
	fmt.Println(res.StatusCode)
	//must gzip
	//fmt.Println(res.Header.Get("Content-encoding"))
	//fmt.Println(res.Header.Get("Content-type"))
	body := res.Body
	buf := make([]byte, buflen)
	if res.Header.Get("Content-Encoding") == "gzip" {
		body, err = gzip.NewReader(body)
		if err != nil {
			return "", err
		}

	}
	reader := bufio.NewReader(body)
	for {
		n, err := reader.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			break
		}
		ret += string(buf[:n])
	}
	return ret, rerr
}

func getToolBarUrl(date string) string {
	reg := regexp.MustCompile(`getToolbarData:.+success`)
	ret := reg.FindString(date)
	tmp := strings.Split(ret, ",")
	// for _, v := range tmp {
	// 	fmt.Println(v)
	// }

	reg = regexp.MustCompile(`url:"[a-zA-Z]+://[^\s]*"`)
	ret = reg.FindString(tmp[0])
	//url:"https://img-home.csdnimg.cn/data_json/toolbar/toolbar1217.json"
	reg = regexp.MustCompile(`(?s:"(.+?)")`)
	ret_1 := reg.FindAllStringSubmatch(ret, -1)
	return ret_1[0][1]
}
func main() {
	fmt.Println("hello csdn spider!")
	ret, err := getRes(urlCsdn)

	if err != nil {
		panic(err)
	}
	ret = findToolbarJs(ret)
	if ret != "" {
		ret, _ = getToolBarRes(ret)
		toolBarUrl := getToolBarUrl(ret)
		fmt.Println(toolBarUrl)
	}

	//fmt.Println(gbk2utf.GetStrCodeType([]byte(ret)))
	//fmt.Println(ret)
	// var fpath string = "csdnRes.txt"
	// cnt, err := saveFile(fpath, ret)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%d saveded to file:%s\r\n", cnt, fpath)
	// gbk2utf.Test()
}
