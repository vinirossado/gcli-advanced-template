package slice

import "sort"

func Filter[T any](s []T, f func(T) bool) []T {
    var r []T
    for _, v := range s {
        if f(v) {
            r = append(r, v)
        }
    }
    return r
}

func ForEach[T any](s []T, f func(T)) {
    for _, v := range s {
        f(v)
    }
}

func ForeachOrdered[T any](m map[string]string, f func(string, string)) {
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, k := range keys {
        f(k, m[k])
    }
}

func In[Item comparable](item Item, items []Item) bool {
    for i := 0; i < len(items); i++ {
        if items[i] == item {
            return true
        }
    }
    return false
}

func Some[T any](s []T, f func(T) bool) bool {
    for _, v := range s {
        if f(v) {
            return true
        }
    }
    return false
}

func Map[T any, U any](s []T, f func(T) U) []U {
    var r []U
    for _, v := range s {
        r = append(r, f(v))
    }
    return r
}