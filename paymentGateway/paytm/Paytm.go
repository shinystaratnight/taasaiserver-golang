package paytm

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Paytm struct {
}

const (
	PaytmMerchantKey     = `tYU5n!ObgphLvA80`
	MID                  = `tstUYr79743442324211`
	INDUSTRY_TYPE_ID     = `Retail`
	CHANNEL_ID           = `WAP`
	WEBSITE              = `APPSTAGING`
	CALLBACK_URL         = `https://securegw-stage.paytm.in/theia/paytmCallback?ORDER_ID=`
	TransactionStatusAPI = `https://securegw-stage.paytm.in/merchant-status/getTxnStatus`
)

func (p *Paytm) GetChecksum(paramsMap map[string]string) (checksum string, err error) {
	var keys = make([]string, 0, 0)
	for k, v := range paramsMap {
		if v != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	var arrayList = make([]string, 0, 0)
	for _, key := range keys {
		if value, ok := paramsMap[key]; ok && value != "" {
			arrayList = append(arrayList, value)
		}
	}
	arrayStr := getArray2Str(arrayList)
	salt := generateSalt(4)
	finalString := arrayStr + "|" + salt
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(finalString)))
	hashString := hash + salt
	crypt, err := Encrypt([]byte(hashString))
	if err != nil {
		return
	}
	checksum = base64.StdEncoding.EncodeToString(crypt)
	return
}

func getArray2Str(arrayList []string) (str string) {
	findme := "REFUND"
	findmepipe := "|"
	flag := 1
	for _, v := range arrayList {
		pos := strings.Index(v, findme)
		pospipe := strings.Index(v, findmepipe)
		if pos != -1 || pospipe != -1 {
			continue
		}
		if flag > 0 {
			str += strings.TrimSpace(v)
			flag = 0
		} else {
			str += "|" + strings.TrimSpace(v)
		}
	}
	return
}

func generateSalt(length int) (salt string) {
	rand.Seed(time.Now().UnixNano())
	data := "AbcDE123IJKLMN67QRSTUVWXYZ"
	data += "aBCdefghijklmn123opq45rs67tuv89wxyz"
	data += "0FGH45OP89"
	for i := 0; i < length; i++ {
		salt += string(data[int(rand.Int()%len(data))])
	}
	return
}

func Encrypt(input []byte) (output []byte, err error) {
	iv := "@@@@&&&&####$$$$"
	crypter, _ := NewCrypter([]byte(PaytmMerchantKey), []byte(iv))
	output, err = crypter.Encrypt(input)
	return
}

func Decrypt(input []byte) (output []byte, err error) {
	iv := "@@@@&&&&####$$$$"
	crypter, err := NewCrypter([]byte(PaytmMerchantKey), []byte(iv))
	output, err = crypter.Decrypt(input)
	return
}

type TransactionStatus struct {
	TXNID       string `json:"TXNID"`
	BANKTXNID   string `json:"BANKTXNID"`
	ORDERID     string `json:"ORDERID"`
	TXNAMOUNT   string `json:"TXNAMOUNT"`
	STATUS      string `json:"STATUS"`
	TXNTYPE     string `json:"TXNTYPE"`
	GATEWAYNAME string `json:"GATEWAYNAME"`
	RESPCODE    string `json:"RESPCODE"`
	RESPMSG     string `json:"RESPMSG"`
	BANKNAME    string `json:"BANKNAME"`
	MID         string `json:"MID"`
	PAYMENTMODE string `json:"PAYMENTMODE"`
	REFUNDAMT   string `json:"REFUNDAMT"`
	TXNDATE     string `json:"TXNDATE"`
}

func GetTransactionStatus(orderId string, checksum string) (success bool, err error) {
	var (
		req  *http.Request
		resp *http.Response
		body []byte
	)

	jsonStr := fmt.Sprintf(`{"MID":"%s","ORDERID":"%s","CHECKSUMHASH":"%s"}`, MID, orderId, checksum)

	req, err = http.NewRequest("POST", TransactionStatusAPI, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cache-Control", "no-cache")
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var txnStatus TransactionStatus
	if err = json.Unmarshal(body, &txnStatus); err != nil {
		return false, err
	}
	if txnStatus.STATUS == "TXN_SUCCESS" {
		return true, nil
	}
	return false, err
}
