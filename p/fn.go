package p

import (
 "encoding/json"
 "net/http"
)

type Currency string

const (
 USD Currency = "USD"
 CHF Currency = "CHF"
 EUR Currency = "EUR"
 GBP Currency = "GBP"
 JPY Currency = "JPY"
)

type vip struct {
 StartDateTime string `json:"startDateTime"`
 EndDateTime string `json:"endDateTime"`
 Currency Currency `json:"currency"`
 RelatedCurrency []Currency `json:"relatedCurrency"`
 Title string `json:"title"`
 IsClose bool `json:"isClose"`
 IsDelete bool `json:"isDelete"`
}

func GetVipData(w http.ResponseWriter, r *http.Request) {
 if r.Method != "POST" {
  http.Error(w, "not post", http.StatusInternalServerError)
  return
 }
 if err := r.ParseForm(); err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
  return
 }
 symbol := r.FormValue("sy")
 v1 := vip{
  StartDateTime: "2019/02/01 00:00:00",
  EndDateTime: "2019/02/01 02:00:00",
  Currency: USD,
  RelatedCurrency: []Currency{USD, CHF},
  Title: "なんか大変そうな発表, " + symbol,
  IsClose:true,
  IsDelete:true,
 }
 v2 := vip{
  StartDateTime: "2019/02/01 03:00:00",
  EndDateTime: "2019/02/01 05:00:00",
  Currency: USD,
  RelatedCurrency: []Currency{USD, CHF, GBP},
  Title: "なんか大変そうな発表2, " + symbol[3:],
  IsClose:true,
  IsDelete:false,
 }
 v3 := vip{
  StartDateTime: "2019/02/01 22:00:00",
  EndDateTime: "2019/02/01 23:00:00",
  Currency: USD,
  RelatedCurrency: []Currency{EUR, JPY},
  Title: "なんか大変そうな発表3, " + symbol[:3],
  IsClose:false,
  IsDelete:true,
 }
 vp := []vip{v1, v2, v3,}
 res, err := json.Marshal(vp)
 if err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
  return
 }
 w.Header().Set("Content-Type", "application/json")
 w.Write(res)
}