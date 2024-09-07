package main

import (
	"AndroidToolServer-Go/model"
	"AndroidToolServer-Go/roof/db"
	"context"
	"encoding/xml"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"github.com/chromedp/chromedp"
	"github.com/wangshizebin/jiebago"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

var cutter *jiebago.JieBaGo
var stopWords []string

func init() {
	cutter = jiebago.NewJieBaGo()
	stopWords = append(stopWords, "的")
}

type SubFile struct {
	Name         string `xml:"name,attr"`
	TxtName      string `xml:"txtName,attr"`
	PartNo       string `xml:"partNo,attr"`
	ChapterNo    string `xml:"chapterNo,attr"`
	OriginalNo   string `xml:"originalNo,attr"`
	OriginalName string `xml:"originalName,attr"`
	TopicID      string `xml:"topicId,attr"`
}

type File struct {
	Name         string    `xml:"name,attr"`
	TxtName      string    `xml:"txtName,attr"`
	PartNo       string    `xml:"partNo,attr"`
	ChapterNo    string    `xml:"chapterNo,attr"`
	OriginalNo   string    `xml:"originalNo,attr"`
	OriginalName string    `xml:"originalName,attr"`
	TopicID      string    `xml:"topicId,attr"`
	SubFiles     []SubFile `xml:"file"`
}

type Zip struct {
	Name      string `xml:"name,attr"`
	NodeID    string `xml:"nodeId,attr"`
	PartNo    string `xml:"partNo,attr"`
	ChapterNo string `xml:"chapterNo,attr"`
	Files     []File `xml:"file"`
}

const FILE_PATH = "D:/Go/troubles2.xml"

func main() {
	//writeToXml()

	//fmt.Println(string(body))

	//getList()

	/*url := "https://support.huawei.com/supportgateway/view/v1/enterprise/doc/main-content?nid=EDOC1000015874&partNo=j006"
	parse(url)*/

	parseXml()
}

func parseXml() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	// 打开 XML 文件
	file, err := os.Open(FILE_PATH)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 定义 XML 解析器
	decoder := xml.NewDecoder(file)

	// 解析 XML 文件
	var zip Zip
	err = decoder.Decode(&zip)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
		return
	}

	// 输出解析结果
	fmt.Println("Zip Name:", zip.Name)
	fmt.Println("Node ID:", zip.NodeID)
	fmt.Println("Part No:", zip.PartNo)
	fmt.Println("Chapter No:", zip.ChapterNo)
	cnt := 1
	for _, file := range zip.Files {
		fmt.Println("File Name:", file.Name)
		fmt.Println("Text Name:", file.TxtName)
		decimal, _ := strconv.ParseInt(file.TxtName[2:], 16, 64)
		fmt.Println("Alarm ID:", decimal)
		fmt.Println("Part No:", file.PartNo)
		fmt.Println("Chapter No:", file.ChapterNo)
		fmt.Println("Original No:", file.OriginalNo)
		fmt.Println("Original Name:", file.OriginalName)
		fmt.Println("Topic ID:", file.TopicID)

		for _, subFile := range file.SubFiles {
			//start := strings.Index(subFile.Name, "t")
			//end := strings.Index(subFile.Name, "（")

			fmt.Println("  Sub File Name:", subFile.Name)
			//fmt.Println("  Text Name:", subFile.TxtName)
			fmt.Println("  Part No:", subFile.PartNo)

			baseUrl := "https://support.huawei.com/supportgateway/view/v1/enterprise/doc/main-content?nid=EDOC1000015874&partNo="
			//parse(baseUrl)
			//resp, _ := http.Get(baseUrl + subFile.PartNo)

			tx := parse(baseUrl+subFile.PartNo, subFile.TopicID)
			fmt.Println("第", cnt, "个故障正在写入数据库")
			if tx.Error != nil {
				fmt.Println("第", cnt, "个故障写入数据库完成!")
				cnt++
				//wg.Done()
			}

			fmt.Println()

			/*go func() {
				baseUrl := "https://support.huawei.com/supportgateway/view/v1/enterprise/doc/main-content?nid=EDOC1000015874&partNo="
				//parse(baseUrl)
				//resp, _ := http.Get(baseUrl + subFile.PartNo)
				tx := parse(baseUrl+subFile.PartNo, subFile.TopicID)
				fmt.Println("第", i, "个故障正在写入数据库")
				if tx.Error != nil {
					fmt.Println("第", i, "个故障写入数据库完成!")
					wg.Done()
				}

				fmt.Println()
			}()*/
			//fmt.Println("  Chapter No:", subFile.ChapterNo)
			//fmt.Println("  Original No:", subFile.OriginalNo)
			//fmt.Println("  Original Name:", subFile.OriginalName)
			fmt.Println("  Topic ID:", subFile.TopicID)
			fmt.Println()
		}

		fmt.Println()
	}

	wg.Wait()
}

func contains(target string, slice []string) bool {
	return strings.Contains(strings.Join(slice, ""), target)
}

// 分词
func testCutWords(str string) []string {
	var words []string

	words = cutter.ExtractKeywords(str, 10)
	return words
}

func parse(url string, topicId string) (tx *gorm.DB) {
	//textTitles := []string{"告警解释", "对系统的影响"}
	resp, _ := http.Get(url)
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	falure := &model.DgsGeneralFailure{}
	idTmp, _ := strconv.ParseInt(topicId, 16, 32)
	id := int32(idTmp)
	falure.ID = &id
	//告警标题
	title := doc.Find("h1").Text()
	//start := strings.Index(title, "：")
	//end := strings.Index(title, "（")
	//title = title[start+3 : end]
	fmt.Println("告警标题:", title)
	falure.AlarmTitle = title
	//关键字
	fmt.Println("关键字:")
	words := testCutWords(title)
	str := strings.Join(words, "|")
	fmt.Println(str)
	falure.AlarmKeyword = &str
	doc1, _ := doc.Find("div").Html()

	//网站源代码 alarm_handle
	falure.AlarmDetail = &doc1
	fmt.Println(len(doc1))
	fmt.Println()

	doc.Find(".idp-ltr-html-section").Each(func(i int, selection *goquery.Selection) {
		alarmTitle := selection.Find("h2").Text()
		fmt.Println(alarmTitle)
		fmt.Println("=====================================")
		//fmt.Println(selection.Html())
		var resons []string
		if alarmTitle == "对系统的影响" {
			influence := selection.Find("p").Text()
			falure.AlarmSystemInfluence = &influence
			fmt.Println(influence)
		}
		if alarmTitle == "可能原因" {
			selection.Find("li").Each(func(i int, selection *goquery.Selection) {
				resons = append(resons, selection.Text())
			})
			fmt.Println(resons)
			falure.AlarmPossibleReasons = resons
		}
		if alarmTitle == "处理步骤" {
			//解决步骤带标题
			doc2, _ := selection.Html()
			falure.AlarmHandle = doc2
			//fmt.Println(doc2)
			//获取解决步骤内容
			//var doc2 strings.Builder
			//doc2.WriteString("<ol>")
			//doc3, _ := selection.Find("ol").Html()
			//doc2.WriteString(doc3)
			//doc2.WriteString("</ol>")
			//fmt.Println(doc2.String())
		}
	})

	return db.PostgreSQLOrm.DB().Create(falure)
}

func writeToXml() {
	resp, _ := http.Get("https://support.huawei.com/supportgateway/view/v1/enterprise/doc/catalogue?nid=EDOC1000015874")
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	err := os.WriteFile(FILE_PATH, body, 0644)
	if err == nil {
		fmt.Println("写入成功！")
	}
}

func getList() {
	// 创建新的 chromedp 上下文
	ctx, cancel := chromedp.NewContext(context.Background(),
		chromedp.WithLogf(func(format string, args ...interface{}) {
			fmt.Printf(format, args...)
		}),
	)
	defer cancel()

	// 导航到目标网页
	var htmlContent string
	err := chromedp.Run(ctx, chromedp.Navigate("https://support.huawei.com/enterprise/zh/doc/EDOC1000015874/9586e444"))
	if err != nil {
		fmt.Println("Error navigating to webpage:", err)
		return
	}

	// 等待页面完全加载
	err = chromedp.Run(ctx, chromedp.WaitReady("body"))
	if err != nil {
		fmt.Println("Error waiting for page to load:", err)
		return
	}

	// 获取页面 HTML 内容
	err = chromedp.Run(ctx, chromedp.OuterHTML("html", &htmlContent))
	if err != nil {
		fmt.Println("Error getting page HTML:", err)
		return
	}

	fmt.Println(htmlContent)
	// 使用 htmlquery 库解析 HTML 内容
	_, err = htmlquery.Parse(strings.NewReader(htmlContent))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

}
