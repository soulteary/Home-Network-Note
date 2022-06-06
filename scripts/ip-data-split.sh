#!/bin/sh
# author soulteary

mkdir -p result
for region in $(cat data/region.txt)
do
    echo "${region}"
    echo "cat data/pure.txt | awk -F '|' '/${region}/&&/ipv4/ {print \$4 \"/\" 32-log(\$5)/log(2)}' | cat > result/${region}-ipv4.txt" | bash - 
    echo "cat data/pure.txt | awk -F '|' '/${region}/&&/ipv6/ {print \$4 \"/\" \$5}' | cat > result/${region}-ipv6.txt" | bash -
done
