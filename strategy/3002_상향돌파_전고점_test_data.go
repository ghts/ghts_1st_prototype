package base

import (
	공통 "pts/common"
	"time"
)

func 고점_찾기_테스트용_가장_간단한_모의_데이터() *공통.S일일가격정보_모음 {
	/*
		시간이 흐르면서 계속 상승하다가,
		어느 시점부터 계속 하락하는 모의 데이터.
		고점이 어디인지가 너무나 명확하므로 고점 찾는 로직을 테스트 하기에 좋다.
	*/

	일일가격정보_모음 := make([]*공통.C일일가격정보, 0)

	종목 := 공통.F종목_가상_1()
	var 식별코드 uint64 = 0
	일자, _ := time.Parse("2006-01-02", "2010-01-01")
	var 전일시가, 전일종가 float64 = 100000, 100500
	var 시가, 종가 float64
	최대반복횟수 := 1000
	중간지점 := 최대반복횟수 / 2

	for i := 0; i < 최대반복횟수; i++ {
		식별코드 = 식별코드 + 1
		일자 = 일자.AddDate(0, 0, 1)

		나머지 := i % 7

		if 나머지 == 0 || 나머지 == 1 {
			// 토, 일요일  주말에는 가격 데이터가 없는 것을 흉내냄.
			시가 = 0
			종가 = 0
		} else if i < 중간지점 {
			시가 = 전일시가 + 10
			전일시가 = 시가

			종가 = 전일종가 + 10
			전일종가 = 종가
		} else {
			시가 = 전일시가 - 10
			전일시가 = 시가

			종가 = 전일종가 - 10
			전일종가 = 종가
		}

		일일시세 := new(공통.S일일가격정보)
		일일시세.S식별코드(식별코드)
		일일시세.S종목(종목)
		일일시세.S일자(일자)
		일일시세.S시가(시가)
		일일시세.S종가(종가)

		if 시가 > 종가 {
			일일시세.S고가(시가)
			일일시세.S저가(종가)
		} else {
			일일시세.S고가(종가)
			일일시세.S저가(시가)
		}

		일일시세.S조정종가(종가)
		일일시세.S거래량(1000)
		//일일시세.S전일종가(??)
		일일시세.M조정가격재계산()

		일일가격정보_모음 = append(일일가격정보_모음, 일일시세.M상수형구조체생성())
	}

	반환값 := new(공통.S일일가격정보_모음)
	반환값.S내용(일일가격정보_모음)

	return 반환값
}

func 고점_찾기_테스트용_주기적으로_하루씩_등락이_있는_모의_데이터() *공통.S일일가격정보_모음 {
	/*
		시간이 흐르면서 계속 상승하다가,
		어느 시점부터 계속 하락하지만 며칠 간격으로 하루씩 등락이 있는 모의 데이터
		고점이 어디인지가 너무나 명확하므로 고점 찾는 로직을 테스트 하기에 좋다.
	*/

	일일가격정보_모음 := make([]*공통.C일일가격정보, 0)

	종목 := 공통.F종목_가상_1()
	var 식별코드 uint64 = 0
	일자, _ := time.Parse("2006-01-02", "2010-01-01")
	var 전일시가, 전일종가 float64 = 100000, 100500
	var 시가, 종가 float64
	최대반복횟수 := 1000
	중간지점 := 최대반복횟수 / 2

	for i := 0; i < 최대반복횟수; i++ {
		식별코드 = 식별코드 + 1
		일자 = 일자.AddDate(0, 0, 1)

		나머지 := i % 7

		if 나머지 == 0 || 나머지 == 1 {
			// 토, 일요일  주말에는 가격 데이터가 없는 것을 흉내냄.
			시가 = 0
			종가 = 0
		} else if i < 중간지점 {
			if 나머지 == 3 || 나머지 == 5 {
				시가 = 전일시가 - 5
				전일시가 = 시가

				종가 = 전일종가 - 5
				전일종가 = 종가
			} else {
				시가 = 전일시가 + 10
				전일시가 = 시가

				종가 = 전일종가 + 10
				전일종가 = 종가
			}
		} else {
			if 나머지 == 4 || 나머지 == 6 {
				시가 = 전일시가 + 5
				전일시가 = 시가

				종가 = 전일종가 + 5
				전일종가 = 종가
			} else {
				시가 = 전일시가 - 10
				전일시가 = 시가

				종가 = 전일종가 - 10
				전일종가 = 종가
			}
		}

		일일시세 := new(공통.S일일가격정보)
		일일시세.S식별코드(식별코드)
		일일시세.S종목(종목)
		일일시세.S일자(일자)
		일일시세.S시가(시가)
		//일일시세.S고가(??)
		//일일시세.S저가(??)
		일일시세.S종가(종가)
		일일시세.S조정종가(종가)
		일일시세.S거래량(1000)
		//일일시세.S전일종가(??)
		일일시세.M조정가격재계산()

		일일가격정보_모음 = append(일일가격정보_모음, 일일시세.M상수형구조체생성())
	}

	반환값 := new(공통.S일일가격정보_모음)
	반환값.S내용(일일가격정보_모음)

	return 반환값
}

func 고점_찾기_테스트용_주기적으로_이틀_연속_등락이_있는_모의_데이터() *공통.S일일가격정보_모음 {
	/*ㄹ
	시간이 흐르면서 계속 상승하다가,
	어느 시점부터 계속 하락하지만 며칠 간격으로 이틀 연속 등락이 있는 모의 데이터.
	고점이 어디인지가 너무나 명확하므로 고점 찾는 로직을 테스트 하기에 좋다.
	*/

	일일가격정보_모음 := make([]*공통.C일일가격정보, 0)

	종목 := 공통.F종목_가상_1()
	var 식별코드 uint64 = 0
	일자, _ := time.Parse("2006-01-02", "2010-01-01")
	var 전일시가, 전일종가 float64 = 100000, 100500
	var 시가, 종가 float64
	최대반복횟수 := 1000
	중간지점 := 최대반복횟수 / 2

	for i := 0; i < 최대반복횟수; i++ {
		식별코드 = 식별코드 + 1
		일자 = 일자.AddDate(0, 0, 1)

		나머지 := i % 7

		if 나머지 == 0 || 나머지 == 1 {
			// 토, 일요일  주말에는 가격 데이터가 없는 것을 흉내냄.
			시가 = 0
			종가 = 0
		} else if i < 중간지점 {
			if 나머지 == 3 || 나머지 == 4 {
				시가 = 전일시가 - 4
				전일시가 = 시가

				종가 = 전일종가 - 4
				전일종가 = 종가
			} else {
				시가 = 전일시가 + 10
				전일시가 = 시가

				종가 = 전일종가 + 10
				전일종가 = 종가
			}
		} else {
			if 나머지 == 5 || 나머지 == 6 {
				시가 = 전일시가 + 4
				전일시가 = 시가

				종가 = 전일종가 + 4
				전일종가 = 종가
			} else {
				시가 = 전일시가 - 10
				전일시가 = 시가

				종가 = 전일종가 - 10
				전일종가 = 종가
			}
		}

		일일시세 := new(공통.S일일가격정보)
		일일시세.S식별코드(식별코드)
		일일시세.S종목(종목)
		일일시세.S일자(일자)
		일일시세.S시가(시가)
		//일일시세.S고가(??)
		//일일시세.S저가(??)
		일일시세.S종가(종가)
		일일시세.S조정종가(종가)
		일일시세.S거래량(1000)
		//일일시세.S전일종가(??)
		일일시세.M조정가격재계산()

		일일가격정보_모음 = append(일일가격정보_모음, 일일시세.M상수형구조체생성())
	}

	반환값 := new(공통.S일일가격정보_모음)
	반환값.S내용(일일가격정보_모음)

	return 반환값
}
