package model

import (
	"log"
	"time"
	"https://github.com/Stellasarah04/GroupieTracker_AUZET_Sarah/tree/main/SRC/go/config"
)

type Pays []struct {
	Name struct {
		Common     string `json:"common"`
		Official   string `json:"official"`
		NativeName struct {
			Deu struct {
				Official string `json:"official"`
				Common   string `json:"common"`
			} `json:"deu"`
		} `json:"nativeName"`
	} `json:"name,omitempty"`
	Tld         []string `json:"tld,omitempty"`
	Cca2        string   `json:"cca2"`
	Ccn3        string   `json:"ccn3,omitempty"`
	Cca3        string   `json:"cca3"`
	Cioc        string   `json:"cioc,omitempty"`
	Independent bool     `json:"independent,omitempty"`
	Status      string   `json:"status"`
	UnMember    bool     `json:"unMember"`
	Currencies  struct {
		Chf struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"CHF"`
	} `json:"currencies,omitempty"`
	Idd struct {
		Root     string   `json:"root"`
		Suffixes []string `json:"suffixes"`
	} `json:"idd,omitempty"`
	Capital      []string `json:"capital,omitempty"`
	AltSpellings []string `json:"altSpellings"`
	Region       string   `json:"region"`
	Subregion    string   `json:"subregion,omitempty"`
	Languages    struct {
		Deu string `json:"deu"`
	} `json:"languages,omitempty"`
	Translations struct {
		Ara struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ara"`
		Bre struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"bre"`
		Ces struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ces"`
		Cym struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"cym"`
		Deu struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"deu"`
		Est struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"est"`
		Fin struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fin"`
		Fra struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fra"`
		Hrv struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hrv"`
		Hun struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hun"`
		Ita struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ita"`
		Jpn struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"jpn"`
		Kor struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"kor"`
		Nld struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"nld"`
		Per struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"per"`
		Pol struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"pol"`
		Por struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"por"`
		Rus struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"rus"`
		Slk struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"slk"`
		Spa struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"spa"`
		Srp struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"srp"`
		Swe struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"swe"`
		Tur struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"tur"`
		Urd struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"urd"`
		Zho struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"zho"`
	} `json:"translations,omitempty"`
	Latlng     []float64 `json:"latlng"`
	Landlocked bool      `json:"landlocked"`
	Borders    []string  `json:"borders,omitempty"`
	Area       float64   `json:"area"`
	Demonyms   struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
		Fra struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"fra"`
	} `json:"demonyms,omitempty"`
	Flag string `json:"flag"`
	Maps struct {
		GoogleMaps     string `json:"googleMaps"`
		OpenStreetMaps string `json:"openStreetMaps"`
	} `json:"maps"`
	Population int    `json:"population"`
	Fifa       string `json:"fifa,omitempty"`
	Car        struct {
		Signs []string `json:"signs"`
		Side  string   `json:"side"`
	} `json:"car,omitempty"`
	Timezones  []string `json:"timezones"`
	Continents []string `json:"continents"`
	Flags      struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
		Alt string `json:"alt"`
	} `json:"flags,omitempty"`
	CoatOfArms struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"coatOfArms,omitempty"`
	StartOfWeek string `json:"startOfWeek"`
	CapitalInfo struct {
		Latlng []float64 `json:"latlng"`
	} `json:"capitalInfo,omitempty"`
	PostalCode struct {
		Format string `json:"format"`
		Regex  string `json:"regex"`
	} `json:"postalCode,omitempty"`
	Currencies0 struct {
		Bzd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"BZD"`
	} `json:"currencies,omitempty"`
	Languages0 struct {
		Bjz string `json:"bjz"`
		Eng string `json:"eng"`
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini struct {
		Num1999 float64 `json:"1999"`
	} `json:"gini,omitempty"`
	Currencies1 struct {
		Mkd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"MKD"`
	} `json:"currencies,omitempty"`
	Languages1 struct {
		Mkd string `json:"mkd"`
	} `json:"languages,omitempty"`
	Gini0 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies2 struct {
		Gbp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"GBP"`
		Jep struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"JEP"`
	} `json:"currencies,omitempty"`
	Languages2 struct {
		Eng string `json:"eng"`
		Fra string `json:"fra"`
		Nrf string `json:"nrf"`
	} `json:"languages,omitempty"`
	Flags0 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies3 struct {
		Mop struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"MOP"`
	} `json:"currencies,omitempty"`
	Languages3 struct {
		Por string `json:"por"`
		Zho string `json:"zho"`
	} `json:"languages,omitempty"`
	Translations0 struct {
		Ara struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ara"`
		Bre struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"bre"`
		Ces struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ces"`
		Cym struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"cym"`
		Deu struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"deu"`
		Est struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"est"`
		Fin struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fin"`
		Fra struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fra"`
		Hrv struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hrv"`
		Hun struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hun"`
		Ita struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ita"`
		Jpn struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"jpn"`
		Kor struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"kor"`
		Nld struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"nld"`
		Per struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"per"`
		Pol struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"pol"`
		Por struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"por"`
		Rus struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"rus"`
		Slk struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"slk"`
		Spa struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"spa"`
		Srp struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"srp"`
		Swe struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"swe"`
		Tur struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"tur"`
		Urd struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"urd"`
	} `json:"translations,omitempty"`
	Flags1 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CapitalInfo0 struct {
	} `json:"capitalInfo,omitempty"`
	Currencies4 struct {
		Pyg struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"PYG"`
	} `json:"currencies,omitempty"`
	Languages4 struct {
		Grn string `json:"grn"`
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini1 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies5 struct {
		Awg struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"AWG"`
	} `json:"currencies,omitempty"`
	Languages5 struct {
		Nld string `json:"nld"`
		Pap string `json:"pap"`
	} `json:"languages,omitempty"`
	Car0 struct {
		Side string `json:"side"`
	} `json:"car,omitempty"`
	Flags2 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies6 struct {
		Lkr struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"LKR"`
	} `json:"currencies,omitempty"`
	Languages6 struct {
		Sin string `json:"sin"`
		Tam string `json:"tam"`
	} `json:"languages,omitempty"`
	Gini2 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	Currencies7 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages7 struct {
		Ita string `json:"ita"`
	} `json:"languages,omitempty"`
	Gini3 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies8 struct {
		Ang struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"ANG"`
	} `json:"currencies,omitempty"`
	Languages8 struct {
		Eng string `json:"eng"`
		Nld string `json:"nld"`
		Pap string `json:"pap"`
	} `json:"languages,omitempty"`
	Translations1 struct {
		Ara struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ara"`
		Bre struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"bre"`
		Ces struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ces"`
		Cym struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"cym"`
		Deu struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"deu"`
		Est struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"est"`
		Fin struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fin"`
		Fra struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fra"`
		Hun struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hun"`
		Ita struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ita"`
		Kor struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"kor"`
		Nld struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"nld"`
		Per struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"per"`
		Pol struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"pol"`
		Por struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"por"`
		Rus struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"rus"`
		Slk struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"slk"`
		Spa struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"spa"`
		Srp struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"srp"`
		Swe struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"swe"`
		Tur struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"tur"`
		Urd struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"urd"`
		Zho struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"zho"`
	} `json:"translations,omitempty"`
	Flags3 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies9 struct {
		Ang struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"ANG"`
	} `json:"currencies,omitempty"`
	Languages9 struct {
		Eng string `json:"eng"`
		Fra string `json:"fra"`
		Nld string `json:"nld"`
	} `json:"languages,omitempty"`
	Flags4 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms0 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies10 struct {
		Xof struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XOF"`
	} `json:"currencies,omitempty"`
	Languages10 struct {
		Por string `json:"por"`
		Pov string `json:"pov"`
	} `json:"languages,omitempty"`
	Gini4 struct {
		Num2010 float64 `json:"2010"`
	} `json:"gini,omitempty"`
	Currencies11 struct {
		Egp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EGP"`
	} `json:"currencies,omitempty"`
	Languages11 struct {
		Ara string `json:"ara"`
	} `json:"languages,omitempty"`
	Gini5 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies12 struct {
		Mru struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"MRU"`
	} `json:"currencies,omitempty"`
	Languages12 struct {
		Ara string `json:"ara"`
	} `json:"languages,omitempty"`
	Gini6 struct {
		Num2014 float64 `json:"2014"`
	} `json:"gini,omitempty"`
	Currencies13 struct {
		Nzd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"NZD"`
	} `json:"currencies,omitempty"`
	Languages13 struct {
		Eng string `json:"eng"`
		Smo string `json:"smo"`
		Tkl string `json:"tkl"`
	} `json:"languages,omitempty"`
	Demonyms0 struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
	} `json:"demonyms,omitempty"`
	Flags5 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms1 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies14 struct {
		Xof struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XOF"`
	} `json:"currencies,omitempty"`
	Languages14 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini7 struct {
		Num2015 float64 `json:"2015"`
	} `json:"gini,omitempty"`
	Currencies15 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages15 struct {
		Slk string `json:"slk"`
	} `json:"languages,omitempty"`
	Gini8 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies16 struct {
		Amd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"AMD"`
	} `json:"currencies,omitempty"`
	Languages16 struct {
		Hye string `json:"hye"`
	} `json:"languages,omitempty"`
	Gini9 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies17 struct {
		Ern struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"ERN"`
	} `json:"currencies,omitempty"`
	Languages17 struct {
		Ara string `json:"ara"`
		Eng string `json:"eng"`
		Tir string `json:"tir"`
	} `json:"languages,omitempty"`
	Currencies18 struct {
		Lsl struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"LSL"`
		Zar struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"ZAR"`
	} `json:"currencies,omitempty"`
	Languages18 struct {
		Eng string `json:"eng"`
		Sot string `json:"sot"`
	} `json:"languages,omitempty"`
	Gini10 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies19 struct {
		Pkr struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"PKR"`
	} `json:"currencies,omitempty"`
	Languages19 struct {
		Eng string `json:"eng"`
		Urd string `json:"urd"`
	} `json:"languages,omitempty"`
	Gini11 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies20 struct {
		Gbp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"GBP"`
		Ggp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"GGP"`
	} `json:"currencies,omitempty"`
	Languages20 struct {
		Eng string `json:"eng"`
		Fra string `json:"fra"`
		Nfr string `json:"nfr"`
	} `json:"languages,omitempty"`
	Flags6 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies21 struct {
		Gip struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"GIP"`
	} `json:"currencies,omitempty"`
	Languages21 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Flags7 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies22 struct {
		Isk struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"ISK"`
	} `json:"currencies,omitempty"`
	Languages22 struct {
		Isl string `json:"isl"`
	} `json:"languages,omitempty"`
	Gini12 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies23 struct {
		Bhd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"BHD"`
	} `json:"currencies,omitempty"`
	Languages23 struct {
		Ara string `json:"ara"`
	} `json:"languages,omitempty"`
	Currencies24 struct {
		Uzs struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"UZS"`
	} `json:"currencies,omitempty"`
	Languages24 struct {
		Rus string `json:"rus"`
		Uzb string `json:"uzb"`
	} `json:"languages,omitempty"`
	Gini13 struct {
		Num2003 float64 `json:"2003"`
	} `json:"gini,omitempty"`
	Currencies25 struct {
		Mur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"MUR"`
	} `json:"currencies,omitempty"`
	Languages25 struct {
		Eng string `json:"eng"`
		Fra string `json:"fra"`
		Mfe string `json:"mfe"`
	} `json:"languages,omitempty"`
	Gini14 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies26 struct {
		Gtq struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"GTQ"`
	} `json:"currencies,omitempty"`
	Languages26 struct {
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini15 struct {
		Num2014 float64 `json:"2014"`
	} `json:"gini,omitempty"`
	Currencies27 struct {
		Zwl struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"ZWL"`
	} `json:"currencies,omitempty"`
	Languages27 struct {
		Bwg string `json:"bwg"`
		Eng string `json:"eng"`
		Kck string `json:"kck"`
		Khi string `json:"khi"`
		Ndc string `json:"ndc"`
		Nde string `json:"nde"`
		Nya string `json:"nya"`
		Sna string `json:"sna"`
		Sot string `json:"sot"`
		Toi string `json:"toi"`
		Tsn string `json:"tsn"`
		Tso string `json:"tso"`
		Ven string `json:"ven"`
		Xho string `json:"xho"`
		Zib string `json:"zib"`
	} `json:"languages,omitempty"`
	Gini16 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies28 struct {
		Xcd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XCD"`
	} `json:"currencies,omitempty"`
	Languages28 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Flags8 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies29 struct {
		Uah struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"UAH"`
	} `json:"currencies,omitempty"`
	Languages29 struct {
		Ukr string `json:"ukr"`
	} `json:"languages,omitempty"`
	Gini17 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies30 struct {
		Xpf struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XPF"`
	} `json:"currencies,omitempty"`
	Languages30 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Demonyms1 struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
	} `json:"demonyms,omitempty"`
	Flags9 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms2 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies31 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages31 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Flags10 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies32 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages32 struct {
		Est string `json:"est"`
	} `json:"languages,omitempty"`
	Gini18 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies33 struct {
		Bdt struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"BDT"`
	} `json:"currencies,omitempty"`
	Languages33 struct {
		Ben string `json:"ben"`
	} `json:"languages,omitempty"`
	Gini19 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	Currencies34 struct {
		Mwk struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"MWK"`
	} `json:"currencies,omitempty"`
	Languages34 struct {
		Eng string `json:"eng"`
		Nya string `json:"nya"`
	} `json:"languages,omitempty"`
	Gini20 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	Currencies35 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages35 struct {
		Slv string `json:"slv"`
	} `json:"languages,omitempty"`
	Gini21 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies36 struct {
		Aud struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"AUD"`
		Kid struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"KID"`
	} `json:"currencies,omitempty"`
	Languages36 struct {
		Eng string `json:"eng"`
		Gil string `json:"gil"`
	} `json:"languages,omitempty"`
	Gini22 struct {
		Num2006 float64 `json:"2006"`
	} `json:"gini,omitempty"`
	Currencies37 struct {
		Nok struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"NOK"`
	} `json:"currencies,omitempty"`
	Languages37 struct {
		Nno string `json:"nno"`
		Nob string `json:"nob"`
		Smi string `json:"smi"`
	} `json:"languages,omitempty"`
	Gini23 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies38 struct {
		Bob struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"BOB"`
	} `json:"currencies,omitempty"`
	Languages38 struct {
		Aym string `json:"aym"`
		Grn string `json:"grn"`
		Que string `json:"que"`
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini24 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies39 struct {
		Xaf struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XAF"`
	} `json:"currencies,omitempty"`
	Languages39 struct {
		Fra string `json:"fra"`
		Por string `json:"por"`
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Currencies40 struct {
		Szl struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"SZL"`
		Zar struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"ZAR"`
	} `json:"currencies,omitempty"`
	Languages40 struct {
		Eng string `json:"eng"`
		Ssw string `json:"ssw"`
	} `json:"languages,omitempty"`
	Gini25 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	CoatOfArms3 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies41 struct {
		Aoa struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"AOA"`
	} `json:"currencies,omitempty"`
	Languages41 struct {
		Por string `json:"por"`
	} `json:"languages,omitempty"`
	Gini26 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies42 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages42 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Flags11 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies43 struct {
		Egp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EGP"`
		Ils struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"ILS"`
		Jod struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"JOD"`
	} `json:"currencies,omitempty"`
	Languages43 struct {
		Ara string `json:"ara"`
	} `json:"languages,omitempty"`
	Gini27 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	Flags12 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Languages44 struct {
		Fra string `json:"fra"`
		Gsw string `json:"gsw"`
		Ita string `json:"ita"`
		Roh string `json:"roh"`
	} `json:"languages,omitempty"`
	Gini28 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies44 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages45 struct {
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini29 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies45 struct {
		Mxn struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"MXN"`
	} `json:"currencies,omitempty"`
	Languages46 struct {
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini30 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies46 struct {
		Xcd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XCD"`
	} `json:"currencies,omitempty"`
	Languages47 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Currencies47 struct {
		Nio struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"NIO"`
	} `json:"currencies,omitempty"`
	Languages48 struct {
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini31 struct {
		Num2014 float64 `json:"2014"`
	} `json:"gini,omitempty"`
	Currencies48 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages49 struct {
		Cat string `json:"cat"`
	} `json:"languages,omitempty"`
	Currencies49 struct {
		Jmd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"JMD"`
	} `json:"currencies,omitempty"`
	Languages50 struct {
		Eng string `json:"eng"`
		Jam string `json:"jam"`
	} `json:"languages,omitempty"`
	Gini32 struct {
		Num2004 float64 `json:"2004"`
	} `json:"gini,omitempty"`
	Currencies50 struct {
		Xof struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XOF"`
	} `json:"currencies,omitempty"`
	Languages51 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini33 struct {
		Num2009 float64 `json:"2009"`
	} `json:"gini,omitempty"`
	Currencies51 struct {
		Syp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"SYP"`
	} `json:"currencies,omitempty"`
	Languages52 struct {
		Ara string `json:"ara"`
	} `json:"languages,omitempty"`
	Gini34 struct {
		Num2003 float64 `json:"2003"`
	} `json:"gini,omitempty"`
	Currencies52 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages53 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Demonyms2 struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
	} `json:"demonyms,omitempty"`
	Flags13 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms4 struct {
	} `json:"coatOfArms,omitempty"`
	CapitalInfo1 struct {
	} `json:"capitalInfo,omitempty"`
	Currencies53 struct {
		Gbp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"GBP"`
		Shp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"SHP"`
	} `json:"currencies,omitempty"`
	Languages54 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Flags14 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms5 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies54 struct {
		Tnd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"TND"`
	} `json:"currencies,omitempty"`
	Languages55 struct {
		Ara string `json:"ara"`
	} `json:"languages,omitempty"`
	Gini35 struct {
		Num2015 float64 `json:"2015"`
	} `json:"gini,omitempty"`
	Currencies55 struct {
		Dkk struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"DKK"`
	} `json:"currencies,omitempty"`
	Languages56 struct {
		Kal string `json:"kal"`
	} `json:"languages,omitempty"`
	Flags15 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies56 struct {
		Xaf struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XAF"`
	} `json:"currencies,omitempty"`
	Languages57 struct {
		Eng string `json:"eng"`
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini36 struct {
		Num2014 float64 `json:"2014"`
	} `json:"gini,omitempty"`
	Currencies57 struct {
		Jpy struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"JPY"`
	} `json:"currencies,omitempty"`
	Languages58 struct {
		Jpn string `json:"jpn"`
	} `json:"languages,omitempty"`
	Gini37 struct {
		Num2013 float64 `json:"2013"`
	} `json:"gini,omitempty"`
	Currencies58 struct {
		Gyd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"GYD"`
	} `json:"currencies,omitempty"`
	Languages59 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Gini38 struct {
		Num1998 float64 `json:"1998"`
	} `json:"gini,omitempty"`
	Currencies59 struct {
		Nzd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"NZD"`
	} `json:"currencies,omitempty"`
	Languages60 struct {
		Eng string `json:"eng"`
		Niu string `json:"niu"`
	} `json:"languages,omitempty"`
	Flags16 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms6 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies60 struct {
		Gbp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"GBP"`
	} `json:"currencies,omitempty"`
	Languages61 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Gini39 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies61 struct {
		Iqd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"IQD"`
	} `json:"currencies,omitempty"`
	Languages62 struct {
		Ara string `json:"ara"`
		Arc string `json:"arc"`
		Ckb string `json:"ckb"`
	} `json:"languages,omitempty"`
	Gini40 struct {
		Num2012 float64 `json:"2012"`
	} `json:"gini,omitempty"`
	Currencies62 struct {
		Mad struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"MAD"`
	} `json:"currencies,omitempty"`
	Languages63 struct {
		Ara string `json:"ara"`
		Ber string `json:"ber"`
	} `json:"languages,omitempty"`
	Gini41 struct {
		Num2013 float64 `json:"2013"`
	} `json:"gini,omitempty"`
	Currencies63 struct {
		Pgk struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"PGK"`
	} `json:"currencies,omitempty"`
	Languages64 struct {
		Eng string `json:"eng"`
		Hmo string `json:"hmo"`
		Tpi string `json:"tpi"`
	} `json:"languages,omitempty"`
	Gini42 struct {
		Num2009 float64 `json:"2009"`
	} `json:"gini,omitempty"`
	Currencies64 struct {
		Ugx struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"UGX"`
	} `json:"currencies,omitempty"`
	Languages65 struct {
		Eng string `json:"eng"`
		Swa string `json:"swa"`
	} `json:"languages,omitempty"`
	Gini43 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	Currencies65 struct {
		Azn struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"AZN"`
	} `json:"currencies,omitempty"`
	Languages66 struct {
		Aze string `json:"aze"`
		Rus string `json:"rus"`
	} `json:"languages,omitempty"`
	Gini44 struct {
		Num2005 float64 `json:"2005"`
	} `json:"gini,omitempty"`
	Currencies66 struct {
		Czk struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"CZK"`
	} `json:"currencies,omitempty"`
	Languages67 struct {
		Ces string `json:"ces"`
		Slk string `json:"slk"`
	} `json:"languages,omitempty"`
	Gini45 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies67 struct {
		Pen struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"PEN"`
	} `json:"currencies,omitempty"`
	Languages68 struct {
		Aym string `json:"aym"`
		Que string `json:"que"`
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini46 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies68 struct {
		Tjs struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"TJS"`
	} `json:"currencies,omitempty"`
	Languages69 struct {
		Rus string `json:"rus"`
		Tgk string `json:"tgk"`
	} `json:"languages,omitempty"`
	Gini47 struct {
		Num2015 float64 `json:"2015"`
	} `json:"gini,omitempty"`
	Currencies69 struct {
		Ron struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"RON"`
	} `json:"currencies,omitempty"`
	Languages70 struct {
		Ron string `json:"ron"`
	} `json:"languages,omitempty"`
	Gini48 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies70 struct {
		Top struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"TOP"`
	} `json:"currencies,omitempty"`
	Languages71 struct {
		Eng string `json:"eng"`
		Ton string `json:"ton"`
	} `json:"languages,omitempty"`
	Gini49 struct {
		Num2015 float64 `json:"2015"`
	} `json:"gini,omitempty"`
	Currencies71 struct {
		Ssp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"SSP"`
	} `json:"currencies,omitempty"`
	Languages72 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Gini50 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	Currencies72 struct {
		Mvr struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"MVR"`
	} `json:"currencies,omitempty"`
	Languages73 struct {
		Div string `json:"div"`
	} `json:"languages,omitempty"`
	Gini51 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	Currencies73 struct {
		Try struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"TRY"`
	} `json:"currencies,omitempty"`
	Languages74 struct {
		Tur string `json:"tur"`
	} `json:"languages,omitempty"`
	Gini52 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies74 struct {
		Rwf struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"RWF"`
	} `json:"currencies,omitempty"`
	Languages75 struct {
		Eng string `json:"eng"`
		Fra string `json:"fra"`
		Kin string `json:"kin"`
	} `json:"languages,omitempty"`
	Gini53 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	Currencies75 struct {
		Aud struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"AUD"`
	} `json:"currencies,omitempty"`
	Languages76 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Demonyms3 struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
	} `json:"demonyms,omitempty"`
	Flags17 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies76 struct {
		Aud struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"AUD"`
	} `json:"currencies,omitempty"`
	Languages77 struct {
		Eng string `json:"eng"`
		Nau string `json:"nau"`
	} `json:"languages,omitempty"`
	Gini54 struct {
		Num2012 float64 `json:"2012"`
	} `json:"gini,omitempty"`
	Currencies77 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages78 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Flags18 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms7 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies78 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages79 struct {
		Por string `json:"por"`
	} `json:"languages,omitempty"`
	Gini55 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies79 struct {
		Irr struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"IRR"`
	} `json:"currencies,omitempty"`
	Languages80 struct {
		Fas string `json:"fas"`
	} `json:"languages,omitempty"`
	Translations2 struct {
		Ara struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ara"`
		Bre struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"bre"`
		Ces struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ces"`
		Cym struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"cym"`
		Deu struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"deu"`
		Est struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"est"`
		Fin struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fin"`
		Fra struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fra"`
		Hrv struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hrv"`
		Hun struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hun"`
		Ita struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ita"`
		Jpn struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"jpn"`
		Kor struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"kor"`
		Nld struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"nld"`
		Pol struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"pol"`
		Por struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"por"`
		Rus struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"rus"`
		Slk struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"slk"`
		Spa struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"spa"`
		Srp struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"srp"`
		Swe struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"swe"`
		Tur struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"tur"`
		Urd struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"urd"`
		Zho struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"zho"`
	} `json:"translations,omitempty"`
	Gini56 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies80 struct {
		Uyu struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"UYU"`
	} `json:"currencies,omitempty"`
	Languages81 struct {
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini57 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies81 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages82 struct {
		Cnr string `json:"cnr"`
	} `json:"languages,omitempty"`
	Gini58 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	Currencies82 struct {
		Aud struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"AUD"`
	} `json:"currencies,omitempty"`
	Languages83 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Demonyms4 struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
	} `json:"demonyms,omitempty"`
	Flags19 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms8 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies83 struct {
		Tmt struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"TMT"`
	} `json:"currencies,omitempty"`
	Languages84 struct {
		Rus string `json:"rus"`
		Tuk string `json:"tuk"`
	} `json:"languages,omitempty"`
	Gini59 struct {
		Num1998 float64 `json:"1998"`
	} `json:"gini,omitempty"`
	Currencies84 struct {
		Nzd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"NZD"`
	} `json:"currencies,omitempty"`
	Languages85 struct {
		Eng string `json:"eng"`
		Mri string `json:"mri"`
		Nzs string `json:"nzs"`
	} `json:"languages,omitempty"`
	Currencies85 struct {
		Kpw struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"KPW"`
	} `json:"currencies,omitempty"`
	Languages86 struct {
		Kor string `json:"kor"`
	} `json:"languages,omitempty"`
	Currencies86 struct {
		Thb struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"THB"`
	} `json:"currencies,omitempty"`
	Languages87 struct {
		Tha string `json:"tha"`
	} `json:"languages,omitempty"`
	Gini60 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies87 struct {
		Inr struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"INR"`
	} `json:"currencies,omitempty"`
	Languages88 struct {
		Eng string `json:"eng"`
		Hin string `json:"hin"`
		Tam string `json:"tam"`
	} `json:"languages,omitempty"`
	Gini61 struct {
		Num2011 float64 `json:"2011"`
	} `json:"gini,omitempty"`
	Currencies88 struct {
		Dkk struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"DKK"`
		Fok struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"FOK"`
	} `json:"currencies,omitempty"`
	Languages89 struct {
		Dan string `json:"dan"`
		Fao string `json:"fao"`
	} `json:"languages,omitempty"`
	Flags20 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies89 struct {
		Lyd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"LYD"`
	} `json:"currencies,omitempty"`
	Languages90 struct {
		Ara string `json:"ara"`
	} `json:"languages,omitempty"`
	Currencies90 struct {
		Mzn struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"MZN"`
	} `json:"currencies,omitempty"`
	Languages91 struct {
		Por string `json:"por"`
	} `json:"languages,omitempty"`
	Gini62 struct {
		Num2014 float64 `json:"2014"`
	} `json:"gini,omitempty"`
	Currencies91 struct {
		Xpf struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XPF"`
	} `json:"currencies,omitempty"`
	Languages92 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Flags21 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies92 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages93 struct {
		Por string `json:"por"`
		Tet string `json:"tet"`
	} `json:"languages,omitempty"`
	Gini63 struct {
		Num2014 float64 `json:"2014"`
	} `json:"gini,omitempty"`
	CoatOfArms9 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies93 struct {
		Etb struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"ETB"`
	} `json:"currencies,omitempty"`
	Languages94 struct {
		Amh string `json:"amh"`
	} `json:"languages,omitempty"`
	Gini64 struct {
		Num2015 float64 `json:"2015"`
	} `json:"gini,omitempty"`
	Currencies94 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages95 struct {
		Lit string `json:"lit"`
	} `json:"languages,omitempty"`
	Gini65 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies95 struct {
		Lbp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"LBP"`
	} `json:"currencies,omitempty"`
	Languages96 struct {
		Ara string `json:"ara"`
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini66 struct {
		Num2011 float64 `json:"2011"`
	} `json:"gini,omitempty"`
	Currencies96 struct {
		Kyd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"KYD"`
	} `json:"currencies,omitempty"`
	Languages97 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Flags22 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies97 struct {
		Mdl struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"MDL"`
	} `json:"currencies,omitempty"`
	Languages98 struct {
		Ron string `json:"ron"`
	} `json:"languages,omitempty"`
	Gini67 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies98 struct {
		Lak struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"LAK"`
	} `json:"currencies,omitempty"`
	Languages99 struct {
		Lao string `json:"lao"`
	} `json:"languages,omitempty"`
	Gini68 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies99 struct {
		Ngn struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"NGN"`
	} `json:"currencies,omitempty"`
	Languages100 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Gini69 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies100 struct {
		Clp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"CLP"`
	} `json:"currencies,omitempty"`
	Languages101 struct {
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini70 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies101 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages102 struct {
		Hrv string `json:"hrv"`
	} `json:"languages,omitempty"`
	Gini71 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies102 struct {
		Bbd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"BBD"`
	} `json:"currencies,omitempty"`
	Languages103 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Currencies103 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages104 struct {
		Ita string `json:"ita"`
		Lat string `json:"lat"`
	} `json:"languages,omitempty"`
	Currencies104 struct {
		Xaf struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XAF"`
	} `json:"currencies,omitempty"`
	Languages105 struct {
		Fra string `json:"fra"`
		Sag string `json:"sag"`
	} `json:"languages,omitempty"`
	Gini72 struct {
		Num2008 float64 `json:"2008"`
	} `json:"gini,omitempty"`
	Currencies105 struct {
		Zmw struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"ZMW"`
	} `json:"currencies,omitempty"`
	Languages106 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Gini73 struct {
		Num2015 float64 `json:"2015"`
	} `json:"gini,omitempty"`
	Currencies106 struct {
		Bif struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"BIF"`
	} `json:"currencies,omitempty"`
	Languages107 struct {
		Fra string `json:"fra"`
		Run string `json:"run"`
	} `json:"languages,omitempty"`
	Gini74 struct {
		Num2013 float64 `json:"2013"`
	} `json:"gini,omitempty"`
	Currencies107 struct {
		Sll struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"SLL"`
	} `json:"currencies,omitempty"`
	Languages108 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Gini75 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies108 struct {
		Sos struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"SOS"`
	} `json:"currencies,omitempty"`
	Languages109 struct {
		Ara string `json:"ara"`
		Som string `json:"som"`
	} `json:"languages,omitempty"`
	Gini76 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies109 struct {
		Sar struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"SAR"`
	} `json:"currencies,omitempty"`
	Languages110 struct {
		Ara string `json:"ara"`
	} `json:"languages,omitempty"`
	Currencies110 struct {
		Aed struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"AED"`
	} `json:"currencies,omitempty"`
	Languages111 struct {
		Ara string `json:"ara"`
	} `json:"languages,omitempty"`
	Gini77 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies111 struct {
		Khr struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"KHR"`
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages112 struct {
		Khm string `json:"khm"`
	} `json:"languages,omitempty"`
	Currencies112 struct {
		Nad struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"NAD"`
		Zar struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"ZAR"`
	} `json:"currencies,omitempty"`
	Languages113 struct {
		Afr string `json:"afr"`
		Deu string `json:"deu"`
		Eng string `json:"eng"`
		Her string `json:"her"`
		Hgm string `json:"hgm"`
		Kwn string `json:"kwn"`
		Loz string `json:"loz"`
		Ndo string `json:"ndo"`
		Tsn string `json:"tsn"`
	} `json:"languages,omitempty"`
	Gini78 struct {
		Num2015 float64 `json:"2015"`
	} `json:"gini,omitempty"`
	Currencies113 struct {
		Tzs struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"TZS"`
	} `json:"currencies,omitempty"`
	Languages114 struct {
		Eng string `json:"eng"`
		Swa string `json:"swa"`
	} `json:"languages,omitempty"`
	Gini79 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies114 struct {
		Fjd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"FJD"`
	} `json:"currencies,omitempty"`
	Languages115 struct {
		Eng string `json:"eng"`
		Fij string `json:"fij"`
		Hif string `json:"hif"`
	} `json:"languages,omitempty"`
	Gini80 struct {
		Num2013 float64 `json:"2013"`
	} `json:"gini,omitempty"`
	Currencies115 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages116 struct {
		Eng string `json:"eng"`
		Nld string `json:"nld"`
		Pap string `json:"pap"`
	} `json:"languages,omitempty"`
	Flags23 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies116 struct {
		Sgd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"SGD"`
	} `json:"currencies,omitempty"`
	Languages117 struct {
		Zho string `json:"zho"`
		Eng string `json:"eng"`
		Msa string `json:"msa"`
		Tam string `json:"tam"`
	} `json:"languages,omitempty"`
	Translations3 struct {
		Ara struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ara"`
		Bre struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"bre"`
		Ces struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ces"`
		Cym struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"cym"`
		Deu struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"deu"`
		Est struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"est"`
		Fin struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fin"`
		Fra struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fra"`
		Hrv struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hrv"`
		Hun struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hun"`
		Ita struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ita"`
		Jpn struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"jpn"`
		Kor struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"kor"`
		Nld struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"nld"`
		Per struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"per"`
		Pol struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"pol"`
		Por struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"por"`
		Rus struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"rus"`
		Slk struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"slk"`
		Spa struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"spa"`
		Srp struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"srp"`
		Swe struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"swe"`
		Tur struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"tur"`
		Urd struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"urd"`
	} `json:"translations,omitempty"`
	Currencies117 struct {
		Gmd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"GMD"`
	} `json:"currencies,omitempty"`
	Languages118 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Gini81 struct {
		Num2015 float64 `json:"2015"`
	} `json:"gini,omitempty"`
	Currencies118 struct {
		Dop struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"DOP"`
	} `json:"currencies,omitempty"`
	Languages119 struct {
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini82 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies119 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages120 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Demonyms5 struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
	} `json:"demonyms,omitempty"`
	Flags24 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms10 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies120 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages121 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Gini83 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies121 struct {
		Ves struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"VES"`
	} `json:"currencies,omitempty"`
	Languages122 struct {
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini84 struct {
		Num2006 float64 `json:"2006"`
	} `json:"gini,omitempty"`
	Currencies122 struct {
		Byn struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"BYN"`
	} `json:"currencies,omitempty"`
	Languages123 struct {
		Bel string `json:"bel"`
		Rus string `json:"rus"`
	} `json:"languages,omitempty"`
	Gini85 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies123 struct {
		Npr struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"NPR"`
	} `json:"currencies,omitempty"`
	Languages124 struct {
		Nep string `json:"nep"`
	} `json:"languages,omitempty"`
	Gini86 struct {
		Num2010 float64 `json:"2010"`
	} `json:"gini,omitempty"`
	Currencies124 struct {
		Zar struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"ZAR"`
	} `json:"currencies,omitempty"`
	Languages125 struct {
		Afr string `json:"afr"`
		Eng string `json:"eng"`
		Nbl string `json:"nbl"`
		Nso string `json:"nso"`
		Sot string `json:"sot"`
		Ssw string `json:"ssw"`
		Tsn string `json:"tsn"`
		Tso string `json:"tso"`
		Ven string `json:"ven"`
		Xho string `json:"xho"`
		Zul string `json:"zul"`
	} `json:"languages,omitempty"`
	Gini87 struct {
		Num2014 float64 `json:"2014"`
	} `json:"gini,omitempty"`
	Currencies125 struct {
		Pln struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"PLN"`
	} `json:"currencies,omitempty"`
	Languages126 struct {
		Pol string `json:"pol"`
	} `json:"languages,omitempty"`
	Gini88 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies126 struct {
		Myr struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"MYR"`
	} `json:"currencies,omitempty"`
	Languages127 struct {
		Eng string `json:"eng"`
		Msa string `json:"msa"`
	} `json:"languages,omitempty"`
	Gini89 struct {
		Num2015 float64 `json:"2015"`
	} `json:"gini,omitempty"`
	Currencies127 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages128 struct {
		Swe string `json:"swe"`
	} `json:"languages,omitempty"`
	Flags25 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies128 struct {
		Cve struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"CVE"`
	} `json:"currencies,omitempty"`
	Languages129 struct {
		Por string `json:"por"`
	} `json:"languages,omitempty"`
	Gini90 struct {
		Num2015 float64 `json:"2015"`
	} `json:"gini,omitempty"`
	Currencies129 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages130 struct {
		Eng string `json:"eng"`
		Pau string `json:"pau"`
	} `json:"languages,omitempty"`
	Currencies130 struct {
		Ghs struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"GHS"`
	} `json:"currencies,omitempty"`
	Languages131 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Gini91 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	Currencies131 struct {
		Bam struct {
			Name string `json:"name"`
		} `json:"BAM"`
	} `json:"currencies,omitempty"`
	Languages132 struct {
		Bos string `json:"bos"`
		Hrv string `json:"hrv"`
		Srp string `json:"srp"`
	} `json:"languages,omitempty"`
	Gini92 struct {
		Num2011 float64 `json:"2011"`
	} `json:"gini,omitempty"`
	Currencies132 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages133 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Currencies133 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Gini93 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	Currencies134 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages134 struct {
		Cha string `json:"cha"`
		Eng string `json:"eng"`
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Demonyms6 struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
	} `json:"demonyms,omitempty"`
	Flags26 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies135 struct {
		Jod struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"JOD"`
	} `json:"currencies,omitempty"`
	Languages135 struct {
		Ara string `json:"ara"`
	} `json:"languages,omitempty"`
	Gini94 struct {
		Num2010 float64 `json:"2010"`
	} `json:"gini,omitempty"`
	Currencies136 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages136 struct {
		Deu string `json:"deu"`
		Fra string `json:"fra"`
		Nld string `json:"nld"`
	} `json:"languages,omitempty"`
	Gini95 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies137 struct {
		Bsd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"BSD"`
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages137 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Currencies138 struct {
		Kgs struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"KGS"`
	} `json:"currencies,omitempty"`
	Languages138 struct {
		Kir string `json:"kir"`
		Rus string `json:"rus"`
	} `json:"languages,omitempty"`
	Gini96 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies139 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages139 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Demonyms7 struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
	} `json:"demonyms,omitempty"`
	Flags27 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies140 struct {
		Vuv struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"VUV"`
	} `json:"currencies,omitempty"`
	Languages140 struct {
		Bis string `json:"bis"`
		Eng string `json:"eng"`
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini97 struct {
		Num2010 float64 `json:"2010"`
	} `json:"gini,omitempty"`
	Currencies141 struct {
		Aud struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"AUD"`
		Tvd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"TVD"`
	} `json:"currencies,omitempty"`
	Languages141 struct {
		Eng string `json:"eng"`
		Tvl string `json:"tvl"`
	} `json:"languages,omitempty"`
	Gini98 struct {
		Num2010 float64 `json:"2010"`
	} `json:"gini,omitempty"`
	Currencies142 struct {
		Hnl struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"HNL"`
	} `json:"currencies,omitempty"`
	Languages142 struct {
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini99 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies143 struct {
		Sbd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"SBD"`
	} `json:"currencies,omitempty"`
	Languages143 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Gini100 struct {
		Num2012 float64 `json:"2012"`
	} `json:"gini,omitempty"`
	Currencies144 struct {
		Xof struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XOF"`
	} `json:"currencies,omitempty"`
	Languages144 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini101 struct {
		Num2015 float64 `json:"2015"`
	} `json:"gini,omitempty"`
	Currencies145 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages145 struct {
		Eng string `json:"eng"`
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Flags28 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms11 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies146 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages146 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Flags29 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms12 struct {
	} `json:"coatOfArms,omitempty"`
	PostalCode0 struct {
		Format string `json:"format"`
	} `json:"postalCode,omitempty"`
	Currencies147 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages147 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Flags30 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies148 struct {
		Dkk struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"DKK"`
	} `json:"currencies,omitempty"`
	Languages148 struct {
		Dan string `json:"dan"`
	} `json:"languages,omitempty"`
	Gini102 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies149 struct {
		Wst struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"WST"`
	} `json:"currencies,omitempty"`
	Languages149 struct {
		Eng string `json:"eng"`
		Smo string `json:"smo"`
	} `json:"languages,omitempty"`
	Gini103 struct {
		Num2013 float64 `json:"2013"`
	} `json:"gini,omitempty"`
	Currencies150 struct {
		Xof struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XOF"`
	} `json:"currencies,omitempty"`
	Languages150 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini104 struct {
		Num2014 float64 `json:"2014"`
	} `json:"gini,omitempty"`
	Currencies151 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages151 struct {
		Cal string `json:"cal"`
		Cha string `json:"cha"`
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Flags31 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms13 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies152 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages152 struct {
		Eng string `json:"eng"`
		Smo string `json:"smo"`
	} `json:"languages,omitempty"`
	Flags32 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms14 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies153 struct {
		Bmd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"BMD"`
	} `json:"currencies,omitempty"`
	Languages153 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Flags33 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies154 struct {
		Xcd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XCD"`
	} `json:"currencies,omitempty"`
	Languages154 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Currencies155 struct {
		Dzd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"DZD"`
		Mad struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"MAD"`
		Mru struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"MRU"`
	} `json:"currencies,omitempty"`
	Languages155 struct {
		Ber string `json:"ber"`
		Mey string `json:"mey"`
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Demonyms8 struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
	} `json:"demonyms,omitempty"`
	Flags34 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms15 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies156 struct {
		Bnd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"BND"`
		Sgd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"SGD"`
	} `json:"currencies,omitempty"`
	Languages156 struct {
		Msa string `json:"msa"`
	} `json:"languages,omitempty"`
	Currencies157 struct {
		Afn struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"AFN"`
	} `json:"currencies,omitempty"`
	Languages157 struct {
		Prs string `json:"prs"`
		Pus string `json:"pus"`
		Tuk string `json:"tuk"`
	} `json:"languages,omitempty"`
	Currencies158 struct {
		Djf struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"DJF"`
	} `json:"currencies,omitempty"`
	Languages158 struct {
		Ara string `json:"ara"`
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini105 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies159 struct {
		Sdg struct {
			Name string `json:"name"`
		} `json:"SDG"`
	} `json:"currencies,omitempty"`
	Languages159 struct {
		Ara string `json:"ara"`
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Gini106 struct {
		Num2014 float64 `json:"2014"`
	} `json:"gini,omitempty"`
	Currencies160 struct {
		Vnd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"VND"`
	} `json:"currencies,omitempty"`
	Languages160 struct {
		Vie string `json:"vie"`
	} `json:"languages,omitempty"`
	Gini107 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies161 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages161 struct {
		Eng string `json:"eng"`
		Mlt string `json:"mlt"`
	} `json:"languages,omitempty"`
	Gini108 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies162 struct {
		Xcd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XCD"`
	} `json:"currencies,omitempty"`
	Languages162 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Gini109 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	Currencies163 struct {
		Krw struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"KRW"`
	} `json:"currencies,omitempty"`
	Languages163 struct {
		Kor string `json:"kor"`
	} `json:"languages,omitempty"`
	Gini110 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	Currencies164 struct {
		Btn struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"BTN"`
		Inr struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"INR"`
	} `json:"currencies,omitempty"`
	Languages164 struct {
		Dzo string `json:"dzo"`
	} `json:"languages,omitempty"`
	Gini111 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies165 struct {
		Cdf struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"CDF"`
	} `json:"currencies,omitempty"`
	Languages165 struct {
		Fra string `json:"fra"`
		Kon string `json:"kon"`
		Lin string `json:"lin"`
		Lua string `json:"lua"`
		Swa string `json:"swa"`
	} `json:"languages,omitempty"`
	Gini112 struct {
		Num2012 float64 `json:"2012"`
	} `json:"gini,omitempty"`
	Currencies166 struct {
		Huf struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"HUF"`
	} `json:"currencies,omitempty"`
	Languages166 struct {
		Hun string `json:"hun"`
	} `json:"languages,omitempty"`
	Gini113 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies167 struct {
		Omr struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"OMR"`
	} `json:"currencies,omitempty"`
	Languages167 struct {
		Ara string `json:"ara"`
	} `json:"languages,omitempty"`
	Currencies168 struct {
		Rub struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"RUB"`
	} `json:"currencies,omitempty"`
	Languages168 struct {
		Rus string `json:"rus"`
	} `json:"languages,omitempty"`
	Gini114 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies169 struct {
		Cny struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"CNY"`
	} `json:"currencies,omitempty"`
	Languages169 struct {
		Zho string `json:"zho"`
	} `json:"languages,omitempty"`
	Translations4 struct {
		Ara struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ara"`
		Bre struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"bre"`
		Ces struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ces"`
		Cym struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"cym"`
		Deu struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"deu"`
		Est struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"est"`
		Fin struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fin"`
		Fra struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fra"`
		Hrv struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hrv"`
		Hun struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hun"`
		Ita struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ita"`
		Jpn struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"jpn"`
		Kor struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"kor"`
		Nld struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"nld"`
		Per struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"per"`
		Pol struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"pol"`
		Por struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"por"`
		Rus struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"rus"`
		Slk struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"slk"`
		Spa struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"spa"`
		Srp struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"srp"`
		Swe struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"swe"`
		Tur struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"tur"`
		Urd struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"urd"`
	} `json:"translations,omitempty"`
	Gini115 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	Currencies170 struct {
		Kwd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"KWD"`
	} `json:"currencies,omitempty"`
	Languages170 struct {
		Ara string `json:"ara"`
	} `json:"languages,omitempty"`
	Currencies171 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages171 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Flags35 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms16 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies172 struct {
		Sek struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"SEK"`
	} `json:"currencies,omitempty"`
	Languages172 struct {
		Swe string `json:"swe"`
	} `json:"languages,omitempty"`
	Gini116 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies173 struct {
		Bwp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"BWP"`
	} `json:"currencies,omitempty"`
	Languages173 struct {
		Eng string `json:"eng"`
		Tsn string `json:"tsn"`
	} `json:"languages,omitempty"`
	Gini117 struct {
		Num2015 float64 `json:"2015"`
	} `json:"gini,omitempty"`
	Idd0 struct {
	} `json:"idd,omitempty"`
	Languages174 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Demonyms9 struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
	} `json:"demonyms,omitempty"`
	Flags36 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms17 struct {
	} `json:"coatOfArms,omitempty"`
	CapitalInfo2 struct {
	} `json:"capitalInfo,omitempty"`
	Currencies174 struct {
		Kmf struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"KMF"`
	} `json:"currencies,omitempty"`
	Languages175 struct {
		Ara string `json:"ara"`
		Fra string `json:"fra"`
		Zdj string `json:"zdj"`
	} `json:"languages,omitempty"`
	Gini118 struct {
		Num2014 float64 `json:"2014"`
	} `json:"gini,omitempty"`
	Currencies175 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages176 struct {
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini119 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies176 struct {
		Dzd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"DZD"`
	} `json:"currencies,omitempty"`
	Languages177 struct {
		Ara string `json:"ara"`
	} `json:"languages,omitempty"`
	Gini120 struct {
		Num2011 float64 `json:"2011"`
	} `json:"gini,omitempty"`
	Currencies177 struct {
		Htg struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"HTG"`
	} `json:"currencies,omitempty"`
	Languages178 struct {
		Fra string `json:"fra"`
		Hat string `json:"hat"`
	} `json:"languages,omitempty"`
	Gini121 struct {
		Num2012 float64 `json:"2012"`
	} `json:"gini,omitempty"`
	Currencies178 struct {
		Aud struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"AUD"`
	} `json:"currencies,omitempty"`
	Languages179 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Gini122 struct {
		Num2014 float64 `json:"2014"`
	} `json:"gini,omitempty"`
	Currencies179 struct {
		Bgn struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"BGN"`
	} `json:"currencies,omitempty"`
	Languages180 struct {
		Bul string `json:"bul"`
	} `json:"languages,omitempty"`
	Gini123 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies180 struct {
		Rsd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"RSD"`
	} `json:"currencies,omitempty"`
	Languages181 struct {
		Srp string `json:"srp"`
	} `json:"languages,omitempty"`
	Gini124 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies181 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages182 struct {
		Ell string `json:"ell"`
	} `json:"languages,omitempty"`
	Gini125 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies182 struct {
		Xaf struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XAF"`
	} `json:"currencies,omitempty"`
	Languages183 struct {
		Fra string `json:"fra"`
		Kon string `json:"kon"`
		Lin string `json:"lin"`
	} `json:"languages,omitempty"`
	Gini126 struct {
		Num2011 float64 `json:"2011"`
	} `json:"gini,omitempty"`
	CoatOfArms18 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies183 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages184 struct {
		Lav string `json:"lav"`
	} `json:"languages,omitempty"`
	Gini127 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies184 struct {
		Cuc struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"CUC"`
		Cup struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"CUP"`
	} `json:"currencies,omitempty"`
	Languages185 struct {
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Currencies185 struct {
		Nzd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"NZD"`
	} `json:"currencies,omitempty"`
	Languages186 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Flags37 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms19 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies186 struct {
		Shp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"SHP"`
	} `json:"currencies,omitempty"`
	Languages187 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Demonyms10 struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
	} `json:"demonyms,omitempty"`
	Flags38 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms20 struct {
	} `json:"coatOfArms,omitempty"`
	Name0 struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name,omitempty"`
	Idd1 struct {
	} `json:"idd,omitempty"`
	Flags39 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CapitalInfo3 struct {
	} `json:"capitalInfo,omitempty"`
	Currencies187 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages188 struct {
		Eng string `json:"eng"`
		Gle string `json:"gle"`
	} `json:"languages,omitempty"`
	Gini128 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies188 struct {
		Xcd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XCD"`
	} `json:"currencies,omitempty"`
	Languages189 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Flags40 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies189 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages190 struct {
		Fin string `json:"fin"`
		Swe string `json:"swe"`
	} `json:"languages,omitempty"`
	Gini129 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies190 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages191 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini130 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies191 struct {
		Brl struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"BRL"`
	} `json:"currencies,omitempty"`
	Languages192 struct {
		Por string `json:"por"`
	} `json:"languages,omitempty"`
	Gini131 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies192 struct {
		Xof struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XOF"`
	} `json:"currencies,omitempty"`
	Languages193 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini132 struct {
		Num2011 float64 `json:"2011"`
	} `json:"gini,omitempty"`
	Currencies193 struct {
		Cop struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"COP"`
	} `json:"currencies,omitempty"`
	Languages194 struct {
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini133 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies194 struct {
		Xof struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XOF"`
	} `json:"currencies,omitempty"`
	Languages195 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini134 struct {
		Num2015 float64 `json:"2015"`
	} `json:"gini,omitempty"`
	Currencies195 struct {
		Stn struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"STN"`
	} `json:"currencies,omitempty"`
	Languages196 struct {
		Por string `json:"por"`
	} `json:"languages,omitempty"`
	Gini135 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies196 struct {
		Php struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"PHP"`
	} `json:"currencies,omitempty"`
	Languages197 struct {
		Eng string `json:"eng"`
		Fil string `json:"fil"`
	} `json:"languages,omitempty"`
	Gini136 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies197 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages198 struct {
		Sqi string `json:"sqi"`
		Srp string `json:"srp"`
	} `json:"languages,omitempty"`
	Translations5 struct {
		Ara struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ara"`
		Bre struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"bre"`
		Ces struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ces"`
		Cym struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"cym"`
		Deu struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"deu"`
		Est struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"est"`
		Fin struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fin"`
		Fra struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fra"`
		Hrv struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hrv"`
		Hun struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hun"`
		Ita struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ita"`
		Kor struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"kor"`
		Nld struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"nld"`
		Per struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"per"`
		Pol struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"pol"`
		Por struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"por"`
		Rus struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"rus"`
		Slk struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"slk"`
		Spa struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"spa"`
		Srp struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"srp"`
		Swe struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"swe"`
		Tur struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"tur"`
		Urd struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"urd"`
		Zho struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"zho"`
	} `json:"translations,omitempty"`
	Gini137 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Flags41 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies198 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages199 struct {
		Ita string `json:"ita"`
	} `json:"languages,omitempty"`
	Currencies199 struct {
		All struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"ALL"`
	} `json:"currencies,omitempty"`
	Languages200 struct {
		Sqi string `json:"sqi"`
	} `json:"languages,omitempty"`
	Gini138 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies200 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages201 struct {
		Deu string `json:"deu"`
		Fra string `json:"fra"`
		Ltz string `json:"ltz"`
	} `json:"languages,omitempty"`
	Gini139 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies201 struct {
		Nok struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"NOK"`
	} `json:"currencies,omitempty"`
	Languages202 struct {
		Nor string `json:"nor"`
	} `json:"languages,omitempty"`
	Demonyms11 struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
	} `json:"demonyms,omitempty"`
	Flags42 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms21 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies202 struct {
		Srd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"SRD"`
	} `json:"currencies,omitempty"`
	Languages203 struct {
		Nld string `json:"nld"`
	} `json:"languages,omitempty"`
	Currencies203 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages204 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Flags43 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms22 struct {
	} `json:"coatOfArms,omitempty"`
	PostalCode1 struct {
		Format string `json:"format"`
	} `json:"postalCode,omitempty"`
	Currencies204 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages205 struct {
		Bar string `json:"bar"`
	} `json:"languages,omitempty"`
	Gini140 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies205 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages206 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Demonyms12 struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
	} `json:"demonyms,omitempty"`
	Flags44 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms23 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies206 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages207 struct {
		Eng string `json:"eng"`
		Mah string `json:"mah"`
	} `json:"languages,omitempty"`
	Currencies207 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages208 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Flags45 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies208 struct {
		Lrd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"LRD"`
	} `json:"currencies,omitempty"`
	Languages209 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Gini141 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	Currencies209 struct {
		Ars struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"ARS"`
	} `json:"currencies,omitempty"`
	Languages210 struct {
		Grn string `json:"grn"`
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini142 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies210 struct {
		Idr struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"IDR"`
	} `json:"currencies,omitempty"`
	Languages211 struct {
		Ind string `json:"ind"`
	} `json:"languages,omitempty"`
	Gini143 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies211 struct {
		Xof struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XOF"`
	} `json:"currencies,omitempty"`
	Languages212 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini144 struct {
		Num2014 float64 `json:"2014"`
	} `json:"gini,omitempty"`
	Currencies212 struct {
		Xcd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XCD"`
	} `json:"currencies,omitempty"`
	Languages213 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Currencies213 struct {
		Scr struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"SCR"`
	} `json:"currencies,omitempty"`
	Languages214 struct {
		Crs string `json:"crs"`
		Eng string `json:"eng"`
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini145 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies214 struct {
		Qar struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"QAR"`
	} `json:"currencies,omitempty"`
	Languages215 struct {
		Ara string `json:"ara"`
	} `json:"languages,omitempty"`
	Currencies215 struct {
		Cad struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"CAD"`
	} `json:"currencies,omitempty"`
	Languages216 struct {
		Eng string `json:"eng"`
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini146 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies216 struct {
		Hkd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"HKD"`
	} `json:"currencies,omitempty"`
	Languages217 struct {
		Eng string `json:"eng"`
		Zho string `json:"zho"`
	} `json:"languages,omitempty"`
	Translations6 struct {
		Ara struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ara"`
		Bre struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"bre"`
		Ces struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ces"`
		Cym struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"cym"`
		Deu struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"deu"`
		Est struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"est"`
		Fin struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fin"`
		Fra struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fra"`
		Hrv struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hrv"`
		Hun struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hun"`
		Ita struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ita"`
		Jpn struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"jpn"`
		Kor struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"kor"`
		Nld struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"nld"`
		Per struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"per"`
		Pol struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"pol"`
		Por struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"por"`
		Rus struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"rus"`
		Slk struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"slk"`
		Spa struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"spa"`
		Srp struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"srp"`
		Swe struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"swe"`
		Tur struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"tur"`
		Urd struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"urd"`
	} `json:"translations,omitempty"`
	Flags46 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies217 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages218 struct {
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini147 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies218 struct {
		Ils struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"ILS"`
	} `json:"currencies,omitempty"`
	Languages219 struct {
		Ara string `json:"ara"`
		Heb string `json:"heb"`
	} `json:"languages,omitempty"`
	Gini148 struct {
		Num2016 float64 `json:"2016"`
	} `json:"gini,omitempty"`
	Currencies219 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages220 struct {
		Nld string `json:"nld"`
	} `json:"languages,omitempty"`
	Gini149 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies220 struct {
		Kes struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"KES"`
	} `json:"currencies,omitempty"`
	Languages221 struct {
		Eng string `json:"eng"`
		Swa string `json:"swa"`
	} `json:"languages,omitempty"`
	Gini150 struct {
		Num2015 float64 `json:"2015"`
	} `json:"gini,omitempty"`
	Currencies221 struct {
		Ckd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"CKD"`
		Nzd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"NZD"`
	} `json:"currencies,omitempty"`
	Languages222 struct {
		Eng string `json:"eng"`
		Rar string `json:"rar"`
	} `json:"languages,omitempty"`
	Flags47 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies222 struct {
		Xaf struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XAF"`
	} `json:"currencies,omitempty"`
	Languages223 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini151 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies223 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages224 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Demonyms13 struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
	} `json:"demonyms,omitempty"`
	Flags48 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms24 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies224 struct {
		Mnt struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"MNT"`
	} `json:"currencies,omitempty"`
	Languages225 struct {
		Mon string `json:"mon"`
	} `json:"languages,omitempty"`
	Gini152 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies225 struct {
		Xpf struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XPF"`
	} `json:"currencies,omitempty"`
	Languages226 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Flags49 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies226 struct {
		Ttd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"TTD"`
	} `json:"currencies,omitempty"`
	Languages227 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Gini153 struct {
		Num1992 float64 `json:"1992"`
	} `json:"gini,omitempty"`
	Currencies227 struct {
		Yer struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"YER"`
	} `json:"currencies,omitempty"`
	Languages228 struct {
		Ara string `json:"ara"`
	} `json:"languages,omitempty"`
	Gini154 struct {
		Num2014 float64 `json:"2014"`
	} `json:"gini,omitempty"`
	Currencies228 struct {
		Aud struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"AUD"`
	} `json:"currencies,omitempty"`
	Languages229 struct {
		Eng string `json:"eng"`
		Pih string `json:"pih"`
	} `json:"languages,omitempty"`
	Flags50 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms25 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies229 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages230 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Flags51 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms26 struct {
	} `json:"coatOfArms,omitempty"`
	Currencies230 struct {
		Pab struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"PAB"`
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages231 struct {
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini155 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies231 struct {
		Twd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"TWD"`
	} `json:"currencies,omitempty"`
	Languages232 struct {
		Zho string `json:"zho"`
	} `json:"languages,omitempty"`
	Translations7 struct {
		Ara struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ara"`
		Bre struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"bre"`
		Ces struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ces"`
		Cym struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"cym"`
		Deu struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"deu"`
		Est struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"est"`
		Fin struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fin"`
		Fra struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fra"`
		Hrv struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hrv"`
		Hun struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hun"`
		Ita struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ita"`
		Jpn struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"jpn"`
		Kor struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"kor"`
		Nld struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"nld"`
		Per struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"per"`
		Pol struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"pol"`
		Por struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"por"`
		Rus struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"rus"`
		Slk struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"slk"`
		Spa struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"spa"`
		Srp struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"srp"`
		Swe struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"swe"`
		Tur struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"tur"`
		Urd struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"urd"`
	} `json:"translations,omitempty"`
	Flags52 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies232 struct {
		Eur struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"EUR"`
	} `json:"currencies,omitempty"`
	Languages233 struct {
		Ell string `json:"ell"`
		Tur string `json:"tur"`
	} `json:"languages,omitempty"`
	Gini156 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies233 struct {
		Xcd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XCD"`
	} `json:"currencies,omitempty"`
	Languages234 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Currencies234 struct {
		Xcd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XCD"`
	} `json:"currencies,omitempty"`
	Languages235 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Languages236 struct {
		Nor string `json:"nor"`
	} `json:"languages,omitempty"`
	Flags53 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	CoatOfArms27 struct {
	} `json:"coatOfArms,omitempty"`
	CapitalInfo4 struct {
	} `json:"capitalInfo,omitempty"`
	Currencies235 struct {
		Mmk struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"MMK"`
	} `json:"currencies,omitempty"`
	Languages237 struct {
		Mya string `json:"mya"`
	} `json:"languages,omitempty"`
	Gini157 struct {
		Num2017 float64 `json:"2017"`
	} `json:"gini,omitempty"`
	Currencies236 struct {
		Xaf struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"XAF"`
	} `json:"currencies,omitempty"`
	Languages238 struct {
		Ara string `json:"ara"`
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini158 struct {
		Num2011 float64 `json:"2011"`
	} `json:"gini,omitempty"`
	Currencies237 struct {
		Fkp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"FKP"`
	} `json:"currencies,omitempty"`
	Languages239 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Flags54 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
	Currencies238 struct {
		Kzt struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"KZT"`
	} `json:"currencies,omitempty"`
	Languages240 struct {
		Kaz string `json:"kaz"`
		Rus string `json:"rus"`
	} `json:"languages,omitempty"`
	Gini159 struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini,omitempty"`
	Currencies239 struct {
		Gel struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"GEL"`
	} `json:"currencies,omitempty"`
	Languages241 struct {
		Kat string `json:"kat"`
	} `json:"languages,omitempty"`
	Gini160 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies240 struct {
		Usd struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"USD"`
	} `json:"currencies,omitempty"`
	Languages242 struct {
		Eng string `json:"eng"`
	} `json:"languages,omitempty"`
	Gini161 struct {
		Num2013 float64 `json:"2013"`
	} `json:"gini,omitempty"`
	Currencies241 struct {
		Crc struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"CRC"`
	} `json:"currencies,omitempty"`
	Languages243 struct {
		Spa string `json:"spa"`
	} `json:"languages,omitempty"`
	Gini162 struct {
		Num2019 float64 `json:"2019"`
	} `json:"gini,omitempty"`
	Currencies242 struct {
		Mga struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"MGA"`
	} `json:"currencies,omitempty"`
	Languages244 struct {
		Fra string `json:"fra"`
		Mlg string `json:"mlg"`
	} `json:"languages,omitempty"`
	Gini163 struct {
		Num2012 float64 `json:"2012"`
	} `json:"gini,omitempty"`
	Currencies243 struct {
		Gnf struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"GNF"`
	} `json:"currencies,omitempty"`
	Languages245 struct {
		Fra string `json:"fra"`
	} `json:"languages,omitempty"`
	Gini164 struct {
		Num2012 float64 `json:"2012"`
	} `json:"gini,omitempty"`
	Currencies244 struct {
		Gbp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"GBP"`
		Imp struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"IMP"`
	} `json:"currencies,omitempty"`
	Languages246 struct {
		Eng string `json:"eng"`
		Glv string `json:"glv"`
	} `json:"languages,omitempty"`
	Demonyms14 struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
	} `json:"demonyms,omitempty"`
	Flags55 struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags,omitempty"`
}

type Pays []Pays

func NewPays(c *Pays) {
	if c == nil {
		log.Fatal(c)
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	err := config.Db().QueryRow("INSERT INTO pays (Nom, Langue, Capital, Région, sous-région, monnaie) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id;", c.Names, c.Capital, c.Monnaie, c.Languages, c.Région, c.SousRégion).Scan(&c.Id)

	if err != nil {
		log.Fatal(err)
	}
}

func FindPaysById(id int) *Pays {
	var pays Pays

	row := config.Db().QueryRow("SELECT * FROM pays WHERE id = $1;", id)
	err := row.Scan(&pays.Id, &pays.Names, &pays.Capital, &pays.Monnaie, &pays.Languages, &pays.Région, &pays.SousRégion)

	if err != nil {
		log.Fatal(err)
	}

	return &pays
}

func AllPays() *Pays {
	var pays Pays

	rows, err := config.Db().Query("SELECT * FROM pays")

	if err != nil {
		log.Fatal(err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var c Pays

		err := rows.Scan(&c.Id, &c.Names, &c.Capital, &c.Monnaie, &c.Languages, &c.Région, &c.SousRégion)

		if err != nil {
			log.Fatal(err)
		}

		pays = append(pays, c)
	}

	return &pays
}

func UpdatePays(pays *Pays) {
	pays.UpdatedAt = time.Now()

	stmt, err := config.Db().Prepare("UPDATE pays SET Names=$1, Capital=$2, Monnaie=$3, Languages=$4, updated_at=$5 WHERE id=$6;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(pays.Names, pays.Capital, pays.Monnaie, pays.Languages, pays.Région, pays.SousRégion, pays.id)

	if err != nil {
		log.Fatal(err)
	}
}

func DeletePaysById(id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM pays WHERE id=$1;")

	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(id)

	return err
}
