set datafile separator ","
set grid

set ylabel "-Im(H) / Ohm"
set xlabel "Re(H)/ Ohm"
plot 'Nyq.csv' using 2:3 with lines
