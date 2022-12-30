package main

import "fmt"

// 错误处理(异常处理)
func main() {
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100 / 10 = ", result)
	}

	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}

type DivideError struct {
	dividee int
	divider int
}

// 实现error接口
func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	dividee: %d
	divider: 0
	`
	return fmt.Sprintf(strFormat, de.dividee)
}

// int类型除法
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}
