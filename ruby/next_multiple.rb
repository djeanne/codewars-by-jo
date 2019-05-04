def round_to_next_5(n)
	a = n % 5
	unless n == 0
		if a == 0
			b = 5 * (n / 5)
		else
			b = 5 * ((n / 5) + 1)
		end
	else
		b = 0
	end
	return b
end