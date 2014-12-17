#!/usr/bin/env python
from __future__ import division
from __future__ import print_function

import cv2
import cv2.cv as cv
import json
import math
import numpy as np
import os
import sys

basePath = os.path.dirname(sys.argv[0])
cachePath = sys.argv[1]
imageName = sys.argv[2]

ranges = [
	# bronze
	[
		# [15, 91, 188]
		np.array([5, 81, 178]),
		np.array([25, 101, 198]),
		# [10, 215, 51]
		np.array([0, 205, 41]),
		np.array([20, 225, 61]),
	],
	# silver
	[
		# [90, 33, 46]
		np.array([90, 23,36]),
		np.array([100, 43, 56]),
		# [91, 62, 194]
		np.array([81, 52, 184]),
		np.array([110, 53, 66]),
	],
	# gold
	[
		# [16, 158, 76]
		np.array([6, 148, 66]),
		np.array([26, 168, 86]),
		# [25, 88, 216]
		np.array([15, 78, 206]),
		np.array([35, 98, 226]),
	],
	# platinum
	[
		# [90, 9, 30]
		np.array([80, 0, 20]),
		np.array([100, 19, 40]),
		# [87, 27, 186]
		np.array([77, 17, 176]),
		np.array([97, 37, 196]),
	],
	# onyx
	[
		# [84, 56, 113]
		np.array([74, 46, 103]),
		np.array([94, 66, 123]),
		# [0, 0, 2]
		np.array([0, 0, 0]),
		np.array([10, 10, 12]),
	],
]

img = cv2.imread(cachePath + "/" + imageName)

# remove the menu
i = 0
crop_top = 0
for row in img:
	i += 1
	if row.sum() < 3000:
		crop_top = i
		break

img = img[crop_top:, :img.shape[1]]
imgCrop = img[:(img.shape[0] // 2), :img.shape[1]]
viewCrop = imgCrop.copy()

view = img.copy()
gray = cv2.cvtColor(imgCrop, cv2.COLOR_BGR2GRAY)

# Matching
query = cv2.imread(basePath + '/template.png', 0)
sift = cv2.SIFT()
kp1, des1 = sift.detectAndCompute(query, None)

FLANN_INDEX_KDTREE = 0
index_params = dict(algorithm = FLANN_INDEX_KDTREE, trees = 5)
search_params = dict(checks = 50)

flann = cv2.FlannBasedMatcher(index_params, search_params)

ret, bin = cv2.threshold(gray, 55, 150, cv2.THRESH_BINARY)

# Badges
contours, h = cv2.findContours(bin, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)

r = 0
badge_top = np.inf
badge_bottom = 0

innovator = []
for cnt in contours:
	r += 1
	approx = cv2.approxPolyDP(cnt, 0.01 * cv2.arcLength(cnt, True), True)

	if (len(approx) == 8 or len(approx) == 4 or len(approx) == 7 or len(approx) == 5) and cv2.isContourConvex(approx) and cv2.contourArea(approx) > 2000:
		x1, x2 = np.inf, 0
		y1, y2 = np.inf, 0

		for p in approx:
			x1 = min(x1, p[0][0])
			y1 = min(y1, p[0][1])

			x2 = max(x2, p[0][0])
			y2 = max(y2, p[0][1])

		cv2.rectangle(view, (x1 - 10, y1 - 10), (x2 + 10, y2 + 10), (0, 0, 0), cv.CV_FILLED)
	elif len(approx) == 6 and cv2.isContourConvex(approx) and cv2.contourArea(approx) > 2000:
		ret, mask = cv2.threshold(gray, 0, 0, cv2.THRESH_BINARY)
		cv2.drawContours(mask, [cnt], 0, (255, 255, 255), -1)

		p0 = approx[0][0]
		p1 = approx[1][0]
		p2 = approx[2][0]
		p3 = approx[3][0]

		top = p1[1] - p0[1]
		bottom = p3[1] - p2[1]

		masked = cv2.bitwise_and(gray, gray, mask = mask)

		x, y, w, h = cv2.boundingRect(cnt)

		crop = masked[y+top:y+h-bottom, x:x+w]

		kp2, des2 = sift.detectAndCompute(crop, None)
		if des2 is None:
			continue

		badge_top = min(badge_top, y)
		badge_bottom = max(badge_bottom, y + h)

		matches = flann.knnMatch(des1,des2,k=2)

		count = 0
		for m, n in matches:
			if m.distance < 0.7 * n.distance:
				count += 1

		if count / len(matches) > 0.09:
			masked = cv2.bitwise_and(imgCrop, imgCrop, mask = mask)
			badge = masked[y:y+h, x:x+w]

			Z = badge.reshape((-1, 3))
			Z = np.float32(Z)
			criteria = (cv2.TERM_CRITERIA_EPS + cv2.TERM_CRITERIA_MAX_ITER, 10, 1.0)
			K = 2
			ret, label, center=cv2.kmeans(Z, K, criteria, 10, cv2.KMEANS_RANDOM_CENTERS)

			center = np.uint8(center)

			res = center[label.flatten()]
			res2 = res.reshape((badge.shape))

			hsv = cv2.cvtColor(res2, cv2.COLOR_BGR2HSV)

			c0 = cv2.cvtColor(np.uint8([[center[0]]]), cv2.COLOR_BGR2HSV)
			c1 = cv2.cvtColor(np.uint8([[center[1]]]), cv2.COLOR_BGR2HSV)

			for k, v in enumerate(ranges):
				m1 = cv2.inRange(hsv, v[0], v[1])
				m2 = cv2.inRange(hsv, v[2], v[3])

				res4 = cv2.bitwise_and(hsv, hsv, mask = m1+m2)
				if res4.mean() > 4:
					break

			innovator.append((count, len(matches), k))

out = {}

# Sort matches
innovator.sort(key=lambda tup: tup[1], reverse=True)
for i, v in enumerate(innovator):
	out = {
		'good': v[0],
		'total': v[1],
		'rank': v[2],
	}

	break

# Mission circles
mission_bottom = 0
grayCrop = gray[badge_bottom:gray.shape[0], 0:gray.shape[1]]
blur = cv2.GaussianBlur(grayCrop, (9, 9), 2, 2)
circles = cv2.HoughCircles(blur, cv.CV_HOUGH_GRADIENT, 1, 10)
if type(circles) is np.ndarray:
	for i in circles[0]:
		mission_bottom = max(mission_bottom, i[1] + badge_bottom + i[2])

achieve_bottom = max(badge_bottom, mission_bottom)

# Lines above
top = view[:badge_top, :view.shape[1]]

i = 0
offset_bottom = 0
for row in reversed(top):
	i += 1
	if row.sum() < 1e4:
		offset_bottom = len(top) - i
		break

top = top[0:offset_bottom, :view.shape[1]]

# cover up the share button
cv2.rectangle(top, (top.shape[1] // 8 * 6, 0), (top.shape[1], top.shape[0]//2), (0, 0, 0), cv.CV_FILLED)

# Lines below
bottom = view[achieve_bottom:, 0:view.shape[1]]

maxSwitches = 7
if mission_bottom > 0:
	maxSwitches = 5

switches = 0
wasBlack = True
i = 0
for x, row in enumerate(bottom):
	i += 1

	mean = row.mean()
	if mean < 8 and wasBlack == False:
		switches += 1
		wasBlack = True
	elif mean >= 8 and wasBlack == True:
		switches += 1
		wasBlack = False

	if switches == maxSwitches:
		break

bottom = bottom[i - 1:, :view.shape[1]]

h1, w1 = top.shape[:2]
h2, w2 = bottom.shape[:2]

view = np.zeros((top.shape[0]+bottom.shape[0], top.shape[1], 3), np.uint8)

view[:h1, :w1] = top
view[h1:h1+h2, :w1] = bottom

print(json.dumps(out))
cv2.imwrite(cachePath + "/cv_" + imageName, view)
