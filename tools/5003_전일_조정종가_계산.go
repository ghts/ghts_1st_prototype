package tools

import (
	"log"
	공통 "pts/common"
)

func F전종목_전일_조정종가_맵() (map[uint64]float64, error) {
	전종목_전일_조정종가_맵 := make(map[uint64]float64)

	종목정보_맵, 에러 := 공통.F종목정보_맵()
	if 에러 != nil {
		log.Println("tools.F전일_조정종가_계산() : 공통.F종목정보_맵() 에러", 에러)

		return nil, 에러
	}

	for _, 종목 := range 종목정보_맵 {
		종목별_전일_조정종가_맵, 에러 := F종목별_전일_조정종가_맵(종목)
		if 에러 != nil {
			log.Println("tools.F전종목_전일_조정종가() : "+종목.G종목코드()+" F종목별_전일_조정종가() 에러 발생.", 에러)

			return nil, 에러
		}

		for 키, 전일_조정종가 := range 종목별_전일_조정종가_맵 {
			전종목_전일_조정종가_맵[키] = 전일_조정종가
		}
	}

	return 전종목_전일_조정종가_맵, nil
}

func F종목별_전일_조정종가_맵(종목 *공통.C종목) (map[uint64]float64, error) {
	var 전일_조정종가 float64 = 0.0
	전일_조정종가_맵 := make(map[uint64]float64)

	일일가격정보_모음, 에러 := 공통.F종목별_일일가격정보_모음(종목)
	if 에러 != nil {
		log.Println("tools.F종목별_전일_조정종가() : common.F종목별_일일가격정보_모음() 에러 발생.")
		log.Println(에러)

		return nil, 에러
	}

	일일가격정보_슬라이스 := 일일가격정보_모음.G슬라이스()

	for _, 일일가격정보 := range 일일가격정보_슬라이스 {
		if 일일가격정보.G시가() == 0.0 ||
			일일가격정보.G고가() == 0.0 ||
			일일가격정보.G저가() == 0.0 ||
			일일가격정보.G종가() == 0.0 ||
			일일가격정보.G조정종가() == 0.0 {
			continue
		}

		전일_조정종가_맵[일일가격정보.G식별코드()] = 전일_조정종가
		전일_조정종가 = 일일가격정보.G조정종가()
	}

	return 전일_조정종가_맵, nil
}
