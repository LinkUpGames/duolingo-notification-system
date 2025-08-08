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
  int probabilities_length;

  // The probabilities for each notification in this decision
  Notification **probabilities;

} Decision;

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
 * Free the memory for the notification struct that was allocated
 * @param notification The notification to free
 */
void free_notification(Notification *notification);

/**
 * Update the scores of each arm given the prior chosen arm
 * @param arms Array with each arm
 * @param length The length of the arms array
 * @param chosen_arm The name of the arm that was chosen by algorithm previously
 * @param selected Whether the arm by algorithm was selected by the user (1 or
 * 0);
 * @param alpha The learning rate
 * @param beta The temperature
 */
void update_scores(Arm **arms, int length, const char *chosen_arm, int selected,
                   float alpha, float beta);

/**
 * Create an empty list of arms of a specific size
 * @param length The number of arms to create
 */
Arm **create_arm_list(int size);

/**
 * Create a new arm struct with the following
 * @param name The name of the arm
 * @param reward The reward value of the arm at the moment
 */
Arm *create_arm(const char *name, float reward);

/**
 * Free an arm and it's memory after it's been created
 */
void free_arm(Arm *arm);

/**
 * Free the arm list
 * @param arms The arm list
 * @param length The length of the list
 */
void free_arm_list(Arm **arms, int length);

/**
 * Parse the python input provided and return an Arm array with the input
 * provided
 * @param args The arguments provided
 * @param list_size The size of the array
 */
Arm **parse_python_input(PyObject *input_list, Py_ssize_t list_size);

/**
 * Return the arms as a python list
 * @param arms The arms to be converted
 * @param length The length of the arms array
 */
PyObject *return_arms(Arm **arms, int length);

#endif
