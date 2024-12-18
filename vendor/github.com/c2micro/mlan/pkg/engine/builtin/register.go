package builtin

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/c2micro/mlan/pkg/engine/object"
	"github.com/c2micro/mlan/pkg/engine/storage"
	"github.com/c2micro/mlan/pkg/engine/utils"
)

// регистрация встроенных в язык функций
func Register() {
	// assert: если аргумент false - выход из программы
	storage.BuiltinFunctions["assert"] = object.NewBuiltinFunc("assert", Assert)
	// print: печать без переноса строки
	storage.BuiltinFunctions["print"] = object.NewBuiltinFunc("print", Print)
	// println: печать с переносом строки
	storage.BuiltinFunctions["println"] = object.NewBuiltinFunc("println", Println)
	// is_bool: является ли объект Bool
	storage.BuiltinFunctions["is_bool"] = object.NewBuiltinFunc("is_bool", IsBool)
	// is_dict: является ли объект Dict
	storage.BuiltinFunctions["is_dict"] = object.NewBuiltinFunc("is_dict", IsDict)
	// is_float: является ли объект Float
	storage.BuiltinFunctions["is_float"] = object.NewBuiltinFunc("is_float", IsFloat)
	// is_int: является ли объект Int
	storage.BuiltinFunctions["is_int"] = object.NewBuiltinFunc("is_int", IsInt)
	// is_list: является ли объект List
	storage.BuiltinFunctions["is_list"] = object.NewBuiltinFunc("is_list", IsList)
	// is_null: является ли объект Null
	storage.BuiltinFunctions["is_null"] = object.NewBuiltinFunc("is_null", IsNull)
	// is_str: является ли объект Str
	storage.BuiltinFunctions["is_str"] = object.NewBuiltinFunc("is_str", IsStr)
	// len: получение длины объекта
	storage.BuiltinFunctions["len"] = object.NewBuiltinFunc("len", Len)
	// str_len: получение длины строки (количество рун)
	storage.BuiltinFunctions["str_len"] = object.NewBuiltinFunc("str_len", StrLen)
	// bool: кастинг в Bool
	storage.BuiltinFunctions["bool"] = object.NewBuiltinFunc("bool", Bool)
	// float: кастинг в Float
	storage.BuiltinFunctions["float"] = object.NewBuiltinFunc("float", Float)
	// int: кастинг в Int
	storage.BuiltinFunctions["int"] = object.NewBuiltinFunc("int", Int)
	// str: кастинг в Str
	storage.BuiltinFunctions["str"] = object.NewBuiltinFunc("str", Str)
	// reverse: разворот объекта задом-наперед
	storage.BuiltinFunctions["reverse"] = object.NewBuiltinFunc("reverse", Reverse)
	// chr: получение символа на базе int кода
	storage.BuiltinFunctions["chr"] = object.NewBuiltinFunc("chr", Chr)
	// ord: получение int кода символа
	storage.BuiltinFunctions["ord"] = object.NewBuiltinFunc("ord", Ord)
	// str_chr_at: получение руны в строке по индексу
	storage.BuiltinFunctions["str_chr_at"] = object.NewBuiltinFunc("str_chr_at", StrChrAt)
	// list_pop_index: убираем из списка значение по индексу
	storage.BuiltinFunctions["list_pop_index"] = object.NewBuiltinFunc("list_pop_index", ListPopIndex)
	// dict_pop_key: убираем из списка значение по ключу
	storage.BuiltinFunctions["dict_pop_key"] = object.NewBuiltinFunc("dict_pop_key", DictPopKey)
	// base64_enc: кодирование строки в base64
	storage.BuiltinFunctions["base64_enc"] = object.NewBuiltinFunc("base64_enc", Base64Enc)
	// base64_dec: декодирование строки из base64
	storage.BuiltinFunctions["base64_dec"] = object.NewBuiltinFunc("base64_dec", Base64Dec)
	// base32_enc: кодирование строки в base32
	storage.BuiltinFunctions["base32_enc"] = object.NewBuiltinFunc("base32_enc", Base32Enc)
	// base32_dec: декодирование строки из base32
	storage.BuiltinFunctions["base32_dec"] = object.NewBuiltinFunc("base32_dec", Base32Dec)
	// md5: получение md5 хеша от строки
	storage.BuiltinFunctions["md5"] = object.NewBuiltinFunc("md5", Md5)
	// sha1: получение sha1 хеша от строки
	storage.BuiltinFunctions["sha1"] = object.NewBuiltinFunc("sha1", Sha1)
	// sha256: получение sha256 хеша от строки
	storage.BuiltinFunctions["sha256"] = object.NewBuiltinFunc("sha256", Sha256)
	// gzip: сжатие строки
	storage.BuiltinFunctions["gzip"] = object.NewBuiltinFunc("gzip", Gzip)
	// gunzip: распаковка архива в виде строки
	storage.BuiltinFunctions["gunzip"] = object.NewBuiltinFunc("gunzip", Gunzip)
	// fread: чтение файла с ФС
	storage.BuiltinFunctions["fread"] = object.NewBuiltinFunc("fread", Fread)
	// fwrite: запись данных в файл на ФС
	storage.BuiltinFunctions["fwrite"] = object.NewBuiltinFunc("fwrite", Fwrite)
	// str_split: сплиттинг строки по символу
	storage.BuiltinFunctions["str_split"] = object.NewBuiltinFunc("str_split", StrSplit)
}

func Assert(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Bool); !ok {
		return nil, fmt.Errorf("expecting bool, got %s", args[0].TypeName())
	}
	if !args[0].(*object.Bool).GetValue().(bool) {
		return nil, fmt.Errorf("assertion occured")
	}
	return object.NewNull(), nil
}

func Print(args ...object.Object) (object.Object, error) {
	for _, arg := range args {
		fmt.Print(arg.String())
	}
	return object.NewNull(), nil
}

func Println(args ...object.Object) (object.Object, error) {
	for _, arg := range args {
		fmt.Print(arg.String())
	}
	fmt.Println()
	return object.NewNull(), nil
}

func IsBool(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Bool); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func IsDict(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Dict); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func IsFloat(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Float); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func IsInt(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Int); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func IsList(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.List); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func IsNull(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Null); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func IsStr(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Str); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func Len(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch args[0].(type) {
	case *object.Dict:
		return object.NewInt(int64(len(args[0].(*object.Dict).GetValue().(map[string]object.Object)))), nil
	case *object.List:
		return object.NewInt(int64(len(args[0].(*object.List).GetValue().([]object.Object)))), nil
	case *object.Str:
		return object.NewInt(int64(len(args[0].(*object.Str).GetValue().(string)))), nil
	}
	return nil, fmt.Errorf("unable get len of '%s' type", args[0].TypeName())
}

func StrLen(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch args[0].(type) {
	case *object.Str:
		return object.NewInt(int64(utf8.RuneCountInString(args[0].(*object.Str).GetValue().(string)))), nil
	}
	return nil, fmt.Errorf("unable get str len of '%s' type", args[0].TypeName())
}

func Bool(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch args[0].(type) {
	case *object.Bool:
		return args[0], nil
	case *object.Dict:
		if len(args[0].(*object.Dict).GetValue().(map[string]object.Object)) == 0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	case *object.Float:
		if args[0].(*object.Float).GetValue().(float64) == 0.0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	case *object.Int:
		if args[0].(*object.Int).GetValue().(int64) == 0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	case *object.List:
		if len(args[0].(*object.List).GetValue().([]object.Object)) == 0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	case *object.Null:
		return object.NewBool(false), nil
	case *object.Str:
		if len(args[0].(*object.Str).GetValue().(string)) == 0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func Float(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch args[0].(type) {
	case *object.Bool:
		return object.NewFloat(utils.BoolToFloat(args[0].(*object.Bool).GetValue().(bool))), nil
	case *object.Float:
		return args[0], nil
	case *object.Int:
		return object.NewFloat(utils.IntToFloat(args[0].(*object.Int).GetValue().(int64))), nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func Int(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch args[0].(type) {
	case *object.Bool:
		return object.NewInt(utils.BoolToInt(args[0].(*object.Bool).GetValue().(bool))), nil
	case *object.Float:
		return object.NewInt(utils.FloatToInt(args[0].(*object.Float).GetValue().(float64))), nil
	case *object.Int:
		return args[0], nil
	case *object.Str:
		val, err := strconv.Atoi(args[0].(*object.Str).GetValue().(string))
		if err != nil {
			return nil, fmt.Errorf("unable convert str to int: %v", err)
		}
		return object.NewInt(int64(val)), nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func Str(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch args[0].(type) {
	case *object.Bool:
		return object.NewStr(args[0].(*object.Bool).String()), nil
	case *object.Dict:
		return object.NewStr(args[0].(*object.Dict).String()), nil
	case *object.Float:
		return object.NewStr(args[0].(*object.Float).String()), nil
	case *object.Int:
		return object.NewStr(args[0].(*object.Int).String()), nil
	case *object.List:
		return object.NewStr(args[0].(*object.List).String()), nil
	case *object.Null:
		return object.NewStr(args[0].(*object.Null).String()), nil
	case *object.Str:
		return args[0], nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func Reverse(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch args[0].(type) {
	case *object.List:
		var temp []object.Object
		l := args[0].(*object.List).GetValue().([]object.Object)
		for i := len(l) - 1; i >= 0; i-- {
			temp = append(temp, l[i])
		}
		return object.NewList(temp), nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func Chr(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch args[0].(type) {
	case *object.Bool:
		return object.NewStr(string(rune(utils.BoolToInt(args[0].(*object.Bool).GetValue().(bool))))), nil
	case *object.Int:
		return object.NewStr(string(rune(args[0].(*object.Int).GetValue().(int64)))), nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func Ord(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch args[0].(type) {
	case *object.Str:
		if utf8.RuneCountInString(args[0].(*object.Str).GetValue().(string)) != 1 {
			return nil, fmt.Errorf("str must have only one char")
		}
		r, _ := utf8.DecodeRuneInString(args[0].(*object.Str).GetValue().(string))
		return object.NewInt(int64(r)), nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func StrChrAt(args ...object.Object) (object.Object, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("expecting 2 arguments, got %d", len(args))
	}
	// 1ый аргумент - строка
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	// 2ой аргумент - int/bool
	var idx int64
	switch args[1].(type) {
	case *object.Bool:
		idx = utils.BoolToInt(args[1].(*object.Bool).GetValue().(bool))
	case *object.Int:
		idx = args[1].(*object.Int).GetValue().(int64)
	default:
		return nil, fmt.Errorf("expecting int/bool as 2nd argument, got '%s'", args[1].TypeName())
	}
	// валидируем индекс
	if idx < 0 || idx >= int64(utf8.RuneCountInString(str.GetValue().(string))) {
		return nil, object.ErrIndexOutOfRange
	}
	// ширина руны (в количестве байт)
	w := 0
	// количество рун
	c := 0
	for i := 0; i < len(str.GetValue().(string)); i += w {
		var o rune
		o, w = utf8.DecodeRuneInString(str.GetValue().(string)[i:])
		if c == int(idx) {
			return object.NewStr(string(o)), nil
		}
		c++
	}
	return nil, object.ErrNotImplemented
}

func ListPopIndex(args ...object.Object) (object.Object, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("expecting 2 arguments, got %d", len(args))
	}
	list, ok := args[0].(*object.List)
	if !ok {
		return nil, fmt.Errorf("expecting list as 1st argument, got '%s'", args[0].TypeName())
	}
	idx, ok := args[1].(*object.Int)
	if !ok {
		return nil, fmt.Errorf("expecting int as 2nd argument, got '%s'", args[1].TypeName())
	}
	if idx.GetValue().(int64) < 0 || int(idx.GetValue().(int64)) >= len(list.GetValue().([]object.Object)) {
		return nil, object.ErrIndexOutOfRange
	}
	l := list.GetValue().([]object.Object)
	i := idx.GetValue().(int64)
	return object.NewList(append(l[:i], l[i+1:]...)), nil
}

func DictPopKey(args ...object.Object) (object.Object, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("expecting 2 arguments, got %d", len(args))
	}
	dict, ok := args[0].(*object.Dict)
	if !ok {
		return nil, fmt.Errorf("expecting dict as 1st argument, got '%s'", args[0].TypeName())
	}
	key, ok := args[1].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 2nd argument, got '%s'", args[1].TypeName())
	}
	d := dict.GetValue().(map[string]object.Object)
	delete(d, key.GetValue().(string))
	return object.NewDict(d), nil
}

func Base64Enc(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	return object.NewStr(base64.StdEncoding.EncodeToString([]byte(str.GetValue().(string)))), nil
}

func Base64Dec(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	val, err := base64.StdEncoding.DecodeString(str.GetValue().(string))
	if err != nil {
		return nil, err
	}
	return object.NewStr(string(val)), nil
}

func Base32Enc(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	return object.NewStr(base32.StdEncoding.EncodeToString([]byte(str.GetValue().(string)))), nil
}

func Base32Dec(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	val, err := base32.StdEncoding.DecodeString(str.GetValue().(string))
	if err != nil {
		return nil, err
	}
	return object.NewStr(string(val)), nil
}

func Md5(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	md5sum := md5.Sum([]byte(str.GetValue().(string)))
	return object.NewStr(hex.EncodeToString(md5sum[:])), nil
}

func Sha1(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	sha1sum := sha1.Sum([]byte(str.GetValue().(string)))
	return object.NewStr(hex.EncodeToString(sha1sum[:])), nil
}

func Sha256(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	sha256sum := sha256.Sum256([]byte(str.GetValue().(string)))
	return object.NewStr(hex.EncodeToString(sha256sum[:])), nil
}

func Gzip(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	b := &bytes.Buffer{}
	w := gzip.NewWriter(b)
	_, err := w.Write([]byte(str.GetValue().(string)))
	if err != nil {
		return nil, err
	}
	_ = w.Close()
	return object.NewStr(b.String()), nil
}

func Gunzip(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	r, err := gzip.NewReader(bytes.NewBufferString(str.GetValue().(string)))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = r.Close()
	}()
	res, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return object.NewStr(string(res)), nil
}

func Fread(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	data, err := os.ReadFile(str.GetValue().(string))
	if err != nil {
		return nil, err
	}
	return object.NewStr(string(data)), nil
}

func Fwrite(args ...object.Object) (object.Object, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("expecting 2 arguments, got %d", len(args))
	}
	path, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	data, ok := args[1].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 2nd argument, got '%s'", args[1].TypeName())
	}
	if err := os.WriteFile(path.GetValue().(string), []byte(data.GetValue().(string)), 0640); err != nil {
		return nil, err
	}
	return object.NewNull(), nil
}

func StrSplit(args ...object.Object) (object.Object, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("expecting 2 arguments, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting 1st argument str, got '%s'", args[0].TypeName())
	}
	d, ok := args[1].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting 2nd argument str, got '%s'", args[1].TypeName())
	}
	l := strings.Split(str.GetValue().(string), d.GetValue().(string))
	var temp []object.Object
	for _, v := range l {
		temp = append(temp, object.NewStr(v))
	}
	return object.NewList(temp), nil
}
