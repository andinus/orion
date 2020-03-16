package hibp

import (
	"fmt"
	"strings"
)

// GetPwned takes SHA-1 Hash as input & returns Pwned Passwords list
// returned by the Have I Been Pwned API
func GetPwned(hsh string) (map[string]string, error) {
	api := "https://api.pwnedpasswords.com/range"
	list := make(map[string]string)

	pfx := hsh[:5]

	reqApi := fmt.Sprintf("%s/%s", api, pfx)
	body, err := reqHIBP(reqApi)
	if err != nil {
		return list, fmt.Errorf("reqHIBP failed\n%s",
			err.Error())
	}

	for _, v := range strings.Split(body, "\r\n") {
		s := strings.Split(v, ":")
		list[s[0]] = s[1]
	}
	return list, err
}

// ChkPwn takes list, hash as input & returns if the hash is in list,
// the frequency
func ChkPwn(list map[string]string, hsh string) (bool, string) {
	sfx := hsh[5:]
	for k, fq := range list {
		if sfx == k {
			return true, fq
		}
	}
	return false, ""
}
