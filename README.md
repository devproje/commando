# Commando
- Go 언어로 만든 간단한 Argument 파싱 & 핸들링 라이브러리 입니다!

## Milestone
- Short option 기능 (ex: -a, -b 등등)

## How to use
- 더 자세한 예제는 [example](https://github.com/devproje/commando/tree/master/example)로 가시면 확인 하실 수 있어요!

### General
- code
```go
package main

import (
	"fmt"
	"github.com/devproje/commando"
	"github.com/devproje/commando/option"
	"os"
)

func main() {
	// 파일 이름은 제외한 상태로 arguments 주입
	command := commando.NewCommando(os.Args[1:])

	// 입력 argument가 없을 경우
	command.Root("test", "테스트 명령어 입니다!", func(n *commando.Node) error {
		fmt.Println("Hello, World!")
		return nil
	})

	// 입력 argument가 있을 경우
	command.Root("print", "테스트 명령어 입니다!", func(n *commando.Node) error {
		name, err := option.ParseString(*n.MustGetOpt("name"), n)
		if err != nil {
			return err
		}
		
		fmt.Printf("입력받은 이름: %s\n", name)
		return nil
	})

	// 입력받은 arguments를 파싱 후 핸들링
	err := command.Execute()
	if err != nil {
		panic(err)
    }
}
```

- in terminal
```bash
~$ ./sample test
Hello, World!
~$ ./sample print --name "Eungyo Lee"
입력받은 이름: Eungyo Lee
```

### Complex
- code
```go
package main

import (
	"fmt"
	"github.com/devproje/commando"
	"github.com/devproje/commando/option"
	"github.com/devproje/commando/types"
	"os"
)

func main() {
	// 파일 이름은 제외한 상태로 arguments 주입
	command := commando.NewCommando(os.Args[1:])

	command.ComplexRoot("test", "테스트 명령어 입니다!", []commando.Node{
		command.Then("print", "Hello, World를 출력 합니다", func(n *commando.Node) error {
			fmt.Println("Hello, World!")
			return nil
		}),
		command.Then("sum", "두 수의 합을 구합니다", func(n *commando.Node) error {
			var a, b int64
			var err error

			a, err = option.ParseInt(*n.MustGetOpt("a"), n)
			if err != nil {
				return err
			}

			b, err = option.ParseInt(*n.MustGetOpt("b"), n)
			if err != nil {
				return err
			}
			
			fmt.Printf("%d + %d = %d\n", a, b, a + b)
			return nil
		}, types.OptionData{
			Name: "a",
			Desc: "첫번째 숫자",
			Type: types.INTEGER,
		}, types.OptionData{
			Name: "b",
			Desc: "두번째 숫자",
			Type: types.INTEGER,
		}),
	})

	// 입력받은 arguments를 파싱 후 핸들링
	err := command.Execute()
	if err != nil {
		panic(err)
	}
}
```

- in terminal
```bash
~$ ./sample test print
Hello, World!
~$ ./sample test sum --a 10 --b 20
10 + 20 = 30
```
