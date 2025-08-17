package fastdfs

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
)

type storageInfo struct {
	ipAddr         string
	port           int64
	storePathIndex byte
}

type trackerTcpConn struct {
	header
	groupName      string
	remoteFilename string
	storageInfo    storageInfo
}

// @brief 发送数据给tracker
func (t *trackerTcpConn) Send(conn net.Conn) error {
	if err := t.header.sendHeader(conn); err != nil {
		return err
	}
	return nil
}

func (t *trackerTcpConn) Receive(conn net.Conn) error {
	if err1 := t.receiveHeader(conn); err1 != nil {
		return errors.New(ERROR_HEADER_RECEV_ERROR + err1.Error())
	}

	buf := make([]byte, t.pkgLen)
	if _, err2 := conn.Read(buf); err2 != nil {
		return err2
	}

	// 开始解析响应
	t.groupName = string(getBytesByPosition(buf, 0, 16))
	t.storageInfo.ipAddr = string(getBytesByPosition(buf, 16, 46))
	t.storageInfo.port = int64(binary.BigEndian.Uint32(getBytesByPosition(buf, 65, 4)))
	t.storageInfo.storePathIndex = getBytesByPosition(buf, 39, 1)[0]

	return nil
}

func getBytesByPosition(bys []byte, start, num int) []byte {
	newBytes := bys[start:]
	endPosition := bytes.IndexByte(newBytes, 0x0)
	if endPosition > 0 {
		num = endPosition
	}
	return newBytes[:num]
}
