package common

import (
//"bytes"
//"encoding/csv"
//"errors"
//"path/filepath"
	"log"
//"os"
//"strings"
//"time"
//"strconv"
)

func F나를_위한_문구() {
	log.Println("")
	log.Println("----------------------------------------------------------")
	log.Println("	쉽고 간단하게, 테스트로 검증해 가면서 마음 편하게.")
	log.Println("----------------------------------------------------------")
	log.Println("")
}
func F메모() {
	F나를_위한_문구()
	log.Println("----------------------------------------------------------")
	log.Println("                  메            모")
	log.Println("----------------------------------------------------------")
	log.Println("")
	log.Println("TODO : I통화, I통화_구조체, c통화, s통화, TestC통화(). TestS통화(). 마무리.")
	log.Println("TODO : TestS종목별_포트폴리오(), TestS포트폴리오_통합관리(), TestS종목별_포트폴리오_통합관리().")
	log.Println("TODO : I종목별_포트폴리오_통합관리. G단가() 기준.")	
	log.Println("TODO : C포트폴리오내역구성원, C포트폴리오내역구성원 수정.")
	log.Println("       0203 C포트폴리오내역구성원 : G현재단가(), G매입금액().")
	log.Println("       0203 C포트폴리오내역 : GSharpe비율(), G연평균수익률().")	
	log.Println("TODO : 위험관리 전략 구현. I포트폴리오 구현체가 완성된 후 진행.")
	log.Println("TODO : I전략그룹과 I전략을 하나로 통일. ")
	log.Println("		대부분의 기능이 겹치므로, 하위 원소 I전략을 포함할 수 있기만 하면 될 듯함.")
	log.Println("		별도의 I전략그룹이 꼭 필요하다면 I전략을 포함하는 형태로 할 것.")
	log.Println("		단, 구현체는 S전략그룹과 S개별전략 및 각 개별전략별로 나누어질 수 있음.")
	log.Println("TODO : 0900_테스트용_샘플데이터_test.go 할 것.")
	log.Println("		테스트의 기반이 되는 샘플데이터에 테스트가 없으면 테스트 자체가 취약해 짐.")
	log.Println("TODO : 자주 사용되는 함수 중 panic 가능성이 높은 함수에 recover() 추가.")
	log.Println("TODO : interface{}의 deep copy가 가능한 지 조사해 보고,")
	log.Println("       만약, 가능하다면 common.F맵_복사() 구현할 것.")
	log.Println("TODO : go test pts/common 병목지점 해결.")
	log.Println("")
	log.Println("PLAN : 순수 Go언어 데이터베이스. tiedot, kv, ql, leveldb-go.")
	log.Println("			MySQL에 대한 외부 의존성을 제거하여서 향후 배포할 때 편리할 것이다.")
	log.Println("PLAN : 사용자 UI는 HTML5 기반으로 한다.")
	log.Println("		a. GopherJS, CoffeeScript : Javascript에 적응하는 어려움을 덜어줄 가능성이 있음.")
	log.Println("		b. AngularJS : DOM을 직접 조작해야 하는 어려움을 덜어줄 가능성이 있음.")
	log.Println("			 			GopherJS용 바인딩도 존재함.")
	log.Println("")
}