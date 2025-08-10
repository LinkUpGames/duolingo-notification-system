#ifndef DECISION_H
#define DECISION_H

#include "notification.h"
#include <Python.h>

// The decision log
typedef struct Decision {
  // The log id
  char *id;

  // The notification id that was decided
  char *notification_id;

  // Whether the notification was selected or not
  int selected;

  // The probabilities for each notification in this decision
  NotificationArray *probabilities;

} Decision;

// DecisionArray The decision array
typedef struct DecisionArray {
  Decision **array;

  long length;

  size_t current_id;

} DecisionArray;

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
 * Create a list of decisions given the length to initialize
 * @param length The length of the array
 */
DecisionArray *create_decision_list(long length);

/**
 * Free the array of decisions
 * @param decisions The decisions array
 * @param length The length of the array
 */
void free_decision_list(DecisionArray *decisions);

/**
 * Parse a python list and return an array with Decision structs
 * @param list The python list
 */
DecisionArray *parse_python_decision_list(PyObject *list);

/**
 * Parse the decision object inside of the list
 * @param obj The obj to parse
 */
Decision *parse_python_decision_obj(PyObject *obj);

#endif
