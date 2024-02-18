## 基本構文

### :=  // 変数宣言

func main() {
    var a int = 1
    b := 2
    fmt.Println(a, b)
}

### &  // アドレス演算子（ポインタ）の作成
func main() {
    a := 1
    b := &a
    fmt.Println(a, b)
}

