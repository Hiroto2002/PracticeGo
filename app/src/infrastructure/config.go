// この層は、アプリケーションの技術的な基盤を構成する要素に関連するコードを含んでいます。
// データベース接続の管理、外部サービスとの通信、ファイルシステムの操作など、アプリケーションの核となる機能をサポートする低レベルの詳細を扱います。
// infrastructureパッケージは、アプリケーションの他の部分（ビジネスロジックやUI）から技術的詳細を抽象化し、分離することで、コードの再利用性とメンテナンス性を高める役割を持ちます。
package infrastructure

// struct:構造体、Javascriptでいう型エイリアス
type Config struct {
    DB struct {
        Production struct {
            Host string
            Username string
            Password string
            DBName string
			Port string
        }
        Test struct {
            Host string
            Username string
            Password string
            DBName string
			Port string
        }
    }
    Routing struct {
        Port string
    }
}

// インスタンス生成
func NewConfig() *Config {

    c := new(Config)

    c.DB.Production.Host = "postgres"
    c.DB.Production.Username = "postgres"
    c.DB.Production.Password = "postgres"
    c.DB.Production.DBName = "practicego"
	c.DB.Production.Port = "5432"


    // c.DB.Test.Host = "localhost"
    // c.DB.Test.Username = "username"
    // c.DB.Test.Password = "password"
    // c.DB.Test.DBName = "db_name_test"
    
    c.Routing.Port = ":8080"

    return c
}
