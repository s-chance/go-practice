package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type DictRequest struct {
}

type DictResponse struct {
	NormalizedSource string `json:"normalizedSource"`
	DisplaySource    string `json:"displaySource"`
	Translations     []struct {
		NormalizedTarget string  `json:"normalizedTarget"`
		DisplayTarget    string  `json:"displayTarget"`
		PosTag           string  `json:"posTag"`
		Confidence       float64 `json:"confidence"`
		PrefixWord       string  `json:"prefixWord"`
		BackTranslations []struct {
			NormalizedText string `json:"normalizedText"`
			DisplayText    string `json:"displayText"`
			NumExamples    int    `json:"numExamples"`
			FrequencyCount int    `json:"frequencyCount"`
		} `json:"backTranslations"`
		Transliteration string `json:"transliteration"`
	} `json:"translations"`
}

func main() {
	client := &http.Client{}
	var data = strings.NewReader(`&from=en&to=zh-Hans&text=sky&token=q9B7xegeNgTmtJAWUwFzNn85kVW4RocT&key=1674722119071`)
	req, err := http.NewRequest("POST", "https://cn.bing.com/tlookupv3?isVertical=1&&IG=E4B5DF32851B421C91509E7ED0EBAF5D&IID=translator.5023.1", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "cn.bing.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("cookie", "SRCHD=AF=ANNTA1; SRCHUID=V=2&GUID=6B9DDC1F0BD04818A9D20E88FBB46B72&dmnchg=1; PPLState=1; KievRPSSecAuth=FACCBBRaTOJILtFsMkpLVWSG6AN6C/svRwNmAAAEgAAACIjJG0oK2gmCQAR8E/UnaFDtS2Gq6akcs2sMzyZ1xiiY8+PWMPABTAPBFi+eaNnOYvf49hVDZqOafHt4oW/sPnUuGCauLIbD/IJc6gfQ+Jzhme6AUXH3AIEW2g5nGnPZyifXjyBxOZCDEfSPgetnYxBtOQXefv4zDskt/NKd8/Il/k/93bNObkp9KgKcQrcI+6TYUS/sDLErg/vdiLqQ2TMeEHjOnhvANsXvA3Z+mWC4/7qCo3Q14ZTqaHRcXufb4ZwYQbkOvDH/Vebpn5zCNYmUoVdDA7/g8thFYeGKuNdnu3QKiOwKn9NoRR3w7KQQWpDkgluIz8KbdROJmKDZFIQcoYlc+GZMWb+eCW0OuA6Di8USkiiq6fQfqe7xxGgCkrh8O+TdDt4gkXuyVfyH8vWbkId38XhVpKWjs8TbeJGlM6O1TkRv9nLWzcOAydMAHqj9IJgG2+IEeeSJ6JczjqAZFdTCPhneY4tGaypIJEYXv1uLedTc6xz3fxRUT5fewOh5YD4LuIpUwtQ6BKUtlnLeyYSqmsERGTuQUAxGg6/KbleZW6W5DTbYqNLRM4y6L/A0LwdXvJcmxns/YCjS6NRVZd8UP4/6/KoCs35Gg+qC5TbQatSvig1Z1IvqYfGQgV7RgMiDCzRN2RxtdUPC3OVrAK2/63NrbC2FZgKMGfCvFXukmlQ/O/JYH+RSMs/hHh9ouwWUA20Ts7/kNK3RId/rhjkSc4QGh7UgDAY096OGBpMHG12cGJE7ZODKs95ZwwIiGpJvSgdLuD3M4bAJmlij7UjLSwikOZNuaPDTPur8A8GnU2Bjc2z9LalCNJ1JXyL/hQR+fyHApEZ0CXMFvLAZ9RbM8+eQEzTgNpzoOabRCw4k41v4xbMs/3eOg5yv/SaKbL+3gtzmOrSwhByJOcjJu979crI7MX3lbx6kN40ldeApm0I7HQ6pnNOZa653LgKrE4OIiNzI7Y0g/1JEurHJ8ts7aKrhofDX/rbCx/xhV5n8o7qOFwDShQ3gkM/Db26Zl3Ow5KO7XNSxZVP7pv+6veH+0gSmCCuVZE6W9BuRGxsZQnCBt5B/97IFubzQZEYkjuYhhwMuIuCar5ePYnkHvLQoncTk/nuUe6zuPOUR3Mz3hh2ELOe1Dt45Jz5hm9vxzP4NOEPVYM9U3d0WYejbblkK6qSEZo3kOnbfBL3aw3px9I7r6l0XflTPAQOks2oX1UaWJKUILOMVncrbNwf//y/6/28nfc8MTJrXTtN09GeiOlrIbaEQG9hmAlmOyTcnrCDcFsDyMk9Buiz0H/PKtUkRyCmlFNps3m7bCzrxjxAlOk5zdfZby1Zys+jhWLo6FaWN5MK4SBLvSouDOAUUQ1tvzH/YnSj2PnZzh4UzTLRI+QGzPjNEoZIJ53sMFaij1O3rMuTMU4+HzhjLVHvLzzgrrDdWkQ5BxZ7O5JAIc/L/rAHXJvtLCxQAdC4Lh7WtqOSw7S7eojfdaBF/IQE=; _EDGE_V=1; MUID=1117B923440860470806AB26456E61E7; MUIDB=1117B923440860470806AB26456E61E7; _UR=QS=0&TQS=0; ANON=A=E90F2DBB752ED3626FDB6213FFFFFFFF&E=1b9c&W=1; NAP=V=1.9&E=1b42&C=MACrhfo7eS_AcTEOepRwFcTjT3qVDlrxwDqWkjHT43_3IlIiFiGATg&W=1; MMCASM=ID=E9F3A2B6951C429BAEA248EE14B94B9D; TTRSL=zh-Hans; _HPVN=CS=eyJQbiI6eyJDbiI6NSwiU3QiOjAsIlFzIjowLCJQcm9kIjoiUCJ9LCJTYyI6eyJDbiI6NSwiU3QiOjAsIlFzIjowLCJQcm9kIjoiSCJ9LCJReiI6eyJDbiI6NSwiU3QiOjAsIlFzIjowLCJQcm9kIjoiVCJ9LCJBcCI6dHJ1ZSwiTXV0ZSI6dHJ1ZSwiTGFkIjoiMjAyMi0xMi0wMlQwMDowMDowMFoiLCJJb3RkIjowLCJHd2IiOjAsIkRmdCI6bnVsbCwiTXZzIjowLCJGbHQiOjAsIkltcCI6MTR9; USRLOC=HS=1&CLOC=LAT=30.232887657466044|LON=120.03079445594473|A=733.4464586120832|TS=230126060726|SRC=W; ANIMIA=FRE=1; TRBDG=FIMPR=1; ZHCHATSTRONGATTRACT=TRUE; ZHCHATWEAKATTRACT=TRUE; WLS=C=ce57f9b49f2570e0&N=Sky; SUID=A; _SS=SID=100882846CF6656B2D4790216D906490&R=200&RB=0&GB=0&RG=200&RP=200&PC=U531; SRCHS=PC=U531; SRCHUSR=DOB=20220819&T=1674718614000&POEX=W; ipv6=hit=1674722218554&t=4; _U=1emYr1BMExVl4Q7dWX78V5RAIIzi2coa-NyRrH8Beif8BiqQ35rwBn0GDRvY-MA1_f0bX6eUgFIKnZ6YSkmvFmn-CqEGutNDXm2TosBL8h7E6J4ebKPiAfhsUXmqZjmHFJ2bRA6XtItzxcSG1nGW5iCudKGWrqF0SQ2kVlNVt2q7YSXmDM_ZXWu1jVCRCW4sn; _EDGE_S=ui=zh-cn&SID=100882846CF6656B2D4790216D906490; btstkn=B57lyClYQze%252Fmo5ni3MYjMcPEHgCnlO8bQQw%252FmORHaLkTKyz%252BDxdEI%252BWa2aXU1cfbHwvgTwqsvVZRjlL3puhEE2CO0yjtyqVy6wwj4GrlwY%253D; SNRHOP=I=&TS=; ABDEF=V=13&ABDV=11&MRNB=1674722082196&MRB=0; _RwBf=ilt=3001&ihpd=0&ispd=6&rc=200&rb=0&gb=0&rg=200&pc=200&mtu=0&rbb=0&g=0&cid=&clo=0&v=7&l=2023-01-26T08:00:00.0000000Z&lft=2023-01-17T00:00:00.0000000-08:00&aof=0&o=2&p=&c=&t=0&s=0001-01-01T00:00:00.0000000+00:00&ts=2023-01-26T08:34:43.2112772+00:00&rwred=0&wls=&lka=0&lkt=0&TH=; SRCHHPGUSR=BRH=M&CW=1013&UTC=480&CH=849&SCW=1164&SCH=3416&DPR=2.3&DM=0&SRCHLANG=zh-Hans&PV=10.0.0&BZA=0&PRVCW=1536&PRVCH=849&BRW=HTP&EXLTT=31&HV=1674722120&PR=2.25; _TTSS_IN=hist=WyJpZCIsInpoLUhhbnMiLCJlbiIsImF1dG8tZGV0ZWN0Il0=; _TTSS_OUT=hist=WyJlbiIsInpoLUhhbnMiXQ==; _tarLang=default=zh-Hans")
	req.Header.Set("origin", "https://cn.bing.com")
	req.Header.Set("referer", "https://cn.bing.com/translator")
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="99", "Microsoft Edge";v="109", "Chromium";v="109"`)
	req.Header.Set("sec-ch-ua-arch", `"x86"`)
	req.Header.Set("sec-ch-ua-bitness", `"64"`)
	req.Header.Set("sec-ch-ua-full-version", `"109.0.1518.61"`)
	req.Header.Set("sec-ch-ua-full-version-list", `"Not_A Brand";v="99.0.0.0", "Microsoft Edge";v="109.0.1518.61", "Chromium";v="109.0.5414.87"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-model", "")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-ch-ua-platform-version", `"10.0.0"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Edg/109.0.1518.61")
	req.Header.Set("x-client-data", "eyIxIjoiMTciLCIxMCI6IiIsIjIiOiIxIiwiMyI6IjAiLCI0IjoiLTcyMDAwOTA0OTU1NzMyNjk1NDAiLCI1IjoiXCJlVG50SmdjZVM4VE1KUVE5V1h3T0ZaV1orUjY5bDVXYkZSTHdPTXBMUUtjPVwiIiwiNiI6InN0YWJsZSIsIjciOiIxMjI4MzYwNjQ2NjYzIiwiOSI6ImRlc2t0b3AifQ==")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

}
