package main

import (
	"../bank"
	"fmt"
	"log"
	"os"
	"strconv"
	//"reflect"
	"strings"
)

func deposit(a *bank.Account, b int) {
	a.Balance += b
	fmt.Println(a.Balance)
}

func withdraw(a *bank.Account, b int) {
	a.Balance -= b
	fmt.Println(a.Balance)
}
func main() {
	//선택
	var selnum int

	// 구조체 슬라이스 생성
	var accounts []bank.Account

	//기존에 있던 파일의 정보들 불러오기
	lines, err := bank.FileRead("account.txt")

	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		splits := strings.Split(line, ",")
		name := splits[0]
		num := splits[1]
		balance, err := strconv.Atoi(splits[2])
		if err != nil {
			log.Fatal(err)
		}
		s := bank.Account{Num: num, Name: name, Balance: balance}
		accounts = append(accounts, s)
		//fmt.Println(name,num,balance)
	}
	fmt.Println(accounts)
	for {
		fmt.Println("----------------------------------------------------")
		fmt.Println("1.계좌생성 | 2.계좌목록 | 3. 예금 | 4.출금 | 5.종료")
		fmt.Println("-----------------------------------------------------")

		fmt.Print("선택 > ")

		fmt.Scanf("%d", &selnum)

		switch selnum {
		case 1:
			var num, name string
			var balance int
			fmt.Println("=========")
			fmt.Println("계좌생성")
			fmt.Println("=========")
			fmt.Printf("계좌번호(000-000): ")
			fmt.Scanf("%s", &num)

			if !strings.Contains(num, "-") {
				fmt.Println("계좌번호를 바르게 입력하세요")
				break
			}
			fmt.Printf("계좌주: ")
			fmt.Scanf("%s", &name)
			fmt.Printf("초기입금액 : ")
			fmt.Scanf("%d", &balance)

			a := bank.Account{Name: name, Num: num, Balance: balance}
			accounts = append(accounts, a)

		case 2:
			fmt.Println("=========")
			fmt.Println("계좌목록")
			fmt.Println("=========")

			for _, account := range accounts {
				fmt.Printf("%s\t\t%s\t\t%d\n", account.Name, account.Num, account.Balance)
			}

		case 3:
			var num string
			var bal int
			ok := false
			fmt.Println("=========")
			fmt.Println("예금")
			fmt.Println("=========")
			fmt.Printf("계좌번호(000-000): ")
			fmt.Scanf("%s", &num)
			fmt.Printf("예금액 : ")
			fmt.Scanf("%d", &bal)

			for i, account := range accounts {
				if account.Num == num {
					deposit(&accounts[i], bal)
					ok = true
				}
			}
			if ok {
				fmt.Println("결과: 예금이 성공하였습니다")
			} else {
				fmt.Println("계좌번호를 찾지 못하였습니다.")
			}

		case 4:
			var num string
			var bal int
			ok := false
			fmt.Println("=========")
			fmt.Println("출금")
			fmt.Println("=========")
			fmt.Printf("계좌번호(000-000): ")
			fmt.Scanf("%s", &num)
			for i, account := range accounts {
				if account.Num == num {
					fmt.Println("현재 잔고 : ", account.Balance)

					fmt.Printf("출금액 : ")
					fmt.Scanf("%d", &bal)
					if account.Balance < bal {
						fmt.Println("잔액이 부족합니다")
						ok=false
					} else {
						withdraw(&accounts[i], bal)
						ok=true
					}
				}
			}
			if !ok {
				fmt.Println("결과 : 출금에 실패하였습니다")
				return
			} else {
				fmt.Println("결과 : 출금을 성공했습니다")
			}

		case 5:
			os.Exit(0)
		default:
			fmt.Println("잘못 입력하셨습니다\n프로그램을 종료합니다.")
			os.Exit(0)
		}

	}

}
