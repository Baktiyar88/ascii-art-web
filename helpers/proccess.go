package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// isValidFile проверяет, действителен ли файл, сравнивая его хэш SHA-256 с ожидаемым значением.
// Возвращает bool (валидность файла) и error (если произошла ошибка).
func isValidFile(str string) (bool, error) {
	data, err := os.ReadFile(str)
	if err != nil {
		return false, fmt.Errorf("failed to read file: %w", err)
	}

	hasher := sha256.New()
	hasher.Write(data)
	hashCalculated := hex.EncodeToString(hasher.Sum(nil))

	hashDeclaredMap := map[string]bool{
		"e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf": true,
		"242fdef5fad0fe9302bad1e38f0af4b0f83d086e49a4a99cdf0e765785640666": true,
		"78ccd616680eb9068fe1465db1c852ceaffd8c0f318e3aa0414e1635508e85bf": true,
	}

	if _, ok := hashDeclaredMap[hashCalculated]; !ok {
		return false, nil
	}
	return true, nil
}

// getBannerContent читает файл баннера и возвращает его содержимое в виде среза строк.
// Возвращает []string (содержимое файла) и error (если произошла ошибка).
func getBannerContent(sFileName string) ([]string, error) {
	data, err := os.ReadFile(sFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read banner file: %w", err)
	}
	sReturn := strings.Split(string(data), "\n")
	return sReturn, nil
}

// DrawAscii преобразует строку в ASCII-арт, используя указанный файл баннера.
// Возвращает string (результат ASCII-арта) и error (если произошла ошибка).
func DrawAscii(str, sBanner string) (string, error) {
	// Проверяем целостность файла баннера
	if len(sBanner) > 0 {
		sBanner = "./banners/" + sBanner + ".txt"
	} else {
		sBanner = "./banners/standard.txt"
	}

	valid, err := isValidFile(sBanner)
	if err != nil {
		return "", fmt.Errorf("error validating banner file: %w", err)
	}
	if !valid {
		return "", fmt.Errorf("banner file is corrupted")
	}

	sBannerContent, err := getBannerContent(sBanner)
	if err != nil {
		return "", fmt.Errorf("error reading banner file: %w", err)
	}

	for i, line := range sBannerContent {
		sBannerContent[i] = strings.ReplaceAll(line, "\r", "") // Удаляем символ возврата каретки
	}

	strs := strings.ReplaceAll(str, "\\n", "\n")
	sContent := strings.Split(strs, "\n")
	var sBuilder strings.Builder

	// Проверяем, все ли строки пусты
	// allEmpty := true
	// for _, s := range sContent {
	// 	if s != "" {
	// 		allEmpty = false
	// 		break
	// 	}
	// }

	// 	if allEmpty {
	// 		for i := 0; i < len(sContent)-1; i++ {
	// 			sBuilder.WriteString("\n")
	// 		}
	// 	} else {
	// 		for _, wordInput := range sContent {
	// 			if len(wordInput) == 0 {
	// 				sBuilder.WriteString("\n")
	// 				continue
	// 			}
	// 			for h := 1; h < 9; h++ {
	// 				bFlag := false
	// 				for _, char := range wordInput {
	// 					for rowIndex, bannerLine := range sBannerContent {
	// 						if rowIndex == (int(char)-32)*9+h {
	// 							bFlag = true
	// 							sBuilder.WriteString(bannerLine)
	// 						}
	// 					}
	// 				}
	// 				if bFlag {
	// 					sBuilder.WriteString("\n")
	// 				}
	// 			}
	// 		}
	// 	}
	// 	return sBuilder.String(), nil
	for _, wordInput := range sContent {
		if len(wordInput) == 0 {
			sBuilder.WriteString("\n")
			continue
		}
		for h := 1; h < 9; h++ {
			for _, char := range wordInput {
				for rowIndex, bannerLine := range sBannerContent {
					if rowIndex == (int(char)-32)*9+h {
						sBuilder.WriteString(bannerLine)
					}
				}
			}
			sBuilder.WriteString("\n")
		}
	}
	return sBuilder.String(), nil
}

func AsciiValid(str string) bool {
	for _, c := range str {
		if (32 <= c && c <= 126) || c == '\r' || c == '\n' {
			continue
		}
		return false
	}
	return true
}
