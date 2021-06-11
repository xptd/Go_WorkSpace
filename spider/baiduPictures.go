package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

const targetUrl string = `https://image.baidu.com/search/albumsdetail?tn=albumsdetail&word=%E8%AE%BE%E8%AE%A1%E7%B4%A0%E6%9D%90&fr=albumslist
&album_tab=%E5%BB%BA%E7%AD%91&album_id=1&rn=30`

const cookie string = `BIDUPSID=8357087683D6F2C9572D909248A1F31A; PSTM=1602318818; BAIDUID=7CD2A6F101C6A94E8690C931141B4D10:FG=1; BDUSS=BITElzdWlWYnJkbFFYZHRYRkg5TkVCeWdOdWdWUTV6NEZ6MDZWMThVSG9rN0JmSVFBQUFBJCQAAAAAAAAAAAEAAAC4TKJ-8J-RvfCfkb3nmoTwn5EAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAOgGiV~oBolfbl; BDUSS_BFESS=BITElzdWlWYnJkbFFYZHRYRkg5TkVCeWdOdWdWUTV6NEZ6MDZWMThVSG9rN0JmSVFBQUFBJCQAAAAAAAAAAAEAAAC4TKJ-8J-RvfCfkb3nmoTwn5EAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAOgGiV~oBolfbl; BCLID_BFESS=11251630079535006508; BDSFRCVID_BFESS=fMDOJeC62x57iiReHKOabQmZh2ddxfRTH6aoM8jxVCLdgV_NGdc7EG0PVf8g0Kub1DMtogKKL2OTHm_F_2uxOjjg8UtVJeC6EG0Ptf8g0f5; H_BDCLCKID_SF_BFESS=tJIJ_ID2JCD3enTmbt__-P4DeUju0URZ5mAqotPaJbTEVR5SyUnH-l_DWxAHKP3O5HrnaIQqahCWj4nKLUvfbpKmX4Jb55b43bRT-JLy5KJvfJ_lQMcMhP-UyPkHWh37aGOlMKoaMp78jR093JO4y4Ldj4oxJpOJ5JbMonLafJOKHICwej-53e; __yjs_duid=1_07c5a064c5f1d0ac4377777c4c422e0c1617774602509; firstShowTip=1; ab_sr=1.0.0_NjdkY2Y3MjY0MzM2OTU2MDQ0MDYxZjEzZWQxNjJlNzM0OGM4YTY3NTgwZGFiNDBiYWNjNGNkZjM0Y2RkMzc2N2I1NmU4NTFkZDk2Mjk3ZGEzMGY4YjljYjFhYWYzMGVhNGMwYzA4MTJkYThiZWU0ZDdhNTVlYWI2MGU2MzIzZDE=; H_PS_PSSID=33798_33820_33752_33272_33691_33855_33713_26350; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; BA_HECTOR=ag8ga52ka0a12h2g6k1g6qs250q; BDRCVFR[feWj1Vr5u3D]=I67x6TjHwwYf0; delPer=0; PSINO=2; BAIDUID_BFESS=7CD2A6F101C6A94E8690C931141B4D10:FG=1; BDRCVFR[Q5XHKaSBNfR]=mk3SLVN4HKm; userFrom=null`
const bufLen int = 512

func httpGet(url string) (ret string, rerr error) {
	res, err := http.Get(url)
	if err != nil {
		rerr = err
		return
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			rerr = err
		}
	}()
	buf := make([]byte, bufLen)
	for {
		c, err := res.Body.Read(buf)
		if c == 0 {
			break
		}
		if err != nil && err != io.EOF {
			rerr = err
			return
		}
		ret += string(buf[:c])
	}
	return ret, err
}

func getResponse(url string) (ret string, rerr error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Host", "image.baidu.com")
	req.Header.Set("Accept-Encoding", "zip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", cookie)
	res, err := client.Do(req)
	if err != nil {
		rerr = err
		return
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			rerr = err
		}
	}()
	fmt.Println(res.Status)
	buf := make([]byte, bufLen)
	for {
		c, err := res.Body.Read(buf)
		if c == 0 {
			break
		}
		if err != nil && err != io.EOF {
			rerr = err
			return
		}
		ret += string(buf[:c])
	}
	return ret, err
}
func main() {
	decodeUrl, _ := url.QueryUnescape(strings.Replace(targetUrl, "\n", "", 1))
	//fmt.Println(decodeUrl)
	//res, err := httpGet(decodeUrl)
	res, err := getResponse(decodeUrl)
	if err != nil {
		panic(err)
	}

	file, ferr := os.Create("res.txt")
	if ferr != nil {
		panic(err)
	}
	defer func() {
		fmt.Println("close file")
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	//	fmt.Println(res)
	n, _ := file.WriteString(res)
	fmt.Println(n)
	reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	tmp := reg.FindAllStringSubmatch(res, -1)
	fmt.Println(tmp)

}
