package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

var cookie string

func init() {
	flag.StringVar(&cookie, "cookie", "", "cnblog of cookie")
}
func main() {
	flag.Parse() //暂停获取参数
	cookie = strings.Replace(cookie, "_semicolon_", ";", -1)
	cookie = strings.Replace(cookie, "_vertical_", "|", -1)
	cookie = strings.Replace(cookie, "_frontbracket_", "(", -1)
	cookie = strings.Replace(cookie, "_backbracket_", ")", -1)
	cookie = strings.Replace(cookie, "_space_", " ", -1)
	println(cookie)
	if len(cookie) > 0 {
		fmt.Printf("开始执行")
		getBlogList(1)
	} else {
		fmt.Printf("未输入正确的CNBLOGS_COOKIE，请到github项目/Settings/Secrets/Actions下通过New repository secret添加一个CNBLOGS_COOKIE并填入正确的cookie值")
	}
}
func geturl(pageno int) string {
	return fmt.Sprintf("https://i.cnblogs.com/api/posts/list?p=%s&cid=&tid=&t=1&cfg=0&search=&orderBy=&s=&scid=", strconv.Itoa(pageno))
}
func getBlogList(pageindex int) {
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
		if childval.IsPublished {
			childbody := getData(fmt.Sprintf("https://i.cnblogs.com/api/posts/%s", strconv.Itoa(childval.Id)))
			var jsconf blogbodyConf
			err := json.Unmarshal(childbody, &jsconf)
			if err != nil {
				fmt.Println("error:", err)
			}
			var tagbody = ""
			for _, tag := range jsconf.BlogPost.Tags {
				if tagbody != "" {
					tagbody = fmt.Sprintf("%s,\"%s\"", tagbody, tag)
				} else {
					tagbody = fmt.Sprintf("\"%s\"", tag)
				}
			}
			fmt.Printf("\r\n tagbody:%s", tagbody)
			var tagstr = fmt.Sprintf("[%s]", tagbody)
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
				fmt.Printf("\r\n fileName:%s", fileName)
				downloadImage(imgurl, strconv.Itoa(jsconf.BlogPost.Id), fileName)
				articleBody = strings.Replace(articleBody, imgurl, fmt.Sprintf("/cnblogs/%s/%s", strconv.Itoa(jsconf.BlogPost.Id), fileName), -1)
			}

			reg1, _ := regexp.Compile(`http://im.*.png`)
			imgurls1 := reg1.FindAllString(articleBody, -1)
			for _, imgurl := range imgurls1 {
				fileName := path.Base(imgurl)
				fmt.Printf("\r\n fileName:%s", fileName)
				downloadImage(imgurl, strconv.Itoa(jsconf.BlogPost.Id), fileName)
				articleBody = strings.Replace(articleBody, imgurl, fmt.Sprintf("/cnblogs/%s/%s", strconv.Itoa(jsconf.BlogPost.Id), fileName), -1)
			}

			downloadFile(strings.NewReader(articleBody), strconv.Itoa(jsconf.BlogPost.Id), fmt.Sprintf("%s.md", strconv.Itoa(jsconf.BlogPost.Id)))
		}
	}
	if conf.PageIndex > 0 && conf.PageIndex*conf.PageSize <= conf.PostsCount {
		getBlogList(conf.PageIndex + 1)
	}
	fmt.Println("执行完毕")
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
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
func downloadImage(imgurl string, rootpath string, fileName string) {
	filepath := fmt.Sprintf("../cnblogs/%s/%s", rootpath, fileName)
	isexists, err := exists(filepath)
	if isexists {
		return
	}
	res, err := http.Get(imgurl)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

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
	isexists, err := exists(filepath)
	if isexists {
		return
	}
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
	PageIndex  int `json:"pageIndex"`
	PageSize   int `json:"pageSize"`
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
	Id            int      `json:"id"`
	AutoDesc      string   `json:"autoDesc"`
	DatePublished string   `json:"datePublished"`
	PostBody      string   `json:"postBody"`
	Title         string   `json:"title"`
	Url           string   `json:"url"`
	Author        string   `json:"author"`
	Tags          []string `json:"tags"`
}
