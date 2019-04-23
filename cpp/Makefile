ifeq (${CXX},)
CXX=g++
endif
LINK=${CXX}

TOPDIR=`pwd`
COMMONPATH=`pwd`/..
LIBPATH=${TOPDIR}/libs
#-static-libstdc++
LDFLAGS= -static-libgcc -std=c++11 -Wl,-rpath,$(LIBPATH)
CXXFLAGS  = -pipe -std=c++0x -fPIC -g -fno-omit-frame-pointer \
			-DNDEBUG=1 -Wconversion -O3 -Wall -W -fvisibility=hidden

LIB	   = -pthread -lpthread -L$(LIBPATH) -lrecorder -lrt
INCPATH =-I. -I${TOPDIR}/include

SRC=$(wildcard ./agorasdk/*.cpp *.cpp)
OBJ=$(addsuffix .o, $(basename $(SRC)))

#test:
#	echo "$(OBJ)"


TARGET=libmyagora.a


.PHONY: all clean

libmyagora.a: $(OBJ)
	ar r "$@" $(OBJ)

%.o: %.cpp
	$(CXX) -c $(CXXFLAGS) $(INCPATH) -o "$@" "$<"

clean:
	rm -f $(TARGET)
	rm -f ${OBJ}
