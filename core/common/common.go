package common

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/csv"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"sync"
)

var mutex sync.Mutex

//整形转换成字节
func IntToBytes(n int32) []byte {
	x := n
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

//把协议名称转为唯一协议编号
func ProtocalNumber(replacement string) int32 {
	var h int32
	h = 0
	for _, char := range []rune(replacement) {
		h = 31*h + int32(char)
	}
	return h
}

//获取协议名称
func GetProtoName(t proto.Message) string {
	x := proto.MessageName(t)
	proto_ := strings.Split(x, ".")
	if len(proto_) > 0 {
		return proto_[1]
	} else {
		return ""
	}
}

//连个字符串的key合并
func MergeMapString(varA map[string]string, varB map[string]string) map[string]string {
	data := make(map[string]string, len(varA)+len(varB))
	for k, v := range varA {
		data[k] = v
	}
	for k, v := range varB {
		data[k] = v
	}
	return data
}

//读取csv数据
func ReadCSVData(_file string) [][]string {
	mutex.Lock()
	csvFile, err := os.Open(_file)
	defer csvFile.Close()
	if nil != err {
		logrus.Error("fail to read csv file, err msg:", err.Error())
		return nil
	}
	data := csv.NewReader(bufio.NewReader(csvFile))
	var LineNumber int
	var dataRecords [][]string
	for {
		record, err := data.Read()
		// 如果读到文件的结尾，EOF的优先级居然比nil还高！
		if err == io.EOF {
			break
		} else if err != nil {
			return nil
		}
		// Read返回的是一个数组，它已经帮我们分割了，
		if LineNumber > 2 {
			s := make([]string, len(record))
			for i := 0; i < len(record); i++ {
				s[i] = record[i]
			}
			dataRecords = append(dataRecords, s)
		}
		LineNumber++
	}
	mutex.Unlock()
	return dataRecords
}

//读取txt文件
func ReadTXTData(_file string) []string {
	mutex.Lock()
	file, err := os.Open(_file)
	if nil != err {
		logrus.Error("fail to read txt file, err msg:", err.Error())
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var dataRecords []string
	//是否有下一行
	for scanner.Scan() {
		dataRecords = append(dataRecords, scanner.Text())
	}
	mutex.Unlock()
	return dataRecords
}
