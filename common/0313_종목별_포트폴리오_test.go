package common

import (
	"testing"
)

func TestS종목별_포트폴리오(테스트 *testing.T) {
	var 종목별_포트폴리오_인터페이스 I종목별_포트폴리오 = new(S종목별_포트폴리오)
	종목별_포트폴리오_인터페이스.G종목()	// 컴파일 에러 막는 용도이며, 다른 의미 없음.
}