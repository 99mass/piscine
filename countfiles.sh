#find . -type d | wc -l; find -type f |wc -l
find . \( -type d,f \)  |wc -l
