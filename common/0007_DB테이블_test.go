package common

import (
	"strings"
	"testing"
)

func TestF테스트_모드(테스트 *testing.T) {
	F테스트_모드_시작()
	if !strings.HasPrefix(F종목정보_테이블(), 테스트_테이블_접두사) {
		테스트.Error("TestF테스트_모드() 테스트 에러. 테스트 모드인데 접두사 test_가 붙지 않았습니다.")
	}

	F테스트_모드_종료()
	if strings.HasPrefix(F종목정보_테이블(), 테스트_테이블_접두사) {
		테스트.Error("TestF테스트_모드() 테스트 에러. 정상 모드인데 접두사 test_가 붙었습니다.")
	}
}

func TestDb_종목정보_테이블이_존재하는_지_확인(테스트 *testing.T) {
	F테스트_모드_종료()

	_, 에러 := F_SQL질의_정수("SELECT COUNT(*) FROM " + F종목정보_테이블())

	if 에러 != nil {
		F종목정보_테이블_생성()
	}

	_, 에러 = F_SQL질의_정수("SELECT COUNT(*) FROM " + F종목정보_테이블())

	if 에러 != nil {
		테스트.Error("common.TestDb_종목정보_테이블이_존재하는_지_확인() : 종목정보 테이블이 존재하지 않으며, 생성하는 데 실패했습니다.")
	}
}

func TestDb_일일가격정보_테이블이_존재하는_지_확인(테스트 *testing.T) {
	F테스트_모드_종료()

	_, 에러 := F_SQL질의_정수("SELECT COUNT(*) FROM " + F일일가격정보_테이블())

	if 에러 != nil {
		F일일가격정보_테이블_생성()
	}

	_, 에러 = F_SQL질의_정수("SELECT COUNT(*) FROM " + F일일가격정보_테이블())

	if 에러 != nil {
		테스트.Error("common.TestDb_일일가격정보_테이블이_존재하는_지_확인() : 일일가격정보 테이블이 존재하지 않으며, 생성하는 데 실패했습니다.")
	}
}
