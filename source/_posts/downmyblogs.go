package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"regexp"
	"path"
)
const cookie = "__gads=ID=bbfb10578bf6ebf7:T=1649728913:S=ALNI_MbOOVEJx5s2Puf6CEs5qYSkW_aHwg; UM_distinctid=1805aabf9bd238-08b22c569116a1-3a67551f-144000-1805aabf9be934; Hm_lpvt_5e692fc8cef6db021a22bd470d660e4c=1651548396; Hm_lvt_5e692fc8cef6db021a22bd470d660e4c=1651548396; .AspNetCore.Antiforgery.b8-pDmTq1XM=CfDJ8AuMt_3FvyxIgNOR82PHE4kGyR-4aSs739sPowluSIbWBXRwvByozasPNheZKXorGKFxlQ2XMwL2PoW9QSwXcJS2PX1srKH7Fa9naKg7dhqQywIeSl4dYqyikYKZH4e61r02vJbatsXG76kQEqh_vUo; .AspNetCore.Session=CfDJ8AuMt%2F3FvyxIgNOR82PHE4luRKm0AX8iUW7HmvQ8B%2BVY3t4ucprrn4Mu3jD4G%2B%2BufQxZZtdKQbvlaZnSDa6wkIP3oZiLhlEaGmF0MJPSm%2FjvekMt%2FlPeLQ5fSxmz5rtmzj7AGjgKu1bCViLfVGink8N4v4fsafdaldUsUF%2BkpcWz; __utma=2642927.1176019125.1649728912.1652923332.1652923332.1; __utmc=2642927; __utmz=2642927.1652923332.1.1.utmcsr=cn.bing.com|utmccn=(referral)|utmcmd=referral|utmcct=/; Hm_lvt_0daf1d1987de95558f12b56df149bfda=1653285692; Hm_lpvt_0daf1d1987de95558f12b56df149bfda=1653285692; Hm_lvt_851cdd44d7a836d43196b0bfa8c0c3bb=1654159502; Hm_lpvt_851cdd44d7a836d43196b0bfa8c0c3bb=1654475406; Hm_lvt_c897bd902d294bb3778d9ab55b85a256=1655199408; Hm_lpvt_c897bd902d294bb3778d9ab55b85a256=1655199408; _ga_4CQQXWHK3C=GS1.1.1655860922.2.0.1655860929.0; Hm_lvt_866c9be12d4a814454792b1fd0fed295=1655862584; _gid=GA1.2.635015422.1656895021; __gpi=UID=000004c28bf2c485:T=1649729335:RT=1657432736:S=ALNI_MY4bFzL-Q4DF7EivtB4voDkSrLb2w; .Cnblogs.AspNetCore.Cookies=CfDJ8EOBBtWq0dNFoDS-ZHPSe5251Pvp1fD4MFhyd_PD4tNcGD_u2h0PSkDcNGhyUMRd1zyFDyEZ6kSQ9PGFTgwOqqzt9s7WozoAgC3wf2pU7fkRY59RV9cCP6HrSQeE-jHTo-1AFJNcEE1on6pVxzRETNegcBcpYbBBK97Dyz0wwErnYlgrkbkK6jUUFD4ajAUAP7O9nhT6keL_3mmLaLd9VA3JJHk22rvNVoYAWJ0c_nQFq4ysvSLS4UnM3LpKqmeV8XdOj9PJ59-GX_gXq6IVIJJcyxRiQmBRSgItn6GDXgg5WW3skY2HcTYlOXQzDYg_DssEnqLp5ETJhK6n_xBYIkM7z3oZNp-94VnYZuNHSolvNGnjE4fCNIjKGUYRNu-p1QEJGpcCdNrct9oEKwNtNAC9juCaB1B-dZkIX9GDds7w8uAeJP-PP2hNGJRoXwwvqhSxwvMfnNPKuUUfHM4E35TDeh0Uc5iM0dR92oao4_IPxx_yZbPmYmOzJ0K9-hFGfcRvlm0yPt0nALLNA1fM1QRyfxna5zWXz1TJkLs0hVX3ETtjIZ_Bo02WlyldpKs2_Q; .CNBlogsCookie=00570C45AB9DECCA0F43D2D63ECBA80584565487C0150E1B16B249C3C4C77A4B092D76EED817FF9F26485C86C144A03FC85262706ED3A218614FA4A93EBD1AD55B64EF37F16DFBB40D9244F428F757D00C08BCFD; _ga_3Q0DVSGN10=GS1.1.1657457482.3.1.1657457559.57; _ga=GA1.2.1176019125.1649728912; Hm_lpvt_866c9be12d4a814454792b1fd0fed295=1657457926; _gat_gtag_UA_48445196_1=1; XSRF-TOKEN=CfDJ8EOBBtWq0dNFoDS-ZHPSe539x5kZvZAb6VuAhygBUDenTzI2W0r6bPhJYfj93Z48RwxDhshFuqCZvuBdHhOysmFnEyA-xIp8TNGoMG005rDOQ8D0Qnmb-hGFhsfpLxjdQwMc8rY2FeSxKHPNs-Rwe0HIcOQckNWgdxPJsge9n_LohuQJAVqduolVgAWYCcIeqw"

func main() {
    fmt.Printf("开始执行")
	getBlogList(1)
}
func geturl(pageno int) string{
    return fmt.Sprintf("https://i.cnblogs.com/api/posts/list?p=%s&cid=&tid=&t=1&cfg=0&search=&orderBy=&s=&scid=",strconv.Itoa(pageno))
}
func getBlogList(pageindex int){
	var urlstr = geturl(pageindex)

	recordbody := getData(urlstr)
	fmt.Printf("\r\n recordbody:%s \n", recordbody)

	var conf blogList
	err := json.Unmarshal(recordbody, &conf)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("\r\n PageIndex:%s，PageSize:%s，PostsCount:%s \n", strconv.Itoa(conf.PageIndex), strconv.Itoa(conf.PageSize), strconv.Itoa(conf.PostsCount))
	for _, childval := range conf.PostList {
		if(childval.IsPublished){
			childbody := getData(fmt.Sprintf("https://i.cnblogs.com/api/posts/%s", strconv.Itoa(childval.Id)))
			fmt.Printf("childbody:%s \n", childbody)
			var jsconf blogbodyConf
			err := json.Unmarshal(childbody, &jsconf)
			if err != nil {
				fmt.Println("error:", err)
			}
			var tagbody = ""
			for _, tag := range jsconf.BlogPost.Tags {
				if(tagbody!=""){
					tagbody = fmt.Sprintf("%s,\"%s\"",tagbody,tag)
				}else{
					tagbody = fmt.Sprintf("\"%s\"",tag)
				}
			}
			var tagstr = fmt.Sprintf("[%s]",tagbody)
			var articleBody = fmt.Sprintf("---\r\ntitle: %s\r\ndate: %s\r\nauthor: %s\r\ntags: %s\r\n---\r\n%s",
			    jsconf.BlogPost.Title,
				jsconf.BlogPost.DatePublished,
				jsconf.BlogPost.Author,
				tagstr,
			    string(jsconf.BlogPost.PostBody))
			//添加文章信息

			reg, _ := regexp.Compile(`https://im.*.png`)
			imgurls := reg.FindAllString(articleBody, -1)
			for _, imgurl := range imgurls {
				fileName := path.Base(imgurl)
				downloadImage(imgurl,strconv.Itoa(jsconf.BlogPost.Id),fileName)
				articleBody = strings.Replace(articleBody, imgurl, fmt.Sprintf("/cnblogs/%s/%s",strconv.Itoa(jsconf.BlogPost.Id),fileName), -1)
			}
			fmt.Printf("articleBody:%s \n", articleBody)

			downloadFile(strings.NewReader(articleBody), strconv.Itoa(jsconf.BlogPost.Id), fmt.Sprintf("%s.md",  strconv.Itoa(jsconf.BlogPost.Id)))
		}
	}
    if(conf.PageIndex>0 && conf.PageIndex*conf.PageSize<=conf.PostsCount){
		getBlogList(conf.PageIndex+1)
	}
	fmt.Println("执行完毕")
}
func getData(urlstr string) []byte {
	client := &http.Client{}
	fmt.Printf("\r\n urlstr:%s \n", urlstr)
	req, _ := http.NewRequest("GET", urlstr, nil)
	req.Header.Add("cookie", cookie)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
func downloadImage(imgurl string, rootpath string, fileName string){
	filepath := fmt.Sprintf("../cnblogs/%s/%s", rootpath, fileName)
	res, err := http.Get(imgurl)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32 * 1024)

    if _, err := os.Stat(fmt.Sprintf("../cnblogs/%s", rootpath)); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.MkdirAll(fmt.Sprintf("../cnblogs/%s", rootpath), 0777) //0777也可以os.ModePerm
		os.Chmod(fmt.Sprintf("../cnblogs/%s", rootpath), 0777)
	}
	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
}
func downloadFile(body io.Reader, rootpath string, name string) {
    filepath := fmt.Sprintf("./cnblogs/%s", name)
	// Create output file
	if rootpath != "" {
		if _, err := os.Stat("./cnblogs"); os.IsNotExist(err) {
			// 必须分成两步：先创建文件夹、再修改权限
			os.MkdirAll("./cnblogs", 0777) //0777也可以os.ModePerm
			os.Chmod("./cnblogs", 0777)
		}
	}
	out, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	// copy stream
	_, err = io.Copy(out, body)
	if err != nil {
		panic(err)
	}
}

type blogList struct {
	PageIndex int `json:"pageIndex"`
	PageSize int `json:"pageSize"`
	PostsCount int `json:"postsCount"`

	PostList []blogbodymsg `json:"postList"`
}
type blogbodymsg struct {
	Id int `json:"id"`

	DatePublished string `json:"datePublished"`

	DateUpdated string `json:"dateUpdated"`

	Title string `json:"title"`

	IsPublished bool `json:"isPublished"`
}

type blogbodyConf struct {
	BlogPost blogPostEntity `json:"blogPost"`
}
type blogPostEntity struct {
	Id int `json:"id"`
	AutoDesc string `json:"autoDesc"`
	DatePublished string `json:"datePublished"`
	PostBody string `json:"postBody"`
	Title string `json:"title"`
	Url string `json:"url"`
	Author string `json:"author"`
	Tags []string `json:"tags"` 
}