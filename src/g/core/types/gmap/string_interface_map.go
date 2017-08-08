package gmap

import (
	"sync"
)

type StringInterfaceMap struct {
	m sync.RWMutex
	M map[string]interface{}
}

func NewStringInterfaceMap() *StringInterfaceMap {
	return &StringInterfaceMap{
		M: make(map[string]interface{}),
	}
}

// 哈希表克隆
func (this *StringInterfaceMap) Clone() *map[string]interface{} {
    m := make(map[string]interface{})
    this.m.RLock()
    for k, v := range this.M {
        m[k] = v
    }
    this.m.RUnlock()
    return &m
}

// 设置键值对
func (this *StringInterfaceMap) Set(key string, val interface{}) {
	this.m.Lock()
	this.M[key] = val
	this.m.Unlock()
}

// 批量设置键值对
func (this *StringInterfaceMap) BatchSet(m map[string]interface{}) {
    todo := make(map[string]interface{})
    this.m.RLock()
    for k, v := range m {
        old, exists := this.M[k]
        if exists && v == old {
            continue
        }
        todo[k] = v
    }
    this.m.RUnlock()

    if len(todo) == 0 {
        return
    }

    this.m.Lock()
    for k, v := range todo {
        this.M[k] = v
    }
    this.m.Unlock()
}

// 获取键值
func (this *StringInterfaceMap) Get(key string) interface{} {
	this.m.RLock()
	val, _ := this.M[key]
	this.m.RUnlock()
	return val
}

// 删除键值对
func (this *StringInterfaceMap) Remove(key string) {
	this.m.Lock()
	delete(this.M, key)
	this.m.Unlock()
}

// 批量删除键值对
func (this *StringInterfaceMap) BatchRemove(keys []string) {
    this.m.Lock()
    for _, key := range keys {
        delete(this.M, key)
    }
    this.m.Unlock()
}

// 返回对应的键值，并删除该键值
func (this *StringInterfaceMap) GetAndRemove(key string) interface{} {
	this.m.Lock()
	val, exists := this.M[key]
	if exists {
		delete(this.M, key)
	}
	this.m.Unlock()
	return val
}

// 返回键列表
func (this *StringInterfaceMap) Keys() []string {
	this.m.RLock()
	keys := make([]string, 0)
	for key, _ := range this.M {
		keys = append(keys, key)
	}
    this.m.RUnlock()
	return keys
}

// 返回值列表(注意是随机排序)
func (this *StringInterfaceMap) Values() []interface{} {
	this.m.RLock()
	vals := make([]interface{}, 0)
	for _, val := range this.M {
		vals = append(vals, val)
	}
	this.m.RUnlock()
	return vals
}

// 是否存在某个键
func (this *StringInterfaceMap) Contains(key string) bool {
	this.m.RLock()
	_, exists := this.M[key]
	this.m.RUnlock()
	return exists
}

// 哈希表大小
func (this *StringInterfaceMap) Size() int {
	this.m.RLock()
	len := len(this.M)
	this.m.RUnlock()
	return len
}

// 哈希表是否为空
func (this *StringInterfaceMap) IsEmpty() bool {
	this.m.RLock()
	empty := (len(this.M) == 0)
	this.m.RUnlock()
	return empty
}

// 清空哈希表
func (this *StringInterfaceMap) Clear() {
    this.m.Lock()
    this.M = make(map[string]interface{})
    this.m.Unlock()
}