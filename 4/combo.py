
def is_increasing(key):
	last = 0
	for d in key:
		if d < last:
			return False
		last = d
	return True

def is_valid(key):
	last = 0
	has_dupe = False
	for d in key:
		if d < last:
			return False
		if d == last:
			has_dupe = True
		last = d
	return has_dupe

def is_valid_strict(key):
	last = -1
	dupe_len = 0
	has_dupe = False

	for d in key:
		if d < last:
			return False
		if d == last:
			dupe_len += 1
		else:
			if dupe_len == 2:
				has_dupe = True
			dupe_len = 1
		last = d

	return has_dupe or dupe_len == 2


	return has_dupe

def generate_keys(valid_f):
	keys = []
	queue = [()]
	while queue:
		k = queue.pop(0)
		for i in range(10):
			newK = tuple(k + (i,))
			if not is_increasing(newK):
				continue
			if len(newK) == 6:
				if valid_f(newK):
					keys.append(newK)
				continue
			queue.append(newK)

	return keys

def keys_in_range(start, end):
	keys = [
		int("".join((str(c) for c in k))) 
		for k in generate_keys(is_valid_strict)
	]
	valid = 0
	for k in keys:
		if k < start:
			continue
		if k > end:
			return valid
		print k
		valid += 1

if __name__ == "__main__":
	print(keys_in_range(146810, 612564))
