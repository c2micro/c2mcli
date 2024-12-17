package engine

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
)

// RegisterBuiltinFunctions регистрация встроенных в язык функций
func RegisterBuiltinFunctions() {
	// assert если аргумент false - выход из программы
	BuiltinFunctions["assert"] = object.NewBuiltinFunc("assert", builtinAssert)
	// print печать без переноса строки
	BuiltinFunctions["print"] = object.NewBuiltinFunc("print", builtinPrint)
	// println печать с переносом строки
	BuiltinFunctions["println"] = object.NewBuiltinFunc("println", builtinPrintln)
	// is_bool является ли объект Bool
	BuiltinFunctions["is_bool"] = object.NewBuiltinFunc("is_bool", builtinIsBool)
	// is_dict является ли объект Dict
	BuiltinFunctions["is_dict"] = object.NewBuiltinFunc("is_dict", builtinIsDict)
	// is_dict является ли объект Float
	BuiltinFunctions["is_float"] = object.NewBuiltinFunc("is_float", builtinIsFloat)
	// is_int является ли объект Int
	BuiltinFunctions["is_int"] = object.NewBuiltinFunc("is_int", builtinIsInt)
	// is_list является ли объект List
	BuiltinFunctions["is_list"] = object.NewBuiltinFunc("is_list", builtinIsList)
	// is_null является ли объект Null
	BuiltinFunctions["is_null"] = object.NewBuiltinFunc("is_null", builtinIsNull)
	// is_str является ли объект Str
	BuiltinFunctions["is_str"] = object.NewBuiltinFunc("is_str", builtinIsStr)
	// len получение длины объекта
	BuiltinFunctions["len"] = object.NewBuiltinFunc("len", builtinLen)
	// str_len получение длины строки (количество рун)
	BuiltinFunctions["str_len"] = object.NewBuiltinFunc("str_len", builtinStrLen)
	// bool кастинг в Bool
	BuiltinFunctions["bool"] = object.NewBuiltinFunc("bool", builtinBool)
	// float кастинг в Float
	BuiltinFunctions["float"] = object.NewBuiltinFunc("float", builtinFloat)
	// int кастинг в Int
	BuiltinFunctions["int"] = object.NewBuiltinFunc("int", builtinInt)
	// str кастинг в Str
	BuiltinFunctions["str"] = object.NewBuiltinFunc("str", builtinStr)
	// reverse разворот объекта задом-наперед
	BuiltinFunctions["reverse"] = object.NewBuiltinFunc("reverse", builtinReverse)
	// chr получение символа на базе int кода
	BuiltinFunctions["chr"] = object.NewBuiltinFunc("chr", builtinChr)
	// ord получение int кода символа
	BuiltinFunctions["ord"] = object.NewBuiltinFunc("ord", builtinOrd)
	// str_chr_at получение руны в строке по индексу
	BuiltinFunctions["str_chr_at"] = object.NewBuiltinFunc("str_chr_at", builtinStrChrAt)
	// list_pop_index убираем из списка значение по индексу
	BuiltinFunctions["list_pop_index"] = object.NewBuiltinFunc("list_pop_index", builtinListPopIndex)
	// dict_pop_key убираем из списка значение по ключу
	BuiltinFunctions["dict_pop_key"] = object.NewBuiltinFunc("dict_pop_key", builtinDictPopKey)
	// base64_enc кодирование строки в base64
	BuiltinFunctions["base64_enc"] = object.NewBuiltinFunc("base64_enc", builtinBase64Enc)
	// base64_dec декодирование строки из base64
	BuiltinFunctions["base64_dec"] = object.NewBuiltinFunc("base64_dec", builtinBase64Dec)
	// base32_enc кодирование строки в base32
	BuiltinFunctions["base32_enc"] = object.NewBuiltinFunc("base32_enc", builtinBase32Enc)
	// base32_dec декодирование строки из base32
	BuiltinFunctions["base32_dec"] = object.NewBuiltinFunc("base32_dec", builtinBase32Dec)
	// md5 получение md5 хеша от строки
	BuiltinFunctions["md5"] = object.NewBuiltinFunc("md5", builtinMd5)
	// sha1 получение sha1 хеша от строки
	BuiltinFunctions["sha1"] = object.NewBuiltinFunc("sha1", builtinSha1)
	// sha256 получение sha256 хеша от строки
	BuiltinFunctions["sha256"] = object.NewBuiltinFunc("sha256", builtinSha256)
	// gzip сжатие строки
	BuiltinFunctions["gzip"] = object.NewBuiltinFunc("gzip", builtinGzip)
	// gunzip распаковка архива в виде строки
	BuiltinFunctions["gunzip"] = object.NewBuiltinFunc("gunzip", builtinGunzip)
	// fread чтение файла с ФС
	BuiltinFunctions["fread"] = object.NewBuiltinFunc("fread", builtinFread)
	// fwrite запись данных в файл на ФС
	BuiltinFunctions["fwrite"] = object.NewBuiltinFunc("fwrite", builtinFwrite)
	// str_split сплиттинг строки по символу
	BuiltinFunctions["str_split"] = object.NewBuiltinFunc("str_split", builtinStrSplit)
}

func builtinAssert(args ...object.Object) (object.Object, error) {
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

func builtinPrint(args ...object.Object) (object.Object, error) {
	for _, arg := range args {
		fmt.Print(arg.String())
	}
	return object.NewNull(), nil
}

func builtinPrintln(args ...object.Object) (object.Object, error) {
	for _, arg := range args {
		fmt.Print(arg.String())
	}
	fmt.Println()
	return object.NewNull(), nil
}

func builtinIsBool(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Bool); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func builtinIsDict(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Dict); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func builtinIsFloat(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Float); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func builtinIsInt(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Int); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func builtinIsList(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.List); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func builtinIsNull(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Null); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func builtinIsStr(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Str); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func builtinLen(args ...object.Object) (object.Object, error) {
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

func builtinStrLen(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch args[0].(type) {
	case *object.Str:
		return object.NewInt(int64(utf8.RuneCountInString(args[0].(*object.Str).GetValue().(string)))), nil
	}
	return nil, fmt.Errorf("unable get str len of '%s' type", args[0].TypeName())
}

func builtinBool(args ...object.Object) (object.Object, error) {
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

func builtinFloat(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch args[0].(type) {
	case *object.Bool:
		return object.NewFloat(boolToFloat(args[0].(*object.Bool).GetValue().(bool))), nil
	case *object.Float:
		return args[0], nil
	case *object.Int:
		return object.NewFloat(intToFloat(args[0].(*object.Int).GetValue().(int64))), nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func builtinInt(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch args[0].(type) {
	case *object.Bool:
		return object.NewInt(boolToInt(args[0].(*object.Bool).GetValue().(bool))), nil
	case *object.Float:
		return object.NewInt(floatToInt(args[0].(*object.Float).GetValue().(float64))), nil
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

func builtinStr(args ...object.Object) (object.Object, error) {
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

func builtinReverse(args ...object.Object) (object.Object, error) {
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

func builtinChr(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch args[0].(type) {
	case *object.Bool:
		return object.NewStr(string(rune(boolToInt(args[0].(*object.Bool).GetValue().(bool))))), nil
	case *object.Int:
		return object.NewStr(string(rune(args[0].(*object.Int).GetValue().(int64)))), nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func builtinOrd(args ...object.Object) (object.Object, error) {
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

func builtinStrChrAt(args ...object.Object) (object.Object, error) {
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
		idx = boolToInt(args[1].(*object.Bool).GetValue().(bool))
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

func builtinListPopIndex(args ...object.Object) (object.Object, error) {
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

func builtinDictPopKey(args ...object.Object) (object.Object, error) {
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

func builtinBase64Enc(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	return object.NewStr(base64.StdEncoding.EncodeToString([]byte(str.GetValue().(string)))), nil
}

func builtinBase64Dec(args ...object.Object) (object.Object, error) {
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

func builtinBase32Enc(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	return object.NewStr(base32.StdEncoding.EncodeToString([]byte(str.GetValue().(string)))), nil
}

func builtinBase32Dec(args ...object.Object) (object.Object, error) {
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

func builtinMd5(args ...object.Object) (object.Object, error) {
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

func builtinSha1(args ...object.Object) (object.Object, error) {
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

func builtinSha256(args ...object.Object) (object.Object, error) {
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

func builtinGzip(args ...object.Object) (object.Object, error) {
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

func builtinGunzip(args ...object.Object) (object.Object, error) {
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

func builtinFread(args ...object.Object) (object.Object, error) {
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

func builtinFwrite(args ...object.Object) (object.Object, error) {
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

func builtinStrSplit(args ...object.Object) (object.Object, error) {
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
