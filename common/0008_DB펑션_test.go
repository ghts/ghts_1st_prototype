package common

import (
	"bytes"
	"testing"
	"time"
	//"log"
)

func TestF데이터베이스_연결(테스트 *testing.T) {
	// 테스트하는 더 좋은 방법이 있을까나?
	데이터베이스, 에러 := F데이터베이스_연결()
	defer 데이터베이스.Close()

	if 에러 != nil {
		테스트.Error("F데이터베이스_연결()가 데이터베이스 연결.")
	}

}

func TestF_SQL실행(테스트 *testing.T) {
	테스트용_테이블 := "test_table"

	sql := new(bytes.Buffer)
	sql.WriteString("DROP TABLE IF EXISTS " + 테스트용_테이블)
	_, 에러 := F_SQL실행(sql.String())
	if 에러 != nil {
		//테스트.Error("TestF_SQL실행() DROP TABLE 에러 : ", sql.String(), 에러)
	}

	sql = new(bytes.Buffer)
	sql.WriteString("CREATE TABLE " + 테스트용_테이블 + " (")
	sql.WriteString("int_column	INTEGER,")
	sql.WriteString("string_column	VARCHAR(200)")
	sql.WriteString(") ENGINE=MEMORY DEFAULT CHARSET = utf8mb4;")
	_, 에러 = F_SQL실행(sql.String())
	if 에러 != nil {
		테스트.Error("TestF_SQL실행() CREATE TABLE 에러 : ", sql.String(), 에러)
	}

	sql = new(bytes.Buffer)
	sql.WriteString("INSERT INTO ")
	sql.WriteString(테스트용_테이블 + " (")
	sql.WriteString("int_column,")
	sql.WriteString("string_column")
	sql.WriteString(") VALUES (")
	sql.WriteString("?,?")
	sql.WriteString(")")
	결과, 에러 := F_SQL실행(sql.String(), 1, "first")
	if 에러 != nil {
		테스트.Error("TestF_SQL실행() INSERT 에러 1 : ", sql.String(), 에러)
	}

	영향받은_행_갯수, 에러 := 결과.RowsAffected()
	if 에러 != nil {
		테스트.Error("TestF_SQL실행() INSERT RowsAffected() 에러 1", 에러)
	}

	if 영향받은_행_갯수 != 1 {
		테스트.Error("TestF_SQL실행() INSERT RowsAffected() 에러 2", 에러)
	}

	결과, 에러 = F_SQL실행(sql.String(), 2, "second")
	if 에러 != nil {
		테스트.Error("TestF_SQL실행() INSERT 에러 2 : ", sql.String(), 에러)
	}

	영향받은_행_갯수, 에러 = 결과.RowsAffected()
	if 에러 != nil {
		테스트.Error("TestF_SQL실행() INSERT RowsAffected() 에러 3", 에러)
	}

	if 영향받은_행_갯수 != 1 {
		테스트.Error("TestF_SQL실행() INSERT RowsAffected() 에러 4", 에러)
	}

	sql = new(bytes.Buffer)
	sql.WriteString("SELECT COUNT(*) FROM " + 테스트용_테이블)
	행_갯수, 에러 := F_SQL질의_정수(sql.String())
	if 에러 != nil {
		테스트.Error("TestF_SQL실행() 행 갯수 질의 에러 : ", sql.String(), 에러)
	}

	if 행_갯수 != 2 {
		테스트.Error("TestF_SQL실행() 행 갯수가 예상과 다름 : ", sql.String(), 에러)
	}

	sql = new(bytes.Buffer)
	sql.WriteString("DELETE FROM " + 테스트용_테이블)
	결과, 에러 = F_SQL실행(sql.String())
	if 에러 != nil {
		테스트.Error("TestF_SQL실행() DELETE 에러 : ", sql.String(), 에러)
	}

	영향받은_행_갯수, 에러 = 결과.RowsAffected()
	if 에러 != nil {
		테스트.Error("TestF_SQL실행() DELETE RowsAffected() 에러 1", 에러)
	}

	if 영향받은_행_갯수 != 2 {
		테스트.Error("TestF_SQL실행() DELETE RowsAffected() 에러 2", 에러)
	}

	sql = new(bytes.Buffer)
	sql.WriteString("SELECT COUNT(*) FROM " + 테스트용_테이블)
	행_갯수, 에러 = F_SQL질의_정수(sql.String())
	if 에러 != nil {
		테스트.Error("TestF_SQL실행() 행 갯수 질의 에러 2 : ", sql.String(), 에러)
	}

	if 행_갯수 != 0 {
		테스트.Error("TestF_SQL실행() 행 갯수 질의 에러 2 : ", sql.String(), 에러)
	}

	sql = new(bytes.Buffer)
	sql.WriteString("DROP TABLE " + 테스트용_테이블)
	_, 에러 = F_SQL실행(sql.String())
	if 에러 != nil {
		테스트.Error("TestF_SQL실행() DROP TABLE 에러 2 : ", sql.String(), 에러)
	}
}

func TestF_SQL질의_부호없는_정수(테스트 *testing.T) {
	테스트용_테이블 := "test_table"

	sql := new(bytes.Buffer)
	sql.WriteString("CREATE TABLE " + 테스트용_테이블 + " (")
	sql.WriteString("uint_column	INTEGER UNSIGNED,")
	sql.WriteString("string_column	VARCHAR(200)")
	sql.WriteString(") ENGINE=MEMORY DEFAULT CHARSET = utf8mb4;")
	_, 에러 := F_SQL실행(sql.String())
	if 에러 != nil {
		테스트.Error("TestF_SQL질의_부호없는_정수() CREATE TABLE 에러 : ", sql.String(), 에러)
	}
	
	defer F테이블_삭제(테스트용_테이블)

	sql = new(bytes.Buffer)
	sql.WriteString("INSERT INTO ")
	sql.WriteString(테스트용_테이블 + " (")
	sql.WriteString("uint_column,")
	sql.WriteString("string_column")
	sql.WriteString(") VALUES (")
	sql.WriteString("?,?")
	sql.WriteString(")")
	
	var 부호없는_정수64 uint64 = 1
	_, 에러 = F_SQL실행(sql.String(), 부호없는_정수64, "first")
	if 에러 != nil {
		테스트.Error("TestF_SQL질의_부호없는_정수() INSERT 에러 1 : ", sql.String(), 에러)
	}

	부호없는_정수64 = 2
	_, 에러 = F_SQL실행(sql.String(), 부호없는_정수64, "second")
	if 에러 != nil {
		테스트.Error("TestF_SQL질의_부호없는_정수() INSERT 에러 2 : ", sql.String(), 에러)
	}

	sql = new(bytes.Buffer)
	sql.WriteString("SELECT ")
	sql.WriteString("	uint_column ")
	sql.WriteString("FROM " + 테스트용_테이블 + " ")
	sql.WriteString("WHERE string_column = ?")
	결과값, 에러 := F_SQL질의_부호없는_정수(sql.String(), "first")
	if 에러 != nil {
		테스트.Error("TestF_SQL질의_부호없는_정수() : 결과값 질의 에러 : ", 에러)
	}

	if 결과값 != 1 {
		테스트.Errorf("TestF_SQL질의_부호없는_정수() : 결과값 불일치. 예상값 1, 실제값 %v.", 결과값)
	}
}

func TestF_SQL질의_정수(테스트 *testing.T) {
	테스트용_테이블 := "test_table"

	sql := new(bytes.Buffer)
	sql.WriteString("CREATE TABLE " + 테스트용_테이블 + " (")
	sql.WriteString("int_column	INTEGER,")
	sql.WriteString("string_column	VARCHAR(200)")
	sql.WriteString(") ENGINE=MEMORY DEFAULT CHARSET = utf8mb4;")
	_, 에러 := F_SQL실행(sql.String())
	if 에러 != nil {
		테스트.Error("TestF_SQL질의_정수() CREATE TABLE 에러 : ", sql.String(), 에러)
	}
	
	defer F테이블_삭제(테스트용_테이블)

	sql = new(bytes.Buffer)
	sql.WriteString("INSERT INTO ")
	sql.WriteString(테스트용_테이블 + " (")
	sql.WriteString("int_column,")
	sql.WriteString("string_column")
	sql.WriteString(") VALUES (")
	sql.WriteString("?,?")
	sql.WriteString(")")
	_, 에러 = F_SQL실행(sql.String(), 1, "first")
	if 에러 != nil {
		테스트.Error("TestF_SQL질의_정수() INSERT 에러 1 : ", sql.String(), 에러)
	}

	_, 에러 = F_SQL실행(sql.String(), 2, "second")
	if 에러 != nil {
		테스트.Error("TestF_SQL질의_정수() INSERT 에러 2 : ", sql.String(), 에러)
	}

	sql = new(bytes.Buffer)
	sql.WriteString("SELECT COUNT(*) FROM " + 테스트용_테이블)
	행_갯수, 에러 := F_SQL질의_정수(sql.String())
	if 에러 != nil {
		테스트.Error("TestF_SQL질의_정수() 행 갯수 질의 에러 : ", sql.String(), 에러)
	}

	if 행_갯수 != 2 {
		테스트.Error("TestF_SQL질의_정수() 행 갯수가 예상과 다름 : ", sql.String(), 에러)
	}
}

func TestF_SQL질의_실수(테스트 *testing.T) {
	테스트용_테이블 := "test_table"

	sql := new(bytes.Buffer)
	sql.WriteString("CREATE TABLE " + 테스트용_테이블 + " (")
	sql.WriteString("float_column	FLOAT,")
	sql.WriteString("string_column	VARCHAR(200)")
	sql.WriteString(") ENGINE=MEMORY DEFAULT CHARSET = utf8mb4;")
	_, 에러 := F_SQL실행(sql.String())
	if 에러 != nil {
		테스트.Error("TestF_SQL질의_실수() CREATE TABLE 에러 : ", sql.String(), 에러)
	}
	
	defer F테이블_삭제(테스트용_테이블)

	sql = new(bytes.Buffer)
	sql.WriteString("INSERT INTO ")
	sql.WriteString(테스트용_테이블 + " (")
	sql.WriteString("float_column,")
	sql.WriteString("string_column")
	sql.WriteString(") VALUES (")
	sql.WriteString("?,?")
	sql.WriteString(")")
	_, 에러 = F_SQL실행(sql.String(), 1.1, "one point one")
	if 에러 != nil {
		테스트.Error("TestF_SQL질의_실수() INSERT 에러 1 : ", sql.String(), 에러)
	}

	_, 에러 = F_SQL실행(sql.String(), 2.2, "이쩜이")
	if 에러 != nil {
		테스트.Error("TestF_SQL질의_실수() INSERT 에러 2 : ", sql.String(), 에러)
	}

	sql = new(bytes.Buffer)
	sql.WriteString("SELECT ")
	sql.WriteString("	float_column ")
	sql.WriteString("FROM " + 테스트용_테이블 + " ")
	sql.WriteString("WHERE string_column = ?")
	결과값, 에러 := F_SQL질의_실수(sql.String(), "이쩜이")
	if 에러 != nil {
		테스트.Error("TestF_SQL질의_실수() : 결과값 질의 에러.", 에러)
	}
	if F반올림_통화(결과값) != 2.2 {
		테스트.Errorf("TestF_SQL질의_실수() : 결과값 불일치. 예상값 2.2,실제값 %v.", F반올림_통화(결과값))
	}
}

func TestF_SQL질의_문자열(테스트 *testing.T) {
	테스트용_테이블 := "test_table"

	sql := new(bytes.Buffer)
	sql.WriteString("CREATE TABLE " + 테스트용_테이블 + " (")
	sql.WriteString("int_column	INTEGER,")
	sql.WriteString("string_column	VARCHAR(200)")
	sql.WriteString(") ENGINE=MEMORY DEFAULT CHARSET = utf8mb4;")
	_, 에러 := F_SQL실행(sql.String())
	if 에러 != nil {
		테스트.Error("common.TestF_SQL질의_문자열() CREATE TABLE 에러 : ", 에러, sql.String())
	}

	defer F테이블_삭제(테스트용_테이블)

	sql = new(bytes.Buffer)
	sql.WriteString("INSERT INTO ")
	sql.WriteString(테스트용_테이블 + " (")
	sql.WriteString("int_column,")
	sql.WriteString("string_column")
	sql.WriteString(") VALUES (")
	sql.WriteString("?,?")
	sql.WriteString(")")
	_, 에러 = F_SQL실행(sql.String(), 1, "one")
	if 에러 != nil {
		테스트.Error("common.TestF_SQL질의_문자열() : INSERT문 에러.", 에러, sql.String())
	}

	질의값, 에러 := F_SQL질의_문자열("SELECT string_column FROM " + 테스트용_테이블 + " WHERE int_column = ? AND string_column = ?", 1, "one")
	if 에러 != nil {
		테스트.Error("TestF_SQL질의_문자열() : 질의 에러 : ", 에러, sql.String())
	}

	if 질의값 != "one" {
		테스트.Errorf("TestF_SQL질의_문자열() : 값 불일치. 예상값 'one', 실제값 %v", 질의값)
	}
}

func TestF_SQL질의_시점(테스트 *testing.T) {
	테스트용_테이블 := "test_table"

	sql := new(bytes.Buffer)
	sql.WriteString("CREATE TABLE " + 테스트용_테이블 + " (")
	sql.WriteString("date_column	DATE")
	sql.WriteString(") ENGINE=MEMORY DEFAULT CHARSET = utf8mb4;")
	_, 에러 := F_SQL실행(sql.String())
	if 에러 != nil {
		테스트.Error("common.TestF_SQL질의() CREATE TABLE 에러 : ", 에러, sql.String())
	}

	defer F테이블_삭제(테스트용_테이블)

	일자, _ := F문자열2일자("2000-01-05")

	sql = new(bytes.Buffer)
	sql.WriteString("INSERT INTO ")
	sql.WriteString(테스트용_테이블 + " (")
	sql.WriteString("date_column")
	sql.WriteString(") VALUES (")
	sql.WriteString("?")
	sql.WriteString(")")
	_, 에러 = F_SQL실행(sql.String(), 일자)
	if 에러 != nil {
		테스트.Error("common.TestF_SQL질의_정수() : INSERT문 에러.", 에러, sql.String(), 일자)
	}

	질의값, 에러 := F_SQL질의_시점("SELECT date_column FROM " + 테스트용_테이블)
	if 에러 != nil {
		테스트.Error("TestF_SQL질의_시점() : 질의 에러 : ", 에러, sql.String())
	}
	if !질의값.Equal(일자) {
		테스트.Errorf("TestF_SQL질의_시점() : 값 불일치. 예상값 %v, 실제값 %v", 일자, 질의값)
	}
}

func TestF_SQL질의(테스트 *testing.T) {	
	테스트용_테이블 := "test_table"

	sql := new(bytes.Buffer)
	sql.WriteString("CREATE TABLE " + 테스트용_테이블 + " (")
	sql.WriteString("uint_column INTEGER UNSIGNED,")
	sql.WriteString("int_column	INTEGER,")
	sql.WriteString("float_column	FLOAT,")
	sql.WriteString("string_column	VARCHAR(200),")
	sql.WriteString("date_column	DATE")
	sql.WriteString(") ENGINE=MEMORY DEFAULT CHARSET = utf8mb4;")
	_, 에러 := F_SQL실행(sql.String())
	if 에러 != nil {
		테스트.Error("common.TestF_SQL질의() CREATE TABLE 에러 : ", 에러, sql.String())
	}

	defer F테이블_삭제(테스트용_테이블)
	
	일자, _ := F문자열2일자("2000-01-01")

	sql = new(bytes.Buffer)
	sql.WriteString("INSERT INTO ")
	sql.WriteString(테스트용_테이블 + " (")
	sql.WriteString("uint_column,")
	sql.WriteString("int_column,")
	sql.WriteString("float_column,")
	sql.WriteString("string_column,")
	sql.WriteString("date_column")
	sql.WriteString(") VALUES (")
	sql.WriteString("?,?,?,?,?")
	sql.WriteString(")")
	_, 에러 = F_SQL실행(sql.String(), 1, -1, 1.1, "one", 일자)
	if 에러 != nil {
		테스트.Error("common.TestF_SQL질의_정수() : INSERT문 에러.", 에러, sql.String())
	}
	
	// 부호없는 정수 테스트
	var 부호없는_정수 uint = 0
	질의값, 에러 := F_SQL질의(부호없는_정수, "SELECT uint_column FROM " + 테스트용_테이블)
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : 부호없는 정수 칼럼 질의 에러 : ", 에러, sql.String())
	}
	
	부호없는_정수_결과값, 에러 := 질의값.G부호없는_정수()
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : S가변형.G부호없는_정수() 에러.", 에러)
	}
	if 부호없는_정수_결과값 != 1 {
		테스트.Errorf("TestF_SQL질의() : 부호없는_정수 불일치. 예상값 1, 실제값 %v", 부호없는_정수_결과값)
	}
	
	// 부호없는_정수8 테스트
	var 부호없는_정수8 uint8 = 0
	질의값, 에러 = F_SQL질의(부호없는_정수8, "SELECT uint_column FROM " + 테스트용_테이블)
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : 부호없는 정수 칼럼 질의 에러 : ", 에러, sql.String())
	}
	
	부호없는_정수8_결과값, 에러 := 질의값.G부호없는_정수8()
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : S가변형.G부호없는_정수8() 에러.", 에러)
	}
	if 부호없는_정수8_결과값 != 1 {
		테스트.Errorf("TestF_SQL질의() : 부호없는_정수8 불일치. 예상값 1, 실제값 %v", 부호없는_정수8_결과값)
	}
	
	// 부호없는_정수16 테스트
	var 부호없는_정수16 uint16 = 0
	질의값, 에러 = F_SQL질의(부호없는_정수16, "SELECT uint_column FROM " + 테스트용_테이블)
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : 부호없는 정수 칼럼 질의 에러 : ", 에러, sql.String())
	}
	
	부호없는_정수16_결과값, 에러 := 질의값.G부호없는_정수16()
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : S가변형.G부호없는_정수16() 에러.", 에러)
	}
	if 부호없는_정수16_결과값 != 1 {
		테스트.Errorf("TestF_SQL질의() : 부호없는_정수16 불일치. 예상값 1, 실제값 %v", 부호없는_정수16_결과값)
	}
	
	// 부호없는_정수32 테스트
	var 부호없는_정수32 uint32 = 0
	질의값, 에러 = F_SQL질의(부호없는_정수32, "SELECT uint_column FROM " + 테스트용_테이블)
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : 부호없는 정수 칼럼 질의 에러 : ", 에러, sql.String())
	}	
	부호없는_정수32_결과값, 에러 := 질의값.G부호없는_정수32()
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : S가변형.G부호없는_정수32() 에러.", 에러)
	}
	if 부호없는_정수32_결과값 != 1 {
		테스트.Errorf("TestF_SQL질의() : 부호없는_정수32 불일치. 예상값 1, 실제값 %v", 부호없는_정수32_결과값)
	}
	
	// 부호없는_정수64 테스트
	var 부호없는_정수64 uint64 = 0
	질의값, 에러 = F_SQL질의(부호없는_정수64, "SELECT uint_column FROM " + 테스트용_테이블)
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : 부호없는 정수 칼럼 질의 에러 : ", 에러, sql.String())
	}
	
	부호없는_정수64_결과값, 에러 := 질의값.G부호없는_정수64()
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : S가변형.G부호없는_정수64() 에러.", 에러)
	}
	if 부호없는_정수64_결과값 != 1 {
		테스트.Errorf("TestF_SQL질의() : 부호없는_정수64 불일치. 예상값 1, 실제값 %v", 부호없는_정수64_결과값)
	}
	
	// 정수 테스트
	var 정수 int = 0
	질의값, 에러 = F_SQL질의(정수, "SELECT int_column FROM " + 테스트용_테이블)
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : 정수 칼럼 질의 에러 : ", 에러, sql.String())
	}

	정수_결과값, 에러 := 질의값.G정수()
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : S가변형.G정수() 에러.", 에러)
	}
	if 정수_결과값 != -1 {
		테스트.Errorf("TestF_SQL질의() : 정수 불일치. 예상값 -1, 실제값 %v", 정수_결과값)
	}
	
	// 정수8 테스트
	var 정수8 int8 = 0
	질의값, 에러 = F_SQL질의(정수8, "SELECT int_column FROM " + 테스트용_테이블)
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : 정수 칼럼 질의 에러 : ", 에러, sql.String())
	}

	정수8_결과값, 에러 := 질의값.G정수8()
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : S가변형.G정수8() 에러.", 에러)
	}
	if 정수8_결과값 != -1 {
		테스트.Errorf("TestF_SQL질의() : 정수8 불일치. 예상값 -1, 실제값 %v", 정수8_결과값)
	}
	
	// 정수16 테스트
	var 정수16 int16 = 0
	질의값, 에러 = F_SQL질의(정수16, "SELECT int_column FROM " + 테스트용_테이블)
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : 정수 칼럼 질의 에러 : ", 에러, sql.String())
	}

	정수16_결과값, 에러 := 질의값.G정수16()
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : S가변형.G정수16() 에러.", 에러)
	}
	if 정수16_결과값 != -1 {
		테스트.Errorf("TestF_SQL질의() : 정수16 불일치. 예상값 -1, 실제값 %v", 정수16_결과값)
	}
	
	// 정수32 테스트
	var 정수32 int32 = 0
	질의값, 에러 = F_SQL질의(정수32, "SELECT int_column FROM " + 테스트용_테이블)
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : 정수 칼럼 질의 에러 : ", 에러, sql.String())
	}

	정수32_결과값, 에러 := 질의값.G정수32()
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : S가변형.G정수32() 에러.", 에러)
	}
	if 정수32_결과값 != -1 {
		테스트.Errorf("TestF_SQL질의() : 정수32 불일치. 예상값 -1, 실제값 %v", 정수32_결과값)
	}
	
	// 정수64 테스트
	var 정수64 int64 = 0
	질의값, 에러 = F_SQL질의(정수64, "SELECT int_column FROM " + 테스트용_테이블)
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : 정수 칼럼 질의 에러 : ", 에러, sql.String())
	}

	정수64_결과값, 에러 := 질의값.G정수64()
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : S가변형.G정수64() 에러.", 에러)
	}
	if 정수64_결과값 != -1 {
		테스트.Errorf("TestF_SQL질의() : 정수64 불일치. 예상값 -1, 실제값 %v", 정수64_결과값)
	}
	
	// 실수32 테스트
	var 실수32 float32 = 0.0
	질의값, 에러 = F_SQL질의(실수32, "SELECT float_column FROM " + 테스트용_테이블 + " WHERE int_column = ?", -1)
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : 실수 칼럼 질의 에러 : ", 에러, sql.String())
	}

	실수32_결과값, 에러 := 질의값.G실수32()
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : S가변형.G실수32() 에러.", 에러)
	}
	if F반올림_통화(float64(실수32_결과값)) != 1.1 {
		테스트.Errorf("TestF_SQL질의() : 실수32 불일치. 예상값 1.1, 실제값 %v", 실수32_결과값)
	}

	// 실수64 테스트
	var 실수64 float64 = 0.0
	질의값, 에러 = F_SQL질의(실수64, "SELECT float_column FROM " + 테스트용_테이블 + " WHERE int_column = ?", -1)
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : 실수 칼럼 질의 에러 : ", 에러, sql.String())
	}

	실수64_결과값, 에러 := 질의값.G실수64()
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : S가변형.G실수64() 에러.", 에러)
	}
	if F반올림_통화(실수64_결과값) != 1.1 {
		테스트.Errorf("TestF_SQL질의() : 실수64 불일치. 예상값 1.1, 실제값 %v", 실수64_결과값)
	}

	// 문자열 테스트
	var 문자열 string = ""
	질의값, 에러 = F_SQL질의(문자열, "SELECT string_column FROM " + 테스트용_테이블)
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : 문자열 칼럼 질의 에러 : ", 에러, sql.String())
	}

	문자열_결과값, 에러 := 질의값.G문자열()
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : S가변형.G문자열() 에러.", 에러)
	}
	if 문자열_결과값 != "one" {
		테스트.Errorf("TestF_SQL질의() : 문자열값 불일치. 예상값 'one', 실제값 %v", 문자열_결과값)
	}
	
	// 시점 테스트
	var 시점 time.Time
	질의값, 에러 = F_SQL질의(시점, "SELECT date_column FROM " + 테스트용_테이블)
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : 시점 칼럼 질의 에러 : ", 에러, sql.String())
	}

	시점_결과값, 에러 := 질의값.G시점()
	if 에러 != nil {
		테스트.Error("TestF_SQL질의() : S가변형.G시점() 에러.", 에러)
	}
	if !시점_결과값.Equal(일자) {
		테스트.Errorf("TestF_SQL질의() : 문자열값 불일치. 예상값 %v, 실제값 %v", F일자2문자열(일자), F일자2문자열(시점_결과값))
	}
}