#include "date.h"
#include <time.h>

struct timespec get_current_time() {
  struct timespec current;
  clock_gettime(CLOCK_REALTIME, &current);

  return current;
}

timestamp get_milliseconds(struct timespec time) {
  timestamp current = (timestamp)time.tv_sec * 1000 + time.tv_nsec / 1000000;

  return current;
}
