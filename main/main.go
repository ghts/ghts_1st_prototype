package main

import (
	"log"
	공통 "pts/common"
	도구 "pts/tools"
)

func main() {
	일일가격정보_확보()
}

func 일일가격정보_확보() {
	공통.F테스트_모드_종료()

	에러 := 도구.F종목코드정보_확보()
	if 에러 != nil {
		log.Println("main() : tools.F종목코드정보_확보() 에러.", 에러)

		return
	}

	종목코드별_에러내역_맵 := 도구.F전종목_일일가격정보_확보()
	log.Printf("main() : tools.F전종목_일일가격정보_확보() 중 %v 종목에서 에러 발생.", 종목코드별_에러내역_맵.G에러난_종목코드_수량())

	식별코드별_에러내역_맵, 에러 := 도구.F전종목_일일가격정보_체크_한국()
	if 에러 != nil {
		log.Println("main() : tools.F전종목_일일가격정보_체크_한국() 에러.", 에러)

		return
	}
	log.Printf("main() : tools.F전종목_일일가격정보_체크_한국() 중 %v개 항목에서 에러 발생.", 식별코드별_에러내역_맵.G에러난_식별코드_수량())
}
