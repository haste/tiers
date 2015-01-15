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

FLANN_INDEX_KDTREE = 1  # bug: flann enums are missing
FLANN_INDEX_LSH    = 6

def init_feature(name):
	chunks = name.split('-')
	if chunks[0] == 'sift':
		detector = cv2.SIFT()
		norm = cv2.NORM_L2
	elif chunks[0] == 'surf':
		detector = cv2.SURF(800)
		norm = cv2.NORM_L2
	elif chunks[0] == 'orb':
		detector = cv2.ORB(400)
		norm = cv2.NORM_HAMMING
	elif chunks[0] == 'akaze':
		detector = cv2.AKAZE()
		norm = cv2.NORM_HAMMING
	elif chunks[0] == 'brisk':
		detector = cv2.BRISK()
		norm = cv2.NORM_HAMMING
	else:
		return None, None
	if 'flann' in chunks:
		if norm == cv2.NORM_L2:
			flann_params = dict(algorithm = FLANN_INDEX_KDTREE, trees = 5)
		else:
			flann_params= dict(algorithm = FLANN_INDEX_LSH,
					table_number = 6, # 12
					key_size = 12,     # 20
					multi_probe_level = 1) #2
		matcher = cv2.FlannBasedMatcher(flann_params, {})  # bug : need to pass empty dict (#1329)
	else:
		matcher = cv2.BFMatcher(norm)
	return detector, matcher

def filter_matches(kp1, kp2, matches, ratio = 0.7):
	mkp1, mkp2 = [], []
	for m in matches:
		if len(m) == 2 and m[0].distance < m[1].distance * ratio:
			m = m[0]
			mkp1.append( kp1[m.queryIdx] )
			mkp2.append( kp2[m.trainIdx] )
	p1 = np.float32([kp.pt for kp in mkp1])
	p2 = np.float32([kp.pt for kp in mkp2])
	kp_pairs = zip(mkp1, mkp2)
	return p1, p2, kp_pairs

def angle_cos(p0, p1, p2):
	dx1 = p0[0][0] - p2[0][0]
	dy1 = p0[0][1] - p2[0][1]
	dx2 = p1[0][0] - p2[0][0]
	dy2 = p1[0][1] - p2[0][1]

	return (dx1*dx2 + dy1*dy2)/math.sqrt((dx1*dx1 + dy1*dy1)*(dx2*dx2 + dy2*dy2) + 1e-10)

def fill_view(view, approx):
	x1, x2 = np.inf, 0
	y1, y2 = np.inf, 0

	for p in approx:
		x1 = min(x1, p[0][0])
		y1 = min(y1, p[0][1])

		x2 = max(x2, p[0][0])
		y2 = max(y2, p[0][1])

	cv2.rectangle(view, (x1 - 10, y1 - 10), (x2 + 10, y2 + 10), (0, 0, 0), cv.CV_FILLED)

basePath = os.path.dirname(sys.argv[0])
cachePath = sys.argv[1]
imageName = sys.argv[2]

ranges = [
	# bronze
	[
		# [15, 91, 188]
		np.array([0, 71, 168]),
		np.array([35, 111, 208]),
		# [10, 215, 51]
		np.array([0, 195, 31]),
		np.array([30, 235, 71]),
	],
	# silver
	[
		# [90, 33, 46]
		np.array([70, 13, 26]),
		np.array([110, 53, 66]),
		# [91, 62, 194]
		np.array([61, 42, 174]),
		np.array([111, 82, 214]),
	],
	# gold
	[
		# [16, 158, 76]
		np.array([0, 138, 56]),
		np.array([36, 178, 96]),
		# [25, 88, 216]
		np.array([5, 68, 196]),
		np.array([45, 108, 236]),
	],
	# platinum
	[
		# [90, 9, 30]
		np.array([70, 0, 10]),
		np.array([110, 29, 50]),
		# [87, 27, 186]
		np.array([67, 7, 166]),
		np.array([107, 47, 206]),
	],
	# onyx
	[
		# [84, 56, 113]
		np.array([64, 36, 93]),
		np.array([104, 76, 133]),
		# [0, 0, 2]
		np.array([0, 0, 0]),
		np.array([20, 20, 22]),
	],
]

img = cv2.imread(cachePath + "/" + imageName)
imgCrop = img[:(img.shape[0] // 2), :img.shape[1]]
viewCrop = imgCrop.copy()

view = img.copy()
gray = cv2.cvtColor(imgCrop, cv2.COLOR_BGR2GRAY)

# Matching
query = cv2.imread(basePath + '/template.png', 0)
detector, matcher = init_feature("sift")

kp1, des1 = detector.detectAndCompute(query, None)

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
		fill_view(view, approx)
	elif len(approx) == 6 and cv2.isContourConvex(approx) and cv2.contourArea(approx) > 2000:
		max_cosine = 0.0
		for r in range(2, len(approx) + 1):
			max_cosine = max(max_cosine, math.fabs(angle_cos(approx[r%4], approx[r-2], approx[r-1])))

		if max_cosine >= 0.65:
			fill_view(view, approx)
			continue

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

		crop = masked[y:y+h, x:x+w]
		cropThres, bin = cv2.threshold(crop, 100, 255, cv2.THRESH_BINARY)

		kp2, des2 = detector.detectAndCompute(crop, None)
		if des2 is None:
			continue

		badge_top = min(badge_top, y)
		badge_bottom = max(badge_bottom, y + h)

		matches = matcher.knnMatch(des1, trainDescriptors = des2, k = 2)

		ratio = 0.0
		p1, p2, kp_pairs = filter_matches(kp1, kp2, matches)
		if len(p1) >= 4:
			H, status = cv2.findHomography(p1, p2, cv2.RANSAC, 1.0)
			if np.sum(H) == 0.0:
				continue

			#print('%d / %d  = %.3f inliers/matched' % (np.sum(status), len(status), np.sum(status) / len(status)))
			#print('%.3f' % (np.sum(status) / len(kp2)))

			ratio = np.sum(status) / len(kp2)

		if ratio >= 0.15:
			masked = cv2.bitwise_and(imgCrop, imgCrop, mask = mask)
			badge = cv2.blur(masked[y:y+h, x:x+w], (3, 3))

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

			tier = -1
			for k, v in reversed(list(enumerate(ranges))):
				m1 = cv2.inRange(hsv, v[0], v[1])
				m2 = cv2.inRange(hsv, v[2], v[3])

				res4 = cv2.bitwise_and(hsv, hsv, mask = m1+m2)
				if res4.mean() > 4:
					tier = k
					break

			innovator.append((ratio, len(matches), tier))

out = {}

# Sort matches
innovator.sort(key=lambda tup: tup[0], reverse=True)
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
