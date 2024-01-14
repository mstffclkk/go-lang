package channel

func Channel1(evenCh chan int) {
	sum:=0
	for i:=0; i<=10; i+=2 {
		sum+=i
	}

	evenCh <- sum
}

func Channel2(oddCh chan int) {
	sum:=0
	for i:=1; i<=10; i+=2 {
		sum+=i
	}

	oddCh <- sum
}


