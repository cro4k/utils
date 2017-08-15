/**
 * Created by 93201 on 2017/5/26.
 */
package utils

import (
	"net/url"
	"sort"
	"strings"
)

func SignatureParam(v url.Values, salt string) string {
	var keys = []string{}
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	str := ""
	for _, key := range keys {
		if key == "signature" || v.Get(key) == "" {
			continue
		}
		str = str + key + "=" + v.Get(key) + "&"
	}
	str = strings.Trim(str, "&")
	switch strings.ToLower(v.Get("sign_method")) {
	case "md5":
		return MD5String([]byte(str + salt))
	default:
		return Sha1String([]byte(str + salt))
	}
}

func VerifySignature(v url.Values, salt string) bool {
	return SignatureParam(v, salt) == v.Get("signature")
}
