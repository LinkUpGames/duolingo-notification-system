#ifndef ALGORITHM_H
#define ALGORITHM_H

#include <Python.h>

// An arm that corresponds to a notification to be there
typedef struct Arm {
  // The name of the arm
  // NOTE: You must free this once you are done
  char *name;

  // The reward [0,1] for the arm
  float reward;
} Arm;

// Notification The notification structure used for
typedef struct Notification {
  // Id of the notification
  char *id;

  // The score for the notification
  float score;

  // The probability that the notification was selected
  float probability;

} Notification;

// The decision log
typedef struct Decision {
  // The log id
  char *id;

  // The notification id that was decided
  char *notification_id;

  // Whether the notification was selected or not
  int selected;

  // The length of the probabilities
  long probabilities_length;

  // The probabilities for each notification in this decision
  Notification **probabilities;

} Decision;

/**
 * Create a list of decisions given the length to initialize
 * @param length The length of the array
 */
Decision **create_decision_list(long length);

/**
 * Free the array of decisions
 * @param decisions The decisions array
 * @param length The length of the array
 */
void free_decision_list(Decision **decisions, long length);

/**
 * Create a Decision struct with an empty array (no array allocation)
 * @param id The id for the decision
 * @param notification_id The notification id
 * @param selected Whether the notification was selectged by the user or not (1
 * or 0)
 */
Decision *create_decision(char const *id, char const *notification_id,
                          int selected);

/**
 * Free the decision struct and any allocated components
 * @param decision The decision struct to free
 */
void free_decision(Decision *decision);

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
Notification **create_notification_list(long length);

/**
 * Free the notification list
 * @param length The length of the array
 */
void free_notification_list(Notification **notifications, int length);

/**
 * Free the memory for the notification struct that was allocated
 * @param notification The notification to free
 */
void free_notification(Notification *notification);

/**
 * Parse a python list and return an array with Decision structs
 * @param list The python list
 */
Decision **parse_python_list(PyObject *list, Py_ssize_t length);

/**
 * Parse a python dict in the form of an object and return the decision log from
 * it
 * @param obj The python object
 */
Decision *parse_python_obj(PyObject *obj);

#endif
