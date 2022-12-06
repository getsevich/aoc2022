import os, sys
import codecs

data = '';

for filename in os.listdir('CSV_zh_tw'):
    print( filename )
    if filename.endswith(".csv"):
        path = os.path.join('CSV_zh_tw', filename)
        curData = codecs.open( path, "r", "utf16" ).read()
        data = data + ''.join(set(curData))


f = codecs.open("input.txt", "w", "utf16" )
f.write( data )
f.close()


print( "----------------------------------------------------" )
print( "all done, " + str(len(data)) + " unique glyphs found" )