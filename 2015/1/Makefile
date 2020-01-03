CC :=g++
CFLAGS := -std=gnu++11 -Wall 
LIBS := 
EXE := main-1 main-2
OBJ := main-1.o main-2.o 

all: $(EXE) 
	
$(EXE): %: $(OBJ) 
	$(CC) $(LIBS) -o $@ $< 

$(OBJ): %.o: %.cpp
	$(CC) -c $(CFLAGS) $<

clean:
	rm -rf $(OBJ) $(EXE) 

.PHONY: clean all


