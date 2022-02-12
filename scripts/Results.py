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
        if "wrappedhack" in lines[i]:
            goSnakeWins += 1
        if "anti-blobby-live" in lines[i]:
            antiBlobWins += 1


print("Results:")
print("wrappedhack: ", goSnakeWins, " wins")
print("anti-blobby-live", antiBlobWins, " wins")
