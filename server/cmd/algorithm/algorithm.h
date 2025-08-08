#ifndef __ALGORITHM_H__
#define __ALGORITHM_H__

#include <stdint.h>

// The notification being used
typedef struct Notification {
  char *notification_id;
  int days;
  float score;
} Notification;

/**
 * Compute decay factor for the notification
 * @param penalty The base penalty for the day
 * @param factor Scale the base factor
 * @param The number of days
 */
void compute_decay(float penalty, float factor, float days,
                   Notification **notification);

#endif
