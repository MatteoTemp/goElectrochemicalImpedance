set term qt persist size 700,500
set datafile separator ','

set xtics font 'arial, 10' nomirror
set ytics font 'arial, 13' nomirror
set y2tics font 'arial, 13' nomirror

set grid
set size square

set format x "%0.0e"
set logscale x 10

set xlabel "log f"
set zlabel "-Im Z"
set ylabel "Re Z"


splot 'OutputFiles/Nyq.csv' using 1:2:3 with lines

pause mouse close
