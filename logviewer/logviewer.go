package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	_ "gopkg.in/goracle.v2"

	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
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

	ora, err := sql.Open("goracle", "oraspam/pltspam10!@asfs")
	if err != nil {
		panic(err)
	}
	defer ora.Close()

	tblName := fmt.Sprintf("TM_SFS_SMS_%03d ", iphonenumber%256+1)

	//fmt.Printf("tblName : %s\n", tblName)

	///*
	sql := fmt.Sprintf("SELECT CUST_NUM,SMS_CLC,SRC_NUM,CB_NUM,SMS_KIND,TEL_ID,SMS_LENG, RCV_DT from %s WHERE SUBSTR(RCV_DT,1,8) = '%s' AND CUST_NUM = '%s'", tblName, date, phonenumber)
	//*/

	/* SMS_MSG euc-kr Test ( SMS_KIND -> SMS_MSG  )
	sql := fmt.Sprintf("SELECT CUST_NUM,SMS_CLC,SRC_NUM,CB_NUM,SMS_MSG,TEL_ID,SMS_LENG, RCV_DT from %s WHERE SUBSTR(RCV_DT,1,8) = '%s' AND CUST_NUM = '%s'", tblName, date, phonenumber)
	*/

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

		//SMS_MSG Test
		/*
			eucKRBytes, err := utf8ToEUCKR(col5)
			if err != nil {
				fmt.Println("변환 오류:", err)
				return
			}

			fmt.Printf("%03d \t%s \t%s \t%s \t%s \t%s \t\t%s \t%s \t\t%s\n", cnt, col1, col2, col3, col4, eucKRBytes, col6, col7, col8)
		*/
	}

	fmt.Printf("\n================================================================================================================================================================================================\n\n")

	// 에러 처리
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// UTF-8에서 EUC-KR로 변환하는 함수
func utf8ToEUCKR(utf8String string) ([]byte, error) {
	utf8Bytes := []byte(utf8String)

	// EUC-KR 인코딩 설정
	eucKREncoding := korean.EUCKR.NewEncoder()

	// 변환 수행
	eucKRBytes, _, err := transform.Bytes(eucKREncoding, utf8Bytes)
	if err != nil {
		return nil, err
	}

	return eucKRBytes, nil
}

//실행 명령어
///home/sfs/go/src/logviewer/logviewer 01059132451 20240102

//테스트 필요한 항목.
//1.(x) SMS_MSG : utf-8 -> euc-kr test
//2. 여러줄 : 역따옴표.
