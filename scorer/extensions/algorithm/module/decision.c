#include "decision.h"
#include "notification.h"
#include <Python.h>

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

DecisionArray *create_decision_list(long length) {
  Decision **decisions = (Decision **)calloc(length, sizeof(Decision *));

  if (decisions == NULL) {
    return NULL;
  }

  for (size_t i = 0; i < length; i++) {
    decisions[i] = NULL;
  }

  DecisionArray *array = (DecisionArray *)malloc(sizeof(DecisionArray));
  if (array == NULL) {
    free(decisions);

    return NULL;
  }

  array->array = decisions;
  array->current_id = 0;
  array->length = length;

  return array;
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
    free_notification_list(decision->probabilities);

    // Null pointer
    decision->probabilities = NULL;
  }

  free(decision);
}

void free_decision_list(DecisionArray *decisions) {
  if (decisions != NULL) {
    long length = decisions->length;

    for (size_t i = 0; i < length; i++) {
      // Free the decision
      free_decision(decisions->array[i]);
      decisions->array[i] = NULL;
    }

    free(decisions);
    decisions = NULL;
  }
}

DecisionArray *parse_python_decision_list(PyObject *list) {
  Py_ssize_t length = PyList_Size(list);
  DecisionArray *array = create_decision_list(length);

  if (array == NULL) {
    PyErr_NoMemory();

    return NULL;
  }

  for (Py_ssize_t i = 0; i < length; i++) {
    PyObject *item = PyList_GetItem(list, i);

    if (!PyDict_Check(item)) {
      free_decision_list(array);
      PyErr_Format(PyExc_ValueError, "Item %zd not a dict", i);

      return NULL;
    }

    // Create the decision struct
    // Decision *decision =
  }

  return array;
}

Decision *parse_python_decision_obj(PyObject *obj) {
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

  NotificationArray *notifications =
      parse_python_notification_dict(probabilities_obj);

  if (notifications == NULL) {
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

  return decision;
}

NotificationArray *parse_python_notification_dict(PyObject *dict) {
  if (!PyDict_Check(dict)) {
    PyErr_Format(PyExc_ValueError, "Probabilities is not a dictionary");

    return NULL;
  }

  PyObject *keys = PyDict_Keys(dict);
  if (!keys) {
    PyErr_Format(PyExc_ValueError, "Could not get keys");
    return NULL;
  }

  Py_ssize_t keys_length = PyList_Size(keys);

  NotificationArray *array = create_notification_list(keys_length);
  if (array == NULL) {
    PyErr_NoMemory();

    return NULL;
  }

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
      free_notification_list(array);
      PyErr_Format(PyExc_ValueError, "Could not parse notification id or "
                                     "probability from probability dict!");

      return NULL;
    }

    // Create Notification
    Notification *notification = create_notification(notification_id, 0, value);
    if (notification == NULL) {
      Py_DECREF(keys);
      free_notification_list(array);
      PyErr_NoMemory();

      return NULL;
    }

    // Add the notification reference
    array->array[i] = notification;
  }

  Py_DECREF(keys);

  return array;
}
