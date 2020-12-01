#!/usr/local/bin/python
# coding: utf-8

def get_layer(in_list, width, height):
	i = 0
	layer = []

	for y in range(height):
		this_row = []
		for x in range(width):
			idx = x + width*y
			this_row.append(in_list[idx])
		layer.append(this_row)
	return layer

def test_layer():
	img = list(range(16))
	l = get_layer(img, 4,4)
	for row in l:
		print row

def process_image(in_list, width, height):
	block = width * height
	if len(in_list) % block != 0:
		raise Exception("Invalid image size")

	return [
		get_layer(in_list[idx:idx+block], width, height)
		for idx in range(0, len(in_list), block)
	]

def get_number_of(layer, val):
	i = 0
	for row in layer:
		for cell in row:
			if cell == val:
				i += 1
	return i

def get_image_data(s):
	return [int(c) for c in s]

def find_answer(image):
	l = min(image, key=lambda layer: get_number_of(layer, 0))
	ones = get_number_of(l, 1)
	twos = get_number_of(l, 2)
	return ones*twos

def example():
	in_arr = [1,2,3,4,5,6,7,8,9,0,1,2]
	image = process_image(in_arr, 3, 2)
	print(find_answer(image))

def compose(layers):
	height = len(layers[0])
	width = len(layers[0][0])

	image = [
		[0 for _ in range(width)]
		for _ in range(height)
	]
	for layer in reversed(layers):
		for y, row in enumerate(layer):
			for x, cell in enumerate(row):
				if cell == 0:
					image[y][x] = 0
				elif cell == 1:
					image[y][x] = 1
	return image

def render(image):
	for row in image:
		line = ""
		for cell in row:
			if cell == 0:
				line += u"\u001b[40m" + " "
			elif cell == 1:
				line += u"\u001b[47m" + " "
		line += u"\u001b[0m"
		print(line)

def example_render():
	data = get_image_data("0222112222120000")
	layers = process_image(data, 2, 2)
	pixels = compose(layers)
	render(pixels)

def run():
	with open("data.txt") as in_f:
		in_arr = get_image_data(in_f.read())
		img = process_image(in_arr, 25, 6)
		print(find_answer(img))

def run_render():
	with open("data.txt") as in_f:
		in_arr = get_image_data(in_f.read())
		img = process_image(in_arr, 25, 6)
		pixels = compose(img)
		render(pixels)

if __name__ == "__main__":
	run()
	run_render()