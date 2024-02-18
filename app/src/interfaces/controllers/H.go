// この構造体は、WebアプリケーションのAPIレスポンスを標準化するために使用されます。
// 具体的には、JSON形式のレスポンスボディを生成する際に利用されることが想定されています。
package controllers

type H struct {
    Message string `json:"message"`
    Data interface{} `json:"data"`
}

func NewH(message string, data interface{}) *H {
    H := new(H)
    H.Message = message
    H.Data = data
    return H
}