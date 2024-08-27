num = {}
count_even = 0
count_even = tonumber(count_even)

for i = 1, 100, 1 do
  num[i] = math.random(1, 100)
  if num[i] % 2 == 0 then
    count_even = count_even + 1
  end
end

print(count_even)
