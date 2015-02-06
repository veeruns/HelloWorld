#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <sys/time.h>
#include <sys/timex.h>
#include <string.h>
#include <signal.h>

#define NSEC_PER_SEC 1000000000ULL
// #define COUNT 1000
#define COUNT 100

/* returns 1 if a <= b, 0 otherwise */
static inline int in_order(struct timespec a, struct timespec b)
{
        if(a.tv_sec < b.tv_sec)
                return 1;
        if(a.tv_sec > b.tv_sec)
                return 0;
        if(a.tv_nsec > b.tv_nsec)
                return 0;
        return 1;
}

struct timespec timespec_add(struct timespec ts, unsigned long long ns)
{
    ts.tv_nsec += ns;
    while(ts.tv_nsec >= NSEC_PER_SEC) {
        ts.tv_nsec -= NSEC_PER_SEC;
        ts.tv_sec++;
    }
    return ts;
}


char* time_state_str(int state)
{
    switch (state) {
        case TIME_OK:   return "TIME_OK";
        case TIME_INS:  return "TIME_INS";
        case TIME_DEL:  return "TIME_DEL";
        case TIME_OOP:  return "TIME_OOP";
        case TIME_WAIT: return "TIME_WAIT";
        case TIME_BAD:  return "TIME_BAD";
    }
    return "ERROR"; 
}

/* clear NTP time_status & time_state */
void clear_time_state(void)
{
    struct timex tx;
    int ret;

    /*
     * XXX - The fact we have to call this twice seems
     * to point to a slight issue in the kernel's ntp state
     * managment. Needs to be investigated further.
     */

    tx.modes = ADJ_STATUS;
    tx.status = STA_PLL;
    ret = adjtimex(&tx);

    tx.modes = ADJ_STATUS;
    tx.status = 0;
    ret = adjtimex(&tx);
}

/* Make sure we cleanup on ctrl-c */
void handler(int unused)
{
    clear_time_state();
    exit(0);
}


/* Test for known hrtimer failure */
void test_hrtimer_failure(void)
{
    struct timespec now, target;

    clock_gettime(CLOCK_REALTIME, &now);
    target = timespec_add(now, NSEC_PER_SEC/2);
    clock_nanosleep(CLOCK_REALTIME, TIMER_ABSTIME, &target, NULL);
    clock_gettime(CLOCK_REALTIME, &now);

    if (!in_order(target, now)) {
        printf("Note: hrtimer early expiration failure observed.\n");
    }

}



int main(int argc, char** argv) 
{
    struct timeval tv;
    struct timex tx;
    struct timespec ts;
    int settime = 0;
    int x;
    int i = 0;

    signal(SIGINT, handler);
    signal(SIGKILL, handler);
    printf("This runs continuously. Press ctrl-c to stop\n");


    /* Process arguments */
    if (argc > 1) {
        if (!strncmp(argv[1], "-s", 2)) {
            printf("Setting time to speed up testing\n");
            settime = 1;
        } else {
            printf("Usage: %s [-s]\n", argv[0]);
            printf("    -s: Set time to right before leap second each iteration\n");
        }
    }

    printf("\n");
    while (i < COUNT) {
        int ret;
        time_t now, next_leap;
        /* Get the current time */
        gettimeofday(&tv, NULL);

        /* Calculate the next possible leap second 23:59:60 GMT */
        tv.tv_sec += 86400 - (tv.tv_sec % 86400);
        next_leap = ts.tv_sec = tv.tv_sec;

        if (settime) {
            tv.tv_sec -= 10;
            settimeofday(&tv, NULL);
            printf("Setting time to %s", ctime(&ts.tv_sec));
        }

        /* Reset NTP time state */
        clear_time_state();

        /* Set the leap second insert flag */
        tx.modes = ADJ_STATUS;
        tx.status = STA_INS;
        ret = adjtimex(&tx);
/*
        if (ret) {
            printf("Error: Problem setting STA_INS!: %s\n",
                            time_state_str(ret));
            return -1;
        }
*/

        /* Validate STA_INS was set */
        ret = adjtimex(&tx);
        if (tx.status != STA_INS) {
            printf("Error: STA_INS not set!: %s\n",
                            time_state_str(ret));
            return -1;
        }

        printf("Scheduling leap second for %s", ctime(&next_leap));

        /* Wake up 3 seconds before leap */
        ts.tv_sec -= 3;
        if (clock_nanosleep(CLOCK_REALTIME, TIMER_ABSTIME, &ts, NULL))
            printf("Something woke us up, returning to sleep\n");

        /* Validate STA_INS is still set */
        ret = adjtimex(&tx);
        if (tx.status != STA_INS) {
            printf("Something cleared STA_INS, setting it again.\n");
            tx.modes = ADJ_STATUS;
            tx.status = STA_INS;
            ret = adjtimex(&tx);
        }

        /* Check adjtimex output every half second */
        now = tx.time.tv_sec;
        while (now < next_leap+2) {
            char buf[26];
            ret = adjtimex(&tx);

            ctime_r(&tx.time.tv_sec, buf);
            buf[strlen(buf)-1] = 0; /*remove trailing\n */

            printf("%s + %6ld us\t%s\n",
                    buf,
                    tx.time.tv_usec, 
                    time_state_str(ret));
            now = tx.time.tv_sec;
            /* Sleep for another half second */
            ts.tv_sec = 0;
            ts.tv_nsec = NSEC_PER_SEC/2;
            clock_nanosleep(CLOCK_MONOTONIC, 0, &ts, NULL);         
        }

        /* Note if kernel has known hrtimer failure */
        test_hrtimer_failure();

        printf("Leap complete\n\n");

        i++;
    }

    clear_time_state();
    return 0;
}

