#include "algorithm.h"
#include "date.h"
#include <time.h>

char *select_arm(Arm **arms, int length, float alpha, float beta) {
  // Get current time
  struct tm *time = get_utc_time();

  // Iterate through every arm
  for (int i = 0; i < length; i++) {
    Arm *arm = arms[i];

    float count = arm->count;
    float reward = arm->reward;

    float average_reward = 0;
    float recovery = 30; // a month old, this could be anything though

    // If the arm is not new and has a history of being selected
    if (reward > 0) {
      average_reward = reward / count;
    }

    // Get the recovery for the arm
    if (last_sent > 0) {
      recovery = time->tm_isdst
    }
  }

  return NULL;
}
