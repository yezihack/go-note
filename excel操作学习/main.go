package main

import (
	//"github.com/yezihack/studyGo/excel操作学习/library"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/davecgh/go-spew/spew"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func main() {
	//list := library.GetDirFiles("/Users/wangzl/工作文档/评论数据分析/评论")
	//fmt.Println(len(list))
	//for _, val := range list {
	//	InsertDB(val)
	//}
	GetData()
}

type Comment struct {
	Id          string `json:"id"`
	State       string `json:"state"`
	UserId      string `json:"user_id"`
	ReceiverId  string `json:"receiver_id"`
	SenderId    string `json:"sender_id"`
	Refer       string `json:"refer"`
	IsMessage   string `json:"is_message"`
	Content     string `json:"content"`
	Created     string `json:"created"`
	RaiseFundId string `json:"raisefund_id"`
	Entity      string `json:"entity"`
	EntityId    string `json:"entity_id"`
	Context     string `json:"context"`
	Anonymous   string `json:"anonymous"`
}

func GetData() {
	obj := mysql.DataComment{}
	data, err1 := obj.Select(DB)
	if err1 != nil {
		panic(err1)
	}
	xlsx := excelize.NewFile()
	// Create a new sheet.
	index := xlsx.NewSheet("Sheet1")

	i := 1
	for _, item := range data {
		child := Comment{}
		err := json.Unmarshal([]byte(item.Data), &child)
		if err != nil {
			panic(err)
		}
		if child.Id == "id" {
			continue
		}
		xlsx.SetCellValue("Sheet1", "A"+fmt.Sprint(i), child.Id)
		xlsx.SetCellValue("Sheet1", "B"+fmt.Sprint(i), child.State)
		xlsx.SetCellValue("Sheet1", "C"+fmt.Sprint(i), child.UserId)
		xlsx.SetCellValue("Sheet1", "D"+fmt.Sprint(i), child.ReceiverId)
		xlsx.SetCellValue("Sheet1", "E"+fmt.Sprint(i), child.SenderId)
		xlsx.SetCellValue("Sheet1", "F"+fmt.Sprint(i), child.Refer)
		xlsx.SetCellValue("Sheet1", "G"+fmt.Sprint(i), child.IsMessage)
		xlsx.SetCellValue("Sheet1", "H"+fmt.Sprint(i), child.Content)
		xlsx.SetCellValue("Sheet1", "I"+fmt.Sprint(i), child.Created)
		xlsx.SetCellValue("Sheet1", "J"+fmt.Sprint(i), child.RaiseFundId)
		xlsx.SetCellValue("Sheet1", "K"+fmt.Sprint(i), child.Entity)
		xlsx.SetCellValue("Sheet1", "L"+fmt.Sprint(i), child.EntityId)
		xlsx.SetCellValue("Sheet1", "M"+fmt.Sprint(i), child.Context)
		xlsx.SetCellValue("Sheet1", "N"+fmt.Sprint(i), child.Anonymous)
		i++
	}
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./Data_Comment2.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func GetExcelData(path string) []*Comment {
	fmt.Println(path)
	if len(path) == 0 {
		return nil
	}
	xlsx, err := excelize.OpenFile(path)
	if err != nil {
		spew.Dump(err)
	}
	rows := xlsx.GetRows("SheetJS")

	if rows == nil {
		return nil
	}
	resultList := make([]*Comment, 0)
	for _, row := range rows {
		resultList = append(resultList, &Comment{
			Id:          row[0],
			State:       row[1],
			UserId:      row[2],
			ReceiverId:  row[3],
			SenderId:    row[4],
			Refer:       row[5],
			IsMessage:   row[6],
			Content:     row[7],
			Created:     row[8],
			RaiseFundId: row[9],
			Entity:      row[10],
			EntityId:    row[11],
			Context:     row[12],
			Anonymous:   row[13],
		})
	}
	return resultList
}

func init() {
	//%s:%s@tcp(%s:%s)/%s?charset=%s
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3308)/kindled?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	DB = db
}
