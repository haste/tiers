#mkdir -p /home/haste/opt
mkdir /tmp/py
cd /tmp/py
wget https://www.python.org/ftp/python/2.7.3/Python-2.7.3.tgz
wget http://www.netlib.org/lapack/lapack-3.5.0.tgz
wget http://downloads.sourceforge.net/project/numpy/NumPy/1.9.1/numpy-1.9.1.tar.gz?r=http%3A%2F%2Fsourceforge.net%2Fprojects%2Fnumpy%2Ffiles%2F&ts=1417212043&use_mirror=optimate -O numpy-1.9.1.tar.gz
wget http://downloads.sourceforge.net/project/opencvlibrary/opencv-unix/2.4.9/opencv-2.4.9.zip?r=http%3A%2F%2Fopencv.org%2Fdownloads.html&ts=1417212086&use_mirror=cznic -O opencv-2.4.9.zip

# Python
tar xvzf Python-2.7.3.tgz
cd Python-2.7.3
./configure --prefix /home/haste/opt/ --enable-unicode=ucs4 -enable-shared
make -j21
make install
cd ..

# lapack
tar xvzf lapack-3.5.0.tgz
cd lapack-3.5.0
mkdir rel
cd rel
cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/home/haste/opt -D BUILD_SHARED_LIBS:BOOL=ON ..
make -j21
make install
cd ..

# numpy
tar xvzf numpy-1.9.1.tar.gz
cd numpy-1.9.1
export LAPACK=/home/haste/opt/lib/liblapack.so
export BLAS=/home/haste/opt/lib/libblas.so
python setup.py install --prefix /home/haste/opt
cd ..

# OpenCV
unzip opencv-2.4.9.zip
cd opencv-2.4.9
mkdir rel
cd rel
LD_PRELOAD=/home/haste/opt/lib/liblapack.so:/home/haste/opt/lib/libblas.so cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/home/haste/opt -D WITH_TBB=ON -D BUILD_NEW_PYTHON_SUPPORT=ON -D WITH_V4L=OFF PYTHON_PACKAGES_PATH=/home/haste/opt/lib/python2.7/site-packages/ -D PYTHON_INCLUDE_DIR=/home/haste/opt/include/python2.7/ -D PYTHON_LIBRARY=/home/haste/opt/lib/libpython2.7.so -D PYTHON_EXECUTABLE=/home/haste/opt/bin/python2.7 ..
make -j21
make install
cd ..

LD_LIBRARY_PATH=/home/haste/opt/lib /home/haste/opt/bin/python
