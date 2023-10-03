package dblite

import (
	"fmt"
	"os"
	"testing"
)

var (
	DataFile  *os.File
	IndexFile *os.File
	indexs    *CachedIndexs
)

func init() {
	indexs = &CachedIndexs{}
}

// testing all data storeing functions
func Test_Append(t *testing.T) {

	DataFile, _ = os.OpenFile("mok/data.page", os.O_RDWR|os.O_CREATE, 0644)
	IndexFile, _ = os.OpenFile("mok/primary.indexs", os.O_RDWR|os.O_CREATE, 0644) // why not "os.O_APPEND" ?

	var at int64
	for i := 0; i <= 11; i++ {

		data := "hello world ok " + fmt.Sprint(i)

		lenByte, err := Append(DataFile, data)

		if err != nil {
			fmt.Println("error is : ", err)
		}

		_ = Get(DataFile, at, lenByte)
		//fmt.Printf("Data is %s: \nlen byte is %d, at is %d\n\n", myData, lenByte, at)
		at += int64(lenByte)
	}
	fmt.Println("Done")
}

// testing all index functions
func Test_All_Index_Funcs(t *testing.T) {

	fmt.Println("Test newIndex func")
	state, _ := IndexFile.Stat()
	fmt.Println("index file size is : ", state.Size())
	fmt.Println("file name ", IndexFile.Name())
	fmt.Println("file name ", IndexFile)
	fmt.Println("=========================")

	// testing NewIndex func
	for i := 0; i < 1111; i++ {
		NewIndex(IndexFile, i, i)
		fmt.Println("at", i, "size", i)
	}

	// testing GetIndex func
	//"input 140 return 2800
	pageName, indx, size := GetIndex(IndexFile, 140)
	if pageName != "0" {
		t.Error("pageName must be 1")
	}
	if indx != 140 { // 2800
		t.Error("index must be 2800")
	}
	if size != 140 { // 2800
		t.Errorf("size is %d, must be %d", size, 140)
	}

	//"input 1111: 2220
	pageName, indx, size = GetIndex(IndexFile, 1111)
	if pageName != "1" {
		t.Error("pageName must be 1")
	}
	if indx != 1111 {
		t.Errorf("index s %d, must be %d\n", indx, 1111)
	}
	if size != 1111 {
		t.Error("size must be ", 1111)
	}

	// testing UpdateIndex func
	for i := 10; i <= 1111; i++ {
		UpdateIndex(IndexFile, i, int64(i+5), int64(i+10))
	}

	//"input 1111: 2220
	pageName, indx, size = GetIndex(IndexFile, 1111)
	if pageName != "1" {
		t.Error("pageName must be 1")
	}

	if indx != 1111+5 {
		t.Errorf("index s %d, must be %d\n", indx, 1111+5)
	}
	if size != 1121 {
		t.Error("size must be ", 1121)
	}

	// testing DeleteIndex func
	DeleteIndex(IndexFile, 1091)

	pageName, indx, size = GetIndex(IndexFile, 1091)
	if pageName != "1" {
		t.Error("pageName must be 1")
	}

	if indx != 0 {
		t.Errorf("index s %d, must be %d\n", indx, 0)
	}
	if size != 0 {
		t.Error("size must be ", size)
	}

}

func Test_LastIndex(t *testing.T) {
	lIndex := lastIndex("data.page")
	if lIndex == 0 {
		t.Errorf("lastindex is %d must be greater then 0\n", lIndex)
	}
	println("lastindex is ", lIndex)

	lastPageIndex := lastIndex("primary.indexs")
	if lastPageIndex == 0 {
		t.Errorf("lastindex is %d must be greater then 0\n", lIndex)
	}
	println("last Data index is ", lastPageIndex)

}

func Test_lastIndex(t *testing.T) {
	index := lastIndex("primary.indexs")
	fmt.Println("index is ", index)
	if index != 1112 {
		t.Errorf("index is %d, must be greater then 1112\n", index)
	}
}

func Test_finish(t *testing.T) {
	DataFile.Close()
	os.Remove("mok/primary.indexs")

	IndexFile.Close()
	os.Remove("mok/primary.indexs")

}
