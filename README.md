en: example of implementing encryption for a 64-byte array based on an 8x8 cube slice rotation algorithm. It will then be used to practice writing a process for goroutine interactions on a stream of encryption and decryption of a large text divided into 64-byte chunks.

<img width="200" height="200" alt="019cb54e-9b51-740b-ac61-27cfa8dfcb4c" src="https://github.com/user-attachments/assets/5216e9c7-9ce3-448e-a296-672c869b5f6d" />


random text 60 bt. 펅꼫톼칽췻펍슈픠뿭됝곐쉱쬹콿붕껒볌뫽팂덝

byte array:
[237 142 133 234 188 171 237 134 188 236 185 189 236 183 187 237 142 141 236 138 136 237 148 160 235 191 173 235 144 157 234 179 144 236 137 177 236 172 185 236 189 191 235 182 149 234 187 146 235 179 140 235 171 189 237 140 130 235 141 157 0 0 0 60]

byte array (encrypted):
[163 91 121 59 227 83 41 21 65 125 89 221 85 47 255 211 137 57 155 59 91 88 31 115 83 95 241 125 22 201 109 149 85 25 209 223 83 81 19 157 85 171 75 57 207 102 109 123 159 83 71 123 133 82 17 243 83 123 117 143 115 25 213 73]

as Base64: o1t5O+NTKRVBfVndVS//04k5mztbWB9zU1/xfRbJbZVVGdHfU1ETnVWrSznPZm17n1NHe4VSEfNTe3WPcxnVSQ==

as int64 array:
	0x-5CA486C41CACD6EB		0x417D59DD552FFFD3		0x-76C664C4A4A7E08D		0x535FF17D16C96D95
	0x5519D1DF5351139D		0x55AB4B39CF666D7B		0x-60ACB8847AADEE0D		0x537B758F7319D549

byte array (decrypted):
[237 142 133 234 188 171 237 134 188 236 185 189 236 183 187 237 142 141 236 138 136 237 148 160 235 191 173 235 144 157 234 179 144 236 137 177 236 172 185 236 189 191 235 182 149 234 187 146 235 179 140 235 171 189 237 140 130 235 141 157 0 0 0 60]

펅꼫톼칽췻펍슈픠뿭됝곐쉱쬹콿붕껒볌뫽팂덝
