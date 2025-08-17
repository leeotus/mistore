package fastdfs

import "time"

const (
	TCP_HEADER_LEN                                          = 10
	TRACKER_PROTO_CMD_RESP                                  = 100
	TRACKER_PROTO_CMD_SERVICE_QUERY_STORE_WITHOUT_GROUP_ONE = 101
	STORAGE_PROTO_CMD_UPLOAD_FILE                           = 11
	STORAGE_PROTO_CMD_DELETE_FILE                           = 12
	FDFS_GROUP_NAME_FIX_LEN                                 = 16
	FILE_EXTNAME_FIX_LEN                                    = 6
	TCP_CONN_TIMEOUT                                        = time.Second * 10
)

const (
	ERROR_HEADER_RECEV_ERROR       = "收到的消息头有误"
	ERROR_CONN_POOL_NO_ACTIVE_CONN = "连接池无有效连接"
)
