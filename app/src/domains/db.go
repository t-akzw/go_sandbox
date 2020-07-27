package domains
import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime/debug"
	"strings"
	"sync"
	"time"
	"github.com/jinzhu/gorm"
	"local.packages/lib"

)
var db = lib.NewDbConnection()
func BeginTransaction() {
	fmt.Println("[TRANSACTION] BEGIN: ", contextName())
	contexts.Mux.Lock()
	defer contexts.Mux.Unlock()
	contexts.Contexts[contextName()] = db.Begin()
}
func RollbackTransactionIfNeeded() {
	if err := recover(); err != nil {
		contexts.Mux.Lock()
		defer contexts.Mux.Unlock()
		Rollback()
		panic(err)
	}
}
func Rollback() {
	if val, ok := contexts.Contexts[contextName()]; ok {
		fmt.Println("[TRANSACTION] ROLLBACK: ", contextName())
		val.Rollback()
		delete(contexts.Contexts, contextName())
	}
}
func CommitTransaction() {
	contexts.Mux.Lock()
	defer contexts.Mux.Unlock()
	if val, ok := contexts.Contexts[contextName()]; ok {
		fmt.Println("[TRANSACTION] COMMIT: ", contextName())
		val.Commit()
		delete(contexts.Contexts, contextName())
	}
}
func contextName() string {
	line := strings.SplitN(string(debug.Stack()), "\n", 2)[0]
	return "goroutine_" + strings.Split(line, " ")[1]
}
type SafeContexts struct {
	Contexts map[string]*gorm.DB
	Mux      sync.Mutex
}
var contexts = SafeContexts{Contexts: map[string]*gorm.DB{}}
func tx() *gorm.DB {
	contexts.Mux.Lock()
	defer contexts.Mux.Unlock()
	if val, ok := contexts.Contexts[contextName()]; ok {
		return val
	} else {
		return db
	}
}
func txSave(value interface{}) *gorm.DB {
	context := tx().Save(value)
	if err := context.Error; err != nil {
		panic(err)
	}
	return context
}
func txSaveWithError(value interface{}) (*gorm.DB, error) {
	context := tx().Save(value)
	if err := context.Error; err != nil {
		return context, err
	}
	return context, nil
}
func GetDbTime() time.Time {
	type Result struct {
		Now time.Time `json:"now"`
	}
	result := Result{}
	tx().Raw("SELECT now() as now").Scan(&result)
	return result.Now
}
func mapping(value interface{}, properties interface{}) {
	domapping(reflect.ValueOf(value), reflect.StructField{}, reflect.ValueOf(properties))
}
func domapping(value reflect.Value, field reflect.StructField, source reflect.Value) (bool, error) {
	if source.Kind() == reflect.Ptr {
		source = source.Elem()
	}
	var vKind = value.Kind()
	if vKind == reflect.Ptr {
		var isNew bool
		vPtr := value
		if value.IsNil() {
			isNew = true
			vPtr = reflect.New(value.Type().Elem())
		}
		isSetted, err := domapping(vPtr.Elem(), field, source)
		if err != nil {
			return false, err
		}
		if isNew && isSetted {
			value.Set(vPtr)
		}
		return isSetted, nil
	}
	if vKind != reflect.Struct || !field.Anonymous {
		ok := tryToSetValue(value, field, source)
		if ok {
			return true, nil
		}
	}
	if vKind == reflect.Struct {
		tValue := value.Type()
		var isSetted bool
		for i := 0; i < value.NumField(); i++ {
			sf := tValue.Field(i)
			if sf.PkgPath != "" && !sf.Anonymous { // unexported
				continue
			}
			ok, err := domapping(value.Field(i), tValue.Field(i), source)
			if err != nil {
				return false, err
			}
			isSetted = isSetted || ok
		}
		return isSetted, nil
	}
	return false, nil
}
func tryToSetValue(value reflect.Value, field reflect.StructField, source reflect.Value) bool {
	sv := source.FieldByName(field.Name)
	if !sv.IsValid() {
		return false
	}
	if field.Type.Kind() == reflect.Ptr {
		return true
	}
	value.Set(sv)
	return true
}
var seedLetters = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
func GenerateRandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = seedLetters[rand.Intn(len(seedLetters))]
	}
	return string(b)
}
type EventSearchConfig struct {
	Offset          uint `form:"offset"`
	Limit           uint `form:"limit"`
	SendFlexMessage bool `form:"sendFlexMessage"`
}