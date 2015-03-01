package base

import (
	공통 "pts/common"
	"time"
)

func 종료시점_이전에_정상적인_가격데이터_존재여부_확인(일일가격정보_모음 *공통.S일일가격정보_모음, 시작시점 time.Time, 종료시점 time.Time) bool {
	if 시작시점.After(종료시점) {
		return false
	}

	일일가격정보_내용 := 일일가격정보_모음.G내용()

	for _, 일일가격정보 := range 일일가격정보_내용 {
		if 일일가격정보.G일자().Before(시작시점) {
			continue
		}

		// 종료시점 이전만을 고려할려면 종료시점을 포함하지 않는 것이 중요함.
		if 일일가격정보.G일자().Equal(종료시점) ||
			일일가격정보.G일자().After(종료시점) {
			break
		}

		// 비정상적인 가격 데이터 무시
		if 일일가격정보.G시가() == 0.0 ||
			일일가격정보.G종가() == 0.0 ||
			일일가격정보.G조정시가() == 0.0 ||
			일일가격정보.G조정종가() == 0.00 ||
			일일가격정보.G거래량() == 0 {
			continue
		}

		// 정상적인가격데이터_찾음
		return true
	}

	return false
}
