package common

import (
	"testing"
	"log"
)

func TestS포트폴리오_모음(테스트 *testing.T) {
	log.Println("TODO : TestS포트폴리오_통합관리()")
	
	/*
	type S테스트용_종목별_포트폴리오 struct {
		M식별코드 uint64
		M종목 *C종목	
		M단가 *C통화
		
		M롱포지션_수량 int64
		M숏포지션_수량 int64
		M순_수량 int64
		M총_수량 int64
		
		M롱포지션_금액() *S통화
		M숏포지션_금액() *S통화
		M순_금액() *S통화
		M총_금액() *S통화
	}
	
	type S테스트용_포트폴리오 struct {
		M식별코드 uint64

		M포트폴리오_내용 []I종목별_포트폴리오

		M롱포지션_금액 *S통화
		M숏포지션_금액 *S통화
		M순_금액() *S통화
		M총_금액() *S통화
	} */
	
	포트폴리오_모음 := F포트폴리오_모음_생성(100, "테스트용_포트폴리오_모음")
	
	var 포트폴리오_인터페이스 I포트폴리오 = 포트폴리오_모음; 포트폴리오_인터페이스.G식별코드()
}