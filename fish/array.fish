# 配列の操作
set a 1 2 3 4 5

# インデックス
echo $a[0]
echo $a[-1]

# スライス
echo $a[1..2]
echo $a[3..-1]

# 範囲外
echo $a[1000]
