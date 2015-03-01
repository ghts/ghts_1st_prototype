package common

import (
	"bytes"
	"testing"
)

func TestF종목_생성(테스트 *testing.T) {
	식별코드 := uint64(3)
	종목코드 := "A000020"
	종목코드2 := "KR7000020008"
	종목명칭 := "동화약품보통주"
	종목명칭2 := "DongwhaPharm"
	발행기관코드 := "2"
	시장구분 := "유가증권시장상장"

	종목 := F종목_생성(식별코드,
		종목코드,
		종목코드2,
		종목명칭,
		종목명칭2,
		발행기관코드,
		시장구분)

	식별코드1 := 종목.G식별코드()
	if 식별코드1 != uint64(3) ||
		식별코드1 != 식별코드 ||
		식별코드1 != 종목.식별코드.G값() {
		테스트.Error("F종목_생성().G식별코드() 에러 1")
	}
	식별코드1 = uint64(999)

	식별코드2 := 종목.G식별코드()
	if 식별코드2 != uint64(3) ||
		식별코드2 != 식별코드 ||
		식별코드2 != 종목.식별코드.G값() {
		테스트.Error("F종목_생성().G식별코드() 에러 2")
	}

	종목코드1_1 := 종목.G종목코드()
	if 종목코드1_1 != "A000020" ||
		종목코드1_1 != 종목코드 ||
		종목코드1_1 != 종목.종목코드.G값() {
		테스트.Error("F종목_생성().G종목코드() 에러 1")
	}
	종목코드1_1 = "변경된 종목코드"

	종목코드1_2 := 종목.G종목코드()
	if 종목코드1_2 != "A000020" ||
		종목코드1_2 != 종목코드 ||
		종목코드1_2 != 종목.종목코드.G값() {
		테스트.Error("F종목_생성().G종목코드() 에러 2")
	}

	종목코드2_1 := 종목.G종목코드2()
	if 종목코드2_1 != "KR7000020008" ||
		종목코드2_1 != 종목코드2 ||
		종목코드2_1 != 종목.종목코드2.G값() {
		테스트.Error("F종목_생성().G종목코드2() 에러 1")
	}
	종목코드2_1 = "변경된 종목코드2"

	종목코드2_2 := 종목.G종목코드2()
	if 종목코드2_2 != "KR7000020008" ||
		종목코드2_2 != 종목코드2 ||
		종목코드2_2 != 종목.종목코드2.G값() {
		테스트.Error("F종목_생성().G종목코드2() 에러 2")
	}

	종목명칭1_1 := 종목.G종목명칭()
	if 종목명칭1_1 != "동화약품보통주" ||
		종목명칭1_1 != 종목명칭 ||
		종목명칭1_1 != 종목.종목명칭.G값() {
		테스트.Error("F종목_생성().G종목명칭() 에러 1")
	}
	종목명칭1_1 = "변경된 종목명칭"

	종목명칭1_2 := 종목.G종목명칭()
	if 종목명칭1_2 != "동화약품보통주" ||
		종목명칭1_2 != 종목명칭 ||
		종목명칭1_2 != 종목.종목명칭.G값() {
		테스트.Error("F종목_생성().G종목명칭() 에러 2")
	}

	종목명칭2_1 := 종목.G종목명칭2()
	if 종목명칭2_1 != "DongwhaPharm" ||
		종목명칭2_1 != 종목명칭2 ||
		종목명칭2_1 != 종목.종목명칭2.G값() {
		테스트.Error("F종목_생성().G종목명칭2() 에러 1")
	}
	종목명칭2_1 = "변경된 종목명칭2"

	종목명칭2_2 := 종목.G종목명칭2()
	if 종목명칭2_2 != "DongwhaPharm" ||
		종목명칭2_2 != 종목명칭2 ||
		종목명칭2_2 != 종목.종목명칭2.G값() {
		테스트.Error("F종목_생성().G종목명칭2() 에러 2")
	}

	발행기관코드1 := 종목.G발행기관코드()
	if 발행기관코드1 != "2" ||
		발행기관코드1 != 발행기관코드 ||
		발행기관코드1 != 종목.발행기관코드.G값() {
		테스트.Error("F종목_생성().G발행기관코드() 에러 1")
	}
	발행기관코드1 = "변경된 발행기관코드"

	발행기관코드2 := 종목.G발행기관코드()
	if 발행기관코드2 != "2" ||
		발행기관코드2 != 발행기관코드 ||
		발행기관코드2 != 종목.발행기관코드.G값() {
		테스트.Error("F종목_생성().G발행기관코드() 에러 2")
	}

	시장구분1 := 종목.G시장구분()
	if 시장구분1 != 시장구분 {
		테스트.Error("F종목_생성().G시장구분() 에러 1")
	}
	시장구분1 = "변경된 시장구분"

	시장구분2 := 종목.G시장구분()
	if 시장구분2 != 시장구분 {
		테스트.Error("F종목_생성().G시장구분() 에러 2")
	}
}

func TestF종목정보_맵(테스트 *testing.T) {
	F테스트_모드_시작()
	defer F테스트_모드_종료()
	F종목정보_테이블_생성()
	defer F테이블_삭제(F종목정보_테이블())

	F종목정보_맵_초기화()

	종목 := F종목_동화약품()

	sql := new(bytes.Buffer)
	sql.WriteString("INSERT INTO ")
	sql.WriteString(F종목정보_테이블() + " (")
	sql.WriteString("id,")
	sql.WriteString("issuer,")
	sql.WriteString("name1,")
	sql.WriteString("name2,")
	sql.WriteString("code1,")
	sql.WriteString("code2,")
	sql.WriteString("market_status")
	sql.WriteString(") VALUES (")
	sql.WriteString("?,?,?,?,?,?,?")
	sql.WriteString(")")

	_, 에러 := F_SQL실행(sql.String(),
		종목.G식별코드(),
		종목.G발행기관코드(),
		종목.G종목명칭(),
		종목.G종목명칭2(),
		종목.G종목코드(),
		종목.G종목코드2(),
		종목.G시장구분())

	if 에러 != nil {
		테스트.Error("TestF종목정보_맵() INSERT 에러 1.", 에러, sql.String())
	}

	종목 = F종목_삼성전자()

	_, 에러 = F_SQL실행(sql.String(),
		종목.G식별코드(),
		종목.G발행기관코드(),
		종목.G종목명칭(),
		종목.G종목명칭2(),
		종목.G종목코드(),
		종목.G종목코드2(),
		종목.G시장구분())

	if 에러 != nil {
		테스트.Error("TestF종목정보_맵() INSERT 에러 2", 에러, sql.String())
	}

	F종목정보_맵_초기화()
	종목정보_맵, 에러 := F종목정보_맵()

	if 에러 != nil {
		테스트.Error("TestF종목정보_맵() : F종목정보_맵()에서 에러가 발생하였습니다.", 에러)
	}

	if 종목정보_맵 == nil {
		테스트.Error("TestF종목정보_맵() : 종목정보_맵이 nil입니다.")
	}

	if len(종목정보_맵) != 2 {
		테스트.Error("TestF종목정보_맵() 항목 수량이 예상과 다릅니다. 예상값 1개. 실제값 :", len(종목정보_맵))
	}

	종목 = F종목_동화약품()
	종목1 := 종목정보_맵[종목.G종목코드()]

	종목 = F종목_삼성전자()
	종목2 := 종목정보_맵[종목.G종목코드()]

	if 종목1 == nil {
		테스트.Error("TestF종목정보_맵() : 종목1가 nil입니다.")
	}

	if 종목2 == nil {
		테스트.Error("TestF종목정보_맵() : 종목2가 nil입니다.")
	}

	종목 = F종목_동화약품()
	if 종목.G종목코드() != 종목1.G종목코드() ||
		종목.G종목코드2() != 종목1.G종목코드2() ||
		종목.G종목명칭() != 종목1.G종목명칭() ||
		종목.G종목명칭2() != 종목1.G종목명칭2() ||
		종목.G발행기관코드() != 종목1.G발행기관코드() ||
		종목.G시장구분() != 종목1.G시장구분() {
		테스트.Error("TestF종목정보_맵(): 종목1 내용 불일치.")
	}

	종목 = F종목_삼성전자()
	if 종목.G종목코드() != 종목2.G종목코드() ||
		종목.G종목코드2() != 종목2.G종목코드2() ||
		종목.G종목명칭() != 종목2.G종목명칭() ||
		종목.G종목명칭2() != 종목2.G종목명칭2() ||
		종목.G발행기관코드() != 종목2.G발행기관코드() ||
		종목.G시장구분() != 종목2.G시장구분() {
		테스트.Error("TestF종목정보_맵(): 종목2 내용 불일치.")
	}

	// 중앙저장소 변경 불가능 확인

	// 반환값을 외부에서 변경함.
	종목정보_맵[F종목_동화약품().G종목코드()] = F종목_삼성전자()

	종목_원본 := 종목정보_저장소[F종목_동화약품().G종목코드()]
	종목_복사본 := 종목정보_맵[F종목_동화약품().G종목코드()]

	if 종목_원본.G종목코드() == 종목_복사본.G종목코드() {
		테스트.Error("0301 common.TestF종목정보_맵() : 종목정보 저장소 원본이 변경되었습니다.")
	}
}

func TestF종목_검색(테스트 *testing.T) {
	F테스트_모드_시작()
	defer F테스트_모드_종료()
	F종목정보_테이블_생성()
	defer F테이블_삭제(F종목정보_테이블())

	종목정보_맵_원본 := F종목정보_맵_테스트용()

	에러 := F종목정보_맵_DB기록(종목정보_맵_원본)
	if 에러 != nil {
		테스트.Error("TestF종목검색() : F종목정보_맵_DB기록() 에러 발생.")
	}

	for _, 종목1 := range 종목정보_맵_원본 {
		종목2 := F종목_검색(종목1.G종목코드())

		if 종목2 == nil {
			테스트.Error("TestF종목검색() : 검색결과 nil.",
				종목1.G종목코드()+" "+종목1.G종목명칭())
		}

		if 종목1.G종목코드() != 종목2.G종목코드() ||
			종목1.G종목코드2() != 종목2.G종목코드2() ||
			종목1.G종목명칭() != 종목2.G종목명칭() ||
			종목1.G종목명칭2() != 종목2.G종목명칭2() ||
			종목1.G발행기관코드() != 종목2.G발행기관코드() ||
			종목1.G시장구분() != 종목2.G시장구분() {
			테스트.Error("TestF종목검색() : 내용 불일치.",
				종목1.G종목코드()+" "+종목1.G종목명칭())
		}
	}
}

func TestF종목정보_맵_DB기록(테스트 *testing.T) {
	F테스트_모드_시작()
	defer F테스트_모드_종료()
	F종목정보_테이블_생성()
	defer F테이블_삭제(F종목정보_테이블())

	종목정보_맵_원본 := F종목정보_맵_테스트용()
	에러 := F종목정보_맵_DB기록(종목정보_맵_원본)
	if 에러 != nil {
		테스트.Error("TestF종목정보_맵_DB기록() : F종목정보_맵_DB기록() 에러 발생.", 에러)
	}

	F종목정보_맵_초기화()
	종목정보_맵, 에러 := F종목정보_맵()
	if 에러 != nil {
		테스트.Error("TestF종목정보_맵_DB기록() : F종목정보_맵() 에러 발생.", 에러)
	}

	if len(종목정보_맵) != len(종목정보_맵_원본) {
		테스트.Error("TestF종목정보_맵_DB기록() : 기록된 레코드 수량이 예상과 다릅니다. " +
			"예상값 : " + string(len(종목정보_맵_원본)) +
			", 실제값 : " + string(len(종목정보_맵)))
	}

	for _, 종목1 := range 종목정보_맵_원본 {
		종목2 := 종목정보_맵[종목1.G종목코드()]
		if 종목2 == nil {
			테스트.Error("TestF종목정보_맵_DB기록() : 존재하지 않는 종목.",
				종목1.G종목코드()+" "+종목1.G종목명칭())
		}

		if 종목1.G종목코드() != 종목2.G종목코드() ||
			종목1.G종목코드2() != 종목2.G종목코드2() ||
			종목1.G종목명칭() != 종목2.G종목명칭() ||
			종목1.G종목명칭2() != 종목2.G종목명칭2() ||
			종목1.G발행기관코드() != 종목2.G발행기관코드() ||
			종목1.G시장구분() != 종목2.G시장구분() {
			테스트.Error("TestF종목정보_맵_DB기록() : 내용 불일치.",
				종목1.G종목코드()+" "+종목1.G종목명칭())
		}
	}
}
