package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	go_ora "github.com/sijms/go-ora"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Invalid arguments ")
		return
	}

	phonenumber := os.Args[1]
	date := os.Args[2]

	iphonenumber, err := strconv.Atoi(phonenumber)

	fmt.Println("phonenumber : " + phonenumber)
	fmt.Println("date : " + date)

	//username/password@hostname:port/service_name
	//
	//ora, err := sql.Open("goracle", "oraspam/pltspam10!@asfs")
	//ora, err := sql.Open("ora", "oraspam/pltspam10!@asfs:1521/ORASPAM")

	server := "210.217.178.121"
	port := 1521 // 오라클 서버의 포트 번호
	serviceName := "spamdb"
	username := "oraspam"
	password := "pltspam10!"

	// 오라클 연결 문자열 생성
	connStr := go_ora.BuildUrl(server, port, serviceName, username, password, nil)

	// 오라클 데이터베이스에 연결
	ora, err := sql.Open("oracle", connStr)

	if err != nil {
		panic(err)
	}
	defer ora.Close()

	tblName := fmt.Sprintf("TM_SFS_SMS_%03d ", iphonenumber%256+1)

	//fmt.Printf("tblName : %s\n", tblName)

	sql := fmt.Sprintf("SELECT CUST_NUM,SMS_CLC,SRC_NUM,CB_NUM,SMS_KIND,TEL_ID,SMS_LENG, RCV_DT from %s WHERE SUBSTR(RCV_DT,1,8) = '%s' AND CUST_NUM = '%s'", tblName, date, phonenumber)

	//SMS_MSG : utf-8 -> euc-kr test
	//여러줄 : 역따옴표.

	//fmt.Printf("sql : %s\n", sql)

	// 쿼리 실행
	rows, err := ora.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Printf("\n================================================================================================================================================================================================\n\n")

	cnt := 0
	// 결과 처리
	for rows.Next() {
		var col1, col2, col3, col4, col5, col6, col7, col8 string
		if err := rows.Scan(&col1, &col2, &col3, &col4, &col5, &col6, &col7, &col8); err != nil {
			log.Fatal(err)
		}
		if cnt == 0 {
			fmt.Printf("[%s] \t[%s] \t[%s] \t[%s] \t[%s] \t[%s] \t[%s] \t[%s] \t[%s]", "NO", "CUST_NUM", "SMS_CLC", "SRC_NUM", "CB_NUM", "SMS_KIND", "TEL_ID", "SMS_LENG", "RCV_DT")
			fmt.Printf("\n\n")
		}
		cnt++
		fmt.Printf("%03d \t%s \t%s \t%s \t%s \t%s \t\t%s \t%s \t\t%s\n", cnt, col1, col2, col3, col4, col5, col6, col7, col8)
	}

	fmt.Printf("\n================================================================================================================================================================================================\n\n")

	// 에러 처리
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

///Linux : home/sfs/go/src/logviewer/logviewer 01059132451 20240102
///Windows : ./logviewer 01059132451 20240102
