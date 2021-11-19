package cheapcash

import "strings"

type restriction struct {
	Key         string
	ReplaceWith string
}

// Sanitize path that is a valid string when the user
// submitted it, but it's invalid when it's written
// as a file name on local filesystem.
func sanitizePath(path string) string {
	restricted := []restriction{
		{Key: " ", ReplaceWith: "_s_"},
		{Key: "^", ReplaceWith: "_p_"},
		{Key: "*", ReplaceWith: "_a_"},
		{Key: "\"", ReplaceWith: "_dq_"},
		{Key: "'", ReplaceWith: "_sq_"},
		{Key: "?", ReplaceWith: "_qm_"},
		{Key: ">", ReplaceWith: "_gt_"},
		{Key: "<", ReplaceWith: "_lt_"},
	}

	for _, v := range restricted {
		path = strings.ReplaceAll(path, v.Key, v.ReplaceWith)
	}
	return path
}
