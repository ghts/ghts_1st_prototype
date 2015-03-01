package common

import (
	"math/big"
)

type I정수_식별코드 interface { G식별코드() uint64 }
type I문자열_식별코드 interface { G식별코드() string }

type I기본_문자열 interface { String() string }

type I자료형_공통 interface {
	G같음(비교값 interface{}) bool
	I기본_문자열
}

type I상수형 interface {
	I공통
	상수형임()
}

type I변수형 interface {
	I공통
	변수형임()
}

type I큰정수형 interface { G큰정수() big.Int }
type I정밀수형 interface { G정밀수() big.Rat }
type I문자열형 interface { G문자열() string }
type I고정소숫점형 interface { G고정소숫점(소숫점_이하_자릿수 int) C고정소숫점 }
type I통화형 interface { G통화() C통화 }


type I통화종류 interface { G통화종류() P통화 }

type I환율 interface {
	I통화종류
	G기준통화() P통화
	G환율() C고정소숫점
}

type I종목별_포트폴리오 interface {
	I정수_식별코드
	I통화종류

	G종목() C종목
	G단가() C통화

	G롱포지션_수량() uint64
	G숏포지션_수량() uint64
	G순_수량() uint64
	G총_수량() uint64

	G롱포지션_금액() C통화
	G숏포지션_금액() C통화
	G순_금액() C통화
	G총_금액() C통화
}

type C종목별_포트폴리오 interface {
	I상수형
	I종목별_포트폴리오
}

type V종목별_포트폴리오 interface {
	I변수형
	I종목별_포트폴리오
	
	S종목(종목 C종목)
	S단가(단가 C통화)
	
	S롱포지션_수량(수량 uint64)
	S숏포지션_수량(수량 uint64)
	S순_수량(수량 uint64)
	S총_수량(수량 uint64)

	S롱포지션_금액(금액 C통화)
	S숏포지션_금액(금액 C통화)
	S순_금액(금액 C통화)
	S총_금액(금액 C통화)
}

type I포트폴리오 interface {
	I정수_식별코드
	I통화종류

	G보유_종목_모음() []C종목
	G종목별_포트폴리오(종목코드 string) C종목별_포트폴리오
	G전종목_포트폴리오() []C종목별_포트폴리오

	G롱포지션_금액() C통화
	G숏포지션_금액() C통화
	G순_금액() C통화
	G총_금액() C통화
}

type C포트폴리오 interface {
	I상수형
	I포트폴리오
}

type V포트폴리오 interface {
	I변수형
	I포트폴리오
	
	S종목별_포트폴리오(종목별_포트폴리오 C종목별_포트폴리오)
	G상수형() C포트폴리오
}

type I위험관리 interface {
	I정수_식별코드	// 이게 필요한가? DB에 저장할 것인가?
	I통화종류
	G이름() string
	G필요한_매개변수_모음() []string
	G검토(위험관리_매개변수 C위험관리_매개변수) (C위험관리_검토결과, error)
}
