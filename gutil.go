package gutils
import (
    "math/rand"
    "time"
    "crypto/md5"
    "encoding/hex"
)

func getTimeNow() string{
   var timeLayout = "2006-01-02 15:04:05.0000" 
    a:=time.Now().Format(timeLayout)
    return a
}
func getNowDateTime(pat string)string{
    timer := time.Now()
    var timeLayout = "2006-01-02 15:04:05"
    switch(pat){
    case "YYYY-MM-DD HH:mm:ss":return timer.Format(timeLayout);
        case "YYYY-MM-DD": timeLayout = "2006-01-02";return timer.Format(timeLayout);
        case "YYYY.MM.DD": timeLayout = "2006.01.02";return timer.Format(timeLayout);
        case "HH:mm:ss": timeLayout = "15:04:05";return timer.Format(timeLayout);
    
    default: return timer.Format(timeLayout);
    }
}
func s2md5(str string) string  {
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}

func s2md5V2(str string) string {
    data := []byte(str)
    has := md5.Sum(data)
    md5str := fmt.Sprintf("%x", has)
    return md5str
}



func getRandom(MIN int, MAX int) int {
	rand1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	return MIN + rand1.Intn(MAX-MIN)
}

func genCode(format string, cha int) string { //9,15,35
	alphabet := "1234567890abcdefghijklmnopqrstuvwxyz"
	FMT := "XXXXX-XXXXX-XXXXX"
	CHA := 15
	if format != "" {
		FMT = format

	}
	if cha != 0 {
		CHA = cha
	}
	c := ""
	section := strings.Split(FMT, "-")
	for _, v := range section {
		for i := 0; i < len(v); i++ {
			d := (getRandom(0, CHA+1))
			b := alphabet[d : d+1]
			//fmt.Println(b)
			c += b
		}
		c = c + "-"
	}
	return c[0 : len(c)-1]

}
