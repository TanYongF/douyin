package global

import (
	"crypto/md5"
	"encoding/hex"

	"go.uber.org/zap"
)

var (
	Secrete = "douyinxiangmu" //salt for MD5 encryption
	Logger  *zap.Logger       //global log
)

func Md5(pwd string) string {
	str := fmt.Sprintf("%c%c%c%s%c%c%c", Secrete[9], Secrete[0], Secrete[4],
		pwd, Secrete[3], Secrete[12], Secrete[6]) //combine salt and password.
	data := []byte(str)              //convert string to byte slice.
	hash := md5.Sum(data)             //MD5 encryption
	md5str := fmt.Sprintf("%x", hash) //convert []byte to hexadecimal.
	return md5str
}
