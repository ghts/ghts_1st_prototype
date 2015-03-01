package common

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
	"strings"
	"time"
)

func F데이터베이스_연결() (*sql.DB, error) {
	연결정보 := F_MySQL연결정보()

	데이터베이스, 에러 := sql.Open(연결정보.G드라이버(), 연결정보.G연결문자열())

	if 에러 != nil {
		log.Println("common.F데이터베이스_연결() : sql.Open() 에러")
		log.Println(에러)

		return nil, 에러
	}

	// sql.Open이 연결을 열은 것이 아니다. Ping으로 연결을 확인한다.
	// DB 서버 초기화 될 때, 시간초과로 연결에러가 발생하므로 몇 번 재시도 해 본다.
	var Ping성공 = false
	for 반복횟수 := 0; 반복횟수 < 20; 반복횟수++ {
		에러 = 데이터베이스.Ping()

		if 에러 == nil {
			Ping성공 = true
			break
		}

		time.Sleep(1 * time.Second)
	}

	if Ping성공 == false {
		log.Println("common.F데이터베이스_연결() : 데이터베이스.Ping() 에러")
		log.Println(에러)

		return nil, 에러
	}

	return 데이터베이스, nil
}

func F_SQL실행(SQL문 string, 파라메터 ...interface{}) (sql.Result, error) {
	데이터베이스, 에러 := F데이터베이스_연결()
	if 에러 != nil {
		log.Println("common.F_SQL실행() 데이터베이스 열기 에러 : ", SQL문, 파라메터, 에러)
		return nil, 에러
	}
	defer 데이터베이스.Close()

	트랜잭션, 에러 := 데이터베이스.Begin()
	if 에러 != nil {
		log.Println("common.F_SQL실행() 트랜잭션 시작에러 : ", SQL문, 파라메터, 에러)
		return nil, 에러
	}

	DB명령, 에러 := 트랜잭션.Prepare(SQL문)
	if 에러 != nil {
		log.Println("common.F_SQL실행() DB명령 준비 에러 : ", SQL문, 파라메터, 에러)
		return nil, 에러
	}
	defer DB명령.Close()

	var 결과 sql.Result
	if 파라메터 == nil || len(파라메터) == 0 {
		결과, 에러 = DB명령.Exec()
	} else {
		파라메터_모음 := make([]interface{}, 0)
		for _, 개별_파라메터 := range 파라메터 {
			파라메터_모음 = append(파라메터_모음, 개별_파라메터)
		}

		결과, 에러 = DB명령.Exec(파라메터_모음...)
	}

	if 에러 != nil {
		SQL문 = strings.ToUpper(SQL문)
		SQL문 = strings.Replace(SQL문, " ", "", -1)
		SQL문 = strings.Replace(SQL문, "\t", "", -1)

		if strings.Contains(SQL문, "DROPTABLE") {
			// 테이블이 존재하지 않는 DROP TABLE 에러는
			// 유닛 테스트 할 때 테스트용 임시테이블을 준비하는 과정에서
			// 생기는 것이므로 무시하자.
			// 실제 프로그램 운용에서는 테이블을 삭제할 일이 없다.
		} else {
			log.Println("common.F_SQL실행() DB명령.Exec() 에러 : ", SQL문, 파라메터, 에러)
		}

		트랜잭션.Rollback()
		return nil, 에러
	}

	트랜잭션.Commit()

	return 결과, nil
}

func F_SQL질의_부호없는_정수(SQL문 string, 파라메터 ...interface{}) (uint64, error) {
	var 부호없는_정수64 uint64
	질의값, 에러 := F_SQL질의(부호없는_정수64, SQL문, 파라메터)
	if 에러 != nil {
		log.Println("common.F_SQL질의_부호없는_정수() : common.F_SQL질의() 에러.", 에러)
	}

	부호없는_정수64_결과값, 에러 := 질의값.G부호없는_정수64()
	if 에러 != nil {
		log.Println("common.F_SQL질의_부호없는_정수() : common.S가변형.G부호없는_정수64() 에러.", 에러)
	}

	return 부호없는_정수64_결과값, nil
}

func F_SQL질의_정수(SQL문 string, 파라메터 ...interface{}) (int64, error) {
	var 정수64 int64
	질의값, 에러 := F_SQL질의(정수64, SQL문, 파라메터)
	if 에러 != nil {
		log.Println("common.F_SQL질의_정수() : common.F_SQL질의() 에러.", 에러)
	}

	정수64_결과값, 에러 := 질의값.G정수64()
	if 에러 != nil {
		log.Println("common.F_SQL질의_정수() : common.S가변형.G정수64() 에러.", 에러)
	}

	return 정수64_결과값, nil
}

func F_SQL질의_실수(SQL문 string, 파라메터 ...interface{}) (float64, error) {
	var 실수64 float64
	질의값, 에러 := F_SQL질의(실수64, SQL문, 파라메터)
	if 에러 != nil {
		log.Println("common.F_SQL질의_실수() : common.F_SQL질의() 에러.", 에러)
	}

	실수64_결과값, 에러 := 질의값.G실수64()
	if 에러 != nil {
		log.Println("common.F_SQL질의_실수() : common.S가변형.G실수64() 에러.", 에러)
	}

	return 실수64_결과값, nil
}

func F_SQL질의_문자열(SQL문 string, 파라메터 ...interface{}) (string, error) {
	var 문자열 string
	질의값, 에러 := F_SQL질의(문자열, SQL문, 파라메터)
	if 에러 != nil {
		log.Println("common.F_SQL질의_문자열() : common.F_SQL질의() 에러.", 에러)
	}

	문자열_결과값, 에러 := 질의값.G문자열()
	if 에러 != nil {
		log.Println("common.F_SQL질의_문자열() : common.S가변형.G문자열() 에러.", 에러)
	}

	return 문자열_결과값, nil
}

func F_SQL질의_시점(SQL문 string, 파라메터 ...interface{}) (time.Time, error) {
	var 시점 time.Time
	질의값, 에러 := F_SQL질의(시점, SQL문, 파라메터)
	if 에러 != nil {
		log.Println("common.F_SQL질의_시점() : common.F_SQL질의() 에러.", 에러)
	}

	시점_결과값, 에러 := 질의값.G시점()
	if 에러 != nil {
		log.Println("common.F_SQL질의_시점() : common.S가변형.G시점() 에러.", 에러)
	}

	return 시점_결과값, nil
}

func F_SQL질의(반환값 interface{}, 
				SQL문 string, 
				파라메터 ...interface{}) (*S가변형, error) {		
	데이터베이스, 에러 := F데이터베이스_연결()
	if 에러 != nil {
		log.Println("common.F_SQL질의() : F데이터베이스_연결() 에러")

		return nil, 에러
	}
	defer 데이터베이스.Close()

	var 행 *sql.Row
	파라메터_모음 := F중첩된_외부_슬라이스_제거(파라메터)
	
	if len(파라메터_모음) == 0 {
		행 = 데이터베이스.QueryRow(SQL문)		
	} else {
		행 = 데이터베이스.QueryRow(SQL문, 파라메터_모음...)
	}
	
	var 가변형_변수 *S가변형 
	
	switch 반환값.(type) {
	case bool:
		var 참거짓 bool		
		에러 = 행.Scan(&참거짓)
		if 에러 == nil {
			가변형_변수 = F가변형_생성(참거짓)
		}
	case uint:
		var 부호없는_정수 uint
		에러 = 행.Scan(&부호없는_정수)
		if 에러 == nil {
			가변형_변수 = F가변형_생성(부호없는_정수)
		}
	case uint8:
		var 부호없는_정수8 uint8
		에러 = 행.Scan(&부호없는_정수8)
		if 에러 == nil {
			가변형_변수 = F가변형_생성(부호없는_정수8)
		}
	case uint16:
		var 부호없는_정수16 uint16
		에러 = 행.Scan(&부호없는_정수16)
		if 에러 == nil {
			가변형_변수 = F가변형_생성(부호없는_정수16)
		}
	case uint32:
		var 부호없는_정수32 uint32
		에러 = 행.Scan(&부호없는_정수32)
		if 에러 == nil {
			가변형_변수 = F가변형_생성(부호없는_정수32)
		}		
	case uint64:
		var 부호없는_정수64 uint64
		에러 = 행.Scan(&부호없는_정수64)
		if 에러 == nil {
			가변형_변수 = F가변형_생성(부호없는_정수64)
		}
	case int:
		var 정수 int
		에러 = 행.Scan(&정수)
		if 에러 == nil {
			가변형_변수 = F가변형_생성(정수)
		}
	case int8:
		var 정수8 int8
		에러 = 행.Scan(&정수8)
		if 에러 == nil {
			가변형_변수 = F가변형_생성(정수8)
		}
	case int16:
		var 정수16 int16
		에러 = 행.Scan(&정수16)
		if 에러 == nil {
			가변형_변수 = F가변형_생성(정수16)
		}
	case int32:
		var 정수32 int32
		에러 = 행.Scan(&정수32)
		if 에러 == nil {
			가변형_변수 = F가변형_생성(정수32)
		}
	case int64:
		var 정수64 int64
		에러 = 행.Scan(&정수64)
		if 에러 == nil {
			가변형_변수 = F가변형_생성(정수64)
		}
	case float32:
		var 실수32 float32
		에러 = 행.Scan(&실수32)
		if 에러 == nil {
			가변형_변수 = F가변형_생성(실수32)
		}
	case float64:
		var 실수64 float64
		에러 = 행.Scan(&실수64)
		if 에러 == nil {
			가변형_변수 = F가변형_생성(실수64)
		}
	case string:
		var 문자열 string
		에러 = 행.Scan(&문자열)
		if 에러 == nil {
			가변형_변수 = F가변형_생성(문자열)
		}
	case time.Time:
		var 시점 time.Time
		에러 = 행.Scan(&시점)
		if 에러 == nil {
			가변형_변수 = F가변형_생성(시점)
		}
	default:
		log.Println("common.F_SQL질의() 디버깅 : 예상하지 못한 반환값 형식.", reflect.TypeOf(반환값).String())

		return nil, errors.New("common.F_SQL질의() 디버깅 : 예상하지 못한 반환값 형식. " + reflect.TypeOf(반환값).String())
	}
	
	if 에러 == sql.ErrNoRows {
		log.Println("common.F_SQL질의() : 결과값이 존재하지 않음.")
		log.Println(에러)
		log.Println(SQL문)
		log.Println(파라메터)

		return nil, 에러
	} else if 에러 != nil {
		log.Println("common.F_SQL질의() : 질의 에러.")
		log.Println(에러)
		log.Printf("SQL : %v", SQL문)
		log.Printf("파라메터 : %v", 파라메터)

		return nil, 에러
	}
	
	return 가변형_변수, nil
}