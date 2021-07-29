package core

import "time"

const sampleTime string = "2006-01-02 15:04:05"

type LocalTime time.Time

// func (lt *LocalTime) MarshaJSON() ([]byte, error)  {
// 	b := make([]byte,0,len(sampleTime)+2)
// 	b = append(b, '"')
// 	if len(lt.String()) >0 {
// 		b= time.Time(*lt).AppendFormat(b,sampleTime)
// 	}
// 	b = append(b, '"')
// 	return b, nil
// }

func (lt LocalTime) MarshalJSON() ([]byte, error)  {
	b := make([]byte,0,len(sampleTime)+2)
	b = append(b, '"')
	if len(lt.String()) >0 {
		b= time.Time(lt).AppendFormat(b,sampleTime)
	}
	b = append(b, '"')
	return b, nil
}

func (lt *LocalTime)UnmarshalJSON(b []byte) error {
	t,err := time.ParseInLocation(`"`+sampleTime+`"`,string(b),time.Local)
	*lt =LocalTime(t)
	return err
}

func (lt LocalTime)String() string  {
	str  := time.Time(lt).Format(sampleTime)
	if  str == "0001-01-01 00:00:00" {
		return ""	
	}
	return str
}