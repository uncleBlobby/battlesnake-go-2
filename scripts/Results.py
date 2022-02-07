with open('./results.txt') as f:
    lines = f.readlines()

f.close()

goSnakeWins = 0
antiBlobWins = 0

for i in range(0, len(lines)):
    #if len(lines[i]) < 100:
        #print(lines[i], end = "")
    if "DONE" in lines[i]:
        #print(lines[i], end = "")
        if "20blobby2" in lines[i]:
            goSnakeWins += 1
        if "anti-blobby" in lines[i]:
            antiBlobWins += 1


print("Results:")
print("20blobby2: ", goSnakeWins, " wins")
print("anti-blobby", antiBlobWins, " wins")
