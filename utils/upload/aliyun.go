package upload

import (
	"fmt"
	"github.com/Ocyss/douyin/internal/conf"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

func Aliyun(uploadName string, file io.Reader) (string, error) {
	var AliyunAccessKeyId = conf.Conf.Oss.AccessKeyID
	var AliyunAccessKeySecret = conf.Conf.Oss.AccessKeySecret
	var AliyunEndpoint = conf.Conf.Oss.Endpoint
	var AliyunBucketName = conf.Conf.Oss.BucketName

	client, err := oss.New(AliyunEndpoint, AliyunAccessKeyId, AliyunAccessKeySecret)
	if err != nil {
		return "", err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(AliyunBucketName)
	if err != nil {
		return "", err
	}
	err = bucket.PutObject(uploadName, file)
	if err != nil {
		return "", err
	}
	//拼接链接,默认使用https
	return fmt.Sprintf("https://%s.%s/%s", AliyunBucketName, AliyunEndpoint, uploadName), nil
}
