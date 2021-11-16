package cheapcash

import "strings"

type restriction struct {
	Key         string
	ReplaceWith string
}

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
