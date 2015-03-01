package common

import (
	"bytes"
	"log"
)

type C종목 struct {
	식별코드   *C부호없는_정수64
	종목코드   *C문자열 // 단축코드
	종목코드2  *C문자열 // 표준코드
	종목명칭   *C문자열 // 해당국가 명칭
	종목명칭2  *C문자열 // 영문명칭
	발행기관코드 *C문자열
	시장구분   *C문자열 // 코스피, 코스닥, 상장유지, 상장폐지
}

func (c *C종목) G식별코드() uint64   { return c.식별코드.G값() }
func (c *C종목) G종목코드() string   { return c.종목코드.G값() }
func (c *C종목) G종목코드2() string  { return c.종목코드2.G값() }
func (c *C종목) G종목명칭() string   { return c.종목명칭.G값() }
func (c *C종목) G종목명칭2() string  { return c.종목명칭2.G값() }
func (c *C종목) G발행기관코드() string { return c.발행기관코드.G값() }
func (c *C종목) G시장구분() string   { return c.시장구분.G값() }
func F종목_생성(
	식별코드 uint64,
	종목코드 string,
	종목코드2 string,
	종목명칭 string,
	종목명칭2 string,
	발행기관코드 string,
	시장구분 string) *C종목 {
	c := new(C종목)
	c.식별코드 = F부호없는_정수64_생성(식별코드)
	c.종목코드 = F문자열_생성(종목코드)
	c.종목코드2 = F문자열_생성(종목코드2)
	c.종목명칭 = F문자열_생성(종목명칭)
	c.종목명칭2 = F문자열_생성(종목명칭2)
	c.발행기관코드 = F문자열_생성(발행기관코드)
	c.시장구분 = F문자열_생성(시장구분)

	return c
}

var 종목정보_저장소 map[string]*C종목 = nil

func F종목_검색(종목코드 string) *C종목 {
	var 에러 error

	if 종목정보_저장소 == nil {
		종목정보_저장소, 에러 = F종목정보_맵()

		if 에러 != nil {
			log.Println("common.F종목검색() : F종목정보_맵() 에러 발생.")
		}

		if 종목정보_저장소 == nil {
			log.Println("common.F종목검색() : F종목정보_맵() 실행 후에도 nil.")
		}

		if len(종목정보_저장소) == 0 {
			log.Println("common.F종목검색() : F종목정보_맵() 실행 후 맵이 비어있음.")
		}
	}

	return 종목정보_저장소[종목코드]
}

func F종목정보_맵_초기화() {
	종목정보_저장소 = nil
	F종목정보_맵()
}

func F종목정보_맵() (map[string]*C종목, error) {
	// 복사본을 만들어서 넘겨줘서 중앙 저장소가 변경되는 것을 예방한다.
	종목정보_맵 := make(map[string]*C종목)

	if 종목정보_저장소 == nil {
		종목정보_저장소 = make(map[string]*C종목)
	}

	if len(종목정보_저장소) > 0 {
		for 키, 종목정보 := range 종목정보_저장소 {
			종목정보_맵[키] = 종목정보
		}

		return 종목정보_맵, nil
	}

	데이터베이스, 에러 := F데이터베이스_연결()
	if 에러 != nil {
		log.Print("F종목정보_맵() DB열기 에러 : ", 에러)
		종목정보_저장소 = nil

		return nil, 에러
	}
	defer 데이터베이스.Close()

	sql := new(bytes.Buffer)
	sql.WriteString("SELECT")
	sql.WriteString("	id,")
	sql.WriteString("	issuer,")
	sql.WriteString("	name1,")
	sql.WriteString("	name2,")
	sql.WriteString("	code1,")
	sql.WriteString("	code2,")
	sql.WriteString("	market_status ")
	sql.WriteString("FROM " + F종목정보_테이블())

	행_모음, 에러 := 데이터베이스.Query(sql.String())
	if 에러 != nil {
		log.Print("F종목정보_맵() 데이터베이스.Query() 에러 : ", 에러)
		종목정보_저장소 = nil

		return nil, 에러
	}
	defer 행_모음.Close()

	var 식별코드 uint64
	var 종목코드, 종목코드2 string
	var 종목명칭, 종목명칭2 string
	var 발행기관코드, 시장구분 string

	for 행_모음.Next() {
		에러 := 행_모음.Scan(
			&식별코드,
			&발행기관코드,
			&종목명칭,
			&종목명칭2,
			&종목코드,
			&종목코드2,
			&시장구분)

		if 에러 != nil {
			log.Print("common.F종목정보_맵() : 행_모음.Scan() 에러 : ", 에러)
			종목정보_저장소 = nil

			return nil, 에러
		}

		종목정보_저장소[종목코드] = F종목_생성(식별코드,
			종목코드,
			종목코드2,
			종목명칭,
			종목명칭2,
			발행기관코드,
			시장구분)
	}

	for 키, 종목정보 := range 종목정보_저장소 {
		종목정보_맵[키] = 종목정보
	}

	return 종목정보_맵, nil
}

func F종목정보_맵_DB기록(종목정보_맵 map[string]*C종목) error {
	sql := new(bytes.Buffer)
	sql.WriteString("INSERT INTO ")
	sql.WriteString(F종목정보_테이블() + " ")
	sql.WriteString("(")
	sql.WriteString("issuer,")
	sql.WriteString("name1,")
	sql.WriteString("name2,")
	sql.WriteString("code1,")
	sql.WriteString("code2,")
	sql.WriteString("market_status")
	sql.WriteString(") VALUES (")
	sql.WriteString("?,?,?,?,?,?")
	sql.WriteString(")")
	sql추가 := sql.String()

	sql = new(bytes.Buffer)
	sql.WriteString("UPDATE ")
	sql.WriteString(F종목정보_테이블() + " ")
	sql.WriteString("SET")
	sql.WriteString("	issuer = ?,")
	sql.WriteString("	name1 = ?,")
	sql.WriteString("	name2 = ?,")
	sql.WriteString("	code1 = ?,")
	sql.WriteString("	code2 = ?,")
	sql.WriteString("	market_status = ? ")
	sql.WriteString("WHERE")
	sql.WriteString("	id = ?")
	sql갱신 := sql.String()

	데이터베이스, 에러 := F데이터베이스_연결()
	if 에러 != nil {
		log.Println("common.F종목정보_맵_DB기록() : F데이터베이스_연결() 에러.", 에러)
		return 에러
	}
	defer 데이터베이스.Close()

	F종목정보_맵_초기화()
	기존_종목정보_맵, 에러 := F종목정보_맵()
	if 에러 != nil {
		log.Println("common.F종목정보_맵_DB기록() : F종목정보_맵() 에러.", 에러)
		return 에러
	}

	트랜잭션, 에러 := 데이터베이스.Begin()
	if 에러 != nil {
		log.Println("common.F종목정보_맵_DB기록() : 트랜잭션 초기화 에러.", 에러)
		return 에러
	}

	stmt추가, 에러 := 트랜잭션.Prepare(sql추가)
	if 에러 != nil {
		log.Println("common.F종목정보_맵_DB기록() : stmt추가 준비 에러.", 에러)
		return 에러
	}
	defer stmt추가.Close()

	stmt갱신, 에러 := 트랜잭션.Prepare(sql갱신)
	if 에러 != nil {
		log.Println("common.F종목정보_맵_DB기록() : stmt갱신 초기화 에러.", 에러)
		return 에러
	}
	defer stmt갱신.Close()

	// 이미 존재하는 종목은 업데이트, 새 종목은 추가.
	for _, 종목 := range 종목정보_맵 {
		기존_종목, 이미_존재함 := 기존_종목정보_맵[종목.G종목코드()]

		if 이미_존재함 {
			// 모든 정보가 바뀐 게 없으면 굳이 새로 쓸 필요없으므로 건너뜀.
			if 종목.G발행기관코드() == 기존_종목.G발행기관코드() &&
				종목.G종목명칭() == 기존_종목.G종목명칭() &&
				종목.G종목명칭2() == 기존_종목.G종목명칭2() &&
				종목.G종목코드() == 기존_종목.G종목코드() &&
				종목.G종목코드2() == 기존_종목.G종목코드2() &&
				종목.G시장구분() == 기존_종목.G시장구분() {
				continue
			}

			_, 에러 = stmt갱신.Exec(
				종목.G발행기관코드(),
				종목.G종목명칭(),
				종목.G종목명칭2(),
				종목.G종목코드(),  // 실제로는 표준코드보다 단축코드를 더 많이 쓴다.
				종목.G종목코드2(), // 심지어 단축코드의 첫글자 "A"도 생략하곤 한다.
				종목.G시장구분(),  // 그러므로, 종목코드는 단축코드를 기준으로 한다.
				기존_종목.G식별코드())

			if 에러 != nil {
				log.Println("F종목정보_맵_DB기록() : stmt갱신.Exec() 에러. "+
					종목.G종목코드()+", "+종목.G종목명칭(), 에러)

				롤백_에러 := 트랜잭션.Rollback()
				if 롤백_에러 != nil {
					log.Println("common.F종목정보_맵_DB기록() : 롤백 에러가 발생하였습니다.", 롤백_에러)
				}

				return 에러
			}
		} else {
			_, 에러 = stmt추가.Exec(
				종목.G발행기관코드(),
				종목.G종목명칭(),
				종목.G종목명칭2(),
				종목.G종목코드(),  // 실제로는 표준코드보다 단축코드를 더 많이 쓴다.
				종목.G종목코드2(), // 심지어 단축코드의 첫글자 "A"도 생략하곤 한다.
				종목.G시장구분())  // 그러므로, 종목코드는 단축코드를 기준으로 한다.

			if 에러 != nil {
				log.Println("common.F종목정보_맵_DB기록() : stmt추가.Exec() 에러. "+
					종목.G종목코드()+", "+종목.G종목명칭(), 에러)

				롤백_에러 := 트랜잭션.Rollback()
				if 롤백_에러 != nil {
					log.Println("common.F종목정보_맵_DB기록() : 롤백 에러가 발생하였습니다.", 롤백_에러)
				}

				return 에러
			}
		}
	}

	에러 = 트랜잭션.Commit()
	if 에러 != nil {
		log.Println("common.F종목정보_맵_DB기록() : 커밋 에러가 발생하였습니다.", 에러)

		return 에러
	}

	F종목정보_맵_초기화()

	return nil
}
