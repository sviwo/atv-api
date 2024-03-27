package file

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/gutil"
	"sviwo/internal/consts/enums"
	"sync"
)

type singleBucket struct {
	*oss.Bucket
}

var once sync.Once

var singleBucketInstance *singleBucket

func getInstance() *singleBucket {
	if singleBucketInstance == nil {
		once.Do(
			func() {
				// 创建OSSClient实例。
				ctx := gctx.GetInitCtx()
				ct, err := oss.New(
					g.Cfg().MustGet(ctx, "aliyun.oss.endpoint").String(),
					g.Cfg().MustGet(ctx, "aliyun.accessKeyID").String(),
					g.Cfg().MustGet(ctx, "aliyun.accessKeySecret").String(),
				)
				if err != nil {
					panic(err)
				}
				// 设置存储空间的读写权限为公共读。
				err = ct.SetBucketACL(g.Cfg().MustGet(ctx, "aliyun.oss.bucket").String(), oss.ACLPublicRead)
				if err != nil {
					panic(err)
				}
				bt, err := ct.Bucket(g.Cfg().MustGet(ctx, "aliyun.oss.bucket").String())
				if err != nil {
					panic(err)
				}
				singleBucketInstance = &singleBucket{Bucket: bt}
			})
	}
	return singleBucketInstance
}

func init() {
	getInstance()
}

/*
UploadFile 上传文件
*/
func UploadFile(file *ghttp.UploadFile) (uri string, err error) {
	if gutil.IsEmpty(file.Filename) {
		return "", gerror.NewCode(enums.RequestMissingParam)
	}
	open, err := file.Open()
	if err != nil {
		panic(err)
	}
	filename := getPath("", gstr.SubStrFrom(file.Filename, "."))
	err = getInstance().PutObject(filename, open)
	if err != nil {
		panic(err)
	}
	return filename, err
}

/*
DeleteFile 删除文件
*/
func DeleteFile(fileName string) error {
	if gutil.IsEmpty(fileName) {
		return gerror.NewCode(enums.RequestMissingParam)
	}
	err := getInstance().DeleteObject(fileName)
	if err != nil {
		panic(err)
	}
	return nil
}

/*
 * 文件路径
 * @param prefix 前缀
 * @param suffix 后缀
 * @return 返回上传路径
 */
func getPath(prefix string, suffix string) string {
	// 文件路径
	path := gtime.Now().Format("Ymd") + "/" + grand.S(32)
	if prefix != "" {
		path = prefix + "/" + path
	}
	return path + suffix
}
