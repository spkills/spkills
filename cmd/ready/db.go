package main

import (
	"bytes"
	"flag"
	"os"
	"os/exec"

	"github.com/BurntSushi/toml"
	"github.com/spkills/spkills/config"
)

type SqlBoilerConfig struct {
	Output  string `toml:"output"`
	Pkgname string `toml:"pkgname"`
	NoTests bool   `toml:"no-tests"`
	Mysql   struct {
		DbName  string `toml:"dbname"`
		Host    string `toml:"host"`
		Port    string `toml:"port"`
		User    string `toml:"user"`
		Pass    string `toml:"pass"`
		SslMode bool   `toml:"sslmode"`
	} `toml:"mysql"`
}

const (
	SqlBoilerConfigName = "sqlboiler.toml"
)

func main() {
	cli := &Ready{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}

func (r *Ready) Run(args []string) int {

	// オプション引数のパース
	var infile string
	flags := flag.NewFlagSet("awesome-cli", flag.ContinueOnError)
	flags.SetOutput(r.errStream)
	flags.StringVar(&infile, "file", "db.toml", "Path to db.toml file")

	var conf SqlBoilerConfig
	_, err := toml.DecodeFile(infile, &conf)
	if err != nil {
		return ExitCodeParseFlagError
	}

	// アプリケーションの設定からDBの情報を補完
	var appConf config.Config
	_, err = toml.DecodeFile("config/config.toml", &appConf)
	if err != nil {
		return ExitCodeParseFlagError
	}
	conf.Mysql.DbName = appConf.Database.DbName
	conf.Mysql.Host = appConf.Database.Server
	conf.Mysql.Port = appConf.Database.Port
	conf.Mysql.User = appConf.Database.User
	conf.Mysql.Pass = appConf.Database.Password

	// sqlboilder.tomlという名前で出力
	var b bytes.Buffer
	enc := toml.NewEncoder(&b)
	err = enc.Encode(conf)
	if err != nil {
		return ExitCodeParseFlagError
	}
	file, err := os.Create(SqlBoilerConfigName)
	if err != nil {
		return ExitCodeNG
	}
	file.Write(b.Bytes())
	file.Close()

	// generateする
	err = exec.Command("sqlboiler", "mysql").Run()
	if err != nil {
		return ExitCodeNG
	}

	// ファイルを消す
	err = os.Remove(SqlBoilerConfigName)
	if err != nil {
		return ExitCodeNG
	}
	return ExitCodeOK
}
