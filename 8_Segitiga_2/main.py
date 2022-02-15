x1 = input()
x2 = input()
x3 = input()

ix1 = int(x1)
ix2 = int(x2)
ix3 = int(x3)

if (ix1 + ix2 > ix3) or (ix2 + ix3 > ix1) or (ix1 + ix3 > ix2):
    print("Sisi segitiga")
else:
    print("Bukan sisi segitiga")
