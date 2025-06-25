package descussion

type Descuss struct {
	Question  string
	YesAnswer string
	NoAnswer  string
	Jebaited  string
}

type Dialog = []Descuss

var GotHim bool = false
var HeDoesntLoveHisMom bool = false

var DialogFlow = Dialog{
	Descuss{
		Question:  "Привет",
		YesAnswer: "Привет",
		NoAnswer:  "Здарова",
		Jebaited:  "",
	},
	Descuss{
		Question:  "Погода в Санкт-Петербургеу не очень, верно?",
		YesAnswer: "Ну да, ну да",
		NoAnswer:  "Ну нет, ну нет",
		Jebaited:  "",
	},
	Descuss{
		Question:  "Зачарованные топ?",
		YesAnswer: "Конечено",
		NoAnswer:  "Конечено нет",
		Jebaited:  "",
	},
	Descuss{
		Question:  "Маму любишь?",
		YesAnswer: "Да!",
		NoAnswer:  "Нет",
		Jebaited:  "Сосал?",
	},
}
