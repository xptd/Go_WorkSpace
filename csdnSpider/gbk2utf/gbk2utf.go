package gbk2utf

import (
	"fmt"

	"golang.org/x/text/encoding/simplifiedchinese"
)

const (
	GBK     string = "GBK"
	UTF8    string = "UTF-8"
	UNKNOWN string = "UNKNOWN"
)

func IsGBK(bytes []byte) bool {
	// fmt.Println("IsGBK")
	len := len(bytes)
	if len <= 0 {
		return false
	}
	var i int = 0
	for i < len {
		if bytes[i] <= 127 {
			i++
			continue
		} else {
			if bytes[i] >= 0x80 &&
				bytes[i] <= 0xFE &&
				bytes[i+1] >= 0x40 &&
				bytes[i+1] <= 0xfe &&
				bytes[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}
func enumMask(dat byte) (num int) {
	var mask byte = 0x80
	for i := 0; i < 8; i++ {
		if (dat & mask) == mask {
			num++
			mask >>= 1
		} else {
			break
		}
	}
	return num
}

/*
0XXX_XXXX
110X_XXXX 10XX_XXXX
1110_XXXX 10XX_XXXX 10XX_XXXX
1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
*/
func IsUTF8(bytes []byte) bool {
	// fmt.Println("IsUTF8")
	len := len(bytes)
	if len <= 0 {
		return false
	}
	var i int = 0
	for i < len {
		if bytes[i]&0x80 == 0x00 {
			i++
			continue
		} else {
			if num := enumMask(bytes[i]); num >= 2 {
				i++
				for j := 0; j < num-1; j++ {
					if (bytes[i] & 0xc0) != 0x80 {
						return false
					}
					i++
				}
			} else {
				return false
			}

		}
	}
	return true
}

func GetStrCodeType(data []byte) (ret string) {

	len := len(data)
	if len <= 0 {
		return UNKNOWN
	}
	if IsGBK(data) == true {
		return GBK
	} else {
		if IsUTF8(data) == true {
			return UTF8
		}
	}
	return UNKNOWN
}

func Test() {
	testStr := "我爱你 中国 i love you china"
	var Type string = GetStrCodeType([]byte(testStr))
	fmt.Println(Type)
	// 	simplifiedchinese.GBK.NewEncoder().Bytes()   //utf-8 转 gbk
	// simplifiedchinese.GBK.NewDecoder().Bytes()  //gbk 转 utf-8
	gbkdata, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(testStr))
	fmt.Println(GetStrCodeType(gbkdata))
	fmt.Println(string(gbkdata))
	utf8data, _ := simplifiedchinese.GBK.NewDecoder().Bytes(gbkdata)
	fmt.Println(GetStrCodeType(utf8data))
	fmt.Println(string(utf8data))
}
