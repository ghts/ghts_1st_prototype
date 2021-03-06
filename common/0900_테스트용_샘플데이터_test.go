package common

import (
	"testing"
)

func TestF종목정보_맵_테스트용(테스트 *testing.T) {
	맵 := F종목정보_맵_테스트용()

	종목1 := 맵[F종목_동화약품().G종목코드()]
	종목2 := 맵[F종목_삼성전자().G종목코드()]

	if 종목1 == nil || 종목2 == nil {
		테스트.Error("TestF종목정보_맵_테스트용() : 종목검색 결과가 nil 입니다.")
	}
}

func TestF테스트용_비어있는_전략(테스트 *testing.T) {
	var 전략_인터페이스 I전략 = F테스트용_비어있는_전략(100)
	
	if 전략_인터페이스.G식별코드() != 100 {
		테스트.Error("common.TestF테스트용_비어있는_전략() : G식별코드() 불일치. 예상값 100, 실제값 %v.", 전략_인터페이스.G식별코드())
	}
}

func TestF테스트용_비어있는_위험관리(테스트 *testing.T) {
	var 위험관리_인터페이스 I위험관리 = F테스트용_비어있는_위험관리()
	위험관리_인터페이스.G식별코드()
}

func TestF테스트용_종목별_포트폴리오(테스트 *testing.T) {
	var 종목별_포트폴리오_인터페이스 I종목별_포트폴리오 = new(S테스트용_종목별_포트폴리오)
	종목별_포트폴리오_인터페이스.G식별코드()
}

func TestF테스트용_포트폴리오(테스트 *testing.T) {
	var 포트폴리오_인터페이스 I포트폴리오 = new(S테스트용_포트폴리오)
	포트폴리오_인터페이스.G식별코드()
}