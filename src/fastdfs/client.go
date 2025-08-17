package fastdfs

import (
	"errors"
	"os"
	"strconv"
	"sync"
)

// storageServerInfo 存储服务器信息
type storageServerInfo struct {
	addrPort         string
	storagePathIndex byte
}

// Client FastDFS客户端
type Client struct {
	config       *TrackerStorageServerConfig
	trackerPools map[string]*tcpConnPool
	storagePools map[string]*tcpConnPool
	mu           sync.Mutex
}

// CreateClient 创建客户端
func CreateClient(config *TrackerStorageServerConfig) (*Client, error) {
	client := &Client{
		config:       config,
		trackerPools: make(map[string]*tcpConnPool),
		storagePools: make(map[string]*tcpConnPool),
	}

	for _, addr := range config.TrackerServer {
		pool, err := initTcpConnPool(addr, config.MaxConns)
		if err != nil {
			return nil, err
		}
		client.trackerPools[addr] = pool
	}

	return client, nil
}

// getTrackerConn 获取Tracker连接
func (c *Client) getTrackerConn() (*tcpConnPool, *tcpConnBaseInfo, error) {
	for _, pool := range c.trackerPools {
		conn, err := pool.get()
		if err == nil {
			return pool, conn, nil
		}
	}
	return nil, nil, errors.New(ERROR_CONN_POOL_NO_ACTIVE_CONN)
}

// getStorageInfo 获取存储服务器信息
func (c *Client) getStorageInfo() (*storageServerInfo, error) {
	trackerPool, trackerConn, err := c.getTrackerConn()
	if err != nil {
		return nil, err
	}
	defer trackerPool.put(trackerConn)

	req := &trackerTcpConn{
		header: header{
			cmd: TRACKER_PROTO_CMD_SERVICE_QUERY_STORE_WITHOUT_GROUP_ONE,
		},
	}

	if err := req.Send(trackerConn.Conn); err != nil {
		return nil, err
	}

	if err := req.Receive(trackerConn.Conn); err != nil {
		return nil, err
	}

	return &storageServerInfo{
		addrPort:         req.storageInfo.ipAddr + ":" + strconv.FormatInt(req.storageInfo.port, 10),
		storagePathIndex: req.storageInfo.storePathIndex,
	}, nil
}

// UploadByFileName 通过文件名上传
func (c *Client) UploadByFileName(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return "", err
	}

	ext := getFileExtName(fileName)
	fileInfo := &fileInfo{
		filePtr:     file,
		fileSize:    info.Size(),
		fileExtName: ext,
	}

	storageInfo, err := c.getStorageInfo()
	if err != nil {
		return "", err
	}

	uploadReq := &storageServerUploadHeaderBody{
		fileInfo:         fileInfo,
		storagePathIndex: storageInfo.storagePathIndex,
	}

	storagePool, storageConn, err := c.getStorageConn(storageInfo.addrPort)
	if err != nil {
		return "", err
	}
	defer storagePool.put(storageConn)

	if err := uploadReq.Send(storageConn.Conn); err != nil {
		return "", err
	}

	if err := uploadReq.Receive(storageConn.Conn); err != nil {
		return "", err
	}

	return uploadReq.fileId, nil
}

// getStorageConn 获取存储服务器连接
func (c *Client) getStorageConn(addr string) (*tcpConnPool, *tcpConnBaseInfo, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	pool, exists := c.storagePools[addr]
	if !exists {
		var err error
		pool, err = initTcpConnPool(addr, c.config.MaxConns)
		if err != nil {
			return nil, nil, err
		}
		c.storagePools[addr] = pool
	}

	conn, err := pool.get()
	return pool, conn, err
}

// getFileExtName 获取文件扩展名
func getFileExtName(fileName string) string {
	idx := len(fileName)
	for i := len(fileName) - 1; i >= 0; i-- {
		if fileName[i] == '.' {
			idx = i
			break
		}
	}
	if idx == len(fileName) {
		return ""
	}
	return fileName[idx+1:]
}

// Destroy 销毁客户端
func (c *Client) Destroy() {
	for _, pool := range c.trackerPools {
		pool.Destroy()
	}
	for _, pool := range c.storagePools {
		pool.Destroy()
	}
}
