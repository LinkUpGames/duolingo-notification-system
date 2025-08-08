#include "algorithm.h"
#include <Python.h>
#include <math.h>
#include <stdlib.h>
#include <string.h>

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

  if (decisions == NULL) {
    return NULL;
  }

  for (size_t i = 0; i < length; i++) {
    decisions[i] = NULL;
  }

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

    // Free all probabilities struct pointers
    free_notification_list(decision->probabilities,
                           decision->probabilities_length);

    // Null pointer
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

void free_decision_list(Decision **decisions, long length) {
  if (decisions != NULL) {
    for (size_t i = 0; i < length; i++) {
      // Free the decision
      free_decision(decisions[i]);
      decisions[i] = NULL;
    }

    free(decisions);
    decisions = NULL;
  }
}

Notification **create_notification_list(long length) {
  Notification **notifications =
      (Notification **)calloc(length, sizeof(Notification *));

  if (notifications == NULL) {
    return NULL;
  }

  for (size_t i = 0; i < length; i++) {
    notifications[i] = NULL;
  }

  return notifications;
}

void free_notification_list(Notification **notifications, int length) {
  if (notifications != NULL) {
    for (size_t i = 0; i < length; i++) {
      free_notification(notifications[i]);
      notifications[i] = NULL;
    }

    free(notifications);
  }
}

Decision **parse_python_list(PyObject *list, Py_ssize_t length) {
  // Initialize the decision array
  Decision **decisions = create_decision_list(length);
  if (decisions == NULL) {
    PyErr_NoMemory();

    return NULL;
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

    // Create the decision struct
    Decision *decision = parse_python_obj(item);
    if (decision == NULL) {
      PyErr_Format(PyExc_ValueError, "Error with parsing decision object");

      return NULL;
    }

    decisions[i] = decision;
  }

  return decisions;
}

Decision *parse_python_obj(PyObject *obj) {
  // Retrieve field values
  PyObject *id_obj = PyDict_GetItemString(obj, "id");
  PyObject *notification_obj = PyDict_GetItemString(obj, "notification_id");
  PyObject *selected_obj = PyDict_GetItemString(obj, "selected");

  if (!id_obj || !notification_obj || !selected_obj) {
    PyErr_Format(PyExc_ValueError, "Invalid or missing keys in dict!");
    return NULL;
  }

  // Cast variables
  const char *id = PyUnicode_AsUTF8(id_obj);
  const char *notification_id = PyUnicode_AsUTF8(notification_obj);
  const int selected = selected_obj == Py_True ? 1 : 0;

  // Convert to notification array
  PyObject *probabilities_obj = PyDict_GetItemString(obj, "probabilities");
  if (!PyDict_Check(probabilities_obj)) {
    PyErr_Format(PyExc_ValueError, "probabilities is not a dictionary!");
    return NULL;
  }

  long notification_length = 0;
  Notification **notifications =
      parse_python_notification_dict(probabilities_obj, &notification_length);

  if (notifications == NULL) {
    PyErr_NoMemory();
    return NULL;
  }

  // Create decision struct
  Decision *decision = create_decision(id, notification_id, selected);
  if (decision == NULL) {
    PyErr_Format(PyExc_ValueError, "Could not allocate decision memory");
    PyErr_NoMemory();
    return NULL;
  }

  decision->probabilities = notifications;
  decision->probabilities_length = notification_length;

  return decision;
}

Notification **compute_scores(Decision **decisions, long length) {
  for (size_t i = 0; i < length; i++) {
    Decision *decision = decisions[i];

    int was_selected = decision->selected;
    char *selected_notification_id = decision->notification_id;

    long notification_length = decision->probabilities_length;

    // Set the weights when selected and when not selected
    for (size_t j = 0; j < notification_length; j++) {
      Notification *notification = decision->probabilities[j];
      double weight = 1 / notification->probability;
      double value = was_selected * weight;

      if (strcmp(notification->id, selected_notification_id) == 0) {
        //  This is the notificaiton that was selected
        notification->m_plus_count += value * weight;
        notification->m_plus_sum += weight;
      } else {
        // Not selected, reweight score
        notification->m_minus_count += value * weight;
        notification->m_minus_sum += weight;
      }
    }

    // Calculate m_plus and m_minus
    for (size_t j = 0; j < notification_length; j++) {
      Notification *notification = decision->probabilities[j];
      float m_plus = 0;
      float m_minus = 1;

      m_plus = notification->m_plus_count > 0
                   ? notification->m_plus_sum / notification->m_plus_count
                   : m_plus;

      m_minus = notification->m_minus_count > 0
                    ? notification->m_minus_sum / notification->m_minus_count
                    : m_minus;

      // Scores
      notification->m_plus = m_plus;
      notification->m_minus = m_minus;
      notification->score = (m_plus / m_minus) / m_minus;
    }
  }

  return NULL;
}

Notification **parse_python_notification_dict(PyObject *dict, long *length) {
  // Check if it's a dictionary
  if (!PyDict_Check(dict)) {
    PyErr_Format(PyExc_ValueError, "probabilities is not a dictionary");

    return NULL;
  }

  PyObject *keys = PyDict_Keys(dict); // Create a new list
  if (!keys) {
    PyErr_Format(PyExc_ValueError, "Could not get keys");
    return NULL;
  }
  Py_ssize_t keys_length = PyList_Size(keys);
  *length = keys_length;

  Notification **notifications = create_notification_list(keys_length);
  if (notifications == NULL) {
    PyErr_NoMemory();
    Py_DECREF(keys);
    return NULL;
  }

  // Iterate through the keys
  for (Py_ssize_t i = 0; i < keys_length; i++) {
    // Get the notification id
    PyObject *key_obj = PyList_GetItem(keys, i);

    // Get probability value
    PyObject *value_obj = PyDict_GetItem(dict, key_obj);

    const char *notification_id = PyUnicode_AsUTF8(key_obj);
    const float value = PyFloat_AsDouble(value_obj);

    // Error Check
    if (!notification_id || !value) {
      Py_DECREF(keys); // remove reference of newly created list
      free_notification_list(notifications, keys_length);
      PyErr_Format(PyExc_ValueError, "Could not parse notification id or "
                                     "probability from probability dict!");

      return NULL;
    }

    // Create Notification
    Notification *notification = create_notification(notification_id, 0, value);
    if (notification == NULL) {
      Py_DECREF(keys);
      free_notification_list(notifications, keys_length);
      PyErr_NoMemory();

      return NULL;
    }

    // Add the notification reference
    notifications[i] = notification;
  }

  Py_DECREF(keys);

  return notifications;
}
