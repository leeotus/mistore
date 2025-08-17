package fastdfs

import (
	"bytes"
	"encoding/binary"
	"net"
	"os"
)

// 文件信息类
type fileInfo struct {
	filePtr     *os.File
	buffer      []byte
	fileSize    int64
	fileExtName string
}

// 存储服务器上传结构体
type storageServerUploadHeaderBody struct {
	header
	fileInfo         *fileInfo
	storagePathIndex byte
	fileId           string
}

// Send 发送上传请求
func (s *storageServerUploadHeaderBody) Send(conn net.Conn) error {
	s.header.pkgLen = s.fileInfo.fileSize + 15
	s.header.cmd = STORAGE_PROTO_CMD_UPLOAD_FILE
	s.header.status = 0

	if err := s.header.sendHeader(conn); err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	buffer.WriteByte(s.storagePathIndex)

	if err := binary.Write(buffer, binary.BigEndian, s.fileInfo.fileSize); err != nil {
		return err
	}

	buffer.Write(specialFileExtNameConvBytes(s.fileInfo.fileExtName))
	if _, err := conn.Write(buffer.Bytes()); err != nil {
		return err
	}

	// 发送文件内容
	if s.fileInfo.filePtr != nil {
		if _, err := sendFileContent(s.fileInfo.filePtr, conn); err != nil {
			return err
		}
	} else {
		if _, err := conn.Write(s.fileInfo.buffer); err != nil {
			return err
		}
	}

	return nil
}

// Receive 接收上传响应
func (s *storageServerUploadHeaderBody) Receive(conn net.Conn) error {
	if err := s.header.receiveHeader(conn); err != nil {
		return err
	}

	buf := make([]byte, s.pkgLen)
	if _, err := conn.Read(buf); err != nil {
		return err
	}

	s.fileId = string(getBytesByPosition(buf, 0, 16)) + "/" + string(getBytesByPosition(buf, 16, int(s.pkgLen)-16))
	return nil
}

// specialFileExtNameConvBytes 转换文件扩展名为字节
func specialFileExtNameConvBytes(extName string) []byte {
	b := make([]byte, FILE_EXTNAME_FIX_LEN)
	copy(b, extName)
	return b
}

// sendFileContent 发送文件内容
func sendFileContent(file *os.File, conn net.Conn) (int64, error) {
	buf := make([]byte, 4096)
	var total int64
	for {
		n, err := file.Read(buf)
		if n > 0 {
			if _, err := conn.Write(buf[:n]); err != nil {
				return total, err
			}
			total += int64(n)
		}
		if err != nil {
			break
		}
	}
	return total, nil
}
