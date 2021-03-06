package server

import (
	공통 "pts/common"
	"runtime"
	"time"
)

// 증권사별로 1초에 접속횟수 한도가 있을 경우 (예 : 이트레이드 증권 1초당 20회)
// 수백개 종목의 가격정보를 받는 데 시간이 너무 많이 걸린다.
// 그럴 경우 가격정보를 여러 증권사 서버에서 종목별로 나누어 받으면 된다.
// 이렇게 여러 증권사에서 받은 정보를 가격정보제공모듈에서 모두 모아서 취합한 후,
// 각 전략그룹에 배포하는 일종의 물류센터 역할을 한다.
type S가격정보제공 struct {
	전략그룹모음 []공통.I전략그룹

	전략그룹정보구독채널 chan []공통.I전략그룹
	가격정보채널     chan *공통.C가격정보
	공통이벤트채널    chan int8

	가격정보전송_처리권한 chan int8
}

func (s *S가격정보제공) G공통이벤트채널() chan int8 { return s.공통이벤트채널 }

func (s *S가격정보제공) M실행(설정 *C서버설정) {
	설정.G종료대기열().Add(1)
	defer 설정.G종료대기열().Done()

	s.초기화(설정)

	for {
		select {
		case 전략그룹모음 := <-s.전략그룹정보구독채널:
			s.전략그룹모음 = 전략그룹모음
		case 가격정보 := <-s.가격정보채널:
			s.가격정보전송(가격정보)
		case 이벤트 := <-s.공통이벤트채널:
			switch 이벤트 {
			case 공통.비상탈출이벤트:
				s.비상탈출()
			case 공통.종료이벤트:
				s.종료()
			default:
				panic("S가격정보제공.M실행() : 예상치 못한 공통 이벤트.")
			}
		default:
			// 자주 실행되는 모듈이니 대기시간을 짧게 잡자.
			time.Sleep(30 * time.Millisecond)
		}
	}
}

func (s *S가격정보제공) 초기화(설정 *C서버설정) {
	s.전략그룹정보구독채널 = make(chan []공통.I전략그룹, 10)
	s.가격정보채널 = 설정.G가격정보채널()
	s.공통이벤트채널 = make(chan int8, 10)

	// 전략그룹정보 구독신청.
	설정.G전략그룹정보구독신청채널() <- s.전략그룹정보구독채널

	// 채널을 세마포어처럼 이용해서 최대 동시처리 숫자를 조절하는 패턴 (Effective Go 참조)
	가격정보전송_최대동시처리숫자 := 설정.G가격정보제공모듈설정().G가격정보전송_최대동시처리숫자()
	s.가격정보전송_처리권한 = make(chan int8, 가격정보전송_최대동시처리숫자)
}

func (s *S가격정보제공) 가격정보전송(가격정보 *공통.C가격정보) {
	var 전략그룹 전략.I전략그룹

	for _, 전략그룹 = range s.전략그룹모음 {
		<-s.가격정보전송_처리권한 // 처리권한 획득

		// 변수를 복사하여 go루틴 간의 독립성 보장.
		채널 := 전략그룹.G가격정보채널()
		정보 := 가격정보

		go func() {
			채널 <- 정보
			s.가격정보전송_처리권한 <- 1 // 처리권한 반환.
		}()
	}
}

func (s *S가격정보제공) 비상탈출() {} // TODO. 어떤 작업이 수행되어야 하는가?
func (s *S가격정보제공) 종료()   { runtime.Goexit() }
