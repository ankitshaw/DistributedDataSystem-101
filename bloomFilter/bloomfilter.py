import hashlib
import base64

class BloomFilter:
    def __init__(self, bits_size, hash_functions):
        self.bits_size = bits_size
        self.hash_functions = hash_functions
        self.bits_Array = [0]*bits_size
        self.n = 0

    def insert(self, element):
        hash1 = h1(element.encode('utf-8')) % self.bits_size
        self.bits_Array[hash1] = 1

        hash2 = h2(element.encode('utf-8')) % self.bits_size
        self.bits_Array[hash2] = 1

        self.n += 1

    def search(self, element):
        hash1 = h1(element.encode('utf-8')) % self.bits_size
        hash2 = h2(element.encode('utf-8')) % self.bits_size

        if self.bits_Array[hash1] == 0 or self.bits_Array[hash2] == 0:
            print("Not in Bloom Filter")
        else:
            prob = (1.0 - ((1.0 - 1.0/self.bits_size)**(self.hash_functions*self.n))) ** self.hash_functions
            print("Might be in Bloom Filter with false positive probability "+str(prob))

def h1(w):
    h = hashlib.md5(w)
    hashed_bytes = h.digest()
    encoded_bytes = base64.b64encode(hashed_bytes)
    return hash(encoded_bytes[:6])%10

def h2(w):
    h = hashlib.sha256(w)
    hashed_bytes = h.digest()
    encoded_bytes = base64.b64encode(hashed_bytes)
    return hash(encoded_bytes[:6])%10




bf = BloomFilter(10, 1)
bf.insert("Hello")
print(f"Bits: {bf.bits_Array}")

print("\nSearch: Hello")
bf.search("Hello")

print("\nSearch: Not Present")
bf.search("No Way")

print("\nSearch: World")
bf.search("World")