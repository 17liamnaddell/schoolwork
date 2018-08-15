package main

import "encoding/json"
import "fmt"
import "io/ioutil"
import "net/http"
import "strconv"

type MyData struct {
	History struct {
		Observations []struct {
			Pressurei string `json:"pressurei"`
			Tempi     string `json:"tempi"`
			Icon      string `json:"icon"`
			Precipm   string `json:"precipm"`
		} `json:"observations"`
	} `json:"history"`
}

func main() {
	fmt.Println("| Date | Pressure | temp | cloudy? | Precipm |")
	base := "201711"
	var morning string
	var night string
	var n = 2
	fmt.Println(string(n))

	for i := 2; i <= 9; i++ {
		var si = strconv.Itoa(i)
		md := request(base + "0" + si)
		fmt.Println(base + "0" + si)
		var first = md.History.Observations[0]
		var last = md.History.Observations[len(md.History.Observations)-1]
		pi1 := first.Pressurei
		pi2 := last.Pressurei
		tempi1 := first.Tempi
		tempi2 := last.Tempi
		icon1 := first.Icon
		icon2 := last.Icon
		precipm1 := first.Precipm
		precipm2 := last.Precipm
		morning = morning + fmt.Sprintf("| %s | %s | %s | %s | %s |\n", base+"0"+si, pi1, tempi1, icon1, precipm1)
		night = night + fmt.Sprintf("| %s | %s | %s | %s | %s |\n", base+"0"+si, pi2, tempi2, icon2, precipm2)

	}
	fmt.Println(morning)

	fmt.Println("\n| Date | Pressurei | tempi | icon | Precipm |")
	fmt.Println(night)

}

//yyyymmdd
func request(date string) MyData {
	var res []byte
	resp, err := http.Get("http://api.wunderground.com/api/20a4a46f24b3cd03/history_" + date + "/q/NC/Charlotte.json")
	if err != nil {
		panic(err)
	}
	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
	var buf MyData
	err = json.Unmarshal(res, &buf)
	if err != nil {
		panic(err)
	}
	return buf
}
