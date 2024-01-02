package apitp

type APIdata struct {
	Meta struct {
		Id        string   `json:"id"`
		Uuid      string   `json:"uuid"`
		Sort      string   `json:"sort"`
		Src       string   `json:"src"`
		Section   string   `json:"section"`
		Stems     []string `json:"stems"`
		Offensive bool     `json:"offensive"`
	} `json:"meta"`
	Hom int `json:"hom,omitempty"`
	Hwi struct {
		Hw  string `json:"hw"`
		Prs []struct {
			Mw    string `json:"mw"`
			Sound struct {
				Audio string `json:"audio"`
				Ref   string `json:"ref"`
				Stat  string `json:"stat"`
			} `json:"sound"`
		} `json:"prs,omitempty"`
	} `json:"hwi"`
	Fl  string `json:"fl"`
	Ins []struct {
		If  string `json:"if"`
		Prs []struct {
			Mw    string `json:"mw"`
			Sound struct {
				Audio string `json:"audio"`
				Ref   string `json:"ref"`
				Stat  string `json:"stat"`
			} `json:"sound"`
		} `json:"prs"`
	} `json:"ins,omitempty"`
	Def []struct {
		Sseq [][][]interface{} `json:"sseq"`
		Vd   string            `json:"vd,omitempty"`
	} `json:"def"`
	Uros []struct {
		Ure string `json:"ure"`
		Prs []struct {
			Mw    string `json:"mw"`
			Sound struct {
				Audio string `json:"audio"`
				Ref   string `json:"ref"`
				Stat  string `json:"stat"`
			} `json:"sound,omitempty"`
		} `json:"prs"`
		Fl string `json:"fl"`
	} `json:"uros,omitempty"`
	Dros []struct {
		Drp string `json:"drp"`
		Def []struct {
			Sseq [][][]interface{} `json:"sseq"`
		} `json:"def"`
		Prs []struct {
			Mw    string `json:"mw"`
			Sound struct {
				Audio string `json:"audio"`
				Ref   string `json:"ref"`
			} `json:"sound"`
		} `json:"prs,omitempty"`
		Vrs []struct {
			Vl string `json:"vl"`
			Va string `json:"va"`
		} `json:"vrs,omitempty"`
	} `json:"dros,omitempty"`
	Usages []struct {
		Pl string          `json:"pl"`
		Pt [][]interface{} `json:"pt"`
	} `json:"usages,omitempty"`
	Et       [][]interface{} `json:"et,omitempty"`
	Date     string          `json:"date"`
	Shortdef []string        `json:"shortdef"`
	Cxs      []struct {
		Cxl   string `json:"cxl"`
		Cxtis []struct {
			Cxt string `json:"cxt"`
		} `json:"cxtis"`
	} `json:"cxs,omitempty"`
	Lbs []string `json:"lbs,omitempty"`
}

type Bs struct {
	Sense struct {
		Sn string     `json:"sn"`
		Dt [][]string `json:"dt"`
	} `json:"sense"`
}

type Sense struct {
	Sn string          `json:"sn"`
	Dt [][]interface{} `json:"dt"`
}

type Sen struct {
	Sn  string   `json:"sn"`
	Sls []string `json:"sls"`
}
