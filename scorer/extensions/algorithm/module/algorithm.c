#include "algorithm.h"
#include <Python.h>
#include <math.h>
#include <stdlib.h>
#include <string.h>

/**
 * Get the softmax values for the arms based on the temperature
 * NOTE: This returns a pointer that must be FREED!
 * @param arms The array of arms
 * @param length The length of the array
 * @param beta The temperature chosen
 */
float *softmax_probabilities(Arm **arms, int length, float beta) {
  float *values;
  float sum_exps = 0;
  float max_reward = 0;

  // Calculate the max score
  for (int i = 0; i < length; i++) {
    Arm *arm = arms[i];
    max_reward = arm->reward > max_reward ? arm->reward : max_reward;
  }

  // Calculate the exponential function
  values = (float *)calloc(length, sizeof(float));
  if (values == NULL) {
    return NULL;
  }

  for (int i = 0; i < length; i++) {
    Arm *arm = arms[i];
    float reward = arm->reward;

    float exp_p = exp((reward - max_reward) / beta);
    values[i] = exp_p;
    sum_exps += exp_p;
  }

  // Calculate softmax probability
  float *probabilities = (float *)calloc(length, sizeof(float));
  if (probabilities == NULL) {
    {
      return NULL;
    }
  }

  for (int i = 0; i < length; i++) {
    float e = values[i];

    probabilities[i] = e / sum_exps;
  }

  return probabilities;
}

void update_scores(Arm **arms, int length, const char *chosen_arm, int selected,
                   float alpha, float beta) {
  // Compute Softwmax probabilities
  float *probabilities = softmax_probabilities(arms, length, beta);

  // Compute baseline
  float baseline = 0.0;
  for (int i = 0; i < length; i++) {
    baseline += probabilities[i] * arms[i]->reward;
  }

  // Update the Scores of each arm
  for (int i = 0; i < length; i++) {
    Arm *arm = arms[i];
    if (strcmp(arm->name, chosen_arm) == 0) {
      // Update based on previous selection
      arm->reward = alpha * (selected - baseline) * (1 - probabilities[i]);
    } else {
      arm->reward = alpha * (selected - baseline) * probabilities[i];
    }
  }
}

Arm *create_arm(const char *name, float reward) {
  Arm *arm = malloc(sizeof(Arm));

  arm->name = strdup(name);
  arm->reward = reward;

  return arm;
}

Arm **create_arm_list(int size) {
  Arm **arms = (Arm **)calloc(size, sizeof(Arm *));

  return arms;
}

void free_arm(Arm *arm) {
  // Free the name
  free(arm->name);
  free(arm);

  arm = NULL;
}

void free_arm_list(Arm **arms, int length) {
  for (int i = 0; i < length; i++) {
    free_arm(arms[i]);
  }

  free(arms);

  arms = NULL;
}

Arm **parse_python_input(PyObject *input_list, Py_ssize_t list_size) {

  // Intialize the struct that we will be working with
  Arm **arms = create_arm_list(list_size);

  // Check for correct allocation
  if (arms == NULL) {
    PyErr_NoMemory();

    return NULL;
  }

  for (Py_ssize_t i = 0; i < list_size; i++) {
    PyObject *item = PyList_GetItem(input_list, i);

    // Check if the item is an object
    if (!PyDict_Check(item)) {
      PyErr_Format(PyExc_TypeError, "Element %zd is not an object", i);

      free_arm_list(arms, i);
      return NULL;
    }

    // Retrieve the key values
    PyObject *name_obj = PyDict_GetItemString(item, "name");
    PyObject *reward_obj = PyDict_GetItemString(item, "reward");

    // Check if the values exists
    if (!name_obj || !PyUnicode_Check(name_obj) || !reward_obj ||
        !PyFloat_Check(reward_obj)) {
      PyErr_Format(PyExc_ValueError,
                   "Invalid or missing keys in dict at index %zd", i);

      free_arm_list(arms, i);
      return NULL;
    }

    const char *name = PyUnicode_AsUTF8(name_obj);
    float reward = (float)PyFloat_AsDouble(reward_obj);

    Arm *arm = create_arm(name, reward);
    if (!arm) {
      PyErr_NoMemory();

      // Clean up
      free_arm_list(arms, i);
      return NULL;
    }

    arms[i] = arm;
  }

  return arms;
}

PyObject *return_arms(Arm **arms, int length) {
  PyObject *py_list = PyList_New(length);

  // Could not allocate memory for the pylist
  if (!py_list) {
    PyErr_NoMemory();
  }

  for (Py_ssize_t i = 0; i < length; i++) {
    PyObject *dict = PyDict_New();
    Arm *arm = arms[i];

    if (!dict) {
      PyErr_NoMemory();
      Py_DECREF(py_list);
    }

    int status;

    // Add the name
    PyObject *name_obj = PyUnicode_FromString(arm->name);
    status = PyDict_SetItemString(dict, "name", name_obj);

    if (!name_obj || status == -1) {
      PyErr_NoMemory();
      Py_XDECREF(name_obj);
      Py_DECREF(py_list);
      Py_DECREF(dict);
      PyErr_SetString(PyExc_RuntimeError, "Failed to set the 'name'");

      return NULL;
    }
    Py_DECREF(name_obj);

    // Add the reward
    PyObject *reward_obj = PyFloat_FromDouble((double)arm->reward);
    status = PyDict_SetItemString(dict, "reward", reward_obj);

    if (!reward_obj || status == -1) {
      PyErr_NoMemory();
      Py_DECREF(py_list);
      Py_DECREF(dict);
      PyErr_SetString(PyExc_RuntimeError, "Failed to set the 'reward'");

      return NULL;
    }
    Py_DECREF(reward_obj);

    // Insert the dict into the list
    PyList_SET_ITEM(py_list, i, dict);
  }

  return py_list;
}

Decision *create_decision(char const *id, char const *notification_id,
                          int selected) {
  // Allocate
  Decision *decision = (Decision *)malloc(sizeof(Decision));

  // Check for memory allocation
  if (decision == NULL) {
    return NULL;
  }

  // Copy Values
  decision->id = strdup(id);
  decision->notification_id = strdup(notification_id);
  decision->selected = selected;

  // Set to 0
  decision->probabilities = NULL;
  decision->probabilities_length = 0;

  return decision;
}

Decision **create_decision_list(long length) {
  Decision **decisions = (Decision **)calloc(length, sizeof(Decision *));

  return decisions;
}

void free_decision(Decision *decision) {
  if (decision->id != NULL) {
    free(decision->id);
    decision->id = NULL;
  }

  if (decision->notification_id != NULL) {
    free(decision->notification_id);
    decision->notification_id = NULL;
  }

  if (decision->probabilities != NULL) {
    int length = decision->probabilities_length;

    // Free all structs
    for (size_t i = 0; i < length; i++) {
      if (decision->probabilities[i] != NULL) {
        free_notification(decision->probabilities[i]);
        decision->probabilities[i] = NULL;
      }
    }

    // Free array
    free(decision->probabilities);
    decision->probabilities = NULL;
  }

  free(decision);
}

Notification *create_notification(char const *id, float score,
                                  float probability) {
  Notification *notification = (Notification *)malloc(sizeof(Notification));

  if (notification == NULL) {
    return NULL;
  }

  notification->id = strdup(id);
  notification->probability = probability;
  notification->score = score;

  return notification;
}

void free_notification(Notification *notification) {
  if (notification->id != NULL) {
    free(notification->id);
    notification->id = NULL;
  }

  free(notification);
}

Decision **parse_python_list(PyObject *list) {
  // Get the length of the list
  Py_ssize_t length = PyList_Size(list);

  // Initialize the decision array
  Decision **decisions = create_decision_list(length);
  if (decisions == NULL) {
    PyErr_NoMemory();
  }

  // Iterate through the list
  for (Py_ssize_t i = 0; i < length; i++) {
    PyObject *item = PyList_GetItem(list, i);

    // Check that the item is an object
    if (!PyDict_Check(item)) {
      PyErr_Format(PyExc_TypeError, "Element %zd is not an object", i);
      free_decision_list(decisions, length);

      return NULL;
    }
  }

  return decisions;
}

void free_decision_list(Decision **decisions, long length) {
  if (decisions != NULL) {
    for (size_t i = 0; i < length; i++) {
      // Free the decision
    }
  }
}
