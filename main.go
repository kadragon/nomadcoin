package main

import (
	"flag"
	"fmt"
	"os"
)

// usage 사용자 메뉴얼을 출력하는 함수
func usage() {
	fmt.Printf("Welcome to NomadCoin\n\n")
	fmt.Printf("Please use the follwing commands:\n\n")
	fmt.Printf("explorer:	Start the HTML Explorer\n")
	fmt.Printf("rest: 		Start the REST API (recommended)\n")

	// 메뉴얼을 출력했다면 프로그램을 종료. 종료코드 0의 경우 정상종료
	os.Exit(0)
}

func main() {
	// Args의 첫번째 원소는 프로그램 파일명,
	// 길이가 2가 되지 않으면 매개변수가 입력되지 X
	if len(os.Args) < 2 {
		usage()
	}

	// Flag package를 이용해서 매개변수 파싱
	rest := flag.NewFlagSet("rest", flag.ExitOnError)
	portflag := rest.Int("port", 4000, "Set the port of the server")

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		rest.Parse(os.Args[2:])
	default:
		usage()
	}

	if rest.Parsed() {
		fmt.Println(*portflag)
		fmt.Println("Start Server")
	}
}
