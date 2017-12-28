package gregx

import (
    "regexp"
)

// 正则表达式是否匹配
func IsMatch(pattern string, src []byte) bool {
    match, err := regexp.Match(pattern, src)
    if err != nil {
        return false
    }
    return match
}

func IsMatchString(pattern string, src string) bool {
    return IsMatch(pattern, []byte(src))
}

// 正则替换(全部替换)
func Replace(pattern string, src, replace []byte) ([]byte, error) {
    reg, err := regexp.Compile(pattern)
    if err != nil {
        return src, err
    }
    return reg.ReplaceAll(src, replace), nil
}

// 正则替换(全部替换)，字符串
func ReplaceString(pattern, src, replace string) (string, error) {
    r, e := Replace(pattern, []byte(src), []byte(replace))
    return string(r), e
}