package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
)

/**
 * メイン関数。カレントディレクトリを取得し、出力用のファイルを生成する。
 */
func main() {
	//開始メッセージ出力
	fmt.Print("ファイル一覧の出力を開始します。")

	//カレントディレクトリ取得
	cur, _ := os.Getwd()

	//list.txtの生成
	output, err := os.Create(cur + "/list.txt")
	if err != nil {
		fmt.Print("カレントディレクトリ取得、またはlist.txtの生成時にエラーが発生しました。")
		return
	}
	defer output.Close()

	//ファイル一覧出力関数
	showFiles(cur, output)

	//終了メッセージ出力
	fmt.Print("ファイル一覧の出力を完了しました。")
}

/**
 * ファイル一覧出力関数。
 * cur:カレントディレクトリ,output:出力ファイル
 */
func showFiles(cur string, output *os.File) {
	//ディレクトリの中身を読み込み開始
	files, _ := ioutil.ReadDir(cur)

	//中身一覧をforループ
	for _, file := range files {
		//取得したものがディレクトリであった場合、再帰させる
		if file.IsDir(){
			absolutePath, _ := filepath.Abs(file.Name())
			showFiles(absolutePath, output)
			continue
		}
		//取得したものがexeファイル、出力ファイルであった場合はスキップする
		if file.Name() == "listFiles.exe" || file.Name() == "list.txt"{
			continue
		}

		//ファイル名、ファイルパスの取得
		fileName := file.Name()
		filePath, _ := filepath.Abs(fileName)

		//ファイルへの書き込み
		output.WriteString(filePath + "\t" + fileName + "\r")
	}
}
