package conf

import "os"

const Version = "7.12.1"

const (
	CONTENT_TYPE_JSON      = "application/json"
	CONTENT_TYPE_FORM      = "application/x-www-form-urlencoded"
	CONTENT_TYPE_OCTET     = "application/octet-stream"
	CONTENT_TYPE_MULTIPART = "multipart/form-data"

	disableQiniuTimestampSignatureEnvKey = "DISABLE_QINIU_TIMESTAMP_SIGNATURE"
)

func IsDisableQiniuTimestampSignature() bool {
	return os.Getenv(disableQiniuTimestampSignatureEnvKey) == "true"
}
