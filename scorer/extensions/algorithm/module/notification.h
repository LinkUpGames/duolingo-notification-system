#ifndef NOTIFICATION_H
#define NOTIFICATION_H

#include "map.h"
#include <Python.h>

// Notification The notification structure used for
typedef struct Notification {
  // Id of the notification
  char *id;

  // The score for the notification
  float score;

  // The probability that the notification was selected
  float probability;

  // rewarded weight sum when selected
  float m_plus_sum;

  // weight sum when selcted
  float m_plus_count;

  // reward weight sum when not selected
  float m_minus_sum;

  // weight sum when not selected
  float m_minus_count;

  // The average weight when the notification is selected
  float m_plus;

  // The average weight when the notification was not selected;
  float m_minus;

} Notification;

// The notification array that holds the information for a notification
typedef struct NotificationArray {
  Notification **array;

  long length;

  size_t current_id;
} NotificationArray;

/**
 * Create a notification struct with the following values
 * @param id The notification id
 * @param score The score for the notification
 * @param probability The probability that this notification is selected
 */
Notification *create_notification(char const *id, float score,
                                  float probability);

/**
 * Create a notificaiton list of length
 * @param length The length of the array
 */
NotificationArray *create_notification_list(long length);

/**
 * Free the notification list
 * @param length The length of the array
 */
void free_notification_list(NotificationArray *array);

/**
 * Free the memory for the notification struct that was allocated
 * @param notification The notification to free
 */
void free_notification(Notification *notification);

/**
 * Parse a python list object and return a notification array struct
 * @param list The list
 */
NotificationArray *parse_python_notification_list(PyObject *list);

/**
 * The notifications as a dict
 * @parma dict the dictionary object
 */
NotificationArray *parse_python_notification_dict(PyObject *dict);

/**
 * The notification map from list
 * NOTE: that this only create a reference to key and value, so you must free
 * the source
 * @param notifications The notificatio map
 */
hashmap *create_notification_map_from_list(NotificationArray *notifications);

/**
 * Return a python object
 * @param notifications The notification list
 */
PyObject *notification_list_to_python_list(NotificationArray *list);

#endif
