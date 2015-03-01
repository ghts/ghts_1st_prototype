package tools

import (
	"log"
	"os"
	공통 "pts/common"
	"testing"
)

func TestF메모(테스트 *testing.T) {
	공통.F메모()
}

func TestCsv파일이_존재하는_지_확인(테스트 *testing.T) {
	파일, 에러 := os.Open("./test_data/stock_code_info_korea.csv")

	if 파일 != nil {
		defer 파일.Close()
	}

	if 에러 != nil {
		테스트.Error("tools.TestCsv파일이_존재하는_지_확인(). CSV파일이 존재하지 않습니다.")
	}
}

func TestCsv파일_읽기(테스트 *testing.T) {
	csv파일_경로명 = 테스트용_CSV파일_경로명

	종목정보_모음, 에러 := f_KRX표준코드정보_CSV파일_읽기()

	if 에러 != nil {
		테스트.Error("tools.TestCsv파일_읽기(). 실행 중 에러 발생.", 에러)
	}

	if len(종목정보_모음) != 1 {
		테스트.Error("tools.TestCsv파일_읽기(). 종목 갯수가 예상과 다름.", len(종목정보_모음))
	}

	종목 := 종목정보_모음["A000010"]

	if 종목 == nil {
		테스트.Error("TestCsv파일_읽기(). 읽어들인 종목코드가 예상과 다름.")
	}

	if 종목.G종목코드() != "A000010" ||
		종목.G종목코드2() != "KR7000010009" ||
		종목.G종목명칭() != "신한은행보통주" ||
		종목.G종목명칭2() != "SHINHAN BANK" ||
		종목.G발행기관코드() != "1" ||
		종목.G시장구분() != "유가증권시장상장폐지" {
		테스트.Error("TestCsv파일_읽기(). 읽어들인 종목의 내용이 예상과 다름.")
		log.Println("종목코드 :", 종목.G종목코드())
		log.Println("종목코드2 :", 종목.G종목코드2())
		log.Println("종목명칭 :", 종목.G종목명칭())
		log.Println("종목명칭2 :", 종목.G종목명칭2())
		log.Println("발행기관 코드 :", 종목.G발행기관코드())
		log.Println("시장구분 :", 종목.G시장구분())
	}
}

func TestF종목코드정보_확보(테스트 *testing.T) {
	공통.F테스트_모드_시작()
	defer 공통.F테스트_모드_종료()

	공통.F종목정보_테이블_생성()
	defer 공통.F테이블_삭제(공통.F종목정보_테이블())

	csv파일_경로명 = 테스트용_CSV파일_경로명

	에러 := F종목코드정보_확보()
	if 에러 != nil {
		테스트.Error("tools.TestF종목코드정보_확보() : F종목코드정보_확보() 에러.", 에러)
	}

	수량, 에러 := 공통.F_SQL질의_정수("SELECT COUNT(*) FROM " + 공통.F종목정보_테이블())
	if 에러 != nil {
		테스트.Error("tools.TestF종목코드정보_확보() : common.F_SQL질의_정수() 에러.", 에러)
	}

	if 수량 == 0 {
		테스트.Error("tools.TestF종목코드정보_확보() : 종목코드 정보가 기록되지 않았습니다.")
	}
}

/*
func TestF실제로_종목코드정보_확보(테스트 *testing.T) {
	공통.F테스트_모드_종료()
	공통.F종목정보_테이블_생성()

	csv파일_경로명 = 실전용_CSV파일_경로명

	에러 := F종목코드정보_확보()
	if 에러 != nil {
		테스트.Error("tools.TestF종목코드정보_확보() : F종목코드정보_확보() 에러.", 에러)
	}

	수량, 에러 := 공통.F_SQL질의_정수("SELECT COUNT(*) FROM " + 공통.F종목정보_테이블())
	if 에러 != nil {
		테스트.Error("tools.TestF종목코드정보_확보() : common.F_SQL질의_정수() 에러.", 에러)
	}

	log.Printf("%v개의 종목정보를 확보되었습니다.", 수량)
} */
