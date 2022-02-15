b1 = input()
baterai = int(b1)

if(baterai == 100):
    print("Cabut charger")
elif (baterai >= 75 and baterai <= 99):
    print("Charge dalam 30 menit")
elif (baterai >= 20 and baterai <= 74):
    print("Charge dalam 10 menit")
else:
    print("Segera charge")
