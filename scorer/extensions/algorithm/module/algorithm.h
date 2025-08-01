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
