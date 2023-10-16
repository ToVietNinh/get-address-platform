package main

import (
	"regexp"
	"strings"
)

func removeVietnameseDiacritics(input string) string {
	// Create a map to replace diacritic characters with their corresponding non-diacritic characters
	diacriticsMap := map[rune]rune{
		'á': 'a', 'à': 'a', 'ả': 'a', 'ã': 'a', 'ạ': 'a',
		'ă': 'a', 'ắ': 'a', 'ằ': 'a', 'ẳ': 'a', 'ẵ': 'a', 'ặ': 'a',
		'â': 'a', 'ấ': 'a', 'ầ': 'a', 'ẩ': 'a', 'ẫ': 'a', 'ậ': 'a',
		'é': 'e', 'è': 'e', 'ẻ': 'e', 'ẽ': 'e', 'ẹ': 'e',
		'ê': 'e', 'ế': 'e', 'ề': 'e', 'ể': 'e', 'ễ': 'e', 'ệ': 'e',
		'í': 'i', 'ì': 'i', 'ỉ': 'i', 'ĩ': 'i', 'ị': 'i',
		'ó': 'o', 'ò': 'o', 'ỏ': 'o', 'õ': 'o', 'ọ': 'o',
		'ô': 'o', 'ố': 'o', 'ồ': 'o', 'ổ': 'o', 'ỗ': 'o', 'ộ': 'o',
		'ơ': 'o', 'ớ': 'o', 'ờ': 'o', 'ở': 'o', 'ỡ': 'o', 'ợ': 'o',
		'ú': 'u', 'ù': 'u', 'ủ': 'u', 'ũ': 'u', 'ụ': 'u',
		'ư': 'u', 'ứ': 'u', 'ừ': 'u', 'ử': 'u', 'ữ': 'u', 'ự': 'u',
		'ý': 'y', 'ỳ': 'y', 'ỷ': 'y', 'ỹ': 'y', 'ỵ': 'y',
		'đ': 'd',
	}

	// Create a function to replace diacritic characters
	replaceDiacritic := func(r rune) rune {
		if val, ok := diacriticsMap[r]; ok {
			return val
		}
		return r
	}

	// Use strings.Map to replace diacritic characters in the input string
	return strings.Map(replaceDiacritic, input)
}

func standardizeProvinceName(name string) string {
	// Define a regular expression pattern for province names
	regex := regexp.MustCompile(`(Tỉnh|Thành phố|TP|Tp|tp)\s+`)

	// Replace matched patterns with an empty string
	standardizedName := regex.ReplaceAllString(name, "")

	standardizedName = strings.ToLower(standardizedName)

	standardizedName = removeVietnameseDiacritics(standardizedName)

	return standardizedName
}

func standardizeDistrictName(name string) string {
	// Define a regular expression pattern for district names
	regex := regexp.MustCompile(`(Huyện|Quận|Thị xã|TX|tx|Tx)\s+`)

	// Replace matched patterns with an empty string
	standardizedName := regex.ReplaceAllString(name, "")

	standardizedName = strings.ToLower(standardizedName)

	standardizedName = removeVietnameseDiacritics(standardizedName)

	return standardizedName
}

func standardizeWardName(name string) string {
	// Define a regular expression pattern for ward names
	regex := regexp.MustCompile(`(Phường|Xã|P|p|X|x|XP|xp|Xp)\s+`)

	// Replace matched patterns with an empty string
	standardizedName := regex.ReplaceAllString(name, "")

	standardizedName = strings.ToLower(standardizedName)

	standardizedName = removeVietnameseDiacritics(standardizedName)

	return standardizedName
}
