package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
)

//Data is
type Data struct {
	Bid      int      `json:"bid"`
	BlockCc  string   `json:"blockCc"`
	Bname    string   `json:"bname"`
	Bpic     string   `json:"bpic"`
	Cid      int      `json:"cid"`
	Cname    string   `json:"cname"`
	Files    []string `json:"files"`
	Finished bool     `json:"finished"`
	Len      int      `json:"len"`
	NextID   int      `json:"nextId"`
	Path     string   `json:"path"`
	PrevID   int      `json:"prevId"`
	Sl       struct {
		Md5 string `json:"md5"`
	} `json:"sl"`
}

func main() {
	ExampleScrape()

	// DownloadJpg()
}

//ExampleScrape is
func ExampleScrape() {
	imageRoot := "https://i.hamreus.com"
	// Request the HTML page.
	// res, err := http.Get("https://tw.manhuagui.com/comic/7620")

	res, err := http.Get("https://tw.manhuagui.com/comic/7620/354279.html")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("start")
	s := doc.Find("script").Each(func(i int, sf *goquery.Selection) {
		if strings.Contains(sf.Text(), "window") {
			fmt.Println(sf.Text())
			fmt.Println()
			fmt.Println()

			vm := otto.New()
			_, err = vm.Run(`
				var args = 'abc';
				var SMH = {};
				SMH.imgData = function(i){
					args = JSON.stringify(i);
					return {
						preInit: function() {
						}
					}
				}
			`)

			if err != nil {
				log.Fatal(err)
			}

			splicScript := `var LZString=(function(){var f=String.fromCharCode;var keyStrBase64="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";var baseReverseDic={};function getBaseValue(alphabet,character){if(!baseReverseDic[alphabet]){baseReverseDic[alphabet]={};for(var i=0;i<alphabet.length;i++){baseReverseDic[alphabet][alphabet.charAt(i)]=i}}return baseReverseDic[alphabet][character]}var LZString={decompressFromBase64:function(input){if(input==null)return"";if(input=="")return null;return LZString._0(input.length,32,function(index){return getBaseValue(keyStrBase64,input.charAt(index))})},_0:function(length,resetValue,getNextValue){var dictionary=[],next,enlargeIn=4,dictSize=4,numBits=3,entry="",result=[],i,w,bits,resb,maxpower,power,c,data={val:getNextValue(0),position:resetValue,index:1};for(i=0;i<3;i+=1){dictionary[i]=i}bits=0;maxpower=Math.pow(2,2);power=1;while(power!=maxpower){resb=data.val&data.position;data.position>>=1;if(data.position==0){data.position=resetValue;data.val=getNextValue(data.index++)}bits|=(resb>0?1:0)*power;power<<=1}switch(next=bits){case 0:bits=0;maxpower=Math.pow(2,8);power=1;while(power!=maxpower){resb=data.val&data.position;data.position>>=1;if(data.position==0){data.position=resetValue;data.val=getNextValue(data.index++)}bits|=(resb>0?1:0)*power;power<<=1}c=f(bits);break;case 1:bits=0;maxpower=Math.pow(2,16);power=1;while(power!=maxpower){resb=data.val&data.position;data.position>>=1;if(data.position==0){data.position=resetValue;data.val=getNextValue(data.index++)}bits|=(resb>0?1:0)*power;power<<=1}c=f(bits);break;case 2:return""}dictionary[3]=c;w=c;result.push(c);while(true){if(data.index>length){return""}bits=0;maxpower=Math.pow(2,numBits);power=1;while(power!=maxpower){resb=data.val&data.position;data.position>>=1;if(data.position==0){data.position=resetValue;data.val=getNextValue(data.index++)}bits|=(resb>0?1:0)*power;power<<=1}switch(c=bits){case 0:bits=0;maxpower=Math.pow(2,8);power=1;while(power!=maxpower){resb=data.val&data.position;data.position>>=1;if(data.position==0){data.position=resetValue;data.val=getNextValue(data.index++)}bits|=(resb>0?1:0)*power;power<<=1}dictionary[dictSize++]=f(bits);c=dictSize-1;enlargeIn--;break;case 1:bits=0;maxpower=Math.pow(2,16);power=1;while(power!=maxpower){resb=data.val&data.position;data.position>>=1;if(data.position==0){data.position=resetValue;data.val=getNextValue(data.index++)}bits|=(resb>0?1:0)*power;power<<=1}dictionary[dictSize++]=f(bits);c=dictSize-1;enlargeIn--;break;case 2:return result.join('')}if(enlargeIn==0){enlargeIn=Math.pow(2,numBits);numBits++}if(dictionary[c]){entry=dictionary[c]}else{if(c===dictSize){entry=w+w.charAt(0)}else{return null}}result.push(entry);dictionary[dictSize++]=w+entry.charAt(0);enlargeIn--;w=entry;if(enlargeIn==0){enlargeIn=Math.pow(2,numBits);numBits++}}}};return LZString})();String.prototype.splic=function(f){return LZString.decompressFromBase64(this).split(f)}`
			srcValue, scriptErr := vm.Run(splicScript)
			if scriptErr != nil {
				log.Fatal(scriptErr)
			}

			newScript := strings.Replace(sf.Text(), `window["\x65\x76\x61\x6c"]`, "", 1)
			fmt.Println()
			fmt.Println()
			fmt.Println(newScript)

			srcValue, scriptErr = vm.Run(newScript)
			if scriptErr != nil {
				log.Fatal(scriptErr)
			}

			fmt.Println()
			fmt.Println("srcValue===>")
			fmt.Println(srcValue)

			srcValue, scriptErr = vm.Run(`SMH.imgData({"bid":7620,"bname":"斗破苍穹","bpic":"7620.jpg","cid":354279,"cname":"第667话 风雷动（上）","files":["001.jpg.webp","002.jpg.webp","003.jpg.webp","004.jpg.webp","005.jpg.webp","006.jpg.webp","007.jpg.webp","008.jpg.webp"],"finished":false,"len":8,"path":"/ps3/d/斗破苍穹/第667话/","status":0,"block_cc":"","nextId":354398,"prevId":353411,"sl":{"md5":"xCB85s_1D8quSPYxiq-F6w"}}).preInit();`)
			if scriptErr != nil {
				log.Fatal(scriptErr)
			}
			// vm.Run("SMH.imgData.preInit()")

			// 2018/12/10 18:17:36 TypeError: 'splic' is not a function

			if value, err := vm.Get("args"); err == nil {
				fmt.Println("args結果===>")
				fmt.Println(value.IsString())
				fmt.Println(value)

				data := &Data{
					// Votes: &Votes{},
				}

				// argsValue := []byte(value.Export())
				err := json.Unmarshal([]byte(value.String()), data)
				fmt.Println(err)

				// s2, _2 := json.Marshal(data)
				// fmt.Println(_2)
				// fmt.Println(string(s2))

				for i := 0; i < data.Len; i++ {
					filename := data.Files[i]
					fullURL := imageRoot + data.Path + filename + "?cid=" + fmt.Sprint(data.Cid) + "&md5=" + data.Sl.Md5
					fmt.Println(fullURL)
					DownloadJpg(fullURL, strings.Replace(filename, `.webp`, "", 1), data.Cname)
					// do something
				}
				// goData, _ := value.Export()
			}
		}

	})
	fmt.Println(s)

	// script := s.Nodes[1].FirstChild.Data
	fmt.Println("script===>")
	// fmt.Println(script)
	// vm := otto.New()

	prefix := "https://tw.manhuagui.com"

	doc.Find("#chapter-list-1 ul li a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		fmt.Printf(s.Text())
		fmt.Println()
		href, hasvalue := s.Attr("href")
		fmt.Println(href, hasvalue)
		fmt.Println()

		url := prefix + href
		fmt.Println(url)

		// GetImage(url)

	})
}

//GetImage is
func GetImage(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc)

	fmt.Println("mangaFile個數")
	fmt.Println(doc.Find("#imgLoading"))

	fmt.Println(len(doc.Find("#imgLoading").Nodes))
	fmt.Println("mangaFile Src==>")
	fmt.Println(doc.Find("#mangaFile").Attr("src"))

}

//DownloadJpg is
func DownloadJpg(imageURL string, fileName string, folderName string) {

	// don't worry about errors
	// res, err := http.Get("https://tw.manhuagui.com/comic/7620/354279.html")
	// if err != nil {
	// 	log.Fatal(res)
	// 	log.Fatal(err)
	// }

	// http.SetCookie(ressss, ressss.Cookies)
	response, e := http.Get(imageURL)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create("/AAA/" + folderName + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success!")
}
