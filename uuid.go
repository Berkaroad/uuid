/*
import (
	"github.com/berkaroad/uuid"
)

func main(){
	id := uuid.New()
	println(id.String())

	idString := "8eb2a95c-846b-11e5-8550-8bf2f1cec1ce"
	if thisID, err := uuid.Parse(idString); err == nil {
	    println(thisID.String())
	}
}

*/
package uuid

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

var locker *sync.Mutex
var mBuffer [16]byte
var uuidRegex *regexp.Regexp = regexp.MustCompile(`^\{?([a-fA-F0-9]{8})-?([a-fA-F0-9]{4})-?([a-fA-F0-9]{4})-?([a-fA-F0-9]{4})-?([a-fA-F0-9]{12})\}?$`)
var emptyUUID = UUID{}

func init() {
	locker = &sync.Mutex{}
	hostname, _ := os.Hostname()
	mBuffer = md5.Sum([]byte(hostname))
}

type UUID [16]byte

// 创建新的UUID
func New() UUID {
	defer locker.Unlock()
	locker.Lock()

	var uuid [16]byte
	now := time.Now().UTC().UnixNano()
	rand.Seed(now)
	// Timestamp
	binary.BigEndian.PutUint64(uuid[0:8], uint64(now))
	// Machine ID
	copy(uuid[8:12], mBuffer[6:10])
	// Random
	binary.BigEndian.PutUint32(uuid[12:16], uint32(rand.Int63()))
	return uuid
}

// 解析UUID字符串
func Parse(uuidString string) (UUID, error) {
	var uuid [16]byte
	if uuidString == "" {
		return uuid, errors.New("Empty string")
	}

	parts := uuidRegex.FindStringSubmatch(uuidString)
	if parts == nil {
		return uuid, errors.New("Invalid string format")
	}

	slice, _ := hex.DecodeString(strings.Join(parts[1:], ""))
	copy(uuid[:], slice)
	return uuid, nil
}

// 放入到字节分片中
func PutToBytes(data []byte, uuid UUID) {
	data[0] = uuid[0]
	data[1] = uuid[1]
	data[2] = uuid[2]
	data[3] = uuid[3]
	data[4] = uuid[4]
	data[5] = uuid[5]
	data[6] = uuid[6]
	data[7] = uuid[7]
	data[8] = uuid[8]
	data[9] = uuid[9]
	data[10] = uuid[10]
	data[11] = uuid[11]
	data[12] = uuid[12]
	data[13] = uuid[13]
	data[14] = uuid[14]
	data[15] = uuid[15]
}

// 从字节分片中加载
func LoadFromBytes(data []byte) UUID {
	var uuid [16]byte
	if len(data) == 16 {
		uuid[0] = data[0]
		uuid[1] = data[1]
		uuid[2] = data[2]
		uuid[3] = data[3]
		uuid[4] = data[4]
		uuid[5] = data[5]
		uuid[6] = data[6]
		uuid[7] = data[7]
		uuid[8] = data[8]
		uuid[9] = data[9]
		uuid[10] = data[10]
		uuid[11] = data[11]
		uuid[12] = data[12]
		uuid[13] = data[13]
		uuid[14] = data[14]
		uuid[15] = data[15]
	}
	return uuid
}

// 是否为空值
func IsEmpty(uuid UUID) bool {
	return uuid == emptyUUID
}

// UUID的字符串形式
func (self UUID) String() string {
	uuidString := fmt.Sprintf("%x-%x-%x-%x-%x", self[0:4], self[4:6], self[6:8], self[8:10], self[10:16])
	return uuidString
}
