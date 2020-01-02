CC=g++
CFLAGS= -std=gnu++11 -Wall 
LIBS = 
EXE = main
OBJ = main.o

all: $(EXE) 
	
$(EXE): $(OBJ) 
	$(CC) $(LIBS) -o$(EXE) $(OBJ) 

$(OBJ): main.cpp
	$(CC) -c $(CFLAGS) main.cpp

clean:
	rm -rf $(OBJ) $(EXE) 


