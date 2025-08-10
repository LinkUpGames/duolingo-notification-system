#include "notification.h"
#include <stdlib.h>
#include <string.h>

Notification *create_notification(char const *id, float score,
                                  float probability) {

  Notification *notification = (Notification *)malloc(sizeof(Notification));

  if (notification == NULL) {
    return NULL;
  }

  notification->id = strdup(id);
  notification->probability = probability;
  notification->score = score;

  // Setup the compute sums for later
  notification->m_plus = 0;
  notification->m_plus_count = 0;
  notification->m_plus_sum = 0;

  notification->m_minus = 1; // Avoid division by zero later
  notification->m_minus_count = 0;
  notification->m_minus_sum = 0;

  return notification;
}

void free_notification(Notification *notification) {
  if (notification->id != NULL) {
    free(notification->id);
    notification->id = NULL;
  }

  free(notification);
}

NotificationArray *create_notification_list(long length) {

  Notification **notifications =
      (Notification **)calloc(length, sizeof(Notification *));

  if (notifications == NULL) {
    return NULL;
  }

  NotificationArray *array =
      (NotificationArray *)malloc(sizeof(NotificationArray));
  if (array == NULL) {
    return NULL;
  }
  array->length = length;
  array->array = notifications;
  array->current_id = 0;

  return array;
}

void free_notification_list(NotificationArray *notifications) {
  if (notifications != NULL) {
    for (size_t i = 0; i < notifications->length; i++) {
      free_notification(notifications->array[i]);
      notifications->array[i] = NULL;
    }

    free(notifications);
  }
}

NotificationArray *parse_python_notification_list(PyObject *list) {}
