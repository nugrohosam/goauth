package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"unicode"

	"github.com/gorilla/sessions"
	resource "github.com/nugrohosam/gosampleapi/services/http/resources/v1"
	viper "github.com/spf13/viper"
	redisStore "gopkg.in/boj/redistore.v1"
)

// MaxDepth ...
const MaxDepth = 32

// SetAuth ..
func SetAuth(auth *Auth) {
	AuthUser = auth
}

// GetAuth ..
func GetAuth() *Auth {
	return AuthUser
}

// MergeMap ...
func MergeMap(dst, src map[string]interface{}) map[string]interface{} {
	return merge(dst, src, 0)
}

func merge(dst, src map[string]interface{}, depth int) map[string]interface{} {
	if depth > MaxDepth {
		panic("too deep!")
	}
	for key, srcVal := range src {
		if dstVal, ok := dst[key]; ok {
			srcMap, srcMapOk := mapify(srcVal)
			dstMap, dstMapOk := mapify(dstVal)
			if srcMapOk && dstMapOk {
				srcVal = merge(dstMap, srcMap, depth+1)
			}
		}
		dst[key] = srcVal
	}
	return dst
}

// UcFirst ..
func UcFirst(s string) string {
	for index, value := range s {
		return string(unicode.ToUpper(value)) + s[index+1:]
	}
	return ""
}

// LcFirst ..
func LcFirst(s string) string {
	for index, value := range s {
		return string(unicode.ToLower(value)) + s[index+1:]
	}
	return ""
}

// GenerateLimitOffset ..
func GenerateLimitOffset(perPage, page string) (string, string) {
	limit := perPage
	pageInt, _ := strconv.Atoi(page)
	perPageInt, _ := strconv.Atoi(perPage)
	offset := perPageInt * pageInt

	return limit, fmt.Sprint(offset)
}

// BuildPaginate ..
func BuildPaginate(perPage string, page string, total int, listItems interface{}, listItemResource interface{}) interface{} {
	perPageInt, _ := strconv.Atoi(perPage)
	pageInt, _ := strconv.Atoi(page)

	return resource.Paginate{
		Items:       listItemResource,
		PerPage:     perPageInt,
		Total:       total,
		CurrentPage: pageInt,
	}
}

// TypeName ..
func TypeName(t reflect.Type) string {
	return t.Name()
}

// SessionRedis ..
func SessionRedis(key string) {
	store, err := redisStore.NewRediStore(10, "tcp", ":6379", "", []byte(key))
	if err != nil {
		panic(err)
	}
	defer store.Close()
}

// Find ..
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// Encrypt ..
func Encrypt(stringToEncrypt string, keyString string) (encryptedString string) {
	//Since the key is in string, we need to convert decode it to bytes
	key, err := hex.DecodeString(keyString)
	if err != nil {
		panic(err)
	}

	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

// Decrypt ..
func Decrypt(encryptedString string, keyString string) (decryptedString string) {

	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}

func redisStoreSesssion() *redisStore.RediStore {
	redisKey := viper.GetString("reids.key")
	store, err := redisStore.NewRediStore(10, "tcp", ":6379", "", []byte(redisKey))
	if err != nil {
		panic(err)
	}

	return store
}

// StoreSessionString ..
func StoreSessionString(request *http.Request, writer http.ResponseWriter, nameSession string, data string) {
	if viper.GetString("session.driver") == "redis" {
		store := redisStoreSesssion()
		defer store.Close()
	}

	sessionStore := sessions.NewCookieStore([]byte(viper.GetString("app.key")))
	sessionNow, err := sessionStore.Get(request, nameSession)
	if err != nil {
		panic(err)
	}

	sessionNow.Values["data"] = data
	sessionNow.Save(request, writer)
}

// GetSessionDataString ..
func GetSessionDataString(request *http.Request, writer http.ResponseWriter, nameSession string) string {
	if viper.GetString("session.driver") == "redis" {
		store := redisStoreSesssion()
		defer store.Close()
	}

	sessionStore := sessions.NewCookieStore([]byte(viper.GetString("app.key")))
	sessionNow, err := sessionStore.Get(request, nameSession)
	if err != nil {
		return ""
	}

	return sessionNow.Values["data"].(string)
}

// DeleteSessionDataString ..
func DeleteSessionDataString(request *http.Request, writer http.ResponseWriter, nameSession string) error {
	if viper.GetString("session.driver") == "redis" {
		store := redisStoreSesssion()
		defer store.Close()
	}

	sessionStore := sessions.NewCookieStore([]byte(viper.GetString("app.key")))
	sessionNow, err := sessionStore.Get(request, nameSession)
	if err != nil {
		panic(err)
	}

	sessionNow.Options.MaxAge = -1
	err = sessionNow.Save(request, writer)

	return err
}

// GetSessionData ..
func GetSessionData(request *http.Request, writer http.ResponseWriter, nameSession string) interface{} {
	sessionStore := sessions.NewCookieStore([]byte(viper.GetString("app.key")))
	sessionNow, err := sessionStore.Get(request, nameSession)
	if err != nil {
		return ""
	}

	return sessionNow.Values["data"]
}

func mapify(i interface{}) (map[string]interface{}, bool) {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		m := map[string]interface{}{}
		for _, k := range value.MapKeys() {
			m[k.String()] = value.MapIndex(k).Interface()
		}
		return m, true
	}
	return map[string]interface{}{}, false
}
