package vinpack

type Employee struct {
	Id         int     `xml:"empno"`
	Name       string  `xml:"empname"`
	Salary     float64 `xml:"salary"`
	Department string  `xml:"dept"`
}


