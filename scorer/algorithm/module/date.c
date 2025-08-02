#include <stdlib.h>
#include <time.h>

struct tm *get_utc_time() {
  struct tm *utc_time = NULL;

  // Get the current time as a time_t object
  time_t rawtime;
  time(&rawtime);

  // Get the current utc time
  struct tm *gmtime_object = gmtime(&rawtime);

  // Copy it over
  if (gmtime_object != NULL) {
    utc_time = malloc(sizeof(struct tm));

    *utc_time = *gmtime_object;
  }

  return utc_time;
}
