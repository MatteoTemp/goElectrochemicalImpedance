set datafile separator ','
set grid
set size square
plot 'OutputFiles/Lasajous.csv' using 1:2 with lines
