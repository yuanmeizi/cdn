package main

// 存储相关功能的引入包只有这两个，后面不再赘述
import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	accessKey string = "ePHgWBBpBj6IZCd2cNuRBCaOgNEMFEctdWcGnAt2"
	secretKey string = "1ZXwosZWQBH5n01NtoPOnQ4B4TUTZ-AcTIY3WWzR"
	bucket    string = "aschv"
)

func main() {

	//设置访问路由
	http.HandleFunc("/uptoken", clientUpload)
	http.HandleFunc("/upload", serviveUpload)
	//创建监听端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Panic("ListenAndServe:	", err)
	}

}

// 生成客户端上传凭证
func clientUpload(w http.ResponseWriter, r *http.Request) {
	log.Println("收到请求")

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	fmt.Println("生成的客户端上传凭证是:", upToken)
	w.Write([]byte(upToken))
}

// 服务器上传
func serviveUpload(w http.ResponseWriter, r *http.Request) {
	localFile := "./aaa.png"
	key := strconv.FormatInt(time.Now().UnixNano(), 10) + ".png"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		// 	Params: map[string]string{
		// 		"x:name": "github logo",
		// 	},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
	w.Write([]byte("hash:" + ret.Hash + ";key:" + ret.Key))
}
