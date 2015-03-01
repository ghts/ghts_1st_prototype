package money_management

import (
	공통 "common"
	"testing"
)

/*
type I전략그룹_자금관리 interface {
	G주문수량(신호 *공통.C신호) int64
} */

type 모의_자금관리_모듈_1 struct{}

func (s 모의_자금관리_모듈_1) G주문수량(신호 *공통.C신호) int64 { return 1000 }

type 모의_자금관리_모듈_2 struct{}

func (s 모의_자금관리_모듈_2) G주문수량(신호 *공통.C신호) int64 { return 900 }

type 모의_자금관리_모듈_3 struct{}

func (s 모의_자금관리_모듈_3) G주문수량(신호 *공통.C신호) int64 { return 800 }

func TestS전략_자금관리_모듈_모음(테스트 *testing.T) {
	s := new(S전략_자금관리_모듈_모음)

	s.M자금관리_모듈_추가(new(모의_자금관리_모듈_1))
	if s.G주문수량(nil) != 1000 {
		테스트.Error("TestS전략_자금관리_모듈_모음.G주문수량() 에러 1 : 주문수량이 예상과 다릅니다.", s.G주문수량(nil))
	}

	s.M자금관리_모듈_추가(new(모의_자금관리_모듈_2))
	if s.G주문수량(nil) != 900 {
		테스트.Error("TestS전략_자금관리_모듈_모음.G주문수량() 에러 2 : 주문수량이 예상과 다릅니다.", s.G주문수량(nil))
	}

	s.M자금관리_모듈_추가(new(모의_자금관리_모듈_3))
	if s.G주문수량(nil) != 800 {
		테스트.Error("TestS전략_자금관리_모듈_모음.G주문수량() 에러 3 : 주문수량이 예상과 다릅니다.", s.G주문수량(nil))
	}
}
