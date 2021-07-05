package common

import (
	"bytes"
	"encoding/json"
	"sort"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	config "github.com/kukkar/mta-hosting-optimizer/tests/conf"

	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const HTTP_REQUEST_GET = "GET"
const HTTP_REQUEST_POST = "POST"
const HTTP_REQUEST_PUT = "PUT"
const HTTP_REQUEST_DELETE = "DELETE"

const UNAUTHORIZED_ACCESS = "Unauthorized access"

const HEADER_CONTENT_TYPE = "Content-Type"
const HEADER_CACHE_CONTROL = "Cache-Control"
const HEADER_AUTHORIZATION = "Authorization"

func HttpRequest(
	method string,
	endPoint string,
	customHeaders map[string]string,
	body []byte,
	responsePlaceholder interface{}) error {

	client := &http.Client{}
	req := &http.Request{}

	switch method {
	case "GET":
		var errR error
		req, errR = http.NewRequest(HTTP_REQUEST_GET, endPoint, nil)
		if errR != nil {
			return errR
		}

	case "PUT":
		var errR error
		// http.POST expects an io.Reader, which a byte buffer does
		bPostContent := bytes.NewBuffer(body)
		req, errR = http.NewRequest(HTTP_REQUEST_PUT, endPoint, bPostContent)
		if errR != nil {
			return errR
		}

	case "POST":
		var errR error
		// http.POST expects an io.Reader, which a byte buffer does
		bPostContent := bytes.NewBuffer(body)
		req, errR = http.NewRequest(HTTP_REQUEST_POST, endPoint, bPostContent)
		if errR != nil {
			return errR
		}

	case "DELETE":
		var errR error
		// http.POST expects an io.Reader, which a byte buffer does
		bPostContent := bytes.NewBuffer(body)
		req, errR = http.NewRequest(HTTP_REQUEST_DELETE, endPoint, bPostContent)
		if errR != nil {
			return errR
		}
	default:
	}

	req.Header.Add(HEADER_CONTENT_TYPE, "application/json")

	// Added custom headers
	for hKey, hVal := range customHeaders {
		req.Header.Add(hKey, hVal)
	}

	// Execute request
	response, errRes := client.Do(req)
	if errRes != nil {
		return errRes
	}

	defer response.Body.Close()

	// Prepare response
	if response.StatusCode == 403 {
		return fmt.Errorf(UNAUTHORIZED_ACCESS)
	}

	//if response.StatusCode == 200 || response.StatusCode == 500 {
	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(response.Body).Decode(&responsePlaceholder); err != nil {
		return err
	}
	if response.StatusCode != 200 {
		errdataBytes, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf("Got response code: %d, Error: %s", response.StatusCode, errdataBytes)
	}
	return nil

}

func GetSqlConn() (*gorm.DB, error) {
	con, err := config.GetConfig()
	if err != nil {
		fmt.Println("Error in fetching configuration")
		return nil, err
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=1&loc=%s",
		con.Mysql.User,
		con.Mysql.Password,
		con.Mysql.Host,
		con.Mysql.Port,
		con.Mysql.DBName,
		url.QueryEscape(con.Mysql.DefaultTimeZone),
	)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(con.Mysql.MaxIdleConnections)
	db.DB().SetMaxOpenConns(con.Mysql.MaxOpenConnections)
	//Check if we have encountered any errors
	if err != nil {
		return nil, err
	}
	//On DB error
	if err = db.Error; err != nil {
		return nil, err
	}
	return db, nil
}

func GetHMacRequest(input map[string]interface{}) (string, error) {
	var data string
	if input == nil {
		return data, fmt.Errorf("request body can not be empty")
	}
	sortedKeys := getSortedKey(input)
	craftedData := seprateValueWithPipe(input, sortedKeys)
	data = craftedData
	return data, nil
}

func getSortedKey(m map[string]interface{}) []string {

	// To store the keys in slice in sorted order
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func seprateValueWithPipe(m map[string]interface{}, sortedKey []string) string {
	var output string
	for _, eachKey := range sortedKey {
		if val, ok := m[eachKey]; ok {
			if val != nil {
				switch v := val.(type) {
				case int:
					output += fmt.Sprintf("%d", v) + "|"
				case int32:
					output += fmt.Sprintf("%d", v) + "|"
				case int64:
					output += fmt.Sprintf("%d", v) + "|"
				case float64:
					output += fmt.Sprintf("%f", v) + "|"
				case string:
					output += fmt.Sprintf("%s", v) + "|"
				case uint64:
					output += fmt.Sprintf("%d", v) + "|"
				case uint32:
					output += fmt.Sprintf("%d", v) + "|"
				case uint:
					output += fmt.Sprintf("%d", v) + "|"
				}
			}
		}
	}

	return strings.TrimRight(output, "|")
}
