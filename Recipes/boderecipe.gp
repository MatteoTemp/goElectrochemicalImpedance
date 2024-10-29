set datafile separator ","
set grid

set xtics font 'arial, 10' nomirror
set ytics font 'arial, 13' nomirror
set y2tics font 'arial, 13' nomirror
set format x "%0.0e"
set logscale x 10

set y2tics 10
set ytics nomirror

set ylabel "log|H|"
set y2label "phi(H) / deg"
set xlabel "log freq / Hz"

plot 'OutputFiles/Bode.csv' using 1:2 with lines axis x1y1 , 'OutputFiles/Bode.csv' using 1:3 with lines axis x1y2
