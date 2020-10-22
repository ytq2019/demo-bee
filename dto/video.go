/*
 * @Time : 2020/2/22
 * @Author : tianqi.yu
 * @Software : GoLand
 * @Description :
 */
package dto

type KankanVideo struct {
	Title string `json:"title"`
	Img   string `json:"img"`
	Href  string `json:"link"`
	Hint  string `json:"hint"`
	Star  string `json:"star"`
	Id    string `json:"id"`
	S2    string `json:"s2"`
}
type VideoList struct {
	Dianying []*KankanVideo `json:"dianying"`
	Dianshi  []*KankanVideo `json:"dianshi"`
	Zongyi   []*KankanVideo `json:"zongyi"`
	Dongman  []*KankanVideo `json:"dongman"`
}

type VideoDetail struct {
	Nav      string `json:"nav"`
	Title    string `json:"title"`
	Star     string `json:"star"`
	Thumb    string `json:"thumb"`
	Director string `json:"director"`
	Year     string `json:"year"`
	Area     string `json:"area"`
	Type     string `json:"type"`
	Actor    string `json:"actor"`
	Desc     string `json:"desc"`
}

type M360Resp struct {
	Data struct {
		List     string `json:"list"`
		HaveMore int    `json:"have_more"`
		NextURL  string `json:"next_url"`
	} `json:"data"`
	Error int    `json:"error"`
	Msg   string `json:"msg"`
}
type ListResp struct {
	Data []*KankanVideo `json:"list"`
}
