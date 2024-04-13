## wc-tool

# Steps to build:

1. clone the wc-tool repo locally
2. go build
3. cat <filename> | ./wc-tool.exe -l

# Other examples:

$ cat test.txt | ./wc-tool.exe -m
  342181 

$ cat test.txt | ./wc-tool.exe -l
    7143 

$ cat test.txt | ./wc-tool.exe -m
  342181 

$ ./wc-tool.exe test.txt 
7143 58164 342181 test.txt

$ ./wc-tool.exe -l test.txt
    7143 test.txt


$ ./wc-tool.exe -w test.txt
   58164 test.txt


$ ./wc-tool.exe -c test.txt
  342181 test.txt



