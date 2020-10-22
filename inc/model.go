package inc

import (
	"demo-bee/dto"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

var jar, _ = cookiejar.New(nil)
var hClient = &http.Client{
	Jar: jar,
}

func GetRankTopList(typeName string, pageNo int) []*dto.KankanVideo {
	url := fmt.Sprintf("http://www.360kan.com/%s/list.php?cat=all&year=all&area=all&act=all&rank=%s&pageno=%d", typeName, "rankhot", pageNo)
	//client := &http.Client{}
	//jar, err := cookiejar.New(nil)
	//client := &http.Client{
	//	Jar: jar,
	//}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.116 Mobile Safari/537.36")
	req.Header.Add("Referer", "http://www.baidu.com")
	req.Header.Add("Cookie", "__guid=121874957.2736009159732329500.1575374130551.6943; bottom_uid=4604b678b7d1b7525d0a5826f0ac84d0; loc_code=10000000; __huid=11NTGG3eHIVfK9M37A4wG3LJr8IalWVBMlLd4QDNGpAjI%3D; sv_rel=a; sv_hot=a; sv_gus=a; sv_web=b; sv_trivia=a; sv_list=a; sv_detail=a; sv_detail_v3=a; search_page=b; dy_rel=b; tv_rel=b; zy_rel=b; dm_rel=b; __pcclient1=1; __count=22; __bottomad_key=11%7C1582359720618")
	res, err := hClient.Do(req)
	defer res.Body.Close()

	x, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		panic(err)
	}
	var videos []*dto.KankanVideo

	x.Find(".s-tab-main ul li").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Find(".js-tongjic").Attr("href")
		id, _ := s.Find(".js-tongjic").Attr("href")
		img, _ := s.Find("img").Attr("src")
		title := s.Find(".s1").Text()
		hint := s.Find(".hint").Text()
		star := s.Find(".star").Text()
		tmpVideo := &dto.KankanVideo{
			Title: title,
			Img:   img,
			Href:  href,
			Hint:  hint,
			Star:  star,
			Id:    id,
			S2:    "",
		}
		videos = append(videos, tmpVideo)
	})
	if len(videos) == 0 {
		log.Println("error")
		log.Println(x.Html())
	} else {
		log.Println("success")
	}
	return videos
}

func GetVideoDetail(vUrl, vType string) *dto.VideoDetail {
	url := fmt.Sprintf("http://www.360kan.com/%s", vUrl)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Add("Referer", "http://baidu.com")
	res, err := client.Do(req)
	defer res.Body.Close()

	x, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var tmpVideo = new(dto.VideoDetail)
	if vType == "dianshi" || vType == "dianying" {
		x.Find(".c-top-main-wrap").Each(func(i int, s *goquery.Selection) {
			nav := x.Find("").Text()
			title := x.Find("h1").Text()
			star := x.Find(".item-actor").Text()
			thumb, _ := x.Find(".s-cover-img img").Attr("src")
			director := x.Find("#js-desc-switch .item-wrap .item").Eq(3).Find(".name").Text()
			year := x.Find("#js-desc-switch .item-wrap .item").Eq(1).Text()
			area := x.Find("#js-desc-switch .item-wrap .item").Eq(2).Text()
			vtype := x.Find(".tag").Text()
			actor := x.Find("#js-desc-switch .item").Eq(3).Find("a").Text()
			desc := x.Find(".js-close-wrap").Text()
			tmpVideo = &dto.VideoDetail{
				Nav:      nav,
				Title:    title,
				Star:     strings.Replace(strings.Replace(strings.Replace(star, " ", "", -1), "\n", "", -1), "/", "", -1),
				Thumb:    thumb,
				Director: director,
				Year:     year,
				Area:     area,
				Type:     vtype,
				Actor:    actor,
				Desc:     strings.Replace(desc, "收起<<", "", -1),
			}
		})

	} else {

	}
	return tmpVideo
}

func GetRankTopListByMobile(typeName string, pageNo int) []*dto.KankanVideo {
	url := fmt.Sprintf("http://m.360kan.com/list/%sData?pageno=%d", typeName, pageNo)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.116 Mobile Safari/537.36")
	req.Header.Add("Referer", "http://www.baidu.com")
	res, err := hClient.Do(req)
	defer res.Body.Close()
	m360Resp := new(dto.M360Resp)
	bytes, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bytes, m360Resp)
	x, err := goquery.NewDocumentFromReader(strings.NewReader(m360Resp.Data.List))
	//x, err := goquery.NewDocumentFromReader(m360Resp.Data.List)
	if err != nil {
		log.Fatal(err)
	}
	var videos []*dto.KankanVideo

	x.Find(".item").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Find("a").Attr("href")
		id, _ := s.Find("a").Attr("href")
		img, _ := s.Find(".img img").Attr("src")
		title := s.Find(".title").Text()
		hint := s.Find(".score em").Text()
		star := s.Find(".info p").Eq(1).Text()
		tmpVideo := &dto.KankanVideo{
			Title: title,
			Img:   img,
			Href:  href,
			Hint:  hint,
			Star:  strings.Replace(star, "\n", "", -1),
			Id:    id,
			S2:    "",
		}
		videos = append(videos, tmpVideo)
	})
	if len(videos) == 0 {
		log.Println("error")
		log.Println(x.Html())
	} else {
		log.Println("success")
	}
	return videos
}

//电视
//function dianshi_url($url)
//{
//$page = 'http://m.360kan.com' . $url;
//$data = QueryList::Query($page,
//array("site" => array(".wrap option", "data-site"))
//)->data;
//if (count($data) == 0) {
//$data = QueryList::Query($page,
//array("site" => array(".wrap span", "class"))
//)->data;
//$site = substr($data[0]['site'], 14);
//$data[0]['site'] = $site;
//}
//$sites = [];
//foreach ($data as $k => $v) {
//$siteList = site();
//array_push($sites, $siteList[$v['site']]);
//}
//
//return $sites;
//}
//电影
//function caiji_url($url)
//{
//$url = "http://www.360kan.com" . $url;
//$data = QueryList::Query($url, array("link" => array(".top-list-zd:eq(1) a", "href"), "title" => array(".top-list-zd:eq(1) a", "text")))->data;
//if (empty($data)) {
//$data = QueryList::Query($url, array("link" => array(".top-list-zd a", "href"), "title" => array(".top-list-zd a", "text")))->data;
//}
//
//return $data;
//}
//
//function juji_url($url, $site)
//{
//$id = explode("/", str_substr("/", ".", $url));
//$url = "http://www.360kan.com/cover/switchsite?site=" . $site["0"] . "&id=" . $id["1"] . "&category=2";
//$html = file_get_contents($url);
//$html = json_decode($html, true);
//$html = $html["data"];
//if (empty($html)) {
//$url = "http://www.360kan.com/cover/switchsite?site=leshi&id=" . $id["1"] . "&category=2";
//$html = file_get_contents($url);
//$html = json_decode($html, true);
//$html = $html["data"];
//}
//$data = QueryList::Query($html, array("link" => array(".num-tab-main:eq(1) a", "href")))->data;
//if (empty($data)) {
//$data = QueryList::Query($html, array("link" => array(".js-tab a", "href"), "jishu" => array(".js-tab a", "text"), "yugao" => array(".ico-yugao", "text")))->data;
//}
//return $data;
//}
//动漫
//function dongman_url($url, $site)
//{
//$id = explode("/", str_substr("/", ".", $url));
//if ($site["0"] == "levp") {
//$site["0"] = "leshi";
//}
//$url = "http://www.360kan.com/cover/switchsite?site=" . $site["0"] . "&id=" . $id["1"] . "&category=4";
//$html = file_get_contents($url);
//$html = json_decode($html, true);
//$html = $html["data"];
//$data = QueryList::Query($html, array("link" => array(".num-tab-main\t a", "href"), "jishu" => array(".num-tab-main:gt(0) a", "text")))->data;
//$enddata = end($data);
//if ($enddata["link"] != "#") {
//$data = QueryList::Query($html, array("link" => array(".num-tab-main a", "href")))->data;
//} else {
//$data = QueryList::Query($html, array("link" => array(".num-tab-main:gt(0) a", "href"), "jishu" => array(".num-tab-main:gt(0) a", "text")))->data;
//}
//return $data;
//}
//综艺
//function zongyi_url($url)
//{
//$url = "http://www.360kan.com" . $url;
//$data = QueryList::Query($url, array("link" => array(".zd-down a", "href"), "title" => array(".zd-down a", "text")))->data;
//if (empty($data)) {
//$data = QueryList::Query($url, array("link" => array(".ea-site", "href"), "title" => array("#js-siteact .ea-site", "text")))->data;
//}
//return $data;
//}
