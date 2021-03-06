package main

import (
	"fmt"
	"net/http"
	"time"
	"path/filepath"
	"os"
	"strings"
	inix "github.com/go-ini/ini"
	termbox "github.com/nsf/termbox-go"
	"strconv"
	"io/ioutil"
)
var p = fmt.Println
/* init 请按任意键继续*/
func init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	termbox.SetCursor(0, 0)
	termbox.HideCursor()
	//defer termbox.Close()

}
func main(){
	iniPath :=filepath.Join( getCurrentDirectory(),"config.ini")
	p("iniPath:",iniPath)
	if IsDirFileExist(iniPath) !=true {
		p(iniPath,"找不到配置文件，请检查!")
		pause()
		return
	}

	cfg,_ := inix.LoadSources(inix.LoadOptions{IgnoreInlineComment: true}, iniPath)
	url :=cfg.Section("URL").Key("url").String()

	cycles1 := cfg.Section("URL").Key("cycles").String()
	cycles,_ := strconv.Atoi(cycles1)

	sleepMillisecond1 :=cfg.Section("URL").Key("sleepMillisecond").String()
	sleepMillisecond,_ := strconv.Atoi(sleepMillisecond1)

	isPrint1 := cfg.Section("URL").Key("isPrint").String()
	isPrint,_ := strconv.ParseBool(isPrint1)

	isPrintwebBody1 := cfg.Section("URL").Key("isPrintwebBody").String()
	isPrintwebBody,_ := strconv.ParseBool(isPrintwebBody1)

	p("          -------------- config.ini by xiaohai 2018.10.12 Ver:0.2 --------------           ")
	p("url: ", url)
	p("cycles: ", cycles)
	p("sleepMillisecond: ", sleepMillisecond)
	p("isPrint: ", isPrint)
	p("isPrintwebBody: ", isPrintwebBody)
	p("           -------------- config.ini --------------         ")
	p("\r\n")
	p(time.Now().Format("2006-01-02 15:04:05.000000"),"[   Get.Url运行中...  ]:",url)
	sleepXs(url,cycles,sleepMillisecond,isPrint,isPrintwebBody)
	pause()
}
func get(url string,isPrint bool,isPrintwebBody bool){
	response,_:=http.Get(url)
	if isPrintwebBody {
		defer response.Body.Close()
		body,_:=ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}

if isPrint {
	if response.StatusCode == 200 {
		p(time.Now().Format("2006-01-02 15:04:05.000000"),"[Refresh_web OK]:",url)
	}else{
		p(time.Now().Format("2006-01-02 15:04:05.000000"),"[Refresh_web Error]:",url)
	}
    }
}


// 休眠
func sleepXs(url string,num int,sleepMillisecond int,isPrint bool,isPrintwebBody bool) {
	// time.Millisecond    表示1毫秒
	// 休眠100毫秒
	//time.Sleep(100 * time.Millisecond)
	i:=0
	t := time.Now()
	sleepSecondTimeX := time.Millisecond * time.Duration(sleepMillisecond)
	for{
		i++
		get(url,isPrint,isPrintwebBody)
		if i>=num{
			p("ForNum:",num," sleepMillisecond:",sleepSecondTimeX," Use:",time.Now().Sub(t).String())
			break
		}

		// 休眠1秒
		time.Sleep(sleepSecondTimeX)
	}
}



func pause() {
	fmt.Println("请按任意键继续...")
Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			//case termbox.KeyEsc:  // case termbox.KeyF1:
			break Loop
		}
	}
}

//判断文件或文件夹是否存在
func IsDirFileExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

/*获取程序运行路径*/
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		p(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
