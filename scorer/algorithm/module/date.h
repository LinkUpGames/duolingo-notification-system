#include <stdint.h>
#include <time.h>

#define timestamp int64_t

/**
 * Get the current time object
 */
struct timespec get_current_time();

/**
 * Given a timespec object, return the time in milliseconds since epoch
 * @param time The timespec object to convert into milliseconds
 */
timestamp get_milliseconds(struct timespec time);
