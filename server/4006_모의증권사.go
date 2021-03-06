package server

import (
	"runtime"
	"time"
	공통 "pts/common"
)

// 지금은 일봉자료 밖에 없으니, 굳이 복잡하게 하지 말고,
// 일봉 데이터 기준으로 모의증권사 모듈을 간단하게 만들고,
// 나중에 더 정교하게 하던가, 설정을 통해서 주문체결 기능을 적절히 교체할 수 있도록 하자.
type S모의증권사_일봉 struct {
	증권사 *C증권사
	// 가격정보전송, 주문접수 및 처리 모두 서버를 통해서 하므로 전략그룹정보를 알 필요가 없다.
	//전략그룹모음 []공통.I전략그룹	
	//가격정보_수집종목_모음 []*C종목	// 모의서버는 굳이 필요없을 듯.
	
	// 가격정보전송, 주문접수 및 처리 모두 서버를 통해서 하므로 전략그룹정보를 알 필요가 없다.
	//전략그룹정보구독채널 chan []공통.I전략그룹
	// 모의서버는 주기적으로 가격정보를 갱신하는 게 아니라, 
	// 파일에서 예전 가격정보를 한꺼번에 읽어서 순차적으로 전송하면 되니까 틱은 필요없다.
	//틱채널 chan int8	// 가격정보 갱신주기
	주문접수채널 chan *공통.C주문
	공통이벤트채널 chan int8
	
	// 모의서버는 순차적으로 처리해야 하므로 일반 메소드로 처리하므로 
	// go루틴의 최대동시처리갯수를 제한하기 위한 처리권한 자체가 필요없다.
	//주문접수_처리권한 chan int8
}

func (s *S모의증권사) M실행(설정 *C서버설정) {
	설정.G종료대기열().Add(1)
	defer 설정.G종료대기열().Done()
	
	s.초기화(설정)
	
	go s.가격정보전송()

	for {
		select {
		// 가격정보전송, 주문접수 및 처리 모두 서버를 통해서 하므로 전략그룹정보를 알 필요가 없다.
		//case 전략그룹모음 := <-s.전략그룹정보구독채널:	
		//	s.전략그룹모음 = 전략그룹모음
		// 모의서버는 별도의 go루틴에서 파일을 읽어서 가격정보 전송하면 되므로, 틱이 필요없다.
		// 그리고, 실제 서버에서도 가격정보갱신은 별도의 go루틴으로 독립시키는 게 좋을 듯 하다.
		//case <- 틱채널:
		//	s.가격정보갱신()
		// 거래전략그룹에서도 가격정보를 순차적으로 1개씩만 처리하면 
		// 가격정보의 시점 및 주문접수 및 처리의 시점 차이 문제는 없지 않을까 싶다.
		// (예, 가격정보 전송한 지 한참 지난 후, 한참 이전 시점의 주문이 접수되는 문제, 
		//   이전 시점의 주문과 그 이후 시점의 주문이 뒤엉키는 문제.)		
		// 거래전략그룹에서 개별전략인스턴스가 가격정보를 1개씩 처리하도록 보장하는 방법은
		// 더 연구해 볼 것.
		// 개별 전략인스턴스 혹은 미청산초도매수신호 go루틴이 채널을 통해서 가격정보을 받고, 
		// 처리완료 후 채널로 전략그룹에게 처리완료 회신을 보내는 것이 간단하고 확실한 방법일 듯 하다.
		// 처리는 2가지로 나뉘는 데, 신호가 발생되지 않는 경우에는 그냥 처리완료 회신을 보내면 되고,
		// 신호를 발생된 경우, 주문을 내고, 주문처리결과를 회신받은 후 처리완료 회신을 보내면 된다.
		// 어느 방법이던 간에 백테스팅 때는 순차실행을 위해서 채널의 버퍼를 0으로 설정하면 되고,
		// 실제 운용 시에 동시처리 성능을 높이고자 하면 버퍼를 높여주면 된다.
		// 이렇게 하면 같은 로직으로 채널 버퍼만 조절하여 순차처리와 동시처리 모두 구현할 수 있다.
		// 또, 버퍼로 인해서 문제가 생기더라도 버퍼만 없애면 순차처리로 원상복귀되므로 간단하게 대처할 수 있다.
		case 주문 := <-s.주문접수채널:
			// 모의서버는 순차적으로 진행되어야 하므로 go루틴을 사용하지 않고, 
			// 일반적인 메소드로 처리하자.
			s.주문처리(주문)
		case 이벤트 := <-s.공통이벤트채널:
			switch 이벤트 {
			case 공통.비상탈출이벤트:
				s.비상탈출()
			case 공통.종료이벤트:
				s.종료()				
			default:
				panic("S모의증권사.M실행() : 예상치 못한 공통 이벤트.")
			}
		default: 
			time.Sleep(50 * time.Millisecond)
		}
	}
}


func (s *S모의증권사) 초기화(설정 *C서버설정) {
	s.전략그룹정보구독채널 = make(chan []공통.I전략그룹, 10)
	s.공통이벤트채널 = make(chan int8, 10)
	s.주문전송채널 = make(chan *공통.C주문, 100)
	
	// 전략그룹정보 구독신청.
	서버.G전략그룹정보구독신청채널() <- s.전략그룹정보구독채널
	
	// 채널을 세마포어처럼 이용해서 최대 동시처리 숫자를 조절하는 패턴 (Effective Go 참조)
	//가격정보전송_최대동시처리숫자 = 1	// 모의서버는 굳이 많이 할 필요없을 듯.
	s.가격정보전송_처리권한 = make(chan int8, 가격정보전송_최대동시처리숫자)
}

// 모의증권사의 가격정보는 서버에 접속하여 가격정보를 받는 게 아니라,
// 파일에서 가격정보를 읽어서 전송하는 것임.
// TODO 파일에서 가격정보 읽어들여서 시간역순으로 정렬해서 가격정보모음 슬라이스([])에 저장하기.
// func 가격정보전송(가격정보 *공통.C가격정보) {
func (s *S모의증권사) 가격정보전송() {
	// 가격정보는 순차처리를 신경쓰지 않고 그냥 보내면 된다. 
	// 가격정보를 접수 혹은 수신하는 거래전략그룹 측의 채널 버퍼를 조절하면 순차처리와 동시처리를 조절가능.
	// 백테스팅 때 거래전략그룹이 수신채널의 버퍼를 0으로 설정하면 1개씩 순차적으로 처리하게 되고,
	// 실제 거래할 때 거래전략그룹이 수신채널의 버퍼를 늘리면 동시처리 성능이 높아진다.
	
	// 1. 파일에서 가격정보 읽어들이기. (한 종목이건 여러 종목이건 상관없이 한 군데로 모은다.)
	// 2. 모든 가격정보를 시간순서해도 정렬하기 (종목에 상관없이 시간 순서대로.)
	// 3. 시간 순서대로 가격정보를 거래전략그룹의 가격정보 수신채널로 전송한다.
	//    (가격정보의 실제 처리는 거래전략그룹에게 맡기고, 서버는 전송하는 것으로 임무완료.)
	
	var 가격정보모음 []*공통.C가격정보
	
	// TODO 파일에서 가격정보 읽어들여서 시간역순으로 정렬해서 가격정보모음 슬라이스([])에 저장하기.
	
	for _, 가격정보 := range 가격정보모음 {
		for _, 전략그룹 = range s.전략그룹모음 {
			<-s.가격정보전송_처리권한	// 처리권한 획득
		
			// 변수를 복사하여 go루틴 간의 독립성 보장.
			채널 := 전략그룹.G가격정보채널()
			정보 := 가격정보
			
			go func() {
				채널 <- 정보
				s.가격정보전송_처리권한 <- 1	// 처리권한 반환.
			}()
		}
	}
	
	// 여기까지
}

// TODO 주문체결지연 구현.
// 가격정보시점이 시가이면 종가, 종가이면 다음 거래일 시가로 하면 되는가?
// 백테스팅에서는 주문생성시점이 별 도움이 안 되는 데, 가격정보시점으로 시가, 종가 구분이 되는가?
// 가격정보 보낼 때 구분이 되게 보내면 된다.
// func (s *S모의증권사) 가격정보전송()  에서 가격정보의 시가, 종가가 구분되게 하자.
func 주문처리(주문 *공통.C주문) {
	// 식별코드는 어떻게 생성하지?
	// 데이터베이스에 저장할 때 자동으로 생성되는 id로 하면 될 듯 한데.
	주문처리결과 := 공통.F주문처리결과생성(
		식별코드 uint64,
		주문,
		공통.체결완료,
		주문.G단가(),
		주문.G수량(),
		time.Now())
	
	주문.G처리결과통보채널() <- 주문처리결과	
}

func (s *S가격정보제공) 비상탈출() {}	// 현재로서는 모의서버는 비상탈출이 필요없을 듯 하다.
func (s *S가격정보제공) 종료() { runtime.Goexit() }