import sys
import pyvips

print("Running test %s" % sys.argv[1])

def thumbnail():
    iterations = int(sys.argv[2])
    for i in range(1, iterations):
        print("Iteration %d" % i)
        thumbnailStep()

def thumbnailStep():
    image = pyvips.Image.thumbnail("images/image50.jpg", 5000)
    image.write_to_file('out/thumbnail.jpg')


match sys.argv[1]:
    case "thumbnail":
        thumbnail()
    case _:
        raise NameError("Unknown test")
