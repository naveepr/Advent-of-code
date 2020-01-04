CC :=g++
CFLAGS := -std=gnu++11 -Wall 
LIBS := 
EXE1 := main-1 
EXE2 := main-2
EXE := $(EXE1) $(EXE2) 
OBJ1 := main-1.o 
OBJ2 := main-2.o 
OBJ := $(OBJ1) $(OBJ2)

all: $(EXE1) $(EXE2)
	
$(EXE1): %: $(OBJ1) 
	$(CC) $(LIBS) -o $@ $< 

$(EXE2): %: $(OBJ2) 
	$(CC) $(LIBS) -o $@ $< 

$(OBJ): %.o: %.cpp
	$(CC) -c $(CFLAGS) $<

clean:
	rm -rf $(OBJ) $(EXE)  

.PHONY: clean all


