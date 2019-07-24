package temp

import (
	"bufio"
	"encoding/csv"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sync"
)

var mutex sync.Mutex


//读取csv数据
func ReadCSVData(fileName string) [][]string {
	mutex.Lock()
	csvFile, err := os.Open(ConfigCacheData.YamlConfigData.ExcelPath+fileName+".csv")
	defer csvFile.Close()
	if nil != err{
		logrus.Error("fail to read csv file, err msg:",err.Error())
		return nil
	}
	data := csv.NewReader(bufio.NewReader(csvFile))
	var LineNumber int
	var dataRecords [][]string
	for{
		record, err := data.Read()
		// 如果读到文件的结尾，EOF的优先级居然比nil还高！
		if err == io.EOF {
			break
		} else if err != nil {
			return nil
		}
		// Read返回的是一个数组，它已经帮我们分割了，
		if LineNumber > 2{
			s := make([]string,len(record))
			for i := 0; i < len(record); i++ {
				s[i] = record[i]
			}
			dataRecords = append(dataRecords,s)
		}
		LineNumber++
	}
	mutex.Unlock()
	return dataRecords
}

//读取txt文件
func ReadTXTData(fileName string)[]string{
	mutex.Lock()
	file, err := os.Open(ConfigCacheData.YamlConfigData.ExcelPath+fileName+".txt")
	if nil != err{
		logrus.Error("fail to read txt file, err msg:",err.Error())
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var dataRecords []string
	//是否有下一行
	for scanner.Scan() {
		dataRecords = append(dataRecords,scanner.Text() )
	}
	mutex.Unlock()
	return dataRecords
}