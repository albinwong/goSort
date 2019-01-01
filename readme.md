`cd ./bin`
`./sorter -i unsort.dat -o sorted.dat -a qsort`
infile = unsort.dat outfile = sorted.dat algorithms = qsort
The sorting process costs 666ns to complete.
   
`./sorter -i unsorted.dat -o sorted.dat -a bubblesort`
infile = unsorted.dat outfile = sorted.dat algorithm = bubblesort
The sorting process costs 2us to complete.

`cat sorted.dat`
