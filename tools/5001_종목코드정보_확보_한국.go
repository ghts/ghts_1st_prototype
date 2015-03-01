package tools

import (
	"encoding/csv"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	공통 "pts/common"
	"strings"
)

var csv파일_경로명 = 실전용_CSV파일_경로명

func 실전용_CSV파일_경로명() string { return "/test_data/stock_code_info_korea.csv" }
func 테스트용_CSV파일_경로명() string {
	return "/test_data/stock_code_info_korea_for_test_only.csv"
}

func F종목코드정보_확보() error {
	// 종목정보
	종목정보_모음, 에러 := f_KRX표준코드정보_CSV파일_읽기()

	if 에러 != nil {
		log.Println("종목정보 읽어들이기 에러", 에러)
		return 에러
	}

	에러 = 공통.F종목정보_맵_DB기록(종목정보_모음)
	if 에러 != nil {
		log.Println("종목정보 DB기록 에러", 에러)
		return 에러
	}

	return nil
}

// 	http://isin.krx.co.kr 참조
func f_KRX표준코드정보_CSV파일_읽기() (map[string]*공통.C종목, error) {
	종목_모음 := make(map[string]*공통.C종목)

	CSV파일경로, 에러 := 공통.F파일경로로_파일찾기(csv파일_경로명())
	if 에러 != nil {
		log.Println("f_KRX표준코드정보_읽기() : common.F파일경로로_파일찾기() 에러.", 에러)
		return nil, 에러
	}

	파일, 에러 := os.Open(CSV파일경로)

	if 에러 != nil {
		log.Println("f_KRX표준코드정보_읽기() : os.Open() (파일열기) 에러.", 에러)
		return nil, 에러
	}
	defer 파일.Close()

	레코드_모음, 에러 := csv.NewReader(파일).ReadAll()

	if 에러 != nil {
		log.Println("f_KRX표준코드정보_읽기() CSV파일 읽기에러 : ", 에러)
		return nil, 에러
	}

	레코드_모음 = 레코드_모음[1:] // 첫째 줄 제목 제거.

	for _, 레코드 := range 레코드_모음 {
		종목 := 공통.F종목_생성(
			0, // 식별코드 (아직까지 알 수 없으며, DB에 기록할 때 자동으로 부여됨. 일단 0으로 설정)
			strings.TrimSpace(레코드[4]), // 종목코드 (단축코드)
			strings.TrimSpace(레코드[3]), // 종목코드2 (표준코드)
			strings.TrimSpace(레코드[1]), // 종목명칭
			strings.TrimSpace(레코드[2]), // 종목명칭2
			strings.TrimSpace(레코드[0]), // 발행기관 코드
			strings.TrimSpace(레코드[5])) // 시장구분

		종목_모음[종목.G종목코드()] = 종목
	}

	return 종목_모음, 에러
}
