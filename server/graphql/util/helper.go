// Package util impls functions as utilities for Graphql API
package util

// StringPointerHelper return value of string iff string pointer is not nil
// else return empty string
func StringPointerHelper(a *string) string {
	if a != nil {
		return *a
	}
	return ""
}
