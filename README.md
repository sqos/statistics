```
version: 0.1, statistics cvs column range

example: statistics -in=input.csv -out=output.csv -column=3 -value=8888.88 -delta=6.66 -border=true
         will statistics 3'd column value, range [8888.88-6.66, 8888.88+6.66]

  -border
        value range include border, default: true (default true)
  -column int
        set column number which will be statistics
  -delta float
        the delta value for statistics
  -help
        get help information
  -in string
        set the input file name, default: input.csv (default "input.csv")
  -out string
        set the output file name, default: output.csv (default "output.csv")
  -stdout
        set results output to stdout, default: true (default true)
  -value float
        the targe value for statistics

input.csv
xiaoming,80,88,100,6.65
lihong,99.6,99,99,6.33
wangfei,100,88,10,6.45
lilei,99.88,100,10,6.55

./statistics -column 5 -value 6.45 -delta 0.1 -border=false

output.csv
1: wangfei,100,88,10,6.45

./statistics -column 5 -value 6.45 -delta 0.1
output.csv
1: wangfei,100,88,10,6.45
2: lilei,99.88,100,10,6.55


```
