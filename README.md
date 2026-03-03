en: example of implementing encryption for a 64-byte array based on an 8x8 cube slice rotation algorithm. It will then be used to practice writing a process for goroutine interactions on a stream of encryption and decryption of a large text divided into 64-byte chunks.


<img width="1024" height="1024" alt="019cb54e-9b51-740b-ac61-27cfa8dfcb4c" src="https://github.com/user-attachments/assets/54733ab8-4827-40ce-978b-313f2844e61b" />
random text 60 bt. 펅꼫톼칽췻펍슈픠뿭됝곐쉱쬹콿붕껒볌뫽팂덝

byte array:
[237 142 133 234 188 171 237 134 188 236 185 189 236 183 187 237 142 141 236 138 136 237 148 160 235 191 173 235 144 157 234 179 144 236 137 177 236 172 185 236 189 191 235 182 149 234 187 146 235 179 140 235 171 189 237 140 130 235 141 157 0 0 0 60]

byte array (encrypted):
[163 91 121 59 227 83 41 21 65 125 89 221 85 47 255 211 137 57 155 59 91 88 31 115 83 95 241 125 22 201 109 149 85 25 209 223 83 81 19 157 85 171 75 57 207 102 109 123 159 83 71 123 133 82 17 243 83 123 117 143 115 25 213 73]

 <p>as Base64:<font size="6" color="#fa8e47" face="serif">o1t5O+NTKRVBfVndVS//04k5mztbWB9zU1/xfRbJbZVVGdHfU1ETnVWrSznPZm17n1NHe4VSEfNTe3WPcxnVSQ==</font></p>
 
byte array (decrypted):
[237 142 133 234 188 171 237 134 188 236 185 189 236 183 187 237 142 141 236 138 136 237 148 160 235 191 173 235 144 157 234 179 144 236 137 177 236 172 185 236 189 191 235 182 149 234 187 146 235 179 140 235 171 189 237 140 130 235 141 157 0 0 0 60]

펅꼫톼칽췻펍슈픠뿭됝곐쉱쬹콿붕껒볌뫽팂덝

