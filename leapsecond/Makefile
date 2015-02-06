CC=gcc
CFLAGS=-I.
LIBS=-lrt
leapsecond:leap_second_stress_test.c 
	$(CC) -o leapsecond leap_second_stress_test.c $(CFLAGS) $(LIBS)

clean:
	rm -rf *.o *~ core leapsecond

