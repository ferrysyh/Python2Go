a = input()
b = input()

a1 = int(a)
b1 = int(b)

hasil = a1 - b1
if hasil < 0:
    hasil = hasil * -1
print("Nilai mutlak dari", a, "dikurangi", b, "adalah", hasil)
