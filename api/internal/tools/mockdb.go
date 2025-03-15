package tools
import (
	"time"
)
type mockDB struct {}
var mockLoginDetails = map[string]LoginDetails{
	"alex":{
		AuthToken :"123ABC",
		Username :"alex",
},
"jason":{
		AuthToken :"456DEF",
		Username :"jason",
},
"marie":{
	AuthToken :"789GHI",
	Username :"marie",
},

}
var mockCoinDetails = map[string]CoinDetails{
	"alex":CoinDetails{
		Coins:10,
		Username:"alex",
	},
	"jason":CoinDetails{
		Coins:20,
		Username:"jason",
	},
	"marie":CoinDetails{
		Coins:30,
		Username:"marie",
	},

}
func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails{
	time.Sleep(1*time.Second)
	var clientData = LoginDetails{}
	clientData ,ok := mockLoginDetails[username]
	if !ok{
		return nil
	}
	return &clientData
}
func (d *mockDB) GetUserCoins(username string) *CoinDetails{
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{} 
	clientData ,ok := mockCoinDetails[username]
	if !ok{
		return nil
	}
	return &clientData
}
func (d *mockDB) SetUpDatabase()error{
	return nil
}