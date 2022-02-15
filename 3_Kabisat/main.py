tahun = int(input())

kabisat = tahun % 400 == 0 or tahun % 4 == 0 and tahun % 100 != 0
print(kabisat)
