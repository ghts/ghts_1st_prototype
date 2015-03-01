package common

import (
	"bytes"
	"log"
	"strings"
)

var 테스트_인덱스 int64 = 0
var 테스트_테이블_접두사 string = "test"
var 테스트_모드 bool = false
var 테스트_테이블_타입 string = "MEMORY"
var 실전_테이블_타입 string = "InnoDB"
var 종목정보_테이블 string = "stock_info"
var 일일가격정보_테이블 string = "stock_daily_price"
var 테이블_생성_인덱스 uint64 = 0

func F테스트_모드_시작()   { 테스트_모드 = true }
func F테스트_모드_종료()   { 테스트_모드 = false; 테스트_인덱스++ }
func F테스트_모드() bool { return 테스트_모드 }

func 접두사() string {
	if 테스트_모드 {
		return 테스트_테이블_접두사 + F정수64to문자열(테스트_인덱스) + "_"
	} else {
		return ""
	}
}

func 테이블_타입() string {
	if 테스트_모드 {
		return 테스트_테이블_타입
	} else {
		return 실전_테이블_타입
	}
}

func F종목정보_테이블() string   { return 접두사() + 종목정보_테이블 }
func F일일가격정보_테이블() string { return 접두사() + 일일가격정보_테이블 }

func F테이블_삭제(테이블_이름 string) {
	// 실수로 DB테이블을 지우는 것을 방지하기 위해서 테스트 모드일 때만 실행
	if strings.HasPrefix(테스트_테이블_접두사, 테이블_이름) {
		log.Println("F테이블_삭제() : 테스트용 테이블이 아니므로, 테이블 삭제를 취소합니다.", 테이블_이름)
		return
	}

	F_SQL실행("DROP TABLE IF EXISTS " + 테이블_이름)

	//log.Println("테이블_삭제() :", 테이블_이름)
}

func F종목정보_테이블_생성() error {
	sql := new(bytes.Buffer)
	sql.WriteString("CREATE TABLE IF NOT EXISTS " + F종목정보_테이블() + " (")
	sql.WriteString("id	INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,")
	sql.WriteString("issuer  VARCHAR(150),")
	sql.WriteString("name1	VARCHAR(150),")
	sql.WriteString("name2	VARCHAR(150),")
	sql.WriteString("code1	VARCHAR(40),")
	sql.WriteString("code2	VARCHAR(40),")
	sql.WriteString("market_status VARCHAR(250),")
	sql.WriteString("UNIQUE INDEX idx_code1 (code1),")
	sql.WriteString("INDEX idx_code2 (code2),")
	sql.WriteString("INDEX idx_name1 (name1)")
	sql.WriteString(") ENGINE=" + 테이블_타입() + " DEFAULT CHARSET = utf8mb4;")
	_, 에러 := F_SQL실행(sql.String())

	if 에러 != nil {
		//log.Println("F종목정보_테이블_생성() 에러.", 에러)
		return 에러
	}

	return nil
}

func F일일가격정보_테이블_생성() error {
	sql := new(bytes.Buffer)
	sql.WriteString("CREATE TABLE IF NOT EXISTS " + F일일가격정보_테이블() + " (")
	sql.WriteString("id	BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,")
	sql.WriteString("stock_info_id	INT UNSIGNED,")
	sql.WriteString("priced_on	DATE DEFAULT 0 NOT NULL,")
	sql.WriteString("open		DECIMAL(65,5) NOT NULL DEFAULT '0',")
	sql.WriteString("high		DECIMAL(65,5) NOT NULL DEFAULT '0',")
	sql.WriteString("low			DECIMAL(65,5) NOT NULL DEFAULT '0',")
	sql.WriteString("close		DECIMAL(65,5) NOT NULL DEFAULT '0',")
	sql.WriteString("adj_coeff   DOUBLE NOT NULL DEFAULT '0',")
	sql.WriteString("adj_open    DECIMAL(65,5) NOT NULL DEFAULT '0',")
	sql.WriteString("adj_high	DECIMAL(65,5) NOT NULL DEFAULT '0',")
	sql.WriteString("adj_low		DECIMAL(65,5) NOT NULL DEFAULT '0',")
	sql.WriteString("adj_close	DECIMAL(65,5) NOT NULL DEFAULT '0',")
	sql.WriteString("volumn      DOUBLE NOT NULL DEFAULT '0',")
	sql.WriteString("FOREIGN KEY (stock_info_id) REFERENCES " + F종목정보_테이블() + " (id),")
	sql.WriteString("INDEX idx_daily_price (stock_info_id, priced_on)")
	sql.WriteString(") ENGINE=" + 테이블_타입() + " DEFAULT CHARSET = utf8mb4;")
	_, 에러 := F_SQL실행(sql.String())

	if 에러 != nil {
		//log.Println("F일일가격정보_테이블_생성() 에러.", 에러)
		return 에러
	}

	return nil
}
