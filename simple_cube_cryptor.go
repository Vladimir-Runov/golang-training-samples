package main

/* base6464 chunks (encrypted data)
random text 60 bt. 1 chunks...펅꼫톼칽췻펍슈픠뿭됝곐쉱쬹콿붕껒볌뫽팂덝
byte array:
[237 142 133 234 188 171 237 134 188 236 185 189 236 183 187 237 142 141 236 138 136 237 148 160 235 191 173 235 144 157 234 179 144 236 137 177 236 172 185 236 189 191 235 182 149 234 187 146 235 179 140 235 171 189 237 140 130 235 141 157 0 0 0 60]
byte array encrypted:
[163 91 121 59 227 83 41 21 65 125 89 221 85 47 255 211 137 57 155 59 91 88 31 115 83 95 241 125 22 201 109 149 85 25 209 223 83 81 19 157 85 171 75 57 207 102 109 123 159 83 71 123 133 82 17 243 83 123 117 143 115 25 213 73]
as Base64: o1t5O+NTKRVBfVndVS//04k5mztbWB9zU1/xfRbJbZVVGdHfU1ETnVWrSznPZm17n1NHe4VSEfNTe3WPcxnVSQ==
as int64 array:
	0x-5CA486C41CACD6EB		0x417D59DD552FFFD3		0x-76C664C4A4A7E08D		0x535FF17D16C96D95
	0x5519D1DF5351139D		0x55AB4B39CF666D7B		0x-60ACB8847AADEE0D		0x537B758F7319D549
byte array (decrypted):
[237 142 133 234 188 171 237 134 188 236 185 189 236 183 187 237 142 141 236 138 136 237 148 160 235 191 173 235 144 157 234 179 144 236 137 177 236 172 185 236 189 191 235 182 149 234 187 146 235 179 140 235 171 189 237 140 130 235 141 157 0 0 0 60]

펅꼫톼칽췻펍슈픠뿭됝곐쉱쬹콿붕껒볌뫽팂덝
Ок
*/
import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"math/rand"
)

type TextPart struct {
	Index int
	Str   string
}
type DataPart struct {
	Index int
	data  [64]byte
}

const (
	chunkSize     = 63
	cps_scrollMin = 7    // параметр шифрования, мин. сдвиг данных (1...256)
	cps_scrollDiv = 16   // остаток от деления, расч. количество сдвигов буффера (1...99)
	cps_matrixA   = true // параметр шифрования, порядок обхоода массива (x*8+y или x + y*8 )

)

func getBit(data64 *[]byte, byteNo int, bitNo int) byte {
	return ((*data64)[byteNo] >> bitNo) & 1
}

func setBit(data64 *[]byte, byteNo int, bitNo int, newVal byte) byte {
	//	mutex.Lock()         // Блокируем доступ к массиву
	//	defer mutex.Unlock() // Разблокируем в конце функции

	bitn := ((*data64)[byteNo] >> bitNo) & 1
	if bitn != newVal {
		if newVal != 0 {
			(*data64)[byteNo] |= (1 << bitNo) // Устанавливаем бит в 1
		} else {
			(*data64)[byteNo] &^= (1 << bitNo) // Устанавливаем бит в 0
		}
	}
	return bitn
}

// GetBit возвращает true если 1= значение X-го бита в N-м байте массива байт
func GetBitVal(data []byte, N int, X int) bool {
	byteValue := data[N]
	return (byteValue & (1 << X)) != 0
}

func Bytes2Long64Array(data64 []byte) [8]int64 {
	var long64Array [8]int64
	for i := 0; i < 8; i++ {
		long64Array[i] = int64(binary.BigEndian.Uint64(data64[i*8 : (i+1)*8]))
	}
	return long64Array
}

func Long64Array2Bytes64(long64Array [8]int64) []byte {
	data64 := make([]byte, 64) // Создаем массив байт длиной 64
	for i := 0; i < 8; i++ {
		binary.BigEndian.PutUint64(data64[i*8:(i+1)*8], uint64(long64Array[i]))
	}
	return data64
}

func String2Bytes64(inputStr string) []byte {
	bytesOfstr := []byte(inputStr)
	data := make([]byte, 64)

	for i := range data {
		data[i] = byte(rand.Intn(256)) // Генерируем случайное число от 0 до 255
	}
	lenStr := len(bytesOfstr)
	if lenStr > 63 {
		copy(data[:63], bytesOfstr[:63])
		data[63] = byte(63)
	} else {
		copy(data[:lenStr], bytesOfstr)
		data[63] = byte(lenStr)
	}
	length := data[63]
	str := string(data[:length])
	fmt.Println(str)
	return data
}

func ByteNo(x int, y int) int {
	if cps_matrixA {
		return x*8 + y
	}
	return y*8 + x
}

func isEven(n byte) bool {
	return n%2 == 0
}

// GetBit возвращает значение X-го бита в N-м байте массива байт.
func GetBit(data *[]byte, N int, X int) bool {
	byteValue := (*data)[N]
	return (byteValue & (1 << X)) != 0
}

// / влево (против часовой стрелки)
func Rotate8x8clockWise(data *[]byte, bit int, clockWise bool) {

	var mask byte = (1 << bit)
	var matrix [8][8]byte
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			index := i*8 + j
			if index < len(*data) {
				matrix[i][j] = ((*data)[index] & (mask))
			}
		}
	}
	var rotated [8][8]byte
	if clockWise { // по часовой стрелке,
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				rotated[j][7-i] = matrix[i][j]
			}
		}
	} else {
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				rotated[7-j][i] = matrix[i][j]
			}
		}
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			index := i*8 + j
			bitn := ((*data)[index] >> bit) & 1
			if bitn != rotated[i][j] { //(*data)[index] |= (1 << (7 - j))
				if rotated[i][j] != 0 {
					(*data)[index] |= (1 << bit)
				} else {
					(*data)[index] &^= (1 << bit)
				}
			}
		}
	}
}

// LeftRotate выполняет циклический сдвиг влево на N бит в массиве байт.
func LeftScrollArray64(data *[]byte, N int) {
	if len(*data) == 0 {
		return
	}

	totalBits := len(*data) * 8
	N = N % totalBits

	if N == 0 {
		return
	}

	result := make([]byte, len(*data))

	for i := 0; i < len(*data); i++ {
		newIndex := (i - (N / 8) + len(*data)) % len(*data)
		bitShift := N % 8

		if bitShift != 0 {
			result[newIndex] |= ((*data)[i] << bitShift) // Сдвиг влево
			if i < len(*data)-1 {
				result[newIndex] |= ((*data)[i+1] >> (8 - bitShift)) // Сдвиг вправо из следующего байта
			} else {
				result[newIndex] |= ((*data)[0] >> (8 - bitShift)) // Сдвиг вправо из первого байта
			}
		} else {
			result[newIndex] = (*data)[i]
		}
	}

	copy(*data, result)
}

// RightRotate выполняет циклический сдвиг вправо на N бит в массиве байт.
func RightScrollArray64(data *[]byte, N int) {
	if len(*data) == 0 {
		return
	}

	// Общее количество битов в массиве
	totalBits := len(*data) * 8
	// Уменьшаем N до количества битов в массиве
	N = N % totalBits

	if N == 0 {
		return
	}

	// Сдвигаем данные
	// Сначала создаем новый массив для хранения результата
	result := make([]byte, len(*data))

	// Сдвигаем каждый бит
	for i := 0; i < len(*data); i++ {
		// Вычисляем новый индекс для текущего байта
		newIndex := (i + (N / 8)) % len(*data)
		bitShift := N % 8

		if bitShift != 0 {
			result[newIndex] |= ((*data)[i] >> bitShift) // Сдвиг вправо
			if i > 0 {
				result[newIndex] |= ((*data)[i-1] << (8 - bitShift)) // Сдвиг влево из предыдущего байта
			} else {
				result[newIndex] |= ((*data)[len(*data)-1] << (8 - bitShift)) // Сдвиг влево из последнего байта
			}
		} else {
			result[newIndex] = (*data)[i]
		}
	}

	copy((*data), result)
}

func EncryptData(data []byte, keyStr []byte) []byte {
	for k := 0; k < len(keyStr); k++ {
		var nRollBit = cps_scrollMin + int(keyStr[k])%cps_scrollDiv
		if isEven(keyStr[k]) {
			LeftScrollArray64(&data, nRollBit)
		} else {
			RightScrollArray64(&data, nRollBit)
		}

		for i := 0; i < 8; i++ {
			Rotate8x8clockWise(&data, i, GetBitVal(keyStr, k, i))
		}
	}
	return data
}

func DecryptData(data []byte, keyStr []byte) []byte {
	for k := len(keyStr) - 1; k >= 0; k-- {
		for i := 8 - 1; i >= 0; i-- {
			Rotate8x8clockWise(&data, i, !GetBitVal(keyStr, k, i))
		}
		var nRollBit = cps_scrollMin + int(keyStr[k])%cps_scrollDiv
		if isEven(keyStr[k]) {
			RightScrollArray64(&data, nRollBit)
		} else {
			LeftScrollArray64(&data, nRollBit)
		}
	}
	return data
}

func generateRandomHangulText(size int) string {
	const hangulStart = 0xAC00
	const hangulEnd = 0xD7A3
	const hangulCount = hangulEnd - hangulStart + 1

	text := make([]rune, size)
	for i := 0; i < size; i++ {
		text[i] = rune(rand.Intn(hangulCount) + hangulStart)
	}
	return string(text)
}
func generateRandomArabicText(size int) string {
	const arabicStart = 0x0600
	const arabicEnd = 0x06FF
	const arabicCount = arabicEnd - arabicStart + 1

	text := make([]rune, size)
	for i := 0; i < size; i++ {
		text[i] = rune(rand.Intn(arabicCount) + arabicStart)
	}
	return string(text)
}
func generateRandomHebrewText(size int) string {
	const hebrewStart = 0x0590
	const hebrewEnd = 0x05FF
	const hebrewCount = hebrewEnd - hebrewStart + 1

	text := make([]rune, size)
	for i := 0; i < size; i++ {
		text[i] = rune(rand.Intn(hebrewCount) + hebrewStart)
	}
	return string(text)
}

func base64ToChunk(encoded string) ([]byte, error) {

	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// алгоритм криптографического преобразования
func main() {
	keyStr := []byte(generateRandomHangulText(7)) // = []byte("!寸光阴一寸金，寸金难买寸光 secret key here...")

	shortRandomText := generateRandomHangulText(20)
	fmt.Printf("random text %d bt. %d chunks...", len(shortRandomText), int(math.Ceil(float64(len(shortRandomText))/float64(chunkSize))))
	fmt.Println(shortRandomText)

	hash := sha256.New()

	hash.Write([]byte(shortRandomText))
	checksumHexSrc := hex.EncodeToString(hash.Sum(nil))

	var chunk [64]byte //chunk := make([]byte, 64)
	copy(chunk[:], []byte(shortRandomText[:]))
	chunk[63] = byte(len(shortRandomText))
	fmt.Println("byte array:")
	fmt.Println(chunk[:])
	dataEnc := EncryptData(chunk[:], keyStr)
	fmt.Println("byte array encrypted:")
	fmt.Println(dataEnc)

	var str64enc = base64.StdEncoding.EncodeToString(dataEnc)
	fmt.Println("as Base64: " + str64enc)
	fmt.Println("as int64 array:")
	columns := 4
	for i, value := range Bytes2Long64Array(dataEnc) {
		fmt.Printf("\t0x%X\t", value)
		if (i+1)%columns == 0 {
			fmt.Println()
		}
	}

	decodedBx64Data, err := base64ToChunk(str64enc)
	if err != nil {
		log.Fatalf("Ошибка декодирования: %v", err)
		return
	}

	dataDec := DecryptData(decodedBx64Data, keyStr)
	fmt.Println("byte array (decrypted):")
	fmt.Println(dataDec)
	length := int(dataDec[63])
	decStr := string(dataDec[:length])
	fmt.Println(decStr)

	hash.Reset()
	hash.Write([]byte(decStr))
	checksumHexDec := hex.EncodeToString(hash.Sum(nil))
	if checksumHexDec == checksumHexSrc {
		fmt.Println("Ок")
	}

}
