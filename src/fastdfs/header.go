package fastdfs

import (
	"encoding/binary"
	"net"
)

/**
 * @brief fastdfs协议包头
 */

/**
 * @brief FastDFS采用二进制TCP通信协议,一个数据由包头(header)+包体(body)组成
 * @pkgLen body长度,不包含header,只是body的长度
 * @cmd 命令码
 * @status 状态码,0表示成功,非0表示失败(UNIX错误码)
 */
type header struct {
	pkgLen int64
	cmd    byte
	status byte
}

/**
 * @brief 发送数据包头
 */
func (h *header) sendHeader(conn net.Conn) error {
	buf := make([]byte, TCP_HEADER_LEN)
	binary.BigEndian.PutUint64(buf[:8], uint64(h.pkgLen))
	buf[8] = h.cmd
	buf[9] = h.status
	_, err := conn.Write(buf)
	return err
}

// @brief 接收数据包头
func (h *header) receiveHeader(conn net.Conn) error {
	buf := make([]byte, TCP_HEADER_LEN)
	_, err := conn.Read(buf)
	if err != nil {
		return err
	}
	h.pkgLen = int64(binary.BigEndian.Uint64(buf[:8]))
	h.cmd = buf[8]
	h.status = buf[9]
	return nil
}
