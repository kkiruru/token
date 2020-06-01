package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// 프로그램이 종료될 때까지 대기할 WaitGroup
var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 애플리케이션 진입점
func main() {

	// 고루틴당 하나씩, 총n n개의 카운터를 추가한다.
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go api_call()
	}

	// 호출이 끝날 때까지 기다린다.
	wg.Wait()
}

func api_call() {
	//함수가 리턴될 때 lock counter를 줄여준다.
	defer wg.Done()

	// 몇개까지 이상이 없는지 counter 숫자 찍어본다
	mutex.Lock()
	{
		fmt.Printf("api_call %d \n", atomic.LoadInt64(&counter))
		atomic.AddInt64(&counter, 1)
	}
	mutex.Unlock()

	resp, err := http.Get("http://localhost:8080/api1")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 결과 출력
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data))
}
