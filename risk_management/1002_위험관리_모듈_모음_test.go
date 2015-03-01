package risk_management

import (
	공통 "pts/common"
	"testing"
)

/*
type I전략_위험관리 interface {
	G진입허용(신호 *공통.C신호) bool
	G추가매수허용(신호 *공통.C신호) bool
} */

type 모의_위험관리_모듈_1 struct{}

func (s 모의_위험관리_모듈_1) G진입허용(신호 *공통.C신호) bool   { return true }
func (s 모의_위험관리_모듈_1) G추가매수허용(신호 *공통.C신호) bool { return true }

type 모의_위험관리_모듈_2 struct{}

func (s 모의_위험관리_모듈_2) G진입허용(신호 *공통.C신호) bool   { return true }
func (s 모의_위험관리_모듈_2) G추가매수허용(신호 *공통.C신호) bool { return true }

type 모의_위험관리_모듈_3 struct{}

func (s 모의_위험관리_모듈_3) G진입허용(신호 *공통.C신호) bool   { return false }
func (s 모의_위험관리_모듈_3) G추가매수허용(신호 *공통.C신호) bool { return false }

func TestS전략_위험관리_모듈_모음(테스트 *testing.T) {
	s := new(S전략_위험관리_모듈_모음)
	s.M위험관리_모듈_추가(new(모의_위험관리_모듈_1))
	s.M위험관리_모듈_추가(new(모의_위험관리_모듈_2))

	if !s.G진입허용(nil) {
		테스트.Error("TestS전략_위험관리_모듈_모음.G진입허용() 에러 1 : 진입허용 여부가 예상과 다릅니다.")
	}

	if !s.G추가매수허용(nil) {
		테스트.Error("TestS전략_위험관리_모듈_모음.G추가매수허용() 에러 1 : 추가매수허용 여부가 예상과 다릅니다.")
	}

	s.M위험관리_모듈_추가(new(모의_위험관리_모듈_3))

	if s.G진입허용(nil) {
		테스트.Error("TestS전략_위험관리_모듈_모음.G진입허용() 에러 2 : 진입허용 여부가 예상과 다릅니다.")
	}

	if s.G추가매수허용(nil) {
		테스트.Error("TestS전략_위험관리_모듈_모음.G추가매수허용() 에러 2 : 추가매수허용 여부가 예상과 다릅니다.")
	}
}
