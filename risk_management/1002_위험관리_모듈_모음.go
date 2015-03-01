package risk_management

import (
	공통 "pts/common"
)

/*
type I전략_위험관리 interface {
	G진입허용(신호 *공통.C신호) bool
	G추가매수허용(신호 *공통.C신호) bool
} */

type S전략_위험관리_모듈_모음 struct {
	위험관리_모듈_모음 []공통.I전략_위험관리
}

func (s *S전략_위험관리_모듈_모음) M위험관리_모듈_추가(위험관리_모듈 공통.I전략_위험관리) {
	if s.위험관리_모듈_모음 == nil {
		s.위험관리_모듈_모음 = make([]공통.I전략_위험관리, 0)
	}

	s.위험관리_모듈_모음 = append(s.위험관리_모듈_모음, 위험관리_모듈)
}

func (s *S전략_위험관리_모듈_모음) G진입허용(신호 *공통.C신호) bool {
	for _, 위험관리_모듈 := range s.위험관리_모듈_모음 {
		if !위험관리_모듈.G진입허용(신호) {
			return false
		}
	}

	return true
}

func (s *S전략_위험관리_모듈_모음) G추가매수허용(신호 *공통.C신호) bool {
	for _, 위험관리_모듈 := range s.위험관리_모듈_모음 {
		if !위험관리_모듈.G추가매수허용(신호) {
			return false
		}
	}

	return true
}
