package dbio

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"types"
)

var driver bolt.Driver
var conn bolt.Conn
var defaultConnectionURL string = "bolt://neo4j:quickfoxesdevelop@localhost:7687"

func BeginConnection (connectionURL string) error {

	if connectionURL == "" {connectionURL = defaultConnectionURL}

	driver := bolt.NewDriver()
	if connection, err := driver.OpenNeo(defaultConnectionURL); err != nil {
		panic(err)
		return err
	} else {
		conn = connection
		return nil
	}

}

func RetrieveUserMemories(user string) []types.Memory {
	var memoryList = []types.Memory{}
	if data, _, _, _ := conn.QueryNeoAll("MATCH (n:user {username: {username}}) <- [:SUBMITTED_BY] - (m:Memory) RETURN m.url, m.title, m.timestamp", map[string]interface{} {"username":user}); true {
		for index := range data {
			memoryList = append(memoryList, getMemory(data[index]))
		}
	}

	return memoryList

}

func CheckValidLoginCredentials(username, password string) bool {
	if data, _, _, _ := conn.QueryNeoAll("MATCH (n:user) return n.username, n.password", nil); true {
		for index := range data {
			user := data[index]
			if user[0].(string) == username && user[1].(string) == password {
				return true
			}
		}
	} else {
		println("Something disasterous happened!")
	}
	return false
}

func ShutDownConnection(connectionURL string) error {
	if connectionURL == "" {connectionURL = defaultConnectionURL}
	return conn.Close()
}

// TODO: check for already registered user
func RegisterNewUser(username, password string) bool {
	_, err := conn.ExecNeo("CREATE (n:user {username: {username}, password: {password}})",
		map[string]interface{}{"username": username, "password": password})

	if err != nil {
		println("Something disasterous happened!")
		panic(err)
		return false
	}

	return true
}

func SubmitMemory(memory types.Memory, user string) error {

	_, err := conn.ExecNeo("MATCH (u:user {username: {username}}) CREATE (m:Memory {title: {title}, url: {url}, timestamp: {timestamp}}) - [:SUBMITTED_BY] -> (u)",
		map[string]interface{} {
			"title": memory.Title,
			"url": memory.Url,
			"username": user,
			"timestamp": memory.Submission_timestamp,
		})

	if err != nil {

		println("Something disasterous happened!")
		panic(err)
		return err
	} else {
		return nil
	}

}


/*

	HELPER FUNCTIONS

 */


func getMemory(data []interface{}) types.Memory {
	return types.Memory{
		data[0].(string),
		data[1].(string),
		data[2].(int64),
	}
}
