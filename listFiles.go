package main

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
)

var(
	output *os.File
	outputFlNm string
	exeFlNm string
)
/**
 * メイン関数。カレントディレクトリを取得し、出力用のファイルを生成する。
 */
func main() {
	//開始メッセージ出力
	fmt.Print("ファイル一覧の出力を開始します。")

	//カレントディレクトリ取得
	cur, _ := os.Getwd()

	var err error

	//exeのファイル名取得
	thi, err := os.Executable()
	dirNm := filepath.Dir(thi) + "\\"
	exeFlNm = strings.Replace(thi, dirNm , "", -1)

	//出力ファイルの生成
	outputFlNm = "list.txt"
	output, err = os.Create(cur + "/" + outputFlNm)
	if err != nil {
		fmt.Print("出力ファイルの生成時にエラーが発生しました。")
		return
	}
	defer output.Close()

	//カレントディレクトリをトラバース処理し、ファイル出力関数を呼び出す。
	filepath.Walk(cur, listFiles)

	//終了メッセージ出力
	fmt.Print("ファイル一覧の出力を完了しました。")
}

/**
 * ファイル出力関数。
 * cur:カレントディレクトリ,output:出力ファイル
 */
func listFiles(path string, file os.FileInfo, err error) error {
	if err != nil{
		return err
	}

	//取得したものがexeファイル、出力ファイルであった場合はスキップする
	if file.Name() == exeFlNm || file.Name() == outputFlNm{
		return nil
	}

	//取得したものがディレクトリであった場合、スキップする
	if file.IsDir(){
		return nil
	}
	//ファイル名、ファイルパスの取得
	fileName := file.Name()
	path = strings.Replace(path, fileName , "", -1)
	//ファイルへの書き込み
	output.WriteString(path + "\t" + fileName + "\r")

	return nil
}
