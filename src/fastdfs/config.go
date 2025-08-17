package fastdfs

// @brief 配置结构体
type TrackerStorageServerConfig struct {
	TrackerServer []string // tracker的ip地址
	MaxConns      int      // 最大连接数
}
